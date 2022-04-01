---
title: Java从零开始（143）非阻塞 Java NIO Channel 体系介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 非阻塞 Java NIO Channel 体系介绍


## 1. 前言

java.nio.channels.SocketChannel 和 java.nio.channels.ServerSocketChannel 是编写非阻塞 Java TCP Soccket 程序的重要模块。然而，Channel 是 Java NIO 非常重要的概念，在 java.nio.channels 抽象了完整的、关于 Channel 的接口（Interface） 层次结构。

Channel 表示一个和硬件设备、磁盘文件、网络 Socket 等 I/O 设备、或者组件之间**连接**。Channel 的状态要么是**打开的**，要么是**关闭的**。当 Channel 处于**打开**状态，可以从 Channel 中读取数据并且保存到 ByteBuffer 中，也可以将 ByteBuffer 中的数据写到 Channel 中。当 Channel 处于**关闭**状态，对 Channel 执行读、写操作，会产生 ClosedChannelException 异常。

## 2. Channel 类层次结构

java.nio.channels 包抽象了完整的 Channel 层次结构，如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyt4mqij60l80d2gmu02)

### 2.1 Java NIO Channel 的祖先类

java.nio.channels.Channel 是一个 Java 接口，是整个 Channel 家族的祖先，声明的接口如下：

```java
// 获取 Channel 的状态
public boolean isOpen();
// 关闭 channel
public void close() throws IOException;
```

Java NIO Channel 具有以下特性：

* Java NIO 支持面向字节流的数据读写方式，数据从 ByteBuffer 读取后写到 Channel 中，或者从 Channel 中读取后写入 ByteBuffer 中。
* Java NIO Channel 支持字节流读写双向操作，一个处于 Open 状态的 Channel，既可以进行读操作、也可以进行写操作。
* Java NIO Channel 支持阻塞和非阻塞两种模式。
* Java NIO Channel 是线程安全的。

### 2.2 Java NIO 中其他 Channel 接口

在 Java NIO Channel 体系中，对于 Socket 的抽象、数据读、数据写、数组形式的多缓冲读、数组形式的多缓冲写等功能分别进行了抽象，每一个功能都对应一个 Java 接口，下来我们分别做一个说明。

java.nio.channels.ReadableByteChannel 是一个 Java Interface，是对 Channel 读操作的抽象，声明的接口如下：

```java
public int read(ByteBuffer dst) throws IOException;
```

Channel 的 read 方法是从 I/O 设备读取数据，保存在 ByteBuffer 中，为此调用者必须提供 ByteBuffer 用以保存数据。返回值是读取的字节数、0、或者 -1。如果是阻塞式 Channel，read 至少返回 1 或者 -1；如果是非阻塞式 Chanel，read 可能会返回 0。

java.nio.channels.WritableByteChannel 是一个 Java Interface，是对 Channel 写操作的抽象，声明的接口如下：

```java
public int write(ByteBuffer src) throws IOException;
```

Channel 的 write 方法是从 ByteBuffer 读取数据，写入 I/O 设备中，为此调用者必须将要写出去的数据保存到 ByteBuffer 中。返回值是写入的字节数、0、或者 -1。如果是阻塞式 Channel，write 返回请求写入的字节数 或者 -1；如果是非阻塞式 write 可能会返回 0。

java.nio.channels.GatheringByteChannel 是一个 Java Interface，是对 Channel 写操作的抽象，可以写入一个数组，支持对多个 ByteBuffer 的写入，声明的接口如下：

```java
public long write(ByteBuffer[] srcs, int offset, int length)
        throws IOException;
public long write(ByteBuffer[] srcs) throws IOException;
```

java.nio.channels.ScatteringByteChannel 是一个 Java Interface，是对 Channel 读操作的抽象，可以读入一个数组，支持对多个 ByteBuffer 的读入，声明的接口如下：

```java
public long read(ByteBuffer[] dsts, int offset, int length)
        throws IOException;
public long read(ByteBuffer[] dsts) throws IOException;
```

java.nio.channels.NetworkChannel 是一个 Java Interface，表示一个 Socket。实现此接口的 Channel 表示对 Socket 的封装。

java.nio.channels.SelectableChannel 是一个 Java Interface，用于和 java.nio.channels.Selector 集成。声明的主要接口如下：

```java
public abstract SelectableChannel configureBlocking(boolean block)
        throws IOException;
public final SelectionKey register(Selector sel, int ops)
        throws ClosedChannelException
```

实现了 SelectableChannel 接口的类，可以将 I/O 操作设置为非阻塞模式，同时可以注册到 Selector，通过 I/O 多路复用机制监听事件。

java.nio.channels.ByteChannel 也是一个 Java 接口，只是实现了 java.nio.channels.ReadableByteChannel 和 java.nio.channels.WritableByteChannel 两个接口。ByteChannel 本身没有声明任何接口，实现了读写聚合的功能。

## 3. Channel 实现类

在 Channel 的类的层次结构图中，我们画出四个常用的实现类如下：

* FileChannel

文件 Channel 类是对磁盘文件的抽象，可以读写磁盘文件数据。需要通过 FileInputStream 的 getChannel 方法创建 FileChannel 的对象，你不可以直接创建 FileChannel 的对象。FileChannel 对象的创建方法如下：

```java
FileInputStream inFile = new FileInputStream("D:\\fileChannelTest.txt");
ReadableByteChannel fileChannel = inFile.getChannel();
```

* DatagramChannel

数据报 Channel 是用于抽象 UDP Socket，可以将 UDP 数据的读写集成到 Selector 机制中。DatagramChannel 对象的创建方法：

```java
DatagramChannel ch = DatagramChannel.open();
```

* SocketChannel 是对 TCP 的抽象，用于读写 TCP 数据，用于 TCP 客户端和服务器端

* ServerSocketChannel 是对 TCP 监听 Socket 的抽象，用于 TCP 服务器的创建。

## 4. 总结

本节重点是对 Java NIO Channel 类的体系结构进行了介绍。通过 channel 类的体系结构图，我们可以看出，主要是对数据读、数据写、数组形式的多缓冲读、数组形式的多缓冲写、以及和**多路复用机制 Selector**的集成等功能的抽象。每一个接口都对应了相应功能的实现，将这些接口进行一个组合，就是一个具体的 I/O 功能的实现，这一具体的组合是通过具体实现类来体现的。比如 ，FileChannel 是对磁盘文件的抽象，DatagramChannel 是对 UDP Socket 的抽象，Socket Channel 是对 TCP Socket 的抽象，ServerSocketChannel 是对 TCP Server 监听逻辑的抽象。

通过逐个分析每一个接口所提供的能力，我们就熟悉了完整的 Java NIO Channel 体系。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
