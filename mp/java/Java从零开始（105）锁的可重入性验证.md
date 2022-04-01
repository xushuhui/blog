---
title: Java从零开始（105）锁的可重入性验证
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 锁的可重入性验证


## 1. 前言

本节内容主要是对 Java 锁的可重入性进行验证，锁的可重入性的设计是避免死锁非常好的设计思想。本节内容的知识点如下：

* 什么是锁的可重入性，这是本节课程的基础内容；
* 了解可重入锁与非可重入性锁的不同之处，以凸显可重入性锁的优势所在，为本节基础内容；
* 了解什么情况下使用可重入锁，是本节的重点内容之一；
* synchronized 关键字验证锁的可重入性试验，为本节核心内容之一；
* ReentrantLock 验证锁的可重入性试验，为本节核心内容之一；

其实 synchronized 关键字与 ReentrantLock 都是 Java 常见的可重入锁，本节内容使用 ReentrantLock 和 synchronized 来讲解锁的可重入性。

## 2. 什么是锁的可重入性

**定义**：可重入锁又名递归锁，是指在同一个线程在外层方法获取锁的时候，再进入该线程的内层方法会自动获取锁（前提锁对象得是同一个对象或者 class），不会因为之前已经获取过还没释放而阻塞。

Java 中 ReentrantLock 和 synchronized 都是可重入锁，可重入锁的一个优点是可一定程度避免死锁。

**可重入锁原理**：可重入锁的原理是在锁内部维护一个线程标示，用来标示该锁目前被哪个线程占用，然后关联一个计数器。一开始计数器值为 0，说明该锁没有被任何线程占用。当一个线程获取了该锁时，计数器的值会变成 1，这时其他线程再来获取该锁时会发现锁的所有者不是自己而被阻塞挂起。

但是当获取了该锁的线程再次获取锁时发现锁拥有者是自己，就会把计数器值加＋1, 当释放锁后计数器值－1。当计数器值为 0 时，锁里面的线程标示被重置为 null，这时候被阻塞的线程会被唤醒来竞争获取该锁。

## 3. 可重入锁与非可重入性锁

Java 中 ReentrantLock 和 synchronized 都是可重入锁，可重入锁的一个优点是可一定程度避免死锁。

为了解释可重入锁与非可重入性锁的区别与联系，我们拿可重入锁 ReentrantLock 和 非重入锁 NonReentrantLock 进行简单的分析对比。

**相同点**： ReentrantLock 和 NonReentrantLock 都继承父类 AQS，其父类 AQS 中维护了一个同步状态 status 来计数重入次数，status 初始值为 0。

**不同点**：当线程尝试获取锁时，可重入锁先尝试获取并更新 status 值，如果 status == 0 表示没有其他线程在执行同步代码，则把 status 置为 1，当前线程开始执行。

如果 status != 0，则判断当前线程是否是获取到这个锁的线程，如果是的话执行 status+1，且当前线程可以再次获取锁。

而非可重入锁是直接去获取并尝试更新当前 status 的值，如果 status != 0 的话会导致其获取锁失败，当前线程阻塞，导致死锁发生。

## 4. 什么情况下使用可重入锁

**我们先来看看如下代码**：同步方法 helloB 方法调用了同步方法 helloA。

```java
public class DemoTest{
    public synchronized void helloA(){
        System.out.println("helloA");
    }
    public synchronized void helloB(){
        System.out.println("helloB");
        helloA();
    }
}
```

在如上代码中，调用 helloB 方法前会先获取内置锁，然后打印输出。之后调用 helloA 方法，在调用前会先去获取内置锁，如果内置锁不是可重入的，那么调用线程将会一直被阻塞。

因此，对于同步方法内部调用另外一个同步方法的情况下，一定要使用可重入锁，不然会导致死锁的发生。

## 5. synchronized 验证锁的可重入性

为了更好的理解 synchronized 验证锁的可重入性，我们来设计一个简单的场景。

**场景设计**：

* 创建一个类，该类中有两个方法，helloA 方法和 helloB 方法；
* 将两个方法内部的逻辑进行 synchronized 同步；
* helloA 方法内部调用 helloB 方法，营造可重入锁的场景；
* main 方法创建线程，调用 helloA 方法；
* 观察结果，看是否可以成功进行调用。

**实例**：

```java
public class DemoTest {
    public static void main(String[] args) {
        new Thread(new SynchronizedTest()). start();
    }
}
class SynchronizedTest implements Runnable {
    private final Object obj = new Object();
    public void helloA() { //方法1，调用方法2
        synchronized (obj) {
            System.out.println(Thread.currentThread().getName() + " helloA()");
            helloB();
        }
    }
    public void helloB() {
        synchronized (obj) {
            System.out.println(Thread.currentThread().getName() + " helloB()");
        }
    }
    @Override
    public void run() {
        helloA(); //调用helloA方法
    }
}
```

**结果验证**：

```java
Thread-0 helloA()
Thread-0 helloB()
```

**结果解析**：如果同一线程，锁不可重入的话，helloB 需要等待 helloA 释放 obj 锁，如此一来，helloB 无法进行锁的获取，最终造成无限等待，无法正常执行。此处说明了 synchronized 关键字的可重入性，因此能够正常进行两个方法的执行。

## 6. ReentrantLock 验证锁的可重入性

相同的场景，对代码进行如下改造，将 synchronized 同步代码块修改成 lock 接口同步，我们看代码实例如下：

```java
public class DemoTest {
    public static void main(String[] args) {
        new Thread(new SynchronizedTest()). start();
    }
}
class SynchronizedTest implements Runnable {
    private final Lock lock = new ReentrantLock();
    public void helloA() { //方法1，调用方法2
        lock.lock();
        try {
            System.out.println(Thread.currentThread().getName() + " helloA()");
            helloB();
        } finally {
            lock.unlock();
        }
    }
    public void helloB() {
        lock.lock();
        try {
            System.out.println(Thread.currentThread().getName() + " helloB()");
        } finally {
            lock.unlock();
        }
    }
    @Override
    public void run() {
        helloA();
    }
}
```

**结果验证**：

```java
Thread-0 helloA()
Thread-0 helloB()
```

**结果解析**：ReentrantLock 一样是可重入锁，试验成功。

## 7. 小结

锁的可重入性这一概念对于并发编程非常重要，对于本节内容需要深入的理解并掌握。我们之前已经学习过了 synchronized 关键字和 ReentrantLock 锁，此处知识用两者进行了可重入的验证。

本节关键点在于可重入性的意义所在，需要结合实例进行更加细致的理解和掌握。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
