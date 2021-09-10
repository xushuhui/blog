# 多线程售票案例

## 1. 前言

本节内容主要是使用 Java 的锁机制对多线程售票案例进行实现。售票案例多数情况下主要关注多线程如何安全的减少库存，也就是剩余的票数，当票数为 0 时，停止减少库存。

本节内容除了关注车票库存的减少，还会涉及到退票窗口，能够更加贴切的模拟真实的场景。

本节内容需要学习者关注如下两个重点：

* 掌握多线程的售票机制模型，在后续的工作中如果涉及到类似的场景，能够第一时间了解场景的整体结构；
* 使用 Condition 和 Lock 实现售票机制，巩固我们本章节内容所学习的新的锁机制。

## 2. 售票机制模型

售票机制模型是源于现实生活中的售票场景，从开始的单窗口售票到多窗口售票，从开始的人工统计票数到后续的系统智能在线售票。多并发编程能够实现这一售票场景，多窗口售票情况下保证线程的安全性和票数的正确性。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvwgp85j60jg0awwgs02)

如上图所示，有两个售票窗口进行售票，有一个窗口处理退票，这既是现实生活中一个简单的售票机制。

## 3. 售票机制实现

**场景设计**：

* 创建一个工厂类 TicketCenter，该类包含两个方法，saleRollback 退票方法和 sale 售票方法；
* 定义一个车票总数等于 10 ，为了方便观察结果，设置为 10。学习者也可自行选择数量；
* 对于 saleRollback 方法，当发生退票时，通知售票窗口继续售卖车票；
* 对 saleRollback 进行特别设置，每隔 5000 毫秒退回一张车票；
* 对于 sale 方法，只要有车票就进行售卖。为了更便于观察结果，每卖出一张车票，sleep 2000 毫秒；
* 创建一个测试类，main 函数中创建 2 个售票窗口和 1 个退票窗口，运行程序进行结果观察。
* 修改 saleRollback 退票时间，每隔 25 秒退回一张车票；
* 再次运行程序并观察结果。

**实现要求**：本实验要求使用 ReentrantLock 与 Condition 接口实现同步机制。

**实例**：

```java
public class DemoTest {
        public static void main(String[] args) {
            TicketCenter ticketCenter = new TicketCenter();
            new Thread(new saleRollback(ticketCenter),"退票窗口"). start();
            new Thread(new Consumer(ticketCenter),"1号售票窗口"). start();
            new Thread(new Consumer(ticketCenter),"2号售票窗口"). start();
        }
}

class TicketCenter {
    private int capacity = 10; // 根据需求：定义10涨车票
    private Lock lock = new ReentrantLock(false);
    private Condition saleLock = lock.newCondition();
    // 根据需求：saleRollback 方法创建，为退票使用
    public void saleRollback() {
        try {
            lock.lock();
            capacity++;
            System.out.println("线程("+Thread.currentThread().getName() + ")发生退票。" + "当前剩余票数"+capacity+"个");
            saleLock.signalAll(); //发生退票，通知售票窗口进行售票
        } finally {
            lock.unlock();
        }
    }

    // 根据需求：sale 方法创建
    public void sale() {
        try {
            lock.lock();
            while (capacity==0) { //没有票的情况下，停止售票
                try {
                    System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备售票，但当前没有剩余车票");
                    saleLock.await(); //剩余票数为 0 ，无法售卖，进入 wait
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            capacity-- ; //如果有票，则售卖 -1
            System.out.println("线程("+Thread.currentThread().getName() + ")售出一张票。" + "当前剩余票数"+capacity+"个");
        } finally {
            lock.unlock();
        }
    }
}

class saleRollback implements Runnable {
    private TicketCenter TicketCenter; //关联工厂类，调用 saleRollback 方法
    public saleRollback(TicketCenter TicketCenter) {
        this.TicketCenter = TicketCenter;
    }
    public void run() {
        while (true) {
            try {
                Thread.sleep(5000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            TicketCenter.saleRollback(); //根据需求 ，调用 TicketCenter 的 saleRollback 方法

        }
    }
}
class Consumer implements Runnable {
    private TicketCenter TicketCenter;
    public Consumer(TicketCenter TicketCenter) {
        this.TicketCenter = TicketCenter;
    }
    public void run() {
        while (true) {
            TicketCenter.sale(); //调用sale 方法
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
```

