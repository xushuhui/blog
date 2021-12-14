---
title: Java从零开始（9）Java 基本数据类型
zhihu-url: https://zhuanlan.zhihu.com/p/400050130
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 基本数据类型


前面我们学习了变量，每个变量都是有数据类型的。在 Java 中，一共有两大数据类型：

* 基本数据类型（内置数据类型）
* 引用数据类型

本小节，我们将主要介绍**基本数据类型**，除了我们在示例中经常用到的 `int` 类型外，`Java` 语言还有其他 `7` 种基本数据类型。它们由语言预先定义，也称内置数据类型，顾名思义，它们是 `Java` 语言中最基本的数据类型。另外注意：引用数据类型并不属于基本数据类型的范畴，我们将在后面讲解。

Java 语言支持 8 种基本数据类型：`byte`，`short`，`int`，`long`，`float`，`double`，`char` 和 `boolean`

## 1. 分类和特点

### 1.1 分类

基本数据类型定义在 `Java`**类型系统**中，在类型系统中，又可将基本类型分为三大类：

* **数字型**：包含 `short`、`int`、`long`、`float` 和 `double`，只能存储数字数据；
* **字符型**：`byte` 和 `char`，只能存储**字符数据**；
* **布尔类型**：`boolean`，只能存储 `true`（真） 或 `false`（假）。

### 1.2 范围和默认值

基本类型都有固定的大小，因此，基本类型的值会被限制在一个范围内。

另外，在声明变量时，不一定总是需要分配值。编译器会将已声明但未初始化的字段设置合理的默认值。**但是，建议为每个变量都赋初始值，这是很好的编程习惯。**

具体的范围和默认值，请参考下表：

|类型|大小（位）|最小值|最大值|默认值|
|----|----------|------|------|------|
| byte  |8 |-128                      |+127                         |0         |
|char   |16|0                         |+655,35                      |‘\u0000’|
|short  |16|-32,768                   |+32,767                      |0         |
|int    |32|-2,147,483,648            |+2,147,483,647               |0         |
|long   |64|-9,223,372,036,854,775,808|+9,223,372,036,854,775,807   |0L        |
|float  |32|1.4 E-45                  |3.402,823,5 E+38             |0.0f      |
|double |64|4.9 E-324                 |1.797,693,134,862,315,7 E+308|0.0d      |
|boolean|–|–                        |–                           |false     |

### 1.3 常见错误

在编码过程中，要特别注意变量**不要超出**其类型的范围，请查看如下示例：

```java
public class MemoryOverflow {
    public static void main(String[] args) {
      	// Integer.MAX_VALUE 用来int类型能存储的最大值
        int maximumValue = Integer.MAX_VALUE;
      	System.out.println("int类型能够存储的最大值为：" + maximumValue);
        System.out.println(maximumValue + 1);
    }
}
```

运行结果：

```java
int 类型能够存储的最大值为2147483647
-2147483648
```

我们发现 `int` 类型能存储的最大值加 1 之后，变成了一个负数，这个数其实就是 `int` 类型能够存放的最小值，这是因为加 1 之后变量超出了 `int` 类型能够存储的最大值，这就是我们常常说的**内存溢出错误**。

还要特别注意一点，由于 Java 是**强类型**的，每个变量都有一个类型，只有给定种类的值能够存储到该变量中。例如，不能将浮点型的数字赋值给整型变量：

```java
public class IncompatibleTypeError{
    public static void main(String[] args) {
        int age;
      	age = 20.5f;
    }
}
```

源代码在编译阶段就会报错：

```java
$ javac IncompatibleTypeError.java
IncompatibleTypeError.java:4: 错误: 不兼容的类型: 从float转换到int可能会有损失
      	age = 20.5f;
      	      ^
1 个错误
```

## 2. 整型

整型就是整数类型，和数学中的整数意义相同，例如：

```java
// 声明一个整型变量age
int age = 10;
// 声明一个整型变量 score
int score = 100;
```

