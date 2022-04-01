---
title: Java 从零开始（65）原子操作之 LongAccumulator
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 原子操作之 LongAccumulator

## 1. 前言

今天为大家介绍原子操作之 LongAccumulator。此工具位于 java.util.concurrent.atomic 包中。

本节先介绍 LongAccumulator 工具类的基本概念和最基本用法，之后给出 LongAccumulator 工具类最常用的场合说明，然后通过简单的编码实现一个实际案例，最后带领大家熟悉 LongAccumulator 最常用的一些编程方法，让大家进一步加深对 LongAccumulator 工具类的理解。

下面我们正式开始介绍吧。

## 2. 概念介绍

相比 LongAdder，LongAccumulator 工具类提供了更灵活更强大的功能。不但可以指定计算结果的初始值，相比 LongAdder 只能对数值进行加减运算，LongAccumulator 还能自定义计算规则，比如做乘法运行，或其他任何你想要的计算规则。这样描述是不是有点抽象，别着急，看下面的图示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyxvnc7j60n30fegr202)

我们看下面 LongAccumulator 工具类的基本用法。

## 3. 基本用法

```java

// 首先创建一个双目运算器对象，这个对象实现了计算规则。
LongBinaryOperator longBinaryOperator = new LongBinaryOperator() {
        @Override
        public long applyAsLong(long left, long right) {
            ...
        }
    }

// 接着使用构造方法创建一个 LongAccumulator 对象，这个对象的第 1 个参数就是一个双目运算器对象，第二个参数是累加器的初始值。
LongAccumulator longAccumulator = new LongAccumulator(longBinaryOperator, 0);
...
// 调用累加方法
longAccumulator.accumulate(1000)
// 调用结果获取方法
long result = longAccumulator.get();
...
```

是不是简单又强大！LongAccumulator 在我们日常实践中，到底应该应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

LongAccumulator 经常用于自定义运算规则场景下的多线程并发场合。一些简单的累加计算可以直接使用我们之前课程中介绍的工具类，但是当运行规则比较复杂或者 JDK 没有提供对应的工具类时，可以考虑 LongAccumulator 辅助实现。当然所有可使用 LongAdder 的场合都可使用 LongAccumulator 代替，但是没有必要。

下面我们用 LongAccumulator 工具类实现上一节中的生活实例，为了简化叙述，本节我们只统计男性客户总数量。请看下面的代码。

## 5. 场景案例

```java
import java.util.concurrent.atomic.LongAccumulator;

public class LongAccumulatorTest {

	// 此处的运算规则是累加，所以创建一个加法双目运算器对象作为构造函数的第一个参数。
    // 将第二个参数置为 0，表示累加初始值。
    // maleCount 对象代表当天所有进入商场的男性客户总数量。
    private static LongAccumulator maleCount = new LongAccumulator(new LongBinaryOperator() {
	    // 此方法用于实现计算规则
        @Override
        public long applyAsLong(long left, long right) {
	        // 在本例中使用加法计算规则
            return left + right;
        }
    }, 0);

    public static void main(String[] args) {
        // 定义 30 个商场入口检测设备
        for (int i = 1; i <= 30; i++) {
            MonitoringDevice monitoringDevice = new MonitoringDevice(maleCount, i);
            // 开启检测设备进行检测
            new Thread(monitoringDevice).start();
        }
    }
}
```

在上面的代码中，首先创建一个 LongAccumulator 对象表示统计结果，然后创建了 30 个商场入口检测设备模拟检测识别，接下来每个检测设备如何动作呢，看下面的代码。

```java
import java.util.Random;
import java.util.concurrent.atomic.LongAccumulator;

/**
 * 模拟设备
 */
public class MonitoringDevice implements Runnable {
    private LongAccumulator maleCount;
    private String monitoringDeviceNo;
    public MonitoringDevice(LongAccumulator maleCount, int monitoringDeviceNo) {
        this.maleCount = maleCount;
        this.monitoringDeviceNo = "第" + monitoringDeviceNo + "监控采集处";
    }

	/**
	 * 设备运行的处理逻辑
	 */
    public void run() {
        while (true) {
            // 监测处理 （监测设备输出 1 代表男性，0 代表女性，其他代表未能识别，此处随机产生监测结果）
            try {
                Thread.sleep(new Random().nextInt(3000));
            } catch (Exception e) {}
            int monitoringDeviceOutput = new Random().nextInt(3);

            // 对监测结果进行统计
            switch (monitoringDeviceOutput) {
                case 1: maleCount.accumulate(1);
                    System.out.println("统计结果：maleCount=" + maleCount.get());
                    break;
                default:
                    System.out.println("忽略统计");
                    break;
            }
        }
    }
}
```

在 MonitoringDevice 类中，首先模拟监测设备输出，然后将输出结果使用 add () 进行统计累加，使用 sum () 输出累加结果。运行一段时间后运行结果如下。

```java
...
忽略统计
统计结果：maleCount=50
...
```

上面的示例中，使用 LongAccumulator 实现了上一节中相同的需求。对比观察，能够体会到 LongAccumulator 工具类更灵活的地方，但同时也更复杂一些。

至此，大家对 LongAccumulator 已经有了初步的理解，接下来我们继续丰富对 LongAccumulator 工具类的认识。

## 6. 核心方法介绍

除过上面代码中使用的最基本的 accumulate (int)、get () 方法之外，我们再介绍两个方法的使用。

1. reset () 方法

将累加器值置为 0，即为后继使用重新归位。

1. getThenReset () 方法

此方法逻辑等同于先调用 get () 方法再调用 reset () 方法。

## 7. 小结

本节通过一个简单的例子，介绍了 LongAccumulator 的基本用法。在 java.util.concurrent.atomic 包中还存在类似的工具类，如 DoubleAccumulator，用法大同小异，希望大家在日常研发中多比较多总结，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
