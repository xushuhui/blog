# Java 并发工具之 Semaphore

## 1. 前言

从本节开始，我们学习新一章内容 —— 同步工具。

本节带领大家认识第一个常用的 Java 并发工具类之 Semaphore。

本节先通过一个生活中的例子为大家通俗解释一下什么是 Semaphore 信号量，接着介绍 Semaphore 工具类的最基本用法，有了这些基本认识之后，给出 Semaphore 工具最常用的场合说明，然后通过简单的编码实现文中提到的生活案例，让大家有一个理性的认识，之后带领大家熟悉 Semaphore 最常用的一些编程方法，最后通过同类工具的比较，进一步加深对 Semaphore 工具类的理解。

Semaphore 工具类本身使用很简单，重点是对常用编程方法的准确理解。

当我们遇到各类需要做并发控制的场合时，怎么做到选取最合适的并发工具加以应用呢？唯有多加练习，不断总结各种并发工具之间的区别，透彻理解各类工具的应用场合，才能做到游刃有余，手到擒来。

下面我们正式开始介绍吧。

## 2. 概念解释

从 JDK1.5 开始提供，Java 官方就在 java.util.concurrent 并发包中提供了 Semaphore 工具类。

那什么是 “Semaphore” 呢？单词 “Semaphore” 在计算机世界中被解释为中文 “信号量” ，但更能表述其含义的叫法应该是 “许可证管理器”。不管叫什么中文名称，它就是一种计数信号量，用于管理一组资源，给资源的使用者规定一个量从而控制同一时刻的使用者数目。

这样的解释是不是很抽象？没关系，在此为大家举一个生活中通俗的例子，让大家先对 “信号量” 及其应用有一个感性的认识。

大家先观察一下下面过闸机的图例，回想一下我们平时过闸机的场景。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyy5iqnj60hw0efwf702)

（图片来自网络，图片版权归作者所有）

比如上图中过闸机就是信号量的基本运用。

上图中的乘客就类比是我们程序里面的各类线程，闸机就类比是一类线程需要使用的资源，而信号量就是某一时刻可用的闸机数量。

当某个时刻有乘客需要使用闸机过站时，首先他需要找到一台没有人使用的闸机，现实中他通过眼睛观察即可知道，在我们程序里面就是需要观察信号量，看能不能申请到代表可用闸机的信号量，如果能则表示有空闲闸机 （资源） 可用，否则需要等待其他乘客使用完毕 （信号量释放）后再使用。

概念我们已经了解了，那 Semaphore 工具类最基本的用法是怎样的呢？别急，看下面。

## 3. 基本用法

```java
// 首先创建 Semaphore 对象
Semaphore semaphore = new Semaphore();
// 在资源操作开始之前，先获取资源的使用许可
semaphore.acquire();
...
// 在获取到资源后，利用资源进行业务处理
...
// 在资源操作完毕之后，释放资源的使用许可
semaphore.release();
...
```

是不是很简单，那 Semaphore 信号量在我们日常实践中，到底应该应用在哪些场合比较合适呢？下面我们给出最常用的场景说明。

## 4. 常用场景

Semaphore 经常用于限制同一时刻获取某种资源的线程数量，最为典型的就是做流量控制。

比如 WEB 服务器处理能力有限，需要控制网络请求接入的最大连接数，以防止过大的请求流量压垮我们的服务器，导致整个应用不能正常提供服务。

比如数据库服务器处理能力有限，需要控制数据库最大连接数，以防止大量某个应用过分占有数据库连接数，导致数据库服务器不能为其他的应用提供足够的连接请求。

当在研发过程中遇到类似这些场景时，就可以考虑直接应用 Semaphore 工具类辅助实现。

上面举的生活中过闸机的例子，如果用程序表达，该如何实现呢？在程序中如何使用 Semaphore 信号量达到控制和应用呢？最直接方式就是去感受最简单的例子，下面直接用最明了的代码说明例子中如何应用了信号量。

## 5. 场景案例

```java
import java.util.concurrent.Semaphore;

public class SemaphoreTest {

    // 先定义一个Semaphore信号量对象
    private static Semaphore semaphore = new Semaphore(3);

    // 测试方法
    public static void main(String[] args) {

        // 定义10个人过闸机
        for(int i=0; i<10; i++) {
            Person person = new Person(semaphore, i);
            new Thread(person).start();
        }
    }
}
```

在上面的代码中，先创建了一个 Semaphore 信号量对象，然后赋给了每一位进站旅客 Person ，接下来每一位旅客如何动作呢，看下面的代码。

```java
import java.util.concurrent.Semaphore;

public class Person implements Runnable {

    private Semaphore semaphore;
    private String persionName;

    public Person(Semaphore semaphore, int persionNo) {
        this.semaphore = semaphore;
        this.persionName = "旅客" + persionNo;
    }

    public void run() {
        try {
            // 请求获得信号量，就是请求（寻找）是否有可用的闸机
            semaphore.acquire();
            // 已经等到了可用闸机
            System.out.println(this.persionName + "已经占有一台闸机");
            // 进站
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            // 已经进站
            System.out.println(this.persionName + "已经进站");
            // 让出闸机给别人用
            semaphore.release();
        }
    }
}
```

