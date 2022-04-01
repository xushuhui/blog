---
title: Java从零开始（87）多线程 join 方法
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 多线程 join 方法


## 1. 前言

本节对 join 方法进行深入的剖析，主要内容点如下：

* 了解 join 方法的作用，初步的理解 join 方法的使用带来的效果是学习本节内容的基础；
* 了解 join 方法异常处理，我们在使用 join 方法是，需要对 join 方法进行有效的异常处理；
* 通过匿名内部类创建 Thread 是我们本节代码示例所使用的的方式，对这种方式的掌握在后续工作中非常重要；
* 掌握 join 方法如何使用，这是本节的重点内容，也是本节的核心内容；
* 掌握带参数的 join 方法使用，在后续开发过程中，如果对接口的最大返回时间有要求的话，某些情况下会用到带参数的 join 方法，次重点内容。

## 2. join 方法的作用

**方法定义**：多线程环境下，如果需要确保某一线程执行完毕后才可继续执行后续的代码，就可以通过使用 join 方法完成这一需求设计。

在项目实践中经常会遇到一个场景，就是需要等待某几件事情完成后主线程才能继续往下执行， 比如多个线程加载资源， 需要等待多个线程全部加载完毕再汇总处理。

Thread 类中有一个 join 方法就可以做这个事情，join 方法是 Thread 类直接提供的。join 是无参且返回值为 void 的方法。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz56nt2j60jg05b3zq02)

如上图所示，假如有 3 个线程执行逻辑，线程 1 需要执行 5 秒钟，线程 2 需要执行 10 秒钟，线程 3 需要执行 8 秒钟。 如果我们的开发需求是：必须等 3 条线程都完成执行之后再进行后续的代码处理，这时候我们就需要使用到 join 方法。

**使用 join 方法后**：

* **第 5 秒钟**： 线程 1 执行完毕；线程 2 执行了一半； 线程 3 还差 3 秒执行完毕；
* **第 8 秒钟**：线程 1 等待了 3 秒； 线程 3 执行完毕； 线程 2 还差 2 秒执行完毕；
* **第 10 秒钟**： 线程 1 等待了 5 秒； 线程 3 等待了 2 秒；线程 2 执行完毕；
* **从线程 2 执行结束的那一刻**：三条线程同时进行后续代码的执行。

这就是 join 方法的作用与解释。

## 3. join 方法异常处理

join 方法是 Thread 类中的方法，为了了解该方法的异常处理，我们先来简要的看下 join 方法的 JDK 源码：

```java
public final void join() throws InterruptedException {
        join(0);
}
```

从源代码中我们可以看到， join 方法抛出了异常：

```java
throws InterruptedException
```

所以，我们在使用 join 方法的时候，需要对 join 方法的调用进行 try catch 处理或者从方法级别进行异常的抛出。

**try-catch 处理示例**：

```java
public class DemoTest implements Runnable{
    @Override
    public void run() {
        System.out.println("线程："+Thread.currentThread()+" 正在执行...");
    }
    public static void main(String[] args) {
        Thread t1 = new Thread(new DemoTest());
        t1. start();
        try {
            t1.join();
        } catch (InterruptedException e) {
            // 异常捕捉处理
        }
    }
}
```

**throws 异常处理示例**：

```java
public class DemoTest implements Runnable throws InterruptedException {
    @Override
    public void run() {...}
    public static void main(String[] args) {
        Thread t1 = new Thread(new DemoTest());
        t1. start();
        t1.join();
    }
}
```

## 4. join 方法如何使用

为了更好的了解 join 方法的使用，我们首先来设计一个使用的场景。

**场景设计**：

* **线程 1** ：执行时间 5 秒钟；
* **线程 2** ：执行时间 10 秒钟；
* **线程 3** ：执行 8 秒钟。

**需求**：我们需要等 3 个线程都执行完毕后，再进行后续代码的执行。3 个线程执行完毕后，请打印执行时间。

