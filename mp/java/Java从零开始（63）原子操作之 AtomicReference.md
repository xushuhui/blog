---
title: Java 从零开始（63）原子操作之 AtomicReference
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 原子操作之 AtomicReference

## 1. 前言

今天为大家介绍原子操作之 AtomicReference。此工具位于 java.util.concurrent.atomic 包中。

本节先介绍什么是原子引用，接着展示 AtomicReference 工具类的最基本用法，之后给出 AtomicReference 工具类最常用的场合说明，然后通过简单的编码实现一个实际案例，最后带领大家熟悉 AtomicReference 最常用的一些编程方法，让大家进一步加深对 AtomicReference 工具类的理解。

下面我们正式开始介绍吧。

## 2. 概念介绍

本节介绍的 AtomicReference 工具类直译为 “原子引用”。原子操作的概念我们在之前的章节中已经介绍过了，那什么是引用呢？

引用就是为对象另起一个名字，引用对象本身指向被引用对象，对引用对象的操作都会反映到被引用对象上。在 Java 中，引用对象本身存储的是被引用对象的 “索引值”。如果对引用概念还是比较模糊，请查阅 Java 基础语法知识复习。

AtomicReference 工具类和 AtomicInteger 工具类很相似，只是 AtomicInteger 工具类是对基本类型的原子封装，而 AtomicReference 工具类是对引用类型的原子封装。我们用一张原理图展示其基本逻辑。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyxc1rkj60g00cvdj502)

我们看下面 AtomicReference 工具类的基本用法。

## 3. 基本用法

```java

// 由于 AtomicReference 是对一个对象引用的原子封装，所以首先创建一个对象
Car car1 = new Car(100, 10);

// 接着使用构造方法创建一个 AtomicReference 对象。
AtomicReference<Car> atomicReference = new AtomicReference<>(car);

...
// 当前如果是 car1 对象时，以原子方式变更引用为 car2 对象，当结果是 true 时则更新成功
Car car2 = new Car(200, 2);
Boolean bool = atomicReference.compareAndSet(car1, car2)
...
```

是不是很简单，那 AtomicReference 在我们日常实践中，到底应该应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

AtomicReference 和 AtomicInteger 非常类似，不同之处就在于 AtomicInteger 是对整数的封装，且每次只能对一个整数进行封装，而 AtomicReference 则是对普通的对象引用的封装，可将多个变量作为一个整体对象，操控多个属性的原子性的并发类。

下面我们用 AtomicReference 工具类实现生活中汽车牌照竞拍的例子：假设总共有 10 位客户参与竞拍，每位客户只有一次竞拍机会，竞拍是资格竞拍不以竞拍价格为目的。请看下面的代码。

## 5. 场景案例

```java
import java.util.concurrent.atomic.AtomicReference;

public class AtomicReferenceTest {

    // 代表待拍的车牌
    private static CarLicenseTag carLicenseTag = new CarLicenseTag(80000);
    // 创建一个 AtomicReference 对象，对车牌对象做原子引用封装
    private static AtomicReference<CarLicenseTag> carLicenseTagAtomicReference = new AtomicReference<>(carLicenseTag);
    public static void main(String[] args) {
        // 定义 5 个客户进行竞拍
        for(int i=1; i<=5; i++) {
            AuctionCustomer carAuctionCustomer = new AuctionCustomer(carLicenseTagAtomicReference, carLicenseTag, i);
            // 开始竞拍
            new Thread(carAuctionCustomer).start();
        }
    }
}

/**
 * 车牌
 */
public class CarLicenseTag {
    // 每张车牌牌号事先是固定的
    private String licenseTagNo = "沪 X66666";
    // 车牌的最新拍卖价格
    private double price = 80000.00;
    public CarLicenseTag(double price) {
        this.price += price;
    }
    public String toString() {
        return "CarLicenseTag{ licenseTagNo='" + licenseTagNo + ", price=" + price + '}';
    }
}
```

每个客户是如何动作呢，看下面的代码。

```java
import java.util.Random;
import java.util.concurrent.atomic.AtomicReference;

public class AuctionCustomer implements Runnable {
    private AtomicReference<CarLicenseTag> carLicenseTagReference;
    private CarLicenseTag carLicenseTag;
    private String customerNo;
    public AuctionCustomer(AtomicReference<CarLicenseTag> carLicenseTagReference, CarLicenseTag carLicenseTag, int customerNo) {
        this.carLicenseTagReference = carLicenseTagReference;
        this.carLicenseTag = carLicenseTag;
        this.customerNo = "第" + customerNo + "位客户";
    }

    public void run() {
        // 客户竞拍行为 （模拟竞拍思考准备时间 4 秒钟）
        try {
            Thread.sleep(new Random().nextInt(4000));
        } catch (Exception e) {}

        // 举牌更新最新的竞拍价格
        // 此处做原子引用更新
        boolean bool = carLicenseTagReference.compareAndSet(carLicenseTag,
                new CarLicenseTag(new Random().nextInt(1000)));
        System.out.println("第" + customerNo + "位客户竞拍" + bool + " 当前的竞拍信息" + carLicenseTagReference.get().toString());
    }
}
```

运行后运行结果如下。

```java
...
第第 1 位客户位客户竞拍 true 当前的竞拍信息 CarLicenseTag{ licenseTagNo='沪 X66666, price=80405.0}
第第 5 位客户位客户竞拍 false 当前的竞拍信息 CarLicenseTag{ licenseTagNo='沪 X66666, price=80405.0}
第第 3 位客户位客户竞拍 false 当前的竞拍信息 CarLicenseTag{ licenseTagNo='沪 X66666, price=80405.0}
第第 2 位客户位客户竞拍 false 当前的竞拍信息 CarLicenseTag{ licenseTagNo='沪 X66666, price=80405.0}
第第 4 位客户位客户竞拍 false 当前的竞拍信息 CarLicenseTag{ licenseTagNo='沪 X66666, price=80405.0}
...
```

至此，大家对 AtomicReference 已经有了初步的理解，接下来我们继续丰富对 AtomicReference 工具类的认识。

## 6. 核心方法介绍

除过上面代码中使用的最基本的 AtomicReference (V)、compareAndSet (int, int)、get () 方法之外，我们还再介绍两个方法的使用。下面逐个介绍。

1. set () 方法

可以使用不带参数的构造方法构造好对象后，再使用 set () 方法设置待封装的对象。等价于使用 AtomicReference (V) 构造方法。

1. getAndSet () 方法

此方法以原子方式设置为给定值，并返回旧值。逻辑等同于先调用 get () 方法再调用 set () 方法。

## 7. 小结

本节通过一个简单的例子，介绍了 AtomicReference 的基本用法。其实在 java.util.concurrent.atomic 包中还提供了更多更细场景的原子操作类，此包下的大部分工具类都是基于 CAS 原理实现，正因为如此，有很多相似之处，用法大同小异，希望大家在日常研发中多比较多总结，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
