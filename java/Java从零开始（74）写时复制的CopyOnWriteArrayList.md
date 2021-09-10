# 写时复制的 CopyOnWriteArrayList

## 1. 前言

本节带领大家认识第二个常用的 Java 并发容器类之 CopyOnWriteArrayList。

本节先介绍 CopyOnWriteArrayList 工具类表达的概念和最基本用法，接着通过一个生活中的例子为大家解释 CopyOnWriteArrayList 工具类的使用场合，然后通过简单的编码实现此场景。

下面我们正式开始介绍吧。

## 2. 概念解释

什么是 CopyOnWrite ？ 顾名思义，就是 “写数据的时候先拷贝一份副本，在副本上写数据”。为什么需要在写的时候以这种方式执行呢？当然是为了提高效率。

当多个线程同时操作一个 ArrayList 对象时，为了线程安全需要对操作增加线程安全相关的锁控制。采用 CopyOnWrite 方式，可以做到读操作不用加锁，而只对写操作加锁，且可以很方便地反馈写后的结果给到读操作。CopyOnWriteArrayList 就是采用这种优化思想，对 ArrayList 做的线程安全特性增强。我们通过一张图了解其基本原理。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz1d03qj60ol0h345e02)

概念已经了解了，CopyOnWriteArrayList 工具类最基本的用法是怎样的呢？看下面。

## 3. 基本用法

此工具类和 ArrayList 在使用方式方面很类似。

```java
// 创建一个 CopyOnWriteArrayList 对象
CopyOnWriteArrayList phaser = new CopyOnWriteArrayList();
// 新增
copyOnWriteArrayList.add(1);
// 设置（指定下标）
copyOnWriteArrayList.set(0, 2);
// 获取（查询）
copyOnWriteArrayList.get(0);
// 删除
copyOnWriteArrayList.remove(0);
// 清空
copyOnWriteArrayList.clear();
// 是否为空
copyOnWriteArrayList.isEmpty();
// 是否包含
copyOnWriteArrayList.contains(1);
// 获取元素个数
copyOnWriteArrayList.size();
```

是不是很简单，那 CopyOnWriteArrayList 应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

CopyOnWriteArrayList 并发容器用于读多写少的并发场景。因为采用了写时复制的实现原理，当存在大量写的时候，内存中会频繁复制原有数据的副本，**如果原有数据集很大，则很容易造成内存飙升甚至内存异常**。在日常研发中，可用于静态数据字典的缓存场合，如黑白名单过滤判定。

**注意，CopyOnWriteArrayList 不能保证写入的数据实时读取到，只保证数据的最终一致。是因为写入时需要复制一份原有内容，以及写入后的新老内容互换都需要一定时间。**

我们举一个 IP 黑名单判定的例子：当应用接入外部请求后，为了防范风险，一般会对请求做一些特征判定，如对请求 IP 是否合法的判定就是一种。IP 黑名单偶尔会被系统运维人员做更新。我们使用 CopyOnWriteArrayList 工具类实现此场景，请看下面代码。

## 5. 场景案例

```java
import java.util.Random;
import java.util.concurrent.CopyOnWriteArrayList;

public class CopyOnWriteArrayListTest {

    // 创建一个 CountDownLatch 对象，代表黑名单列表
    private static CopyOnWriteArrayList<String> copyOnWriteArrayList = new CopyOnWriteArrayList<>();
    // 模拟初始化的黑名单数据
    static {
        copyOnWriteArrayList.add("ipAddr0");
        copyOnWriteArrayList.add("ipAddr1");
        copyOnWriteArrayList.add("ipAddr2");
    }

    // 主线程
    public static void main(String[] args) throws InterruptedException {
        Runnable task = new Runnable() {
            public void run() {
                // 模拟接入用时
                try {
                    Thread.sleep(new Random().nextInt(5000));
                } catch (Exception e) {}

                String currentIP = "ipAddr" + new Random().nextInt(5);
                if (copyOnWriteArrayList.contains(currentIP)) {
                    System.out.println(Thread.currentThread().getName() + " IP " + currentIP + "命中黑名单，拒绝接入处理");
                    return;
                }
                System.out.println(Thread.currentThread().getName() + " IP " + currentIP + "接入处理...");
            }
        };
        new Thread(task, "请求1").start();
        new Thread(task, "请求2").start();
        new Thread(task, "请求3").start();

        Runnable updateTask = new Runnable() {
            public void run() {
                // 模拟用时
                try {
                    Thread.sleep(new Random().nextInt(2000));
                } catch (Exception e) {}

                String newBlackIP = "ipAddr3";
                copyOnWriteArrayList.add(newBlackIP);
                System.out.println(Thread.currentThread().getName() + " 添加了新的非法IP " + newBlackIP);
            }
        };
        new Thread(updateTask, "IP黑名单更新").start();

        Thread.sleep(1000000);
    }
}
```

运行上面代码，我们观察一下运行结果。

```java
请求2 IP ipAddr1命中黑名单，拒绝接入处理
IP黑名单更新 添加了新的非法IP ipAddr3
请求3 IP ipAddr3命中黑名单，拒绝接入处理
请求1 IP ipAddr4接入处理...
```

观察结果，和我们的预期一致。

## 6. 小结

本节通过一个简单的例子，介绍了 CopyOnWriteArrayList 的使用场景和基本用法。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
