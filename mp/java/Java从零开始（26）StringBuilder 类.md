---
title: Java 从零开始（26）StringBuilder 类
zhihu-url: https://zhuanlan.zhihu.com/p/408517458
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# StringBuilder

上一节，我们学习了 Java 的 `String` 类，并介绍了其常用方法。本小节我们来介绍字符串的另外一个类：`StringBuilder`，我们将会了解到 `StringBuilder` 与 `String` 的**差异**，`StringBuilder` 的**使用场景**，也会介绍与 `StringBuilder` 类对应的 `StringBuffer` 类，`StringBuilder` 的**使用方法以及其常用方法**是本小节的重点学习内容。

## 1. StringBuilder 概述

### 1.1 什么是 StringBuilder

与 `String` 相似，`StringBuilder` 也是一个与字符串相关的类，Java 官方文档给 `StringBuilder` 的定义是：可变的字符序列。

### 1.2 为什么需要 StringBuilder

在 Java 字符串的学习中，我们知道了**字符串具有不可变性**，当频繁操作字符串时候，会在常量池中产生很多无用的数据（回忆图示）。

而 `StringBuilder` 与 `String` 不同，它具有**可变性**。相较 `String` 类不会产生大量无用数据，性能上会大大提高。

因此对于需要频繁操作字符串的场景，建议使用 `Stringbuilder` 类来代替 `String` 类。

## 2. StringBuffer 概述

### 2.1 定义

了解了 `StringBuilder` 类 ，`StringBuffer` 也是不得不提的一个类，Java 官方文档给出的定义是：线程安全的可变字符序列。

### 2.2 与前者的区别

`StringBuffer` 是 `StringBuilder` 的前身，在早期的 `Java` 版本中应用非常广泛，它是 `StringBuilder` 的线程安全版本（**线程**我们将在后面的小节中介绍），但实现线程安全的代价是**执行效率的下降**。

你可以对比 `StringBuilder` 和 `StringBuffer` 的 [接口文档]()，它们的接口基本上完全一致。为了提升我们代码的执行效率，在如今的实际开发中 `StringBuffer` 并不常用。因此本小节的重点在 `StringBuilder` 的学习。

## 3. StringBuilder 的常用方法

### 3.1 构造方法

`StringBuilder` 类提供了如下 4 个构造方法：

1. `StringBuilder()` 构造一个空字符串生成器，初始容量为 16 个字符；
2. `StringBuilder(int catpacity)` 构造一个空字符串生成器，初始容量由参数 `capacity` 指定；
3. `StringBuilder(CharSequence seq)` 构造一个字符串生成器，该生成器包含与指定的 `CharSequence` 相同的字符。；
4. `StringBuilder(String str)` 构造初始化为指定字符串内容的字符串生成器。

其中第 4 个构造方法最为常用，我们可以使用 `StringBuilder` 这样初始化一个内容为 `hello` 的字符串：

```java
StringBuilder str = new StringBuilder("Hello");
```

### 3.2 成员方法

`StringBuilder` 类下面也提供了很多与 `String` 类相似的成员方法，以方便我们对字符串进行操作。下面我们将举例介绍一些常用的成员方法。

#### 3.2.1 字符串连接

可以使用 `StringBuilder` 的 `StringBuilder append(String str)` 方法来实现字符串的连接操作。

我们知道，`String` 的连接操作是通过 `+` 操作符完成连接的：

```java
String str1 = "Hello";
String str2 = "World";
String str3 = str1 + " " + str2;
```

如下是通过 `StringBuilder` 实现的字符串连接示例：

```java
public class ConnectString1 {
    public static void main(String[] args) {
        // 初始化一个内容为 Hello 的字符串生成器
        StringBuilder str = new StringBuilder("Hello");
        // 调用 append() 方法进行字符串的连接
        str.append(" ");
        str.append("World");
       	System.out.println(str);
    }
}
```

运行结果：

```java
Hello World
```

由于 `append()` 方法返回的是一个 `StringBuilder` 类型，我们可以实现**链式调用**。例如，上述连续两个 `append()` 方法的调用语句，可以简化为一行语句：

```java
str.append(" ").append("World");
```

如果你使用 `IDE` 编写如上连接字符串的代码，可能会有下面这样的提示（**IntelliJ idea** 的代码截图）：

