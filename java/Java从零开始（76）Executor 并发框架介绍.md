---
title: Java从零开始（76）Executor 并发框架介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Executor 并发框架介绍


## 1. 前言

从本节开始，我们学习新一章内容 —— 并发框架。

本章总共介绍两类并发框架，本节带领大家认识第一个并发框架之 Executor。

本节先介绍 Executor 并发框架的整个体系结构，接着介绍各部分中的核心接口和实现类，下一节中使用 Executor 并发框架实现一个综合例子，让大家从整体概念、接口实现、应用有一个较全面的了解。

下面我们正式开始介绍吧。

## 2. 整体结构介绍

从 JDK 1.5 开始，java 中将工作单元和执行机制做了分离，于是 Executor 并行框架出现。

什么是工作单元（或称为任务）呢？其实就是我们需要运行的一段逻辑代码。不管什么逻辑的工作单元，最终都需要通过线程运行。

Executor 并行框架对工作单元、以及工作单元的执行做了高度抽象，形成了一整套完整模型。这个模型包括 3 大部分：

1. 对工作单元的抽象，即任务。
2. 任务的执行机制，即如何组织任务的提交、如何管理提交的任务、如何组织多个线程执行。
3. 对任务执行结果的抽象，即如何跟踪任务执行状态，如何获取任务执行结果。

整体结构已经了解了，接着我们继续了解各部分的核心接口和实现类。

## 3. 核心接口和实现类

整个 Executor 框架的核心接口和实现类型如下：

1. 工作单元：Runnable，Callable
2.  工作单元执行：Executor，ExecutorService
3.  工作单元执行结果：Future，FutureTask

Executor 框架核心接口的使用逻辑如下图：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz22qqqj60mu0cnwi702)

下面，我们继续深入了解各接口和实现类的基本知识。

### 3.1. Runnable & Callable

当不需要关注任务执行结果时，使用 Runnable 很合适，反之使用 Callable。代码举例：

```java
Runnable task = new Runnable() {
	public void run() {
		// 任何想要执行的逻辑
	}
}

Callable<String> task = new Callable<String>() {
    public String call() throws Exception {
        // 任何想要执行的逻辑
        return "任务执行后任何想返回的信息";
    }
};
```

### 3.2. Executor & ExecutorService

Executor 接口定义了一个用于执行任务的 execute 方法。ExecutorService 是 Executor 的子接口，其职责是对一堆用于执行任务的线程做管理，即定义了线程池的基本操作接口，有很多具体的实现子类，其核心操作有：

1. execute (Runnable)：提交 Runnable 任务。
2. submit (Callable 或 Runnable)：提交 Callable 或 Runnable 任务，并返回代表此任务的 Future 对象。
3. shutdown ()：关闭新的外部任务提交。
4. shutdownNow ()：尝试中断正在执行的所有任务，清空并返回待执行的任务列表。
5. isTerminated ()：测试是否所有任务都执行完毕了。
6. isShutdown ()：测试是否该 ExecutorService 已被关闭。

这些核心操作在下一节示例中会有应用。

### 3.3. Future & FutureTask

Future 接口定义了对任务执行结果的取消、状态查询、结果获取方法。FutureTask 是 Future 的唯一实现类，其职责是提供方便地构建带有返回结果的任务。Future 接口的核心操作有：

1. cancel (boolean)：尝试取消已经提交的任务。
2. isCancelled ()：查询任务是否在完成之前被取消了。
3. isDone ()：查询任务是否已经完成。
4. get ()：获取异步任务的执行结果（如果任务没执行完将等待）。

## 4. 小结

本节带领大家对 Executor 框架做整体认识，熟悉其最基本骨架内容，先做到有整体的概念。具体 Executor 框架如何使用呢，我们在下一节讲述。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
