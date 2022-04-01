---
title: Java 从零开始（29）Java 包装类
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
zhihu-url: https://zhuanlan.zhihu.com/p/412913098
---

# Java 包装类

本小节我们将学习 Java 的包装类，我们将了解到**什么是包装类**，**为什么需要包装类**，Java 提供的**包装类有哪些**，各种**包装类的常用方法和常量**介绍，**什么是装箱操作**以及**什么是拆箱操作**等内容。

## 1. 什么是包装类

Java 有 8 种基本数据类型，Java 中的每个基本类型都被包装成了一个类，这些类被称为包装类。

包装类可以分为 3 类：`Number`、`Character`、`Boolean`，包装类的架构图如下所示：

![](https://xushuhui.gitee.io/image/imooc/5ed45b8e09923bf007990474.jpg)

## 2. 为什么需要包装类

我们知道 Java 是面向对象的编程语言，但为了便于开发者上手，Java 沿用了 C 语言的基本数据类型，因此 Java 数据类型被分为了基本数据类型和引用数据类型。

对于简单的运算，开发者可以直接使用基本数据类型。但对于需要对象化交互的场景（例如将基本数据类型存入集合中），就需要将基本数据类型封装成 Java 对象，这是因为基本数据类型不具备对象的一些特征，没有对象的属性和方法，也不能使用面向对象的编程思想来组织代码。出于这个原因，包装类就产生了。

包装类就是一个类，因此它有属性、方法，可以对象化交互。

## 3. 基本数据类型与包装类

下表列出了基本数据类型对应的包装类。这些包装类都位于 `java.lang` 包下，因此使用包装类时，我们不需要手动引入。

|基本数据类型|对应的包装类|
|------------|------------|
| byte  |Byte     |
|short  |Short    |
|int    |Integer  |
|long   |Long     |
|float  |Float    |
|double |Double   |
|char   |Character|
|boolean|Boolean  |

除了 `int` 对应的包装类名称为 `Integer` 以及 `char` 对应的包装类名称 `Character`，其他 `6` 种数据类型对应的包装类，命名都为其基本数据类型的首字母的大写。

## 4. 包装类常用方法

### 4.1 Number 类

Number 类是所有数值类型包装类的父类，这里以其中一个子类 `Integer` 类为例，介绍其构造方法、常用方法以及常量。

#### 4.1.1 构造方法

Integer 类提供两个构造方法：

1. `Integer(int value)`：以 int 型变量作为参数创建 Integer 对象；
2. `Integer(String s)`：以 String 型变量作为参数创建 Integer 对象。

实例如下：

```java
// 以 int 型变量作为参数创建 Integer 对象
Integer num = new Integer(3);
// 以 String 型变量作为参数创建 Integer 对象
Integer num = new Integer("8");
```

#### 4.1.2 常用方法

* `byte byteValue()`：以 byte 类型返回该 Integer 的值；
* `int compareTo(Integer anotherInteger)`：在数值上比较两个 Integer 对象。如果这两个值相等，则返回 0；如果调用对象的数值小于 anotherInteger 的数值，则返回负值；如果调用对象的数值大于 anotherInteger 的数值，则返回正值；
* `boolean equals(Object obj)`：比较此对象与指定对象是否相等；
* `int intValue()`：以 int 类型返回此 Integer 对象；
* `int shortValue()`：以 short 类型返回此 Integer 对象；
* `toString()`：返回一个表示该 Integer 值的 String 对象；
* `static Integer valueOf(String str)`：返回保存指定的 String 值的 Integer 对 象；
* `int parseInt(String str)`：返回包含在由 str 指定的字符串中的数字的等价整数值。

更多常用方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/lang/Integer.html)。

#### 4.1.3 常用常量

1. `MAX_VALUE`: 表示 int 型可取的最大值；
2. `MIN_VALUE`: 表示 int 型可取的最小值；
3. `SIZE`：表示以二进制补码形式表示 int 值的位数；
4. `TYPE`: 表示基本类型 Class 实例。

这些常量的使用实例如下：

```java
public class WrapperClassDemo1 {

    public static void main(String[] args) {
        int maxValue = Integer.MAX_VALUE;
        int minValue = Integer.MIN_VALUE;
        int size = Integer.SIZE;
        System.out.println("int 类型可取的最大值" + maxValue);
        System.out.println("int 类型可取的最小值" + minValue);
        System.out.println("int 类型的二进制位数" + size);
    }

}
```

运行结果：

```java
int 类型可取的最大值2147483647
int 类型可取的最小值-2147483648
int 类型的二进制位数32
```

