# JVM 四种引用

## 1. 前言

延续上节可达性分析法的讲解，本节主要讲解可达性分析法所使用的四种引用类型，本节主要内容如下：

* 强引用的定义以及如何消除强引用，为本节重点内容之一；
* 软引用的定义及使用场景，为本节重点内容之一；
* 弱引用的定义及代码示例，验证任何情况下，只要发生 GC 就会回收弱引用对象，为本节重点内容之一；
* 虚引用的定义以及作用，为本节重点内容之一；

## 2. 可达性分析的四种引用类型

上节课程内容讲解了可达性分析，可达性分析的 GC Roots 均为引用对象，那么引用对象有 4 种引用类型如下：

* **强引用**；
* **软引用**；
* **弱引用**；
* **虚引用**。

本节课程内容与可达性分析相辅相成，学习者务必在学习完可达性分析内容后再学习本节内容。

## 3. 强引用

**定义**：强引用就是指在程序代码之中普遍存在的，类似`Object obj = new Object()`这类的引用，只要强引用还存在，垃圾收集器永远不会回收掉被引用的对象。当内存空间不足，Java 虚拟机宁愿抛出 OutOfMemoryError 错误，使程序异常终止，也不会靠随意回收具有强引用的对象来解决内存不足的问题。

**代码示例**：

```java
public class DemoTest {
    public static void main(String[] args) {
        Object obj = new Object(); // 强引用
    }
}
```

在强引用的定义中有这样一句话：“只要强引用还存在，垃圾收集器永远不会回收掉被引用的对象。” 那么有没有办法将强引用消除呢？

**消除强引用示例代码**：

```java
public class DemoTest {
    public static void main(String[] args) {
        Object obj = new Object(); // 强引用
        obj = null; //消除强引用
    }
}
```

如果不使用强引用时，可以赋值 `obj=null`，显示的设置 obj 为 null，则 gc 认为该对象不存在引用，这时候就可以回收此对象。

## 4. 软引用

**定义**：软引用用来描述一些还有用，但并非必需的对象。对于软引用关联着的对象，如果内存充足，则垃圾回收器不会回收该对象，如果内存不够了，就会回收这些对象的内存。

在 JDK 1.2 之后，提供了 SoftReference 类来实现软引用。软引用可用来实现内存敏感的高速缓存。软引用可以和一个引用队列（ReferenceQueue）联合使用，如果软引用所引用的对象被垃圾回收器回收，Java 虚拟机就会把这个软引用加入到与之关联的引用队列中。

**软引用使用场景**：Android 应用图片

软引用主要应用于内存敏感的高速缓存，在 Android 系统中经常使用到。一般情况下，Android 应用会用到大量的默认图片，这些图片很多地方会用到。如果每次都去读取图片，由于读取文件需要硬件操作，速度较慢，会导致性能较低。所以我们考虑将图片缓存起来，需要的时候直接从内存中读取。

但是，由于图片占用内存空间比较大，缓存很多图片需要很多的内存，就可能比较容易发生 OutOfMemory 异常。这时，我们可以考虑使用软引用技术来避免这个问题发生。

SoftReference 可以解决 OOM 的问题，每一个对象通过软引用进行实例化，这个对象就以 cache 的形式保存起来，当再次调用这个对象时，那么直接通过软引用中的 get() 方法，就可以得到对象中的资源数据，这样就没必要再次进行读取了，直接从 cache 中就可以读取得到，当内存将要发生 OOM 的时候，GC 会迅速把所有的软引用清除，防止 OOM 发生。

## 5. 弱引用

**定义**：弱引用描述非必需对象。被弱引用关联的对象只能生存到下一次垃圾回收之前，垃圾收集器工作之后，无论当前内存是否足够，都会回收掉只被弱引用关联的对象。Java 中的类 WeakReference 表示弱引用。

**代码示例**：

```java
import java.lang.ref.WeakReference;

public class Main {
    public static void main(String[] args) {
        WeakReference<String> sr = new WeakReference<String>(new String("hello"));
        System.out.println(sr.get());
        System.gc();                //通知JVM的gc进行垃圾回收
        System.out.println(sr.get());
    }
}
```

**结果验证**：第二个输出结果是 null，这说明只要 JVM 进行垃圾回收，被弱引用关联的对象必定会被回收掉。

```java
hello
null
```

## 6. 虚引用

**定义**："虚引用"顾名思义，就是形同虚设，与其他几种引用都不同，虚引用并不会决定对象的生命周期。如果一个对象仅持有虚引用，那么它就和没有任何引用一样，在任何时候都可能被垃圾回收。虚引用在 Java 中使用 `java.lang.ref.PhantomReference` 类表示。

**作用**：虚引用主要用来跟踪对象被垃圾回收的活动。

**虚引用与软引用和弱引用的区别**：虚引用必须和引用队列（ReferenceQueue）联合使用。当垃圾回收器准备回收一个对象时，如果发现它还有虚引用，就会在回收对象的内存之前，把这个虚引用加入到与之关联的引用队列中。程序可以通过判断引用队列中是否已经加入了虚引用，来了解被引用的对象是否将要被垃圾回收。程序如果发现某个虚引用已经被加入到引用队列，那么就可以在所引用的对象的内存被回收之前采取必要的行动。

**使用示例**：虚引用必须和引用队列（ReferenceQueue）联合使用

```java
import java.lang.ref.PhantomReference;
import java.lang.ref.ReferenceQueue;
public class Main {
    public static void main(String[] args) {
        ReferenceQueue<String> queue = new ReferenceQueue<String>();
        PhantomReference<String> pr = new PhantomReference<String>(new String("hello"), queue);
        System.out.println(pr.get());
    }
}
```

## 7. 小结

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyiouc3j60jg0afwg302)

本节主要讲解可达性分析的四种对象引用类型，通篇皆为重点内容，需要学习者理解并掌握这四种引用类型。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