在 Person 类中，首先通过 acquire 获取了可用闸机，然后休眠 2 秒代表刷卡过闸机，最后在 finally 使用 release 方法让出闸机。我们观察一下运行结果。

```java
旅客0已经占有一台闸机  <---占据了一台
旅客1已经占有一台闸机  <---占据了一台
旅客6已经占有一台闸机  <---占据了一台
旅客0已经进站         <---0号旅客已经进站释放了闸机
旅客3已经占有一台闸机  <---3号旅客这个时候才拿到了可用闸机
旅客1已经进站         <---1号旅客已经进站释放了闸机
旅客2已经占有一台闸机  <---2号旅客这个时候才拿到了可用闸机
旅客6已经进站         <---6号旅客已经进站释放了闸机
旅客4已经占有一台闸机  <---4号旅客这个时候才拿到了可用闸机
旅客3已经进站
旅客4已经进站
旅客5已经占有一台闸机
旅客7已经占有一台闸机
旅客2已经进站
旅客8已经占有一台闸机
旅客7已经进站
旅客5已经进站
旅客9已经占有一台闸机
旅客8已经进站
旅客9已经进站
```

观察结果发现，同一时刻最多只能有 3 位旅客占用闸机进站，其他旅客需要等待其进站后让出闸机才能刷卡进站。

至此，大家对信号量已经有了初步的理解，接下来我们继续丰富对 Semaphore 工具类的认识。

## 6. 其他方法介绍

除过上面代码中使用的最基本的 acquire 方法和 release 方法之外，我们还需要掌握其他几个核心方法的使用。下面逐个介绍。

1. Semaphore(int permits, boolean fair)

上面的例子中使用了 Semaphore (int permits) 构造方法。

此构造方法也是用于创建信号量对象，第二个参数表示创建的信号量是否秉持公平竞争特性。即对资源的申请使用严格按照申请的顺序给予允许。

一般情况下，我们使用 Semaphore (int permits) 构造方法就可以了。

1. availablePermits()

返回当前还可用的许可数，即还允许多少个线程进行使用资源。套用在上面的例子中，就是返回当前还有多少台闸机空闲可用。

```java
int availablePermits = semaphore.availablePermits();
System.out.println("当前可用闸机数" + availablePermits);

>>运行结果：
当前可用闸机数2
旅客0已经占有一台闸机
当前可用闸机数1
旅客1已经占有一台闸机
......
```

1. hasQueuedThreads()

返回是否有线程正在等待获取资源。也就是返回当前是否有人在排队等待过闸机。

```java
boolean hasQueuedThreads = semaphore.hasQueuedThreads();
System.out.println("当前是否有旅客等待闸机进站："+hasQueuedThreads);

>>运行结果：
当前是否有旅客等待闸机进站：false
旅客0已经占有一台闸机
当前是否有旅客等待闸机进站：false
旅客1已经占有一台闸机
当前是否有旅客等待闸机进站：false
旅客2已经占有一台闸机
```

1. acquire(int permits)

申请指定数目的信号量许可，在获取不到指定数目的许可时将一直阻塞。就好比一个旅客需要同时占用两个闸机过站。类似的 release (int permits) 方法用于释放指定数目的信号量许可。

acquire (int permits) 同上面例子中使用的 acquire () 最大的区别就是用于一次性申请多个许可，当参数 permits = 1 时，两者相同。release (int permits) 和 release () 也是类似。

1. tryAcquire()

尝试申请信号量许可，无论是否申请成功都返回申请结果。当申请成功时返回 true ， 否则返回 false 。程序里面根据申请结果决定后继的处理流程。和 acquire () 的主要区别在于，不会阻塞立刻返回。

同类功能的方法还有 tryAcquire (int permits) 、tryAcquire (long timeout, TimeUnit unit) 、tryAcquire (int permits, long timeout, TimeUnit unit) 。这些方法实现的功能一样，只是可以更加精细化地控制对资源申请，比如申请超时控制、申请许可数量。

## 7. 工具对比

大家可能有一个疑问了，Semaphore 好像和 synchronized 关键字没什么区别，都可以实现同步。

其实不然，synchronized 真正用于并发控制，确保对某一个资源的串行访问；而 Semaphore 限制访问资源的线程数，其实并没有实现同步，只有当 Semaphore 限制的资源同时只允许一个线程访问时，两者达到的效果一样。

大家记住，Semaphore 和 synchronized 最主要的差别是 Semaphore 可以控制一个或多个并发，而 synchronized 只能是一个。这一点需要大家好好琢磨。

还是通过上面的例子的运行结果给大家做一下解释。

```java
>>运行结果：
旅客0已经占有一台闸机  <-------
旅客1已经占有一台闸机  | 观察发现同时有多个并发执行，而非串行的一个旅客过完闸机后才轮到下一个旅客。
旅客2已经占有一台闸机  <-------
......
```

## 8. 小结

本节通过一个简单的例子，介绍了 Semaphore 的基本用法。另外对一些核心方法做了简单介绍并给出应用场景。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
