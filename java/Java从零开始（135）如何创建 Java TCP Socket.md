---
title: Java从零开始（135）如何创建 Java TCP Socket
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 如何创建 Java TCP Socket


## 1. 前言

**TCP** 的英文全称是 **Transmission Control Protocol**，翻译成中文叫做**传输控制协议**，它是 TCP/IP 协议族中非常重要的一个**传输层**协议。TCP 是一个面向连接的、面向字节流的、可靠的传输层协议，有丢包重传机制、有流控机制、有拥塞控制机制。TCP 保证数据包的顺序，并且对重复包进行过滤。相比不可靠传输协议 UDP，TCP 完全是相反的。

对于可靠性要求很高的应用场景来说，选择可靠 TCP 作为传输层协议肯定是正确的。例如，著名的 HTTP 协议和 FTP 协议都是采用 TCP 进行传输。当然 TCP 为了保证传输的可靠性，引入了非常复杂的保障机制，比如：TCP 连接建立时的三次握手和连接关闭时的四次挥手机制，滑动窗口机制，发送流控机制，慢启动和拥塞避免机制等。当然，操作系统的网络协议栈已经实现了这些复杂的机制，

本小节主要是介绍通过 Java 语言编写 TCP 客户端、服务器程序的方法。

编写 TCP 客户端、服务器程序主要分为如下几个步骤：

* 创建客户端 Socket，连接到某个服务器监听的端口，需要指定服务器监听的 host 和 port。host 可以是 IP 地址，也可以是域名。
* 创建服务端 Socket，绑定到一个固定的服务端口，监听客户端的连接请求。
* 客户端发起连接请求，完成三次握手过程。
* TCP 连接建立成功后，双方进行数据流交互。
* 数据流交互完成后，关闭连接。

## 2. 传统 TCP 客户端和服务器建立过程

为了更好地理解编写 TCP 客户端和服务器程序的步骤，下图展示了通过 C 语言 Socket API 编写客户端和服务器程序的过程。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyq2dptj60hw0hj79q02)

图中的矩形方框都是 C 函数，很好的展示了客户端和服务器 Socket 的建立过程。对于 Java 语言来说，只是应用面向对象的思维对上面的过程进行了抽象，下来我们就探讨一下如何编写 Java 客户端和服务器程序。

## 3. Java Socket 类分析

Java 语言抽象了 **java.net.Socket** 类，表示一个 Socket，既可以用在客户端，又可以用在服务器端。其实 **java.net.Socket** 也是一个包装类，对外抽象了一组公共方法，具体实现是在 **java.net.SocketImpl** 类中完成的，它允许用户自定义具体实现。**java.net.Socket** 类包含的主要功能如下：

* 创建 Socket，具体就是创建一个 **java.net.Socket** 类的对象。
* 建立 TCP 连接，可以通过 **java.net.Socket** 类的构造方法完成，也可以调用它的 connect 方法完成。
* 将 Socket 绑定到本地接口 IP 地址或者端口，可以调用 **java.net.Socket** 类的 bind 方法完成。

> 提示：
>
>
> 服务器需要做 bind 操作，客户端一般不需要做 bind 操作。

* 关闭连接，可以调用 **java.net.Socket** 类的 close 方法完成。

* 接收数据，可以通过 **java.net.Socket** 类的 getInputStream 方法，返回一个 java.io.InputStream 对象实现数据接收。

* 发送数据，可以通过 **java.net.Socket** 类的 getOutputStream 方法，返回一个 java.io.OutputStream 对象实现数据发送。

**java.net.Socket** 类提供了一组重载的构造方法，方便程序员选择，大体分为四类：

* 可以传入服务器的 host 和 port 参数

原型如下：

```java
  public Socket(String host, int port)
        throws UnknownHostException, IOException
  public Socket(InetAddress address, int port) throws IOException
```

对于 host 参数，你可以传入 IP 地址或者是**域名**。当然，你可以传入构造好的 InetAddress 地址结构。

在 **java.net.Socket** 的构造方法中，首先会构造一个 InetAddress 地址结构，然后进行域名解析，最后调用它的 connect 方法和服务器建立连接。

* 可以传入绑定的本地地址参数

原型如下：

```java
  public Socket(String host, int port, InetAddress localAddr,  int localPort) throws IOException
  public Socket(InetAddress address, int port, InetAddress localAddr,  int localPort) throws IOException
```

这类构造方法也可以传入 host 和 port 外，功能和上面类似。另外，还可以传入 localAddr 和 localPort，会调用 **java.net.Socket** 类的 bind 方法，绑定在本地的接口地址和端口。

* 无参构造方法

```java
  public Socket()
```

此构造方法，除了构造一个 **java.net.Socket** 类的对象，并不会去 connect 服务器。你需要调用它的 connect 方法连接服务器。

```java
public void connect(SocketAddress endpoint, int timeout) throws IOException
```

自己调用 connect 方法，需要构造 SocketAddress 结构，当然你可以设置连接的超时时间，单位是毫秒（milliseconds）。

