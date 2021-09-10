# JUC 包介绍

## 1. 前言

java.util.concurrent （缩写 JUC）并发编程包是专门为 Java 并发编程设计的，在正式介绍 Java 并发工具之前，本节先带领大家认识 Java 并发工具包，对 Java 并发工具在组织形式上有一个直观的认识。

本节先介绍 JUC 包的版本历史，接着介绍 JUC 包的组织形式和内容结构。

在了解了 JUC 包之后，心中有了整体概念，当我们应用到并发工具时，就可以很快定位选择最恰当的工具加以应用，或者能够很快定位查阅相关工具的源代码。

下面我们正式开始介绍。

## 2. JUC 包版本变迁

从 JDK1.5 开始，Java 官方在 rt.jar 核心 jar 包文件中增加了 java.util.concurrent 并发包，由 Doug Lea 大牛编写实现，并在后继的主要版本中不断对其增强、优化。

在 JDK1.6 中，主要对基础数据结构类进行了并发特性增强。

在 JDK1.7 中，主要对并发框架工具类进行了增强，新增了 ForkJoin 系列。

在 JDK1.8 中，主要对原子操作工具类进行了增强，增加了适用于更多场景的工具类。

## 3. JUC 包组织结构

JUC 包在形式上是如何组织的呢？看下面表格。

|包路径|主要内容|典型类型|
|------|--------|--------|
| java.util.concurrent      | 提供很多种最基本的并发工具类，包括对各类数据结构的并发封装，并发框架主要接口| CountDownLatch，CyclicBarrier，Semaphore，Exchanger，Phaser，BlockingQueue，ConcurrentHashMap，ThreadPoolExecutor，ForkJoinPool|
|java.util.concurrent.atomic| 提供各类原子操作工具类                                                      | AtomicInteger， DoubleAdder，LongAccumulator，AtomicReference                                                                  |
|java.util.concurrent.locks | 提供各类锁工具                                                              | Lock，ReadWriteLock，ReentrantLock，StampedLock                                                                                |

了解了形式上的组织方式，我们接下来从内容角度看看 JUC 包的组织方式。

## 4. JUC 包内容结构

JUC 包提供了下面五大方面的内容：

**锁（locks）部分**：提供适合各类场合的锁工具；

**原子变量（atomic）部分**：原子变量类相关，是构建非阻塞算法的基础；

**并发框架（executor）部分**：提供线程池相关类型；

**并发容器（collections） 部分**：提供一系列并发容器相关类型；

**同步工具（tools）部分**：提供相对独立，且场景丰富的各类同步工具，如信号量、闭锁、栅栏等功能；

下面我们通过思维导图直观展示 JUC 包内容的内在联系。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnywm24hj60u00h8qbi02)

## 5. 小结

通过本节介绍，大家对 java 并发工具包有了全局的直观认识，本课程后继主要针对上述表格中的典型类型展开介绍。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
