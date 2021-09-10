# synchronized 关键字

## 1. 前言

本节内容主要是对 synchronized 关键字的使用进行讲解，具体内容点如下：

* 了解 synchronized 关键字的概念，从总体层面对 synchronized 关键字进行了解，是我们本节课程的基础知识；
* 了解 synchronized 关键字的作用，知道 synchronized 关键字使用的意义，使我们学习本节内容的出发点；
* 掌握 synchronized 关键字的 3 中使用方式，使我们本节课程的核心内容，所有的内容讲解都是围绕这一知识点进行的；
* 了解 synchronized 关键字的内存语义，将 synchronized 关键字与 Java 的线程内存模型进行关联，更好的了解 synchronized 关键字的作用及意义，为本节重点内容。

## 2. synchronized 关键字介绍

**概念**：synchronized 同步块是 Java 提供的一种原子性内置锁，Java 中的每个对象都可以把它当作一个同步锁来使用，这些 Java 内置的使用者看不到的锁被称为内部锁，也叫作监视器锁。

**线程的执行**：代码在进入 synchronized 代码块前会自动获取内部锁，这时候其他线程访问该同步代码块时会被阻塞挂起。拿到内部锁的线程会在正常退出同步代码块或者抛出异常后或者在同步块内调用了该内置锁资源的 wait 系列方法时释放该内置锁。

**内置锁**：即排它锁，也就是当一个线程获取这个锁后，其他线程必须等待该线程释放锁后才能获取该锁。

> **Tips**：由于 Java 中的线程是与操作系统的原生线程一一对应的，所以当阻塞一个线程时，需要从用户态切换到内核态执行阻塞操作，这是很耗时的操作，而 synchronized 的使用就会导致上下文切换。
>
>
> 后续章节我们会引入 Lock 接口和 ReadWriteLock 接口，能在一定场景下很好地避免 synchronized 关键字导致的上下文切换问题。

## 3. synchronized 关键字的作用

**作用**：在并发编程中存在线程安全问题，使用 synchronized 关键字能够有效的避免多线程环境下的线程安全问题，线程安全问题主要考虑以下三点：

* 存在共享数据，共享数据是对多线程可见的，所有的线程都有权限对共享数据进行操作；
* 多线程共同操作共享数据。关键字 synchronized 可以保证在同一时刻，只有一个线程可以执行某个同步方法或者同步代码块，同时 synchronized 关键字可以保证一个线程变化的可见性；
* 多线程共同操作共享数据且涉及增删改操作。如果只是查询操作，是不需要使用 synchronized 关键字的，在涉及到增删改操作时，为了保证数据的准确性，可以选择使用 synchronized 关键字。

## 4. synchronized 的三种使用方式

Java 中每一个对象都可以作为锁，这是 synchronized 实现同步的基础。synchronized 的三种使用方式如下：

* **普通同步方法（实例方法）**：锁是当前实例对象 ，进入同步代码前要获得当前实例的锁；
* **静态同步方法**：锁是当前类的 class 对象 ，进入同步代码前要获得当前类对象的锁；
* **同步方法块**：锁是括号里面的对象，对给定对象加锁，进入同步代码库前要获得给定对象的锁。

接下来会对这三种使用方式进行详细的讲解，也是本节课程的核心内容。

## 5. synchronized 作用于实例方法

为了更加深刻的体会 synchronized 作用于实例方法的使用，我们先来设计一个场景，并根据要求，通过代码的实例进行实现。

**场景设计**：

* 创建两个线程，分别设置线程名称为 threadOne 和 threadTwo；
* 创建一个共享的 int 数据类型的 count，初始值为 0；
* 两个线程同时对该共享数据进行增 1 操作，每次操作 count 的值增加 1；
* 对于 count 数值加 1 的操作，请创建一个单独的 increase 方法进行实现；
* increase 方法中，先打印进入的线程名称，然后进行 1000 毫秒的 sleep，每次加 1 操作后，打印操作的线程名称和 count 的值；
* 运行程序，观察打印结果。

**结果预期**：因为 increase 方法有两个打印的语句，不会出现 threadOne 和 threadTwo 的交替打印，一个线程执行完 2 句打印之后，才能给另外一个线程执行。

**实例**：

```java
public class DemoTest extends Thread {
    //共享资源
    static int count = 0;

    /**
     * synchronized 修饰实例方法
     */
    public synchronized void increase() throws InterruptedException {
	    sleep(1000);
	    count++;
	    System.out.println(Thread.currentThread().getName() + ": " + count);
	}
    @Override
    public void run() {
        try {
            increase();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
    public static void main(String[] args) throws InterruptedException {
        DemoTest test = new DemoTest();
        Thread t1 = new Thread(test);
        Thread t2 = new Thread(test);
        t1.setName("threadOne");
        t2.setName("threadTwo");
        t1. start();
        t2. start();
    }
```

