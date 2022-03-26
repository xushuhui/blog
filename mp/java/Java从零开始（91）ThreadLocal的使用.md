---
title: Java从零开始（91）ThreadLocal的使用
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# ThreadLocal 的使用


## 1. 前言

本节内容主要是对 ThreadLocal 进行深入的讲解，具体内容点如下：

* 了解 ThreadLocal 的诞生，以及总体概括，是学习本节知识的基础；
* 了解 ThreadLocal 的作用，从整体层面理解 ThreadLocal 的程序作用，为本节的次重点；
* 掌握 ThreadLocal set 方法的使用，为本节重点内容；
* 掌握 ThreadLocal get 方法的使用，为本节重点内容；
* 掌握 ThreadLocal remove 方法的使用，为本节重点内容；
* 掌握多线程下的 ThreadLocal 的使用，为本节内容的核心。

## 2. ThreadLocal 概述

**诞生**：早在 JDK 1.2 的版本中就提供 java.lang.ThreadLocal，ThreadLocal 为解决多线程程序的并发问题提供了一种新的思路。使用这个工具类可以很简洁地编写出优美的多线程程序。

**概述**：ThreadLocal 很容易让人望文生义，想当然地认为是一个 “本地线程”。其实，ThreadLocal 并不是一个 Thread，而是 Thread 的局部变量，也许把它命名为 ThreadLocalVariable 更容易让人理解一些。

当使用 ThreadLocal 维护变量时，ThreadLocal 为每个使用该变量的线程提供独立的变量副本，所以每一个线程都可以独立地改变自己的副本，而不会影响其它线程所对应的副本。

**总体概括**：从线程的角度看，目标变量就象是线程的本地变量，这也是类名中 “Local” 所要表达的意思。

了解完 ThreadLocal 的总体介绍后，对其有了一个总体的了解，那我们接下来继续探究 ThreadLocal 的真实面貌以及使用。

## 3. ThreadLocal 的作用

**作用**：ThreadLocal 是 JDK 包提供的，它提供了线程本地变量，也就是如果你创建了一个 ThreadLocal 变量，那么访问这个变量的每个线程都会有这个变量的一个本地副本。当多个线程操作这个变量时，实际操作的是自己本地内存里面的变量，从而避免了线程安全问题。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz6cxr1j60jg07ata602)

ThreadLocal 是线程本地存储，在每个线程中都创建了一个 ThreadLocalMap 对象，每个线程可以访问自己内部 ThreadLocalMap 对象内的 value。通过这种方式，避免资源在多线程间共享。

**使用场景**：如为每个线程分配一个 JDBC 连接 Connection。这样就可以保证每个线程的都在各自的 Connection 上进行数据库的操作，不会出现 A 线程关了 B 线程正在使用的 Connection。还有 Session 管理等问题。

## 4. ThreadLocal set 方法

**方法介绍**：set 方法是为了设置 ThreadLocal 变量，设置成功后，该变量只能够被当前线程访问，其他线程不可直接访问操作改变量。

**实例**：

```java
public class DemoTest{
    public static void main(String[] args){
        ThreadLocal<String> localVariable = new ThreadLocal<> () ;
        localVariable.set("Hello World");
    }
}
```

> **Tips**：set 方法可以设置任何类型的值，无论是 String 类型 ，Integer 类型，Object 等类型，原因在于 set 方法的 JDK 源码实现是基于泛型的实现，此处只是拿 String 类型进行的举例。

**实例**：

```java
public void set(T value) { // T value , 泛型实现，可以 set 任何对象类型
        Thread t = Thread.currentThread();
        ThreadLocalMap map = getMap(t);
        if (map != null)
            map.set(this, value);
        else
            createMap(t, value);
    }
```

## 5. ThreadLocal get 方法

**方法介绍**：get 方法是为了获取 ThreadLocal 变量的值，get 方法没有任何入参，直接调用即可获取。

**实例**：

```java
public class DemoTest{
    public static void main(String[] args){
        ThreadLocal<String> localVariable = new ThreadLocal<> () ;
        localVariable.set("Hello World");
        System.out.println(localVariable.get());
    }
}
```

**结果验证**：

```java
Hello World
```

**探究**：请看如下程序，并给出输出结果

**实例**：

