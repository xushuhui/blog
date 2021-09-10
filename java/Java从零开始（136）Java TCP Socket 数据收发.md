# Java TCP Socket 数据收发

## 1. 前言

TCP 是面向**字节流**的传输协议。所谓**字节流**是指 TCP 并不理解它所传输的数据的含义，在它眼里一切都是**字节**，1 **字节**是 8 **比特**。比如，TCP 客户端向服务器发送“Hello Server，I’m client。How are you？”，TCP 客户端发送的是具有一定含义的数据，但是对于 TCP 协议栈来说，传输的是一串**字节流**，具体如何解释这段数据需要 TCP 服务器的应用程序来完成，这就涉及到“应用层协议设计”的问题。

在 TCP/IP 协议栈的四层协议模型中，操作系统内核协议栈实现了链路层、网络层、传输层，将应用层留给了应用程序来实现。在编程实践中，通常有**文本协议**和**二进制协议**两种类型，前者通常通过一个**分隔符**区分消息语义，而后者通常是需要通过一个 length 字段指定消息体的大小。比如著名的 HTTP 协议就是文本协议，通过 “\r\n” 来区分 HTTP Header 的每一行。而 RTMP 协议是一个二进制协议，通过 length 字段来指定消息体的大小。

解析 TCP 字节流的语义通常叫做**消息解析**，如果按照传统 C 语言函数的方式来实现，还是比较麻烦的，有很多细节需要处理。好在 Java 为我们提供了很多工具类，给我们的工作带来了极大地便利。

## 2. Java 字节流结构

