# 线程上下文切换与死锁

## 1. 前言

本节内容主要是对死锁进行深入的讲解，具体内容点如下：

* 理解线程的上下文切换，这是本节的辅助基础内容，从概念层面进行理解即可；
* 了解什么是线程死锁，在并发编程中，线程死锁是一个致命的错误，死锁的概念是本节的重点之一；
* 了解线程死锁的必备 4 要素，这是避免死锁的前提，了解死锁的必备要素，才能找到避免死锁的方式；
* 掌握死锁的实现，通过代码实例，进行死锁的实现，深入体会什么是死锁，这是本节的重难点之一；
* 掌握如何避免线程死锁，我们能够实现死锁，也可以避免死锁，这是本节内容的核心。

## 2. 理解线程的上下文切换

**概述**：在多线程编程中，线程个数一般都大于 CPU 个数，而每个 CPU 同一时－刻只能被一个线程使用，为了让用户感觉多个线程是在同时执行的， CPU 资源的分配采用了时间片轮转的策略，也就是给每个线程分配一个时间片，线程在时间片内占用 CPU 执行任务。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz5gutjj60jg05adh702)

**定义**：当前线程使用完时间片后，就会处于就绪状态并让出 CPU，让其他线程占用，这就是上下文切换，从当前线程的上下文切换到了其他线程。

**问题点解析**：那么就有一个问题，让出 CPU 的线程等下次轮到自己占有 CPU 时如何知道自己之前运行到哪里了？所以在切换线程上下文时需要保存当前线程的执行现场， 当再次执行时根据保存的执行现场信息恢复执行现场。

**线程上下文切换时机**： 当前线程的 CPU 时间片使用完或者是当前线程被其他线程中断时，当前线程就会释放执行权。那么此时执行权就会被切换给其他的线程进行任务的执行，一个线程释放，另外一个线程获取，就是我们所说的上下文切换时机。

## 3. 什么是线程死锁

**定义**：死锁是指两个或两个以上的线程在执行过程中，因争夺资源而造成的互相等待的现象，在无外力作用的情况下，这些线程会一直相互等待而无法继续运行下去。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz5s2czj60jg062wfq02)

如上图所示死锁状态，线程 A 己经持有了资源 2，它同时还想申请资源 1，可是此时线程 B 已经持有了资源 1 ，线程 A 只能等待。

反观线程 B 持有了资源 1 ，它同时还想申请资源 2，但是资源 2 已经被线程 A 持有，线程 B 只能等待。所以线程 A 和线程 B 就因为相互等待对方已经持有的资源，而进入了死锁状态。

## 4. 线程死锁的必备要素

* **互斥条件**：进程要求对所分配的资源进行排他性控制，即在一段时间内某资源仅为一个进程所占有。此时若有其他进程请求该资源，则请求进程只能等待；
* **不可剥夺条件**：进程所获得的资源在未使用完毕之前，不能被其他进程强行夺走，即只能由获得该资源的进程自己来释放（只能是主动释放，如 yield 释放 CPU 执行权）；
* **请求与保持条件**：进程已经保持了至少一个资源，但又提出了新的资源请求，而该资源已被其他进程占有，此时请求进程被阻塞，但对自己已获得的资源保持不放；
* **循环等待条件**：指在发生死锁时，必然存在一个线程请求资源的环形链，即线程集合 {T0,T1,T2,…Tn｝中的 T0 正在等待一个 T1 占用的资源，T1 正在等待 T2 占用的资源，以此类推，Tn 正在等待己被 T0 占用的资源。

**如下图所示**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz64d99j60jg05jabr02)

## 5. 死锁的实现

为了更好的了解死锁是如何产生的，我们首先来设计一个死锁争夺资源的场景。

**场景设计**：

* 创建 2 个线程，线程名分别为 threadA 和 threadB；
* 创建两个资源， 使用 new Object () 创建即可，分别命名为 resourceA 和 resourceB；
* threadA 持有 resourceA 并申请资源 resourceB；
* threadB 持有 resourceB 并申请资源 resourceA ；
* 为了确保发生死锁现象，请使用 sleep 方法创造该场景；
* 执行代码，看是否会发生死锁。

**期望结果**：发生死锁，线程 threadA 和 threadB 互相等待。

> **Tips**：此处的实验会使用到关键字 synchronized，后续小节还会对关键字 synchronized 单独进行深入讲解，此处对 synchronized 的使用仅仅为初级使用，有 JavaSE 基础即可。

**实例**：

