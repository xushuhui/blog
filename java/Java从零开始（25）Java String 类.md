---
title: Java 从零开始（25）Java String 类
zhihu-url: https://zhuanlan.zhihu.com/p/408515348
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java String 类

在前面的 [Java 字符串](http://www.imooc.com/wiki/javalesson/javastring.html) 小节，我们就已经接触了`String`类，但并未提及`String`类相关的操作，现在有了面向对象相关前置知识，我们知道了类下面可以有相关的操作，作为`Java`语言的内置类，`String`类也为我们预先定义了很多好用的方法，本小节我们将介绍`String`类的常用方法，并结合示例辅助理解。

## 1. String 对象的创建

String 对象的创建有两种方式。

第 1 种方式就是我们最常见的创建字符串的方式：

```java
String str1 = "Hello, ";
```

第 2 种方式是对象实例化的方式，使用`new`关键字，并将要创建的字符串作为构造参数：

```java
String str2 = new String("Hello, Java");
```

如果调用 String 类的无参构造方法，则会创建一个空字符串：

```java
String str3 = new String();
```

此处的`str3`就是一个空字符串。但注意，这种方式很少使用。

## 2. 获取字符串长度

可以使用`length()`方法来获取字符串的长度。例如：

```java
public class StringMethod1 {
    public static void main(String[] args) {
        // 创建 String 对象 str
        String str = "hello world!";
        // 调用对象下 length() 方法，并使用 int 类型变量接收返回结果
        int length = str.length();
        System.out.println("str 的长度为：" + length);
    }
}
```

运行结果：

```java
str1 的长度为：12
```

注意，`hello world!`中的空格也算一个字符。

## 3. 字符串查找

### 3.1 获取指定位置字符

可以使用`char charAt(int index)`方法获取字符串指定位置的字符。它接收一个整型的`index`参数，指的是**索引位置**，那什么是索引位置呢？例如，有一字符串`I love Java`，其每个字符的索引如下图所示：

![](https://xushuhui.gitee.io/image/imooc/5ea6a741094b29fb25000864.jpg)

可以从图示中看出，索引下标从`0`开始。假如我们要获取字符`J`，则为方法传入参数`7`即可：

```java
public class StringMethod2 {
    public static void main(String[] args) {
        String str = "I love Java";
        char c = str.charAt(7);
        System.out.println("索引位置为 7 的字符为：" + c);
    }
}
```

运行结果：

```java
索引位置为 7 的字符为：J
```

### 3.2 查找字符串位置

这里介绍查找字符串位置的两个方法：

* `indexOf()` 获取字符或子串在字符串中第一次出现的位置。
* `lasIndexOf()` 获取字符或子串在字符串中最后一次出现的位置。

> 这里的**子串**指的就是字符串中的连续字符组成的子序列。例如，字符串`Hello`就是字符串`Hello Java`的子串。

`indexOf()`有多个重载方法，这里我们只演示其中最常用的两个。

1. 获取字符在字符串中第一次出现的位置：

```java
public class StringMethod2 {
    public static void main(String[] args) {
        String str = "I love Java, I love imooc!";
        int i = str.indexOf('a');
        System.out.println("字符 a 在字符串 str 第一次出现的位置为：" + i);
    }
}
```

运行结果：

```java
字符 a 在字符串 str 第一次出现的位置为：8
```

1. 获取子串在字符串中第一次出现的位置：

```java
public class StringDemo2 {
    public static void main(String[] args) {
        String str = "I love Java, I love imooc!";
        int i = str.indexOf("love");
        System.out.println("子串 love 在字符串 str 第一次出现的位置为：" + i);
    }
}
```

运行结果：

```java
子串 love 在字符串 str 第一次出现的位置为：2
```

关于`lastIndexOf()`，我们也只演示最常用的两个重载方法。

1. 获取字符在字符串中最后一次出现的位置：

```java
public class StringMethod2 {
    public static void main(String[] args) {
        String str = "I love Java, I love imooc!";
        int i = str.lastIndexOf('e');
        System.out.println("字符 e 在字符串 str 最后一次出现的位置为：" + i);
    }
}
```

运行结果：

```java
字符 e 在字符串 str 最后一次出现的位置为：18
```

1. 获取子串在字符串中最后一次出现的位置：

```java
public class StringMethod2 {
    public static void main(String[] args) {
        String str = "I love Java, I love imooc!";
        int i = str.lastIndexOf("I love");
        System.out.println("字串 I love 在字符串 str 最后一次出现的位置为：" + i);
    }
}
```

运行结果：

```java
字串 I love 在字符串 str 最后一次出现的位置为：13
```

需要特别注意的是，以上方法的参数都是区分大小写的。这也就意味着，你永远无法在`I love Java`中查找到字符`E`。如果没有查找，上述方法都会返回一个整型值：`-1`。我们来看以下示例：

```java
public class StringMethod2 {
    public static void main(String[] args) {
        String str = "I love Java";
        int i = str.indexOf('E');
        System.out.println(i);
    }
}
```

运行结果：

```java
-1
```

## 4. 字符串截取

字符串的截取也称为**获取子串**，在实际开发中经常用到，可以使用`substring()`方法来获取子串，String 类中有两个重载的实例方法：

* `String substring(int beginIndex)` 获取从`beginIndex`位置开始到结束的子串。
* `String substring(int beginIndex, int endIndex)` 获取从`beginIndex`位置开始到`endIndex`位置的子串（不包含`endIndex`位置字符）。

关于这两个方法的使用，我们来看一个实例：

```java
public class StringMethod3 {
    public static void main(String[] args) {
        String str = "I love Java";
        String substring = str.substring(2);
        String substring1 = str.substring(2, 6);
        System.out.println("从索引位置 2 到结束的子串为：" + substring);
        System.out.println("从索引位置 2 到索引位置 6 的子串为：" + substring1);
    }
}
```

运行结果：

```java
从索引位置 2 到结束的子串为：love Java
从索引位置 2 到索引位置 6 的子串为：love
```

要特别注意，方法签名上有两个参数的`substring(int beginIndex, int endIndex)`方法，截取的子串不包含`endIndex`位置的字符。

## 5. 字符串切割

### 5.1 切割为字串数组

`String[] split(String regex)`方法可将字符串切割为子串，其参数`regex`是一个正则表达式分隔符，返回字符串数组。例如，我们使用空格作为分隔符来切割`I love Java`字符串，结果将返回含有 3 个元素的字符串数组：

```java
public class StringMethod4 {
    public static void main(String[] args) {

        String str1 = "I love Java";
        // 将字符串 str1 以空格分隔，并将分割结果赋值给 strArr 数组
        String[] strArr = str1.split(" ");
        // 遍历数组，打印每一个元素
        for (String str: strArr) {
            System.out.print(str + '\t');
        }

    }
}
```

运行结果：

```java
I	love	Java
```

注意，有几种特殊的分隔符：`*``^``:``|``.``\`，要使用转义字符转义。例如：

```java
// 以*切割
String str2 = "I*love*Java";
String[] strArr2 = str2.split("\\*");

// 以、切割
String str3 = "I\\love\\Java";
String[] strArr4 = str3.split("\\\\");

// 以|切割
String str4 = "I|love|Java";
String[] strArr4 = str4.split("\\|");
```

另外，还有一个重载方法`String[] split(String regex, int limit)`，其第二个参数`limit`用以控制正则匹配被应用的次数，因此会影响结果的长度，此处不再一一举例介绍。

### 5.2 切割为 byte 数组

在实际工作中，网络上的数据传输就是使用二进制字节数据。因此字符串和字节数组之间的相互转换也很常用。

我们可以使用`getBytes()`方法将字符串转换为`byte`数组。例如：

```java
public class StringMethod4 {
    public static void main(String[] args) {
        String str2 = "我喜欢 Java";
        System.out.println("将字符串转换为 byte 数组：");
        // 将字符串转换为字节数组
        byte[] ascii = str2.getBytes();
        // 遍历字节数组，打印每个元素
        for (byte aByte : ascii) {
            System.out.print(aByte + "\t");
        }
    }
}
```

运行结果：

```java
将字符串转换为 byte 数组：
-26	-120	-111	-27	-106	-100	-26	-84	-94	74	97	118	97
```

将字节数组转换为字符串的方法很简单，直接实例化一个字符串对象，将字节数组作为构造方法的参数即可：

```java
// 此处的 ascii 为上面通过字符串转换的字节数组
String s = new String(ascii);
```

## 6. 字符串大小写转换

字符串的大小写转换有两个方法：

* `toLowerCase()` 将字符串转换为小写

* `toUpperCase()` 将字符串转换为大写

我们来看一个实例：

```java
public class StringMethod5 {
    public static void main(String[] args) {
        String str = "HELLO world";
        String s = str.toLowerCase();
        System.out.println("字符串 str 为转换为小写后为：" + s);
        String s1 = s.toUpperCase();
        System.out.println("字符串 s 为转换为大写后为：" + s1);
    }
}
```

运行结果：

```java
字符串 str 为转换为小写后为：hello world
字符串 s 为转换为大写后为：HELLO WORLD
```

试想，如果想把字符串`HELLO world`中的大小写字母互换，该如何实现呢？

这里可以结合字符串切割方法以及字符串连接来实现：

```java
public class StringMethod5 {
    public static void main(String[] args) {
        String str = "HELLO world";
        // 先切割为数组
        String[] strArr = str.split(" ");
        // 将数组中元素转换大小写并连接为一个新的字符串
        String result = strArr[0].toLowerCase() + " " + strArr[1].toUpperCase();
        System.out.println("字符串 str 的大小写互换后为：" + result);
    }
}
```

运行结果：

```java
字符串 str 的大小写互换后为：hello WORLD
```

当然，实现方式不止一种，你可以结合所学写出更多的方式。

## 7. 字符串比较

`String`类提供了`boolean equals(Object object)`方法来比较字符串内容是否相同，返回一个布尔类型的结果。

需要特别注意的是，在比较字符串内容是否相同时，必须使用`equals()`方法而不能使用`==`运算符。我们来看一个示例：

```java
public class StringMethod6 {
    public static void main(String[] args) {
        // 用两种方法创建三个内容相同的字符串
        String str1 = "hello";
        String str2 = "hello";
        String str3 = new String("hello");
        System.out.println("使用 equals() 方法比较 str1 和 str2 的结果为：" + str1.equals(str2));
        System.out.println("使用==运算符比较 str1 和 str2 的结果为：" + (str1 == str2));
        System.out.println("使用==运算符比较 str1 和 str3 的结果为：" + (str1 == str3));
    }
}
```

运行结果：

```java
使用 equals() 方法比较 str1 和 str2 的结果为：true
使用==运算符比较 str1 和 str2 的结果为：true
使用 equals() 方法比较 str1 和 str3 的结果为：true
使用==运算符比较 str1 和 str3 的结果为：false
```

代码中三个字符串`str1`，`str2`和`str3`的内容都是`hello`，因此使用`equals()`方法对它们进行比较，其结果总是为`true`。

注意观察执行结果，其中使用`==`运算符比较`str1`和`str2`的结果为`true`，但使用`==`运算符比较的`str1`和`str3`的结果为`false`。这是因为`==`运算符比较的是两个变量的**地址**而不是内容。

要探究其原因，就要理解上述创建字符串的代码在计算机内存中是如何执行的。下面我们通过图解的形式来描述这三个变量是如何在内存中创建的。

1. 当执行`String str1 = "hello;"`语句时，会在内存的**栈空间**中创建一个`str1`，在**常量池**中创建一个"hello"，并将`str1`指向`hello`。

![](https://xushuhui.gitee.io/image/imooc/5ea6a77009d06d8e18171011.jpg)

1. 当执行`String str2 = "hello";`语句时，栈空间中会创建一个`str2`，由于其内容与`str1`相同，会指向常量池中的同一个对象。所以`str1`与`str2`指向的地址是相同的，这就是`==`运算符比较`str1`和`str2`的结果为`true`的原因。

![](https://xushuhui.gitee.io/image/imooc/5ea6a78a09d7219118861066.jpg)

1. 当执行`String str3 = new String("hello");`语句时，使用了`new`关键字创建字符串对象，由于对象的实例化操作是在内存的**堆空间**进行的，此时会在栈空间创建一个`str3`，在堆空间实例化一个内容为`hello`的字符串对象，并将`str3`地址指向堆空间中的`hello`，这就是`==`运算符比较`str1`和`str3`的结果为`false`的原因。

![](https://xushuhui.gitee.io/image/imooc/5ea6a7a209c7c55c20171125.jpg)

## 8. 小结

本小节我们介绍了 Java `String`类的常用方法：

* 使用`length()`方法可以获取字符串长度；
* 使用`charAt()`、`indexOf()`以及`lastIndexOf()`方法可以对字符串进行查找；
* `substring()`方法可以对字符串的进行截取，`split()`、`getBytes()`方法可以将字符串切割为数组；
* `toLowerCase()`和`toUpperCase()`方法分别用于大小写转换，使用`equals()`方法对字符串进行比较，这里要注意，对字符串内容进行比较时，永远都不要使用`==`运算符。

这些方法大多有重载方法，实际工作中，要根据合适的场景选用对应的重载方法。

当然，本小节还有很多未介绍到的方法，使用到可以翻阅官网文档来进行学习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
