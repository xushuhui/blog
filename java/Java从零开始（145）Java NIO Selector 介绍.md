---
title: Java从零开始（145）Java NIO Selector 介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java NIO Selector 介绍


## 1. 前言

前面小节介绍 Java TCP Socket 编程时，我们遇到了“创建的 TCP server 不能同时接收多个客户端请求”的问题。我们的解决方案是采用多线程模型，即**每线程模型**或者是**线程池模型**。然而多线程模型会创建大量的线程，消耗大量系统资源。后来进入了 Java NIO 编程的学习，我们说编写高性能服务器必须采用**非阻塞式** Socket 编程。然而，通过 Java NIO 的编程示例可以发现：相比**阻塞式 Socket 编程**，**非阻塞式 Socket 编程**的难度大了一个数量级。我们需要应用好 I/O 多路事件处理机制，需要处理好数据收发的各种情况，而 I/O 事件多路复用机制是整个非阻塞 Socket 编程的核心。

其实，**I/O 多路复用机制**（I/O Multiplex) 最早是由操作系统提供的，有一套专用的系统 API。目前主流操作系统提供的 **I/O 多路复用** API 如下：

* select，是通用机制，Windows、Unix-like 系统都支持。
* poll, 是 UNIX-like 系统支持。
* devpoll，是 SUN Solaris 系统支持。当然，SUN 公司已经不存在了。
* epoll, 是 Linux 系统支持的主流机制。
* Kqueue，是 freebsd 内核支持的机制，Mac OS、IOS 系统也支持。
* IOCP，是 Windows 系统支持的机制。

对于 Java 来说，也有自己的 **I/O 多路复用机制**，那就是 **Java NIO Selector**。

## 2. Java NIO Selector 工作原理

Java NIO 四个核心的组件分别是 Selector、SocketChannel、ServerSocketChannel、SelectionKey。Selector 是 I/O 事件反应器，是动力源。SocketChannel、ServerSocketChannel、SelectionKey 都是功能组件，它们之间互相配合，如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnytinvnj60x30io0xo02)

* 首先创建一个 Selector 对象，然后调用它的 select 方法，进入事件等待状态。
* 对于服务器来说，需要创建 ServerSocketChannel 对象，然后调用它的 register 方法，将 SelectionKey.OP_ACCEPT 事件注册到 Selector，准备监听新的客户端连接。
* 如果 Selector 监听到新的客户端连接请求，SelectionKey.OP_ACCEPT 事件就会产生。调用 ServerSocketChannel 的 accept 方法，返回一个 SocketChannel 对象，需要将 SocketChannel 的 SelectionKey.OP_READ 事件注册到 Selector。

> * 在上面两步中， ServerSocketChannel 和 SocketChannel 都提供了 register 方法，返回值是 SelectionKey。SelectionKey 中绑定了上下文信息。

* 如果 Selector 监听到 I/O 事件，它的 select 方法就会返回。可以调用 Selector 的 selectedKeys 方法，返回一个 SelectionKey 数组，包含了所有产生了 I/O 事件的 SocketChannel。遍历这个数组，逐个处理相应的 I/O 事件。

## 3. Java Selector API 介绍

### 3.1 创建实例

Selector 只声明了一个 protected 构造方法，构造 Selector 实例只能通过它的工厂方法 open，声明如下：

```java
public static Selector open() throws IOException
```

### 3.2 注册事件

Selector 接口并没有声明 register 方法，而是通过它的抽象实现类提供了一个 abstract protected 方法，也没有对外暴露。声明如下：

```java
protected abstract SelectionKey register(AbstractSelectableChannel ch,int ops, Object att);
```

Selector 的 register 机制最终委派给了 AbstractSelectableChannel 类。为此，我们要想将 Channel 注册到 Selector，需要调用 AbstractSelectableChannel 的 register 方法。声明如下：

```java
public final SelectionKey register(Selector sel, int ops, Object att) throws ClosedChannelException
```

参数说明：

* sel 是预先创建的 Selector 对象。
* ops 表示需要注册的具体事件。支持的事件类型如下：

```java
- SelectionKey.OP_ACCEPT   表示监听客户端的连接，用于服务器
- SelectionKey.OP_CONNECT  表示非阻塞式客户端连接过程，用于客户端
- SelectionKey.OP_READ     表示监听读事件
- SelectionKey.OP_WRITE    表示监听写事件
```

* att 用于保存上下文对象。

### 3.3 监听事件

Selector 提供了 select 方法用于监听 I/O 事件，声明如下：

```java
public abstract int select() throws IOException
public abstract int select(long timeout) throws IOException
```

当没有 I/O 事件产生时，调用 select 方法的线程会被阻塞。如果你调用无参 select 方法，线程进入等待状态，直到有 I/O 事件发生才返回。如果你调用包含了 timeout 参数的 select 方法，线程会在 timeout 超时，或者是有 I/O 事件发生返回。select 的返回值表示产生了 I/O 事件的 SelectionKey 的个数。

### 3.4 遍历事件

当有 I/O 事件发生，Selector 的 select 方法会返回。可以通过 Selector 的 selectedKeys 方法，获取所有产生了 I/O 事件的 SelectionKey。声明如下：

```java
 public abstract Set<SelectionKey> selectedKeys()
```

方法的返回值是一个 SelectionKey 类型的集合，我们需要遍历此集合，逐个处理。遍历的方法如下：

```java
Set<SelectionKey> selectedKeys = selector.selectedKeys();
Iterator<SelectionKey> keyIterator = selectedKeys.iterator();
while (keyIterator.hasNext()) {
    SelectionKey key = keyIterator.next();
    if (key != null) {
        if (key.isAcceptable()) {
            // ServerSocketChannel 接收了一个新连接
        } else if (key.isConnectable()) {
            // 表示一个新连接建立
        } else if (key.isReadable()) {
            // Channel 有准备好的数据，可以读取
        } else if (key.isWritable()) {
            // Channel 有空闲的 Buffer，可以写入数据
        }
    }
    keyIterator.remove();
}
```

### 3.5 SelectionKey 介绍

SelectionKey 是由 AbstractSelectableChannel 的 register 方法返回的，主要包含一个事件类型和上下文对象。SelectionKey 提供了一组方法，用以识别 I/O 事件类型。声明如下：

```java
public final boolean isAcceptable()
public final boolean isConnectable()
public final boolean isReadable()
public final boolean isWritable()
```

可以通过 SelectionKey 的 channel 方法，获取关联的 Channel，声明如下：

```java
public abstract SelectableChannel channel()
```

可以通过 SelectionKey 的 attachment 方法，获取关联的上下文对象。

```java
public final Object attachment()
```

SelectionKey 的各个方法相对简单，容易理解，我们在前面小节多次提到，不再赘述。

## 4. 总结

本小结主要是介绍 Java NIO Selector 机制的工作原理。关于 Selector 机制，不仅 Java 支持，各大操作系统都有支持，是编写高性能服务器的利器。一般在依赖倒转模型中，充当动力源、反应器的角色。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