```java
public class DemoTest{
    private static  Object resourceA = new Object();//创建资源 resourceA
    private static  Object resourceB = new Object();//创建资源 resourceB

    public static void main(String[] args) throws InterruptedException {
        //创建线程 threadA
        Thread threadA = new Thread(new Runnable() {
            @Override
            public void run() {
                synchronized (resourceA) {
                    System.out.println(Thread.currentThread().getName() + "获取 resourceA。");
                    try {
                        Thread.sleep(1000); // sleep 1000 毫秒，确保此时 resourceB 已经进入run 方法的同步模块
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                    System.out.println(Thread.currentThread().getName() + "开始申请 resourceB。");
                    synchronized (resourceB) {
                        System.out.println (Thread.currentThread().getName() + "获取 resourceB。");
                    }
                }
            }
        });
        threadA.setName("threadA");
        //创建线程 threadB
        Thread threadB = new Thread(new Runnable() { //创建线程 1
            @Override
            public void run() {
                synchronized (resourceB) {
                    System.out.println(Thread.currentThread().getName() + "获取 resourceB。");
                    try {
                        Thread.sleep(1000); // sleep 1000 毫秒，确保此时 resourceA 已经进入run 方法的同步模块
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                    System.out.println(Thread.currentThread().getName() + "开始申请 resourceA。");
                    synchronized (resourceA) {
                        System.out.println (Thread.currentThread().getName() + "获取 resourceA。");
                    }
                }
            }
        });
        threadB.setName("threadB");

        threadA. start();
        threadB. start();
    }
}
```

**代码讲解**：

* 从代码中来看，我们首先创建了两个资源 resourceA 和 resourceB；
* 然后创建了两条线程 threadA 和 threadB。threadA 首先获取了 resourceA ，获取的方式是代码 synchronized (resourceA) ，然后沉睡 1000 毫秒；
* 在 threadA 沉睡过程中， threadB 获取了 resourceB，然后使自己沉睡 1000 毫秒；
* 当两个线程都苏醒时，此时可以确定 threadA 获取了 resourceA，threadB 获取了 resourceB，这就达到了我们做的第一步，线程分别持有自己的资源；
* 那么第二步就是开始申请资源，threadA 申请资源 resourceB，threadB 申请资源 resourceA 无奈 resourceA 和 resourceB 都被各自线程持有，两个线程均无法申请成功，最终达成死锁状态。

**执行结果验证**：

```java
threadA 获取 resourceA。
threadB 获取 resourceB。
threadA 开始申请 resourceB。
threadB 开始申请 resourceA。
```

看下验证结果，发现已经出现死锁，threadA 申请 resourceB，threadB 申请 resourceA，但均无法申请成功，死锁得以实验成功。

## 6. 如何避免线程死锁

要想避免死锁，只需要破坏掉至少一个构造死锁的必要条件即可，学过操作系统的读者应该都知道，目前只有请求并持有和环路等待条件是可以被破坏的。

造成死锁的原因其实和申请资源的顺序有很大关系，使用资源申请的有序性原则就可避免死锁。

我们依然以第 5 个知识点进行讲解，那么实验的需求和场景不变，我们仅仅对之前的 threadB 的代码做如下修改，以避免死锁。

**代码修改**：

```java
Thread threadB = new Thread(new Runnable() { //创建线程 1
            @Override
            public void run() {
                synchronized (resourceA) { //修改点 1
                    System.out.println(Thread.currentThread().getName() + "获取 resourceB。");//修改点 3
                    try {
                        Thread.sleep(1000); // sleep 1000 毫秒，确保此时 resourceA 已经进入run 方法的同步模块
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                    System.out.println(Thread.currentThread().getName() + "开始申请 resourceA。");//修改点 4
                    synchronized (resourceB) { //修改点 2
                        System.out.println (Thread.currentThread().getName() + "获取 resourceA。"); //修改点 5
                    }
                }
            }
        });
```

请看如上代码示例，有 5 个修改点：

* **修改点 1** ：将 resourceB 修改成 resourceA；
* **修改点 2** ：将 resourceA 修改成 resourceB；
* **修改点 3** ：将 resourceB 修改成 resourceA；
* **修改点 4** ：将 resourceA 修改成 resourceB；
* **修改点 5** ：将 resourceA 修改成 resourceB。

请读者按指示修改代码，并从新运行验证。

**修改后代码讲解**：

* 从代码中来看，我们首先创建了两个资源 resourceA 和 resourceB；
* 然后创建了两条线程 threadA 和 threadB。threadA 首先获取了 resourceA ，获取的方式是代码 synchronized (resourceA) ，然后沉睡 1000 毫秒；
* 在 threadA 沉睡过程中， threadB 想要获取 resourceA ，但是 resourceA 目前正被沉睡的 threadA 持有，所以 threadB 等待 threadA 释放 resourceA；
* 1000 毫秒后，threadA 苏醒了，释放了 resourceA ，此时等待的 threadB 获取到了 resourceA，然后 threadB 使自己沉睡 1000 毫秒；
* threadB 沉睡过程中，threadA 申请 resourceB 成功，继续执行成功后，释放 resourceB；
* 1000 毫秒后，threadB 苏醒了，继续执行获取 resourceB ，执行成功。

**执行结果验证**：

```java
threadA 获取 resourceA。
threadA 开始申请 resourceB。
threadA 获取 resourceB。
threadB 获取 resourceA。
threadB 开始申请 resourceB。
threadB 获取 resourceB。
```

我们发现 threadA 和 threadB 按照相同的顺序对 resourceA 和 resourceB 依次进行访问，避免了互相交叉持有等待的状态，避免了死锁的发生。

## 7. 小结

死锁是并发编程中最致命的问题，如何避免死锁，是并发编程中恒久不变的问题。

掌握死锁的实现以及如果避免死锁的发生，是本节内容的重中之重。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
