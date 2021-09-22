---
title: Java从零开始（142）非阻塞 Java Socket 编程示例
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 非阻塞 Java Socket 编程示例


## 1. 前言

在上一节我们介绍了通过 Java NIO 可以编写**非阻塞式** Socket 程序。Java NIO 新引入的几个类是：

ServerSocketChannel、SocketChannel、Selector、SelectionKey、ByteBuffer。其实 Java NIO 既可以编写**阻塞式** Socket 程序，也可以编写**非阻塞式** Socket 程序。本节将会通过一个简单的 Java TCP 客户端、服务器程序演示 Java NIO 编写 Socket 程序的基本步骤。客户端是通过 Java NIO **阻塞式**实现，服务器是通过 Java NIO **非阻塞式**实现。

## 2. 程序功能及协议结构

实现一个简单的服务器“时钟同步”的功能，此功能在音视频实时通信系统中应用非常广泛。不过通常是将数据源的时钟随同数据包一起携带，并不会专门收发包含时钟的数据包。本文为了演示 Java NIO 的使用方式，客户端和服务器之间实现一个交换“时间”的功能。

客户端每隔 1 秒向服务器发送一个请求对方时间的消息，此消息会携带客户端自己的时间；服务器收到客户端的请求以后，将对方的时间打印在控制台上，同时将自己的本地时间发送给客户端；客户端收到服务器的响应后，将服务器的时间打印在控制台上。客户端一共会执行 10 次。

由于是 TCP 程序，我们必须定义协议结构，至少要定义一个 length 字段，标识消息的长度。协议结构如下：

```java
+------------------------+
|   length               |
++++++++++++++++++++++++++
|    消息体               |
++++++++++++++++++++++++++
```

## 3. Java NIO 客户端实现步骤

* 首先创建目标服务器地址结构，这和之前介绍的一致。

```java
    SocketAddress serverAddr = new InetSocketAddress("127.0.0.1", PORT);
```

* 通过 SocketChannel 的 open 方法打开一个客户端 Socket，参数是 SocketAddress 类型的对象。

```java
SocketChannel sock = null;
sock = SocketChannel.open(serverAddr);
```

> * 注意：我们创建的是**阻塞式**客户端 Socket，如果需要创建**非阻塞式**客户端 Socket，需要调用

```java
sock.configureBlocking(false);
```

* 创建用于收发数据的 ByteBuffer，分配 1024 字节大小的 buffer。

```java
    ByteBuffer recvBuff = ByteBuffer.allocate(1024);
    ByteBuffer sendBuff = ByteBuffer.allocate(1024);
```

* 通过 SocketChannel 的 read 方法读取数据，读取的数据保存在 recvBuff 中，这是上面创建的 Bytebuffer 对象，长度是 1024 字节

```java
int rbytes = sock.read(recvBuff);
```

SocketChannel 的 read 方法返回值如果是大于 0，表示读取的字节数；返回值如果是 -1，表示数据读取结束。对于**非阻塞式** Socket，返回值可能是 0。

* 通过 SocketChannel 的 write 方法向 Socket 写入数据，需要发送的数据，需要提前写入 ByteBuffer 中。

```java
sock.write(sendBuff);
```

客户端完整代码：

```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.net.SocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SocketChannel;
import java.text.SimpleDateFormat;
import java.util.Date;

public class NonblockTCPClient {
    // 服务器监听的端口
    private final static int PORT = 9082;

    public static void main(String[] args) {
        SocketChannel sock = null;
        try {
            // 创建服务器地址结构
            SocketAddress serverAddr = new InetSocketAddress("127.0.0.1", PORT);
            sock = SocketChannel.open(serverAddr);

            ByteBuffer recvBuff = ByteBuffer.allocate(1024);
            ByteBuffer sendBuff = ByteBuffer.allocate(1024);

            int rquest_times = 10;

            while (true){
                SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
                String request = "Request time"+ df.format(new Date());
                sendBuff.putInt(request.length());
                sendBuff.put(request.getBytes());

                sendBuff.flip();
                sock.write(sendBuff);

                System.out.println("Send request to server");
                int bodyLen = -1;
                boolean isFlip = true;
                recvBuff.rewind();

                while (true){
                    int rbytes = sock.read(recvBuff);
                    if (rbytes == -1){
                        sock.close();
                        return;
                    }

                    if (bodyLen == -1){
                        if (rbytes < 4){
                            continue;
                        }
                        recvBuff.flip();

                        bodyLen = recvBuff.getInt();
                        isFlip =false;
                    }

                    if (isFlip ){
                        recvBuff.flip();
                    }
                    if (recvBuff.remaining() < bodyLen){
                        recvBuff.compact();
                        continue;
                    }

                    byte[] body = new byte[bodyLen];
                    recvBuff.get(body);

                    System.out.println("Recv server :" + new String(body, 0, bodyLen));
                    break;
                }

                if (rquest_times-- == 0) {
                    break;
                }

                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

                sendBuff.rewind();
            }
        } catch (IOException e) {
            e.printStackTrace();
            try {
                if (sock != null){
                    sock.close();
                }
            } catch (IOException e1) {
                e1.printStackTrace();
            }
        }
    }
}
```

