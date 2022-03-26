---
title: Java从零开始（107）锁机制之Condition接口
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 锁机制之 Condition 接口


## 1. 前言

本节内容主要是对 Java 锁机制之 Condition 接口进行讲解，Condition 接口是配合 Lock 接口使用的，我们已经学习过 Lock 接口的相关知识，那么接下来对 Condition 接口进行讲解。本节内容的知识点如下：

* Condition 接口简介，这是我们认识 Condition 接口的基础；
* Condition 接口定义，整体上先了解 Condition 接口所包含的方法，为基础内容；
* Condition 接口所提供的方法与 Object 所提供的方法的区别与联系，此部分为本节的重点之一；
* Condition 对象的创建方式，这是开始使用 Condition 的前提，需要牢记；
* Condition 接口的常用方法使用，这是本节课程的核心内容，掌握 Condition 接口使用方式，也是我们本节课程的最终目的所在；
* 使用 ReentrantLock 与 Condition 接口，实现生产者与消费者模式。

## 2. Condition 接口简介

任意一个 Java 对象，都拥有一组监视器方法（定义在 java.lang.Object 上），主要包括 wait ()、wait (long timeout)、notify () 以及 notifyAll () 方法。这些方法与 synchronized 同步关键字配合，可以实现等待 / 通知模式。

**定义**：Condition 接口也提供了类似 Object 的监视器方法，与 Lock 配合可以实现等待 / 通知模式。Condition 可以看做是 Obejct 类的 wait ()、notify ()、notifyAll () 方法的替代品，与 Lock 配合使用。

## 3. Condition 接口定义

我们看到，从 JDK 的源码中可以获悉，Condition 接口包含了如下的方法，对于其中常用的方法，我们在后续的内容中会有比较详细的讲解。

```java
public interface Condition {
    void await() throws InterruptedException;
    long awaitNanos(long nanosTimeout) throws InterruptedException;
    boolean await(long time, TimeUnit unit) throws InterruptedException;
    boolean awaitUntil(Date deadline) throws InterruptedException;
    void signal();
    void signalAll();
}
```

## 4. Condition 方法与 Object 方法的联系与区别

**联系 1**：都有一组类似的方法。

* **Object 对象监视器**: Object.wait()、Object.wait(long timeout)、Object.notify()、Object.notifyAll()。
* **Condition 对象**: Condition.await()、Condition.awaitNanos(long nanosTimeout)、Condition.signal()、Condition.signalAll()。

**联系 2**：都需要和锁进行关联。

* **Object 对象监视器**: 需要进入 synchronized 语句块（进入对象监视器）才能调用对象监视器的方法。
* **Condition 对象**: 需要和一个 Lock 绑定。

**区别**：

* Condition 拓展的语义方法，如 awaitUninterruptibly () 等待时忽略中断方法；
* 在使用方法时，Object 对象监视器是进入 synchronized 语句块（进入对象监视器）后调用 Object.wait ()。而 Condition 对象需要和一个 Lock 绑定，并显示的调用 lock () 获取锁，然后调用 Condition.await ()；
* 从等待队列数量看，Object 对象监视器是 1 个。而 Condition 对象是多个。可以通过多次调用 lock.newCondition () 返回多个等待队列。

## 5. Condition 对象的创建

Condition 对象是由 Lock 对象创建出来的 (Lock.newCondition)，换句话说，Condition 是依赖 Lock 对象的。那么我们来看看如果创建 Condition 对象。

此处仅提供示例代码，后续我们在进行方法讲解时，会有全部的代码示例，但在学习使用方法之前，我们必须先学会如何创建。

```java
Lock lock = new ReentrantLock();
Condition condition1 = lock.newCondition();
Condition condition2 = lock.newCondition();
```

## 6. Condition 方法介绍

**等待机制方法简介**：

* **void await() throws InterruptedException**：当前线程进入等待状态，直到被其它线程的唤醒继续执行或被中断；
* **void awaitUninterruptibly()**：当前线程进入等待状态，直到被其它线程被唤醒；
* **long awaitNanos(long nanosTimeout) throws InterruptedException**：当前线程进入等待状态，直到被其他线程唤醒或被中断，或者指定的等待时间结束；nanosTimeout 为超时时间，返回值 = 超时时间 - 实际消耗时间；
* **boolean await(long time, TimeUnit unit) throws InterruptedException**：当前线程进入等待状态，直到被其他线程唤醒或被中断，或者指定的等待时间结束；与上个方法区别：可以自己设置时间单位，未超时被唤醒返回 true，超时则返回 false；
* **boolean awaitUntil(Date deadline) throws InterruptedException**：当前线程等待状态，直到被其他线程唤醒或被中断，或者指定的截止时间结束，截止时间结束前被唤醒，返回 true，否则返回 false。

**通知机制方法简介**：

* **void signal()**：唤醒一个线程；
* **void signalAll()**：唤醒所有线程。

## 7. ReentrantLock 与 Condition 实现生产者与消费者

非常熟悉的场景设计，这是我们在讲解生产者与消费者模型时使用的案例设计，那么此处有细微的修改如下，请学习者进行比照学习，印象更加深刻。

**场景修改**：

