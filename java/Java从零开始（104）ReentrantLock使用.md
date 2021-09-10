# ReentrantLock 使用

## 1. 前言

本节内容主要是对 ReentrantLock 的使用进行讲解，之前对于 Lock 接口进行了讲解，ReentrantLock 是 Lock 接口的常用实现子类，占据着十分重要的地位。本节内容的知识点如下：

* ReentrantLock 基本方法的使用，即 lock 与 unlock 方法的使用，这是最基础的方法使用，为重点内容；
* ReentrantLock lockInterruptibly 与 tryLock 方法的使用，也是经常使用到的方法，为本节重点内容；
* ReentrantLock 公平锁与非公平锁的使用，也是本节的重点内容；
* ReentrantLock 其他方法的介绍与使用。

通篇来看，ReentrantLock 所有的知识点均为重点内容，是必须要掌握的内容。

## 2. ReentrantLock 介绍

ReentrantLock 在 Java 中也是一个基础的锁，ReentrantLock 实现 Lock 接口提供一系列的基础函数，开发人员可以灵活的使用函数满足各种复杂多变应用场景。

**定义**：ReentrantLock 是一个可重入且独占式的锁，它具有与使用 synchronized 监视器锁相同的基本行为和语义，但与 synchronized 关键字相比，它更灵活、更强大，增加了轮询、超时、中断等高级功能。

ReentrantLock，顾名思义，它是支持可重入锁的锁，是一种递归无阻塞的同步机制。除此之外，该锁还支持获取锁时的公平和非公平选择。

**公平性**：ReentrantLock 的内部类 Sync 继承了 AQS，分为公平锁 FairSync 和非公平锁 NonfairSync。

如果在绝对时间上，先对锁进行获取的请求一定先被满足，那么这个锁是公平的，反之，是不公平的。公平锁的获取，也就是等待时间最长的线程最优先获取锁，也可以说锁获取是顺序的。

ReentrantLock 的公平与否，可以通过它的构造函数来决定。

## 3. ReentrantLock 基本方法 lock 与 unlock 的使用

我们使用一个之前涉及到的 synchronized 的场景，通过 lock 接口进行实现。

**场景回顾**：

* 创建两个线程，创建方式可自选；
* 定义一个全局共享的 static int 变量 count，初始值为 0；
* 两个线程同时操作 count，每次操作 count 加 1；
* 每个线程做 100 次 count 的增加操作。

**结果预期**：获取到的结果为 200。之前我们使用了 synchronized 关键字和乐观锁 Amotic 操作进行了实现，那么此处我们进行 ReentrantLock 的实现方式。

**实现步骤**：

* step 1 ：创建 ReentrantLock 实例，以便于调用 lock 方法和 unlock 方法；
* step 2：在 synchronized 的同步代码块处，将 synchronized 实现替换为 lock 实现。

**实例**：

```java
public class DemoTest{
    private static int count = 0; //定义count = 0
    private static ReentrantLock lock = new ReentrantLock();//创建 lock 实例
    public static void main(String[] args) {
        for (int i = 0; i < 2; i++) { //通过for循环创建两个线程
            new Thread(new Runnable() {
                @Override
                public void run() {
                    //每个线程让count自增100次
                    for (int i = 0; i < 100; i++) {
                        try {
                            lock.lock(); //调用 lock 方法
                            count++;
                        } finally {
                            lock.unlock(); //调用unlock方法释放锁
                        }
                    }
                }
            }). start();
        }
        try{
            Thread.sleep(2000);
        }catch (Exception e){
            e.printStackTrace();
        }
        System.out.println(count);
    }
}
```

**代码分析**：

我们通过 try finally 模块，替代了之前的 synchronized 代码块，顺利的实现了多线程下的并发。

## 4. tryLock 方法

我们之前进行过介绍，Lock 接口包含了两种 tryLock 方法，一种无参数，一种带参数。

* **boolean tryLock()**：仅在调用时锁为空闲状态才获取该锁。如果锁可用，则获取锁，并立即返回值 true。如果锁不可用，则此方法将立即返回值 false；
* **boolean tryLock(long time, TimeUnit unit)**：如果锁在给定的等待时间内空闲，并且当前线程未被中断，则获取锁；

为了了解两种方法的使用，我们先来设置一个简单的使用场景。

**场景设置**：

