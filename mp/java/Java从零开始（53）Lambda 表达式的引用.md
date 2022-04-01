---
title: Java 从零开始（53）Lambda 表达式的引用
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式的引用

所谓 Lambda 表达式的方法引用可以理解为 Lambda 表达式的一种快捷写法，相较于通常的 Lambda 表达式而言有着更高的**可读性**和**重用性**。

> **Tips:** 一般而言，方法实现比较简单、复用地方不多的时候推荐使用通常的 Lambda 表达式，否则应尽量使用方法引用。

Lambda 表达式的引用分为：**方法引用** 和 **构造器引用**两类。

方法引用的格式为：

> 类名：: 方法名

`::` 是引用的运算符，其左边是类名，右边则是引用的方法名。

构造器引用的格式为：

> 类名：:new

同样，`::` 是引用的运算符，其左边是类名，右边则是使用关键字 `new` 表示调用该类的构造函数。构造器引用是一种特殊的引用。

通常引用语法格式有以下 3 种：

* 静态方法引用；
* 参数方法引用；
* 实例方法引用。

接下来我们堆上述 3 种引用逐一进行讲解。

## 1. 静态方法引用

所谓静态方法应用就是调用类的静态方法。

> **Tips：**
>
> 1. 被引用的参数列表和接口中的方法参数一致；
> 2. 接口方法没有返回值的，引用方法可以有返回值也可以没有；
> 3. 接口方法有返回值的，引用方法必须有相同类型的返回值。

我们来看一个例子：

```java
public interface Finder {
    public int find(String s1, String s2);
}
```

这里我们定义了一个 `Finder` 接口，其包含一个方法 `find` ，两个 `String` 类型的输入参数，方法返回值为 `int` 类型。

随后，我们创建一个带有静待方法的类 `StaticMethodClass`：

```java
//创建一个带有静态方法的类
public class StaticMethodClass{
    public static int doFind(String s1, String s2){
        return s1.lastIndexOf(s2);
    }
}
```

在 `StaticMethodClass` 类中，我们查找最后一次出现在字符串 `s1` 中的 `s2` 的位置。

> 在这里`Finder` 接口的 `find` 方法和类 `StaticMethodClass` 的 `doFind` 方法有相同的输入参数（参数个数和类型）完全相同，又因为 `doFind` 方法是一个静态方法，于是我们就可以使用静态方法引用了。

最后，我们在 Lambda 表达式使用这个静态引用：

```java
Finder finder = StaticMethodClass :: doFind;
```

此时，`Finder` 接口引用了 StaticMethodClass 的静态方法 doFind。

## 2. 参数方法引用

参数方法引用顾名思义就是可以将参数的一个方法引用到 Lambda 表达式中。

> **Tips：** 接口方法和引用方法必须有相同的 **参数** 和 **返回值**。

同样我们使用前面的 Finder 接口为例：

```java
public interface Finder {
    public int find(String s1, String s2);
}
```

我们希望 `Finder` 接口搜索参数 `s1` 的出现参数 `s2` 的位置，这个时候我们会使用 Java String 的 indexOf 方法 `String.indexOf` 来进行查询，通常我们是这么使用 Lambda 表达式的：

```java
Finder finder =(s1，s2)-> s1.indexOf(s2);
```

我们发现，接口 `Finder` 的 `find` 方法与 `String.indexOf` 有着相同的方法签名（相同的输入和返回值），那么我们就可以使用参数方法引用来进一步简化：

```java
//参数方法引用
Finder finder = String :: indexOf;

//调用 find 方法
int findIndex = finder.find("abc","bc")
//输出 find 结果。
System.out.println("返回结果："+findIndex)
```

输出为：

```java
返回结果：2
```

此时，编译器会使用参数 `s1` 为引用方法的参数，将引用方法与 `Finder` 接口的 `find` 方法进行类型匹配，最终调用 String 的 indexOf 方法。

## 3. 实例方法引用

实例方法引用就是直接调用实例的方法。

> **Tips：** 接口方法和实例的方法必须有相同的参数和返回值。

我们来看一例子：

首先我们定义一个序列化接口：

```java
public interface Serializer {
    public int serialize(String v1);
}
```

然后我们定一个转换类 StringConverter：

```java
public class StringConverter {
    public int convertToInt(String v1){
        return Integer.valueOf(v1);
    }
}
```

这个时候 `Serializer.serialize` 方法和 `StringConvertor.converToInt` 有着相同的方法签名（即，输入和输出都是相同的），那么，我们可以创建一个 `StringConvertor` 的实例，并通过 Lambda 表达式将其并引用给 `convertToInt()` 方法。

```java
StringConverter stringConverter = new StringConverter();
Serializer serializer = StringConverter::convertToInt;
```

我们在第一行代码中创建了 `StringConverter` 的对象，在第二行代码中，通过 实例方法引用来引用 `StringConverter` 的 `convertToInt` 方法。

## 4. 构造器引用

构造器引用便是引用一个类的构造函数

> **Tips：** 接口方法和对象构造函数的参数必须相同。

其格式如下：

```java
类名 :: new
```

我们来看一个例子：

```java
public interfact MyFactory{
    public String create(char[] chars)
}
```

我们定义了 `MyFactory` 接口 有一个 `create` 方法，接收一个 `char[]` 类型的输入参数，返回值类型为 `String`, 它与 `String(char[] chars)` 这个 `String` 的构造函数有着相同的方法签名。这个时候我们就可以使用构造器引用了：

```java
MyFactory myfactory =  String::new;
```

它等价于 Lambda 表达式：

```java
MyFactory myfactory = chars->new String(chars);
```

## 5. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a8e19094aabd107810143.jpg)

本节我们主要学习了 Lambda 表达式的引用，引用是基于方法调用的事实提供一种简短的语法，让我们无需看完整的代码就能弄明白代码的意图。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
