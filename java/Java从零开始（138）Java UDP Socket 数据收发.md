# Java UDP Socket 数据收发

## 1. 前言

UDP 是面向**数据报**的传输协议。UDP 的包头非常简单，总共占用 8 字节长度，格式如下：

```java
+--------------------+--------------------+
| 源端口（16 bits）   | 目的端口（16 bits） |
+--------------------+--------------------+
| 包的长度（16 bits） | 检验和（16 bits）   |
+--------------------+--------------------+
```

源端口号：占用 2 字节长度，用于标识发送端**应用程序**。

目的端口：占用 2 字节长度，用于标识接收端**应用程序**。

包的长度：表示 UDP 数据包的总长度，占用 2 字节长度。包的长度的值是 UDP 包头的长度加上 UDP 包体的长度。 包体最大长度是 65536-8 = 65528 字节。

> 提示：
>
>
> 网络层的 IPv4 Header 也包含了 Length 字段，IPv4 Payload 的最大长度是 65536-60 = 65476 字节。如果我们控制 UDP 数据包总长度不超过 65476 字节，UDP Header 其实是不需要 UDP Length 字段的。因为在实际开发中，程序员会保证传给 UDP 的数据长度不超过 MTU 最大限度，所以 UDP Length 字段显得有点儿多余。

检验和：占用 2 字节长度。UDP 检验和是用于差错检测的，检验计算包含 UDP 包头和 UDP 包体两部分。

从 UDP 的协议格式可以看出，UDP 保证了应用层消息的完整性，比如，UDP 客户端向服务器发送 “Hello Server，I’m client。How are you？”，UDP 客户端发送的是具有一定含义的数据，UDP 服务端收到也是这个完整的消息。不像面向字节流的 TCP 协议，需要应用程序解析消息。为此，UDP 编程会简单很多。

## 2. Java DatagramPacket 分析

Java 抽象了 **java.net.DatagramPacket** 类表示一个 UDP 数据报，主要功能如下：

* 发送：    * 设置发送的数据。
    * 设置接收此数据的目的主机的 IP 地址和端口号。
    * 获取发送此数据的源主机的 IP 地址和端口号。

* 接收：    * 设置接收数据的 byte 数组。
    * 获取发送此数据的源主机的 IP 地址和端口号。
    * 获取接收此数据的主机目的主机的 IP 地址和端口号。

接收数据的构造方法：

```java
public DatagramPacket(byte[] buffer, int length)
public DatagramPacket(byte[] buffer, int offset, int length)
```

当接收数据的时候，需要构造 java.net.DatagramPacket 的实例，并且要传入接收数据的 byte 数组，然后调用 java.net.DatagramSocket 的 receive 方法就可以接收数据。当 receive 方法调用返回以后，发送此数据包的源主机 IP 地址和端口号保存在 java.net.DatagramSocket 的实例中。

发送数据的构造方法：

```java
public DatagramPacket(byte[] data, int length,InetAddress destination, int port)
public DatagramPacket(byte[] data, int offset, int length,InetAddress destination, int port)
public DatagramPacket(byte[] data, int length,SocketAddress destination)
public DatagramPacket(byte[] data, int offset, int length,SocketAddress destination)
```

当发送数据的时候，同样需要构造 java.net.DatagramPacket 的实例，并且要传入将要发送的数据的 byte 数组，同时要传入接收此数据包的目标主机 IP 地址和端口号，然后调用 java.net.DatagramSocket 的 send 方法就可以发送数据。目标主机的 IP 地址和端口号保存在 java.net.DatagramSocket 的实例中，你可以调用它的 getSocketAddress 方法获取。

获取或设置数据：

```java
public byte[] getData()

public void setData(byte[] data)
public void setData(byte[] data, int offset, int length)
```

获取或设置数据的长度：

```java
public int getLength()
public void setLength(int length)
```

获取设置 IP 地址和端口号

```java
public int getPort()
public InetAddress getAddress() // 只能获取 IP
public SocketAddress getSocketAddress()// 同时获取 IP 和 Port

public void setAddress(InetAddress remote)// 只能设置 IP
public void setPort(int port)
public void setAddress(SocketAddress remote)// 设置 SocketAddress，同时设置 IP 和 Port
```

## 3. UDP 消息序列化与反序列化

java.io.ByteArrayInputStream 和 java.io.ByteArrayOutputStream 继承自 java.io.InputStream 和 java.io.OutputStream。可以作为流的源和流的目标类，当你需要解析复杂的协议格式的时候，可以配合 java.io.DataInputStream 和 java.io.DataOutputStream 类实现消息的序列化、反序列化。

下来我们定义一个简单的消息格式，通常在音视频通信中会遇到这样的消息格式，我们这里只是一个演示版本。具体字段如下：

* version 表示协议版本号，这是一般协议格式都会包含的一个字段。
* flag，一些控制标志，主要表现在用不同的 bit 位表示不同的控制标志。
* sequence，对每个消息进行编号，用来检测是否有丢包发生。
* timestamp，每一个消息都携带一个发送时间戳，可以计算网络延迟、抖动。
* 消息体，消息的具体内容。

图示如下：

