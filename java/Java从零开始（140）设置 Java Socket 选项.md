# 设置 Java Socket 选项

## 1. 前言

前面章节介绍了 Java TCP、UDP Socket 编程方法，按照文中介绍的方法去编写 Socket 程序，是完全可以正常工作的。其实，TCP/IP 协议栈允许你对 Socket 做一些定制，比如设置 Socket 的接收、发送缓冲区的大小，这就是常说的 **Socket 选项**。

本文首先会以 Linux 系统为例，介绍操作系统 Socket 选项的基本概念，然后再介绍 Java 中如何去设置 Socket 选项。

## 2. Socket 选项的概念

操作系统协议栈支持的 Socket 选项参数有很多，汇总起来如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnys0e9uj60pg0bltcp02)

从图中可以看出，Socket 选项按照级别进行分类，级别有很多种，但是总结起来分两类：

* 通用 Socket 级别的选项。枚举值为 SOL_SOCKET。
* 协议相关的选项。协议栈为我们提供了控制所有协议的选项，比如 IP、IPv6、TCP、UDP、ICMP 等。枚举值的格式为 IPPROTO_XXX，XXX 代表协议。

每一种选项级别下面包含了很多选项参数。比如，通用 Socket 选项的级别枚举值是 SOL_SOCKET，其下面包含 SO_RCVBUF 和 SO_SNDBUF 选项参数；IP 协议选项的级别的枚举值是 IPPROTO_IP，其下面包含 IP_TTL、IP_TOS 等选项参数。

在 Linux 系统下，所有的选项参数都可以在帮助手册里面查找，具体方法如下：

通用 Socket 级别选项参数

```java
man 7 socket
```

IP 协议级别选项参数：

```java
man 7 ip
```

IPv6 协议级别选项参数：

```java
man 7 ipv6
```

TCP 协议级别选项参数：

```java
man 7 tcp
```

UDP 协议级别选项参数：

```java
man 7 udp
```

Socket 选项参数最终是如何设置到协议栈的呢？协议栈提供了 getsockopt() 和 setsockopt() 两个 C 语言函数，分别用于获取和设置选项参数。

调用两个函数所需要包含的头文件，以及他们的声明如下：

```java
#include <sys/types.h>
#include <sys/socket.h>

int getsockopt(int sockfd, int level, int optname, void *optval, socklen_t *optlen);
int setsockopt(int sockfd, int level, int optname, const void *optval, socklen_t optlen);
```

如果你对系统本身的 Socket 选项感兴趣，可以通过 man 查找相关帮助。本节重点介绍通用 Socket 选项。

## 3. 通用 Socket 选项

通用 Socket 选项的 level 枚举值是 SOL_SOCKET。

表格中选项名称不用多说，数据类型列表示选项值的类型，大多数是整形，还有一些是结构体类型。有的选项是既可以设置值也可以读取值，用 set 表示；有的选项只能读取值，用 get 表示。常见选项参数如下：

|选项名称|数据类型|get 或 set|说明|
|--------|--------|----------|----|
|SO_BROADCAST|int          |set|设置 Socket 可以进行局域网广播，目标 IP 需要填网段的广播地址或者是统一受限广播地址 255.255.255.255。                 |
|SO_KEEPALIVE|int          |set|用于设置 TCP 连接的保活，一般很少用。                                                                                |
|SO_LINGER   |struct linger|set|用于设置当 TCP 连接已经关闭，但是未发送数据等待时间。通常设置 SO_LINGER 等待时间为 0，解决大量 TIME_WAIT 状态的问题。|
|SO_OOBINLINE|int          |set|用于设置将“带外数据”作为普通数据流来处理。                                                                         |
|SO_RCVBUF   |int          |set|设置 Socket 接收缓冲区大小。                                                                                         |
|SO_REUSEADDR|int          |set|用于设置在调用 bind() 函数时，重用已经 bind 的 Socket 地址。                                                         |
|SO_SNDBUF   |int          |set|设置 Socket 发送缓冲区大小。                                                                                         |

## 4. 常用选项说明

下来，我们对 Socket 编程中常用的 Socket 选项重点介绍。

### 4.1 SO_REUSEADDR

TCP 连接关闭过程中，主动关闭的一方会处于 TIME_WAIT 状态，要等待 2MSL 时间。而服务器在工作过程中有可能由于配置的改变而要重启，或者是由于程序异常奔溃要重新启动。在这种情况下，如果服务器监听的 Socket 处于 TIME_WAIT 状态，那么调用 bind 方法绑定 Socket 就会失败。如果要等待 2MSL 时间，对于服务器来说是难以接受的。要想解决此问题，需要给监听 Socket 设置 SO_REUSEADDR 选项。

