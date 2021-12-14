---
title: Java从零开始（103）AQS原理
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# AQS 原理


## 1. 前言

本节内容主要是对 AQS 原理的讲解，之所以需要了解 AQS 原理，是因为后续讲解的 ReentrantLock 是基于 AQS 原理的。本节内容相较于其他小节难度上会大一些，基础薄弱的学习者可以选择性学习本节内容或者跳过本节内容。

* 了解什么是 AQS，这是认识 AQS 原理的前提，是本节的基础知识点；
* 了解 AQS 提供的两种锁功能，对其有一个全局的了解；
* 了解 AQS 的内部框架原理结构，这是本节课程的核心所在，其他所有的知识点讲解都是围绕这一知识点的；
* 释放锁以及添加线程对于 AQS 内部的变化，这是本节课程的重点知识，了解队列的学习者能够更快的掌握这部分知识；
* AQS 与 ReentrantLock 的联系，这是本节课程与 ReentrantLock 之间的过度知识。

## 2. 什么是 AQS

**定义**：AbstarctQueuedSynchronizer 简称 AQS，是一个用于构建锁和同步容器的框架。

事实上 concurrent 包内许多类都是基于 AQS 构建的，例如 ReentrantLock，ReentrantReadWriteLock，FutureTask 等。AQS 解决了在实现同步容器时大量的细节问题。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvtxrlmj60jg0360u002)

AQS 使用一个 FIFO 队列表示排队等待锁的线程，队列头结点称作 “哨兵节点” 或者 “哑结点”，它不与任何线程关联。其他的节点与等待线程关联，每个阶段维护一个等待状态 waitStatus。

## 3. AQS 提供的两种功能

从使用层面来说，AQS 的锁功能分为两种：独占锁和共享锁。

**独占锁**：每次只能有一个线程持有锁，比如前面给大家演示的 ReentrantLock 就是以独占方式实现的互斥锁；

**共享锁**：允许多个线程同时获取锁，并发访问共享资源，比如 ReentrantReadWriteLock。

## 4. AQS 的内部实现

AQS 的实现依赖内部的同步队列，也就是 FIFO 的双向队列，如果当前线程竞争锁失败，那么 AQS 会把当前线程以及等待状态信息构造成一个 Node 加入到同步队列中，同时再阻塞该线程。当获取锁的线程释放锁以后，会从队列中唤醒一个阻塞的节点 （线程）。

如下图所示，一个节点表示一个线程，它保存着线程的引用（thread）、状态（waitStatus）、前驱节点（prev）、后继节点（next），其实就是个双端双向链表，其数据结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvu9okkj60jg0a7di702)

> **Tips**：AQS 队列内部维护的是一个 FIFO 的双向链表，这种结构的特点是每个数据结构都有两个指针，分别指向直接的后继节点和直接前驱节点。所以双向链表可以从任意一个节点开始，很方便的访问前驱和后继。每个 Node 其实是由线程封装，当线程争抢锁失败后会封装成 Node 加入到 ASQ 队列中去。

## 5. 添加线程对于 AQS 队列的变化

当出现锁竞争以及释放锁的时候，AQS 同步队列中的节点会发生变化，首先看一下添加线程的场景。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvut52jj60jg097mzv02)

**这里会涉及到两个变化**：

* **队列操作的变化**：新的线程封装成 Node 节点追加到同步队列中，设置 prev 节点以及修改当前节点的前置节点的 next 节点指向自己；
* **tail 指向变化**：通过同步器将 tail 重新指向新的尾部节点。

## 6. 释放锁移除节点对于 AQS 队列的变化

第一个 head 节点表示获取锁成功的节点，当头结点在释放同步状态时，会唤醒后继节点，如果后继节点获得锁成功，会把自己设置为头结点，节点的变化过程如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvv47izj60jg0aqtb202)

这个过程也是涉及到两个变化：

**head 节点指向**：修改 head 节点指向下一个获得锁的节点；

**新的获得锁的节点**：如图所示，第二个节点被 head 指向了，此时将 prev 的指针指向 null，因为它自己本身就是第一个首节点，所以 pre 指向 null。

## 7. AQS 与 ReentrantLock 的联系

**ReentrantLock 实现**：ReentrantLock 是根据 AQS 实现的独占锁，提供了两个构造方法如下：

```java
 public ReentrantLock() {
        sync = new NonfairSync();
    }
    public ReentrantLock(boolean fair) {
        sync = fair ? new FairSync() : new NonfairSync();
    }
```

**ReentrantLock 有三个内部类**：Sync，NonfairSync，FairSync，继承关系如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnvvixvej60jg09vtaf02)

**总结**：我们可以看到，这三个内部类都是基于 AQS 进行的实现，由此可见，ReentrantLock 是基于 AQS 进行的实现。

ReentrantLock 提供两种类型的锁：公平锁，非公平锁。分别对应 FairSync，NonfairSync。默认实现是 NonFairSync。

## 8. 小结

本节内容为 AQS 原理进行讲解，会涉及到一些原理问题，队列问题，基础薄弱的学习者可以跳过或者选看本节内容，不会影响后续课程的学习。本节内容其实主要为了提供原理性的知识，对本节的知识掌握，使我们不仅仅是一个使用者。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