* 创建一个工厂类 ProductFactory，该类包含两个方法，produce 生产方法和 consume 消费方法（未改变）；
* 对于 produce 方法，当没有库存或者库存达到 10 时，停止生产。为了更便于观察结果，每生产一个产品，sleep 3000 毫秒（5000 变 3000，调用地址也改变了，具体看代码）；
* 对于 consume 方法，只要有库存就进行消费。为了更便于观察结果，每消费一个产品，sleep 5000 毫秒（sleep 调用地址改变了，具体看代码）；
* 库存使用 LinkedList 进行实现，此时 LinkedList 即共享数据内存（未改变）；
* 创建一个 Producer 生产者类，用于调用 ProductFactory 的 produce 方法。生产过程中，要对每个产品从 0 开始进行编号 （新增 sleep 3000ms)；
* 创建一个 Consumer 消费者类，用于调用 ProductFactory 的 consume 方法 （新增 sleep 5000ms)；
* 创建一个测试类，main 函数中创建 2 个生产者和 3 个消费者，运行程序进行结果观察（未改变）。

**实例**：

```java
public class DemoTest {
        public static void main(String[] args) {
            ProductFactory productFactory = new ProductFactory();
            new Thread(new Producer(productFactory),"1号生产者"). start();
            new Thread(new Producer(productFactory),"2号生产者"). start();
            new Thread(new Consumer(productFactory),"1号消费者"). start();
            new Thread(new Consumer(productFactory),"2号消费者"). start();
            new Thread(new Consumer(productFactory),"3号消费者"). start();
        }
}

class ProductFactory {
    private LinkedList<String> products; //根据需求定义库存，用 LinkedList 实现
    private int capacity = 10; // 根据需求：定义最大库存 10
    private Lock lock = new ReentrantLock(false);
    private Condition p = lock.newCondition();
    private Condition c = lock.newCondition();
    public ProductFactory() {
        products = new LinkedList<String>();
    }
    // 根据需求：produce 方法创建
    public void produce(String product) {
        try {
            lock.lock();
            while (capacity == products.size()) { //根据需求：如果达到 10 库存，停止生产
                try {
                    System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备生产产品，但产品池已满");
                    p.await(); // 库存达到 10 ，生产线程进入 wait 状态
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            products.add(product); //如果没有到 10 库存，进行产品添加
            System.out.println("线程("+Thread.currentThread().getName() + ")生产了一件产品:" + product+";当前剩余商品"+products.size()+"个");
            c.signalAll(); //生产了产品，通知消费者线程从 wait 状态唤醒，进行消费
        } finally {
            lock.unlock();
        }
    }

    // 根据需求：consume 方法创建
    public String consume() {
        try {
            lock.lock();
            while (products.size()==0) { //根据需求：没有库存消费者进入wait状态
                try {
                    System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备消费产品，但当前没有产品");
                    c.await(); //库存为 0 ，无法消费，进入 wait ，等待生产者线程唤醒
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            String product = products.remove(0) ; //如果有库存则消费，并移除消费掉的产品
            System.out.println("线程("+Thread.currentThread().getName() + ")消费了一件产品:" + product+";当前剩余商品"+products.size()+"个");
            p.signalAll();// 通知生产者继续生产
            return product;
        } finally {
            lock.unlock();
        }
    }
}

class Producer implements Runnable {
    private ProductFactory productFactory; //关联工厂类，调用 produce 方法
    public Producer(ProductFactory productFactory) {
        this.productFactory = productFactory;
    }
    public void run() {
        int i = 0 ; // 根据需求，对产品进行编号
        while (true) {
            productFactory.produce(String.valueOf(i)); //根据需求 ，调用 productFactory 的 produce 方法
            try {
                Thread.sleep(3000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            i++;
        }
    }
}
class Consumer implements Runnable {
    private ProductFactory productFactory;
    public Consumer(ProductFactory productFactory) {
        this.productFactory = productFactory;
    }
    public void run() {
        while (true) {
            productFactory.consume();
            try {
                Thread.sleep(5000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
```

**结果验证**：

```java
线程(1号生产者)生产了一件产品:0;当前剩余商品1个
线程(2号生产者)生产了一件产品:0;当前剩余商品2个
线程(1号消费者)消费了一件产品:0;当前剩余商品1个
线程(2号消费者)消费了一件产品:0;当前剩余商品0个
警告：线程(3号消费者)准备消费产品，但当前没有产品
线程(2号生产者)生产了一件产品:1;当前剩余商品1个
线程(1号生产者)生产了一件产品:1;当前剩余商品2个
线程(3号消费者)消费了一件产品:1;当前剩余商品1个
线程(2号消费者)消费了一件产品:1;当前剩余商品0个
警告：线程(1号消费者)准备消费产品，但当前没有产品
```

## 8. 小结

本节内容为主要对 Condition 接口进行了讲解，Condition 接口作为 Lock 接口的监视器，是非常重要的接口，我们需要非常重视 Condition 接口的学习。

本节内容最终的目的是使用 Condition 接口和 Lock 配合实现案例，核心内容即为 Condition 接口的使用，请翻看生产者与消费者一节，对比进行学习。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