整型除了 `int` 类型，还有 `short` 和 `long` 类型，也就是短整型和长整型，他们用于存放不同范围的整数值。需要注意的是，`long` 类型的值后面要以大写字母 `L` 或小写字母 `l` 结尾。请阅读以下代码：

```java
// 声明一个短整型变量a
short a = 20;
// 声明一个长整型变量b
long b = 100l;
// 声明一个长整型变量c
long c = 100L;
```

> **Tips**：对于长整型，推荐后面总以大写字母 `L` 结尾，因为小写字母 `l` 与数字 `1` 容易混淆。

和数学一样，计算机中的整型也有**进制**的概念，`Java` 可以通过以下几种进制系统表示：

* **十进制**：基数为 10，由 0 到 9 组成； 这是我们日常生活中使用的数字系统；
* **十六进制**：基数为 16，由数字 0 到 9 和字母 A 到 F 组成；
* **二进制**：基数为 2，由数字 0 和 1 组成。

在 `Java` 中，十六进制和二进制需要以前缀区分，前缀 `0x` 表示十六进制，前缀 `0b` 表示二进制，我们来看一个示例：

```java
// 十进制表示数字10
int decimalValue = 10;
// 十六进制表示数字10
int hexadecimalValue = 0xA;
// 二进制表示数字10
int binaryValue = 0b1010;
```

## 3. 浮点型

浮点类型就是数学中的**小数**，在 Java 中，浮点型又可分为单精度浮点型 (`float`) 和双精度浮点型 (`double`)。需要注意的是，`float` 类型的值必须要以大写字母 `F` 或小写字母 `f` 结尾，我们来看一个示例：

```java
float f1 = 1.2f;
float f2 = 10.56F;
```

`double` 类型的值可以以字母 `D` 或 `d` 结尾，但不做强制：

```java
// 声明三个double类型变量，其值都为123.3
double d1 = 123.3;
double d2 = 123.3D;
double d3 = 123.3d;
```

另外浮点类型也可以使用字母 `E` 或 `e`，后面跟上次方的数值，来表示数学中的科学计数法，我们来看一个示例：

```java
// 科学计数法表示的 123.4
double d2 = 1.234e2;
float f1  = 1.234e2f;
// 科学计数法表示 200022.1
double d3 = 2.000221e5
```

## 4. 布尔类型

布尔类型 `boolean` 只有两个可能的值：`true` 和 `false`，它是用于跟踪判断条件真假的标志。

```java
boolean b1 = true;
boolean b2 = false;
```

我们将在后面学习条件语句的时候经常使用。这时布尔类型才能发挥它真正的价值，这里不再赘述。

## 5. 字符类型

字符类型 `char` 表示**一个字符**。Java 中的 `char` 类型除了可表示标准的 `ASCII` 外，还可以表示一个 `Unicode` 字符：

```java
public class CharType {
    public static void main(String[] args) {
        // 存放一个 ASCII 码
        char a = 97;
        // 存放一个Unicode字符
        char zh = '慕';
        System.out.println(a);
        System.out.println(zh);
    }
}
```

运行结果：

```java
a
慕
```

Java 语言还为 `char` 和 `String` 类型（String 类型将在后面单独讲解），支持一些特殊的转义序列：

|符号|代表字符|
|----|--------|
| \ n   | 换行（0x0a）                 |
|\ r    | 回车（0x0d）                 |
|\F     | 换页（0x0c）                 |
|\ b    | 退格（0x08）                 |
|\ s    | 空格（0x20）                 |
|\ t    | 制表符                       |
| "     | 双引号                       |
|’     |单引号                        |
| \\    | 反斜杠                       |
| \ddd  | 八进制字符（ddd）            |
|\ uxxxx| 十六进制 Unicode 字符（xxxx）|

## 6. 小结

学完了本小节，我们来回顾一下，Java 提供两大类型：基本数据类型和引用数据类型。一共有 8 种基本数据类型。基本类型又可分为数字型（整型和浮点型），字符型以及布尔型。在编码时要注意，为变量所赋的值不要超出其类型的范围。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
