---
title: Java从零开始（106）读写锁ReentrantReadWriteLock
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 读写锁 ReentrantReadWriteLock


## 1. 前言

本节内容主要是对 Java 读写锁 ReentrantReadWriteLock 进行讲解，本节内容几乎全部为重点知识，需要学习者对 ReentrantReadWriteLock 进行理解和掌握。本节内容的知识点如下：

* ReentrantReadWriteLock 简单介绍，对 ReentrantReadWriteLock 进行一个总体的概括；
* ReentrantReadWriteLock 的类结构，从 Java 层面了解 ReentrantReadWriteLock；
* ReentrantReadWriteLock 的特点，相比于上两点知识，该知识点可视为重点；
* ReentrantReadWriteLock 读锁共享的性质验证，为本节核心内容之一；
* ReentrantReadWriteLock 读写互斥的性质验证，为本节核心内容之一。

ReentrantReadWriteLock 在 Java 的锁当中也占据着十分重要的地位，在并发编程中使用频率也是非常的高，一定要对本节内容进行细致的学习和掌握。

## 2. ReentrantReadWriteLock 介绍

JDK 提供了 ReentrantReadWriteLock 读写锁，使用它可以加快效率，在某些不需要操作实例变量的方法中，完全可以使用读写锁 ReemtrantReadWriteLock 来提升该方法的运行速度。

**定义**：读写锁表示有两个锁，一个是读操作相关的锁，也称为共享锁；另一个是写操作相关的锁，也叫排他锁。

**定义解读**：也就是多个读锁之间不互斥，读锁与写锁互斥、写锁与写锁互斥。在没有线程 Thread 进行写入操作时，进行读取操作的多个 Thread 都可以获取读锁，而进行写入操作的 Thread 只有在获取写锁后才能进行写入操作。即多个 Thread 可以同时进行读取操作，但是同一时刻只允许一个 Thread 进行写入操作。

## 3. ReentrantReadWriteLock 的类结构

ReentrantReadWriteLock 是接口 ReadWriteLock 的子类实现，通过 JDK 的代码可以看出这一实现关系。

```java
public class ReentrantReadWriteLock
        implements ReadWriteLock, java.io.Serializable{}
```

我们再来看下接口 ReadWriteLock，该接口只定义了两个方法：

```java
public interface ReadWriteLock {
    Lock readLock();
    Lock writeLock();
}
```

通过调用相应方法获取读锁或写锁，可以如同使用 Lock 接口一样使用。

## 4. ReentrantReadWriteLock 的特点

**性质 1** ：可重入性。

ReentrantReadWriteLock 与 ReentrantLock 以及 synchronized 一样，都是可重入性锁，这里不会再多加赘述所得可重入性质，之前已经做过详细的讲解。

**性质 2** ：读写分离。

我们知道，对于一个数据，不管是几个线程同时读都不会出现任何问题，但是写就不一样了，几个线程对同一个数据进行更改就可能会出现数据不一致的问题，因此想出了一个方法就是对数据加锁，这时候出现了一个问题：

线程写数据的时候加锁是为了确保数据的准确性，但是线程读数据的时候再加锁就会大大降低效率，这时候怎么办呢？那就对写数据和读数据分开，加上两把不同的锁，不仅保证了正确性，还能提高效率。

**性质 3** ：可以锁降级，写锁降级为读锁。

线程获取写入锁后可以获取读取锁，然后释放写入锁，这样就从写入锁变成了读取锁，从而实现锁降级的特性。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvvu7n2j60jg02n75502)

**性质 4** ：不可锁升级。

线程获取读锁是不能直接升级为写入锁的。需要释放所有读取锁，才可获取写锁。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvw5hs2j60jg02c3zb02)

## 5. ReentrantReadWriteLock 读锁共享

我们之前说过，ReentrantReadWriteLock 之所以优秀，是因为读锁与写锁是分离的，当所有的线程都为读操作时，不会造成线程之间的互相阻塞，提升了效率，那么接下来，我们通过代码实例进行学习。

**场景设计**：

* 创建三个线程，线程名称分别为 t1，t2，t3，线程实现方式自行选择；
* 三个线程同时运行获取读锁，读锁成功后打印线程名和获取结果，并沉睡 2000 毫秒，便于观察其他线程是否可共享读锁；
* finally 模块中释放锁并打印线程名和释放结果；
* 运行程序，观察结果。

