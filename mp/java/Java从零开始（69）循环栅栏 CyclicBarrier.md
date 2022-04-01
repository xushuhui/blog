---
title: Java 从零开始（69）循环栅栏 CyclicBarrier
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 循环栅栏 CyclicBarrier

## 1. 前言

本节带领大家认识第三个常用的 Java 并发工具类之 CyclicBarrier。

本节先介绍 CyclicBarrier 工具类的表达的概念和最基本用法，接着通过一个生活中的例子为大家解释 CyclicBarrier 工具类的使用场合，然后通过简单的编码实现此场景，最后带领大家熟悉 CyclicBarrier 工具类的其他重要方法。

下面我们正式开始介绍吧。

## 2. 概念解释

所谓 Cyclic 即循环的意思，所谓 Barrier 即屏障的意思。所以综合起来，CyclicBarrier 指的就是循环屏障，虽然这个叫法很奇怪，但是却能很好地表达其含义。

CyclicBarrier 工具类允许一组线程相互等待，直到所有线程都到达一个公共的屏障点，然后这些线程一起继续执行后继逻辑。之所以称之为 “循环”，是因为在所有线程都释放了对这个屏障的使用后，这个屏障还可以重新使用。我们通过一张图可以直观了解其表达的控制模型。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyypkbmj60oq07udjz02)

现在我们已经了解了基本概念逻辑，CyclicBarrier 工具类最基本的用法是怎样的呢？看下面。

## 3. 基本用法

```java
// 创建一个 CyclicBarrier 对象，初始化相互等待的线程数量
CyclicBarrier cyclicBarrier = new CyclicBarrier（线程个数）;

// 线程 1 开始处理逻辑
...
// 线程 1 等待其他线程执行到屏障点
cyclicBarrier.await();
// 线程 1 等到了其他所有线程达到屏障点后继续处理后继逻辑
...

// 线程 n 开始处理逻辑
...
// 线程 n 等待其他线程执行到屏障点
cyclicBarrier.await();
// 线程 n 等到了其他所有线程达到屏障点后继续处理后继逻辑
...
```

是不是很简单，CyclicBarrier 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

CyclicBarrier 最适合一个由多个线程共同协作完成任务的场合。

这样描述很抽象，我们还是举一个生活中的例子说明：某学习班总共 5 位同学，约定周末一起乘坐大巴出游，约定了共同的集合地点，雇佣了 1 位司机。请看下面代码。

## 5. 场景案例

```java
import lombok.SneakyThrows;
import java.util.Random;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.CyclicBarrier;

public class CyclicBarrierTest {
    // 创建一个 Runnable 对象，用于屏障解除时处理全局逻辑，在此例子中代表大巴司机
    private static Runnable driver = new Runnable() {
        public void run() {
            System.out.println("所有同学已经集合完毕，开始启动车辆出发。");
        }
    };

    // 创建一个 CyclicBarrier 对象，初始化为 5, 代表需要控制同步的线程个数，在此例子中代表 5 位同学
    static int threadCount = 5;
    private static CyclicBarrier cyclicBarrier = new CyclicBarrier(threadCount, driver);

    public static void main(String[] args) throws InterruptedException {
        // 模拟同学
        for(int i=1; i<=threadCount; i++) {
            // 模拟某个同学的动作
            new Thread(new Runnable() {
                @SneakyThrows
                public void run() {
                    System.out.println( Thread.currentThread().getName() + "已经开始出门。..");
                    // 模拟同学出门赶往集合点的用时
                    try {
                        Thread.sleep(new Random().nextInt(10000));
                    } catch (Exception e) {}
                    System.out.println( Thread.currentThread().getName() + "已经到达集合点");
                    // 等待其他同学到达集合点（等待其他线程到达屏障点）
                    cyclicBarrier.await();
                }
            }, i + "号同学").start();
        }
    }
}
```

运行上面代码，我们观察一下运行结果。

```java
1 号同学准备出门。..
2 号同学准备出门。..
3 号同学准备出门。..
4 号同学准备出门。..
5 号同学准备出门。..
5 号同学已经到达集合点
4 号同学已经到达集合点
1 号同学已经到达集合点
2 号同学已经到达集合点
3 号同学已经到达集合点
所有同学已经集合完毕，开始启动车辆出发。
```

观察结果，和我们的预期一致。注意体会 CyclicBarrier 提供的多线程共同协作的模型。

## 6. 其他方法介绍

除过上面代码中使用的最基本的 await() 方法之外，还有下面几个方法大家可以了解一下。

1. CyclicBarrier(int parties)

相比案例中使用的 CyclicBarrier(int parties, Runnable barrierAction) 构造方法，此方法只用于控制并发线程，不做屏障点到达后的其他动作。

2. await(long timeout, TimeUnit unit) 方法

此方法可以设置等待的时限，当时限过后还未被唤起，则直接自行唤醒继续执行后继任务。

3. getNumberWaiting() 方法

调用此方法，可以获得当前还在等待屏障点解除的线程数，一般用于了解整体处理进度。

## 7. 小结

本节通过一个简单的例子，介绍了 CyclicBarrier 相关的概念原理，使用场景和基本用法。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