**结果验证**：

```java
线程(1号售票窗口)售出一张票。当前剩余票数9个
线程(2号售票窗口)售出一张票。当前剩余票数8个
线程(2号售票窗口)售出一张票。当前剩余票数7个
线程(1号售票窗口)售出一张票。当前剩余票数6个
线程(1号售票窗口)售出一张票。当前剩余票数5个
线程(2号售票窗口)售出一张票。当前剩余票数4个
线程(退票窗口)发生退票。当前剩余票数5个
线程(1号售票窗口)售出一张票。当前剩余票数4个
线程(2号售票窗口)售出一张票。当前剩余票数3个
线程(2号售票窗口)售出一张票。当前剩余票数2个
线程(1号售票窗口)售出一张票。当前剩余票数1个
线程(退票窗口)发生退票。当前剩余票数2个
线程(1号售票窗口)售出一张票。当前剩余票数1个
线程(2号售票窗口)售出一张票。当前剩余票数0个
警告：线程(1号售票窗口)准备售票，但当前没有剩余车票
警告：线程(2号售票窗口)准备售票，但当前没有剩余车票
线程(退票窗口)发生退票。当前剩余票数1个
线程(1号售票窗口)售出一张票。当前剩余票数0个
警告：线程(2号售票窗口)准备售票，但当前没有剩余车票
警告：线程(1号售票窗口)准备售票，但当前没有剩余车票
```

**结果分析**：从结果来看，我们正确的完成了售票和退票的机制，并且使用了 ReentrantLock 与 Condition 接口。

**代码片段分析 1**：看售票方法代码。

```java
public void sale() {
        try {
            lock.lock();
            while (capacity==0) { //没有票的情况下，停止售票
                try {
                    System.out.println("警告：线程("+Thread.currentThread().getName() + ")准备售票，但当前没有剩余车票");
                    saleLock.await(); //剩余票数为 0 ，无法售卖，进入 wait
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            capacity-- ; //如果有票，则售卖 -1
            System.out.println("线程("+Thread.currentThread().getName() + ")售出一张票。" + "当前剩余票数"+capacity+"个");
        } finally {
            lock.unlock();
        }
    }
```

主要来看方法中仅仅使用了 await 方法，因为退票是场景触发的，售票窗口无需唤醒退票窗口，因为真实的场景下，可能没有退票的发生，所以无需唤醒。这与生产者与消费者模式存在着比较明显的区别。

**代码片段分析 2**：看退票方法代码。

```java
public void saleRollback() {
        try {
            lock.lock();
            capacity++;
            System.out.println("线程("+Thread.currentThread().getName() + ")发生退票。" + "当前剩余票数"+capacity+"个");
            saleLock.signalAll(); //发生退票，通知售票窗口进行售票
        } finally {
            lock.unlock();
        }
    }
```

退票方法只有 signalAll 方法，通知售票窗口进行售票，无需调用 await 方法，因为只要有退票的发生，就能够继续售票，没有库存上限的定义，这也是与生产者与消费者模式的一个主要区别。

**总结**：售票机制与生产者 - 消费者模式存在着细微的区别，需要学习者通过代码的实现慢慢体会。由于售票方法只需要进入 await 状态，退票方法需要唤醒售票的 await 状态，因此只需要创建一个售票窗口的 Condition 对象。

## 4. 小结

本节内容主要对售票机制模型进行了讲解，核心内容为售票机制的实现。实现的过程使用 ReentrantLock 与 Condition 接口实现同步机制，也是本节课程的重点知识。

至此，并发编程原理课程就结束了，从基本的多线程的创建，synchronized 关键字的使用，再到锁的应用，贯穿了并发编程原理的全部知识，学习完本套课程，可以进一步学习并发编程包的课程。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