```java
public class DemoTest{
    public static void main(String[] args){
        ThreadLocal<String> localVariable = new ThreadLocal<> () ;
        localVariable.set("Hello World");
        localVariable.set("World is beautiful");
        System.out.println(localVariable.get());
        System.out.println(localVariable.get());
    }
}
```

**探究解析**：从程序中来看，我们进行了两次 set 方法的使用。

第一次 set 的值为 Hello World ；第二次 set 的值为 World is beautiful。接下来我们进行了两次打印输出 get 方法，那么这两次打印输出的结果都会是 World is beautiful。 原因在于第二次 set 的值覆盖了第一次 set 的值，所以只能 get 到 World is beautiful。

**结果验证**：

```java
World is beautiful
World is beautiful
```

**总结**：ThreadLocal 中只能设置一个变量值，因为多次 set 变量的值会覆盖前一次 set 的值，我们之前提出过，ThreadLocal 其实是使用 ThreadLocalMap 进行的 value 存储，那么多次设置会覆盖之前的 value，这是 get 方法无需入参的原因，因为只有一个变量值。

## 6. ThreadLocal remove 方法

**方法介绍**：remove 方法是为了清除 ThreadLocal 变量，清除成功后，该 ThreadLocal 中没有变量值。

**实例**：

```java
public class DemoTest{
    public static void main(String[] args){
        ThreadLocal<String> localVariable = new ThreadLocal<> () ;
        localVariable.set("Hello World");
        System.out.println(localVariable.get());
        localVariable.remove();
        System.out.println(localVariable.get());
    }
}
```

> **Tips**：remove 方法同 get 方法一样，是没有任何入参的，因为 ThreadLocal 中只能存储一个变量值，那么 remove 方法会直接清除这个变量值。

**结果验证**：

```java
Hello World
null
```

## 7. 多线程下的 ThreadLocal

对 ThreadLocal 的常用方法我们已经进行了详细的讲解，那么多线程下的 ThreadLocal 才是它存在的真实意义，那么问了更好的学习多线程下的 ThreadLocal，我们来进行场景的创建，通过场景进行代码实验，更好的体会并掌握 ThreadLocal 的使用。

**场景设计**：

* 创建一个全局的静态 ThreadLocal 变量，存储 String 类型变量；
* 创建两个线程，分别为 threadOne 和 threadTwo；
* threadOne 进行 set 方法设置，设置完成后沉睡 5000 毫秒，苏醒后进行 get 方法打印；
* threadTwo 进行 set 方法设置，设置完成后直接 get 方法打印，打印完成后调用 remove 方法，并打印 remove 方法调用完毕语句；
* 开启线程 threadOne 和 threadTwo ；
* 执行程序，并观察打印结果。

**结果预期**：在 threadOne 设置成功后进入了 5000 毫秒的休眠状态，此时由于只有 threadTwo 调用了 remove 方法，不会影响 threadOne 的 get 方法打印，这体现了 ThreadLocal 变量的最显著特性，线程私有操作。

**实例**：

```java
public class DemoTest{
    static ThreadLocal<String> local = new ThreadLocal<>();
    public static void main(String[] args){

        Thread threadOne = new Thread(new Runnable() {
            @Override
            public void run() {
                local.set("threadOne's local value");
                try {
                    Thread.sleep(5000); //沉睡5000 毫秒，确保 threadTwo 执行 remove 完成
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                System.out.println(local.get());
            }
        });
        Thread threadTwo = new Thread(new Runnable() {
            @Override
            public void run() {
                local.set("threadTwo's local value");
                System.out.println(local.get());
                local.remove();
                System.out.println("local 变量执行 remove 操作完毕。");
            }
        });
        threadTwo. start();
        threadOne. start();
    }
}
```

**结果验证**：

```java
threadTwo's local value
local 变量执行 remove 操作完毕。
threadOne's local value
```

从以上结果来看，在 threadTwo 执行完 remove 方法后，threadOne 仍然能够成功打印，这更加证明了 ThreadLocal 的专属特性，线程独有数据，其他线程不可侵犯。

## 8. 小结

ThreadLocal 是解决线程安全问题一个很好的思路，它通过为每个线程提供一个独立的变量副本解决了变量并发访问的冲突问题。在很多情况下，ThreadLocal 比直接使用 synchronized 同步机制解决线程安全问题更简单，更方便，且结果程序拥有更高的并发性。

本节的重中之重是掌握 ThreadLocal 的方法使用以及其特点，核心内容为多线程下的 ThreadLocal 的使用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
