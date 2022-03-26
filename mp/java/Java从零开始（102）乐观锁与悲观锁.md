---
title: Java从零开始（102）乐观锁与悲观锁
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 乐观锁与悲观锁


## 1. 前言

本节内容主要是对 Java 乐观锁与悲观锁进行更加深入的讲解，本节内容更加偏重于对乐观锁的讲解，因为 synchronized 悲观锁对于大部分学习者并不陌生，本节主要内容如下：

* 乐观锁与悲观锁的概念，之前有所讲解，这里用很小的篇幅进行知识的回顾，巩固；
* 乐观锁与悲观锁的使用场景介绍，通过理解悲观锁与乐观锁不同的风格，理解什么场景下需要选择合适的锁，为本节的重点内容之一；
* 了解乐观锁的缺点，乐观锁有自己的特定的缺陷，不同的锁都有自己的优点与缺点；
* 了解乐观锁缺陷的解决方式，作为本节内容的重点之一；
* 通过引入 Atomic 操作，实现乐观锁，为本节内容的核心，通过对比 synchronized 的实现，用两种锁机制实现同一个需求。

本节内容为 CAS 原理的进阶讲解，也是乐观锁与悲观锁的深入讲解。因为对于并发编程，悲观锁与乐观锁的涉及频率非常高，所以对其进行更加深入的讲解。

## 2. 乐观锁与悲观锁的概念

**悲观锁**：总是假设最坏的情况，每次去拿数据的时候都认为别人会修改，所以每次在拿数据的时候都会上锁，这样其他线程想拿这个数据就会阻塞直到它拿到锁（共享资源每次只给一个线程使用，其它线程阻塞，用完后再把资源转让给其它线程）。

**乐观锁**：总是假设最好的情况，每次去拿数据的时候都认为别人不会修改，所以不会上锁，但是在更新的时候会判断一下在此期间别人有没有去更新这个数据，可以使用版本号机制和 CAS 算法实现。

乐观锁适用于多读的应用类型，这样可以提高吞吐量，像数据库提供的类似于 write_condition 机制，其实都是提供的乐观锁。

## 3. 乐观锁与悲观锁的使用场景

简单的来说 CAS 适用于写比较少的情况下（多读场景，冲突一般较少），synchronized 适用于写比较多的情况下（多写场景，冲突一般较多）。

* 对于资源竞争较少（线程冲突较轻）的情况，使用 synchronized 同步锁进行线程阻塞和唤醒切换以及用户态内核态间的切换操作额外浪费消耗 CPU 资源；而 CAS 基于硬件实现，不需要进入内核，不需要切换线程，操作自旋几率较少，因此可以获得更高的性能；
* 对于资源竞争严重（线程冲突严重）的情况，CAS 自旋的概率会比较大，从而浪费更多的 CPU 资源，效率低于 synchronized。

**总结**：乐观锁适用于写比较少的情况下（多读场景），即冲突真的很少发生的时候，这样可以省去了锁的开销，加大了系统的整个吞吐量。

但如果是多写的情况，一般会经常产生冲突，这就会导致上层应用会不断地进行 retry，这样反倒是降低了性能，所以一般多写的场景下用悲观锁就比较合适。

## 4. 乐观锁的缺点

**ABA 问题**：我们之前也对此进行过介绍。

如果一个变量 V 初次读取的时候是 A 值，并且在准备赋值的时候检查到它仍然是 A 值，那我们就能说明它的值没有被其他线程修改过了吗？

很明显是不能的，因为在这段时间它的值可能被改为其他值，然后又改回 A，那 CAS 操作就会误认为它从来没有被修改过。这个问题被称为 CAS 操作的 “ABA” 问题。

**循环时间长开销大**：在特定场景下会有效率问题。

自旋 CAS（也就是不成功就一直循环执行直到成功）如果长时间不成功，会给 CPU 带来非常大的执行开销。

**总结**：我们这里主要关注 ABA 问题。循环时间长开销大的问题，在特定场景下很难避免的，因为所有的操作都需要在合适自己的场景下才能发挥出自己特有的优势。

## 5. ABA 问题解决之版本号机制

讲解 CAS 原理时，对于解决办法进行了简要的介绍，仅仅是一笔带过。这里进行较详细的阐释。其实 ABA 问题的解决，我们通常通过如下方式进行解决：版本号机制。我们一起来看下版本号机制：

**版本号机制**：一般是在数据中加上一个数据版本号 version 字段，表示数据被修改的次数，当数据被修改时，version 值会加 1。当线程 A 要更新数据值时，在读取数据的同时也会读取 version 值，在提交更新时，若刚才读取到的 version 值为当前数据中的 version 值相等时才更新，否则重试更新操作，直到更新成功。

