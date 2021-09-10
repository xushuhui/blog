# 并发锁之 Lock 接口

## 1. 前言

本节内容主要是对 Java 并发锁之 Lock 接口进行介绍，Lock 是类似于 synchronized 的另外一种锁的使用，那么本节我们会对 Lock 进行详细的介绍，主要知识点如下：

* Lock 接口的介绍，这是我们开始认识 Lock 的敲门砖，本节课程的基础知识；
* Lock 接口相比于 synchronized 关键字的优点，这也是我们学习 Lock 接口的意义所在；
* Lock 接口的常用方法介绍，了解 Lock 接口中的常用方法，是本节内容的核心知识点。

Lock 是一个接口，并非一个实现类，本节内容主要对 Lock 接口进行一个意义、结构及方法的介绍，为后续讲解 Lock 接口的实现类常用锁奠定一个扎实的基础。

## 2. Lock 接口的介绍

**Lock 接口的诞生**：在 Java 中锁的实现可以由 synchronized 关键字来完成，但在 Java5 之后，出现了一种新的方式来实现，即 Lock 接口。

**诞生的意义**：Lock 接口支持那些语义不同（重入、公平等）的锁规则，可以在非阻塞式结构的上下文（包括 hand-over-hand 和锁重排算法）中使用这些规则。主要的实现是 ReentrantLock。对于 ReentrantLock，后续有专门的小节进行讲解。

**JDK 1.5 前的 synchronized**：在多线程的情况下，当一段代码被 synchronized 修饰之后，同一时刻只能被一个线程访问，其他线程都必须等到该线程释放锁之后才能有机会获取锁访问这段代码。

**Lock 接口**： 实现提供了比使用 synchronized 方法和语句可获得的更广泛的锁定操作。此实现允许更灵活的结构，可以具有差别很大的属性，可以支持多个相关的 Condition 对象。

Lock 相对于 synchronized 关键字而言更加灵活，你可以自由得选择你想要加锁的地方。当然更高的自由度也带来更多的责任。

**使用示例**：我们通常会在 try catch 模块中使用 Lock 关键字，在 finally 模块中释放锁。

```java
	 Lock lock = new ReentrantLock(); //通过子类进行创建，此处以ReentrantLock进行举例
     lock.lock(); //加锁
     try {
         // 对上锁的逻辑进行操作
     } finally {
         lock.unlock(); //释放锁
     }
```

## 3. Lock 接口与 synchronized 关键字的区别

* **实现**：synchronized 关键字基于 JVM 层面实现，JVM 控制锁的获取和释放。Lock 接口基于 JDK 层面，手动进行锁的获取和释放；
* **使用**：synchronized 关键字不用手动释放锁，Lock 接口需要手动释放锁，在 finally 模块中调用 unlock 方法；
* **锁获取超时机制**：synchronized 关键字不支持，Lock 接口支持；
* **获取锁中断机制**：synchronized 关键字不支持，Lock 接口支持；
* **释放锁的条件**：synchronized 关键字在满足占有锁的线程执行完毕，或占有锁的线程异常退出，或占有锁的线程进入 waiting 状态才会释放锁。Lock 接口调用 unlock 方法释放锁；
* **公平性**：synchronized 关键字为非公平锁。Lock 接口可以通过入参自行设置锁的公平性。

## 4. Lock 接口相比 synchronized 关键字的优势

我们通过两个个案例分析来了解 Lock 接口的优势所在。

**案例 1** ：在使用 synchronized 关键字的情形下，假如占有锁的线程由于要等待 IO 或者其他原因（比如调用 sleep 方法）被阻塞了，但是又没有释放锁，那么其他线程就只能一直等待，别无他法。这会极大影响程序执行效率。

**案例 1 分析**：该案例体现了 synchronized 的缺陷，当线程被占有时，其他线程会陷入无条件的长期等待。这是非常可怕的，因为系统资源有限，最终可能导致系统崩溃。

**案例 1 解决**：Lock 接口中的 tryLock (long time, TimeUnit unit) 方法或者响应中断 lockInterruptibly () 方法，能够解决这种长期等待的情况。

**案例 2** ：我们知道，当多个线程读写文件时，读操作和写操作会发生冲突现象，写操作和写操作也会发生冲突现象，但是读操作和读操作不会发生冲突现象。

但是如果采用 synchronized 关键字实现同步的话，就会导致一个问题，即当多个线程都只是进行读操作时，也只有一个线程可以进行读操作，其他线程只能等待锁的释放而无法进行读操作。

**案例 2 分析**：该案例体现了 synchronized 的缺陷，悲观锁的缺陷。我们说过，如果只是读操作，没有增删改操作的话，多线程环境下无需加锁。但是这种情况下，如果在同一时间多个线程进行读操作，synchronized 会 block 其他的读操作，这是不合理的。

**案例 2 解决**：Lock 接口家族也可以解决这种情况，后续我们会对 ReadWriteLock 接口的一个子类 ReentrantReadWriteLock 进行讲解。

**总结**：Lock 接口实现提供了比使用 synchronized 方法和语句可获得的更广泛的锁定操作，能够解决 synchronized 不能够避免的问题。

## 5. Lock 接口的常用方法

我们来简单的看下，JDK 中 Lock 接口的源码中所包含的方法：

```java
public interface Lock {
    void lock();
    void lockInterruptibly() throws InterruptedException;
    boolean tryLock();
    boolean tryLock(long time, TimeUnit unit) throws InterruptedException;
    void unlock();
    Condition newCondition();
}
```

**方法介绍**：

1. **void lock()**：获取锁。如果锁不可用，出于线程调度目的，将禁用当前线程，并且在获得锁之前，该线程将一直处于休眠状态；
2. **void lockInterruptibly()**：如果当前线程未被中断，则获取锁；
3. **boolean tryLock()**：仅在调用时锁为空闲状态才获取该锁。如果锁可用，则获取锁，并立即返回值 true。如果锁不可用，则此方法将立即返回值 false；
4. **boolean tryLock(long time, TimeUnit unit)**：如果锁在给定的等待时间内空闲，并且当前线程未被中断，则获取锁；
5. **void unlock()**：释放锁。在等待条件前，锁必须由当前线程保持。调用 Condition.await () 将在等待前以原子方式释放锁，并在等待返回前重新获取锁；
6. **Condition newCondition()**：返回绑定到此 Lock 实例的新 Condition 实例。

> **Tips**：对 Lock 接口方法的使用，我们必须基于子类进行 Lock 的创建来展示，由于目前我们还未接触 Lock 接口的实现子类，此处只做方法的介绍。后续对 ReentrantLock 进行讲解时，会进行深入讲解。

## 6. 小结

本节主要是对 Lock 接口的常用方法进行了介绍，为本节内容的核心知识。除了方法的介绍外，本节内容不容忽视的一个重点内容是 synchronized 关键字与 Lock 接口的区别，以及 Lock 接口的优势所在。

掌握本节内容，有助于同学对后续实现类锁的学习，为后续的学习奠定了良好的基础。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
