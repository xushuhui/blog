---
title: Java从零开始（114）JVM 参数：配置堆空间与栈空间
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# JVM 参数：配置堆空间与栈空间


## 1. 前言

本节内容主要是学习 JVM 配置堆空间与栈空间的常用参数配置，堆空间和栈空间这两块内存区域是非常重要的运行时数据存放区，掌握堆空间与栈空间的参数配置，在实际工作中非常重要。本节主要知识点如下：

* 理解并掌握配置堆空间的参数 -Xms 和 -Xmx，并配和跟踪垃圾回收参数 -XX:+PrintGCDetails 验证堆空间是否配置成功，为本节重点内容；
* 理解并掌握配置年轻代的参数 -Xmn，并配和跟踪垃圾回收参数 -XX:+PrintGCDetails 验证堆空间是否配置成功，为本节重点内容；
* 理解并掌握配置元空间的参数 -XX:MetaspaceSize 和 -XX:MaxMetaspaceSize，并配和跟踪垃圾回收参数 -XX:+PrintGCDetails 验证堆空间是否配置成功，为本节重点内容；
* 理解并掌握配置栈空间的参数 -Xss，为本节次重点内容；

JVM 配置堆空间与栈空间的常用参数，是非常重要的知识点，需要在理解的基础上，并掌握参数的使用方法。堆空间更加详细的内部结构会在后续的课程中做专门讲解，此处先掌握配置方法即可。

## 2. 示例代码准备

本节主要是为了学习配置堆空间和栈空间的参数，因此不需要像跟踪垃圾回收那样手动调用 gc 方法，也不需要像跟踪类的加载与卸载那样建立一个 ArrayList 来观察 ArrayList 类的加载。所以此处的实例代码非常简单，随意打印一行字符串即可，我们主要的关注点在配置完堆空间和栈空间之后是否生效。

**示例**：

```java
public class HeapAndStackParamsDemo {
	 public static void main(String[] args) {
		 System.out.println("Heap and Stack!");
	 }
}
```

## 3. -Xms 和 -Xmx 参数

**参数作用**：

* **-Xms**：设置堆的初始空间大小；
* **-Xmx**：设置堆的最大空间大小。

> **Tips**：多数情况下，这两个参数是配合使用的，设置完初始空间大小后，为了对堆空间的最大值有一个把控，还需要设置下堆空间的最大值。

**场景设置**：设置堆的初始空间大小为 10 M，设置堆的最大空间大小为 20 M。（此处设置的空间大小为实验数据，具体值的设置，需要根据不同项目的实际情况而定。）

* **步骤 1**：在 VM Options 中配置参数 -Xms10m -Xmx20m -XX:+PrintGCDetails 并保存；

> **Tips**：为了验证参数设置是否成功，我们需要配合使用 - XX:+PrintGCDetails 来获取堆空间的空间大小，因此此处的参数需要添加上 - XX:+PrintGCDetails，此处仅为验证，正常情况下，堆空间的设置是单独使用的，如： -Xms10m -Xmx20m。

* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
Heap
 PSYoungGen      total 2560K, used 2012K [0x00000000ffd00000, 0x0000000100000000, 0x0000000100000000)
  eden space 2048K, 98% used [0x00000000ffd00000,0x00000000ffef7388,0x00000000fff00000)
  from space 512K, 0% used [0x00000000fff80000,0x00000000fff80000,0x0000000100000000)
  to   space 512K, 0% used [0x00000000fff00000,0x00000000fff00000,0x00000000fff80000)

 ParOldGen       total 7168K, used 0K [0x00000000ff600000, 0x00000000ffd00000, 0x00000000ffd00000)
  object space 7168K, 0% used [0x00000000ff600000,0x00000000ff600000,0x00000000ffd00000)

 Metaspace       used 3354K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 367K, capacity 388K, committed 512K, reserved 1048576K
