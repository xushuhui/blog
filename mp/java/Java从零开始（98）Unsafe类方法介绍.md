---
title: Java从零开始（98）Unsafe类方法介绍
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Unsafe 类方法介绍


## 1. 前言

本节内容主要是对 Unsafe 类方法进行介绍，JDK jar 包中的 Unsafe 类提供了硬件级别的原子性操作，Unsafe 类中的方法都是 native 方法，它们使用 JNI 的方式访问本地 C＋＋实现库。

本节我们来了解一下 Unsafe 提供的几个主要的方法进行介绍。为我们后续对 Unsafe 方法的使用奠定良好的基础。

## 2. 方法介绍

|方法|作用|
|----|----|
|** objectFieldOffset(Field)**                | 返回指定的变量在所属类中的内存偏移地址，该偏移地址仅仅在该 UnSafe 函数中访问指定字段时使用。                                                                                                                 |
|**arrayBaseOffset(Class)**                   | 获取取数组中第一个元系的地址。                                                                                                                                                                               |
|**arrayIndexScale(Class)**                   | 获取数组中一个元素占用的字节。                                                                                                                                                                               |
|**compareAndSwapLong(Object,long,long,long)**| 比较对象 obj 中偏移量为 offset 的变量的值是否与 expect 相等，相等则使用 update 值更新，然后返回 true，否则返回 false。                                                                                       |
|**longgetLongvolatile(Object,long)**         | 获取对象 obj 中偏移量为 offset 的变量对应 volatile 语义的值。                                                                                                                                                |
|**void putLongvolatile(Object,long,long)**   | 设置 obj 对象中 offset 偏移的类型为 long 的 field 的值为 value, 支持 volatile 语义。                                                                                                                         |
|**putOrderedLong(Object,long,long)**         | 设置 obj 对象中 offset 偏移地址对应的 long 型 field 的值为 value。这是一个有延迟的 putLongvolatile 方法，并且不保证值修改对其他线程立刻可见。只有在变量使用 volatile 修饰并且预计会被意外修改时才使用该方法。|
|**unpark(Object)**                           | 唤醒调用 park 后阻塞的线程。                                                                                                                                                                                 |

## 3. park 方法介绍

**方法描述**：** void park(booleanisAbsolute,longtime)**：阻塞当前线程，其中参数 isAbsolute 等于 false 且 time 等于 0 表示一直阻塞。

**方法解读**：time 大于 0 表示等待指定的 time 后阻塞线程会被唤醒，这个 time 是个相对值，是个增量值，也就是相对当前时间累加 time 后当前线程就会被唤醒。如果 isAbsolute 等于 true，并且 time 大于 0，则表示阻塞的线程到指定的时间点后会被唤醒。

这里 time 是个绝对时间，是将某个时间点换算为 ms 后的值。另外，当其他线程调用了当前阻塞线程的 interrupt 方法而中断了当前线程时，当前线程也会返回，而当其他线程调用了 unPark 方法并且把当前线程作为参数时当前线程也会返回。

## 4. JDK8 新增的函数

|方法|作用|
|----|----|
|** getAndSetLong(Object, long, long)**| 获取对象 obj 中偏移量为 offset 的变量 volaile 语义的当前值，并设置变量 volaile 语义的值为 update。|
|**getAndAddLong(Object,long,long)**   | 方法获取 object 中偏移量为 offset 的 volatile 变量的当前值，并设置变量值为原始值加上 addValue     |

## 5. 小结

本节的核心内容即 Usafe 方法的了解，为下边讲解 Unsafe 方法的使用奠定一个良好的基础。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
