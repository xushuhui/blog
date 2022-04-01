---
title: Java从零开始（112）JVM 参数：跟踪垃圾回收
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# JVM 参数：跟踪垃圾回收


## 1. 前言

本节内容主要是学习 JVM 跟踪垃圾回收的常用参数配置，这是工作中跟踪垃圾回收情况时 JVM 中最常用的参数配置，需要重点对本节内容进行学习。本节主要知识点如下：

* 掌握 IntelliJ IDEA 如何配置 JVM 参数，这是我们学习 JVM 参数配置的前提。对于有部分同学使用 Eclipse 作为开发工具的，可以自行搜索配置方式，此处只针对目前使用广泛的 IntelliJ IDEA 工具进行介绍；
* 理解并掌握跟踪垃圾回收的参数 -XX:+PrintGC，为本节重点内容；
* 理解并掌握跟踪垃圾回收的参数 -XX:+PrintGCDetails，为本节重点内容；
* 理解并掌握跟踪垃圾回收的参数 -XX:+PrintHeapAtGC，为本节重点内容；
* 理解并掌握跟踪垃圾回收的参数 -XX:+PrintGCTimeStamps，为本节重点内容。

JVM 跟踪垃圾回收的常用参数配置是使用 JVM 所必须的知识点，通篇皆为重点掌握内容，需要在理解的基础上并掌握参数的使用方法。

## 2. IntelliJ IDEA 配置 JVM 参数

通过开发工具 IntelliJ IDEA 配置 JVM 参数，需要打开 “Run->Edit Configurations” 菜单，然后在 VM Options 中添加相应的 JVM 参数。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxj217pj60jg0f1wk402)

如上图为添加 JVM 参数 -XX:+PrintGC 的图片说明示例，掌握添加参数的方式，是学习具体参数配置的前提。

## 3. 示例代码准备

为了能够更好的体会，垃圾回收参数的执行效果，我们需要准备一段简易的代码，供我们执行使用，并在代码中手动执行垃圾回收操作，触发垃圾回收机制，使我们的参数能够追踪垃圾回收的动作，并打印相应的日志。

**实例**：准备测试代码，创建一个 String 类型的 ArrayList，并在 list 中添加三个元素，分别是 “Hello”，“World”，“！！！”。然后直接手动执行垃圾回收操作。

```java
public class PrintGCParamsDemo {
    public static void main(String[] args) {
        ArrayList<String> list = new ArrayList<String>();
        list.add("Hello");
        list.add("World");
        list.add("!!!");
        System.gc(); //手动执行 gc 垃圾回收
    }
}
```

> **Tips**： 之所以要手动执行 gc 垃圾回收，是因为 JVM 自动的执行垃圾回收是需要一定的条件的，简单的 main 函数是不能够达到触发垃圾回收的临界值的。所以这里手动进行 gc 方法的调用，是为了展示我们的参数锁带来的作用。

## 4. -XX:+PrintGC 参数

**参数作用**：-XX:+PrintGC 参数是垃圾回收跟踪中十分常用的参数。使用这个参数启动 Java 虚拟机后，只要遇到 GC，就会打印日志。

为了更好的理解并掌握 -XX:+PrintGC 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+PrintGC 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
[GC (System.gc())  3933K->792K(251392K), 0.0054898 secs]
[Full GC (System.gc())  792K->730K(251392K), 0.0290579 secs]
```

**结果分析**：由于这是第一次接触 JVM 打印日志，我们按照字段逐一进行分析。

* **GC 与 Full GC**：代表了垃圾回收的类型，后续会有详细的讲解，这里了解日志意义即可；
* **System.gc()**：代表了引发方式，是通过调用 gc 方法进行的垃圾回收；
* **3933K->792K(251392K)**：代表了之前使用了 3933k 的空间，回收之后使用 792k 空间，言外之意这次垃圾回收节省了 3933k - 792k = 3141k 的容量。251392K 代表总容量；
* **792K->730K(251392K)**：分析同上；
* **0.0054898 secs**：代表了垃圾回收的执行时间，以秒为单位。

那么我们通过上边的结果分析，可以非常顺利的解读这两行日志。

## 5. -XX:+PrintGCDetails 参数

**参数作用**：-XX:+PrintGCDetails 参数是垃圾回收跟踪中十分常用的参数，而且日志更加详细。获取比 -XX:+PrintGC 参数更详细的 GC 信息。

为了更好的理解并掌握 -XX:+PrintGCDetails 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+PrintGCDetails 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
[GC (System.gc()) [PSYoungGen: 3933K->792K(76288K)] 3933K->800K(251392K), 0.0034601 secs] [Times: user=0.00 sys=0.00, real=0.00 secs]
[Full GC (System.gc()) [PSYoungGen: 792K->0K(76288K)] [ParOldGen: 8K->730K(175104K)] 800K->730K(251392K), [Metaspace: 3435K->3435K(1056768K)], 0.0217628 secs] [Times: user=0.03 sys=0.00, real=0.02 secs]
```

**结果分析**： 这里只做新增部分的解释。

* **PSYoungGen**：代表了「年轻代」的回收，在这里只需要了解即可，因为后续讲解 JVM 堆内存以及垃圾回收原理时才会涉及到年轻代；
* **ParOldGen**：「老年代」这里只需要了解即可，因为后续讲解 JVM 堆内存以及垃圾回收原理时才会涉及到老年代；
* **Metaspace**：「元空间」，JDK 的低版本也称之为永久代，依然，此处了解即可。

我们看到 -XX:+PrintGCDetails 参数打印了更加详细的日志内容，把空间清理的部分也表达的更加详细了。

## 6. -XX:+PrintHeapAtGC 参数