```

**结果分析**：我们实验场景的设计中，设置了堆空间的大小初始化为 10 M，那么换算成 Kb 为 10 M = 10240 Kb。

> **Tips**：我们从打印的结果中看到了三部分内存，PSYoungGen （年轻代），ParOldGen（老年代），Metaspace （元空间）。从 JDK1.8 开始，Metaspace （元空间）不属于堆空间，目前我使用的 JDK 大版本号为 1.8 ，因此对于堆空间的初始化大小 10 M，应该只分配给了 PSYoungGen （年轻代）和 ParOldGen（老年代）。

**提出问题**：我们来进行下计算，（PSYoungGen total）2560K + （ParOldGen total）7168K = 9728 K。为什么不等于 10240 K? 是什么原因呢？

**问题解决**：其实是因为这里的 total 指的是可用内存，我们来看下 PSYoungGen （年轻代）的全部日志：

```java
PSYoungGen      total 2560K, used 2012K [0x00000000ffd00000, 0x0000000100000000, 0x0000000100000000)
  eden space 2048K, 98% used [0x00000000ffd00000,0x00000000ffef7388,0x00000000fff00000)
  from space 512K, 0% used [0x00000000fff80000,0x00000000fff80000,0x0000000100000000)
  to   space 512K, 0% used [0x00000000fff00000,0x00000000fff00000,0x00000000fff80000)
```

我们可以看到，PSYoungGen（年轻代）包含了 eden 区以及 from space 和 to space 两个区域，同一时间，from space 和 to space 只有一个区域是可以用的。所以分配给 PSYoungGen（年轻代）的总内存是 2560 K+ 512 K= 3072 K。

我们再从新做下计算，（PSYoungGen total）3072 K+ （ParOldGen total）7168K = 10240 K。到此，证明了参数设置的有效性。

## 4. -Xmn 参数

**参数作用**：专门的设置年轻代 PSYoungGen 大小的参数。

**场景设置**：为了更好的理解并掌握 -Xmn 参数，我们沿用上一知识点的 -Xms10m -Xmx20m -XX:+PrintGCDetails 参数，在此参数的基础上，添加 -Xmn5m ，单独设置年轻代 PSYoungGen 的大小为 5m。

> Tips：前文讲解过，堆空间大小 = 年轻代空间大小 + 老年代空间大小，此处设置堆空间初始大小为 10m，年轻代大小为 5m， 那么通过简单的计算，老年代的空间大小为 10m - 5m = 5m。我们继续来看实验步骤和结果验证。

* **步骤 1**：在 VM Options 中配置参数 -Xms10m -Xmx20m -Xmn5m -XX:+PrintGCDetails 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
Heap
 PSYoungGen      total 4608K, used 2142K [0x00000000ffb00000, 0x0000000100000000, 0x0000000100000000)
  eden space 4096K, 52% used [0x00000000ffb00000,0x00000000ffd179c0,0x00000000fff00000)
  from space 512K, 0% used [0x00000000fff80000,0x00000000fff80000,0x0000000100000000)
  to   space 512K, 0% used [0x00000000fff00000,0x00000000fff00000,0x00000000fff80000)

 ParOldGen       total 5120K, used 0K [0x00000000ff600000, 0x00000000ffb00000, 0x00000000ffb00000)
  object space 5120K, 0% used [0x00000000ff600000,0x00000000ff600000,0x00000000ffb00000)

 Metaspace       used 3441K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 374K, capacity 388K, committed 512K, reserved 1048576K
```

**结果分析**：我们主要来关注下 PSYoungGen（年轻代）的大小，看是否为 5m，换算成 Kb 为 5120 Kb。前文提到过，total 仅代表可用内存，而同一时间 from space 和 to space 只有一个是可用的，所以 PSYoungGen（年轻代）总内存的大小为 4608K + 512K = 5120K = 5m。

## 5. -XX:MetaspaceSize 和 -XX:MaxMetaspaceSize 参数