Java NIO API 同样会抛出 java.io.IOException 异常，需要 catch。TCP 数据接收逻辑是比较难处理的，100 ~ 130 行都是数据读取的逻辑处理，主要是进行协议解析，通过 length 字段识别一个完整的消息。

## 4. Java NIO 服务器端实现步骤

因为服务端采用**非阻塞**模式，需要用到 Java NIO 的 Selector 组件，这是 Java NIO 的 I/O 多路复用机制，可以同时监听多个 SocketChannel 是否有读写事件。

* 创建 Java NIO 的 Selector 实例

```java
selector = Selector.open();
```

* 打开服务器 ServerSocketChannel

```java
serverChannel = ServerSocketChannel.open();
```

* 给 ServerSocketChannel 绑定监听的 socket 地址，监听 any_addr

```java
    serverChannel.socket().bind(new InetSocketAddress(PORT));
```

* 设置 SO_REUSEADDR 选项，作为服务器，这是基本的要求

```java
    serverChannel.socket().setReuseAddress(true);
```

* 设置非阻塞模式，这是服务器的基本要求，也是本小节的重点

```java
    serverChannel.configureBlocking(false);
```

* 向 Selector 注册 accept 事件

```java
    serverChannel.register(selector, SelectionKey.OP_ACCEPT, serverChannel);
```

* 编写事件循环。所有需要读写数据的 SocketChannel，需要将读写事件注册到 Selector。调用 Selector 的 select 方法，调用线程会进入 I/O 事件监听状态。如果没有事件发生，调用线程会被阻塞，进入事件等待状态；如果有事件发生，Selector 的 select 方法会返回发生了 I/O 事件的 SocketChannel 个数。Selector 的 selectedKeys 方法返回一个 java.util.Set 类，集合中包含的是 SelectionKey 结构，SelectionKey 和 SocketChannel 是一一对应的，表示发生了 I/O 事件的 SocketChannel。所以需要 遍历 Set，分别处理每个 SelectionKey。

```java
    while (true) {
       int readyChannels = selector.select();
       if (readyChannels == 0) {
           System.out.println("No socket has i/o events");
           continue;
       }

       Set<SelectionKey> selectedKeys = selector.selectedKeys();
       Iterator<SelectionKey> keyIterator = selectedKeys.iterator();

       while (keyIterator.hasNext()) {
           SelectionKey key = keyIterator.next();
           if (key != null) {

           }
           keyIterator.remove();
       }
    }
```

* 抽象一个内部类 Client，表示一个客户端连接，每当一个新连接建立的时候，创建一个此类的对象。

```java
private static class Client{
    public void sendData();
    public int recvData();
    public void close();
}
```

* 通过 key.isAcceptable() 处理连接**接收**事件。

```java
if (key.isAcceptable()) {
    // a connection was accepted by a ServerSocketChannel.
    ServerSocketChannel ssc = (ServerSocketChannel) key.attachment();
    SocketChannel newSock = ssc.accept();
    newSock.configureBlocking(false);
    Client client = new Client(selector, newSock);
}
```

* 通过 key.isReadable() 处理**读**事件

```java
if (key.isReadable()) {
    // a channel is ready for reading
    Client client = (Client) key.attachment();
    int rc = client.recvData();
    if (rc == 0) {
        client.sendData();
    }
}
```

* 通过 key.isReadable() 处理“写”事件

```java
if (key.isWritable()) {
     // a channel is ready for writing
     Client client = (Client) key.attachment();
     client.cancelEvent(SelectionKey.OP_WRITE);
     client.sendData();
}
```

服务器端完整代码如下：

