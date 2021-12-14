---
title: Java从零开始（82）锁支持工具 LockSupport
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 锁支持工具 LockSupport


## 1. 前言

本节带领大家认识第三个常用的 Java 并发锁工具之 LockSupport。

本节先介绍 LockSupport 工具类的用途，然后介绍关键的编程方法，最后通过一个编程例子为大家展示 StampedLock 工具类的用法。

下面我们正式开始介绍吧。

## 2. 概念解释

LockSupport 是一个线程工具类，提供的方法均是静态方法，可使用类型直接调用。其定义的一组以 park 开头的方法用于阻塞当前线程，定义的 unpark 方法用于唤醒被阻塞的线程。

LockSupport 提供的这种功能，应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 3. 常用场景

LockSupport 可以在任何场合使用它阻塞线程，也可以对指定的任何线程进行唤醒，而不用担心阻塞和唤醒操作的顺序。

JDK 并发包下的锁和其他同步工具在底层实现中大量使用了 LockSupport 工具类 进行线程的阻塞和唤醒，了解其用法和原理，可以更好地理解锁和其它同步工具的底层实现。

下面我们通过一个编程例子体会一下 LockSupport 工具类的用法。

## 4. 编程案例

```java
import java.util.concurrent.locks.LockSupport;

public class LockSupportTest {

    public static void main(String[] args) throws InterruptedException {

        Thread t1 = new Thread(new Runnable() {
            public void run() {
                System.out.println(Thread.currentThread().getName() + "即将被中断");
                LockSupport.park();
                System.out.println(Thread.currentThread().getName() + "已经被中断");
            }
        });
        t1.setName("实验线程");
        // 线程启动立刻执行，在线程内调用了 park 做了线程中断
        t1.start();

        Thread.sleep(5000L);
        // 5 秒后唤醒中断的线程 t1
        LockSupport.unpark(t1);
        System.out.println(t1.getName() + "被恢复中断了");

        Thread.sleep(100000L);
    }
}
```

运行结果如下：

```java
实验线程准备被中断
实验线程已经被中断了
实验线程被恢复中断了
```

用法是不是很简单呢？但其中的原理可没有那么简单，需要我们继续深入学习。

## 5. 小结

本节解释了 LockSupport 的基本概念和常用场景，且通过一个简单的例子展示了其用法，更多关于此工具类的概念和原理介绍，可阅读 “[Java 并发原理入门教程](http://www.imooc.com/wiki/concurrencylesson/reentrantlock.html)” 。

至此，Java 并发工具课程全部结束了。本课程从 Java 并发包的结构开始讲起，然后对内容分类讲解，在讲解每一个知识点时，都解释了基本概念原理，列举了应用场景，并配示例代码辅助大家理解。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