> **Tips**：在 JDK 1.8 之前，所有加载的类信息都放在永久代中。但在 JDK1.8 之时，永久代被移除，取而代之的是元空间（Metaspace）。一些版本比较低的教程或者论坛，经常会在忽略 JDK 版本的前提下谈永久代或者元空间，不要被此类教程迷惑，此处同学要特别注意。

**参数作用**：

* **-XX:MetaspaceSize** ：元空间发生 GC 的初始阈值；

> **Tips**：-XX:MetaspaceSize 这个参数并非设置元空间初始大小，而是设置的发生 GC 的初始阈值。举例来说，如果设置 -XX:MetaspaceSize 为 10m，那么当元空间的数据存储量到达 10m 时，就会发生 GC。

* **-XX:MaxMetaspaceSize** ：设置元空间的最大空间大小。

**场景设置**：设置元空间发生 GC 的初始阈值的大小为 10 M，设置元空间的最大空间大小为 20 M。（此处设置的空间大小为实验数据，具体值的设置，需要根据不同项目的实际情况而定。）

我们通过两步来进行验证：

* **步骤 1**：在 VM Options 中配置参数 -XX:MetaspaceSize=10m -XX:MaxMetaspaceSize=20m -XX:+PrintGCDetails 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
Heap
 PSYoungGen      total 76288K, used 5244K [0x000000076b400000, 0x0000000770900000, 0x00000007c0000000)
  eden space 65536K, 8% used [0x000000076b400000,0x000000076b91f0d8,0x000000076f400000)
  from space 10752K, 0% used [0x000000076fe80000,0x000000076fe80000,0x0000000770900000)
  to   space 10752K, 0% used [0x000000076f400000,0x000000076f400000,0x000000076fe80000)

 ParOldGen       total 175104K, used 0K [0x00000006c1c00000, 0x00000006cc700000, 0x000000076b400000)
  object space 175104K, 0% used [0x00000006c1c00000,0x00000006c1c00000,0x00000006cc700000)

 Metaspace       used 3392K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 368K, capacity 388K, committed 512K, reserved 1048576K
```

**结果分析**：从上面的执行结果可以看到，Metaspace 空间的初始大小为 3392K ，并不是我们设置的 10M。那是因为 -XX:MetaspaceSize 设置的是元空间发生 GC 的初始阈值。当达到这个值时，元空间发生 GC 操作。如果不进行设置，这个值默认是 20.8M。

而 -XX:MaxMetaspaceSize 则是设置元空间的最大值，如果不手动设置，默认基本是机器的物理内存大小。虽然可以不设置，但还是建议设置一下，因为如果一直不断膨胀，那么 JVM 进程可能会被 OS kill 掉。

## 6. -Xss 参数

**参数作用**：设置单个线程栈大小，一般默认 512 - 1024kb。

> **Tips**：由于单个线程栈大小跟操作系统和 JDK 版本都有关系，因此其默认大小是一个范围值， 512 - 1024kb。在平时工作中，-Xss 参数使用到的场景是非常少的，因为单个线程的栈空间大小使用默认的 512 - 1024kb 就能够满足需求。

如果在某些个别场景下，单个线程的栈空间发生内存溢出，多数情况是由于迭代的深度达到了栈的最大深度，导致内存溢出。这种异常情况，多数会选择优化方法，并不是立刻提升栈空间大小，因为盲目提升栈空间大小，是一种资源浪费。

-Xss 参数的使用为本节课程的次重点，学习者只要了解并掌握该参数的作用即可，万一工作中碰到设置栈空间大小的场景，也不至于束手无措。

## 7. 小结

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxkivoaj60jg09n40c02)

本小节的重点内容即我们所讲述的几种设置堆空间的参数，学习者需要对这些常用参数的意义以及使用方式进行掌握，在实际工作中，使用频次非常高。对于栈空间的参数设置，为本节次重点，因为使用场景比较少，学习者只要有此参数的印象即可，做到用时可用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
