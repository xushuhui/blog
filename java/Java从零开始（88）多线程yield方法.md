---
title: Java从零开始（88）多线程yield方法
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 多线程 yield 方法


## 1. 前言

本节对 yield 方法进行深入的剖析，主要内容点如下：

* 首先要了解什么是 CPU 执行权，因为 yield 方法与 CPU 执行权息息相关；
* 了解 yield 方法的作用，要明确 yield 方法的使用所带来的运行效果；
* 了解什么是 native 方法，由于 yield 方法是 native 方法的调用，在学习 yield 方法之前，要了解什么是 native 方法；
* 掌握 yield 方法如何使用，这是本节知识点的重中之重，一定要着重学习；
* 了解 yield 方法和 sleep 方法的区别，进行对比记忆，更有助于掌握该方法的独有特性。

## 2. 什么是 CPU 执行权

我们知道操作系统是为每个线程分配一个时间片来占有 CPU 的，正常情况下当一个线程把分配给自己的时间片使用完后，线程调度器才会进行下一轮的线程调度，这里所说的 “自己占有的时间片” 即 CPU 分配给线程的执行权。

**那进一步进行探究，何为让出 CPU 执行权呢？**

当一个线程通过某种可行的方式向操作系统提出让出 CPU 执行权时，就是在告诉线程调度器自己占有的时间片中还没有使用完的部分自己不想使用了，主动放弃剩余的时间片，并在合适的情况下，重新获取新的执行时间片。

## 3. yield 方法的作用

**方法介绍**：Thread 类中有一个静态的 yield 方法，当一个线程调用 yield 方法时，实际就是在暗示线程调度器当前线程请求让出自己的 CPU 使用权。

```java
public static native void yield();
```

> **Tips**：从这个源码中我们能够看到如下两点要点：
>
> 1. yield 方法是一个静态方法，静态方法的特点是可以由类直接进行调用，而不需要进行对象 new 的创建，调用方式为 Thread.yield ()。
> 2. 该方法除了被 static 修饰，还被 **native** 修饰，那么进入主题，什么是 native 方法呢？我们继续来看下文的讲解。

抽象地讲，一个 Native Method 就是一个 Java 调用的非 Java 代码的接口。一个 Native Method 是这样一个 Java 的方法：该方法的实现由非 java 语言实现。

简单的来说，native 方法就是我们自己电脑的方法接口，比如 Windows 电脑会提供一个 yield 方法，Linux 系统的电脑也同样会提供一个 yield 方法，本地方法，可以理解为操作调用操作系统的方法接口。

**作用**：暂停当前正在执行的线程对象（及放弃当前拥有的 cup 资源），并执行其他线程。yield () 做的是让当前运行线程回到可运行状态，以允许具有相同优先级的其他线程获得运行机会。

**目的**：yield 即 “谦让”，使用 yield () 的目的是让具有相同优先级的线程之间能适当的轮转执行。但是，实际中无法保证 yield () 达到谦让目的，因为放弃 CPU 执行权的线程还有可能被线程调度程序再次选中。

## 4. yield 方法如何使用

为了更好的了解 yield 方法的使用，我们首先来设计一个使用的场景。

**场景设计**：

* 创建一个线程，线程名为 threadOne；
* 打印一个数，该数的值为从 1 加到 10000000 的和；
* 不使用 yield 方法正常执行，记录总的执行时间；
* 加入 yield 方法，再次执行程序；
* 再次记录总执行时间。

**期望结果**： 未加入 yield 方法之前打印的时间 < 加入 yield 方法之后的打印时间。因为 yield 方法在执行过程中会放弃 CPU 执行权并从新获取新的 CPU 执行权。

**代码实现 - 正常执行**：

```java
public class DemoTest extends Thread {
    @Override
    public void run() {
        Long start = System.currentTimeMillis();
        int count = 0;
        for (int i = 1; i <= 10000000; i++) {
             count = count + i;
        }
        Long end = System.currentTimeMillis();
        System.out.println("总执行时间： "+ (end-start) + " 毫秒, 结果 count = " + count);
    }

    public static void main(String[] args) throws InterruptedException {
        DemoTest threadOne = new DemoTest();
        threadOne. start();
    }
}
```

**执行结果验证**：

```java
总执行时间： 6 毫秒.
```

**代码实现 - yield 执行**：

```java
public class DemoTest extends Thread {
    @Override
    public void run() {
        Long start = System.currentTimeMillis();
        int count = 0;
        for (int i = 1; i <= 10000000; i++) {
             count = count + i;
             this.yield(); // 加入 yield 方法
        }
        Long end = System.currentTimeMillis();
        System.out.println("总执行时间： "+ (end-start) + " 毫秒. ");
    }

    public static void main(String[] args) throws InterruptedException {
        DemoTest threadOne = new DemoTest();
        threadOne. start();
    }
}
```

**执行结果验证**：

```java
总执行时间： 5377 毫秒.
```

从执行的结果来看，与我们对 yield 方法的理解和分析完全相符，请同学也进行代码的编写和运行，加深学习印象。当加入 yield 方法执行时，线程会放弃 CPU 的执行权，并等待再次获取新的执行权，所以执行时间上会更加的长。

## 5. yield 方法和 sleep 方法的区别

* sleep () 方法给其他线程运行机会时不考虑线程的优先级，因此会给低优先级的线程以运行的机会；
* yield () 方法只会给相同优先级或更高优先级的线程以运行的机会；
* 线程执行 sleep () 方法后转入阻塞 (blocked) 状态，而执行 yield () 方法后转入就绪 (ready) 状态；
* sleep () 方法声明会抛出 InterruptedException, 而 yield () 方法没有声明任何异常；
* sleep () 方法比 yield () 方法具有更好的移植性 （跟操作系统 CPU 调度相关）。

## 6. 小结

在实际的开发场景中，yield 方法的使用场景比较少，但是对于并发原理知识的学习过程，对 yield 方法的了解非常重要，有助于同学了解不同状态下的线程的不同状态。

本节要重点掌握 yield 方法的作用以及如何使用 yield 方法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
