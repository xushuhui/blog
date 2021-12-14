---
title: Java 从零开始（66）本章小结
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 本章小结

## 1. 前言

通过前面几个小节的学习，相信大家已经掌握了不少知识和技巧，本节我们一起做一个回顾总结。

## 2. 内容回顾

本章总共 4 个小节，介绍了 4 个典型的原子操作类，分别是 AtomicInteger、AtomicReference、DoubleAdder 和 LongAccumulator 。每一个工具类介绍了基本的用法，并且搭配了编程案例。由这 4 个典型的工具类，我们可以类比其他同类型的工具类加以学习，如下：

通过 AtomicInteger 和 AtomicReference 类比学习 AtomicBoolean、AtomicLong、AtomicIntegerArray、AtomicLongArray、AtomicReferenceArray、AtomicReference、AtomicReferenceFieldUpdater、AtomicMarkableReference、AtomicIntegeFieldUpdater、AtomicLongFieldUpdater、AtomicStampedReference。

通过 DoubleAdder 和 LongAccumulator 类比学习 LongAdder、DoubleAccumulator。

## 3. 工具对比

在介绍每一类工具类时，没有过多地做彼此之间的对比，为了大家有一个更深刻的认识，下面总结这些工具类之间的应用差异。请看下面表格。

|工具类|基本概念|典型应用场景|
|------|--------|------------|
| AtomicInteger  | 原子整型工具类，封装了基本类型整型变量的细粒度原子操作          |应用在多线程操作同一个整型变量时                    |
| AtomicReference| 原子引用工具类，封装了引用类型变量的细粒度原子操作              |应用在多线程操作同一个引用变量时                    |
| DoubleAdder    | 浮点型加法器， 封装了基本类型浮点型变量的粗粒度原子操作         |应用在多线程统计汇总某一个数值时                    |
| LongAccumulator| 长整型计算器，封装了基本类型长整型变量的自定义运算规则的原子操作|应用在多线程操作同一个长整型变量且需自定义计算规则时|

## 4. 实践建议

JDK 提供了这么多原子操作工具类，且这些工具类在功能上或多或少有重叠，我们该怎么加以选择使用呢？

首先需要对每类工具类的功能彻底理解，然后把握住**最简原则**加以选择。

另外，当我们学习了这些工具类之后，在实际应用中，**应该首先想到使用这些工具类，而不是自己造轮子**。

至此本章介绍完毕了，希望大家能反复琢磨，反复练习，争取早日掌握，早日体会到这些工具类带来的便捷。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