![](https://xushuhui.gitee.io/image/imooc/5ea6abde09b6bdf916300451.jpg)

提示内容说可以将 `StringBuilder` 类型可以替换为 `String` 类型，也就是说可以将上边地代码改为：

```java
String str = "Hello" + " " + "World";
```

这样写并不会导致执行效率的下降，这是因为 Java 编译器在编译和运行期间会自动将字符串连接操作转换为 `StringBuilder` 操作或者数组复制，间接地优化了由于 `String` 的不可变性引发的性能问题。

值得注意的是，`append()` 的重载方法有很多，可以实现各种类型的连接操作。例如我们可以连接 `char` 类型以及 `float` 类型，实例如下：

```java
public class ConnectString2 {
    public static void main(String[] args) {
        StringBuilder str = new StringBuilder("小明的身高为");
        str.append('：').append(172.5f);
        System.out.println(str);
    }
}
```

运行结果：

```java
小明的身高为：172.5
```

上面代码里连续的两个 `append()` 方法分别调用的是重载方法 `StringBuilder append(char c)` 和 `StringBuilder append(float f)`。

#### 3.2.2 获取容量

可以使用 `int capacity()` 方法来获取当前容量，容量指定是可以存储的字符数（包含已写入字符），超过此数将进行自动分配。注意，容量与长度（length）不同，长度指的是已经写入字符的长度。

例如，构造方法 `StringBuilder()` 构造一个空字符串生成器，初始容量为 16 个字符。我们可以获取并打印它的容量，实例如下：

```java
public class GetCapacity {
    public static void main(String[] args) {
        // 调用 StringBuilder 的无参构造方法，生成一个 str 对象
        StringBuilder str = new StringBuilder();
        System.out.println("str 的初始容量为：" + str.capacity());
        // 循环执行连接操作
        for (int i = 0; i < 16; i ++) {
            str.append(i);
        }
        System.out.println("连接操作后，str 的容量为" + str.capacity());
    }
}
```

运行结果：

```java
str 的初始容量为：16
连接操作后，str 的容量为 34
```

#### 3.2.3 字符串替换

可以使用 `StringBuilder replace(int start, int end, String str)` 方法，来用指定字符串替换从索引位置 `start` 开始到 `end` 索引位置结束（不包含 `end`）的子串。实例如下：

```java
public class StringReplace {
    public static void main(String[] args) {
        // 初始化一个内容为 Hello 的字符串生成器
        StringBuilder str = new StringBuilder("Hello World!");
        // 调用字符串替换方法，将 World 替换为 Java
        str.replace(6, 11, "Java");
        // 打印替换后的字符串
        System.out.println(str);
    }
}
```

运行结果：

```java
Hello Java!
```

也可使用 `StringBuilder delete(int start, int end)` 方法，先来删除索引位置 `start` 开始到 `end` 索引位置（不包含 `end`）的子串，再使用 `StringBuilder insert(int offset, String str)` 方法，将字符串插入到序列的 `offset` 索引位置。同样可以实现字符串的替换，例如：

```java
StringBuilder str = new StringBuilder("Hello World!");
str.delete(6, 11);
str.insert(6, "Java");
```

#### 3.2.4 字符串截取

可以使用 `StringBuilder substring(int start)` 方法来进行字符串截取，例如，我们想截取字符串的后三个字符，实例如下：

```java
public class StringSub {
    public static void main(String[] args) {
        StringBuilder str = new StringBuilder("你好，欢迎来到");
        String substring = str.substring(7);
        System.out.println("str 截取后子串为：" + substring);
    }
}
```

运行结果：

```java
str 截取后子串为：
```

如果我们想截取示例中的” 欢迎 “二字，可以使用重载方法 `StringBuilder substring(int start, int end)` 进行截取：

```java
String substring = str.substring(3, 5);
```

#### 3.2.5 字符串反转

可以使用 `StringBuildr reverse()` 方法，对字符串进行反转操作，例如：

```java
public class StringReverse {
    public static void main(String[] args) {
        StringBuilder str = new StringBuilder("Hello Java");
        System.out.println("str 经过反转操作后为：" + str.reverse());
    }
}
```

运行结果：

```java
str 经过反转操作后为：avaJ olleH
```

## 4. 小结

本小节我们介绍了 Java 的 `StringBuilder` 类，它具有可变性，对于频繁操作字符串的场景，使用它来代替 `String` 类可以提高程序的执行效率；也知道了 `StringBuffer` 是 `StringBuilder` 的线程安全版本，官方更推荐使用 `StringBuilder`；最后我们介绍了 `StringBuilder` 的常用构造方法和成员方法，如果你想了解更多关于 `StringBuilder` 的接口，可以翻阅 [官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/lang/StringBuilder.html) 进行学习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
