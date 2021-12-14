---
title: Java 从零开始（62）原子操作之 AtomicInteger
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 原子操作之 AtomicInteger

## 1. 前言

从本节开始，正式带领大家认识 Java 并发工具类，今天为大家介绍原子操作之 AtomicInteger。此工具位于 java.util.concurrent.atomic 包中。

本节先解释什么是原子操作，接着介绍 AtomicInteger 工具类的最基本用法，有了这些基本认识之后，给出 AtomicInteger 工具最常用的场合说明，然后通过简单的编码实现一个实际案例，让大家有一个理性的认识，最后带领大家熟悉 AtomicInteger 最常用的一些编程方法，进一步加深对 AtomicInteger 工具类的理解。

AtomicInteger 工具类本身使用很简单，重点是对常用编程方法的准确理解。

下面我们正式开始介绍吧。

## 2. 概念解释

什么是原子操作呢？**所谓原子操作，就是一个独立且不可分割的操作。**

AtomicInteger 工具类提供了对整数操作的原子封装。为什么要对整数操作进行原子封装呢？

在 java 中，当我们在多线程情况下，对一个整型变量做加减操作时，如果不加任何的多线程并发控制，大概率会出现线程安全问题，也就是说当多线程同时操作一个整型变量的增减时，会出现运算结果错误的问题。AtomicInteger 工具类就是为了**简化整型变量的同步处理**而诞生的。

大家记住，在多线程并发下，所有不是原子性的操作但需要保证原子性时，都需要进行原子操作处理，否则会出现线程安全问题。

概念已经了解了，那么 AtomicInteger 工具类怎么用呢？别急，最基本的用法请看下面的描述。

## 3. 基本用法

```java
// 首先创建一个 AtomicInteger 对象
AtomicInteger atomicInteger = new AtomicInteger();
// 在操作之前先赋值，如果不显式赋值则值默认为 0 ，就像 int 型变量使用前做初始化赋值一样。
atomicInteger.set(1000);
// 之后可以调用各种方法进行增减操作
...
// 获取当前值
atomicInteger.get();
// 先获取当前值，之后再对原值加 100
atomicInteger.getAndAdd(100)
// 先获取当前值，之后再对原值减 1
atomicInteger.getAndDecrement()
...
```

是不是很简单，AtomicInteger 在我们日常实践中，到底应该应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

AtomicInteger 经常用于多线程操作同一个整型变量时，简化对此变量的线程安全控制的场合。当在研发过程中遇到这些场景时，就可以考虑直接使用 AtomicInteger 工具类辅助实现，完全可以放弃使用 synchronized 关键字做同步控制。

下面我们用 AtomicInteger 工具实现电影院某场次电影票销售的例子。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnywzxdrj60ea09wmzn02)

## 5. 场景案例

```java
import java.util.concurrent.atomic.AtomicInteger;

public class AtomicIntegerTest {

    // 首先创建一个 AtomicInteger 对象
    // 代表《神龙晶晶兽》电影上午场次当前可售的总票数 10 张
    private static AtomicInteger currentTicketCount = new AtomicInteger(10);
    // 主程序
    public static void main(String[] args) {
        // 定义 3 个售票窗口
        for(int i=1; i<=3; i++) {
            TicketOffice ticketOffice = new TicketOffice(currentTicketCount, i);
            // 每个售票窗口开始售票
            new Thread(ticketOffice).start();
        }
    }
}
```

在上面的代码中，先创建了一个 AtomicInteger 对象，然后创建了 3 个售票窗口模拟售票动作 ，接下来每个售票窗口如何动作呢，看下面的代码。

```java
import java.util.Random;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * 模拟售票窗口
 */
public class TicketOffice implements Runnable {
    // 当前可售的总票数
    private AtomicInteger currentTicketCount;
    // 窗口名称（编号）
    private String ticketOfficeNo;
	// 售票窗口构造函数
    public TicketOffice(AtomicInteger currentTicketCount, int ticketOfficeNo) {
        this.currentTicketCount = currentTicketCount;
        this.ticketOfficeNo = "第" + ticketOfficeNo + "售票窗口";
    }
	// 模拟售票逻辑
    public void run() {
	    // 模拟不间断的售票工作（生活中有工作时间段控制）
        while (true) {
            // 获取当前可售的总票数，如果没有余票就关闭当前售票窗口结束售票，否则继续售票
            if (currentTicketCount.get() < 1) {
                System.out.println("票已售完，" + ticketOfficeNo + "结束售票");
                return;
            }
            // 模拟售票用时
            try {
                Thread.sleep(new Random().nextInt(1000));
            } catch (Exception e) {}
            // 当总票数减 1 后不为负数时，出票成功
            int ticketIndex = currentTicketCount.decrementAndGet();
            if (ticketIndex >= 0) {
                System.out.println(ticketOfficeNo + "已出票，还剩" + ticketIndex + "张票");
            }
        }
    }
}
```

