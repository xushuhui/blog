---
title: Java 从零开始（50）Lambda 表达式简介
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式简介

大家好，今天我们开始一个新专题 —— Java Lambda 表达式。这是在 Java 8 中出现的一个新特性，但它并不是 Java 独有的，JavaScript、C#、C++ 等在 Java 8 之前就支持 Lambda 表达式的特性，现在的大多数程序语言也都支持 Lambda 表达式。

这个专题中我们学习函数式编程的概念、Lambda 表达式的语法、以及如何在我们的代码中使用 Lambda 表达式。本文我们主要先介绍下 Lambda 表达式是什么？

## 1. 什么是 Lambda 表达式

什么是 Lambda 表达式呢？维基百科是这样定义的：

> Lambda expression in computer programming, also called an anonymous function, is a defined function not bound to an identifier. ——维基百科

翻译过来就是 Lambda 表达式也叫作匿名函数，是一种是未绑定标识符的函数定义，在编程语言中，匿名函数通常被称为 Lambda 抽象。

换句话说， Lambda 表达式通常是在需要一个函数，但是又不想费神去命名一个函数的场合下使用。这种匿名函数，在 JDK 8 之前是通过 Java 的匿名内部类来实现，从 Java 8 开始则引入了 Lambda 表达式——一种紧凑的、传递行为的方式。

## 2. 为什么要引入 Lambda 表达式

Java Lambda 表达式是伴随 Java 8 于 2014 年出现的。这个时候恰好是多核 CPU 和大数据兴起的时候。在这些趋势下，芯片设计者只能采用多核并行的设计思路，而软件开发者必须能够更好地利用底层硬件的并发特性。做过多线程编程的同学应该很清楚，在并发过程中涉及锁相关的编程算法不但容易出错，而且耗时费力。虽然 Java 提供了 `java.util.concurrent` 包以及很多第三方类库来帮助我们写出多核 CPU 上运行良好的程序，但在大数据集合的处理上面这些工具包在高效的并行操作上都有些欠缺，我们很难通过简单的修改就能够在多核 CPU 上进行高效的运行。

为了解决上述的问题，需要在程序语言上修改现有的 Java ——增加 Lambda 表达式，同时在 Java 8 也引入 Stream（java.util.stream） 流——一个来自数据源的元素队列，支持聚合操作，来提供对大数据集合的处理能力。

所以 Lambda 表达式的出现时为了适应多核 CPU 的大趋势，一方面通过它我们方便的高效的并发程序，通过简单地修改就能编写复杂的集合处理算法。

## 3. Lambda 表达式的优点

那么 Lambda 具体有哪些优点呢？

1. **更加紧凑的代码**： Lambda 表达式可以通过省去冗余代码来减少我们的代码量，增加我们代码的可读性；
2. **更好地支持多核处理**： Java 8 中通过 Lambda 表达式可以很方便地并行操作大集合，充分发挥多核 CPU 的潜能，并行处理函数如 filter、map 和 reduce；
3. **改善我们的集合操作**： Java 8 引入 Stream API，可以将大数据处理常用的 map、reduce、filter 这样的基本函数式编程的概念与 Java 集合结合起来。方便我们进行大数据处理。

## 4. 我们的第一个例子

前面说了 Lambda 表达式的优点，我们用一个例子来直观的感受下 Lambda 表达式是如何帮我们减少代码行数，增加可读性的。

Swing 是一个与平台无关的 Java 类库（位于 `java.awt.*` 中），用来编写图形界面（ GUI ）。里面有一个常见的用法：为了响应用户操作，需要注册一个事件监听器，当用户输入时，监听器就会执行一些操作（这类似于我们网页的上的一个 Botton 按钮，当用户点击按钮后，js 代码会执行相应的动作）。

### 4.1 使用使用匿名内部类来将行为和点击按钮进行关联

这是我们在 Java 8 以前，通常的写法：

```java
button.addActionListener(new ActionListener(){
    @Override
    public void actionPerformed(ActionEvent actionEvent) {
        System.out.println("button click");
    }
});
```

在上面的例子中，我们创建了一个对象来实现 ActionListener 接口（这个对象并没有命名，它是一个匿名内部类），这个对象有一个 actionPerformed 方法。当用户点击按钮 button 时，就会调用该方法，输出 `button click`。

### 4.2 使用 Lambda 表达式来将行为和点击按钮进行关联

在使用 Lambda 表达式以后的写法：

```java
button.addActionListener(event -> System.out.println("button click"));
```

我们只用了一行代码就完成了，你是不是一眼就看出来这个点击事件做的就是输出 `button click`。

关于 Lambda 表 达式其他的两个特点我们将在后续的内容中进行解释。

## 5. 前置知识

本教程主要讲解的是 Java 8 新特性中的 Lambda 表达式的语法基础以及应用，所以需要学习本教程的读者至少掌握 Java 基础语法，Java 集合，迭代器，还需要了解部分设计模式以及设计原则为后期部分小节学习做铺垫。

## 6. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a899609c28f5805810286.jpg) 本节主要介绍了：

* 分析了 Java 8 开始支持 Lambda 表达式的动机、以及使用 Lambda 表达式的好处。
* 介绍了 Lambda 表达式的优点，并通过 Swing 的按钮点击事件来直观的感受 Lambda 表达式是如何来简化我们的代码的。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
