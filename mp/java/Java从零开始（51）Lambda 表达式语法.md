---
title: Java 从零开始（51）Lambda 表达式语法
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式的语法

本节我们将介绍 Lambda 表达式的语法，了解 Lambda 表达式的基本形式和几种变体，以及编译器是怎么理解我们的 Lambda 表达式的。

掌握了这些基础的知识，我们就能很容易的辨别出 Lambda 表达式，以及它是怎么运作的。

## 1. 基本语法

现在我们来回顾下第一 个 Lambda 表达式的例子：

```java
button.addActionListener(event -> System.out.println("button click"));
```

这是一个最基本的 Lambda 表达式，它由三部分组成具体格式是这样子的：

```java
参数 -> 具体实现
```

包含一个 Lambda 表达式的运算符 `->`，在运算符的左边是输入参数，右边则是函数主体。

概括来讲就是：

> **Tips:** 一段带有输入参数的可执行语句块

Lambda 表达式有这么几个特点：

1. **可选类型声明：** 不需要声明参数类型，编译器可以自动识别参数类型和参数值。在我们第一个例子中，并没有指定 event 到底是什么类型；
2. **可选的参数圆括号：** 一个参数可以不用定义圆括号，但多个参数需要定义圆括号；
3. **可选的大括号：** 如果函数主体只包含一个语句，就不需要使用大括号；
4. **可选的返回关键字：** 如果主体只有一个表达式返回值则编译器会自动返回值，大括号需要指明表达式返回了一个数值。

这几个特点是不是看着有点晕呢？不要紧，你只要知道除了我们第一个例子中的 Lambda 表达式的基本形式之外，这些特点还能够形成它的几种变体。接下来我们来看下一个 Lambda 表达式的一些变体。

## 2. Lambda 表达式的五种形式

### 2.1 不包含参数

```java
Runnable noArguments = () -> System.out.println("hello world");
```

在这个例子中，Runnable 接口，只有一个 run 方法，没有参数，且返回类型为 void，所以我们的 Lambda 表达式使用 `()` 表示没有输入参数。

### 2.2 有且只有一个参数

```java
ActionListener oneArguments = event -> System.out.println("hello world");
```

在只有一个参数的情况下 我们可以把 `()` 省略。

### 2.3 有多个参数

```java
BinaryOperator<Long> add = (x,y) -> x+y ;
```

使用 `()` 把参数包裹起来，并用 `,` 来分割参数。上面的代码表示。

### 2.4 表达式主体是一个代码块

```java
Runnable noArguments = () -> {
    System.out.print("hello");
    System.out.println("world");
}
```

当有多行代码的时候我们需要使用 `{}` 把表达式主体给包裹起来。

### 2.5 显示声明参数类型

```java
BinaryOperator<Long> add = (Long x, Long y) -> x+y ;
```

通常 Lambda 的参数类型都有编译器推断得出的，也可以显示的声明参数类型。

## 3. 关于 Lambda 表达式的参数类型

我们再来看一下 2.3 的例子：

```java
BinaryOperator<Long> add = (x,y) -> x+y ;
```

在这个例子中，参数 `x`，`y` 和返回值 `x+y` 我们都没有指定具体的类型，但是编译器却知道它是什么类型。原因就在于编译器可以从程序的上下文推断出来，这里的上下文包含下面 3 种情况：

* 赋值上下文；
* 方法调用上下文；
* 类型转换上下文。

通过这 3 种上下文就可以推断出 Lambda 表达式的目标类型。

目标类型并不是一个全新的概念，通常我们在 Java 数据初始化的时候就是根据上下文推断出来的。比如：

```java
String[] array = {"hello","world"};
```

等号右边的代码我们并没有声明它是什么类型，系统会根据上下文推断出类型的信息。

## 4. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a8c99099b657309380350.jpg)

本节主要介绍了：

* Lambda 表达式的基本语法及其几种变体形式。
* 在 Lambda 表达式中，编译器会根据程序的上下文自动推断出目表达式的目标类型。

掌握这些知识可以帮助我们快速的辨别一个 Lambda 表达式，方便的去理解程序的意图。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
