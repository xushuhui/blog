# Lambda 表达式的变量与作用域

本节我们将分析 Lambda 表达式的局部变量及其作用域进行分析，在这基础上我们会探讨其访问规则背后的原因。

在开始之前我们需要明确一句话：

> **引用值，而不是变量！**
>
> **引用值，而不是变量！**
>
> **引用值，而不是变量！**
>
> **重要的事情说三遍！！！**

## 1. 访问局部变量

Lambda 表达式不会从父类中继承任何变量名，也不会引入一个新的作用域。Lambda 表达式基于词法作用域，也就是说 Lambda 表达式函数体里面的变量和它外部环境的变量具有相同的语义。

访问局部变量要注意如下 3 点：

1. 可以直接在 Lambda 表达式中访问外层的局部变量；
2. 在 Lambda 表达式当中被引用的变量的值不可以被更改；
3. 在 Lambda 表达式当中不允许声明一个与局部变量同名的参数或者局部变量。

现在我们来仔细说明下这三点。

### 1.1 可以直接在 Lambda 表达式中访问外层的局部变量

**在 Lambda 表达式中可以直接访问外层的局部变量，但是这个局部变量必须是声明为 `final` 的。**

首先我们来看一个例子：

```java
import java.util.function.BinaryOperator;

public class LambdaTest1 {

    public static void main(String[] args) {
        final int delta = -1;
        BinaryOperator<Integer> add = (x, y) -> x+y+delta;
        Integer apply = add.apply(1, 2);//结果是2
        System.out.println(apply);
    }

}

```

在这个例子中， `delta` 是 Lambda 表达式中的外层局部变量，被声明为 `final`，我们的 Lambda 表达式是对两个输入参数 `x`,`y` 和外层局部变量 `delta` 进行求和。

如果这个变量是一个既成事实上的 final 变量的话，就可以不使用 `final` 关键字。所谓个既成事实上的 final 变量是指只能给变量赋值一次，在我们的第一个例子中，`delta` 只在初始化的时候被赋值，所以它是一个既成事实的 `final` 变量。

```java
import java.util.function.BinaryOperator;

public class LambdaTest2 {

    public static void main(String[] args) {
        int delta = -1;
        BinaryOperator<Integer> add = (x, y) -> x+y+delta;
        Integer apply = add.apply(1, 2);//结果是2
        System.out.println(apply);
    }

}
```

相较于第一个例子，我们删除了 `final` 关键字，程序没有任何问题。

### 1.2 在 Lambda 表达式当中被引用的变量的值不可以被更改

在 Lambda 表达式中试图修改局部变量是不允许的，那么我们在后面对 `delta` 赋值会怎么样呢？

```java
public static void main(String...s){
    int delta = -1;
    BinaryOperator<Integer> add = (x, y) -> x+y+ delta; //编译报错
    add.apply(1,2);
    delta = 2;
}
```

这个时候编译器会报错说：

```java
Variable used in lambda expression should be final or effectively final
```

### 1.3 在 Lambda 表达式当中不允许声明一个与局部变量同名的参数或者局部变量

```java
public static void main(String...s){
    int delta = -1;
    BinaryOperator<Integer> add = (delta, y) -> delta + y + delta; //编译报错
    add.apply(1,2);
}
```

我们将表达式的第一个参数的名称由 `x` 改为 `delta` 时，编译器会报错说：

```java
Variable 'delta' is already defined in the scope
```

## 2. 访问对象字段与静态变量

**Lambda 内部对于实例的字段和静态变量是即可读又可写的。**

```java
import java.util.function.BinaryOperator;

public class Test {
    public static int staticNum;
    private int num;

    public void doTest() {
        BinaryOperator<Integer> add1 = (x, y) -> {
            num = 3;
            staticNum = 4;
            return x + y + num + Test.staticNum;
        };
        Integer apply = add1.apply(1, 2);
        System.out.println(apply);
    }

    public static void main(String[] args) {
        new Test().doTest();
    }

}
```

这里我们在 `Test`类中，定义了一个静态变量 `staticNum` 和 私有变量 `num`。并在 Lambda 表达式 `add1` 中对其作了修改，没有任何问题。

## 3. 关于引用值，而不是变量

通过前面两节我们对于 Lambda 表达式的变量和作用域有了一个概念，总的来说就是：

> **Tips:** Lambda 表达式可以读写实例变量，只能读取局部变量。

有没有想过这是为什么呢？

* 实例变量和局部变量的实现不同：实例变量都存储在堆中，而局部变量则保存在栈上。如果在线程中要直接访问一个非`final`局部变量，可能线程执行时这个局部变量已经被销毁了。因此，Java 在访问自由局部变量时，实际上是在访问它的副本，而不是访问原始变量。如果局部变量仅仅赋值一次那就没有什么区别了——因此就没有这个限制（也就是既成事实的 `final`）。
* 这个局部变量的访问限制也是 Java 为了促使你从命令式编程模式转换到函数式编程模式，这样会很容易使用 Java 做到并行处理（关于命令式编程模式和函数式编程模式我们将在后续内容中做详细的解释）。

## 4. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a8ac809552cdc06870210.jpg)

本节我们主要介绍了 Lambda 表达式的变量作用域，主要有这么 3 点需要记住：

* 引用值，而不是变量；
* 可以读写实例变量；
* 只能读取局部变量。

最后我们对于 Lambda 表达式对于变量为什么会有这样的访问限制做了相应的分析。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
