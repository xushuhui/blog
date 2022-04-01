---
title: Java 从零开始（40）Java 8 新特性
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415757603
---
# Java 8 新特性

同学们可能有个疑问，Java 的版本已经更新到了 15，本教程使用的版本也是最新的 Java 15，为什么还要介绍`Java 8`的新特性呢？为什么不去介绍 `Java 15` 的新特性呢？

这是因为 Java 8 是一个有重大改变的版本，该版本对 Java 做了重大改进。Java 8 由 Oracle 公司在 2014 年 3 月发布，可以看成是自 Java 5 以后的最具革命性的版本。至今仍是企业级应用最优先考虑使用的版本。Java 8 以后大的版本更新速度比较快，同学们没有必要每个版本的新特性都去研习，有些“新特性”只是尝试，不一定好用，也可能在将来的版本被废弃掉，因此我们学习最稳定的 Java 8 新特性即可。（如果你想了解 Java 15 的新特性，建议去 [官方文档](https://www.oracle.com/java/technologies/javase/15all-relnotes.html) 翻阅。）

本小节我们将列举并概述 Java 8 的核心新特性，有一部分特性我们已经在前面的小节中介绍过，还有一部分没有介绍过的新特性，由于内容较多切较为重要，本小节将简要介绍，详细内容在之后的几个小节中展开介绍。

## 1. 函数式接口

Java 8 引入的一个核心概念是函数式接口（Functional Interfaces）。通过在接口里面添加一个抽象方法，这些方法可以直接从接口中运行。

有关函数式接口的内容我们将在学完`Lambda`后详细介绍。

## 2. Lambda 表达式

在 Java 8 之前，编写一个匿名内部类的代码很冗长、可读性很差，Lambda 表达式的应用则使代码变得更加紧凑，可读性增强；Lambda 表达式使并行操作大集合变得很方便，可以充分发挥多核 CPU 的优势，更易于为多核处理器编写代码。

关于 Lambda 表达式我们将在下一小节介绍。

## 3. Stream API

Java 8 引入了流式操作（Stream），通过该操作可以实现对集合（Collection）的并行处理和函数式操作。

根据流的并发性，流又可以分为串行和并行两种。流式操作实现了集合的过滤、排序、映射等功能。

根据操作返回的结果不同，流式操作分为中间操作和最终操作两种。最终操作返回一特定类型的结果，而中间操作返回流本身，这样就可以将多个操作依次串联起来。

关于`Stream API`，我们将在后面的小节中介绍。

## 4. 接口的增强

Java 8 对接口做了进一步的增强。

在接口中可以添加使用 default 关键字修饰的非抽象方法。还可以在接口中定义静态方法。增强后的接口看上去与抽象类的功能越来越类似了。

关于**默认方法**和**静态方法**，我们在 [`Java`接口】(http://www.imooc.com/wiki/javalesson/javainterface.html) 这一小节，已经结合实例详细介绍过了，忘记了的同学可以回去温习一下。

## 5. 注解的更新

对于注解，Java 8 主要有两点改进：类型注解和重复注解。

Java 8 的类型注解扩展了注解使用的范围。在该版本之前，注解只能是在声明的地方使用。现在几乎可以为任何东西添加注解：局部变量、类与接口，就连方法的异常也能添加注解。新增的两个注释的程序元素类型 `ElementType.TYPE_USE` 和 `ElementType.TYPE_PARAMETER` 用来描述注解的新场合。对类型注解的支持，增强了通过静态分析工具发现错误的能力。原先只能在运行时发现的问题可以提前在编译的时候被排查出来。

在该版本之前使用注解的一个限制是相同的注解在同一位置只能声明一次，不能声明多次。Java 8 引入了重复注解机制，这样相同的注解可以在同一地方声明多次。重复注解机制本身必须用 @Repeatable 注解。

关于注解的更多知识点，可以回到 [`Java` 注解】(http://www.imooc.com/wiki/javalesson/javaannotation.html) 这一小节温习。

## 6. IO/NIO 的改进

Java 8 对 `IO/NIO` 也做了一些改进。主要包括：改进了 `java.nio.charset.Charset` 的实现，使编码和解码的效率得以提升，也精简了 `jre/lib/charsets.jar` 包；优化了 String(byte[],*) 构造方法和 String.getBytes() 方法的性能；还增加了一些新的 IO/NIO 方法，使用这些方法可以从文件或者输入流中获取流（java.util.stream.Stream），通过对流的操作，可以简化文本行处理、目录遍历和文件查找。

## 7. 新的日期时间 API

Java 的日期与时间 API 问题由来已久，Java 8 之前的版本中关于时间、日期及其他时间日期格式化类由于线程安全、重量级、序列化成本高等问题而饱受批评。Java 8 吸收了 Joda-Time 的精华，提供了更优秀易用的 API。

关于新的日期时间 API，我们已经在 [`Java`日期时间处理】(http://www.imooc.com/wiki/javalesson/datetime.html) 这一小节具体介绍。

## 8. 小结

通过本小节的学习，我们了解到 Java 8 是一个革命性的版本，新增了诸多好用的新特性，也深得企业级开发的青睐。关于 Java 8 之后的版本的新特性，我们不再深究。同学们可以去 [官网](https://www.oracle.com/java/technologies/javase/8-whats-new.html) 来查看详细的版本更新日志。

关于**函数式接口**、**`Lambda` 表达式**、`Stream API`等新特性我们将在接下来的小节中展开讲解。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
