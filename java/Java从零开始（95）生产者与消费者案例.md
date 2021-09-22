---
title: Java从零开始（95）生产者与消费者案例
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 生产者与消费者案例


## 1. 前言

本节内容是通过之前学习的 synchronized 关键字，实现多线程并发编程中最经典的生产者与消费者模式，这是本节课程的核心内容，所有的知识点都是围绕这一经典模型展开的。本节有如下知识点：

* 生产者与消费者模型介绍，这是打开本节知识大门的钥匙，也是本节内容的基础；
* 了解生产者与消费者案例实现的三种方式，我们本节以 synchronized 关键字联合 wait/notify 机制进行实现；
* wait 方法和 notify 方法介绍，这是我们实现生产者与消费者案例的技术基础；
* 生产者与消费者案例代码实现，这是我们本节内容的核心，一定要对此知识点进行深入的学习和掌握。

## 2. 生产者与消费者模型介绍

**定义**： 生产者消费者模式是一个十分经典的多线程并发协作的模式。

**意义**：弄懂生产者消费者问题能够让我们对并发编程的理解加深。

**介绍**：所谓生产者 - 消费者问题，实际上主要是包含了两类线程，一种是生产者线程用于生产数据，另一种是消费者线程用于消费数据，为了解耦生产者和消费者的关系，通常会采用共享的数据区域。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz7iumej60jg070mzp02)

共享的数据区域就像是一个仓库，生产者生产数据之后直接放置在共享数据区中，并不需要关心消费者的行为。而消费者只需要从共享数据区中去获取数据，就不再需要关心生产者的行为。

## 3. 生产者与消费者三种实现方式

在实现生产者消费者问题时，可以采用三种方式：

* 使用 Object 的 wait/notify 的消息通知机制，本节课程我们采用该方式结合 synchronized 关键字进行生产者与消费者模式的实现；
* 使用 Lock 的 Condition 的 await/signal 的消息通知机制；
* 使用 BlockingQueue 实现。本文主要将这三种实现方式进行总结归纳。

## 4. wait 与 notify

Java 中，可以通过配合调用 Object 对象的 wait () 方法和 notify () 方法或 notifyAll () 方法来实现线程间的通信。

**wait 方法**：我们之前对 wait 方法有了基础的了解，在线程中调用 wait () 方法，将阻塞当前线程，并且释放锁，直至等到其他线程调用了调用 notify () 方法或 notifyAll () 方法进行通知之后，当前线程才能从 wait () 方法出返回，继续执行下面的操作。

**notify 方法**：即唤醒，notify 方法使原来在该对象上 wait 的线程退出 waiting 状态，使得该线程从等待队列中移入到同步队列中去，等待下一次能够有机会获取到对象监视器锁。

**notifyAll 方法**：即唤醒全部 waiting 线程，与 notify 方法在效果上一致。

## 5. 生产者与消费者案例

为了更好地理解并掌握生产者与消费者模式的实现，我们先来进行场景设计，然后再通过实例代码进行实现并观察运行结果。

**场景设计**：

* 创建一个工厂类 ProductFactory，该类包含两个方法，produce 生产方法和 consume 消费方法；
* 对于 produce 方法，当没有库存或者库存达到 10 时，停止生产。为了更便于观察结果，每生产一个产品，sleep 5000 毫秒；
* 对于 consume 方法，只要有库存就进行消费。为了更便于观察结果，每消费一个产品，sleep 5000 毫秒；
* 库存使用 LinkedList 进行实现，此时 LinkedList 即共享数据内存；
* 创建一个 Producer 生产者类，用于调用 ProductFactory 的 produce 方法。生产过程中，要对每个产品从 0 开始进行编号；
* 创建一个 Consumer 消费者类，用于调用 ProductFactory 的 consume 方法；
* 创建一个测试类，main 函数中创建 2 个生产者和 3 个消费者，运行程序进行结果观察。

**实例**：创建一个工厂类 ProductFactory

