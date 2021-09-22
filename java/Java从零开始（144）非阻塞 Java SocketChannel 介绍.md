---
title: Java从零开始（144）非阻塞 Java SocketChannel 介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 非阻塞 Java SocketChannel 介绍


## 1. 前言

前一小节介绍了 Java NIO Channel 体系结构，其中 java.nio.channels.SocketChannel 和 java.nio.channels.ServerSocketChannel 是编写非阻塞 Java TCP Soccket 程序的重要模块。SocketChanel 用在 TCP 的客户端和服务器端，用于数据的读写。SocketChannel 表示一个 Endpoint，并且和远端的 SocketChannel 建立了连接。ServerSocketChannel 用于创建 TCP 服务器，用于监听新的 TCP 连接请求。本小节对 SocketChannel 和 ServerSocketChannel 的基本功能做一个初步的分析。

## 2. SocketChannel 介绍

SocketChannel 的基本功能就是用于 TCP 连接的建立，数据读写，TCP 连接关闭。

### 2.1 SocketChannel 实例创建

SocketChannel 并没有提供 public 构造方法，创建 SocketChannel 的实例需要通过工厂方法 open 实现。open 声明如下：

```java
public static SocketChannel open() throws IOException
public static SocketChannel open(SocketAddress remote) throws IOException
```

SocketChannel 的不带参数的 open 方法只是创建实例，不会 connect 服务器，带参的 open 方法会 connect 远端服务器。如果是创建阻塞式 SocketChannel，可以通过 open 方法传入远端服务器的地址；如果是创建非阻塞式 SocketChannel 需要调用不带参数的 open 方法，然后再调用 connect 方法连接远端服务器。

### 2.2 SocketChannel 的 connect 方法

SocketChannel 提供了专门 connect 方法，可以进行阻塞式、非阻塞式连接。声明如下：

```java
public abstract boolean connect(SocketAddress remote) throws IOException;
public abstract boolean finishConnect() throws IOException
```

如果 SocketChannel 是阻塞模式，调用它的 connect 方法连接远端服务器，connect 方法会执行阻塞式调用，直到连接成功或者失败返回。

如果 SocketChannel 是非阻塞模式，调用它的 connect 方法连接远端服务器，connect 方法会执行非阻塞式调用。connect 发出连接请求后，不管 TCP 连接是否成功，都会马上返回。必须调用 finishConnect 方法，检查连接是否成功，如果连接成功 finishConnect，返回 true；如果连接没有成功，finishConnect 返回 false。

### 2.3 SocketChannel 的数据读取

SocketChannel 提供了读取单片数据的方法，声明如下：

```java
public abstract int read(ByteBuffer dst) throws IOException
```

其实，单片数据的 read 方法是重写了 java.nio.channels.ReadableByteChannel 中的 read 方法。 read 方法是从 I/O 设备读取数据，保存在 ByteBuffer 中，为此调用者必须提供 ByteBuffer 用以保存数据。返回值是读取的字节数、0、或者 -1。如果是阻塞式 Channel，read 至少返回 1 或者 -1；如果是非阻塞式 Chanel，read 可能会返回 0。

SocketChannel 提供了读取多片数据的方法，声明如下：

```java
public final long read(ByteBuffer[] dsts) throws IOException
public final long read(ByteBuffer[] dsts, int offset, int length) throws IOException
```

其实，多片数据的 read 方法是重写了 java.nio.channels.ScatteringByteChannel 中的 read 方法。多片数据 read 方法的返回值和单片数据 read 方法的返回值具有相同的含义。多片数据的 read 方法，其实是将 TCP 字节流保存在不同的 ByteBuffer 中，这些 ByteBuffer 是不同的内存块，通常叫做 Scatter 机制。

### 2.4 SocketChannel 的数据写入

SocketChannel 提供了写入单片数据的方法，声明如下：

```java
public abstract int write(ByteBuffer src) throws IOException
```

其实，单片数据的 write 方法是重写了 java.nio.channels.WritableByteChannel 中的 write 方法。write 方法是从 ByteBuffer 读取数据，写入 I/O 设备中，为此调用者必须将要写出去的数据保存到 ByteBuffer 中。返回值是写入的字节数、0、或者 -1。如果是阻塞式 Channel，write 返回请求写入的字节数 或者 -1；如果是非阻塞式 write 可能会返回 0。

SocketChannel 提供了写入多片数据的方法，声明如下：

```java
public final long write(ByteBuffer[] dsts) throws IOException
public final long write(ByteBuffer[] dsts, int offset, int length) throws IOException
```

多片数据的 write 方法是重写了 java.nio.channels.GatheringByteChannel 中的 write 方法。多片数据 write 方法的返回值和单片数据 write 方法的返回值具有相同的含义。多片数据的 write 方法，其实是将保存在不同的 ByteBuffer 中字节流写入 TCP Socket，这些 ByteBuffer 是不同的内存块，通常叫做 Gathering 机制。

### 2.5 SocketChannel 的关闭

SocketChannel 覆写了 java.nio.channels.Channel 中的 close 方法，完成 TCP 连接的关闭。方法声明如下：

```java
public void close() throws IOException
```

## 3. ServerSocketChannel 介绍

ServerSocketChannel 用于 TCP 创建服务器，监听客户端的连接请求。ServerSocketChannel 也没有提供 public 的构造方法，创建 ServerSocketChannel 的实例，需要调用它的工厂方法 open。声明如下：

```java
public static ServerSocketChannel open() throws IOException
```

ServerSocketChannel 提供了 bind 方法，用于绑定监听的 IP 地址和端口。声明如下：

```java
public final ServerSocketChannel bind(SocketAddress local) throws IOException
public abstract ServerSocketChannel bind(SocketAddress local, int backlog) throws IOException;
```

ServerSocketChannel 提供了 accept 方法，接收新的客户端请求，返回值是 SocketChannel 类型的对象，表示一个新的 TCP 连接。

```java
public abstract SocketChannel accept() throws IOException;
```

通过 ServerSocketChannel 创建一个 TCP 服务器，一般需要如下几个步骤：

```java
try {
    Selector selector = Selector.open();
    ServerSocketChannel serverChannel = ServerSocketChannel.open();
    // 绑定监听的 socket 地址,监听 any_addr
    serverChannel.socket().bind(new InetSocketAddress(PORT));
    // 设置 SO_REUSEADDR 选项，作为服务器，这是基本的要求
    serverChannel.socket().setReuseAddress(true);
    // 设置非阻塞模式，作为服务器，也是基本要求
    serverChannel.configureBlocking(false);
    // 注册 accept 事件
    serverChannel.register(selector, SelectionKey.OP_ACCEPT, serverChannel);
} catch (IOException e) {
    e.printStackTrace();
}
```

后续的逻辑就是通过 Selector 的 select 方法监听 I/O 事件。然后通过 SocketChannel 的 read 和 write 方法进行数据读写。

## 4. 小结

本小节重点是介绍了 SocketChannel 和 ServerSocketChannel 两个类提供的重点 API。关于这两个类的具体应用示例，已经在前面的小节展示，可以参考。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
