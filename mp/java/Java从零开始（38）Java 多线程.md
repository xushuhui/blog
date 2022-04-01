---
title: Java 从零开始（38）Java 多线程
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415744542
---
# Java 多线程

本小节我们将学习 Java 多线程，通过本小节的学习，你将了解到什么是线程，如何创建线程，创建线程有哪几种方式，线程的状态、生命周期等内容。掌握多线程的代码编写，并理解线程生命周期等内容是本小节学习的重点。

## 1. 什么是线程

要了解什么是线程，就要先了解进程的概念。

进程，是指计算机中已运行的程序，它是一个动态执行的过程。假设我们电脑上同时运行了浏览器、QQ 以及代码编辑器三个软件，这三个软件之所以同时运行，就是进程所起的作用。

线程是操作系统能够进行运算调度的**最小单位**。大部分情况下，它被包含在进程之中，是进程中的实际运作单位。也就是说**一个进程可以包含多个线程，** 因此线程也被称为轻量级进程。

如果你还是对于进程和线程的概念有所困惑，推荐一篇比较优秀的 [文章](https://www.ruanyifeng.com/blog/2013/04/processes_and_threads.html)，有助于帮助你理解进程和线程的概念。

## 2. 创建线程

在 Java 中，创建线程有以下 3 种方式：

1. 继承 `Thread` 类，重写 `run()` 方法，该方法代表线程要执行的任务；
2. 实现 `Runnable` 接口，实现 `run()` 方法，该方法代表线程要执行的任务；
3. 实现 `Callable` 接口，实现 `call()` 方法，`call()` 方法作为线程的执行体，具有返回值，并且可以对异常进行声明和抛出。

下面我们分别来看下这 3 种方法的具体实现。

### 2.1 Thread 类

`Thread` 类是一个线程类，位于 `java.lang` 包下。

#### 2.1.1 构造方法

`Thread` 类的常用构造方法如下：

* `Thread()`：创建一个线程对象；
* `Thread(String name)`：创建一个指定名称的线程对象；
* `Thread(Runnable target)`：创建一个基于 `Runnable` 接口实现类的线程对象；
* `Thread(Runnable target, String name)`：创建一个基于 `Runnable` 接口实现类，并具有指定名称的线程对象。

#### 2.1.2 常用方法

`void run()`：线程相关的代码写在该方法中，一般需要重写；

`void start()`：启动当前线程；

`static void sleep(long m)`：使当前线程休眠 `m` 毫秒；

`void join()`：优先执行调用 `join()` 方法的线程。

> **Tips：**`run()` 方法是一个非常重要的方法，它是用于编写**线程执行体**的方法，不同线程之间的一个最主要区别就是 `run()` 方法中的代码是不同的。

可翻阅 [官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/lang/Thread.html) 以查看更多 API。

#### 2.1.3 实例

通过继承 `Thread` 类创建线程可分为以下 3 步：

1. 定义 `Thread` 类的子类，并重写该类的 `run()` 方法。`run()` 方法的方法体就代表了线程要完成的任务；
2. 创建 `Thread` 子类的实例，即创建线程对象；
3. 调用线程对象的 `start` 方法来启动该线程。

具体实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class ThreadDemo1 extends Thread {

    /**
     * 重写 Thread() 的方法
     */
    @Override
    public void run() {
        System.out.println("这里是线程体");
        // 当前打印线程的名称
        System.out.println(getName());
    }

    public static void main(String[] args) {
        // 实例化 ThreadDemo1 对象
        ThreadDemo1 threadDemo1 = new ThreadDemo1();
        // 调用 start() 方法，以启动线程
        threadDemo1.start();
    }

}
```

运行结果：

```java
这里是线程体
Thread-0
```

小伙伴们可能会有疑问，上面这样的代码，和普通的类实例化以及方法调用有什么区别的，下面我们来看一个稍微复杂些的实例：

```java
/**
 * @author colorful@TaleLin
 */
public class ThreadDemo2 {

    /**
     * 静态内部类
     */
    static class MyThread extends Thread {

        private int i = 3;

        MyThread(String name) {
            super(name);
        }

        @Override
        public void run() {
            while (i > 0) {
                System.out.println(getName() + " i = " + i);
                i--;
            }
        }

    }

    public static void main(String[] args) {
        // 创建两个线程对象
        MyThread thread1 = new MyThread("线程 1");
        MyThread thread2 = new MyThread("线程 2");
        // 启动线程
        thread1.start();
        thread2.start();
    }

}
```

运行结果：

```java
线程 2 i = 3
线程 1 i = 3
线程 1 i = 2
线程 2 i = 2
线程 1 i = 1
线程 2 i = 1
```

代码中我们是先启动了线程 1，再启动了线程 2 的，观察运行结果，线程并不是按照我们所预想的顺序执行的。这里就要划重点了，**不同线程，执行顺序是随机的**。如果你再执行几次代码，可以观察到每次的运行结果都可能不同：

![](https://xushuhui.gitee.io/image/imooc/5eeb28290a140dcd23221470.jpg)

### 2.2 Runnable 接口

#### 2.2.1 为什么需要 `Runnable` 接口

通过实现 `Runnable` 接口的方案来创建线程，要优于继承 `Thread` 类的方案，主要有以下原因：

1. Java 不支持多继承，所有的类都只允许继承一个父类，但可以实现多个接口。如果继承了 `Thread` 类就无法继承其它类，这不利于扩展；
2. 继承 `Thread` 类通常只重写 `run()` 方法，其他方法一般不会重写。继承整个 `Thread` 类成本过高，开销过大。

#### 2.2.2 实例

通过实现 `Runnable` 接口创建线程的步骤如下：

1. 定义 `Runnable` 接口的实现类，并实现该接口的 `run()` 方法。这个 `run()` 方法的方法体同样是该线程的线程执行体；
2. 创建 `Runnable` 实现类的实例，并以此实例作为 `Thread` 的 `target` 来创建 `Thread` 对象，该 `Thread` 对象才是真正的线程对象；
3. 调用线程对象的 `start` 方法来启动该线程。

具体实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class RunnableDemo1 implements Runnable {

    private int i = 5;

    @Override
    public void run() {
        while (i > 0) {
            System.out.println(Thread.currentThread().getName() + " i = " + i);
            i--;
        }
    }

    public static void main(String[] args) {
        // 创建两个实现 Runnable 实现类的实例
        RunnableDemo1 runnableDemo1 = new RunnableDemo1();
        RunnableDemo1 runnableDemo2 = new RunnableDemo1();
        // 创建两个线程对象
        Thread thread1 = new Thread(runnableDemo1, "线程 1");
        Thread thread2 = new Thread(runnableDemo2, "线程 2");
        // 启动线程
        thread1.start();
        thread2.start();
    }

}
```

运行结果：

```java
线程 1 i = 5
线程 1 i = 4
线程 1 i = 3
线程 1 i = 2
线程 2 i = 5
线程 1 i = 1
线程 2 i = 4
线程 2 i = 3
线程 2 i = 2
线程 2 i = 1
```

### 2.3 Callable 接口

#### 2.3.1 为什么需要 `Callable` 接口

**继承 Thread 类和实现 Runnable 接口这两种创建线程的方式都没有返回值**。所以，线程执行完毕后，无法得到执行结果。为了解决这个问题，Java 5 后，提供了 `Callable` 接口和 `Future` 接口，通过它们，可以在线程执行结束后，返回执行结果。

#### 2.3.2 实例

通过实现 `Callable` 接口创建线程步骤如下：

1. 创建 `Callable` 接口的实现类，并实现 `call()` 方法。这个 `call()` 方法将作为线程执行体，并且有返回值；
2. 创建 `Callable` 实现类的实例，使用 `FutureTask` 类来包装 `Callable` 对象，这个 `FutureTask` 对象封装了该 `Callable` 对象的 `call()` 方法的返回值；
3. 使用 `FutureTask` 对象作为 `Thread` 对象的 target 创建并启动新线程；
4. 调用 `FutureTask` 对象的 `get()` 方法来获得线程执行结束后的返回值。

具体实例如下：

```java
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

/**
 * @author colorful@TaleLin
 */
public class CallableDemo1 {

    static class MyThread implements Callable<String> {

        @Override
        public String call() { // 方法返回值类型是一个泛型，在上面 Callable<String> 处定义
            return "我是线程中返回的字符串";
        }

    }

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        // 常见实现类的实例
        Callable<String> callable = new MyThread();
        // 使用 FutureTask 类来包装 Callable 对象
        FutureTask<String> futureTask = new FutureTask<>(callable);
        // 创建 Thread 对象
        Thread thread = new Thread(futureTask);
        // 启动线程
        thread.start();
        // 调用 FutureTask 对象的 get() 方法来获得线程执行结束后的返回值
        String s = futureTask.get();
        System.out.println(s);
    }

}
```

运行结果：

```java
我是线程中返回的字符串
```

## 3. 线程休眠

在前面介绍 `Thread` 类的常用方法时，我们介绍了 `sleep()` 静态方法，该方法可以使当前执行的线程睡眠（暂时停止执行）指定的毫秒数。

线程休眠的实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class SleepDemo implements Runnable {

    @Override
    public void run() {
        for (int i = 1; i <= 5; i ++) {
            // 打印语句
            System.out.println(Thread.currentThread().getName() + ": 执行第" + i + "次");
            try {
                // 使当前线程休眠
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        // 实例化 Runnable 的实现类
        SleepDemo sleepDemo = new SleepDemo();
        // 实例化线程对象
        Thread thread = new Thread(sleepDemo);
        // 启动线程
        thread.start();
    }

}
```

运行结果：

```java
Thread-0: 执行第 1 次
Thread-0: 执行第 2 次
Thread-0: 执行第 3 次
Thread-0: 执行第 4 次
Thread-0: 执行第 5 次
```

![](https://xushuhui.gitee.io/image/imooc/5eeb28ec0ab9799823221470.jpg)

## 4. 线程的状态和生命周期

`java.lang.Thread.Starte` 枚举类中定义了 6 种不同的线程状态：

1. `NEW`：新建状态，尚未启动的线程处于此状态；
2. `RUNNABLE`：可运行状态，Java 虚拟机中执行的线程处于此状态；
3. `BLOCK`：阻塞状态，等待监视器锁定而被阻塞的线程处于此状态；
4. `WAITING`：等待状态，无限期等待另一线程执行特定操作的线程处于此状态；
5. `TIME_WAITING`：定时等待状态，在指定等待时间内等待另一线程执行操作的线程处于此状态；
6. `TERMINATED`：结束状态，已退出的线程处于此状态。

值得注意的是，一个线程在给定的时间点只能处于一种状态。这些状态是不反映任何操作系统线程状态的虚拟机状态。

线程的生命周期，实际上就是上述 6 个线程状态的转换过程。下图展示了一个完整的生命周期：

![](https://xushuhui.gitee.io/image/imooc/5eeb22e909ca713014861007.jpg)

## 5. 小结

通过本小节的学习，我们知道了线程是操作系统能够进行运算调度的**最小单位**。线程也被称为轻量级进程。在 Java 中，可以以 3 种方式创建线程，分别是继承 `Thread` 类、实现 `Runnable` 接口以及实现 `Callable` 接口。可以使用静态方法 `sleep()` 让线程休眠。线程状态有 6 种，也有资料上说线程有 5 种，这部分内容我们按照 Java 源码中的定义 6 种来记忆即可。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