* 访问代理服务器

```java
public Socket(Proxy proxy)
```

当你需要访问某个代理服务器时，可以调用此构造方法，Socket 会自动去连接代理服务器。

创建一个简单的 **java.net.Socket** 客户端，示例代码如下：

```java
import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.net.SocketAddress;

public class TCPClient {
    // 服务器监听的端口号
    private static final int PORT = 56002;
    private static final int TIMEOUT = 15000;

    public static void main(String[] args) {
        Socket client = null;
        try {
            // 在构造方法中传入 host 和 port
            // client = new Socket("192.168.43.49", PORT);

            // 调用无参构造方法
            client = new Socket();
            // 构造服务器地址结构
            SocketAddress serverAddr = new InetSocketAddress("192.168.0.101", PORT);
            // 连接服务器，超时时间是 15 毫秒
            client.connect(serverAddr, TIMEOUT);

            System.out.println("Client start:" + client.getLocalSocketAddress().toString());

            // 向服务器发送数据
            OutputStream out = new BufferedOutputStream(client.getOutputStream());
            String req = "Hello Server!\n";
            out.write(req.getBytes());
            // 不能忘记 flush 方法的调用
            out.flush();
            System.out.println("Send to server:" + req);

            // 接收服务器的数据
            BufferedInputStream in = new BufferedInputStream(client.getInputStream());
            StringBuilder inMessage = new StringBuilder();
            while(true){
                int c = in.read();
                if (c == -1 || c == '\n')
                    break;
                inMessage.append((char)c);
            }
            System.out.println("Recv from server:" + inMessage.toString());
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

这里我们创建的是阻塞式的客户端，有几点需要注意的地方：

> * 通过 OutputStream 的对象向服务器发送完数据后，需要调用 flush 方法。
> * BufferedInputStream 的 read 方法会阻塞线程，所以需要设计好**消息边界的识别机制**，示例代码是通过换行符 ‘\n’ 表示一个消息边界。
> * **java.net.Socket** 的各个方法都抛出了 IOException 异常，需要捕获。
> * 注意调用 close 方法，关闭连接。

## 4. Java ServerSocket 类分析

Java 语言抽象了 **java.net.ServerSocket** 类表示服务器监听 Socket，此类只用在服务器端，通过调用它的 accept 方法来获取新的连接。accept 方法的返回值是 **java.net.Socket** 类型，后续服务器和客户端的数据收发，都是通过 accept 方法返回的 Socket 对象完成。

**java.net.ServerSocket** 类也提供了一组重载的构造方法，方便程序员选择。

```java
  public ServerSocket(int port) throws BindException, IOException
  public ServerSocket(int port, int queueLength) throws BindException, IOException
  public ServerSocket(int port, int queueLength, InetAddress bindAddress) throws IOException
  public ServerSocket() throws IOException
```

* port 参数用于传入服务器监听的端口号。如果传入的 port 是 0，系统会随机选择一个端口监听。
* queueLength 参数用于设置连接接收队列的长度。不传入此参数，采用系统默认长度。
* bindAddress 参数用于将监听 Socket 绑定到一个本地接口。如果传入此参数，服务器会监听指定的接口地址；如果不指定此参数，默认会监听通配符 IP 地址，比如 IPv4 是 0.0.0.0。

> * 提示：
> * 可以通过 netstat 命令查看服务器程序监听的 IP 地址和端口号。

如果你是通过无参构造方法构造 **java.net.ServerSocket** 类的对象，需要手动调用它的 bind 方法，绑定监听端口和接口地址。

创建一个简单的服务器监听 Socket，示例代码如下：

```java
import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;

public class TCPServer {
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
            BufferedInputStream in = new BufferedInputStream(conn.getInputStream());
            StringBuilder inMessage = new StringBuilder();
            while(true){
                int c = in.read();
                if (c == -1 || c == '\n')
                    break;
                inMessage.append((char)c);
            }
            System.out.println("Recv from client:" + inMessage.toString());

            // 向客户端发送数据
            String rsp = "Hello Client!\n";
            BufferedOutputStream out = new BufferedOutputStream(conn.getOutputStream());
            out.write(rsp.getBytes());
            out.flush();
            System.out.println("Send to client:" + rsp);
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

我们创建的阻塞式服务端，所以 **java.net.ServerSocket** 的 accept 方法会阻塞线程，直到新连接返回。同样，在接收客户端的消息的时候注意消息边界的处理，最后向客户端发送响应的时候，需要调用 flush 方法。

## 5. 小结

用 Java 语言编写 TCP 客户端和服务器程序非常方便，你只需要创建一个 **java.net.ServerSocket** 实例，然后调用它的 accept 方法监听客户端的请求；你只需要创建一个 **java.net.Socket** 实例，可以通过构造方法或者 connect 连接对应的服务器，然后就可以进行数据的收发，最后数据交互完成后，调用 close 方法关闭连接即可。

示例代码中的服务器功能还不完善，不能持续提供服务，不能同时接收多个客户端的连接请求，需要在后续的小节逐步完善。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
