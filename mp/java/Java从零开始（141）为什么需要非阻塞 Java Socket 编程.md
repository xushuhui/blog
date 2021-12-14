---
title: Java从零开始（141）为什么需要非阻塞 Java Socket 编程
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 为什么需要非阻塞 Java Socket 编程


## 1. 前言

前面小节介绍的都是**阻塞式** Socket 编程。比如，我们最早编写的 TCP Client/Server 示例程序，客户端定时发送消息，服务器只是做一个响应。由于只是服务一个客户端，所以通过**阻塞式** Socket 编程勉强能满足需求。

在 Java 服务器多线程一节，我们介绍了**每线程模型**和**线程池模型**。通过这两种多线程模型，服务器可以同时和多个客户端完成通信。对于**每线程模型**来说，其核心是为每一个新连接创建一个独立的子线程，由这个独立的子线程负责和客户端完成数据收发。每线程模型的特点是服务器上的线程和客户端是一对一的。这种解决方案是有很大弊端的，因为系统能够创建的线程数量是有限的，是无法支撑高并发场景的。对于**线程池模型**来说，尽管限制了创建的线程的总数，但是由于是阻塞式 Socket，一旦某个线程被分配和客户端通信，就只能和此客户端通信，所以在容量上有限制。

对于高并发的应用场景来说，还是得通过**非阻塞式** Socket 编程来解决。

首先我们了解一下**阻塞式**和**非阻塞式**的区别。

## 2. 阻塞式与非阻塞式模型

我们以 Linux 系统为例，介绍**阻塞式**与**非阻塞式**的概念。Linux 程序的执行模式分为**用户态**和**内核态**，应用程序逻辑运行在**用户态**，访问系统资源的逻辑运行在**内核态**。其实现代操作系统都是这种模式。

当程序的执行逻辑从**用户态**切换到**内核态**时，会引发上下文的切换，会涉及到数据从**用户态**到**内核态**，或者是从**内核态**到**用户态**拷贝的问题。这时，系统 API 会提供**阻塞式**和**非阻塞式**两种调用方式。比如，我们调用 recv 函数接收 Socket 数据，recv 函数可以选择**阻塞式**或者是**非阻塞式**调用模式，不同的模式，编程风格是完全不同。假如 Socket 的接收缓冲区没有准备好要接收的数据，如果选择**阻塞式**调用，那么应用线程会被阻塞在 recv 调用上，不能继续执行，线程会处于等待状态，直到系统准备好数据；如果选择**非阻塞式**调用，那么应用线程不会被阻塞，recv 函数会立即返回。当系统准备好数据以后，会触发一个**读事件**，这就要求我们必须通过某种机制监听**读事件**，一般都是通过 I/O 多路复用机制来解决。

我们通过两张图来感受一下**阻塞式**和**非阻塞式**的差异。

阻塞式：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyscdcgj60ua0gowiu02)

非阻塞式：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyspebhj60w10dw43202)

从以上两张图可以看出，如果 read 函数采用**阻塞式**调用 ，当内核没有准备好的数据时，应用线程会被阻塞到 read 调用上，进入等待状态，直到有数据可以读取才返回。如果 read 函数采用**非阻塞式**调用，当内核没有准备好数据时，read 函数会返回 EAGAIN，线程不会被阻塞。当系统准备好数据以后，会触发一个**读事件**。

对于逻辑比较简单的场景，比如逻辑简单的客户端程序，可以采用**阻塞式**编程模型，这样实现简单，容易理解。对于逻辑比较复杂的场景，比如高性能服务器，必须采用**非阻塞式**编程模型，而且要配合 **I/O 多路复用**机制。

下来我们就介绍一下如何进行**非阻塞式** Socket 编程。

## 3. Java 非阻塞式 Socket 编程

介绍 Java 非阻塞式 Socket 编程，就得介绍 Java NIO。Java NIO 是 Java New IO API，有时也解释为 Java Non-blocking IO。通过 Java NIO 可以实现 Java 非阻塞 Socket 编程。

Java NIO 是 Java 1.4 支持的，它将 Socket 数据流抽象为一个 Channel（管道），Socket 数据读写是通过 Channel

实现的，并且提供了 Buffer 机制，提高数据读写的性能。Java NIO 通常用来编写高性能 Java 服务器程序。在 Java 1.7 以后，Java NIO 对磁盘文件处理得到了增强，可以将 Socket I/O 和 文件 I/O 融合在 Java NIO 中。

Java NIO 提供的新的类结构如下：

|类名称|功能说明|
|------|--------|
|ServerSocketChannel|表示服务端 TCP Socket 的监听 Channel。ServerSocketChannel 提供的工厂方法 open，用于创建它的实例；同时它提供了 accept 方法用于在服务器中接收新的客户端连接请求，返回值是 SocketChannel 类的实例。|
|SocketChannel      |SocketChannel 表示一个 TCP 通信 Channel，可以通过它的 open 方法创建，也可以通过 ServerSocketChannel 的 accept 方法创建。                                                                        |
|Selector           |Java I/O 事件多路复用机制，用于同时监听多个 Channel 的读、写、监听事件                                                                                                                          |
|SelectionKey       |用于表示具体的事件对象                                                                                                                                                                          |
|ByteBuffer         |通过 SocketChannel 进行数据读写，依赖 ByteBuffer                                                                                                                                                |

ServerSocketChannel 和 SocketChannel 同时支持**阻塞式**和**非阻塞式**，默认是**阻塞式**。可以通过如下的方法，打开**非阻塞式**。

```java
// 配置监听 ServerSocketChannel 为非阻塞模式
ServerSocketChannel serverChannel = ServerSocketChannel.open();
serverChannel.configureBlocking(false);

// 配置服务器新建立的 SocketChannel 为非阻塞模式
SocketChannel newSock = serverChannel.accept();
newSock.configureBlocking(false);
```

```java
SocketAddress serverAddr = new InetSocketAddress("127.0.0.1", PORT);
SocketChannel sock = SocketChannel.open(serverAddr);
// 配置客户端 SocketChannel 为非阻塞
sock.configureBlocking(false);
```

## 4. 小结

阻塞式 Socket 编程，程序结构简单，容易编写，容易理解。但是由于阻塞式 Socket 编程，在调用 recv、send 读写数据的时候，会阻塞线程，所以只能适应简单的应用场景。对于编写高性能服务器来说，必须采用非阻塞式 Socket 编程。但是非阻塞式 Socket 编程，程序结构要复杂很多，并且不容易理解，要想编写健壮、稳定的程序不是一件容易的事情。

编写 Java 非阻塞 Socket 程序，需要采用 Java NIO API，这些 API 的具体功能、具体用法，在后面小节介绍。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
