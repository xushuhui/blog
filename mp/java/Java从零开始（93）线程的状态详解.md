---
title: Java从零开始（93）线程的状态详解
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 线程的状态详解


## 1. 前言

本节内容主要是对多线程的 6 种状态进行详细讲解，具体内容点如下：

* 抛开语言，谈操作系统的线程的生命周期及线程 5 种状态，这是我们学习 Java 多线程 6 种状态的基础；
* 掌握 Java 的线程生命周期及 6 种线程状态，是我们本节课程的重点内容；
* 理解 Java 线程 6 种状态的定义，并且通过代码实例进行实战演练，更深入的掌握线程的 6 种不同状态，是我们本节内容的核心知识；
* 掌握 Java 线程不同状态之间的转变关系，更好地理解线程的不同状态，是我们本节课程的重点。

## 2. 操作系统线程的生命周期

**定义**：当线程被创建并启动以后，它既不是一启动就进入了执行状态，也不是一直处于执行状态。在线程的生命周期中，它要经过新建 (New)、就绪（Runnable）、运行（Running）、阻塞 (Blocked)，和死亡 (Dead) 5 种状态。

从线程的新建 (New) 到死亡 (Dead)，就是线程的整个生命周期。

下面我们分别对 5 种不同的状态进行概念解析。

**新建 (New)**：操作系统在进程中新建一条线程，此时线程是初始化状态。

**就绪 (Runnable)**：就绪状态，可以理解为随时待命状态，一切已准备就绪，随时等待运行命令。

**运行 (Running)**：CPU 进行核心调度，对已就绪状态的线程进行任务分配，接到调度命令，进入线程运行状态。

**阻塞 (Blocked)**：线程锁导致的线程阻塞状态。共享内存区域的共享文件，当有两个或两个以上的线程进行非读操作时，只允许一个线程进行操作，其他线程在第一个线程未释放锁之前不可进入操作，此时进入的一个线程是运行状态，其他线程为阻塞状态。

**死亡 (Dead)**：线程工作结束，被操作系统回收。

## 3. Java 的线程的生命周期及状态

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz7aoi9j60jg09xgof02)

**定义**： 在 Java 线程的生命周期中，它要经过新建（New），运行（Running），阻塞（Blocked），等待（Waiting），超时等待（Timed_Waiting）和终止状态（Terminal）6 种状态。

从线程的新建（New）到终止状态（Terminal），就是线程的整个生命周期。

> **Tips** ：与操作系统相比， Java 线程是否少了 “就绪” 状态 ？其实 Java 线程依然有就绪状态，只不过 Java 线程将 “就绪（Runnable）" 和 “运行（Running）” 两种状态统一归结为 “运行（Running）” 状态。

我们来看下 Java 线程的 6 种状态的概念。

**新建 (New)**：实现 Runnable 接口或者继承 Thead 类可以得到一个线程类，new 一个实例出来，线程就进入了初始状态。

**运行 (Running)**：线程调度程序从可运行池中选择一个线程作为当前线程时线程所处的状态。这也是线程进入运行状态的唯一方式。

**阻塞 (Blocked)**：阻塞状态是线程在进入 synchronized 关键字修饰的方法或者代码块时，由于其他线程正在执行，不能够进入方法或者代码块而被阻塞的一种状态。

**等待 (Waiting)**：执行 wait () 方法后线程进入等待状态，如果没有显示的 notify () 方法或者 notifyAll () 方法唤醒，该线程会一直处于等待状态。

**超时等待 (Timed_Waiting)**：执行 sleep（Long time）方法后，线程进入超时等待状态，时间一到，自动唤醒线程。

**终止状态 (Terminal)**：当线程的 run () 方法完成时，或者主线程的 main () 方法完成时，我们就认为它终止了。这个线程对象也许是活的，但是，它已经不是一个单独执行的线程。线程一旦终止了，就不能复生。

## 4. 新建（New）状态详解

**实例**：

```java
public class ThreadTest implements Runnable{
    @Override
    public void run() {
        System.out.println("线程："+Thread.currentThread()+" 正在执行...");
    }
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new ThreadTest()); //线程 创建（NEW）状态
    }
}
```

这里仅仅对线程进行了创建，没有执行其他方法。 此时线程的状态就是新建 (New) 状态。

> **Tips**：新建（New）状态的线程，是没有执行 start () 方法的线程。

## 5. 运行（Running）状态详解

**定义**: 线程调度程序从可运行池中选择一个线程作为当前线程时线程所处的状态。这也是线程进入运行状态的唯一方式。

```java
public class ThreadTest implements Runnable{
    .......
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new ThreadTest()); //线程 创建（NEW）状态
        t1. start(); //线程进入 运行（Running）状态
    }
}
```

当线程调用 start () 方法后，线程才进入了运行（Running）状态。

## 6. 阻塞（Blocked）状态详解

**定义**： 阻塞状态是线程阻塞在进入 synchronized 关键字修饰的方法或者代码块时的状态。

我们先来分析如下代码。

**实例**：

