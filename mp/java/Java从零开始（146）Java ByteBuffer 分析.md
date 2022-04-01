---
title: Java从零开始（146）Java ByteBuffer 分析
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java ByteBuffer 分析


## 1. 前言

基于 java.net.Socket 进行 TCP Socket 编程，我们是通过 java.net.SocketInputStream 和 java.net.SocketOutputStream 实现数据读写，这是基于字节流的数据读写，每次读写固定的几个字节，没有缓冲能力。为了提高 I/O 读写性能，我们引入了 BufferedInputStream & BufferedOutputStream。缓冲流本质是提供了一个较大块内存，实现大块数据读写的能力。Java NIO 提供的 SocketChannel 是基于 ByteBuffer 实现数据读写的，天生具有缓冲能力，毕竟 Java NIO 就是为了解决性能问题的嘛。如果从 SocketChannel 读取数据，必须预先分配好 ByteBuffer；如果要想将数据写入 SocketChannel，必须预先将数据写入 ByteBuffer。

Java NIO 定义了一个关于缓冲的抽象类是 java.nio.Buffer，Buffer 有很多实现子类，ByteBuffer 仅仅是其中一个子类。下来我们就对 java.nio.ByteBuffer 的功能做一个介绍。

## 2. java.nio.Buffer 基本结构

java.nio.Buffer 是一个抽象类，定义了 Buffer 的基本结构。Buffer 存放的内容是 Java 的基本类型，针对每一个基本类型，都有一个实现类。比如，LongBuffer，IntBuffer，ByteBuffer 等。Buffer 是一个线性结构，内部实现是一个数组，是有大小限制的。java.nio.Buffer 中定义了几个非常重要的属性，声明如下：

```java
private int mark = -1;
private int position = 0;
private int limit;
private int capacity;
```

* capacity，表示 ByteBuffer 的容量，即 Buffer 总大小。
* position，表示 Buffer 当前数据读写的位置。position 的取值不会是负数，也不会超过 limit 的取值。
* limit，表示 Buffer 读写操作的结束位置。limit 的取值不会是负数，也不会超过 capacity 的取值。
* mark，用于用户自定义的标记位置。

## 3. java.nio.ByteBuffer 介绍

java.nio.ByteBuffer 中存储的内容是 java 的基本类型 byte。一个非空的 ByteBuffer 的结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnytrr0sj60mj0g476102)

ByteBuffer 内部维护了一个 byte 数组，position、limit、capacity 都是数组的下标。

### 3.1 ByteBuffer 创建

ByteBuffer 本身也是一个抽象类，没有提供 public 构造方法，创建 ByteBuffer 必须通过它的工厂方法完成。如果是要创建一个新的、空的 ByteBuffer，可以调用它的 allocate 方法。如果是想把一个已知的 byte 数组包装到 ByteBuffer 中，而不是重新分配内存空间，可以调用它的 wrap 方法。声明如下：

```java
public static ByteBuffer allocate(int capacity)
public static ByteBuffer wrap(byte[] array, int offset, int length)
public static ByteBuffer wrap(byte[] array)
```

allocate 方法包含一个 capacity 参数，需要用户指定 ByteBuffer 的大小。 wrap 方法有两个重载实现，都需要传入数组的引用，另外一个还可以指定一个 offset 和 length。示例代码如下：

```java
ByteBuffer newBuffer = ByteBuffer.allocate(1024);

byte[] tmpByteArray = new byte[512];
ByteBuffer wrapBuffer = ByteBuffer.wrap(tmpByteArray);
```

新创建的 ByteBuffer，capacity、position、limit 的取值如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyubu5pj60as0fwwf102)

position 指向数组开头，capacity 和 limit 都指向数组结尾。

### 3.2 向 ByteBuffer 写入数据

ByteBuffer 提供了一组重载的、写入数据的方法，你可以写入单个 byte，也可以写入一个 byte 数组。声明如下：

```java
   public abstract ByteBuffer put(byte b);
   public final ByteBuffer put(byte[] src)
```

示例代码如下：

```java
tmpByteArray[0] = (byte)0x11;
tmpByteArray[1] = (byte)0x22;
newBuffer.put((byte)0xAA);
newBuffer.put((byte)0xBB);
newBuffer.put(tmpByteArray, 0, 2);
```

经过 put 操作后， newBuffer 的 capacity、position、limit 的值如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyuqo70j60as0fwaax02)

当我们向 newBuffer 中 put 4 个 byte 类型的数据后，position 指向 4，capacity 和 limit 没有变化。

### 3.3 ByteBuffer 的 flip 方法

如果要从一个写入数据的 ByteBuffer 读取数据，需要调用 ByteBuffer 的 flip 方法。flip 方法会改变 capacity、position、limit 的值。示例代码：

```java
newBuffer.flip();
```

调用完 flip 方法后，newBuffer 的 capacity、position、limit 的值如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyv1igej60as0fwdgm02)

### 3.4 从 ByteBuffer 读取数据

ByteBuffer 提供了一组重载的、读取数据的方法，你可以读取单个 byte，也可以读取一个 byte 数组。声明如下：

```java
   public abstract byte get();
   public ByteBuffer get(byte[] dst, int offset, int length)
```

示例代码如下：

```java
newBuffer.get();
newBuffer.get(tmpByteArray, 0, 2);
```

经过 gett 操作后， newBuffer 的 capacity、position、limit 的值如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyvyik6j60as0fwgmk02)

### 3.5 ByteBuffer 的 rewind 方法

ByteBuffer 的 rewind 方法仅仅是将 position 设置为 0。

### 3.6 ByteBuffer 的 compat 方法

ByteBuffer 的 compat 方法将 position 所指向的、长度为 limit - position 的数据拷贝到数组的开头，然后重新设定 position 和 limit 的值。compat 方法是非常有用的，比如你在解析读取的报文的时候，如果消息不完整，你可以调用 compat 方法，然后继续接收。

## 4. 总结

本小节主要是介绍了 ByteBuffer 的常见用法。要想熟练的从事 Java NIO 编程，首先必须理解 ByteBuffer 的原理和基本用法。ByteBuffer 的本质就是维护了一个数组，通过 position、limit、capacity 来维护读写的位置。另外，网络编程是需要关注字节序的，ByteBuffer 默认是网络字节序，你可以调用它的 order 方法设置字节序。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
