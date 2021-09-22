---
title: Java从零开始（80）可重入读写锁 ReentrantReadWrite
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 可重入读写锁 ReentrantReadWriteLock


## 1. 前言

从本节开始，我们学习新一章内容 —— 并发锁。

在 “[Java 并发原理入门教程](http://www.imooc.com/wiki/concurrencylesson/reentrantlock.html)” 中，介绍了锁相关概念和原理知识，本章各小节内容不再过多解释概念，重点为大家介绍具体锁工具的 API 和使用方法。

本节带领大家认识第一个常用的 Java 并发锁工具之 ReentrantReadWriteLock。

本节先简单介绍 ReentrantReadWriteLock 的基本概念，然后介绍关键的编程方法，最后通过一个编程例子为大家展示 ReentrantReadWriteLock 工具类的用法。

下面我们正式开始介绍吧。

## 2. 概念解释

ReadWriteLock 是一个接口类型，翻译为 “读写锁”。多个线程同时读某个资源时，为了满足并发量，不应该加锁处理，但是如果有一个线程写这个资源，就不应该再有其它线程对该资源进行读操作或者写操作。

实现了此接口的类提供了 “读读能共存、读写不能共存、写写不能共存” 的控制逻辑。当同一个资源被大量线程读取，但仅有少数线程修改时，使用 ReadWriteLock 可大大提高并发效率。比如对一个电商网站的商品，回复评论（写）是不频繁的，但是浏览（读）是非常频繁的，这种情况使用 ReadWriteLock 工具类做并发控制非常适合。总之适合读多写少的场景。

ReentrantLock 翻译为 “可重入锁”。“可重入” 是什么意思呢？就是指一个线程可以多次获取该锁，Java 中 在语言语法层提供的 syncrinized 是最常见的用于并发控制的关键字，其构成的锁也是可重入的。可重入锁在一定程度可以避免死锁的发生。

更多关于锁的原理，可阅读 “[Java 并发原理入门教程](http://www.imooc.com/wiki/concurrencylesson/reentrantlock.html)”。

ReentrantReadWriteLock 类实现了 “读写锁” 和 “可重入锁” 的双重功能。其本身不提供加锁服务，只负责提供读锁和写锁。在介绍关键编程方法之前，我们先看一张图整体了解关键方法的使用方式。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz2p7bij60o40i10z702)

下面我们学习其关键的编程方法。

## 3.ReentrantReadWriteLock 的编程方法

1. 构造方法 ReentrantReadWriteLock () 和 ReentrantReadWriteLock (boolean fair)

有两个构造方法，在构造对象时，可选择是否构造为公平锁模式。什么是公平锁呢？公平锁指的是获取所的线程按照申请锁的线程获取，而非乱序随机获取到锁权限。

2. writeLock () 和 readLock ()

这两个方法分别用于获取写锁和获取读锁，对读操作的逻辑使用读锁对象加锁，对写操作的逻辑使用写锁对象加锁，如此可以做到 “读读能共存、读写不能共存、写写不能共存”，在实现线程安全的情况下，提高并发效率。

3. isFair()

获取锁是否具有公平性，此方法返回 boolean 值。

4. getWriteHoldCount () 、 getReadHoldCount () 和 getReadLockCount ()

getWriteHoldCount () 方法用于获取当前线程的写锁计数，getReadHoldCount () 方法用于获取当前线程的读锁计数，getReadLockCount () 方法用于获取总的读锁计数。所有方法均返回 int 值。

## 4.ReentrantReadWriteLock.ReadLock 的编程方法

此类是一个静态内部类，封装在 ReentrantReadWriteLock 中，此类不对外提供构造方法，由 ReentrantReadWriteLock.readLock () 方法获取此类对象。

1. lock () 和 lockInterruptibly ()

两个方法都用于加读锁，第二个方法可被中断。

2. tryLock () 和 tryLock (long timeout, TimeUnit unit)

两个方法都返回 boolean 值，用于尝试加读锁，第一个方法在尝试不成功时立刻返回，第二个方法可在指定时间内尝试加读锁。这两个方法提供了加锁的柔性，提供了更多操作空间。

3. unlock()

释放读锁。

## 5.ReentrantReadWriteLock.WriteLock 的编程方法

此类是一个静态内部类，封装在 ReentrantReadWriteLock 中，此类不对外提供构造方法，由 ReentrantReadWriteLock.writeLock () 方法获取此类对象。

1. lock () 和 lockInterruptibly ()

两个方法都用于加写锁，第二个方法可被中断。

2. tryLock () 和 tryLock (long timeout, TimeUnit unit)

两个方法都返回 boolean 值，用于尝试加写锁，第一个方法在尝试不成功时立刻返回，第二个方法可在指定时间内尝试加写锁。这两个方法提供了加锁的柔性，提供了更多操作空间。

3. unlock()

释放写锁。

## 6. 编程案例

上面介绍了核心编程方法，我们举一个编程案例，实际体会一下 ReentrantReadWriteLock 的用法。

```java
import lombok.SneakyThrows;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class ReadWriteLockTest {

    // 创建读写锁对象
    private final ReadWriteLock readWriteLock = new ReentrantReadWriteLock();
    // 获取读锁对象
    private final Lock readlock = readWriteLock.readLock();
    // 获取写锁对象
    private final Lock writelock = readWriteLock.writeLock();
    // 待控制的资源
    private int account = 0;

    private static ReadWriteLockTest readWriteLockTest = new ReadWriteLockTest();

    public static void main(String[] args) {
        new Thread(new Runnable() {
            @SneakyThrows
            public void run() {
                while(true) {
                    Thread.sleep(1000);
                    int tmp = readWriteLockTest.get();
                    System.out.println("读操作:" + tmp);
                }
            }
        }).start();

        new Thread(new Runnable() {
            @SneakyThrows
            public void run() {
                while(true) {
                    Thread.sleep(2000);
                    readWriteLockTest.add(10);
                }
            }
        }).start();
    }

    public void add(int value) {
        // 加写锁
        writelock.lock();
        try {
            account += 1;
        } finally {
            // 释放写锁
            writelock.unlock();
        }
    }

    public int get() {
        // 加读锁
        readlock.lock();
        try {
            return account;
        } finally {
            // 释放读锁
            readlock.unlock();
        }
    }
}
```

运行上面代码一段时间后结果如下：

```java
读操作:0
读操作:1
读操作:1
读操作:2
读操作:2
```

注意在使用时，获取锁的操作 lock () 应该放在 try 之前，而释放锁的操作 unlock () 需要放在 finally 中，可确保锁释放。

## 7. 小结

本节解释了 ReentrantReadWriteLock 的基本概念和应用场合，且通过一个简单的例子展示了其用法，更多关于此工具类的概念和原理介绍，可阅读 “[Java 并发原理入门教程](http://www.imooc.com/wiki/concurrencylesson/reentrantlock.html)” 。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
