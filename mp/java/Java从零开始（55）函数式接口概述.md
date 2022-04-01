---
title: Java 从零开始（55）函数式接口概述
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 函数式接口概述

在 Java 里面，所有的方法参数都是有固定类型的，比如将数字 9 作为参数传递给一个方法，它的类型是 int；字符串 “9” 作为参数传递给方法，它的类型是 String。那么 Lambda 表达式的类型由是什么呢？通过本节我们学习什么是函数式接口，它与 Lambda 表达式的关系。

## 1. 什么是函数式接口

函数式接口（Functional Interface）就是一个有且仅有一个抽象方法，但是可以有多个非抽象方法的接口，它可以被隐式转换为 Lambda 表达式。

> **Tips：** 换句话说函数式接口就是 Lambda 表达式的类型。

在函数式接口中，单一方法的命名并不重要，只要方法签名和 Lambda 表达式的类型匹配即可。

> **Tips：** 通常我们会为接口中的参数其一个有意义的名字来增加代易读性，便于理解参数的用途。

函数式接口有下面几个特点：

1. 接口有且仅有一个抽象方法；
2. 允许定义静态方法；
3. 允许定义默认方法；
4. 允许 `java.lang.Object` 中的 `public` 方法；
5. 推荐使用 `@FunctionInterface` 注解（如果一个接口符合函数式接口的定义，加不加该注解都没有影响，但加上该注解可以更好地让编译器进行检查）。

我们来看函数式接口的例子：

```java
@FunctionalInterface
interface TestFunctionalInterface
{
    //抽象方法
    public void doTest();
    //java.lang.Object 中的 public 方法
    public boolean equals(Object obj);
    public String toString();
    //默认方法
    public default void doDefaultMethod(){System.out.println("call dodefaultMethod");}
    //静态方法
    public static void doStaticMethod(){System.out.println("call doStaticMethod");}

    public static void main(String...s){
        //实现抽象方法
        TestFunctionalInterface test = ()->{
            System.out.println("call doTest");
        };
        //调用抽象方法
        test.doTest();
        //调用默认方法
        test.doDefaultMethod();
        //调用静态方法
        TestFunctionalInterface.doStaticMethod();
        //调用 toString 方法
        System.out.println(test.toString());
    }
}
```

我们将得到如下结果：

```java
call doTest
call dodefaultMethod
call doStaticMethod
com.github.x19990416.item.TestFunctionalInterface$$Lambda$1/0x00000008011e0840@63961c42
```

我们通过 `toString` 方法可以发现，`test` 对象被便已成为 `TestFunctionalInterface` 的一个 Lambda 表达式。

## 2. @FunctionalInterface

接下来我们再来看下 `@FunctionalInterface`注解的作用：

首先我们定义一个接口 `TestFunctionalInterface`，包含两个方法 `doTest1` 和 `doTest2`：

```java
interfact TestFunctionalInterface{
    //一个抽象方法
    public void doTest1();
    //另一个抽象方法
    public void doTest2();
}
```

此时对于编译期而言我们的代码是没有任何问题的，它会认为这就是一个普通的接口。当我们使用 `@FunctionalInterface` 后：

```java
//这是一个错误的例子！！！！
@FunctionalInterface
interfact TestFunctionalInterface{
    //一个抽象方法
    public void doTest1();
    //另一个抽象方法
    public void doTest2();
}
```

此时，会告诉编译器这是一个函数式接口，但由于接口中有两个抽象方法，不符合函数式接口的定义，此时编译器会报错：

```java
Multiple non-overriding abstract methods found in interface
```

## 3. 常用的函数式接口

JDK 8 之后新增了一个函数接口包 `java.util.function` 这里面包含了我们常用的一些函数接口：

|接口|参数|返回类型|说明|
|----|----|--------|----|
|Predicate     |T     |boolean|接受一个输入参数 `T`，返回一个布尔值结果                              |
|Supplier      |None  |T      |无参数，返回一个结果，结果类型为 `T`                                  |
|Consumer      |T     |void   |代表了接受一个输入参数 `T` 并且无返回的操作                           |
|Function<T,R> |T     |R      |接受一个输入参数 `T`，返回一个结果 `R`                                |
|UnaryOperator |T     |T      |接受一个参数为类型 `T`, 返回值类型也为 `T`                             |
|BinaryOperator|(T，T)|T      |代表了一个作用于于两个同类型操作符的操作，并且返回了操作符同类型的结果|

### 3.1 Predicate

> 条件判断并返回一个 Boolean 值，包含一个抽象方法 (test) 和常见的三种逻辑关系 与 (and) 、或 (or) 、非 (negate) 的默认方法。

Predicate 接口通过不同的逻辑组合能够满足我们常用的逻辑判断的使用场景。

