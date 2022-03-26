---
title: Java从零开始（81）读写锁 StampedLock
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 读写锁 StampedLock


## 1. 前言

本节带领大家认识第二个常用的 Java 并发锁工具之 StampedLock。

本节先简单介绍 StampedLock 的基本概念，然后介绍关键的编程方法，最后通过一个编程例子为大家展示 StampedLock 工具类的用法。

下面我们正式开始介绍吧。

## 2. 概念解释

我们先解释一组概念：悲观锁、乐观锁。

**悲观锁** 指的是对数据修改持保守态度，在整个数据处理时，先将数据锁定状态，然后再进行修改处理。之所以叫做悲观锁，是因为这是一种对数据的修改抱有悲观态度的并发控制方式，认为数据被并发修改的概率比较大，所以需要在修改之前先加锁。

**乐观锁** 正好相反，抱着数据访问一般情况下不会造成冲突的观点，先对数据做修改，在正式提交修改结果是才会做冲突检查，如果发现修改的是旧版本的数据，则返回修改失败，否则提交修改。

StampedLock 是对 ReentrantReadWriteLock 的改进。相比 ReentrantReadWriteLock 采用了悲观锁的思想对数据修改的并发控制，StampedLock 使用了乐观思想的加锁实现，具有更高的并发性。

下面我们学习其关键的编程方法。

## 3. StampedLock 的编程方法

StampedLock 提供了三种并发控制模式，介绍这三种模式过程中，我们穿插介绍关键的编程方法。

### 3.1. 独占写模式

功能和 ReentrantReadWriteLock 的写锁类似。独占写模式相关的几个方法如下。

long stamp = writeLock () 方法：获取独占写锁，可能会被阻塞。如果获取锁成功则返回一个 stamp；

tryWriteLock () 方法：尝试获取独占写锁，类似 writeLock () 方法，只是获取不到时立刻返回不会阻塞；

tryWriteLock (long time, TimeUnit unit) 方法：允许在给定的时间内尝试获取独占写锁，超时仍然未获取到时则返回；

writeLockInterruptibly () 方法：类似 writeLock () 但允许获取锁的过程被打断；

unlockWrite (long stamp) 方法：用于释放独占写锁；

tryUnlockWrite () 方法：类似 unlockWrite (), 但允许不需要 stamp 邮戳参数。

### 3.2. 悲观读模式

功能和 ReentrantReadWriteLock 的读锁类似。悲观读模式相关的几个方法如下。

long stamp = readLock () 方法：获取独占读锁，可能会被阻塞。如果获取锁成功则返回一个 stamp；

unlockRead (long stamp) 方法：用于释放读锁；

tryReadLock () 方法：尝试获取读锁，类似 readLock () 方法，只是获取不到时立刻返回不会阻塞；

tryReadLock (long time, TimeUnit unit) 方法：允许在给定的时间内尝试获取读锁，超时仍然未获取到时则返回；

readLockInterruptibly () 方法：类似 readLock () 但允许获取锁的过程被打断。

### 3.3. 乐观读模式

这是一种优化的读模式。乐观读模式相关的几个方法如下。

tryOptimisticRead () 方法：非阻塞尝试乐观获取读锁，只有当写锁没有被获取时返回一个非 0 的 stamp 。乐观读取模式适用于短时间读取操作，降低竞争和提高吞吐量。在使用时一般需将数据存储到一个副本中，在后继处理中用于对比数据是否是最新状态；

validate (long stamp) 方法：用于检查在获取到读锁 stamp 后，锁有没被其他写线程抢占。如果写锁没有被获取，那么 validate () 方法返回 true。可多次调用验证这一信息。

另外，此类也提供了一组读写锁之间的**转换方法**：

tryConvertToWriteLock (long stamp) 方法：尝试转换为写锁。转换条件：

tryConvertToReadLock (long stamp) 方法：尝试转换为悲观读锁。

tryConvertToOptimisticRead (long stamp) 方法：尝试转换为乐观读锁。

**注意此类的编程方法有这样一个共通特征：**

所有获取锁的方法，都返回一个邮戳（Stamp），Stamp 为 0 表示获取失败，其余都表示成功；

所有释放锁的方法，都需要一个邮戳（Stamp），这个 Stamp 必须是和成功获取锁时得到的 Stamp 一致；

下面我们举一个具体的编程例子。

## 4. 编程示例

上面介绍了核心编程方法，我们给出一个非常简洁明了的官方例子，切实体会一下 StampedLock 的用法。

```java
import java.util.concurrent.locks.StampedLock;

public class StampedLockTest {

    // 成员变量
    private double x, y;

    // 锁实例
    private final StampedLock sl = new StampedLock();

    // 排它锁-写锁（writeLock）
    void move(double deltaX, double deltaY) {
        long stamp = sl.writeLock();
        try {
            x += deltaX;
            y += deltaY;
        } finally {
            sl.unlockWrite(stamp);
        }
    }

    // 一个只读方法
    // 其中存在乐观读锁到悲观读锁的转换
    double distanceFromOrigin() {

        // 尝试获取乐观读锁
        long stamp = sl.tryOptimisticRead();
        // 将全部变量拷贝到方法体栈内
        double currentX = x, currentY = y;
        // 检查在获取到读锁stamp后，锁有没被其他写线程抢占
        if (!sl.validate(stamp)) {
            // 如果被抢占则获取一个共享读锁（悲观获取）
            stamp = sl.readLock();
            try {
                // 将全部变量拷贝到方法体栈内
                currentX = x;
                currentY = y;
            } finally {
                // 释放共享读锁
                sl.unlockRead(stamp);
            }
        }
        // 返回计算结果
        return Math.sqrt(currentX * currentX + currentY * currentY);
    }

    // 获取读锁，并尝试转换为写锁
    void moveIfAtOrigin(double newX, double newY) {
        long stamp = sl.tryOptimisticRead();
        try {
            // 如果当前点在原点则移动
            while (x == 0.0 && y == 0.0) {
                // 尝试将获取的读锁升级为写锁
                long ws = sl.tryConvertToWriteLock(stamp);
                // 升级成功，则更新stamp，并设置坐标值，然后退出循环
                if (ws != 0L) {
                    stamp = ws;
                    x = newX;
                    y = newY;
                    break;
                } else {
                    // 读锁升级写锁失败则释放读锁，显示获取独占写锁，然后循环重试
                    sl.unlockRead(stamp);
                    stamp = sl.writeLock();
                }
            }
        } finally {
            sl.unlock(stamp);
        }
    }
}
```

注意在使用时，获取锁的操作应该放在 try 之前，而释放锁的操作需要放在 finally 中，可确保锁释放。另外需要注意 StampedLock 具有不可重入性。

## 5. 小结

本节解释了 StampedLock 的基本概念和主要的编程方法，且通过一个简单的例子展示了其用法，更多关于此工具类的概念和原理介绍，可阅读 “[Java 并发原理入门教程](http://www.imooc.com/wiki/concurrencylesson/reentrantlock.html)” 。希望大家在学习过程中，多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
