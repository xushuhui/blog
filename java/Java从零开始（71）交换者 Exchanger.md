# 交换者 Exchanger

## 1. 前言

本节带领大家认识第五个常用的 Java 并发工具类之 Exchanger。

本节先介绍 Exchanger 工具类表达的概念和最基本用法，接着通过一个生活中的例子为大家解释 Exchanger 工具类的使用场合，然后通过简单的编码实现此场景，最后带领大家熟悉 Exchanger 工具类的其他重要方法。

下面我们正式开始介绍吧。

## 2. 概念解释

Exchanger 表示 “交换者”，此工具类提供了两个线程在某个时间点彼此交换信息的功能。使用 Exchanger 的重点是成对的线程使用 exchange () 方法，当一对线程都到达了同步点时，彼此会进行信息交换。我们通过一张图可直观了解其逻辑。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz0rnrzj60k708r77g02)

概念已经了解了，Exchanger 工具类最基本的用法是怎样的呢？看下面。

## 3. 基本用法

```java
// 创建一个 Exchanger 对象
Exchanger exchanger = new Exchanger();

// 线程1开始处理逻辑
...
// 线程1将自己的信息交换给线程2，并一直等待线程2的交换动作
exchanger.exchange("待交换信息");
// 线程1等到了线程2的交换结果后继续处理后继逻辑
...

// 线程2开始处理逻辑
...
// 线程2将自己的信息交换给线程2，并一直等待线程1的交换动作
exchanger.exchange("待交换信息");
// 线程2等到了线程1的交换结果后继续处理后继逻辑
...
```

是不是很简单，那 Exchanger 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

Exchanger 工具类提供了成对的线程彼此同步数据的场合。我们举一个生活中的例子说明：快递员为客户派送物品，客户要求订单采用货到付款的方式进行支付。当快递员送货上门后，出示收款二维码（或者 POM 刷卡支付），客户当面扫码（或刷卡）支付。在这个例子中，快递员交换出去的是货物收到的是款项，而客户正好相反。我们用 Exchanger 工具类简单实现这个场景，请看下面代码。

## 5. 场景案例

```java
import java.util.Random;
import java.util.concurrent.Exchanger;

public class ExchangerTest {

    // 创建一个 Exchanger 对象
    private static Exchanger<Object> exchanger = new Exchanger();

    public static void main(String[] args) throws InterruptedException {
        // 模拟快递员
        new Thread(() -> {
            System.out.println( Thread.currentThread().getName() + "送货上门中...");
            // 模拟快递送货用时
            try {
                Thread.sleep(new Random().nextInt(10000));
            } catch (Exception e) {}
            System.out.println( Thread.currentThread().getName() + "货物已经送到，等待客户付款");
            // 进行货款交换
            try {
                Object money = exchanger.exchange("快递件");
                // 收到货款
                System.out.println("已经收到货款" + money + "，继续下一单派送...");
            } catch (Exception e) {}
        }, "快递员").start();

        // 模拟客户
        new Thread(() -> {
            System.out.println( Thread.currentThread().getName() + "工作中...");
            // 模拟工作中用时
            try {
                Thread.sleep(new Random().nextInt(10000));
            } catch (Exception e) {}
            System.out.println( Thread.currentThread().getName() + "接到快递员取件电话，货物已经送到");
            try {
                // 进行货款交换
                Object packagz = exchanger.exchange("1000元");
                // 收到货款
                System.out.println("已经收到货物" + packagz + "...");
            } catch (Exception e) {}
        }, "客户").start();

    }
}
```

运行上面代码，我们观察一下运行结果。

```java
快递员送货上门中...
客户工作中...
快递员货物已经送到，等待客户付款
客户接到快递员取件电话，货物已经送到
已经收到货物快递件...
已经收到货款1000元，继续下一单派送...
```

观察结果，和我们的预期一致。

## 6. 其他方法介绍

Exchanger 工具类的用法比较简单，其提供了两个 exchange 方法，除过上面代码中使用的方法之外，其还对进行了重载。

exchange (V, timeout, TimeUnit) 方法。

允许设置交换等待的超时时间，当时间过后还未交换到需要的对方数据，则不再等待，继续后继逻辑执行。

## 7. 小结

本节通过一个简单的例子，介绍了 Exchanger 的使用场景和基本用法。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
