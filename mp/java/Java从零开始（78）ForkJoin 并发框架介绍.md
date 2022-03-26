---
title: Java从零开始（78）ForkJoin 并发框架介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# ForkJoin 并发框架介绍


## 1. 前言

本节带领大家认识第二个并发框架之 ForkJoin。

本节先介绍 ForkJoin 并发框架的整个体系结构，接着介绍各部分中的核心接口和实现类，下一节中使用 ForkJoin 并发框架实现一个综合例子，让大家从整体概念、接口实现、应用有一个较全面的了解。

下面我们正式开始介绍吧。

## 2. 整体结构介绍

从 JDK 1.7 开始，java 提供了一套大任务分解成小任务并行执行的框架 ForkJoin ，并且在 JDK 1.8 中进一步做了优化。

相比上一节介绍的 Executor 并发框架而言，ForkJoin 框架更倾向于任务拆分并行执行的场合，而 Executor 框架更适合于更一般的任务彼此之间无内在关系的场合。

ForkJoin 框架的基本思想是将一个大任务拆分成多个处理逻辑相同的子任务，最后将这些子任务的结果再汇总起来，从而得到大任务的结果。即在任务处理时，先进行任务切分，然后进行切分后的各子任务的计算，最后做结果合并。担任子任务还可以继续进行切分，这需要根据实际情况而定。

整体结构已经了解了，接着我们继续了解各部分的核心接口和实现类。

## 3. 核心接口和实现类

整个 ForkJoin 框架的核心接口和实现类很简洁，罗列如下：

1. 线程池 ForkJoinPool，代表执行任务的线程池。
2. 执行线程 ForkJoinWorkerThread，代表 ForkJoinPool 线程池中的一个执行任务的线程。
3. 任务 ForkJoinTask ，代表运行在 ForkJoinPool 中的任务。

ForkJoin 框架核心接口的使用逻辑如下图：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz2dn6vj60kv0dan1f02)

下面，我们继续深入了解各接口和实现类的基本知识。

### 3.1. ForkJoinPool

ForkJoinPool 代表了此任务的执行器，当采用 ForkJoin 框架执行我们的任务时，首先需要创建一个 ForkJoinPool 对象，所有后继执行的过程控制都交给此对象完成。代码举例：

```java
// 构建任务执行器
ForkJoinPool pool = new ForkJoinPool();
// 提交待执行的任务
ForkJoinTask<T> result = pool.submit(ForkJoinTask类型的对象);
// 开始任务执行并获取执行结果
result.invoke();
```

### 3.2. ForkJoinWorkerThread

ForkJoinWorkerThread 是 ForkJoin 框架中用于执行任务的线程实现。一般情况下我们无需显式地使用此类，由 ForkJoinPool 类内部自行创建并维护。

### 3.3. ForkJoinTask

ForkJoinTask 是一个抽象类，定义了任务的主要操作接口。

1. fork ()：在当前线程运行的线程池中再创建提交一个子任务。
2. join ()：当任务完成的时候返回计算结果。
3. invoke ()：开始执行任务，如果必要等待计算完成。

共有两个子类：

1. RecursiveAction：提供了无需关心任务执行结果场合下的默认实现
2.  RecursiveTask：提供了最终需获取任务执行结果的场合下的默认实现

一般我们只需要根据实际情况，选择继承上面的两个子类之一，然后实现自己的逻辑就可以了。代码举例：

```java
class myTask extends RecursiveTask<Integer> {
	@Override
	protected Integer compute() {
		// 执行逻辑，可根据情况直接计算返回，也可根据任务大小确定是否继续拆分子任务
		SumTask left = new myTask(子任务待处理的数据范围);
		SumTask right = new SumTask(子任务待处理的数据范围);
		left.fork();
		right.fork();
		return left.join() + right.join();
	}
}
```

## 4. 小结

本节带领大家对 ForkJoin 框架做整体认识，熟悉其最基本骨架内容，先做到有整体的概念。具体 ForkJoin 框架如何使用呢，我们在下一节讲述。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
