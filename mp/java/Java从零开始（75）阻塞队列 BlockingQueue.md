---
title: Java从零开始（75）阻塞队列 BlockingQueue
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 阻塞队列 BlockingQueue


## 1. 前言

本节带领大家认识第三个常用的 Java 并发容器类之 BlockingQueue。

本节先介绍 BlockingQueue 工具类表达的概念和最基本用法，接着通过一个例子为大家解释 BlockingQueue 工具类的使用场合，然后通过简单的编码实现此场景。

下面我们正式开始介绍吧。

## 2. 概念解释

BlockingQueue 顾名思义，就是阻塞队列。队列的概念同我们日常生活中的队列一样。在计算机中，队列具有先入先出的特征，不允许插队的情况出现。在 Java 中，BlockingQueue 是一个 interface 而非 class，它有很多实现类，如 ArrayBlockingQueue、LinkedBlockingQueue 等，这些实现类之间主要区别体现在存储结构或元素操作上，但入队和出队操作却是类似的。

概念已经了解了，BlockingQueue 工具类最基本的用法是怎样的呢，我们以 LinkedBlockingQueue 实现类为例说明？看下面。

## 3. 基本用法

```java
// 创建一个 LinkedBlockingQueue 对象
LinkedBlockingQueue<String> linkedBlockingQueue = new LinkedBlockingQueue();
// 在不违反容量限制的情况下，可立即将指定元素插入此队列，成功返回 true，当无可用空间时候，返回 IllegalStateException 异常。
linkedBlockingQueue.add("car");
// 在不违反容量限制的情况下，可立即将指定元素插入此队列，成功返回 true，当无可用空间时候，返回 false。
linkedBlockingQueue.offer("car");
// 直接在队列中插入元素，当无可用空间时候，阻塞等待。
linkedBlockingQueue.put("car");
// 将给定元素在给定的时间内设置到队列中，如果设置成功返回 true，否则返回 false，超时后返回 fase。
linkedBlockingQueue.offer("car", 10, Timeunit.SECONDS);
// 获取并移除队列头部的元素，无元素时候阻塞等待。
linkedBlockingQueue.take();
// 获取并移除队列头部的元素，无元素时候阻塞等待指定时间。超时后返回 null。
linkedBlockingQueue.poll(10, Timeunit.SECONDS);
```

是不是很简单，那 BlockingQueue 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

BlockingQueue 首先作为一个队列，可以适用于任何需要队列数据结构的场合，其次其具有阻塞操作的特征，可用于线程间协同操作的场合。日常研发中，生产者消费者模型常常使用 BlockingQueue 实现。我们通过一张图了解其基本逻辑。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz1q4stj60qw084dk802)

我们举一个生活中汽车排队加油的例子说明：每一个加油站台就是一个阻塞队列，汽车依次排队进入，先进入的先出站。当站台满了后继车辆就需要排队等待，当前面的汽车加好油离开（出队）后，后面的汽车进入（入队）开始加油。我们用 LinkedBlockingQueue 工具类实现，请看下面代码。

## 5. 场景案例

```java
import lombok.SneakyThrows;

import java.util.Random;
import java.util.concurrent.LinkedBlockingQueue;

public class BlockingQueueTest {

    // 创建一个 LinkedBlockingQueue 对象，用于汽车排队
    private static LinkedBlockingQueue<String> linkedBlockingQueue = new LinkedBlockingQueue<>();

    // 主线程
    public static void main(String[] args) throws InterruptedException {
		// 汽车
        int carNumber = 1;
        while(carNumber < 5) {
            new Thread(new Runnable() {
                @SneakyThrows
                public void run() {
                    // 模拟用时
                    Thread.sleep(new Random().nextInt(1000));
                    // 汽车进站排队等待
                    linkedBlockingQueue.offer(Thread.currentThread().getName());
                    System.out.println(Thread.currentThread().getName() + "：已经排队进入收费站台，等候收费...");
                }
            }, "汽车" + carNumber++).start();
        }

        // 收费员
        new Thread(new Runnable() {
            @SneakyThrows
            public void run() {
                while(true) {
                    // 模拟用时
                    Thread.sleep(new Random().nextInt(1000));
                    // 汽车过收费后出站
                    String car = linkedBlockingQueue.take();
                    System.out.println(Thread.currentThread().getName() + "：汽车" + car + "已经收费完毕离开收费站台");
                }
            }
        }, "收费员").start();

        // 信息展示
        new Thread(new Runnable() {
            @SneakyThrows
            public void run() {
                while(true) {
                    Thread.sleep(1000);
                    // 实时公示当前车流排队情况
                    int howMany = linkedBlockingQueue.size();
                    System.out.println(Thread.currentThread().getName() + "：当前还" + howMany + "辆在等候过站");
                }
            }
        }, "大屏").start();

        Thread.sleep(1000000);
    }
}
```

运行上面代码，我们观察一下运行结果。

```java
汽车1：已经排队进入收费站台，等候收费...
汽车2：已经排队进入收费站台，等候收费...
汽车4：已经排队进入收费站台，等候收费...
收费员：汽车汽车1已经收费完毕离开收费站台
汽车3：已经排队进入收费站台，等候收费...
大屏：当前还3辆在等候过站
收费员：汽车汽车2已经收费完毕离开收费站台
大屏：当前还2辆在等候过站
收费员：汽车汽车4已经收费完毕离开收费站台
收费员：汽车汽车3已经收费完毕离开收费站台
大屏：当前还0辆在等候过站
```

观察结果，和我们的预期一致。

## 6. 小结

本节通过一个简单的例子，介绍了 BlockingQueue 接口以及一个具体的实现类 LinkedBlockingQueue 的使用场景和基本用法，其他实现了 BlockingQueue 接口的实现类都类似。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
