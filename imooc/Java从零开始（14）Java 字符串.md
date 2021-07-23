# Java 字符串

在 [Java 基本数据类型](http://www.imooc.com/wiki/javalesson/primitivetypes.html)一节中，我们已经知道 Java 中有两大数据类型：一个是基本数据类型，另一个是引用数据类型。

字符串 (`String`) 便属于引用数据类型。字符串在 Java 程序中无处不在，其实我们在第一个 Java 程序中就已经使用过了字符串，打印语句中双引号中的`Hello World`就是一个字符串：`System.out.println("Hello World");`

本小节我们将介绍字符串的基础知识，告诉你如何创建一个字符串，也会讲到 `Java 13`支持的多行字符串的创建方法和注意事项，还会介绍很常用的字符串的连接操作，最后我们也会着重讨论字符串的一个重要特性：不可变性。

需要注意的是，`String`是在`java.lang`包中定义的内置类，了解它的相关操作需要面向对象的前置知识，为了减少初学者的困扰，字符串的相关操作不会在本小节讨论。我们将在完成面向对象的学习后专门讲解`String`类。

## 1. 什么是字符串

字符串 (`string`) 是由零个或多个字符组成的有限序列，它是编程语言中表示文本的数据类型。简单来说，字符串就是一串字符。

## 2. 创建字符串

### 2.1 创建字符和字符串

提到字符串，就很容易与字符关联起来，回忆我们已经学过的字符类型`char`，在代码中我们是这样创建字符的：

```java
// 依次创建三个字符 慕 课 网
char c1 = '慕';
char c2 = '慕';
char c3 = '网';
```

与`char`类型不同的是，`String`类型使用双引号来表示一个字符串（例如：`"字符串文字"`），在代码中我们可以这样创建字符串：

```java
String str = "慕课网";
```

我们再来看一个创建并打印字符串的程序示例：

```java
public class StringTest1 {
    public static void main(String[] args) {
        // 创建一个空字符串
        String str1 = "";
        // 创建一个只包含一个字符的字符串
        String str2 = "q";
        // 创建包含多个字符的字符串
        String str3 = "imooc is great!";
        // 依次打印它们
        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);
    }
}
```

运行结果：

```java

q
imooc is great!
```

字符串可以声明为空，因此第一行将打印一个空行。

### 2.2 创建多行字符串

自`Java 13` 以后，我们可以使用三引号来表示一个多行字符串，被官方称为“文本块”。文本块常用来表示多行的或大段的文字。例如：

```java
public class StringTest3 {
    public static void main(String[] args) {
        String str = """
            Java 很棒！
            慕课网很棒！！
            能够在慕课网学 Java 更棒！！！""";
        System.out.println(str);
    }
}
```

> **Tips:**
>
> 这里需要注意的是，文本块中起始的三引号后面要另起一行，也就是说下面这种写法是错误的：
>
> ```java
> String str = """Java 很棒！
> 慕课网很棒！！
> 能够在慕课网学 Java 更棒！！！""";
> ```

如果我们直接使用`javac`命令编译代码，将会报错：

```java
javac StringTest3.java
StringTest3.java:4: 错误: 文本块 是预览功能，默认情况下禁用。
        String str = """
                     ^
  （请使用 --enable-preview 以启用 文本块）
1 个错误
```

报错告诉我们：文本块是预览功能，默认是禁用的。我们可以给编译器加上一些参数来编译执行代码：

```java
$ javac -source 14 --enable-preview StringTest3.java
$ java --enable-preview StringTest3
Java 很棒！
慕课网很棒！！
能够在慕课网学 Java 更棒！！！
```

## 2. 字符串的连接

我们可以使用加法运算符`+`将字符串和任意其他的数据类型进行连接操作。

可以将两个或多个字符串连接在一起，例如：

```java
public class StringJoinTest1 {
    public static void main(String[] args) {
        // 定义两个字符串
        String str1 = "Hello";
        String str2 = "慕课网";
        // 将字符串str1连接一个空格，再连接str2，最后再连接一个感叹号
        String str3 = str1 + " " + str2 + "！";
        System.out.println(str3);
    }
}
```

运行结果：

```java
$ javac StringJoinTest1.java
$ java StringJoinTest1
Hello 慕课网！
```

我们还可以将字符串与其他类型的值相连接，例如：

```java
public class StringJoinTest2 {
    public static void main(String[] args) {
        // 定义整型变量 age
        int age = 12;
        // 定义浮点型变量 height
        float height = 132.5f;
        // 定义一个字符串变量 name
        String name = "小明";
        // 将上面三个变量连接，拼成一个更长的字符串
        String result = name + "今年" + age + "岁了，" + "他的身高为" + height + "cm";
        System.out.println(result);
    }
}
```

运行结果：

```java
小明今年10岁了，他的身高为132.5cm
```

我们知道加法运算符除了可以连接字符串，也可以进行数学的加法运算。那我们再来看一个实例：

```java
public class StringJoinTest3 {
    public static void main(String[] args) {
        // 定义两个整型变量a、b
        int a = 100;
        int b = 200;
        // 打印一个空字符串+a+b的值
        System.out.println("" + a + b);
    }
}
```

试想屏幕上将会得到什么结果呢？编译执行代码：

```java
$javac StringJoinTest3
java StringJoinTest
100200
```

这里结果没有打印`300`的原因很简单： 按照从左到右的执行顺序，（空字符串 + 数值`100`) 的运算结果应该是字符串`100`，然后再运算（字符串`100`+ 数值`200`)，其最终结果为字符串`100200`。由于运算中有字符串的存在，`+`运算符在这里用于字符串的连接而不是数值的计算。

试想，如果将表达式中的操作数倒过来（即`b + a + ""`），将得到什么结果呢？

## 3. 字符串的不可变性

字符串有一个重要特性：不可变性。也就是说，字符串一经创建便无法修改。我们来看一个示例：

```java
public class StringTest4 {
    public static void main(String[] args) {
        // 创建一个字符串 str
        String str = "hello imooc!";
        // 对 str 重新赋值
        str = "Java is great!";
        System.out.println(str);
    }
}
```

运行结果：

```java
Java is great!
```

童鞋们可能发现，咦？不是说字符串一旦创建就无法修改了吗，怎么`str`变量依然被修改成了新的值？

这里直接给出答案：上述例子中，变量`str`没有修改，而是变量`str`的“**指向**”被修改了。

我们通过图示来描述上述代码的执行流程，首先创建一个`str`，指向"hello imooc!"：

![](//img.mukewang.com/wiki/5e9734fd0917747f15620704.jpg)

当对`str`重新赋值后，在内存中先创建了字符串"Java is great"，然后改变了变量`str`的指向，原来的"hello imooc!"并没有发生改变：

![](//img.mukewang.com/wiki/5e9735110944c9ac15990988.jpg)

## 4. 小结

通过本小节的学习，我们了解了字符类型`char` 是基本数据类型，而字符串类型`String`是引用类型。在创建字符串时，我们可以使用双引号表示，从`Java 13`起，加入了多行字符串的表示方法。字符串的连接操作是使用加法运算符实现的。最后要特别注意字符串的不可变性，这在面试中经常会考到。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
