---
title: Java 从零开始（68）同步计数器 CountDownLatch
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 同步计数器 CountDownLatch

## 1. 前言

本节带领大家认识第二个常用的 Java 并发工具类之 CountDownLatch。

本节先介绍 CountDownLatch 工具类表达的概念和最基本用法，接着通过一个生活中的例子为大家解释 CountDownLatch 工具类的使用场合，然后通过简单的编码实现此场景，最后带领大家熟悉 CountDownLatch 工具类的其他重要方法。

下面我们正式开始介绍吧。

## 2. 概念解释

CountDownLatch 工具类从字面理解为 “倒计数锁”，其内部使用一个计数器进行实现，计数器初始值为线程的数量。当每一个线程完成自己的任务后，计数器的值就会减一。当计数器的值为 0 时，表示所有的线程都已经完成了任务，然后在 CountDownLatch 上等待的线程就可以恢复继续执行后继任务。是不是很抽象，其实很简单，看下面的图例。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyyhk0hj60qp07tq7902)

这就是 CountDownLatch 工具类的基本逻辑。概念已经了解了，CountDownLatch 工具类最基本的用法是怎样的呢？看下面。

## 3. 基本用法

```java
// 创建一个 CountDownLatch 对象
CountDownLatch countDownLatch = new CountDownLatch（子线程个数）;

// 子线程 1 开始处理逻辑
...
// 子线程执行完所有逻辑进行计数器减 1
countDownLatch.countDown();

// 子线程 n 开始处理逻辑
...
// 子线程执行完所有逻辑进行计数器减 1
countDownLatch.countDown();

// 主线程等待所有子线程执行完
countDownLatch.await();
// 主线程继续执行后继逻辑
...
```

是不是很简单，CountDownLatch 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

CountDownLatch 经常用于某一线程在开始运行前等待其他关联线程执行完毕的场合。

比如我们制作一张复杂报表，报表的各部分可以安排对应的一个线程进行计算，只有当所有线程都执行完毕后，再由最终的报表输出线程进行报表文件生成。

下面我们使用 CountDownLatch 实现这个例子。假设这张报表有 5 个部分，我们总共安排 5 个子线程分别计算，再设置 1 个报表输出线程用于最终生成报表文件。请看下面代码。

## 5. 场景案例

```java
import java.util.Random;
import java.util.concurrent.CountDownLatch;

public class CountDownLatchTest {
    // 创建一个 CountDownLatch 对象，初始化为 5, 代表需要控制同步的子线程个数
    static int threadCount = 5;
    private static CountDownLatch countDownLatch = new CountDownLatch(threadCount);
    // 报表生成主线程
    public static void main(String[] args) throws InterruptedException {
        // 定义报表子线程
        for(int i=1; i<=threadCount; i++) {
            // 开始报表子线程处理
            new Thread(new Runnable() {
                public void run() {
                    // 模拟报表数据计算时间
                    try {
                        Thread.sleep(new Random().nextInt(5000));
                    } catch (Exception e) {}
                    System.out.println( Thread.currentThread().getName() + "已经处理完毕");
                    countDownLatch.countDown();
                }
            }, "报表子线程" + i).start();
        }
        // 主线程等待所有子线程运行完毕后输出报表文件
        countDownLatch.await();
        System.out.println("报表数据已经全部计算完毕，开始生成报表文件。..");
    }
}
```

运行上面代码，我们观察一下运行结果。

```java
报表子线程 3 已经处理完毕
报表子线程 5 已经处理完毕
报表子线程 1 已经处理完毕
报表子线程 4 已经处理完毕
报表子线程 2 已经处理完毕
报表数据已经全部计算完毕，开始生成报表文件。..
```

观察结果，和我们的预期一致。注意体会 CountDownLatch 提供的多线程共同协作的模型。

## 6. 其他方法介绍

除过上面代码中使用的最基本的 countDown ()、await () 方法之外，还有两个方法大家可以了解一下。

1. await (long, TimeUnit) 方法

此方法提供了更灵活的等待参数，可以设置等待超时时间。当等待超过了设定的时限，则不再阻塞直接开始后继处理。

2. getCount () 方法

调用此方法，可获得当前计数器的数值，了解整体处理进度。

## 7. 小结

本节解释了 CountDownLatch 的基本逻辑模型，且通过一个简单的例子，介绍了 CountDownLatch 的使用场景和基本用法。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
