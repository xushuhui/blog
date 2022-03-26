---
title: Java从零开始（96）volatile关键字
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# volatile 关键字


## 1. 前言

本节内容主要是对 volatile 关键字进行讲解，具体内容点如下：

* volatile 关键字概念介绍，从整体层面了解 volatile 关键字；
* volatile 关键字与 synchronized 关键字的区别，这是本节的重点内容之一，了解 volatile 关键字与 synchronized 关键字的区别，才能更好地区分并掌握两钟关键字的使用；
* volatile 关键字原理介绍，也是本节课程的重点之一；
* volatile 关键字的使用，是本节课程的核心内容，所有的知识点都是围绕这一目的进行讲解的。

## 2. volatile 关键字介绍

**概念**：volatile 关键字解决内存可见性问题，是一种弱形式的同步。

**介绍**：该关键字可以确保当一个线程更新共享变量时，更新操作对其他线程马上可见。当一个变量被声明为 volatile 时，线程在写入变量时不会把值缓存在寄存器或者其他地方，而是会把值刷新回主内存。

当其他线程读取该共享变量时，会从主内存重新获取最新值，而不是使用当前线程的工作内存中的值。

## 3. volatile 与 synchronized 的区别

**相似处**：volatile 的内存语义和 synchronized 有相似之处，具体来说就是，当线程写入了 volatile 变量值时就等价于线程退出 synchronized 同步块（把写入工作内存的变量值同步到主内存），读取 volatile 变量值时就相当于进入 synchronized 同步块（ 先清空本地内存变量值，再从主内存获取最新值）。

**区别**：使用锁的方式可以解决共享变量内存可见性问题，但是使用锁太笨重，因为它会带来线程上下文的切换开销。具体区别如下：

* volatile 本质是在告诉 jvm 当前变量在寄存器（工作内存）中的值是不确定的，需要从主存中读取；synchronized 则是锁定当前变量，只有当前线程可以访问该变量，其他线程被阻塞住；
* volatile 仅能使用在变量级别；synchronized 则可以使用在变量、方法、和类级别的；
* volatile 仅能实现变量的修改可见性，不能保证原子性；而 synchronized 则可以保证变量的修改可见性和原子性；
* volatile 不会造成线程的阻塞；synchronized 可能会造成线程的阻塞；
* volatile 标记的变量不会被编译器优化；synchronized 标记的变量可以被编译器优化

## 4. volatile 原理

**原理介绍**：Java 语言提供了一种弱同步机制，即 volatile 变量，用来确保将变量的更新操作通知到其他线程。

当把变量声明为 volatile 类型后，编译器与运行时都会注意到这个变量是共享的，volatile 变量不会被缓存在寄存器或者对其他处理器不可见的地方，因此在读取 volatile 类型的变量时总会返回最新写入的值。

> **Tips**：在访问 volatile 变量时不会执行加锁操作，因此也就不会使执行线程阻塞，因此 volatile 变量是一种比 sychronized 关键字更轻量级的同步机制。

我们来通过下图对非 volatile 关键字修饰的普通变量的读取方式进行理解，从而更加细致的了解 volatile 关键字修饰的变量。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnz7urfnj60jg08jmz802)

当对非 volatile 变量进行读写的时候，每个线程先从内存拷贝变量到 CPU 缓存中。如果计算机有多个 CPU，每个线程可能在不同的 CPU 上被处理，这意味着每个线程可以拷贝到不同的 CPU cache 中。

而声明变量是 volatile 的，JVM 保证了每次读变量都从内存中读，跳过 CPU cache。

## 5. volatile 关键字的使用

为了对 volatile 关键字有着更深的使用理解，我们通过一个非常简单的场景的设计来进行学习。

**场景设计**：

* 创建一个 Student 类，该类有一个 String 属性，name；
* 将 name 的 get 和 set 方法设置为同步方法；
* 使用 synchronized 关键字实现；
* 使用 volatile 关键字实现。

这是一个非常简单的场景，场景中只涉及到了一个类的两个同步方法，通过对两种关键字的实现，能更好的理解 volatile 关键字的使用。

**实例**： synchronized 关键字实现

```java
public class Student {
    private String name;

    public synchronized String getName() {
        return name;
    }

    public synchronized void setName(String name) {
        this.name = name;
    }
}
```

**实例**： volatile 关键字实现

```java
public class Student {
    private volatile String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
```

**总结**：在这里使用 synchronized 和使用 volatile 是等价的，都解决了共享变量 name 的内存可见性问题。

但是前者是独占锁，同时只能有一个线程调用 get（）方法，其他调用线程会被阻塞，同时会存在线程上下文切换和线程重新调度的开销，这也是使用锁方式不好的地方。

而后者是非阻塞算法，不会造成线程上下文切换的开销。

## 6. 小结

本节内容的核心知识点即 volatile 关键字的使用方式，在掌握核心知识之前，需要对重点内容进行理解和学习，本节内容所有的重点知识如 volatile 关键字原理，与 synchronized 关键字的区别，都是围绕核心知识进行的讲解。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