```java
public class DemoTest implements Runnable{
    @Override
    public void run() {
        testBolockStatus();
    }
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new DemoTest()); //线程 t1创建（NEW）状态
        t1.setName("T-one");
        Thread t2 = new Thread(new DemoTest()); //线程 t2创建（NEW）状态
        t2.setName("T-two");
        t1. start(); //线程 t1 进入 运行（Running）状态
        t2. start(); //线程 t2 进入 运行（Running）状态
    }

    public static synchronized void testBolockStatus(){ // 该方法被 synchronized修饰
        System.out.println("我是被 synchronized 修饰的同步方法， 正在有线程" +
                Thread.currentThread().getName() +
                "执行我，其他线程进入阻塞状态排队。");
    }
}
```

**代码分析**：

**首先，请看关键代码**：

```java
t1. start(); //线程 t1 进入 运行（Running）状态
t2. start(); //线程 t2 进入 运行（Running）状态
```

我们将线程 t1 和 t2 进行 运行状态的启动，此时 t1 和 t2 就会执行 run () 方法下的 sync testBolockStatus () 方法。

**然后，请看关键代码**：

```java
public static synchronized void testBolockStatus(){ // 该方法被 synchronized修饰

```

testBolockStatus () 方法是被 synchronized 修饰的同步方法。当有 2 条或者 2 条以上的线程执行该方法时， 除了进入方法的一条线程外，其他线程均处于 “阻塞” 状态。

**最后，我们看下执行结果**：

```java
我是被 synchronized 修饰的同步方法， 正在有线程T-one执行我，其他线程进入阻塞状态排队。
我是被 synchronized 修饰的同步方法， 正在有线程T-two执行我，其他线程进入阻塞状态排队。
```

**执行结果解析**：我们有两条线程， 线程名称分别为： T-one 和 T-two。

* **执行结果第一条**： T-one 的状态当时为 运行（Running）状态，T-two 状态为 阻塞（Blocked）状态；
* **执行结果第二条**： T-two 的状态当时为 运行（Running）状态，T-one 状态为 阻塞（Blocked）状态。

## 7. 等待（Waiting）状态详解

**定义**: 执行 wait () 方法后线程进入等待状态，如果没有显示的 notify () 方法或者 notifyAll () 方法唤醒，该线程会一直处于等待状态。

我们通过代码来看下，等待（Waiting）状态。

**实例**：

```java
public class DemoTest implements Runnable{
    @Override
    public void run() {
        try {
            testBolockStatus();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new DemoTest()); //线程 t1创建（NEW）状态
        t1.setName("T-one");
        t1. start(); //线程进入 运行 状态
    }
    public synchronized void testBolockStatus() throws InterruptedException {
        System.out.println("我是线程：" + Thread.currentThread().getName() + ". 我进来了。");
        this.wait(); //线程进入 等待状态 ，没有其他线程 唤醒， 会一直等待下去
        System.out.println("我是被 synchronized 修饰的同步方法， 正在有线程" +
                Thread.currentThread().getName() +
                "执行我，其他线程进入阻塞状态排队。");
    }
}
```

**注意看下关键代码**：

```java
this.wait(); //线程进入 等待状态 ，没有其他线程 唤醒， 会一直等待下去
```

这里调用了 wait () 方法。线程进入 等待（Waiting）状态。如果没有其他线程唤醒，会一直维持等待状态。

**运行结果**：

```java
我是线程：T-one. 我进来了。
```

没有办法打印 wait () 方法后边的执行语句，因为线程已经进入了等待状态。

## 8. 超时等待（Timed-Waiting）状态详解

**定义**: 执行 sleep（Long time）方法后，线程进入超时等待状态，时间一到，自动唤醒线程。

我们通过代码来看下，超时等待（Timed-Waiting）状态。

**实例**：

```java
public class DemoTest implements Runnable{
    @Override
    public void run() {
        .....
    }
    public static void main(String[] args) throws InterruptedException {
       .....
    }
    public synchronized void testBolockStatus() throws InterruptedException {
        System.out.println("我是线程：" + Thread.currentThread().getName() + ". 我进来了。");
        Thread.sleep(5000); //超时等待 状态 5 秒后自动唤醒线程。
        System.out.println("我是被 synchronized 修饰的同步方法， 正在有线程" +
                Thread.currentThread().getName() +
                "执行我，其他线程进入阻塞状态排队。");
    }
}
```

**注意看下关键代码**：

```java
Thread.sleep(5000); //超时等待 状态 5 秒后自动唤醒线程。
```

这里调用了 sleep () 方法。线程进入超时等待（Timed-Waiting）状态。超时等待时间结束，自动唤醒线程继续执行。

**运行结果**：5 秒后，打印第二条语句。

```java
我是线程：T-one. 我进来了。
我睡醒了。我是被 synchronized 修饰的同步方法， 正在有线程T-one执行我，其他线程进入阻塞状态排队。
```

## 9. 终止（Terminal）状态定义

**定义**: 当线程的 run () 方法完成时，或者主线程的 main () 方法完成时，我们就认为它终止了。这个线程对象也许是活的，但是，它已经不是一个单独执行的线程。线程一旦终止了，就不能复生。

## 10. 小结

本节的重中之重在于线程的 6 种不同的状态，本节所有的内容都围绕这 6 种不同的状态进行的讲解，这也是本小节的核心内容，也是必须要掌握的内容。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
