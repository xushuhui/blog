---
title: Java从零开始（137）如何创建 Java UDP Socket
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 如何创建 Java UDP Socket


## 1. 前言

**UDP** 的英文全称是：User Datagram Protocol，翻译成中文叫**用户数据报协议**，它是 TCP/IP 协议族中一个非常重要的**传输层**协议。UDP 是一个无连接的、不可靠的传输层协议，没有丢包重传机制、没有流控机制、没有拥塞控制机制。UDP 不保证数据包的顺序，UDP 传输经常出现乱序，UDP 不对重复包进行过滤。

既然 UDP 这么多缺点，我们还有学习的必要吗？其实不然，正是因为 UDP 没有提供复杂的各种保障机制，才使得它具有实时、高效的传输特性。那么 UDP 到底有哪些优势呢？

* 第一，UDP 是面向消息的传输协议，保证数据包的边界，不需要进行消息解析，处理逻辑非常简单。
* 第二，UDP 具有实时、高效的传输特性。
* 第三，协议栈没有对 UDP 进行过多的干预，这给应用层带来了很多便利，应用程序可以根据自己的需要对传输进行控制。比如，自己实现优先级控制、流量控制、可靠性机制等。当然还有其他一些优势，我就不再一一列举。

UDP 在音视频会议、VOIP、音视频实时通信等行业有着广泛的应用。为此，我们是非常有必要学好 UDP 的。由于 UDP 相对简单，学习起来也会轻松很多。

## 2. 传统 UDP 客户端和服务器建立过程

同样，我们展示了通过 C 语言 Socket API 编写 UDP 客户端和服务器程序的步骤，如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyqrojej60hx0hl0yt02)

图中的矩形方框都是 C 函数。对比 TCP 客户端、服务器的建立过程，我们发现 UDP 客户端可以调用 connect 函数，但是并不会去连接服务器，只是和本地接口做绑定；UDP 服务器是没有 listen 和 accept 调用的。对于 UDP 客户端来说，connect 函数的调用是可选的。接下来，我们就探讨一下如何用 Java 语言编写 UDP 客户端和服务器程序。

## 3. Java DatagramSocket 分析

Java 语言抽象了 **java.net.DatagramSocket** 类，表示一个 UDP Socket，既可以用在客户端，又可以用在服务器端。**java.net.DatagramSocket** 是一个包装类，对外抽象了一组方法，具体实现是在 **java.net.DatagramSocketImpl** 类中完成的，它允许用户自定义具体实现。**java.net.DatagramSocket** 类包含的主要功能如下：

* 创建 UDP Socket，具体就是创建一个 **java.net.DatagramSocket** 类的对象。
* 将 Socket 绑定到本地接口 IP 地址或者端口，可以调用 **java.net.DatagramSocket** 类的**构造方法**或 **bind** 方法完成。
* 将客户端 UDP Socket 和远端 Socket 做绑定，可以通过 **java.net.DatagramSocket** 类的 connect 方法完成。

> 提示：
>
>
> UDP 客户端调用 connect 方法，仅仅是将本地 Socket 和远端 Socket 做绑定，并不会有类似 TCP 三次握手的过程。

* 关闭连接，可以调用 **java.net.DatagramSocket** 类的 close 方法完成。

* 接收数据，可以通过 **java.net.DatagramSocket** 类的 receive 方法实现数据接收。

* 发送数据，可以通过 **java.net.DatagramSocket** 类的 send 方法实现数据发送。

**java.net.Socket** 类提供了一组重载的构造方法，方便程序员选择，大体分为四类：

* 无参

```java
public DatagramSocket() throws SocketException
```

绑定到任意可用的端口和通配符 IP 地址，比如 IPv4 的 0.0.0.0。一般用作 UDP 客户端 Socket 的创建。

* 传入 port 参数

```java
public DatagramSocket(int port) throws SocketException
```

绑定到由 port 指定的端口和通配符 IP 地址，比如 IPv4 的 0.0.0.0。一般用作 UDP 服务端 Socket 的创建。

* 传入指定的 IP 和 Port 参数

```java
public DatagramSocket(SocketAddress bindaddr) throws SocketException
public DatagramSocket(int port, InetAddress laddr) throws SocketException
```