**场景示例**：假设商店类 Shop 中有一个 version 字段，当前值为 1 ；而当前商品数量为 50。

* 店员 A 此时将其读出（ version=1 ），并将商品数量扣除 10，更新为 50 - 10 = 40；
* 在店员 A 操作的过程中，店员 B 也读入此信息（ version=1 ），并将商品数量扣除 20，更新为 50 - 20 = 30；
* 店员 A 完成了修改工作，将数据版本号加 1（ version=2 ），商品数量为 40，提交更新，此时由于提交数据版本大于记录当前版本，数据被更新，数据记录 version 更新为 2 ；
* 店员 B 完成了操作，也将版本号加 1（ version=2 ），试图更新商品数量为 30。但此时比对数据记录版本时发现，店员 B 提交的数据版本号为 2 ，数据记录当前版本也为 2 ，不满足 “ 提交版本必须大于记录当前版本才能执行更新 “ 的乐观锁策略，因此，店员 B 的提交被驳回；
* 店员 B 再次重新获取数据，version = 2，商品数量 40。在这个基础上继续执行自己扣除 20 的操作，商品数量更新为 40 - 20 = 20；
* 店员 B 将版本号加 1 ，version = 3，将之前的记录 version 2 更新为 3 ，将之前的数量 40 更新 为 20。

从如上描述来看，所有的操作都不会出现脏数据，关键在于版本号的控制。

> **Tips**：Java 对于乐观锁的使用进行了良好的封装，我们可以直接使用并发编程包来进行乐观锁的使用。本节接下来所使用的 Atomic 操作即为封装好的操作。
>
> 之所以还要对 CAS 原理以及 ABA 问题进行深入的分析，主要是为了让学习者了解底层的原理，以便更好地在不同的场景下选择使用锁的类型。

## 6. Atomic 操作实现乐观锁

为了更好地理解悲观锁与乐观锁，我们通过设置一个简单的示例场景来进行分析。并且我们采用悲观锁 synchronized 和乐观锁 Atomic 操作进行分别实现。

Atomic 操作类，指的是 java.util.concurrent.atomic 包下，一系列以 Atomic 开头的包装类。例如 AtomicBoolean，AtomicInteger，AtomicLong。它们分别用于 Boolean，Integer，Long 类型的原子性操作。

Atomic 操作的底层实现正是利用的 CAS 机制，而 CAS 机制即乐观锁。

**场景设计**：

* 创建两个线程，创建方式可自选；
* 定义一个全局共享的 static int 变量 count，初始值为 0；
* 两个线程同时操作 count，每次操作 count 加 1；
* 每个线程做 100 次 count 的增加操作。

**结果预期**：最终 count 的值应该为 200。

**悲观锁 synchronized 实现**：

```java
public class DemoTest extends Thread{
    private static int count = 0; //定义count = 0
    public static void main(String[] args) {
        for (int i = 0; i < 2; i++) { //通过for循环创建两个线程
            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        Thread.sleep(10);
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                    //每个线程让count自增100次
                    for (int i = 0; i < 100; i++) {
                        synchronized (DemoTest.class){
                            count++;
                        }
                    }
                }
            }). start();
        }
        try{
            Thread.sleep(2000);
        }catch (Exception e){
            e.printStackTrace();
        }
        System.out.println(count);
    }
}
```

**结果验证**：

```java
200
```

**乐观锁 Atomic 操作实现**：

```java
public class DemoTest extends Thread{
    //Atomic 操作，引入AtomicInteger。这是实现乐观锁的关键所在。
    private static AtomicInteger count = new AtomicInteger(0);
    public static void main(String[] args) {
        for (int i = 0; i < 2; i++) {
            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        Thread.sleep(10);
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                    //每个线程让count自增100次
                    for (int i = 0; i < 100; i++) {
                        count.incrementAndGet();
                    }
                }
            }). start();
        }
        try{
            Thread.sleep(2000);
        }catch (Exception e){
            e.printStackTrace();
        }
        System.out.println(count);
    }
}
```

**结果验证**：

```java
200
```

**代码解读**：

此处主要关注两个点，第一个是 count 的创建，是通过 AtomicInteger 进行的实例化，这是使用 Atomic 的操作的入口，也是使用 CAS 乐观锁的一个标志。

第二个是需要关注 count 的增加 1 调用是 AtomicInteger 中 的 incrementAndGet 方法，该方法是原子性操作，遵循 CAS 原理。

## 7. 小结

本节内容所有的知识点讲解都可以作为重点内容进行学习。悲观锁与乐观锁是并发编程中所涉及的非常重要的内容，一定要深入的理解和掌握。

对于课程中 CAS 原理的进阶讲解，也是非常重要的知识点，对于 ABA 问题，是并发编程中所涉及的高频话题、考题，也要对此加以理解和掌握。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