**参数作用**：-XX:+PrintHeapAtGC 参数是垃圾回收跟踪中，对堆空间进行跟踪时十分常用的参数，可以在每次 GC 前后分别打印堆的信息。注意，是 GC 前后均打印，打印两次。

为了更好的理解并掌握 -XX:+PrintHeapAtGC 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+PrintHeapAtGC 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
{Heap before GC invocations=1 (full 0):
 PSYoungGen      total 76288K, used 3933K [0x000000076b400000, 0x0000000770900000, 0x00000007c0000000)
  eden space 65536K, 6% used [0x000000076b400000,0x000000076b7d7480,0x000000076f400000)
  from space 10752K, 0% used [0x000000076fe80000,0x000000076fe80000,0x0000000770900000)
  to   space 10752K, 0% used [0x000000076f400000,0x000000076f400000,0x000000076fe80000)
 ParOldGen       total 175104K, used 0K [0x00000006c1c00000, 0x00000006cc700000, 0x000000076b400000)
  object space 175104K, 0% used [0x00000006c1c00000,0x00000006c1c00000,0x00000006cc700000)
 Metaspace       used 3420K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 371K, capacity 388K, committed 512K, reserved 1048576K
Heap after GC invocations=1 (full 0):
 PSYoungGen      total 76288K, used 792K [0x000000076b400000, 0x0000000770900000, 0x00000007c0000000)
  eden space 65536K, 0% used [0x000000076b400000,0x000000076b400000,0x000000076f400000)
  from space 10752K, 7% used [0x000000076f400000,0x000000076f4c6030,0x000000076fe80000)
  to   space 10752K, 0% used [0x000000076fe80000,0x000000076fe80000,0x0000000770900000)
 ParOldGen       total 175104K, used 0K [0x00000006c1c00000, 0x00000006cc700000, 0x000000076b400000)
  object space 175104K, 0% used [0x00000006c1c00000,0x00000006c1c00000,0x00000006cc700000)
 Metaspace       used 3420K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 371K, capacity 388K, committed 512K, reserved 1048576K
}
{Heap before GC invocations=2 (full 1):
 PSYoungGen      total 76288K, used 792K [0x000000076b400000, 0x0000000770900000, 0x00000007c0000000)
  eden space 65536K, 0% used [0x000000076b400000,0x000000076b400000,0x000000076f400000)
  from space 10752K, 7% used [0x000000076f400000,0x000000076f4c6030,0x000000076fe80000)
  to   space 10752K, 0% used [0x000000076fe80000,0x000000076fe80000,0x0000000770900000)
 ParOldGen       total 175104K, used 0K [0x00000006c1c00000, 0x00000006cc700000, 0x000000076b400000)
  object space 175104K, 0% used [0x00000006c1c00000,0x00000006c1c00000,0x00000006cc700000)
 Metaspace       used 3420K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 371K, capacity 388K, committed 512K, reserved 1048576K
Heap after GC invocations=2 (full 1):
 PSYoungGen      total 76288K, used 0K [0x000000076b400000, 0x0000000770900000, 0x00000007c0000000)
  eden space 65536K, 0% used [0x000000076b400000,0x000000076b400000,0x000000076f400000)
  from space 10752K, 0% used [0x000000076f400000,0x000000076f400000,0x000000076fe80000)
  to   space 10752K, 0% used [0x000000076fe80000,0x000000076fe80000,0x0000000770900000)
 ParOldGen       total 175104K, used 705K [0x00000006c1c00000, 0x00000006cc700000, 0x000000076b400000)
  object space 175104K, 0% used [0x00000006c1c00000,0x00000006c1cb07a0,0x00000006cc700000)
 Metaspace       used 3420K, capacity 4496K, committed 4864K, reserved 1056768K
  class space    used 371K, capacity 388K, committed 512K, reserved 1048576K
}

```

**结果分析**： 由于这是对「堆空间」的日志打印，在学习完 JVM 堆空间之前，了解即可。对于堆空间的参数跟踪，这里进行了更加细致的打印。从结果来看，在 GC 前后，打印了两次堆空间信息，并且将 PSYoungGen 以及 ParOldGen 进行了更加详细的日志打印。

后续学习完 JVM 堆空间之后，回望这部分知识，会非常的简单，此处也必须要先了解，为后续对堆的学习打下良好的基础。

## 7. -XX:+PrintGCTimeStamps 参数

**参数作用**：会在每次 GC 发生时，额外输出 GC 发生的时间，该输出时间为虚拟机启动后的时间偏移量。需要与 -XX:+PrintGC 或 -XX:+PrintGCDetails 配合使用，单独使用 -XX:+PrintGCTimeStamps 参数是没有效果的。

为了更好的理解并掌握 -XX:+PrintGCTimeStamps 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+PrintGC -XX:+PrintGCTimeStamps 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：

```java
0.247: [GC (System.gc())  3933K->760K(251392K), 0.0114098 secs]
0.259: [Full GC (System.gc())  760K->685K(251392K), 0.0079185 secs]
```

**结果分析**：我们看到了，与 -XX:+PrintGC 参数打印的结果，唯一的区别就是日志开头的 0.247 与 0.259。此处 0.247 与 0.259 表示， JVM 开始运行 0.247 秒后发生了 GC，开始运行 0.259 秒后，发生了 Full GC。

## 8. 小结

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxjgqkdj60jg09cgn802)

本小节的重点内容，即我们所讲述的四个常见的跟踪垃圾回收的参数，学习者，需要对这四个常用参数的意义，以及使用方式进行掌握。由于日志输出中部分内容，涉及到的知识点还没有讲到，但是本节对于这些内容的初步接触，能够有利于学习者，在后续的学习中提升学习的效率，以及理解效率。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