```java
import java.io.*;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.*;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Iterator;
import java.util.Set;

public class NonblockTCPServer {
    // 服务器监听的端口
    private final static int PORT = 9082;
    private Selector selector = null;
    private ServerSocketChannel serverChannel = null;

    private static class Client{
        // 接收 buffer 长度
        private final static int RECV_BUF_LEN = 1024;
        // 接收buffer 声明
        private ByteBuffer recvBuff = null;
        // 发送 buffer 长度
        private static final int SEND_BUFF_LEN = 1024;
        // 发送 buffer 声明
        private ByteBuffer sendBuff = null;
        // the Selector
        private Selector selector = null;
        // SocketChannel 引用声明，表示一个连接
        private SocketChannel socketChannel = null;
        private SelectionKey sk_ = null;
        private boolean canSend = true;

        public Client(Selector selector, SocketChannel newSock){
            this.selector = selector;
            this.socketChannel = newSock;
            this.recvBuff = ByteBuffer.allocate(RECV_BUF_LEN);
            this.sendBuff = ByteBuffer.allocate(SEND_BUFF_LEN);
            this.register(SelectionKey.OP_READ);
        }

        private void register(int op){
            try {
                if (sk_ == null){
                    sk_ = this.socketChannel.register(selector, op, this);
                } else {
                    sk_.interestOps(op | sk_.interestOps());
                }
            } catch (ClosedChannelException e) {
                e.printStackTrace();
            }
        }

        public void cancelEvent(int ops){
            if (sk_ == null)
                return;

            sk_.interestOps(sk_.interestOps() & (~ops));
        }

        public void sendData() {
            try {
                int totalSendBytes = 0;
                String resp = null;
                if (canSend){
                    //设置日期格式
                    SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
                    resp = "The server time : " + df.format(new Date());
                    sendBuff.putInt(resp.length());
                    sendBuff.put(resp.getBytes());
                    totalSendBytes = resp.length() + 4;

                    sendBuff.flip();
                }else {
                    totalSendBytes = sendBuff.remaining();
                }

                int sbytes = this.socketChannel.write(sendBuff);
                System.out.println("Send to client about message :" + resp);
                if (sbytes < totalSendBytes) {
                    this.register(SelectionKey.OP_WRITE);
                    canSend = false;
                } else {
                    if (!canSend){
                        canSend = true;
                    }
                    sendBuff.rewind();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

        public int recvData(){
            try {
                int recvBytes = this.socketChannel.read(this.recvBuff);
                if (recvBytes < 0){
                    System.out.println("Meet error or the end of stream");
                    close();
                    return -1;
                }else if (recvBytes == 0){
                    return 0;// eagain
                }

                this.recvBuff.flip();
                while (this.recvBuff.remaining() > 0) {
                    // Incomplete message header
                    if (this.recvBuff.remaining() < 4) {
                        break;
                    }

                    int bodyLen = this.recvBuff.getInt();
                    if (bodyLen > this.recvBuff.remaining()) {
                        // Incomplete message body
                        break;
                    }

                    byte[] body = new byte[bodyLen];
                    this.recvBuff.get(body, 0, bodyLen);
                    System.out.println("Recv message from client: " +
                            new String(body, 0, bodyLen));
                }
                // flip recv buffer
                this.recvBuff.compact();
                return 0;
            } catch (IOException e) {
                e.printStackTrace();
                close();
            }
            return -1;
        }

        public void close(){
            try {
                cancelEvent(SelectionKey.OP_WRITE | SelectionKey.OP_READ);
                if (this.socketChannel != null){
                    this.socketChannel.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public void start(){
        try {
            selector = Selector.open();

            serverChannel = ServerSocketChannel.open();
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
            stop();
        }
    }

    public void process() {
        try {
            while (true) {
                int readyChannels = selector.select();
                if (readyChannels == 0) {
                    System.out.println("No socket has i/o events");
                    continue;
                }

                Set<SelectionKey> selectedKeys = selector.selectedKeys();
                Iterator<SelectionKey> keyIterator = selectedKeys.iterator();

                while (keyIterator.hasNext()) {
                    SelectionKey key = keyIterator.next();
                    if (key != null) {
                        if (key.isAcceptable()) {
                            // a connection was accepted by a ServerSocketChannel.
                            ServerSocketChannel ssc = (ServerSocketChannel) key.attachment();
                            SocketChannel newSock = ssc.accept();
                            newSock.configureBlocking(false);
                            Client client = new Client(selector, newSock);
                        } else if (key.isConnectable()) {
                            // a connection was established with a remote server.
                        } else if (key.isReadable()) {
                            // a channel is ready for reading
                            Client client = (Client) key.attachment();
                            int rc = client.recvData();
                            if (rc == 0) {
                                client.sendData();
                            }
                        } else if (key.isWritable()) {
                            // a channel is ready for writing
                            Client client = (Client) key.attachment();
                            client.cancelEvent(SelectionKey.OP_WRITE);
                            client.sendData();
                        }
                    }
                    keyIterator.remove();
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public void stop(){
        try {
            if (serverChannel != null){
                serverChannel.close();
                serverChannel = null;
            }
            if (selector != null) {
                selector.close();
                selector = null;
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        NonblockTCPServer tcp = new NonblockTCPServer();
        tcp.start();
        tcp.process();
    }
}
```

对于非阻塞式 Socket，需要处理发送 Buffer 满和接收 Buffer 为空的情况。服务器样例代码的 320 ~ 390 行的主要逻辑就是在处理非阻塞模式下，发送 Buffer 满和接收 Buffer 为空的逻辑。

## 5. 小结

本节主要是介绍了通过 Java NIO 编写**阻塞式**和**非阻塞式** Socket 程序的步骤。通过示例代码可以看出，阻塞式 Socket 程序结构简单，容易实现。非阻塞式 Socket 程序结构复杂，不容易实现，数据收发处理的细节较多，容易出错。

在编写非阻塞式 Java NIO Socket 程序，需要将 I/O 事件注册到 Selector，通过 Selector 的调度去处理具体的逻辑。往往是实现**事件驱动 I/O**架构的最佳选择。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
