---
title: Java 从零开始（70）移相器 Phaser
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 移相器 Phaser

## 1. 前言

本节带领大家认识第四个常用的 Java 并发工具类之 Phaser。

本节先介绍 Phaser 工具类表达的概念和最基本用法，接着通过一个生活中的例子为大家解释 Phaser 工具类的使用场合，然后通过简单的编码实现此场景，最后带领大家熟悉 Phaser 工具类的其他重要方法。

下面我们正式开始介绍吧。

## 2. 概念解释

Phaser 表示 “阶段器”，一个可重用的同步 barrier，与 CyclicBarrier 相比，Phaser 更灵活，而且侧重于 “重用”。Phaser 中允许 “注册的同步者（parties）” 随时间而变化。Phaser 可以通过构造器初始化 parties 个数，也可以在 Phaser 运行期间随时加入新的 parties，以及在运行期间注销 parties。

是不是又强大又抽象，没关系，我们通过一张图可以直白了解其提供的逻辑模型。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz0gg9ij60uw08rwl602)

概念已经了解了，Phaser 工具类最基本的用法是怎样的呢？看下面。

## 3. 基本用法

```java
// 创建一个 Phaser 对象
Phaser phaser = new Phaser();

// 将线程 m 作为同步者之一进行同步控制注册
phaser.register();
// 线程 m 开始处理逻辑
...
// 线程 m 等待同一周期内，其他线程到达，然后进入新的周期，并继续同步进行
phaser.arriveAndAwaitAdvance();
...

// 线程 m 执行完毕后做同步控制注销
phaser.arriveAndDeregister();
```

这个工具类相对而言比较复杂，大家不要着急，结合后面的案例仔细体会。Phaser 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

Phaser 适合用于具有多阶段处理的任务，在每个阶段有多个线程并行处理的场景。这样描述很抽象，我们举一个生活中的例子：有一个开发小组总共 4 个人，约定一起去旅游。计划一起出发，先去景点 A 自有活动，3 个小时后去景点 B 自有活动，2 个小时候后活动结束统一集合。这个场景中 4 个人相当于 4 个线程，分了 4 个阶段完成了整个计划。像类似这样的场景很适合用 Phaser 解决。请看下面代码。

## 5. 场景案例

```java
public class PhaserTest {
	// 先构建一个阶段器对象
    private static TravelPhaser travelPhaser = new TravelPhaser();
	// 主逻辑
    public static void main(String[] args) throws InterruptedException {
	    // 创建 5 个线程代表每一位同事
        for (int i = 1; i < 5; i++) {
            // 对每一个需要同步控制的线程进行同步控制注册
            travelPhaser.register();
            // 模拟每一位同事开始旅游行动
            Thread thread = new Thread(new Colleague(travelPhaser), "同事" + i);
            thread.start();
        }
    }
}
```

上述代码在注册好需要同步控制的所有线程之后，开启了每一个线程（每位同事）的处理。每一个线程（每位同事）如何行动呢，代码如下：

```java
import java.util.Random;

/**
 * 模拟人以及旅游的各类状态
 */
public class Colleague implements Runnable {
    private TravelPhaser travelPhaser;
	public Colleague(TravelPhaser travelPhaser) {
        this.travelPhaser = travelPhaser;
    }

	/**
	 * 模拟每位同事的动作
	 */
    @Override
    public void run() {
        doAnything();
        System.out.println(Thread.currentThread().getName() + "到达出发集合地");
        travelPhaser.arriveAndAwaitAdvance();

        doAnything();
        System.out.println(Thread.currentThread().getName() + "已经在景点 A 自由活动结束");
        travelPhaser.arriveAndAwaitAdvance();

        doAnything();
        System.out.println(Thread.currentThread().getName() + "已经在景点 B 自由活动结束");
        travelPhaser.arriveAndAwaitAdvance();

        doAnything();
        System.out.println(Thread.currentThread().getName() + "到达返程集合地");
        travelPhaser.arriveAndAwaitAdvance();
    }

	/**
	 * 模拟用时
	 */
    private void doAnything() {
        try {
            Thread.sleep(new Random().nextInt(10000));
        } catch (Exception e) {}
    }
}
```

上述代码模拟了每位同事的旅游过程。代码中使用了 arriveAndAwaitAdvance () 进行每个旅游阶段的控制。我们再接着看对旅游各个阶段的自定义控制：

```java
import java.util.concurrent.Phaser;

/**
 * 对每一个阶段进行自定义控制
 */
public class TravelPhaser extends Phaser {

    protected boolean onAdvance(int phase, int registeredParties) {
        switch (phase) {
            // 第 1 阶段，旅游前的集合
            case 0:
                System.out.println("出发前小组人员集合完毕，总人数："+getRegisteredParties());
                return false;
            // 第 2 阶段，景点 A 游玩
            case 1:
                System.out.println("景点 A 游玩结束");
                return false;
            // 第 3 阶段，景点 B 游玩
            case 2:
                System.out.println("景点 B 游玩结束");
                return false;
            // 第 4 阶段，旅游结束返程集合
            case 3:
                System.out.println("所有活动结束后小组人员集合完毕，总人数："+getRegisteredParties());
                return true;
            default:
                return true;
        }
    }
}
```

上述代码只是在各个阶段打印了一些描述信息，实际中可以做更多的逻辑控制。运行上面代码，我们观察一下运行结果。

```java
同事 1 到达出发集合地
同事 4 到达出发集合地
同事 2 到达出发集合地
同事 3 到达出发集合地
出发前小组人员集合完毕，总人数：4
同事 3 已经在景点 A 自由活动结束
同事 2 已经在景点 A 自由活动结束
同事 1 已经在景点 A 自由活动结束
同事 4 已经在景点 A 自由活动结束
景点 A 游玩结束
同事 4 已经在景点 B 自由活动结束
同事 2 已经在景点 B 自由活动结束
同事 1 已经在景点 B 自由活动结束
同事 3 已经在景点 B 自由活动结束
景点 B 游玩结束
同事 2 到达返程集合地
同事 3 到达返程集合地
同事 1 到达返程集合地
同事 4 到达返程集合地
所有活动结束后小组人员集合完毕，总人数：4
```

观察结果，和我们的预期一致。注意体会 Phaser 提供的多线程共同协作的模型。

## 6. 其他方法介绍

除过上面代码中使用的最基本的 register ()、arriveAndAwaitAdvance ()、arriveAndDeregister ()、getRegisteredParties () 方法之外，还有下面几个方法大家可以了解一下。

1. awaitAdvance (int phase) 方法。

具有阻塞功能，等待 phase 周期数下其他所有的 parties 都到达后返回。如果指定的 phase 与当前的 phase 不一致，则立即返回。

2. awaitAdvanceInterruptibly (int phase) 方法。

同 awaitAdvance 类似，但支持中断响应，即 waiter 线程如果被外部中断，则此方法立即返回。

3. forceTermination () 方法。

用于强制终止 phase，此后 Phaser 对象将不可用，即 register 等将不再有效。

## 7. 小结

本节介绍了 Phaser 基本概念原理，并且通过一个简单的例子，介绍了其使用场景和基本用法。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