**结果验证**：

```java
threadTwo 获取到锁，其他线程在我执行完毕之前，不可进入。
threadTwo: 1
threadOne 获取到锁，其他线程在我执行完毕之前，不可进入。
threadOne: 2
```

从结果可以看出，threadTwo 进入该方法后，休眠了 1000 毫秒，此时线程 threadOne 依然没有办法进入，因为 threadTwo 已经获取了锁，threadOne 只能等待 threadTwo 执行完毕后才可进入执行，这就是 synchronized 修饰实例方法的使用。

> **Tips**：仔细看 DemoTest test = new DemoTest () 这就话，我们创建了一个 DemoTest 的实例对象，对于修饰普通方法，synchronized 关键字的锁即为 test 这个实例对象。

## 6. synchronized 作用于静态方法

> **Tips**：对于 synchronized 作用于静态方法，锁为当前的 class，要明白与修饰普通方法的区别，普通方法的锁为创建的实例对象。为了更好地理解，我们对第 5 点讲解的代码进行微调，然后观察打印结果。

**代码修改**：其他代码不变，只修改如下部分代码。

* 新增创建一个实例对象 testNew ；
* 将线程 2 设置为 testNew 。

```java
public static void main(String[] args) throws InterruptedException {
        DemoTest test = new DemoTest();
        DemoTest testNew = new DemoTest();
        Thread t1 = new Thread(test);
        Thread t2 = new Thread(testNew);
        t1.setName("threadOne");
        t2.setName("threadTwo");
        t1. start();
        t2. start();
    }
```

**结果验证**：

```java
threadTwo 获取到锁，其他线程在我执行完毕之前，不可进入。
threadOne 获取到锁，其他线程在我执行完毕之前，不可进入。
threadTwo: 1
threadOne: 2
```

**结果分析**：我们发现 threadTwo 和 threadOne 同时进入了该方法，为什么会出现这种问题呢？

因为我们此次的修改是新增了 testNew 这个实例对象，也就是说，threadTwo 的锁是 testNew ，threadOne 的锁是 test。

两个线程持有两个不同的锁，不会产生互相 block。相信讲到这里，同学对实例对象锁的作用也了解了，那么我们再次将 increase 方法进行修改，将其修改成静态方法，然后输出结果。

**代码修改**：

```java
public static synchronized void increase() throws InterruptedException {
        System.out.println(Thread.currentThread().getName() + "获取到锁，其他线程在我执行完毕之前，不可进入。" );
        sleep(1000);
        count++;
        System.out.println(Thread.currentThread().getName() + ": " + count);
    }
```

**结果验证**：

```java
threadOne获取到锁，其他线程在我执行完毕之前，不可进入。
threadOne: 1
threadTwo获取到锁，其他线程在我执行完毕之前，不可进入。
threadTwo: 2
```

**结果分析**：我们看到，结果又恢复了正常，为什么会这样？

关键的原因在于，synchronized 修饰静态方法，锁为当前 class，即 DemoTest.class。

```java
public class DemoTest extends Thread {}
```

无论 threadOne 和 threadTwo 如何进行 new 实例对象的创建，也不会改变锁是 DemoTest.class 的这一事实。

## 7. synchronized 作用于同步代码块

> **Tips**：对于 synchronized 作用于同步代码，锁为任何我们创建的对象，只要是个对象即可，如 new Object () 可以作为锁，new String () 也可作为锁，当然如果传入 this，那么此时代表当前对象。

我们将代码恢复到第 5 点的知识，然后在第 5 点知识的基础上，再次对代码进行如下修改：

**代码修改**：

```java
	/**
     * synchronized 修饰实例方法
     */
    static final Object objectLock = new Object(); //创建一个对象锁
    public static void increase() throws InterruptedException {
        System.out.println(Thread.currentThread().getName() + "获取到锁，其他线程在我执行完毕之前，不可进入。" );
        synchronized (objectLock) {
            sleep(1000);
            count++;
            System.out.println(Thread.currentThread().getName() + ": " + count);
        }
    }
```

**代码解析**：我们创建了一个 objectLock 作为对象锁，除了第一句打印语句，让后三句代码加入了 synchronized 同步代码块，当 threadOne 进入时，threadTwo 不可进入后三句代码的执行。

**结果验证**：

```java
threadOne 获取到锁，其他线程在我执行完毕之前，不可进入。
threadTwo 获取到锁，其他线程在我执行完毕之前，不可进入。
threadOne: 1
threadTwo: 2
```

## 8. 小结

本节内容的核心即 synchronized 关键字的 3 种使用方式，这是必须要掌握的问题。除此之外，不同的使用方法获取到的锁的类型是不一样的，这是本节内容的重点，也是必须要掌握的知识。

对 synchronized 关键字的熟练使用，是并发编程中的一项重要技能。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