Java 的 java.net.ServerSocket 类提供了 setReuseAddress 方法，可以用以设置 SO_REUSEADDR 选项，如下：

```java
        ServerSocket ss = null;
        try {
            ss = new ServerSocket();
            ss.setReuseAddress(true);
            ss.bind(new InetSocketAddress(8022));
            ss.accept();
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
```

> * 注意：
> * SO_REUSEADDR 选项需要在 bind 方法调用前设置，所以要创建一个未绑定的 ServerSocket 对象，然后手动执行 bind 操作。

### 4.2 SO_KEEPALIVE

SO_KEEPALIVE 是协议栈提供的一种连接保活机制，一般是用在 TCP 协议中。主要目的是当通信双方长时间没有数据交互，然而 Socket 还没有被关闭，协议栈会向对方发送一个 Heartbeat 消息期望对方回复一个 ACK，如果对方能回复说明连接是正常的，如果对方不能回复，尝试几次以后就会关闭连接。系统保活的时间一般是 2 小时。

Java 的 java.net.Socket 类提供了 setKeepAlive 方法，可以用以设置 SO_KEEPALIVE 选项，如下：

```java
    sock.setKeepAlive(true);
```

### 4.3 SO_LINGER

SO_LINGER 是用来设置“连接关闭以后，未发送完的数据包还可以在协议栈逗留的时间”。java.net.Socket 提供了 setSoLinger 方法可以设置 SO_LINGER 选项。原型如下：

```java
public void setSoLinger(boolean on, int linger) throws SocketException
```

* 如果设置 on 为 false，则该选项的值被忽略，协议栈会采用默认行为。close 调用会立即返回给调用者，协议栈会尽可能将 Socket 发送缓冲区未发送的数据发送完成。

* 如果设置 on 为 true，但是 linger 为 0，当你调用了 close() 方法以后，协议栈将丢弃保留在 Socket 发送缓冲区中未发送完的数据，然后向对方发送一个 RST。这样连接很快会被关闭，不会进入 TIME_WAIT 状态，这也是一个避免“由于大量 TIME_WAIT 状态的 Socket 导致连接失败“的解决办法。

* 如果设置 on 为 true ，但是 linger 的取值大于 0，当你调用了 close() 方法以后，如果 Socket 发送缓冲区还有未发送完的数据，那么系统会等待一个指定的时间，close() 才返回。注意，这种情况下 close() 方法返回，并不能保证 Socket 发送缓冲区中未发送的数据被成功发送完。

> * 注意：
> * 参数 linger 的单位是**秒**。

```java
sock.setSoLinger(true, 20);
```

### 4.4 SO_RCVBUF

SO_RCVBUF 很好理解，用于设置 Socket 的接收缓冲区大小。TCP 一般不需要设置，UDP 可能需要设置。java.net.Socket 类提供了 setReceiveBufferSize 方法可以设置接收缓冲区的大小。

```java
sock.setReceiveBufferSize(16384);
```

### 4.5 SO_SNDBUF

SO_SNDBUF 也很好理解，用于设置 Socket 的发送缓冲区大小。一般不需要设置，采用系统默认大小即可。java.net.Socket 类提供了 setSendBufferSize 方法可以设置发送缓冲区的大小。

```java
sock.setSendBufferSize(16384);
```

### 4.6 SO_OOBINLINE

SO_OOBINLINE 用于设置将“带外数据”作为普通数据流来处理。java.net.Socket 类提供了 setOOBInline 方法可以设置 SO_OOBINLINE 选项。

```java
sock.setOOBInline(true);
```

### 4.7 TCP_NODELAY

TCP_NODELAY 用于关闭 Nagle 算法，一般是用在实时性要求比较高的场景。java.net.Socket 提供了 setTcpNoDelay 方法用于设置 TCP_NODELAY 选项。

```java
 sock.setTcpNoDelay(true);
```

## 5 小结

本节重点是介绍在 java 中设置常用 Socket 选项的方法。当然，我们是从 Linux 系统本身提供的 Socket 选项开始的，我们也介绍了在 linux 系统中如何查找 Socket 选项的方法。了解操作系统对 Socket 选项的支持，可以让你形成一个完整的认识。

文中列出了常用 Socket 选项的应用场景。SO_REUSEADDR 是服务器必须要设置的一个选项，也只有服务器才需要此功能。TCP_NODELAY 是在开发实时性要求很高的程序时，必须要设置的，比如音视频通信系统。

SO_LINGER 是在服务器端解决“由于 TIME_WAIT 过多，导致连接失败的问题”时的一个常用方法。其他选项，可以根据需要选择是否开启。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
