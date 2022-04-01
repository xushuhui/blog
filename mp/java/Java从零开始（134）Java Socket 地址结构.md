---
title: Java从零开始（134）Java Socket 地址结构
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java Socket 地址结构


## 1. 前言

我们知道计算机网络中连接的设备有很多，比如 PC、手机、打印机、路由器、交换机、网关等，通常把这些网络设备叫做节点（Node）。每一个节点都分配有唯一的 IP 地址，用以标识此设备。IP 地址包含 32 位 IPv4 和 128 位 IPv6 两个版本。由于 IP 地址是一串数字或者是字节序列，对计算机是友好的，但是对我们人类非常不友好，不利于传播、记忆。为此，计算机科学家又开发了一套 DNS 系统，给每一台计算机分配了唯一的、对人类友好的主机名字，通常叫做域名。比如，[www.imooc.com](http://www.imooc.com) 是主站的域名。当然，有的主机会分配多个域名。

人们常说生活没有那么简单，往往是解决了一个老问题，又引出了新问题。当你开发了 DNS 系统以后，我们人类确实方便了，可是域名对计算机来说不方便，计算机更喜欢 IP 地址。这就又需要解决 IP 地址和域名之间相互解析、映射的问题，当然这些问题在 DNS 系统中都得到了妥善的处理。域名解析系统是一个分布式集群系统，是一个树形结构。一次域名解析可能需要经过本地缓存、本地域名服务器、远程域名服务器之间多次交互。

从上面的描述可以看出，IP 地址和域名之间的相互解析是一套非常复杂的机制。好在操作系统将这一套复杂的机制进行了封装，以 API 的形式提供给网络程序员，这样极大的简化了编程的复杂度。

一般操作系统都提供了 C 语言接口 **getaddrinfo** 和 **getnameinfo**，前者的功能是通过域名获取 IP 地址，后者的功能是通过 IP 地址获取域名。

在 Java 平台中，**java.net.InetAddress** 类实现了完整的 IP 地址和域名之间的相互解析机制。

## 2. InetAddress 类的体系结构

**java.net.InetAddress** 类的体系结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnypqw35j60f207aq3202)

各类的功能说明：

* InetAddress 是 Java IP 地址的包装类，也是域名解析的核心类。
* Inet4Address 代表了 IPv4 地址格式的封装，一般程序员不需要关心此类。
* Inet6Address 代表了 IPv6 地址格式的封装，一般程序员不需要关心此类。
* InetSocketAddress 是 Socket 地址的封装，它通过私有内部类 InetSocketAddressHolder 间接包装了 **InetAddress** 结构和** 端口号**（Port）。在网络编程中，通常把 Socket 地址叫做 ** Endpoint**，用 <IP, Port> 的组合来表示。

在网络编程中，应用最为频繁的两个类是 InetSocketAddress 和 InetAddress。其中，InetSocketAddress 类对 InetAddress 和 Port 进行了封装，形成了完整的 Socket 地址。而 InetAddress 核心实现就是域名解析和缓存。

InetAddress 类没有 public 构造方法，提供了一组 public static 工厂方法用以创建 InetAddress 实例。接下来，我们重点分析一下 getByName 和 getByAddress 两类方法。

## 3. getByName 方法

InetAddress 提供了两个公有静态方法 getByName 和 getAllByName 来构造 InetAddress 实例，它们的原型如下：

```java
// 创建单个 InetAddress 实例
public static InetAddress getByName(String host) throws UnknownHostException
// 创建多个 InetAddress 实例
public static InetAddress[] getAllByName(String host) throws UnknownHostException
```

这两个方法都会连接域名解析服务器进行域名解析，具体工作原理如下：

* 首先会检查传入参数 host，也就是域名。如果传入参数为 null，那么会返回以 loopback 地址构造的 InetAddress 结构。
* 如果输入参数 host 是一个 IP 地址，那么根据 IP 地址是 IPv4 还是 IPv6，分别构造 Inet4Address 或 Inet6Address 结构，并且返回。
* 查询本地 Cache，如果本地 Cache 中已经存在 host 相应的地址，则直接返回。
* 如果本地 Cache 查询失败，则遍历本地注册的 name services。如果有定制的 name services 注册，那么会调用此定制的 name services。如果没有定制的 name services，那么会调用 default name services，最终会调用系统的 getaddrinfo 函数。getaddrinfo 是一个 POSIX 标准函数，一般系统都会实现。

