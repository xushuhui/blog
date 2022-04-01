---
title: Java 从零开始（42）函数式接口
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415767199
---
# 函数式接口

在上个小节的最后，我们提到了函数式接口的概念，也知道了想要使用`Lambda`表达式，则必须依赖函数式接口。本小节我们将学习函数式接口相关的知识，包括**什么是函数式接口**，**为什么需要函数式接口**，**如何自定义一个函数式接口**，**如何创建函数式接口的对象**，以及一些 **Java 内置的函数式接口**的详细介绍等。本小节内容较为简单，但需要读者有 [`Lambda`表达式】(http://www.imooc.com/wiki/javalesson/lambda.htm) 前置知识，学习重点是要了解 Java 内置函数式接口。

## 1. 什么是函数式接口

函数是接口（`Functional Interface`）的定义非常容易理解：只有一个抽象方法的接口，就是函数式接口。可以通过`Lambda`表达式来创建函数式接口的对象。

我们来看一个在之前我们就经常使用的`Runnable`接口，`Runnable`接口就是一个函数式接口，下面的截图为 Java 源码：

![](https://xushuhui.gitee.io/image/imooc/5efe86c8093a092215421363.jpg)

我们看到`Runnable`接口中只包含一个抽象的`run()`方法，并且在接口上标注了一个`@FuncationInterface`注解，此注解就是 Java 8 新增的注解，用来标识一个函数式接口。

## 2. 为什么需要函数式接口

学习了这么久的 Java，我们对 **Java 是纯种的面向对象的编程语言**这一概念，可能有了一定的感触，在 Java 中，**一切皆是对象**。但是随着`Python`、`scala`等语言的兴起，函数式编程的概念得到开发者们的推崇，Java 不得不做出调整以支持更广泛的技术要求。

在面向函数编程的语言中，`Lambda`表达式的类型就是函数，但是**在 Java 中，`Lambda`表达式的类型是对象**而不是函数，他们必须依赖于一种特别的对象类型——函数式接口。所以说，Java 中的`Lambda`表达式就是一个函数式接口的对象。我们之前使用匿名实现类表示的对象，都可以使用`Lambda`表达式来表示。

## 3. 自定义函数式接口

想要自定义一个函数式接口也非常简单，在接口上做两件事即可：

1. **定义一个抽象方法**：注意，接口中只能有一个抽象方法；
2. **在接口上标记`@FunctionalInterface`注解**：当然也可以不标记，但是如果错写了多个方法，编辑器就不能自动检测你定义的函数式接口是否有问题了，所以建议还是写上吧。

```java
/**
 * 自定义函数式接口
 * @author colorful@TaleLin
 */
@FunctionalInterface
public interface FunctionalInterfaceDemo {

    void run();

}

```

由于标记了`@FunctionalInterface`注解，下面接口下包含两个抽象方法的这种错误写法，编译器就会给出提示：

![](https://xushuhui.gitee.io/image/imooc/5efe871409f7c28215660650.jpg)

## 4. 创建函数式接口对象

在上面，我们自定义了一个函数式接口，那么如何创建它的对象实例呢？

我们可以使用匿名内部类来创建该接口的对象，实例代码如下：

```java
/**
 * 测试创建函数式接口对象
 * @author colorful@TaleLin
 */
public class Test {

    public static void main(String[] args) {
        // 使用匿名内部类方式创建函数式接口
        FunctionalInterfaceDemo functionalInterfaceDemo = new FunctionalInterfaceDemo() {
            @Override
            public void run() {
                System.out.println("匿名内部类方式创建函数式接口");
            }
        };
        functionalInterfaceDemo.run();
    }

}
```

运行结果：

```java
匿名内部类方式创建函数式接口
```

现在，我们学习了`Lambda`表达式，也可以使用`Lambda`表达式来创建，这种方法相较匿名内部类更加简洁，也更推荐这种做法。实例代码如下：

```java
/**
 * 测试创建函数式接口对象
 * @author colorful@TaleLin
 */
public class Test {

    public static void main(String[] args) {
        // 使用 Lambda 表达式方式创建函数式接口
        FunctionalInterfaceDemo functionalInterfaceDemo = () -> System.out.println("Lambda 表达式方式创建函数式接口");
        functionalInterfaceDemo.run();
    }

}
```

运行结果：

```java
Lambda 表达式方式创建函数式接口
```

当然，还有一种更笨的方法，写一个接口的实现类，通过实例化实现类来创建对象。由于比较简单，而且不符合我们学习函数式接口的初衷，这里就不再做实例演示了。

## 5. 内置的函数式接口介绍

通过上面一系列介绍和演示，相信对于函数式接口的概念和使用，你已经烂熟于心了。但是只知道这些还不够用，下面的内容才是本小节的重点，Java 中内置了丰富的函数式接口，位于`java.util.function`包下，学习这些函数式接口有助于我们理解 Java 函数式接口的真正用途和意义。

Java 内置了 4 个核心函数式接口：

1. `Comsumer<T>`**消费型接口**： 表示接受单个输入参数但不返回结果的操作，包含方法：`void accept(T t)`，可以理解为消费者，只消费（接收单个参数）、不返回（返回为 `void`）；
2. `Supplier<T>`**供给型接口**：表示结果的供给者，包含方法`T get()`，可以理解为供给者，只提供（返回`T`类型对象）、不消费（不接受参数）；
3. `Function<T, R>`**函数型接口**：表示接受一个`T`类型参数并返回`R`类型结果的对象，包含方法`R apply(T t)`；
4. `Predicate<T>`**断言型接口**：确定`T`类型的对象是否满足约束，并返回`boolean`值，包含方法`boolean test(T t)`。

我们在 Java 的 `api` 文档中可以看到有一些方法的形参，会出现上面几类接口，我们在实例化这些接口的时候，就可以使用`Lambda`表达式的方式来实例化。

我们下面看几个实例，消费型接口使用实例：

```java
import java.util.function.Consumer;

/**
 * Java 内置 4 大核心 h 函数式接口 —— 消费型接口
 * Consumer<T> void accept(T t)
 * @author colorful@TaleLin
 */
public class FunctionalInterfaceDemo1 {

    public static void main(String[] args) {
        Consumer<String> consumer = s -> System.out.println(s);
        consumer.accept("只消费，不返回");
    }

}
```

运行结果：

```java
只消费，不返回
```

供给型接口使用实例：

```java
import java.util.function.Consumer;
import java.util.function.Supplier;

/**
 * Java 内置 4 大核心 h 函数式接口 —— 供给型接口
 * Supplier<T> T get()
 * @author colorful@TaleLin
 */
public class FunctionalInterfaceDemo2 {

    public static void main(String[] args) {
        Supplier<String> supplier = () -> "只返回，不消费";
        String s = supplier.get();
        System.out.println(s);
    }

}
```

运行结果：

```java
只返回，不消费
```

下面我们使用断言型接口，来实现一个根据给定的规则，来过滤字符串列表的方法，实例如下：

```java
import java.util.ArrayList;
import java.util.List;
import java.util.function.Predicate;

/**
 * Java 内置 4 大核心函数式接口 —— 断言型接口
 * Predicate<T> boolean test(T t)
 * @author colorful@TaleLin
 */
public class FunctionalInterfaceDemo3 {

    /**
     * 根据 Predicate 断言的结果，过滤 list 中的字符串
     * @param list 待过滤字符串
     * @param predicate 提供规则的接口实例
     * @return 过滤后的列表
     */
    public static List<String> filterStringList(List<String> list, Predicate<String> predicate) {
        // 过滤后的字符串列表
        ArrayList<String> arrayList = new ArrayList<>();
        for (String string: list) {
            if (predicate.test(string)) {
                // 如果 test 是 true，则将元素加入到过滤后的列表中
                arrayList.add(string);
            }
        }
        return arrayList;
    }

    public static void main(String[] args) {
        ArrayList<String> arrayList = new ArrayList<>();
        arrayList.add("Java");
        arrayList.add("PHP");
        arrayList.add("Python");
        arrayList.add("JavaScript");
        System.out.println("过滤前：");
        System.out.println(arrayList);

        List<String> filterResult = filterStringList(arrayList, new Predicate<String>() {
            @Override
            public boolean test(String s) {
                // 返回字符串中是否包含 P
                return s.contains("P");
            }
        });
        System.out.println("过滤后：");
        System.out.println(filterResult);
    }

}
```

运行结果：

```java
过滤前：
[Java, PHP, Python, JavaScript]
过滤后：
[PHP, Python]
```

当然，我们学习了`Lambda`表达式，在`main()`方法中就可以不再使用匿名内部类了，改写`main()`方法中调用`filterStringList()`方法的代码：

```java
List<String> filterResult = filterStringList(arrayList, s -> s.contains("P"));
```

上面的实例代码可能有些难以理解，跟着我的节奏来解读一下：

* 先定义一个方法`List<String> filterStringList(List<String> list, Predicate<String> predicate)`，此方法用于根据指定的规则过滤字符串列表，接收的第一个参数为待过滤列表，第二个参数是一个函数式接口类型的规则，注意，这个参数就是规则的制定者；
* 再看`filterStringList()`方法的方法体，方法体内部对待过滤列表进行了遍历，会调用`Predicate<T>`接口下的`boolean test(T t)`方法，判断每一个字符串是否符合规则，符合规则就追加到新的列表中，最终返回一个新的过滤后的列表；
* 在`main()`方法中，我们调用了上面定义的`filterStringList()`方法，第一个参数就是待过滤列表，这里的第二个参数，是我们创建的一个断言型接口的对象，其重写的`test(String s)`方法就是过滤规则关键所在，方法体就是判断`s`字符串是否包含`P`字符，并一个 boolean 类型的结果；
* 理解了第二个参数通过匿名内部类创建对象的方式，再改写成通过`Lambda`表达式的方式创建对象，就不难理解了。

上面我们介绍了核心的内置函数式接口，理解了这些接口的使用，其他接口就不难理解了。可翻阅 [官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/function/package-summary.html) 来查看更多。

## 6. 小结

通过本小节的学习，我们知道了函数式接口就是只有一个抽象方法的接口，要使用`Lambda`表达式，就必须依赖函数式接口；自定义函数接口建议使用`@FunctionalInterface`注解来进行标注，当然如果通过 Java 内置的函数式接口就可以满足我们的需求，就不需要我们自己自定义函数式接口了。本小节的最后，我们通过一个较为复杂的函数式接口实例，实现了一个过滤字符串列表的方法，如果还是不能完全理解，建议同学下面多加练习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