* 创建两个线程，创建方式自选；
* 两个线程同时执行代码逻辑；
* 代码逻辑使用 boolean tryLock () 方法，如果获取到锁，执行打印当前线程名称，并沉睡 5000 毫秒；如果未获取锁，则打印 timeout，并处理异常信息；
* 观察结果并进行分析；
* 修改代码，使用 boolean tryLock (long time, TimeUnit unit) 方法，设置时间为 4000 毫秒；
* 观察结果并进行分析；
* 再次修改代码，使用 boolean tryLock (long time, TimeUnit unit) 方法，设置时间为 6000 毫秒；
* 观察结果并进行分析。

**实例**：使用 boolean tryLock () 方法

```java
public class DemoTest implements Runnable{
    private static Lock locks = new ReentrantLock();
    @Override
    public void run() {
        try {
            if(locks.tryLock()){ //尝试获取锁，获取成功则进入执行，不成功则执行finally模块
                System.out.println(Thread.currentThread().getName()+"-->");
                Thread.sleep(5000);
            }else{
                System.out.println(Thread.currentThread().getName()+" time out ");
            }
        } catch (InterruptedException e) {
             e.printStackTrace();
        }finally {
            try {
                locks.unlock();
            } catch (Exception e) {
                System.out.println(Thread.currentThread().getName() + "未获取到锁，释放锁抛出异常");
            }
        }
    }
    public static void main(String[] args) throws InterruptedException {
        DemoTest test =new DemoTest();
        Thread t1 = new Thread(test);
        Thread t2 = new Thread(test);
        t1. start();
        t2. start();
        t1.join();
        t2.join();
        System.out.println("over");
    }
}
```

**结果验证**：

```java
Thread-1-->
Thread-0 time out
Thread-0 未获取到锁，释放锁抛出异常
over
```

**结果分析**：从打印的结果来看， Thread-1 获取了锁权限，而 Thread-0 没有获取锁权限，这就是 tryLock，没有获取到锁资源则放弃执行，直接调用 finally。

**实例**：使用 boolean tryLock (4000 ms) 方法

将 if 判断进行修改如下：

```java
 if(locks.tryLock(4000,TimeUnit.MILLISECONDS)){ //尝试获取锁，获取成功则进入执行，不成功则执行finally模块
      System.out.println(Thread.currentThread().getName()+"-->");
      Thread.sleep(5000);
  }
```

**结果验证**：

```java
Thread-1-->
Thread-0 time out
Thread-0 未获取到锁，释放锁抛出异常
over
```

**结果分析**：tryLock 方法，虽然等待 4000 毫秒，但是这段时间不足以等待 Thread-1 释放资源锁，所以还是超时。 我们换成 6000 毫秒试试。

**实例**：使用 boolean tryLock (6000 ms) 方法

将 if 判断进行修改如下：

```java
 if(locks.tryLock(6000,TimeUnit.MILLISECONDS)){ //尝试获取锁，获取成功则进入执行，不成功则执行finally模块
      System.out.println(Thread.currentThread().getName()+"-->");
      Thread.sleep(5000);
  }
```

**结果验证**：

```java
Thread-1-->
Thread-0-->
over
```

**结果分析**：tryLock 方法，等待 6000 毫秒，Thread-1 先进入执行，5000 毫秒后 Thread-0 进入执行，都能够有机会获取锁。

**总结**：以上就是 tryLock 方法的使用，可以指定最长的获取锁的时间，如果获取则执行，未获取则放弃执行。

## 5. 公平锁与非公平锁

**分类**：根据线程获取锁的抢占机制，锁可以分为公平锁和非公平锁。

**公平锁**：表示线程获取锁的顺序是按照线程请求锁的时间早晚来决定的，也就是最早请求锁的线程将最早获取到锁。

**非公平锁**：非公平锁则在运行时闯入，不遵循先到先执行的规则。

**ReentrantLock**：ReentrantLock 提供了公平和非公平锁的实现。

**ReentrantLock 实例**：

```java
//公平锁
ReentrantLock pairLock = new ReentrantLock(true);
//非公平锁
ReentrantLock pairLock1 = new ReentrantLock(false);
//如果构造函数不传递参数，则默认是非公平锁。
ReentrantLock pairLock2 = new ReentrantLock();
```

**场景介绍**：通过模拟一个场景假设，来了解公平锁与非公平锁。