Java 的 [java.io](http://java.io).* 包中包含了 InputStream 和 OutputStream 两个类，是 Java 字节流 I/O 的基础类，其他具体的 Java I/O 字节流功能类都派生自这两个类。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyqfrtdj60pg086dgl02)

图中只列出了我们 Socket 编程中常用的 I/O 字节流类。**java.net.SocketInputStream** 类是 Socket 的输入流实现类，它继承了 java.io.FileInputStream 类。**java.net.SocketOutputStream** 类是 Socket 的输出流实现类，它继承了 java.io.FileOutputStream 类，下来我们逐一介绍这些类的基本功能。

### 2.1 Java InputStream & OutputStream

**java.io.InputStream** 类是一个抽象超类，它提供最小的编程接口和输入流的部分实现。**java.io.InputStream** 类定义的几类方法：

* 读取字节或字节数组，一组 read 方法。
* 标记流中的位置，mark 方法。
* 跳过输入字节，skip 方法。
* 找出可读取的字节数，available 方法。
* 重置流中的当前位置，reset 方法。
* 关闭流，close 方法。

InputStream 流在创建实例时会自动打开，你可以调用 close 方法显式关闭流，也可以选择在垃圾回收 InputStream 时，隐式关闭流。需要注意的是垃圾回收机制关闭流，并不能立刻生效，可能会造成流对象泄漏，所以一般需要主动关闭。

**java.io.OutputStream** 类同样是一个抽象超类，它提供最小的编程接口和输出流的部分实现。**java.io.OutputStream** 定义的几类方法：

* 写入字节或字节数组，一组 write 方法。
* 刷新流，flush 方法。
* 关闭流，close 方法。

OutputStream 流在创建时会自动打开，你可以调用 close 方法显式关闭流，也可以选择在垃圾回收 OutputStream 时，隐式关闭流。

### 2.2 FileInputStream & FileOutputStream

java.io.FileInputStream 和 java.io.FileOutputStream 是文件输入和输出流类，用于从本机文件系统上的文件读取数据或向其写入数据。你可以通过文件名、java.io.File 对象、java.io.FileDescriptor 对象创建一个 FileInputStream 或 FileOutputStream 流对象。

### 2.3 SocketOutputStream & SocketInputStream

java.net.SocketInputStream 和 java.net.SocketOutputStream 代表了 Socket 流的读写，他们分别继承自 java.io.FileInputStream 和 java.io.FileOutputStream 类，这说明 Socket 读写包含了文件读写的特性。另外，这两个类是定义在 [java.net](http://java.net).* 包中，并没有对外公开。

### 2.4 FilterInputStream & FilterOutputStream

java.io.FilterInputStream 和 java.io.FilterOutputStream 分别是 java.io.InputStream 和 java.io.OutputStream 的子类，并且它们本身都是抽象类，为被过滤的流定义接口。java.io.FilterInputStream 和 java.io.FilterOutputStream 的主要作用是为基础流提供一些额外的功能，这些不同的功能都是单独的类，继承了他们的接口。例如，过滤后的流 BufferedInputStream 和 BufferedOutputStream 在读写时会缓冲数据，以加快数据传输速度。

### 2.5 BufferedInputStream & BufferedOutputStream

java.io.BufferedInputStream 类继承自 java.io.FilterInputStream 类，它的作用是为 java.io.FileInputStream、java.net.SocketInputStream 等输入流提供缓冲功能。一般通过 java.io.BufferedInputStream 的构造方法传入具体的输入流，同时可以指定缓冲区的大小。java.io.BufferedInputStream 会从底层 Socket 读取一批数据保存到内部缓冲区中，后续通过 java.io.BufferedInputStream 的 read 方法读取数据，实际上都从缓冲区中读取，等读完缓冲中的这部分数据之后，再从底层 Socket 中读取下一部分的数据。

> * 注意：
> * 当你调用 java.io.BufferedInputStream 的 read 方法读取一个数组时，只有当读取的数据达到数组长度时才会返回，否则线程会被阻塞。

java.io.BufferedOutputStream 类继承自 java.io.FilterOutputStream 类，它的作用是为 java.io.FileOutputStream、java.net.SocketOutputStream 等输出流提供缓冲功能。一般通过 java.io.BufferedOutputStream 的构造方法传入底层输出流，同时可以指定缓冲区的大小。每次调用 java.io.BufferedOutputStream 的 write 方法写数据时，实际上是写入它的内部缓冲区中，当内部缓冲区写满或者调用了 flush 方法，才会将数据写入底层 Socket 的缓冲区。

BufferedInputStream 和 BufferedOutputStream 在读取或写入时缓冲数据，从而减少了对原始数据源所需的访问次数。缓冲流通常比类似的非缓冲流效率更高。

### 2.6 DataInputStream & DataOutputStream

java.io.DataInputStream 和 java.io.DataOutputStream 类继承自 java.io.FilterInputStream 和 java.io.FilterOutputStream 类，同时实现了 java.io.DataInput 和 java.io.DataOutput 接口，功能是以机器无关的格式读取或写入原始 Java 数据类型。

## 3. 数据读写的案例程序

我们设计一个简单的协议，每个消息的开头 4 字节表示消息体的长度，格式如下：

```java
+-----------------+
| 4 字节消息长度   |
+-----------------+
|                 |
|   消息体         |
|                 |
+-----------------+
```

我们通过这个简单的协议演示 java.io.DataInputStream 、java.io.DataOutputStream 和 java.io.BufferedInputStream、java.io.BufferedOutputStream 类的具体用法。TCP 客户端和服务器的编写可以参考上一节内容，本节仅展示数据读写的代码片段。

客户端数据读写代码：

```java
import java.io.*;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.net.SocketAddress;

public class TCPClientIO {
    // 服务器监听的端口号
    private static final int PORT = 56002;
    private static final int TIMEOUT = 15000;

    public static void main(String[] args) {
        Socket client = null;
        try {
            // 调用无参构造方法
            client = new Socket();
            // 构造服务器地址结构
            SocketAddress serverAddr = new InetSocketAddress("127.0.0.1", PORT);
            // 连接服务器，超时时间是 15 毫秒
            client.connect(serverAddr, TIMEOUT);

            System.out.println("Client start:" + client.getLocalSocketAddress().toString());

            // 向服务器发送数据
            DataOutputStream out = new DataOutputStream(
                    new BufferedOutputStream(client.getOutputStream()));
            String req = "Hello Server!\n";
            out.writeInt(req.getBytes().length);
            out.write(req.getBytes());
            // 不能忘记 flush 方法的调用
            out.flush();
            System.out.println("Send to server:" + req + " length:" +req.getBytes().length);

            // 接收服务器的数据
            DataInputStream in = new DataInputStream(
                    new BufferedInputStream(client.getInputStream()));

            int msgLen = in.readInt();
            byte[] inMessage = new byte[msgLen];
            in.read(inMessage);
            System.out.println("Recv from server:" + new String(inMessage) + " length:" + msgLen);
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (client != null){
                try {
                    client.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
```

服务端数据读写代码：

```java
import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.io.BufferedInputStream;
import java.io.DataInputStream;
import java.io.BufferedOutputStream;
import java.io.DataOutputStream;

public class TCPServerIO {
    private static final int PORT =56002;

    public static void main(String[] args) {
        ServerSocket ss = null;
        try {
            // 创建一个服务器 Socket
            ss = new ServerSocket(PORT);
            // 监听新的连接请求
            Socket conn = ss.accept();
            System.out.println("Accept a new connection:" + conn.getRemoteSocketAddress().toString());

            // 读取客户端数据
            DataInputStream in = new DataInputStream(
                    new BufferedInputStream(conn.getInputStream()));
            int msgLen = in.readInt();
            byte[] inMessage = new byte[msgLen];
            in.read(inMessage);
            System.out.println("Recv from client:" + new String(inMessage) + "length:" + msgLen);

            // 向客户端发送数据
            String rsp = "Hello Client!\n";

            DataOutputStream out = new DataOutputStream(
                    new BufferedOutputStream(conn.getOutputStream()));
            out.writeInt(rsp.getBytes().length);
            out.write(rsp.getBytes());
            out.flush();
            System.out.println("Send to client:" + rsp + " length:" + rsp.getBytes().length);
            conn.close();
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (ss != null){
                try {
                    ss.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }

        System.out.println("Server exit!");
    }
}

```

注意读写消息长度需要用 readInt 和 writeInt 方法。

## 4. 总结

通过本节学习，你需要树立一个观念：TCP 是面向字节流的协议，TCP 传输数据的时候并不保证消息边界，消息边界需要程序员设计应用层协议来保证。将字节流解析成自定义的协议格式，需要借助 [java.io](http://java.io).* 中提供的工具类，一般情况下，java.io.DataInputStream 、java.io.DataOutputStream 和 java.io.BufferedInputStream、java.io.BufferedOutputStream 四个类就足以满足你的需求了。DataInputStream 和 DataOutputStream 类主要是用以读写 java 相关的数据类型，BufferedInputStream 和 BufferedOutputStream 解决缓冲读写的问题，目的是提高读写效率。

本节简要介绍了 Socket 编程中常用的 I/O 流类，关于 [java.io](http://java.io).* 包中的各种 I/O 流类不是本节的重点，需要你自己参考相关 Java 书籍。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