绑定到指定的端口和指定的网络接口。如果你的主机有多个网卡，并且你指向在某个指定的网卡上收发数据，可以用此构造方法。既可以用作 UDP 客户端 Socket，也可以用作 UDP 服务端 Socket。

## 4. Java UDP 客户端

我们创建一个简单的 UDP 客户端程序，代码如下：

```java
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetSocketAddress;
import java.net.SocketAddress;

public class UDPClient {
    private static final int PORT = 9002;
    private static final String DST_HOST = "127.0.0.1";
    private static final int RECV_BUFF_LEN = 1500;
    private static byte[] inBuff = new byte[RECV_BUFF_LEN];

    public static void main(String[] args) {
        // 创建 UDP 客户端 Socket，选择无参构造方法，由系统分配本地端口号和网络接口
        try (DatagramSocket udpClient = new DatagramSocket()){
            // 构造发送的目标地址，指定目标 IP 和目标端口号
            SocketAddress to = new InetSocketAddress(DST_HOST, PORT);
            while (true){
                String req = "Hello Server!";
                // 构造发送数据包，需要传入消息内容和目标地址结构 SocketAddress
                DatagramPacket message = new DatagramPacket(req.getBytes(), req.length(), to);
                // 发送消息
                udpClient.send(message);
                System.out.println("Send UDP message:"
                        + req + " to server:" + message.getSocketAddress().toString());

                // 构造接收消息的数据包，需要传入 byte 数组
                DatagramPacket inMessage = new DatagramPacket(inBuff, inBuff.length);
                // 接收消息
                udpClient.receive(inMessage);

                System.out.println("Recv UDP message:"
                        + new String(inMessage.getData(), 0, inMessage.getLength())
                        + " from server:" + inMessage.getSocketAddress().toString());

                // 每隔 2 秒发送一次消息
                try {
                    Thread.sleep(2000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

## 5. Java UDP 服务端

我们创建一个简单的 UDP 服务端程序，代码如下：

```java

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;

public class UDPServer {
    private static final int BIND_PORT = 9002;
    private static final String BIND_HOST = "127.0.0.1";
    private static final int RECV_BUFF_LEN = 1500;
    private static byte[] inBuff = new byte[RECV_BUFF_LEN];

    public static void main(String[] args) {
        // 构造服务器 Socket，绑定到一个固定的端口，监听的 IP 是 0.0.0.0
        try (DatagramSocket udpServer = new DatagramSocket(BIND_PORT)) {
            // 构造接收消息的数据包，需要传入 byte 数组。
            // 我们将这条语句放在循环外，不需要每次消息收发都构造此结构
            DatagramPacket inMessage = new DatagramPacket(inBuff, inBuff.length);
            while (true){
                // 接收客户端消息
                udpServer.receive(inMessage);
                System.out.println("Recv UDP message:"
                        + new String(inMessage.getData(), 0, inMessage.getLength())
                        + " from Client:"
                        + inMessage.getSocketAddress().toString());

                String rsp = "Hello Client!";
                // 构造发送的消息结构
                // 注意！！！对于服务器来说，发送的目标地址一定是接收消息时的源地址，所以从 inMessage 结构获取
                DatagramPacket message = new DatagramPacket(rsp.getBytes(), rsp.length(),
                        inMessage.getSocketAddress());
                // 发送消息
                udpServer.send(message);
                System.out.println("Send UDP message:"
                        + rsp + " to Client:" + message.getSocketAddress().toString());
                // 重置接收数据包消息长度，准备接收下一个消息
                inMessage.setLength(inBuff.length);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

## 6. 小结

用 Java 语言编写 UDP 客户端和服务器程序非常简单，你只需要创建一个 **java.net.DatagramSocket** 实例。如果你构造的是服务器 Socket，需要传入监听的**端口号**，监听的接口 IP 是可选的，默认是监听通配符 IP。如果你创建的是客户端 Socket，你可以传入绑定的本地 Port 和接口地址，也可以不传入任何参数。对于客户端 UDP Socket 来说，你也可以调用 connect 方法，只是和远端 Socket 绑定，没有类似 TCP 的三次握手过程。

示例代码我们采用的是 try-with-resources 写法，对于 Java 7 以后的版本，可以不显式调用 close 方法。如果是非此写法，需要显式调用 close 方法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