* 假设线程 A 已经持有了锁，这时候线程 B 请求该锁将会被挂起；
* 当线程 A 释放锁后，假如当前有线程 C 也需要获取该锁，如果采用非公平锁方式，则根据线程调度策略，线程 B 和线程 C 两者之一可能获取锁，这时候不需要任何其他干涉；
* 而如果使用公平锁则需要把 C 挂起，让 B 获取当前锁，因为 B 先到所以先执行。

> **Tips**：在没有公平性需求的前提下尽量使用非公平锁，因为公平锁会带来性能开销。

## 6. lockInterruptibly 方法

**lockInterruptibly () 方法**：能够中断等待获取锁的线程。当两个线程同时通过 lock.lockInterruptibly () 获取某个锁时，假若此时线程 A 获取到了锁，而线程 B 只有等待，那么对线程 B 调用 threadB.interrupt () 方法能够中断线程 B 的等待过程。

**场景设计**：

* 创建两个线程，创建方式可自选实现；
* 第一个线程先调用 start 方法，沉睡 20 毫秒后调用第二个线程的 start 方法，确保第一个线程先获取锁，第二个线程进入等待；
* 最后调用第二个线程的 interrupt 方法，终止线程；
* run 方法的逻辑为打印 0，1，2，3，4，每打印一个数字前，先沉睡 1000 毫秒；
* 观察结果，看是否第二个线程被终止。

**实例**：

```java
public class DemoTest{
    private Lock lock = new ReentrantLock();

    public void doBussiness() {
        String name = Thread.currentThread().getName();
        try {
            System.out.println(name + " 开始获取锁");
            lock.lockInterruptibly(); //调用lockInterruptibly方法，表示可中断等待
            System.out.println(name + " 得到锁，开工干活");
            for (int i=0; i<5; i++) {
                Thread.sleep(1000);
                System.out.println(name + " : " + i);
            }
        } catch (InterruptedException e) {
            System.out.println(name + " 被中断");
        } finally {
            try {
                lock.unlock();
                System.out.println(name + " 释放锁");
            } catch (Exception e) {
                System.out.println(name + " : 没有得到锁的线程运行结束");
            }
        }
    }

    public static void main(String[] args) throws InterruptedException {
        final DemoTest lockTest = new DemoTest();
        Thread t0 = new Thread(new Runnable() {
                    public void run() {
                        lockTest.doBussiness();
                    }});
        Thread t1 = new Thread(new Runnable() {
                    public void run() {
                        lockTest.doBussiness();
                    }});

        t0. start();
        Thread.sleep(20);
        t1. start();
        t1.interrupt();
    }
}
```

**结果验证**：可以看到，thread -1 被中断了。

```java
Thread-0 开始获取锁
Thread-0 得到锁，开工干活
Thread-1 开始获取锁
Thread-1 被中断
Thread-1 : 没有得到锁的线程运行结束
Thread-0 : 0
Thread-0 : 1
Thread-0 : 2
Thread-0 : 3
Thread-0 : 4
Thread-0 释放锁
```

## 7. ReentrantLock 其他方法介绍

对 ReentrantLock 来说，方法很多样，如下介绍 ReentrantLock 其他的方法，有兴趣的同学可以自行的尝试使用。

* **getHoldCount()**：当前线程调用 lock () 方法的次数；
* **getQueueLength()**：当前正在等待获取 Lock 锁的线程的估计数；
* **getWaitQueueLength(Condition condition)**：当前正在等待状态的线程的估计数，需要传入 Condition 对象；
* **hasWaiters(Condition condition)**：查询是否有线程正在等待与 Lock 锁有关的 Condition 条件；
* **hasQueuedThread(Thread thread)**：查询指定的线程是否正在等待获取 Lock 锁；
* **hasQueuedThreads()**：查询是否有线程正在等待获取此锁定；
* **isFair()**：判断当前 Lock 锁是不是公平锁；
* **isHeldByCurrentThread()**：查询当前线程是否保持此锁定；
* **isLocked()**：查询此锁定是否由任意线程保持。

## 8. 小结

本节内容对 ReentrantLock 进行了比较详细的讲解，通篇内容皆为重点内容，需要同学们进行细致的掌握。核心内容即为 ReentrantLock 的使用，可以根据小节中的实例进行自行的编码和试验，更深刻的理解 ReentrantLock 的使用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
