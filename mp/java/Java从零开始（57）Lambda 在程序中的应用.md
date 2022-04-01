---
title: Java 从零开始（57）Lambda 在程序中的应用
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 程序中的应用

通过前面的内容我们对于 Lambda 表达式以及函数式编程已经有了一定的了解，对于集合方面的使用也有了概念，那么，在本节我们将从一个日志改造的例子出发，探讨下如何在我们的程序中更好地使用 Lambda 表达式。

> **Tips：** 本节内容有点分散，主要是启发大家的思路

## 1. 让我们的类支持 Lambda 表达式

日志记录工具是我们平时用的最多的一个工具，比如 SLF4J、Log4j 等，可以帮助我们快速查看程序的信息、定位问题，也在一定程度上增加了系统开销，通常我们在写 `debug` 日志的时候为了降低会进行日志级别的判定 （在本例中我们使用的是 Log4j 2）

```java
public class DemoLogger {
    public static void main(String[] args) {
        Logger logger = LogManager.getLogger(DemoLogger.class);
        if(logger.isDebugEnabled()){
            logger.debug("这是一个 debug 日志");
        }
    }
}
```

想必上面的代码应该都非常熟悉，在 Log4j 2 中提供了 Lambda 表达式实现的日志记录方法，对于上述代码我们可以简化为：

```java
public class DemoLogger {
    public static void main(String[] args) {
        Logger logger = LogManager.getLogger(DemoLogger.class);
        logger.debug(()->"这是一个 debug 日志");
    }
}
```

通过查看源代码我们可以发现 Logger 对象提供了一个 Supplier 的 `debug` 方法：

```java
@Override
public void debug(final Supplier<?> msgSupplier) {
    logIfEnabled(FQCN, Level.DEBUG, null, msgSupplier, (Throwable) null);
}

@Override
public void logIfEnabled(final String fqcn, final Level level, final Marker marker, final Supplier<?> msgSupplier,
        final Throwable t) {
    if (isEnabled(level, marker, msgSupplier, t)) {
        logMessage(fqcn, level, marker, msgSupplier, t);
    }
}
```

在这个方法中，它通过 `logIfEnabled` 判断是否为 `debug` 进而决定是否调用 `supplier` 对象的内容。 这给了我们一个启发，那就是：

> 我们可以运用 `java.util.funciton` 中的接口来重新封装我们原有的类来支持 Lambda 表达式，进而简化我们的代码。

## 2. 多重继承

在 Java 中接口是允许多重继承的，那么如果多个接口有着相同的默认方法的情况下会怎么样呢？

```java
public class MultipleInterface {
    public interface RedBox{
        public default String color(){
            return "red";
        }
    }

    public interface BlueBox{
        public default String color(){
            return "blue";
        }
    }

    public class CombineBox implements RedBox,BlueBox{

    }
}
```

在上面的代码中，我们定义了两个接口 `RedBox` 和 `BlueBox` 都有相同的默认方法 `color`，类 `CombineBox` 同时实现 `RedBox` 和 `BlueBox`，此时，由于编译器不清楚应该继承哪个接口，所以报错：

```java
MultipleInterface.CombineBox inherits unrelated defaults for color() from types MultipleInterface.RedBox and MultipleInterface.BlueBox
```

此时，我们可以使用同方法重载来明确方法内容，当然我们可以通过 `super` 语法来给编译器明确使用哪一个默认接口：

```java
public class CombineBox implements RedBox,BlueBox{
    public String color(){
        return RedBox.super.color();
    }
}
```

上述的内容我们主要是对默认方法的工作原理做了一个简单的介绍，对于默认方法通常有三条定律来帮助我们使用默认方法：

> 1. 类胜于接口：如果在继承链中有声明的方法，那么就可以忽略接口中定义的方法 （这样可以让我们的代码向后兼容）；
> 2. 子类胜于父类：如果一个接口继承了另外一个接口，而且两个接口都定义了一个默认方法，那么子类中定义的方法将生效；
> 3. 如果上述两条都不适用，那么子类要么需要实现该方法，要么将该方法声明成抽象方法 （ `abstract` ）。

## 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a9468099d396b08980252.jpg)

本节从类的重新和接口继承两方面介绍了我们如何重新封装我们的类来支持 Lambda 表达式，以及在函数接口在多继承的情况下出现默认方法冲突时如何去编写我们的代码。大家可以在平时的编码过程中按照上述的思路逐步练习封装自己原来的代码，自然就会有自己的心得体会。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