```java
import java.util.function.Predicate;

public class DemoPredicate {
    public static void main(String[] args) {
        //条件判断
        doTest(s -> s.length() > 5);
        //逻辑非
        doNegate(s -> s.length() > 5);
        //逻辑与
        boolean isValid = doAnd(s -> s.contains("H"),s-> s.contains("w"));
        System.out.println("逻辑与的结果："+isValid);
        //逻辑或
        isValid = doOr(s -> s.contains("H"),s-> s.contains("w"));
        System.out.println("逻辑或的结果："+isValid);
    }

    private static void doTest(Predicate<String> predicate) {
        boolean veryLong = predicate.test("Hello World");
        System.out.println("字符串长度很长吗：" + veryLong);
    }

    private static boolean doAnd(Predicate<String> resource, Predicate<String> target) {
        boolean isValid = resource.and(target).test("Hello world");
        return isValid;
    }

    private static boolean doOr(Predicate<String> one, Predicate<String> two) {
        boolean isValid = one.or(two).test("Hello world");
        return isValid;
    }
    private static void doNegate(Predicate<String> predicate) {
        boolean veryLong = predicate.negate().test("Hello World");
        System.out.println("字符串长度很长吗：" + veryLong);
    }
}
```

结果如下：

```java
字符串长度很长吗：true
字符串长度很长吗：false
逻辑与的结果：true
逻辑或的结果：true
```

### 3.2 Supplier

> 用来获取一个泛型参数指定类型的对象数据（生产一个数据），我们可以把它理解为一个工厂类，用来创建对象。

Supplier 接口包含一个抽象方法 `get`，通常我们它来做对象转换。

```java
import java.util.function.Supplier;

public class DemoSupplier {
    public static void main(String[] args) {
        String sA = "Hello ";
        String sB = "World ";
        System.out.println(
                getString(
                        () -> sA + sB
                )
        );
    }

    private static String getString(Supplier<String> stringSupplier) {
        return stringSupplier.get();
    }
}
```

结果如下：

```java
Hello World
```

上述例子中，我们把两个 String 对象合并成一个 String。

### 3.3 Consumer

> 与 Supplier 接口相反，Consumer 接口用于消费一个数据。

Consumer 接口包含一个抽象方法 `accept` 以及默认方法 `andThen` 这样 Consumer 接口可以通过 `andThen` 来进行组合满足我们不同的数据消费需求。最常用的 Consumer 接口就是我们的 `for` 循环，`for` 循环里面的代码内容就是一个 Consumer 对象的内容。

```java
import java.util.function.Consumer;

public class DemoConsumer {
    public static void main(String[] args) {
        //调用默认方法
        consumerString(s -> System.out.println(s));
        //consumer 接口的组合
        consumerString(
                // toUpperCase() 方法，将字符串转换为大写
                s -> System.out.println(s.toUpperCase()),
                // toLowerCase() 方法，将字符串转换为小写
                s -> System.out.println(s.toLowerCase())
        );
    }

    private static void consumerString(Consumer<String> consumer) {
        consumer.accept("Hello");
    }

    private static void consumerString(Consumer<String> first, Consumer<String> second) {
        first.andThen(second).accept("Hello");
    }
}
```

结果如下：

```java
Hello
HELLO
hello
```

在调用第二个 `consumerString` 的时候我们通过 `andThen` 把两个 `Consumer` 组合起来，首先把 `Hello` 全部转变成大写，然后再全部转变成小写。

### 3.4 Function

> 根据一个类型的数据得到另一个类型的数据，换言之，根据输入得到输出。

Function 接口有一个抽象方法 `apply` 和默认方法 `andThen`，通过 `andThen` 可以把多个 `Function` 接口进行组合，是适用范围最广的函数接口。

```java
import java.util.function.Function;

public class DemoFunction {
    public static void main(String[] args) {
        doApply(s -> Integer.parseInt(s));
        doCombine(
                str -> Integer.parseInt(str)+10,
                i -> i *= 10
        );
    }

    private static void doApply(Function<String, Integer> function) {
        int num = function.apply("10");
        System.out.println(num + 20);
    }
    private static void doCombine(Function<String, Integer> first, Function<Integer, Integer> second) {
        int num = first.andThen(second).apply("10");
        System.out.println(num + 20);
    }
}
```

结果如下：

```java
30
220
```

上述四个函数接口是最基本最常用的函数接口，需要熟悉其相应的使用场景并能够熟练使用。 `UnaryOperator` 和 `BinaryOperator` 作用与 `Funciton` 类似，大家可以通过 `Java` 的源代码进一步了解其作用。

## 4. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a92e70939869507930502.jpg)

本节，我们主要阐述了函数式接口的定义以及其与 Lambda 表达式的关系。并对新增的 `java.util.function` 包中常用的函数式接口进行了解释。这些接口常用于集合处理中（我们将在后续的内容进一步学习），关于函数式接口主要记住一点，那就是：

> 接口有且仅有一个抽象方法

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
