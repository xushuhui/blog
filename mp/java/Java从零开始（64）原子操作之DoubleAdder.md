---
title: Java 从零开始（64）原子操作之 DoubleAdder
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 原子操作之 DoubleAdder

## 1. 前言

今天为大家介绍原子操作之 DoubleAdder。此工具位于 java.util.concurrent.atomic 包中。

本节先介绍 DoubleAdder 工具类的基本概念和最基本用法，之后给出 DoubleAdder 工具类最常用的场合说明，然后通过简单的编码实现一个实际案例，让大家有一个理性的认识，最后带领大家熟悉 DoubleAdder 最常用的一些编程方法，进一步加深对 DoubleAdder 工具类的理解。

下面我们正式开始介绍吧。

## 2. 概念介绍

DoubleAdder 工具类采用了 “分头计算最后汇总” 的思路，避免每一次（细粒度）操作的并发控制，提高了并发的性能。什么是细粒度的同步控制呢？所谓细粒度的同步控制，指的是对待同步控制对象的每一次操作都需要加以控制，这样描述是不是有点抽象，别着急，看下面的图示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyxkv1aj60mz0er42v02)

我们看下面 DoubleAdder 工具类的基本用法。

## 3. 基本用法

```java
// 首先创建一个 DoubleAdder 对象
DoubleAdder doubleAdder = new DoubleAdder();
...
// 调用累加方法
doubleAdder.add(50.5);
doubleAdder.add(49.5);
// 调用求和方法
double sum = doubleAdder.sum();
...
```

是不是很简单，那 DoubleAdder 在我们日常实践中，到底应该应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

DoubleAdder 经常用于多线程并发做收集统计数据的场合，而不是细粒度的同步控制。

下面我们用 DoubleAdder 工具类实现一个生活案例：某商场为了掌握客流特征，在商场所有出入口架设了人体特征识别设备，此类设备可以有效识别客人性别等信息。基于此，商场管理办公室计划制作一个客流性别流量图表，用于决策商场的服务内容。

## 5. 场景案例

```java
import java.util.concurrent.atomic.DoubleAdder;

public class DoubleAdderTest {

    // 首先创建三个 DoubleAdder 对象分别表示统计结果
    // 代表当天所有进入商场的男性客户总数量
    private static DoubleAdder maleCount = new DoubleAdder();
    // 代表当天所有进入商场的女性客户总数量
    private static DoubleAdder womenCount = new DoubleAdder();
    // 代表当天所有进入商场的未能识别的客户总数量
    private static DoubleAdder unknownGenderCount = new DoubleAdder();

    public static void main(String[] args) {
        // 定义 30 个商场入口检测设备
        for (int i = 1; i <= 30; i++) {
            MonitoringDevice monitoringDevice = new MonitoringDevice(maleCount, womenCount, unknownGenderCount, i);
            // 开启检测设备进行检测
            new Thread(monitoringDevice).start();
        }
    }
}
```

在上面的代码中，首先创建三个 DoubleAdder 对象分别表示统计结果，然后创建了 30 个商场入口检测设备模拟检测识别，接下来每个检测设备如何动作呢，看下面的代码。

```java
import java.util.Random;
import java.util.concurrent.atomic.DoubleAdder;

public class MonitoringDevice implements Runnable {

    private DoubleAdder maleCount;
    private DoubleAdder womenCount;
    private DoubleAdder unknownGenderCount;

    private String monitoringDeviceNo;

    public MonitoringDevice(DoubleAdder maleCount, DoubleAdder womenCount, DoubleAdder unknownGenderCount, int monitoringDeviceNo) {
        this.maleCount = maleCount;
        this.womenCount = womenCount;
        this.unknownGenderCount = unknownGenderCount;
        this.monitoringDeviceNo = "第" + monitoringDeviceNo + "监控采集处";
    }

    public void run() {
        while (true) {
            // 监测处理 （监测设备输出 1 代表男性，0 代表女性，其他代表未能识别，此处随机产生监测结果）
            try {
                Thread.sleep(new Random().nextInt(3000));
            } catch (Exception e) {}
            int monitoringDeviceOutput = new Random().nextInt(3);

            // 对监测结果进行统计
            switch (monitoringDeviceOutput) {
                case 0: womenCount.add(1);
                    System.out.println("统计结果：womenCount=" + womenCount.sum());
                    break;
                case 1: maleCount.add(1);
                    System.out.println("统计结果：maleCount=" + maleCount.sum());
                    break;
                default: unknownGenderCount.add(1);
                    System.out.println("统计结果：unknownGenderCount=" + unknownGenderCount.sum());
                    break;
            }
        }
    }
}
```

在 MonitoringDevice 类中，首先模拟监测设备输出，然后将输出结果使用 add () 进行统计累加，使用 sum () 输出累加结果。运行一段时间后运行结果如下。

```java
...
统计结果：unknownGenderCount=23.0
统计结果：womenCount=24.0
统计结果：maleCount=32.0
...
```

上面的案例中，总共计算了三个统计值，每一个统计值都使用了多个线程同时进行统计计算。在统计过程中，每一个线程只需要累加自己的那份统计结果，所以不需要做同步控制，只要在最后进行汇总统计结果时做同步控制进行汇总即可。像这样的场景使用 DoubleAdder 工具类会非常方便简洁。

至此，大家对 DoubleAdder 已经有了初步的理解，接下来我们继续丰富对 DoubleAdder 工具类的认识。

## 6. 核心方法介绍

除过上面代码中使用的最基本的 add (int)、sum () 方法之外，我们再介绍两个方法的使用。

1. reset () 方法

将累加器值置为 0，即为后继使用重新归位。

1. sumThenReset () 方法

此方法逻辑等同于先调用 sum () 方法再调用 reset () 方法，简化代码编写。

## 7. 小结

本节通过一个简单的例子，介绍了 DoubleAdder 的基本用法。在 java.util.concurrent.atomic 包中还有一个类似的工具类 LongAdder，用法大同小异，希望大家在日常研发中多比较多总结，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