### 4.2 Character 类

Character 类在对象中包装一个基本类型为 char 的值。一个 Character 对象包含类型为 char 的单个字段。

#### 4.2.1 构造方法

Character 类提供了一个构造方法：

`Character(char value)`：很少使用。

#### 4.2.2 常用方法

* `char charValue()`：返回此 Character 对象的值；
* `int compareTo(Character anotherCharacter)`：返回此 Character 对象的值，根据数字比较两个 Character 对象，若这两个对象相等则返回 0 ；
* `boolean equals(Object obj)`：将调用该方法的对象与指定的对象相比较；
* `char toUpperCase(char ch)`：将字符参数转换为大写；
* `char toLowerCase(char ch)`：将字符参数转换为小写；
* `String toString()`：返回一个表示指定 char 值的 String 对象；
* `char charValue()`：返回此 Character 对象的值；
* `boolean isUpperCase(char ch)`：判断指定字符是否是大写字符；
* `boolean isLowerCase(char ch)`：判断指定字符是否是小写字符。

更多常用方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/lang/Character.html)。

### 4.3 Boolean 类

Boolean 类将基本类型为 boolean 的值包装在一个对象中。一个 Boolean 类型的对象只包含一个类型为 boolean 的字段。此外，此类还为 boolean 和 String 的相互转换提供了许多方法，并提供了处理 boolean 时非常有用的其他一些常量和方法。

#### 4.3.1 构造方法

Boolean 类提供了如下两个构造方法：

1. `Boolean(boolean value)`：创建一个表示 value 参数的 boolean 对象（很少使用）；
2. `Boolean(String s)`：以 String 变量作为参数，创建 boolean 对象。此时，如果传入的字符串不为 null，且忽略大小写后的内容等于 “true”，则生成 Boolean 对象值为 true，反之为 false。（很少使用）。

#### 4.3.2 常用方法

* `boolean booleanValue()`：将 Boolean 对象的值以对应的 boolean 值返回；
* `boolean equals(Object obj)`：判断调用该方法的对象与 obj 是否相等，当且仅当参数不是 null，而且与调用该方法的对象一样都表示同一个 boolean 值的 Boolean 对象时， 才返回 true；
* `boolean parseBoolean(Sting)`：将字符串参数解析为 boolean 值；
* `String toString()`：返回表示该 boolean 值的 String 对象；
* `boolean valueOf(String s)`：返回一个用指定的字符串表示值的 boolean 值。

更多常用方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/lang/Boolean.html)。

#### 4.3.3 常用常量

* `TRUE`：对应基值 true 的 Boolean 对象；
* `FALSR`：对应基值 false 的 Boolean 对象；
* `TYPE`：表示基本类型 Class 实例。

## 5. 装箱和拆箱

装箱就是基本数据类型向包装类转换；拆箱就是包装类向基本数据类型转换。装箱和拆箱又有自动和手动之分。

实现装箱的实例如下：

```java
public class WrapperClassDemo2 {

    public static void main(String[] args) {
        // 自动装箱
        int num1 = 19;
        Integer num2 = num1;
        System.out.println("num2=" + num2);

        // 手动装箱
        Integer num3 = new Integer(20);
        System.out.println("num3=" + num3);
    }

}
```

运行结果：

```java
num2=19
num3=20
```

自动装箱就是直接将一个基本数据类型的变量，赋值给对应包装类型的变量；手动装箱就是调用包装类的构造方法（在 Java14 中已经过时，不推荐这样的操作）。

实现拆箱的实例如下：

```java
public class WrapperClassDemo3 {

    public static void main(String[] args) {
        // 自动拆箱
        Integer num1 = 19;
        int num2 = num1;
        System.out.println("num2=" + num2);

        // 手动拆箱
        int num3 = num1.intValue();
        System.out.println("num3=" + num3);
    }

}
```

运行结果：

```java
num2=19
num3=19
```

自动拆箱就是直接将一个包装类型的变量，赋值给对应的基本数据类型变量；手动拆箱通过调用对应包装类下的 `xxxValue()` 方法来实现。

## 6. 小结

通过本小节的学习，我们知道了包装类就是将基本数据类型包装成的类，它有属性、方法，可以对象化交互。除了 `int` 对应的包装类名称为 `Integer` 以及 `char` 对应的包装类名称 `Character`，其他 `6` 种数据类型对应的包装类，命名都为其基本数据类型的首字母的大写。装箱就是基本数据类型向包装类转换，拆箱就是包装类向基本数据类型转换。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