**期望结果**： 10 秒执行时间。

看到这个是不是似曾相识呢？ 这就是我们本节第 2 知识点所举出的示例，现在我们来进行代码实现和验证，体会 join 方法的使用。

**实例**：

```java
public class DemoTest{
    public static void main(String[] args) throws InterruptedException {
        Thread threadOne = new Thread(new Runnable() { //线程 1
            @Override
            public void run() {
                try {
                    Thread.sleep (5000 ); //线程 1 休眠 5 秒钟
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                System.out.println ("线程 1 休眠 5 秒钟，执行完毕。");
            }
        });

        Thread threadTwo = new Thread(new Runnable() { //线程 2
            ...
                    Thread.sleep (10000 ); //线程 2 修眠 10 秒钟
            ...
                System.out.println ("线程 2 修眠 10 秒钟，执行完毕。");
            }
        });

        Thread threadThree = new Thread(new Runnable() {//线程 3
            ...
                    Thread.sleep (8000 ); //线程 3 修眠 8 秒钟
            ...
                System.out.println ("线程 3 修眠 8 秒钟，执行完毕。");
            }
        });

        Long startTime = System.currentTimeMillis();
        threadOne. start();threadTwo. start();threadThree. start();
        System.out.println("等待三个线程全部执行完毕再继续向下执行,我要使用 join 方法了。");
        threadOne.join(); //线程 1 调用 join 方法
        threadTwo.join(); //线程 2 调用 join 方法
        threadThree.join(); //线程 3 调用 join 方法
        Long endTime = System.currentTimeMillis();
        System.out.println("三个线程都执行完毕了，共用时： "+ (endTime - startTime) + "毫秒");
    }
}
```

**执行结果验证**：

```java
等待三个线程全部执行完毕再继续向下执行,我要使用 join 方法了。
线程 1 休眠 5 秒钟，执行完毕。
线程 3 修眠 8 秒钟，执行完毕。
线程 2 修眠 10 秒钟，执行完毕。
三个线程都执行完毕了，共用时： 10002毫秒
```

从执行的结果来看，与我们对 join 方法的理解和分析完全相符，请同学也进行代码的编写和运行，加深学习印象。

## 5. 带参数的 join 方法使用

除了无参的 join 方法以外， Thread 类还提供了有参 join 方法如下：

```java
public final synchronized void join(long millis)
    throws InterruptedException
```

该方法的参数 long millis 代表的是毫秒时间。

**方法作用描述**：等待 millis 毫秒终止线程，假如这段时间内该线程还没执行完，也不会再继续等待。

结合上一个知识点的代码，我们都是调用的无参 join 方法，现在对上一个知识点代码进行如下调整：

```java
threadOne.join(); //线程 1 调用 join 方法
threadTwo.join(3000); //线程 2 调用 join 方法
threadThree.join(); //线程 3 调用 join 方法
```

从代码中我们看到，线程 2 使用 join 方法 3000 毫秒的等待时间，如果 3000 毫毛后，线程 2 还未执行完毕，那么主线程则放弃等待线程 2，只关心线程 1 和线程 3。

**我们来看下执行结果**：

```java
等待三个线程全部执行完毕再继续向下执行,我要使用 join 方法了。
线程 1 休眠 5 秒钟，执行完毕。
线程 3 修眠 8 秒钟，执行完毕。
三个线程都执行完毕了，共用时： 8000毫秒
线程 2 修眠 10 秒钟，执行完毕。
```

从执行结果来看， 总用时 8000 毫秒，因为线程 2 被放弃等待了，所以只考虑线程 1 和线程 3 的执行时间即可。

## 6. 小结

在实际的开发场景中，经常会设计到对 join 方法的使用，无参方法使用更加常见，了解 join 方法的使用非常重要。

本节重中之重，就是掌握 join 方法的使用，join 方法在后续的开发工作中非常关键，很多情况下都会有所使用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
