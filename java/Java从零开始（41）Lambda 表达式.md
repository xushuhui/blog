---
title: Java 从零开始（41）Lambda 表达式
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
---

# Lambda 表达式

Lambda 表达式是一个 Java 8 以后开始支持的一个非常优雅的新特性，本小节我们将学习什么是 Lambda 表达式，为什么需要 Lambda 表达式，Lambda 表达式的基础语法，以及 Lambda 表达式的实际应用等内容。

## 1. 什么是 Lambda 表达式

> Lambda 表达式基于数学中的 [ λ 演算](https://baike.baidu.com/item/%CE%BB%E6%BC%94%E7%AE%97) 得名，直接对应于其中的 lambda 抽象（lambda abstraction），是一个匿名函数，即没有函数名的函数。

Lambda 表达式是一个匿名函数，匿名函数由 `LISP` 语言在 1958 年首先采用，此后，越来越多的编程语言开始陆续采用。

我们可以把 Lambda 表达式理解为是**一段可传递的代码**（将代码像数据一样传递）。使用它可以写出简洁、灵活的代码。作为一种更紧凑的代码风格，使 Java 语言的表达能力得到了提升。

## 2. 为什么需要 Lambda 表达式

在 Java 8 之前，编写一个匿名内部类的代码很冗长、可读性很差，请查看如下实例：

```java
public class LambdaDemo1 {

    public static void main(String[] args) {

        // 实例化一个 Runnable 接口的匿名实现类对象
        Runnable runnable = new Runnable() {
            @Override
            public void run() {
                System.out.println("Hello, 匿名内部类");
            }
        };
        // 执行匿名内部类的 run() 方法
        runnable.run();
    }

}
```

运行结果：

```java
Hello, 匿名内部类
```

Lambda 表达式的应用则使代码变得更加紧凑，可读性增强；另外，Lambda 表达式使并行操作大集合变得很方便，可以充分发挥多核 CPU 的优势，更易于为多核处理器编写代码。

下面我们使用 `Lambda` 表达式改写上面的代码。

如果你使用 `IDEA` 编写代码，可以直接一键智能修改，首先，将鼠标光标移动到灰色的 `new Runnable()` 代码处，此时会弹出一个提示框，提示可以使用 `Lambda` 表达式替换，点击 `Replace with lambda` 按钮即可完成代码替换，截图如下：

![](https://xushuhui.gitee.io/image/imooc/5efe83dd09f1d7a317420892.jpg)

修改后实例如下：

```java
public class LambdaDemo1 {

    public static void main(String[] args) {

        // 实例化一个 Runnable 接口的匿名实现类对象
        Runnable runnable = () -> System.out.println("Hello, 匿名内部类");
        // 执行匿名内部类的 run() 方法
        runnable.run();
    }

}
```

运行结果：

```java
Hello, 匿名内部类
```

通过对比，使用 `lambda` 表达式实现了与匿名内部类同样的功能，并且仅仅用了一行代码，代码变得更加简洁了。对于这样的写法，你可能还非常疑惑，但别担心，我们马上就来详细讲解基础语法。

## 3. 基础语法

Lambda 表达式由三个部分组成：

* **一个括号内用逗号分隔的形式参数列表**：实际上就是接口里面抽象方法的参数；
* **一个箭头符号**：->，这个箭头我们又称为 Lambda 操作符或箭头操作符；
* **方法体**：可以是表达式和代码块，是重写的方法的方法体。语法如下：

1. 方法体为表达式，该表达式的值作为返回值返回。

```java
(parameters) -> expression
```

1. 方法体为代码块，必须用 {} 来包裹起来，且需要一个 return 返回值，但若函数式接口里面方法返回值是 void，则无需返回值。

```java
(parameters) -> {
    statement1;
    statement2;
}
```

## 4. 使用实例

通过上面一系列内容的学习，我们可以得出一个结论：Lambda 表达式本质上就是接口实现类的对象，它简化了之前匿名内部类的冗长代码的编写。

关于 Lambda 表达式的具体使用，我们将根据语法格式分为 5 种展开介绍。

### 4.1 无参数无返回值

无参数无返回值，指的是接口实现类重写的方法是无参数无返回值的，我们一开始提到的 `Runnable` 接口匿名内部类就属于此类：

```java
public class LambdaDemo2 {

    public static void main(String[] args) {
        // 通过匿名内部类实例实例化一个 Runnable 接口的实现类
        Runnable runnable1 = new Runnable() {
            @Override
            public void run() {  // 方法无形参列表，也无返回值
                System.out.println("Hello, 匿名内部类");
            }
        };
        // 执行匿名内部类的 run() 方法
        runnable1.run();

        // 无参数无返回值，通过 lambda 表达式来实例化 Runnable 接口的实现类
        Runnable runnable2 = () -> System.out.println("Hello, Lambda");
        // 执行通过 lambda 表达式实例化的对象下的 run() 方法
        runnable2.run();
    }

}
```

运行结果：

```java
Hello, 匿名内部类
Hello, Lambda
```

### 4.2 单参数无返回值

无参数无返回值，指的是接口实现类重写的方法是单个参数，返回值为 `void` 的，实例如下：

```java
import java.util.function.Consumer;

public class LambdaDemo3 {

    public static void main(String[] args) {

        // 单参数无返回值
        Consumer<String> consumer1 = new Consumer<String>() {
            @Override
            public void accept(String s) {
                System.out.println(s);
            }
        };
        consumer1.accept("Hello World!");

        Consumer<String> consumer2 = (String s) -> {
            System.out.println(s);
        };
        consumer2.accept("你好，世界！");
    }

}
```

运行结果：

```java
Hello World!
你好，世界！
```

### 4.3 省略数据类型

什么叫省略数据类型呢？我们依旧采用上面的实例，使用 `Lambda` 表达式可以省略掉括号中的类型，实例代码如下：

```java
// 省略类型前代码
Consumer<String> consumer2 = (String s) -> {
    System.out.println(s);
};
consumer2.accept("你好，世界！");

// 省略类型后代码
Consumer<String> consumer3 = (s) -> {
    System.out.println(s);
};
consumer3.accept("你好，世界！");
```

> **Tips**：之所以能够省略括号中的数据类型，是因为我们在 `Comsumer<String>` 处已经指定了泛型，编译器可以推断出类型，后面就不用指定具体类型了。称为**类型推断**。

### 4.4 省略参数的小括号

当我们写的 `Lambda` 表达式只需要一个参数时，参数的小括号就可以省略，改写上面实例的代码：

```java
// 省略小括号前代码
Consumer<String> consumer3 = (s) -> {
    System.out.println(s);
};
consumer3.accept("你好，世界！");
// 省略小括号后代码
Consumer<String> consumer4 = s -> {
    System.out.println(s);
};
consumer3.accept("你好，世界！");
```

观察实例代码，之前的 `(s) ->` 可以改写成 `s ->`，这样写也是合法的。

### 4.5 省略 return 和大括号

当 `Lambda` 表达式体只有一条语句时，如果有返回，则 return 和大括号都可以省略，实例代码如下：

```java
import java.util.Comparator;

public class LambdaDemo4 {

    public static void main(String[] args) {

        // 省略 return 和 {} 前代码
        Comparator<Integer> comparator1 = (o1, o2) -> {
            return o1.compareTo(o2);
        };
        System.out.println(comparator1.compare(12, 12));

        // 省略 return 和 {} 后代码
        Comparator<Integer> comparator2 = (o1, o2) -> o1.compareTo(o2);
        System.out.println(comparator2.compare(12, 23));

    }
}
```

运行结果：

```java
0
-1
```

## 5. 小结

通过本小节的学习，我们知道了 `Lambda` 表达式并不是 Java 所特有的特性，很多主流编程语言都支持 `Lambda` 表达式，在 Java 中，Lambda 表达式实际上就是接口实现类的对象，它简化了之前匿名内部类的冗长代码的编写。Lambda 表达式能使代码变得更加紧凑，增强代码的可读性；另外，Lambda 表达式使并行操作大集合变得很方便，可以充分发挥多核 CPU 的优势，更易于为多核处理器编写代码。

我们也通过大量的实例代码介绍了 `Lambda` 的语法格式和实际使用，通过这些实例，可以发现 `Lambda` 表达式对于接口也是有要求的 —— 接口内部只能包含一个抽象方法，如果接口内部包含多个抽象方法，我们就无法使用 `Lambda` 表达式了，这样只包含一个抽象方法的接口，我们称为**函数式接口**，下一小节我们将详细介绍函数式接口。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
