# 守护线程与用户线程

## 1. 前言

本节内容主要是对守护线程与用户线程进行深入的讲解，具体内容点如下：

* 了解守护线程与用户线程的定义及区别，使我们学习本节内容的基础知识点；
* 了解守护线程的特点，是我们掌握守护线程的第一步；
* 掌握守护线程的创建，是本节内容的重点；
* 通过守护线程与 JVM 的退出实验，更加深入的理解守护线程的地位以及作用，为本节内容次重点；
* 了解守护线程的作用及使用场景，为后续开发过程中提供守护线程创建的知识基础。

## 2. 守护线程与用户线程的定义及区别

Java 中的线程分为两类，分别为 daemon 线程（守护线程〉和 user 线程（用户线程）。

在 JVM 启动时会调用 main 函数， main 函数所在的线程就是一个用户线程，其实在 JVM 内部同时还启动了好多守护线程，比如垃圾回收线程。

**守护线程定义**：所谓守护线程，是指在程序运行的时候在后台提供一种通用服务的线程。比如垃圾回收线程就是一个很称职的守护者，并且这种线程并不属于程序中不可或缺的部分。

因此，当所有的非守护线程结束时，程序也就终止了，同时会杀死进程中的所有守护线程。反过来说，只要任何非守护线程还在运行，程序就不会终止。

**用户线程定义**：某种意义上的主要用户线程，只要有用户线程未执行完毕，JVM 虚拟机不会退出。

**区别**：在本质上，用户线程和守护线程并没有太大区别，唯一的区别就是当最后一个非守护线程结束时，JVM 会正常退出，而不管当前是否有守护线程，也就是说守护线程是否结束并不影响 JVM 的退出。

言外之意，只要有一个用户线程还没结束， 正常情况下 JVM 就不会退出。

## 3. 守护线程的特点

Java 中的守护线程和 Linux 中的守护进程是有些区别的，Linux 守护进程是系统级别的，当系统退出时，才会终止。

而 Java 中的守护线程是 JVM 级别的，当 JVM 中无任何用户进程时，守护进程销毁，JVM 退出，程序终止。总结来说，Java 守护进程的最主要的特点有：

* 守护线程是运行在程序后台的线程；
* 守护线程创建的线程，依然是守护线程；
* 守护线程不会影响 JVM 的退出，当 JVM 只剩余守护线程时，JVM 进行退出；
* 守护线程在 JVM 退出时，自动销毁。

## 4. 守护线程的创建

**创建方式**：将线程转换为守护线程可以通过调用 Thread 对象的 setDaemon (true) 方法来实现。

**创建细节**：

* thread.setDaemon (true) 必须在 thread.start () 之前设置，否则会跑出一个 llegalThreadStateException 异常。你不能把正在运行的常规线程设置为守护线程；
* 在 Daemon 线程中产生的新线程也是 Daemon 的；
* 守护线程应该永远不去访问固有资源，如文件、数据库，因为它会在任何时候甚至在一个操作的中间发生中断。

**线程创建代码示例**：

```java
public class DemoTest {
    public static void main(String[] args) throws InterruptedException {
        Thread threadOne = new Thread(new Runnable() {
            @Override
            public void run() {
                //代码执行逻辑
            }
        });
        threadOne.setDaemon(true); //设置threadOne为守护线程
        threadOne. start();
    }
}
```

## 5. 守护线程与 JVM 的退出实验

为了更好的了解守护线程与 JVM 是否退出的关系，我们首先来设计一个守护线程正在运行，但用户线程执行完毕导致的 JVM 退出的场景。

**场景设计**：

* 创建 1 个线程，线程名为 threadOne；
* run 方法线程 sleep 1000 毫秒后，进行求和计算，求解 1 + 2 + 3 + … + 100 的值；
* 将线程 threadOne 设置为守护线程；
* 执行代码，最终打印的结果；
* 加入 join 方法，强制让用户线程等待守护线程 threadOne；
* 执行代码，最终打印的结果。

**期望结果**：

* 未加入 join 方法之前，threadOne 不能执行求和逻辑，无打印输出，因为 main 函数线程执行完毕后，JVM 退出，守护线程也就随之死亡，无打印结果；
* 加入 join 方法后，可以打印求和结果，因为 main 函数线程需要等待 threadOne 线程执行完毕后才继续向下执行，main 函数执行完毕，JVM 退出。

> **Tips**：main 函数就是一个用户线程，main 方法执行时，只有一个用户线程，如果 main 函数执行完毕，用户线程销毁，JVM 退出，此时不会考虑守护线程是否执行完毕，直接退出。

**代码实现 - 不加入 join 方法**：

```java
public class DemoTest {
    public static void main(String[] args){
        Thread threadOne = new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                int sum = 0;
                for (int i = 1; i  <= 100; i++) {
                    sum = sum + i;
                }
                System.out.println("守护线程，最终求和的值为： " + sum);
            }
        });
        threadOne.setDaemon(true); //设置threadOne为守护线程
        threadOne. start();
        System.out.println("main 函数线程执行完毕， JVM 退出。");
    }
}
```

**执行结果验证**：

```java
main 函数线程执行完毕， JVM 退出。
```

从结果上可以看到，JVM 退出了，守护线程还没来得及执行，也就随着 JVM 的退出而消亡了。

**代码实现 - 加入 join 方法**：

```java
public class DemoTest {
    public static void main(String[] args){
        Thread threadOne = new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                int sum = 0;
                for (int i = 1; i  <= 100; i++) {
                    sum = sum + i;
                }
                System.out.println("守护线程，最终求和的值为： " + sum);
            }
        });
        threadOne.setDaemon(true); //设置threadOne为守护线程
        threadOne. start();
        try {
            threadOne.join(); // 加入join 方法
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("main 函数线程执行完毕， JVM 退出。");
    }
}
```

**执行结果验证**：

```java
守护线程，最终求和的值为： 5050
main 函数线程执行完毕， JVM 退出。
```

从结果来看，守护线程不决定 JVM 的退出，除非强制使用 join 方法使用户线程等待守护线程的执行结果，但是实际的开发过程中，这样的操作是不允许的，因为守护线程，默认就是不需要被用户线程等待的，是服务于用户线程的。

## 6. 守护线程的作用及使用场景

**作用**：我们以 GC 垃圾回收线程举例，它就是一个经典的守护线程，当我们的程序中不再有任何运行的 Thread, 程序就不会再产生垃圾，垃圾回收器也就无事可做，所以当垃圾回收线程是 JVM 上仅剩的线程时，垃圾回收线程会自动离开。

它始终在低级别的状态中运行，用于实时监控和管理系统中的可回收资源。

**应用场景**：

* 为其它线程提供服务支持的情况，可选用守护线程；
* 根据开发需求，程序结束时，这个线程必须正常且立刻关闭，就可以作为守护线程来使用；
* 如果一个正在执行某个操作的线程必须要执行完毕后再释放，否则就会出现不良的后果的话，那么这个线程就不能是守护线程，而是用户线程；
* 正常开发过程中，一般心跳监听，垃圾回收，临时数据清理等通用服务会选择守护线程。

## 7. 小结

掌握用户线程和守护线程的区别点非常重要，在实际的工作开发中，对一些服务型，通用型的线程服务可以根据需要选择守护线程进行执行，这样可以减少 JVM 不可退出的现象，并且可以更好地协调不同种类的线程之间的协作，减少守护线程对高优先级的用户线程的资源争夺，使系统更加的稳定。

本节的重中之重是掌握守护线程的创建以及创建需要注意的事项，了解守护线程与用户线程的区别使我们掌握守护线程的前提。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