```java
class ProductFactory {
    private LinkedList<String> products; //根据需求定义库存，用 LinkedList 实现
    private int capacity = 10; // 根据需求：定义最大库存 10
    public ProductFactory() {
        products = new LinkedList<String>();
    }
    // 根据需求：produce 方法创建
    public synchronized void produce(String product) {
        while (capacity == products.size()) { //根据需求：如果达到 10 库存，停止生产
            try {
                System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备生产产品，但产品池已满");
                wait(); // 库存达到 10 ，生产线程进入 wait 状态
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        products.add(product); //如果没有到 10 库存，进行产品添加
        try {
            Thread.sleep(5000); //根据需求为了便于观察结果，每生产一个产品，sleep 5000 ms
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("线程("+Thread.currentThread().getName() + ")生产了一件产品:" + product+";当前剩余商品"+products.size()+"个");
        notify(); //生产了产品，通知消费者线程从 wait 状态唤醒，进行消费
    }

    // 根据需求：consume 方法创建
    public synchronized String consume() {
        while (products.size()==0) { //根据需求：没有库存消费者进入wait状态
            try {
                System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备消费产品，但当前没有产品");
                wait(); //库存为 0 ，无法消费，进入 wait ，等待生产者线程唤醒
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        String product = products.remove(0) ; //如果有库存则消费，并移除消费掉的产品
        try {
            Thread.sleep(5000);//根据需求为了便于观察结果，每消费一个产品，sleep 5000 ms
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("线程("+Thread.currentThread().getName() + ")消费了一件产品:" + product+";当前剩余商品"+products.size()+"个");
        notify();// 通知生产者继续生产
        return product;
    }
}
```

**实例**：Producer 生产者类创建

```java
class Producer implements Runnable {
    private ProductFactory productFactory; //关联工厂类，调用 produce 方法
    public Producer(ProductFactory productFactory) {
        this.productFactory = productFactory;
    }
    public void run() {
        int i = 0 ; // 根据需求，对产品进行编号
        while (true) {
            productFactory.produce(String.valueOf(i)); //根据需求 ，调用 productFactory 的 produce 方法
            i++;
        }
    }
}
```

**实例**：Consumer 消费者类创建

```java
class Consumer implements Runnable {
    private ProductFactory productFactory;
    public Consumer(ProductFactory productFactory) {
        this.productFactory = productFactory;
    }
    public void run() {
        while (true) {
            productFactory.consume();
        }
    }
}
```

**实例**： 创建测试类，2 个生产者，3 个消费者

```java
public class DemoTest extends Thread{
    public static void main(String[] args) {
        ProductFactory productFactory = new ProductFactory();
        new Thread(new Producer(productFactory),"1号生产者"). start();
        new Thread(new Producer(productFactory),"2号生产者"). start();
        new Thread(new Consumer(productFactory),"1号消费者"). start();
        new Thread(new Consumer(productFactory),"2号消费者"). start();
        new Thread(new Consumer(productFactory),"3号消费者"). start();
    }
}
```

**结果验证**：

```java
线程(1号生产者)生产了一件产品:0;当前剩余商品1个
线程(3号消费者)消费了一件产品:0;当前剩余商品0个
警告：线程(2号消费者)准备消费产品，但当前没有产品
警告：线程(1号消费者)准备消费产品，但当前没有产品
线程(2号生产者)生产了一件产品:0;当前剩余商品1个
线程(2号消费者)消费了一件产品:0;当前剩余商品0个
警告：线程(1号消费者)准备消费产品，但当前没有产品
线程(2号生产者)生产了一件产品:1;当前剩余商品1个
线程(3号消费者)消费了一件产品:1;当前剩余商品0个
线程(1号生产者)生产了一件产品:1;当前剩余商品1个
线程(3号消费者)消费了一件产品:1;当前剩余商品0个
线程(2号生产者)生产了一件产品:2;当前剩余商品1个
线程(1号消费者)消费了一件产品:2;当前剩余商品0个
警告：线程(2号消费者)准备消费产品，但当前没有产品
线程(2号生产者)生产了一件产品:3;当前剩余商品1个
...
...
```

**结果分析**：

从结果来看，生产者线程和消费者线程合作无间，当没有产品时，消费者线程进入等待；当产品达到 10 个最大库存是，生产者进入等待。这就是经典的生产者 - 消费者模型。

## 6. 小结

实现多线程并发编程中最经典的生产者与消费者模式，这是本节课程的核心内容，所有的知识点都是围绕这一经典模型展开的。 在掌握 synchronized 关键字，wait 方法和 notify 方法的基础上，理解并掌握生产者与消费者模式是本节课程的最终目标。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