getByName 方法的应用非常简单，示例如下：

```java
public static void testInetAddressByName(String host){
        try {
            InetAddress addr = InetAddress.getByName(host);
            System.out.println("getByName addr=" + addr.toString());

            InetAddress[] addrs = InetAddress.getAllByName(host);
            for (InetAddress a: addrs){
                System.out.println("getAllByName addr=" + a.toString());
            }
        } catch (UnknownHostException e) {
            e.printStackTrace();
        }
    }
```

测试 [wwww.imooc.com](http://wwww.imooc.com) 域名，执行结果如下：

```java
getByName addr=www.imooc.com/115.182.41.103
getAllByName addr=www.imooc.com/115.182.41.103
getAllByName addr=www.imooc.com/117.121.101.144
getAllByName addr=www.imooc.com/115.182.41.180
getAllByName addr=www.imooc.com/117.121.101.40
getAllByName addr=www.imooc.com/117.121.101.134
getAllByName addr=www.imooc.com/115.182.41.163
```

需要注意的是 getByName 方法会抛出 UnknownHostException 异常，需要捕获。

## 4. getByAddress 方法

如果你有明确的 IP 地址，并不需要进行域名解析，可以调用 InetAddress 提供的另一组工厂方法 getByAddress，方法原型如下：

```java
public static InetAddress getByAddress(byte[] addr) throws UnknownHostException

public static InetAddress getByAddress(String host, byte[] addr) throws UnknownHostException
```

这是两个重载的 public static 方法，功能都类似：

* 第一个重载的 getByAddress 方法提供一个参数，即用 byte [] 类型的数组表示的 IP 地址。
* 第二个重载的 getByAddress 方法提供两个参数，用 String 类型表示的域名（host），和用 byte [] 类型的数组表示的 IP 地址。
* 二者都不进行域名解析，只是根据输入参数构造 InetAddress 实例。
* 接收 host 输入参数的 getByAddress 方法不保证**域名和 IP 地址**的对应关系，也不保证域名是否可以访问。

getByAddress 方法应用的示例代码如下：

```java
public static void testInetAddressByAddr()
    {
        byte[] ips = new byte[]{ (byte)192, (byte)168,1,101};
        try {
            InetAddress addr = InetAddress.getByAddress(ips);
            System.out.println("getByAddress addr=" + addr.toString());

            InetAddress addr2 = InetAddress.getByAddress("www.example.com", ips);
            System.out.println("getByAddress with host addr=" + addr2.toString());
        } catch (UnknownHostException e) {
            e.printStackTrace();
        }
    }
```

我们输入 192.168.1.101，执行结果如下：

```java
getByAddress addr=/192.168.1.101
getByAddress with host addr=www.example.com/192.168.1.101
```

## 5. InetAddress 的 Cache 策略

由于域名解析需要客户端和域名服务器经过很多次交互，一般都比较耗费时间，所以 InetAddress 提供了 Cache 机制。这样，当客户程序调用 getByName 解析域名的时候，首先是从 Cache 中查找，这样可以极大提高域名解析的效率。

域名绑定的 IP 地址可能会发生变化，所以 Cache 中存储的 IP 地址也是有生命周期的。Java 提供了两个全局参数可以用来配置 Cache 的有效时间。

* networkaddress.cache.ttl

成功解析的域名在 Cache 中的存活时间。

* networkaddress.cache.negative.ttl

解析失败的域名在 Cache 中的存活时间。

实际上除了 Java 本地有 Cache 机制，域名解析服务器也是有 Cache 机制的，目的都是相同的。

## 6. 小结

InetAddress 类在网络编程中的应用是非常频繁的，了解域名解析机制有利于我们更好的应用此类的功能。在实际产品应用中都是通过 getByName 方法构造 InetAddress 实例的，尽量避免通过 getByAddress 方法构造 InetAddress 实例。这样可以提高程序的维护性。当然，在实验室的内网环境中进行开发测试，往往采用的是私有 IP 地址，这时可以通过 getByAddress 方法来构造 InetAddress 实例。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
