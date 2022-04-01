---
title: Java从零开始（86）Java 多线程的创建
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 多线程的创建


## 1. 前言

本节内容重点需要掌握 Java 多线程的三种创建方式，具体内容如下：

* Java 线程类 Thread 继承结构，这是 JDK Thread 源码的类结构，是了解 Thread 类的第一步；
* 掌握多线程的三种创建方式，这是本节的重点内容。本节所有内容都是围绕改话题进行的讲解；
* 了解多线程不同创建方式的优劣，不同的创建方式有自己的优势和劣势，本节还会推荐同学使用第二种接口实现的创建方式；
* 掌握 Thread 类常用方法，这也是本节的重点内容，其常用方法使我们开发过程中经常涉及到的，必须要熟记于心；
* Thread 类编程实验实战，学习完多线程的创建方式，我们需要进行实战代码巩固本节的内容。

## 2. Thread 类结构介绍

**介绍**： 位于 java.lang 包下的 Thread 类是非常重要的线程类。学习 Thread 类的使用是学习多线程并发编程的基础。它实现了 Runnable 接口，其包集成结构如下图所示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz4w5vlj60jg070abb02)

## 3. 多线程的三种创建方式

Java 多线程有 3 种创建方式如下：

* **方式一**：继承 Thread 类的方式创建线程；
* **方式二**：实现 java.lang.Runnable 接口；
* **方式三**：实现 Callable 接口。

在接下来的内容中，会详细的对这 3 种创建方式进行详细的讲解。

## 4. 多线程实现之继承 Thread 类

**实现步骤**：

* **步骤 1**：继承 Thread 类 extends Thread；
* **步骤 2**：复写 run () 方法，run () 方法是线程具体逻辑的实现方法。

**实例**：

```java
/**
 * 方式一：继承Thread类的方式创建线程
 */
public class ThreadExtendTest extends Thread{ //步骤 1
    @Override
    public void run() { //步骤 2
	    //run方法内为具体的逻辑实现
        System.out.println("create thread by thread extend");
    }
    public static void main(String[] args) {
        new ThreadExtendTest(). start();
    }
}
```

## 5. 多线程实现之实现 Runnable 接口

> **Tips**：由于 Java 是面向接口编程，且可进行多接口实现，相比 Java 的单继承特性更加灵活，易于扩展，所以相比方式一，更推荐使用方式二进行线程的创建。

**实现步骤**：

* **步骤 1**：实现 Runnable 接口，implements Runnable；
* **步骤 2**：复写 run () 方法，run () 方法是线程具体逻辑的实现方法。

**实例**：

```java
/**
 * 方式二：实现java.lang.Runnable接口
 */
public class ThreadRunnableTest implements Runnable{//步骤 1
    @Override
    public void run() {//步骤 2
		//run方法内为具体的逻辑实现
        System.out.println("create thread by runnable implements");
    }
    public static void main(String[] args) {
        new Thread(new ThreadRunnableTest()). start();
   }
}
```

## 6. 多线程实现之实现 Callable 接口

> **Tips**：方式一与方式二的创建方式都是复写 run 方法，都是 void 形式的，没有返回值。但是对于方式三来说，实现 Callable 接口，能够有返回值类型。

**实现步骤**：

* **步骤 1**：实现 Callable 接口，implements Callable；
* **步骤 2**：复写 call () 方法，call () 方法是线程具体逻辑的实现方法。

**实例**：

```java
/**
 * 方式三：实现Callable接口
 */
public class ThreadCallableTest implements Callable<String> {//步骤 1
    @Override
    public String call() throws Exception { //步骤 2
	    //call 方法的返回值类型是 String
	    //call 方法是线程具体逻辑的实现方法
        return "create thread by implements Callable";
    }
    public static void main(String[] args) throws ExecutionException, InterruptedException{
        FutureTask<String> future1 = new FutureTask<String>(new ThreadCallableTest());
        Thread thread1 = new Thread(future1);
        thread1. start();
        System.out.println(future1.get());
    }
}
```

## 7. 匿名内部类创建 Thread

首先确认，这并不是线程创建的第四种方式，先来看如何创建。

**实例**：

```java
Thread t = new Thread(new Runnable() {
            @Override
            public void run() {
                System.out.println("通过匿名内部类创建Thread");
            }
        });
```

我们从代码中可以看出，还是进行了一个 Runnable 接口的使用，所以这并不是新的 Thread 创建方式，只不过是通过方式二实现的一个内部类创建。

> **Tips**： 在后续章节讲解 join 方法如何使用 的时候，我们会采用匿名内部类的方式进行多线程的创建。

## 8. Thread 类的常用方法介绍

|方法|作用|
|----|----|
|** start()**             | 启动当前的线程，调用当前线程的 run ()。                             |
|**run()**                | 通常需要重写 Thread 类中的此方法，将创建要执行的操作声明在此方法中。|
|**currentThread()**      | 静态方法，返回代码执行的线程。                                      |
|**getName()**            | 获取当前线程的名字。                                                |
|**setName()**            | 设置当前线程的名字。                                                |
|**sleep(long millitime)**| 让当前进程睡眠指定的毫秒数，在指定时间内，线程是阻塞状态。          |
|**isAlive()**            | 判断进程是否存活。                                                  |
|**wait()**               | 线程等待。                                                          |
|**notify()**             | 线程唤醒。                                                          |

## 9. Thread 编程测验实验

**实验目的**：对 Thread 的创建方式进行练习，巩固本节重点内容，并在练习的过程中，使用常用的 start 方法和 sleep 方法以及 线程的 setName 方法。

**实验步骤**：

* 使用 Runnable 接口创建两条线程 ：t1 和 t2；
* 请设置线程 t1 和 t2 的线程名称分别为 “ThreadOne” 和 “ThreadTwo”；
* 线程 t1 执行完 run () 方法后，线程睡眠 5 秒；
* 线程 t2 执行完 run () 方法后，线程睡眠 1 秒。

请先自行实现，并将结果与所提供的答案进行复核。

```java
public class ThreadTest implements Runnable{

    @Override
    public void run() {
        System.out.println("线程："+Thread.currentThread()+" 正在执行...");
    }
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new ThreadTest());
        t1.setName("ThreadOne");
        Thread t2 = new Thread(new ThreadTest());
        t2.setName("ThreadTwo");
        t1. start();
        t1.sleep(5000);
        t2. start();
        t1.sleep(1000);
        System.out.println("线程执行结束。");
    }
}
```

**执行结果**：

```java
线程：Thread[ThreadOne,5,main] 正在执行...
线程：Thread[ThreadTwo,5,main] 正在执行...
线程执行结束。
```

> **Tips**： 该测验主要针对线程的创建方式以及线程的执行 start 方法的测验，并附带进行了线程 setName 和线程 sleep 方法的使用。对于线程其他常用方法的使用如 wait 方法等，会在后续小节进行详细讲解。

## 10. 小结

本节课程的重中之重就是掌握线程的 3 中创建方式以及 Thread 类常用方法的使用，一定要掌握并吃透。

线程 Thread 的创建方式以及执行方式是学习多线程并发的前提条件，我们在使用无返回值的多线程创建方式时，推荐使用方式二进行多线程的创建。如果针对具体的业务场景需要使用多线程执行结果的返回值，那我们需要使用方式三进行线程的创建。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