**结果预期**：三条线程能同时获取锁，因为读锁共享。

**实例**：

```java
public class DemoTest {
    private ReentrantReadWriteLock lock = new ReentrantReadWriteLock();// 读写锁
    private int i;
    public String readI() {
        try {
            lock.readLock().lock();// 占用读锁
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 占用读锁,i->" + i);
            Thread.sleep(2000);
        } catch (InterruptedException e) {

        } finally {
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 释放读锁,i->" + i);
            lock.readLock().unlock();// 释放读锁
        }
        return i + "";
    }

    public static void main(String[] args) {
        final DemoTest demo1 = new DemoTest();
        Runnable runnable = new Runnable() {
            @Override
            public void run() {
                demo1.readI();
            }
        };
        new Thread(runnable, "t1"). start();
        new Thread(runnable, "t2"). start();
        new Thread(runnable, "t3"). start();
    }
}
```

**结果验证**：

```java
threadName -> t1 占用读锁,i->0
threadName -> t2 占用读锁,i->0
threadName -> t3 占用读锁,i->0
threadName -> t1 释放读锁,i->0
threadName -> t3 释放读锁,i->0
threadName -> t2 释放读锁,i->0
```

**结果分析**：从结果来看，t1，t2，t3 均在同一时间获取了锁，证明了读锁共享的性质。

## 6. ReentrantReadWriteLock 读写互斥

当共享变量有写操作时，必须要对资源进行加锁，此时如果一个线程正在进行读操作，那么写操作的线程需要等待。同理，如果一个线程正在写操作，读操作的线程需要等待。

**场景设计**：细节操作不详细阐述，看示例代码即可。

* 创建两个线程，线程名称分别为 t1，t2；
* 线程 t1 进行读操作，获取到读锁之后，沉睡 5000 毫秒；
* 线程 t2 进行写操作；
* 开启 t1，1000 毫秒后开启 t2 线程；
* 运行程序，观察结果。

**结果预期**：线程 t1 获取了读锁，在沉睡的 5000 毫秒中，线程 t2 只能等待，不能获取到锁，因为读写互斥。

**实例**：

```java
public class DemoTest {
    private ReentrantReadWriteLock lock = new ReentrantReadWriteLock();// 读写锁
    private int i;
    public String readI() {
        try {
            lock.readLock().lock();// 占用读锁
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 占用读锁,i->" + i);
            Thread.sleep(5000);
        } catch (InterruptedException e) {
        } finally {
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 释放读锁,i->" + i);
            lock.readLock().unlock();// 释放读锁
        }
        return i + "";
    }

    public void addI() {
        try {
            lock.writeLock().lock();// 占用写锁
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 占用写锁,i->" + i);
            i++;
        } finally {
            System.out.println("threadName -> " + Thread.currentThread().getName() + " 释放写锁,i->" + i);
            lock.writeLock().unlock();// 释放写锁
        }
    }

    public static void main(String[] args) throws InterruptedException {
        final DemoTest demo1 = new DemoTest();
        new Thread(new Runnable() {
            @Override
            public void run() {
                demo1.readI();
            }
        }, "t1"). start();
        Thread.sleep(1000);
        new Thread(new Runnable() {
            @Override
            public void run() {
                demo1.addI();
            }
        }, "t2"). start();
    }
}
```

**结果验证**：

```java
threadName -> t1 占用读锁,i->0
threadName -> t1 释放读锁,i->0
threadName -> t2 占用写锁,i->0
threadName -> t2 释放写锁,i->1
```

**结果解析**：验证成功，在线程 t1 沉睡的过程中，写锁 t2 线程无法获取锁，因为锁已经被读操作 t1 线程占据了。

## 7. 小结

本节内容只要是对读写锁 ReentrantReadWriteLock 进行的比较细致的讲解，对于本节的内容几乎通篇为重点内容。

其中核心知识点为读锁共享和读写互斥的验证，所有的知识点都是围绕这两个话题进行讲解的，有兴趣的同学可以根据实例代码进行写锁互斥的验证。唯一不同的地方就是创建两个写线程进行写锁的获取。

掌握本节知识点，有助于我们在特定的场景下对读写锁进行应用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