```java
+-----------------+-----------------+-----------------|-----------------+
| version(8 bits) | flag(8 bits)    |          Sequence(16 bits)        |
+-----------------|-----------------+-----------------------------------+
|                  Timestamp(32 bits)                                   |
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
|                                                                       |
|                       Message Body                                    |
|                                                                       |
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
```

对于这样一个格式，通过 java.net.DatagramPacket 类读取或者是设置的是一个 byte 数组，要想解析数组中消息各个字段的含义，需要借助 java.io.ByteArrayInputStream 和 java.io.ByteArrayOutputStream 类，以及 java.io.DataInputStream 和 java.io.DataOutputStream 类。

我们设计了一个 Message 类用来表示消息结构，当然 Message 类要包含协议格式中的各个字段。除了提供各个属性的 getter/setter 方法外，还提供了 serialize 和 deserialize 方法，实现了消息的**序列化**、**反序列化**。最后，我们覆盖了 toString 方，将 Message 转换成 String 形式。

代码清单如下：

```java
import java.io.*;
import java.net.DatagramPacket;

public class Message implements Serializable {
    private static final int SEND_BUFF_LEN = 512;
    private static ByteArrayOutputStream outArray = new ByteArrayOutputStream(SEND_BUFF_LEN);

    private byte version =1;
    private byte flag;
    private short sequence;
    private int timestamp;
    private byte[] body = null;
    private int bodyLength = 0;

    public byte getVersion() {
        return version;
    }

    public void setVersion(byte version) {
        this.version = version;
    }

    public byte getFlag() {
        return flag;
    }

    public void setFlag(byte flag) {
        this.flag = flag;
    }

    public short getSequence() {
        return sequence;
    }

    public void setSequence(short sequence) {
        this.sequence = sequence;
    }

    public int getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(int timestamp) {
        this.timestamp = timestamp;
    }

    public byte[] getBody() {
        return body;
    }

    public void setBody(byte[] body) {
        this.body = body;
    }

    public DatagramPacket serialize()
    {
        try {
            outArray.reset();
            DataOutputStream out = new DataOutputStream(outArray);

            out.writeByte(this.getVersion());
            out.writeByte(this.getFlag());
            out.writeShort(this.getSequence());
            out.writeInt(this.getTimestamp());
            out.write(this.body);
            // 构造发送数据包，需要传入消息内容和目标地址结构 SocketAddress
            byte[] outBytes = outArray.toByteArray();
            DatagramPacket message = new DatagramPacket(outBytes,  outBytes.length);

            return message;
        } catch (IOException e) {
            e.printStackTrace();
        }

        return null;
    }
    public void deserialize(DatagramPacket inMessage)
    {
        try {
            DataInputStream in = new DataInputStream(
                    new ByteArrayInputStream(inMessage.getData(), 0, inMessage.getLength()));
            this.version = in.readByte();
            this.flag = in.readByte();
            this.sequence = in.readShort();
            this.timestamp = in.readInt();
            this.body = new byte[512];
            this.bodyLength = in.read(this.body);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public String toString() {
        return " version: " + this.getVersion()
                + " flag: " + this.getFlag()
                + " sequence: " + this.getSequence()
                + " timestamp: " + this.getSequence()
                + " message body: " +new String(body, 0, this.bodyLength);
    }
}
```

通过 DataOutputStream 和 ByteArrayOutputStream 的配合，实现 serialize 功能。通过 DataInputStream 和 ByteArrayInputStream 配合，实现 deserialize 功能。

Message 序列化的用法：

```java
private static final int PORT = 9002;
private static final String DST_HOST = "127.0.0.1";
private static short sequence = 1;

SocketAddress to = new InetSocketAddress(DST_HOST, PORT);

String req = "Hello Server!";
Message sMsg = new Message();
sMsg.setVersion((byte)1);
sMsg.setFlag((byte)21);
sMsg.setSequence(sequence++);
sMsg.setTimestamp((int)System.currentTimeMillis()&0xFFFFFFFF);
sMsg.setBody(req.getBytes());
DatagramPacket outMessage = sMsg.serialize();
 outMessage.setSocketAddress(to);
```

Message 反列化的用法：

```java
Message rMsg = new Message();
rMsg.deserialize(inMessage);// inMessage 是一个 DatagramPacket 类型的实例
```

## 4. 小结

本节首先介绍了 **java.net.DatagramPacket** 类的基本功能，这是 Java UDP Socket 程序进行数据读写的基础类。在调用 receive 方法接收数据之前，首先要创建 DatagramPacket 的实例，同时要为他提供一个介绍数据的字节数组。当 receive 方法成功返回后，你可以调用 DatagramPacket 的 getSocketAddress 方法获取发送主机的源 IP 地址和端口号。在调用 send 方法发送数据之前，首先要创建 DatagramPacket 的实例，将要发送的数据传给他，同时要将接收数据的目标主机的 IP 地址和端口号设置给它。

接着我们重点介绍了 UDP 编程中常见的协议格式定义、解析方法，主要是通过 java.io.ByteArrayInputStream 和 java.io.ByteArrayOutputStream 类，以及 java.io.DataInputStream 和 java.io.DataOutputStream 类实现消息的序列化、反序列化功能。我们提供了完整的实现代码，已经序列化、反序列化的具体用法。可以说这一部分内容在实践中经常会遇到，需要好好掌握。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