在 TicketOffice 类中，首先通过 get () 获取了当前可售的总票数，在有余票的情况下继续售票。然后随机休眠代替售票过程，最后使用 decrementAndGet () 尝试出票。我们观察一下运行结果。

```java
第 3 售票窗口已出票，还剩 9 张票
第 1 售票窗口已出票，还剩 8 张票
第 2 售票窗口已出票，还剩 7 张票
第 1 售票窗口已出票，还剩 6 张票
第 3 售票窗口已出票，还剩 5 张票
第 3 售票窗口已出票，还剩 4 张票
第 2 售票窗口已出票，还剩 3 张票
第 1 售票窗口已出票，还剩 2 张票
第 3 售票窗口已出票，还剩 1 张票
第 2 售票窗口已出票，还剩 0 张票
票已售完，第 2 售票窗口结束售票
票已售完，第 1 售票窗口结束售票
票已售完，第 3 售票窗口结束售票
```

在这个案例中，因为存在多个售票窗口同时对一场电影进行售票，如果不对可售票数做并发售票控制，很可能会出现多卖出票的尴尬。例子中没有直接使用 synchronized 关键字做同步控制，而是使用 JDK 封装好的 AtomicInteger 原子工具类实现了并发控制整型变量的操作，是不是很方便呢。

至此，大家对 AtomicInteger 已经有了初步的理解，接下来我们继续丰富对 AtomicInteger 工具类的认识。

## 6. 核心方法介绍

除了上面代码中使用的最基本的 AtomicInteger (int)、AtomicInteger ()、 set () 、get () 和 decrementAndGet () 方法之外，我们还需要掌握其他几组核心方法的使用。下面逐个介绍。

1. getAndAdd (int) 方法与 AddAndGet (int) 方法

第 1 个方法是先获取原值，之后再对原值做增加。注意获取的值是变更之前的值。而第 2 个方法正好相反，是先对原值做增加操作之后再获取更新过的值。

```java
AtomicInteger atomicInteger = new AtomicInteger();
System.out.println(atomicInteger.get());            // 0
System.out.println(atomicInteger.getAndAdd(10));    // 0，获取当前值并加 10
System.out.println(atomicInteger.get());            // 10
System.out.println(atomicInteger.addAndGet(20));    // 30，当前值先加 20 再获取
System.out.println(atomicInteger.get());            // 30
```

1. getAndIncrement () 方法与 incrementAndGet () 方法

第 1 个方法是先获取值，之后再对原值做增 1 操作，注意获取的值是变更之前的值。而第 2 个方法正好相反，是先对原值做增 1 的操作之后再获取更新过的值。

```java
AtomicInteger atomicInteger = new AtomicInteger();
System.out.println(atomicInteger.get());  // 0
System.out.println(atomicInteger.getAndIncrement()); // 0，获取当前值并自增 1
System.out.println(atomicInteger.get());  // 1
System.out.println(atomicInteger.incrementAndGet()); // 2，当前值先自增 1 再获取
System.out.println(atomicInteger.get());  // 2
```

1. compareAndSet(int expect, int update)

原值与 expect 相比较，如果不相等则返回 false 且原有值保持不变，否则返回 true 且原值更新为 update。

```java
AtomicInteger atomicInteger = new AtomicInteger(10);
System.out.println(atomicInteger.get()); // 10
int expect = 12;
int update = 20;
Boolean b =atomicInteger.compareAndSet(expect, update);
System.out.println(b); // 10 不等于 12 不满足期望，所以返回 false，且保持原值不变
System.out.println(atomicInteger.get());
```

## 7. 小结

本节通过一个简单的例子，介绍了 AtomicInteger 的基本用法，另外对一些核心方法做了简单介绍。在这个包下面，还有很多类似的工具类，也是对基本类型原子操作的封装，如 AtomicBoolean、AtomicLong，用法大同小异，希望大家在日常研发中多比较多总结，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
