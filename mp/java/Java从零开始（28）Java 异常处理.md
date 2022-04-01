---
title: Java 从零开始（28）Java 异常处理
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/412912996
---
# Java 异常处理

`Java` 的异常处理是 Java 语言的一大重要特性，也是提高代码健壮性的最强大方法之一。当我们编写了错误的代码时，编译器在编译期间可能会抛出异常，有时候即使编译正常，在运行代码的时候也可能会抛出异常。本小节我们将介绍**什么是异常**、Java 中**异常类的架构**、**如何进行异常处理**、**如何自定义异常**、**什么是异常链**、**如何使用异常链**等内容。

## 1. 什么是异常

异常就是程序上的错误，我们在编写程序的时候经常会产生错误，这些错误划分为**编译期间的错误**和**运行期间的错误。**

下面我们来看几个常见的异常案例。

如果语句漏写分号，程序在编译期间就会抛出异常，实例如下：

```java
public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello World!")
    }
}
```

运行结果：

```java
$ javac Hello.java
Hello.java:3: 错误：需要';'
        System.out.println("Hello World!")
                                          ^
1 个错误
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec62cd90afc2fc815621358.jpg)

由于代码的第 3 行语句漏写了分号，Java 编译器给出了明确的提示。

`static` 关键字写成了 `statci`，实例如下：

```java
Hello.java:2: 错误：需要<标识符>
    public statci void main(String[] args) {
                 ^
1 个错误
```

当数组下标越界，程序在编译阶段不会发生错误，但在运行时会抛出异常。实例如下：

```java
public class ArrayOutOfIndex {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        System.out.println(arr[3]);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.ArrayIndexOutOfBoundsException: Index 3 out of bounds for length 3
	at ArrayOutOfIndex.main(ArrayOutOfIndex.java:4)
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec62e090a36dc6320761358.jpg)

## 2. Java 异常类架构

在 Java 中，通过 `Throwable` 及其子类来描述各种不同类型的异常。如下是 Java 异常类的架构图（不是全部，只展示部分类）：

![](https://xushuhui.gitee.io/image/imooc/5ec46e9f098c833310400588.jpg)

### 2.1 Throwable 类

`Throwable` 位于 `java.lang` 包下，它是 Java 语言中所有错误（`Error`）和异常（`Exception`）的父类。

`Throwable` 包含了其线程创建时线程执行堆栈的快照，它提供了 `printStackTrace()` 等接口用于获取堆栈跟踪数据等信息。

主要方法：

* `fillInStackTrace`： 用当前的调用栈层次填充 `Throwable` 对象栈层次，添加到栈层次任何先前信息中；

* `getMessage`：返回关于发生的异常的详细信息。这个消息在 `Throwable` 类的构造函数中初始化了；

* `getCause`：返回一个 `Throwable` 对象代表异常原因；

* `getStackTrace`：返回一个包含堆栈层次的数组。下标为 0 的元素代表栈顶，最后一个元素代表方法调用堆栈的栈底；

* `printStackTrace`：打印 `toString()` 结果和栈层次到 `System.err`，即错误输出流。

### 2.2 Error 类

`Error` 是 `Throwable` 的一个直接子类，它可以指示合理的应用程序不应该尝试捕获的**严重问题**。这些错误在应用程序的控制和处理能力之外，编译器不会检查 `Error`，对于设计合理的应用程序来说，即使发生了错误，本质上也无法通过异常处理来解决其所引起的异常状况。

常见 `Error`：

* `AssertionError`：断言错误；

* `VirtualMachineError`：虚拟机错误；

* `UnsupportedClassVersionError`：Java 类版本错误；

* `OutOfMemoryError` ：内存溢出错误。

### 2.3 Exception 类

`Exception` 是 `Throwable` 的一个直接子类。它指示合理的应用程序可能希望捕获的条件。

`Exception` 又包括 `Unchecked Exception`（非检查异常）和 `Checked Exception`（检查异常）两大类别。

#### 2.3.1 Unchecked Exception （非检查异常）

`Unchecked Exception` 是编译器不要求强制处理的异常，包含 `RuntimeException` 以及它的相关子类。我们编写代码时即使不去处理此类异常，程序还是会编译通过。

常见非检查异常：

* `NullPointerException`：空指针异常；

* `ArithmeticException`：算数异常；

* `ArrayIndexOutOfBoundsException`：数组下标越界异常；

* `ClassCastException`：类型转换异常。

#### 2.3.2 Checked Exception（检查异常）

`Checked Exception` 是编译器要求必须处理的异常，除了 `RuntimeException` 以及它的子类，都是 `Checked Exception` 异常。我们在程序编写时就必须处理此类异常，否则程序无法编译通过。

常见检查异常：

* `IOException`：IO 异常

* `SQLException`：SQL 异常

## 3. 如何进行异常处理

在 Java 语言中，异常处理机制可以分为两部分：

1. **抛出异常**：当一个方法发生错误时，会创建一个异常对象，并交给运行时系统处理；

2. **捕获异常**：在方法抛出异常之后，运行时系统将转为寻找合适的异常处理器。

Java 通过 5 个关键字来实现异常处理，分别是：`throw`、`throws`、`try`、`catch`、`finally`。

异常总是先抛出，后捕获的。下面我们将围绕着 5 个关键字来详细讲解**如何抛出异常**以及**如何捕获异常**。

## 4. 抛出异常

### 4.1 实例

我们先来看一个除零异常的实例代码：

```java
public class ExceptionDemo1 {
    // 打印 a / b 的结果
    public static void divide(int a, int b) {
        System.out.println(a / b);
    }

    public static void main(String[] args) {
        // 调用 divide() 方法
        divide(2, 0);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.ArithmeticException: / by zero
	at ExceptionDemo1.divide(ExceptionDemo1.java:4)
	at ExceptionDemo1.main(ExceptionDemo1.java:9)
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec62f610a8561c914921382.jpg)

我们知道 `0` 是不能用作除数的，由于 `divide()` 方法中除数 `b` 为 `0`，所以代码将停止执行并显示了相关的异常信息，此信息为堆栈跟踪，上面的运行结果告诉我们：`main` 线程发生了类型为 `ArithmeticException` 的异常，显示消息为 `by zero`，并且提示了可能发生异常的方法和行号。

### 4.2 throw

上面的实例中，程序在运行时引发了错误，那么如何来显示抛出（创建）异常呢？

我们可以使用 `throw` 关键字来抛出异常，`throw` 关键字后面跟异常对象，改写上面的实例代码：

```java
public class ExceptionDemo2 {
    // 打印 a / b 的结果
    public static void divide(int a, int b) {
        if (b == 0) {
            // 抛出异常
            throw new ArithmeticException("除数不能为零");
        }
        System.out.println(a / b);
    }

    public static void main(String[] args) {
        // 调用 divide() 方法
        divide(2, 0);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.ArithmeticException: 除数不能为零
	at ExceptionDemo2.divide(ExceptionDemo2.java:5)
	at ExceptionDemo2.main(ExceptionDemo2.java:12)
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec62fec0a7a2e2714921382.jpg)

代码在运行时同样引发了错误，但显示消息为 “除数不能为零”。我们看到 `divide()` 方法中加入了条件判断，如果调用者将参数 `b` 设置为 `0` 时，会使用 `throw` 关键字来抛出异常，throw 后面跟了一个使用 `new` 关键字实例化的算数异常对象，并且将消息字符串作为参数传递给了算数异常的构造函数。

我们可以使用 `throw` 关键字抛出任何类型的 `Throwable` 对象，它会中断方法，`throw` 语句之后的所有内容都不会执行。除非已经处理抛出的异常。异常对象不是从方法中返回的，而是从方法中抛出的。

### 4.3 throws

可以通过 `throws` 关键字声明方法要抛出何种类型的异常。如果一个方法可能会出现异常，但是没有能力处理这种异常，可以在方法声明处使用 `throws` 关键字来声明要抛出的异常。例如，汽车在运行时可能会出现故障，汽车本身没办法处理这个故障，那就让开车的人来处理。

`throws` 用在方法定义时声明该方法要抛出的异常类型，如下是伪代码：

```java
public void demoMethod() throws Exception1, Exception2, ... ExceptionN {
    // 可能产生异常的代码
}
```

`throws` 后面跟的异常类型列表可以有一个也可以有多个，多个则以 `,` 分割。当方法产生异常列表中的异常时，将把异常抛向方法的调用方，由调用方处理。

throws 有如下使用规则：

1. 如果方法中全部是非检查异常（即 `Error`、`RuntimeException` 以及的子类），那么可以不使用 `throws` 关键字来声明要抛出的异常，编译器能够通过编译，但在运行时会被系统抛出；
2. 如果方法中可能出现检查异常，就必须使用 `throws` 声明将其抛出或使用 `try catch` 捕获异常，否则将导致编译错误；
3. 当一个方法抛出了异常，那么该方法的调用者必须处理或者重新抛出该异常；
4. 当子类重写父类抛出异常的方法时，声明的异常必须是父类所声明异常的同类或子类。

## 5. 捕获异常

**使用 try 和 catch 关键字可以捕获异常**。try catch 代码块放在异常可能发生的地方。它的语法如下：

```java
try {
    // 可能会发生异常的代码块
} catch (Exception e1) {
    // 捕获并处理 try 抛出的异常类型 Exception
} catch (Exception2 e2) {
    // 捕获并处理 try 抛出的异常类型 Exception2
} finally {
    // 无论是否发生异常，都将执行的代码块
}
```

我们来看一下上面语法中的 3 种语句块：

1. **`try` 语句块**：用于监听异常，当发生异常时，异常就会被抛出；
2. **`catch` 语句块**：`catch` 语句包含要捕获的异常类型的声明，当 `try` 语句块发生异常时，`catch` 语句块就会被检查。当 `catch` 块尝试捕获异常时，是按照 `catch` 块的声明顺序从上往下寻找的，一旦匹配，就不会再向下执行。因此，如果同一个 `try` 块下的多个 `catch` 异常类型有父子关系，应该将子类异常放在前面，父类异常放在后面；
3. **`finally` 语句块**：无论是否发生异常，都会执行 `finally` 语句块。`finally` 常用于这样的场景：由于 `finally` 语句块总是会被执行，所以那些在 `try` 代码块中打开的，并且必须回收的物理资源（如数据库连接、网络连接和文件），一般会放在 `finally` 语句块中释放资源。

`try` 语句块后可以接零个或多个 `catch` 语句块，如果没有 `catch` 块，则必须跟一个 `finally` 语句块。简单来说，`try` 不允许单独使用，必须和 `catch` 或 `finally` 组合使用，`catch` 和 `finally` 也不能单独使用。

实例如下：

```java
public class ExceptionDemo3 {
    // 打印 a / b 的结果
    public static void divide(int a, int b) {
        System.out.println(a / b);
    }

    public static void main(String[] args) {
        try {
            // try 语句块
            // 调用 divide() 方法
            divide(2, 0);
        } catch (ArithmeticException e) {
            // catch 语句块
            System.out.println("catch: 发生了算数异常：" + e);
        } finally {
            // finally 语句块
            System.out.println("finally: 无论是否发生异常，都会执行");
        }
    }
}
```

运行结果：

```java
catch: 发生了算数异常：java.lang.ArithmeticException: / by zero
finally: 无论是否发生异常，都会执行
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec6301f0a4b32d514921382.jpg)

`divide()` 方法中除数 `b` 为 `0`，会发生除零异常，我们在方法调用处使用了 `try` 语句块对异常进行捕获；如果捕获到了异常， `catch` 语句块会对 `ArithmeticException` 类型的异常进行处理，此处打印了一行自定义的提示语句；最后的 `finally` 语句块，无论发生异常与否，总会执行。

Java 7 以后，`catch` 多种异常时，也可以像下面这样简化代码：

```java
try {
    // 可能会发生异常的代码块
} catch (Exception | Exception2 e) {
    // 捕获并处理 try 抛出的异常类型
} finally {
    // 无论是否发生异常，都将执行的代码块
}
```

## 6. 自定义异常

自定义异常，就是定义一个类，去继承 `Throwable` 类或者它的子类。

Java 内置了丰富的异常类，通常使用这些内置异常类，就可以描述我们在编码时出现的大部分异常情况。一旦内置异常无法满足我们的业务要求，就可以通过自定义异常描述特定业务产生的异常类型。

实例：

```java
public class ExceptionDemo4 {

    static class MyCustomException extends RuntimeException {
        /**
         * 无参构造方法
         */
        public MyCustomException() {
            super("我的自定义异常");
        }
    }

    public static void main(String[] args) {
      	// 直接抛出异常
        throw new MyCustomException();
    }
}
```

运行结果：

```java
Exception in thread "main" ExceptionDemo4$MyCustomException: 我的自定义异常
	at ExceptionDemo4.main(ExceptionDemo4.java:13)
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec63e250a93f84b14921382.jpg)

在代码中写了一个自定义异常 `MyCustomException`，继承自 `RuntimeException`，它是一个静态内部类，这样在主方法中就可以直接抛出这个异常类了。当然，也可以使用 `catch` 来捕获此类型异常。

## 7. 异常链

异常链是以一个异常对象为参数构造新的异常对象，新的异常对象将包含先前异常的信息。简单来说，就是将异常信息从底层传递给上层，逐层抛出，我们来看一个实例：

```java
public class ExceptionDemo5 {

    /**
     * 第一个自定义的静态内部异常类
     */
    static class FirstCustomException extends Exception {

        // 无参构造方法
        public FirstCustomException() {
            super("第一个异常");
        }
    }

    /**
     * 第二个自定义的静态内部异常类
     */
    static class SecondCustomException extends Exception {

        public SecondCustomException() {
            super("第二个异常");
        }
    }

    /**
     * 第三个自定义的静态内部异常类
     */
    static class ThirdCustomException extends Exception {

        public ThirdCustomException() {
            super("第三个异常");
        }
    }

    /**
     * 测试异常链静态方法 1，直接抛出第一个自定义的静态内部异常类
     * @throws FirstCustomException
     */
    public static void f1() throws FirstCustomException {
        throw new FirstCustomException();
    }

    /**
     * 测试异常链静态方法 2，调用 f1() 方法，并抛出第二个自定义的静态内部异常类
     * @throws SecondCustomException
     */
    public static void f2() throws SecondCustomException {
        try {
            f1();
        } catch (FirstCustomException e) {
            throw new SecondCustomException();
        }
    }

    /**
     * 测试异常链静态方法 3，调用 f2() 方法， 并抛出第三个自定义的静态内部异常类
     * @throws ThirdCustomException
     */
    public static void f3() throws ThirdCustomException {
        try {
            f2();
        } catch (SecondCustomException e) {
            throw new ThirdCustomException();
        }
    }

    public static void main(String[] args) throws ThirdCustomException {
        // 调用静态方法 f3()
        f3();
    }
}
```

运行结果：

```java
Exception in thread "main" ExceptionDemo5$ThirdCustomException: 第三个异常
	at ExceptionDemo5.f3(ExceptionDemo5.java:46)
	at ExceptionDemo5.main(ExceptionDemo5.java:51)
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec640430afdc24e14921382.jpg)

通过运行结果，我们只获取到了静态方法 `f3()` 所抛出的异常堆栈信息，前面代码所抛出的异常并没有被显示。

我们改写上面的代码，让异常信息以链条的方式 “连接” 起来。可以通过改写自定义异常的构造方法，来获取到之前异常的信息。实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class ExceptionDemo6 {

    /**
     * 第一个自定义的静态内部异常类
     */
    static class FirstCustomException extends Exception {

        // 无参构造方法
        public FirstCustomException() {
            super("第一个异常");
        }

    }

    /**
     * 第二个自定义的静态内部异常类
     */
    static class SecondCustomException extends Exception {

        /**
         * 通过构造方法获取之前异常的信息
         * @param cause 捕获到的异常对象
         */
        public SecondCustomException(Throwable cause) {
            super("第二个异常", cause);
        }
    }

    /**
     * 第三个自定义的静态内部异常类
     */
    static class ThirdCustomException extends Exception {

        /**
         * 通过构造方法获取之前异常的信息
         * @param cause 捕获到的异常对象
         */
        public ThirdCustomException(Throwable cause) {
            super("第三个异常", cause);
        }
    }

    /**
     * 测试异常链静态方法 1，直接抛出第一个自定义的静态内部异常类
     * @throws FirstCustomException
     */
    public static void f1() throws FirstCustomException {
        throw new FirstCustomException();
    }

    /**
     * 测试异常链静态方法 2，调用 f1() 方法，并抛出第二个自定义的静态内部异常类
     * @throws SecondCustomException
     */
    public static void f2() throws SecondCustomException {
        try {
            f1();
        } catch (FirstCustomException e) {
            throw new SecondCustomException(e);
        }
    }

    /**
     * 测试异常链静态方法 3，调用 f2() 方法， 并抛出第三个自定义的静态内部异常类
     * @throws ThirdCustomException
     */
    public static void f3() throws ThirdCustomException {
        try {
            f2();
        } catch (SecondCustomException e) {
            throw new ThirdCustomException(e);
        }
    }

    public static void main(String[] args) throws ThirdCustomException {
        // 调用静态方法 f3()
        f3();
    }
}
```

运行结果：

```java
Exception in thread "main" ExceptionDemo6$ThirdCustomException: 第三个异常
	at ExceptionDemo6.f3(ExceptionDemo6.java:74)
	at ExceptionDemo6.main(ExceptionDemo6.java:80)
Caused by: ExceptionDemo6$SecondCustomException: 第二个异常
	at ExceptionDemo6.f2(ExceptionDemo6.java:62)
	at ExceptionDemo6.f3(ExceptionDemo6.java:72)
	... 1 more
Caused by: ExceptionDemo6$FirstCustomException: 第一个异常
	at ExceptionDemo6.f1(ExceptionDemo6.java:51)
	at ExceptionDemo6.f2(ExceptionDemo6.java:60)
	... 2 more
```

运行过程：

![](https://xushuhui.gitee.io/image/imooc/5ec64df40aada54714921382.jpg)

通过运行结果，我们看到，异常发生的整个过程都打印到了屏幕上，这就是一个异常链。

## 8. 小结

通过本小节的学习，我们知道了异常就是程序上的错误，良好的异常处理可以提高代码的健壮性。Java 语言中所有错误（`Error`）和异常（`Exception`）的父类都是 `Throwable`。`Error` 和 `Exception` 是 `Throwable` 的直接子类，我们通常说的异常处理实际上就是处理 `Exception` 及其子类，异常又分为**检查型异常**和**非检查型异常**。通过抛出异常和捕获异常来实现异常处理。我们亦可以通过继承 `Throwable` 类或者它的子类来自定义异常类。通过构造方法获取之前异常的信息可以实现异常链。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
