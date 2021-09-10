# Unsafe 类的使用

## 1. 前言

本节对 Unsafe 类的使用进行讲解，上一小节内容已经对 Unsafe 类的常用方法有了大体的概括，本节主要内容点如下：

Unsafe 类的简介，对 UnSafe 类有一个整体的认识；

Unsafe 类的创建以及创建过程中避免的异常机制，这是开始使用 UnSafe 类的前提；

了解 Unsafe 类操作对象属性方法的使用，这是本节内容的重点之一；

了解 Unsafe 操作数组元素方法的使用，也是本节内容的重点之一。

本节内容意在了解并掌握 Unsafe 类的常用方法的使用。

## 2. Unsafe 类简介

Unsafe 类是 Java 整个并发包底层实现的核心，它具有像 C++ 的指针一样直接操作内存的能力，而这也就意味着其越过了 JVM 的限制。

Unsafe 类有如下的特点：

* Unsafe 不受 JVM 管理，也就无法被自动 GC，需要手动 GC，容易出现内存泄漏；
* Unsafe 的大部分方法中必须提供原始地址 （内存地址） 和被替换对象的地址，偏移量需自行计算，一旦出现问题必然是 JVM 崩溃级别的异常，会导致整个应用程序直接 crash；
* 直接操作内存，也意味着其速度更快，在高并发的条件之下能够很好地提高效率。

## 3. Unsafe 类的创建

Unsafe 类是不可以通过 new 关键字直接创建的。Unsafe 类的构造函数是私有的，而对外提供的静态方法 Unsafe.getUnsafe () 又对调用者的 ClassLoader 有限制 ，如果这个方法的调用者不是由 Boot ClassLoader 加载的，则会报错。

**实例**：通过 main 方法进行调用，报错。

```java
import sun.misc.Unsafe;
import sun.misc.VM;
import sun.reflect.Reflection;

public class DemoTest {
    public static void main(String[] args) {
        getUnsafe();
    }
    public static Unsafe getUnsafe() {
        Class unsafeClass = Reflection.getCallerClass();
        if (!VM.isSystemDomainLoader(unsafeClass.getClassLoader())) {
            throw new SecurityException("Unsafe");
        } else {
            return null;
        }
    }
}
```

**运行结果**：

```java
Exception in thread "main" java.lang.InternalError: CallerSensitive annotation expected at frame 1
	at sun.reflect.Reflection.getCallerClass(Native Method)
	at leeCode.DemoTest.getUnsafe(DemoTest.java:12)
	at leeCode.DemoTest.main(DemoTest.java:9)
```

**报错原因**： Java 源码中由开发者自定义的类都是由 Appliaction ClassLoader 加载的，也就是说 main 函数所依赖的 jar 包都是 ClassLoader 加载的，所以会报错。

所以正常情况下我们无法直接使用 Unsafe ，如果需要使用它，则需要利用反射

**实例**：通过反射，成功加载 Unsafe 类。

```java
import sun.misc.Unsafe;
import java.lang.reflect.Field;

public class DemoTest {
    public static void main(String[] args) {
        Unsafe unsafe = getUnsafe();
        System.out.println("Unsafe 加载成功："+unsafe);
    }
    public static Unsafe getUnsafe() {
        Unsafe unsafe = null;
        try {
            Field field = Unsafe.class.getDeclaredField("theUnsafe");
            field.setAccessible(true);
            unsafe = (Unsafe) field.get(null);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return unsafe;
    }
}
```

**结果验证**：

```java
Unsafe 加载成功：sun.misc.Unsafe@677327b6
```

**总结**：Unsafe 类的加载必须使用反射进行，否则会报错。

## 4. Unsafe 类操作对象属性

**操作对象属性的常用方法有**：

* **public native Object getObject(Object o, long offset)**：获取一个 Java 对象中偏移地址为 offset 的属性的值，此方法可以突破修饰符的限制，类似的方法有 getInt ()、getDouble () 等，同理还有 putObject () 方法；
* **public native Object getObjectVolatile(Object o, long offset)**：强制从主存中获取目标对象指定偏移量的属性值，类似的方法有 getIntVolatile ()，getDoubleVolatile () 等，同理还有 putObjectVolatile ()；
* **public native void putOrderedObject(Object o, long offset, Object x)**：设置目标对象中偏移地址 offset 对应的对象类型属性的值为指定值。这是一个有序或者有延迟的 putObjectVolatile () 方法，并且不保证值的改变被其他线程立即看到。只有在属性被 volatile 修饰并且期望被修改的时候使用才会生效，类似的方法有 putOrderedInt () 和 putOrderedLong ()；
* **public native long objectFieldOffset(Field f)**：返回给定的非静态属性在它的类的存储分配中的位置 （偏移地址），然后可根据偏移地址直接对属性进行修改，可突破属性的访问修饰符限制。

**实例**：

```java
import sun.misc.Unsafe;
import java.lang.reflect.Field;

public class DemoTest {
    private String name;
    public static void main(String[] args) {
        Unsafe unsafe = getUnsafe();
        try {
            DemoTest directMemory = (DemoTest) unsafe.allocateInstance(DemoTest.class);
            //获取name属性
            long nameOffset = unsafe.objectFieldOffset(DemoTest.class.getDeclaredField("name"));
            //设置name属性
            unsafe.putObject(directMemory, nameOffset, "并发编程");
            System.out.println("属性设置成功："+ directMemory.getName());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    public static Unsafe getUnsafe() {
        Unsafe unsafe = null;
        try {
            Field field = Unsafe.class.getDeclaredField("theUnsafe");
            field.setAccessible(true);
            unsafe = (Unsafe) field.get(null);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return unsafe;
    }
    public void setName(String name) {
        this.name = name;
    }
    public String getName() {
        return name;
    }
}
```

结果验证：

```java
属性设置成功：并发编程
```

## 5. Unsafe 操作数组元素

Unsafe 操作数组元素主要有如下两个方法：

* **public native int arrayBaseOffset(Class arrayClass)**：返回数组类型的第一个元素的偏移地址 （基础偏移地址）；
* **public native int arrayIndexScale(Class arrayClass)**：返回数组中元素与元素之间的偏移地址的增量，配合 arrayBaseOffset () 使用就可以定位到任何一个元素的地址。

**实例**：

```java
import sun.misc.Unsafe;
import java.lang.reflect.Field;

public class DemoTest {
    private static String[] names = {"多线程", "Java", "并发编程"};
    public static void main(String[] args) {
        Unsafe unsafe = getUnsafe();
        try {
            Class<?> a = String[].class;
            int base = unsafe.arrayBaseOffset(a);
            int scale = unsafe.arrayIndexScale(a);
            // base + i * scale 即为字符串数组下标 i 在对象的内存中的偏移地址
            System.out.println(unsafe.getObject(names, (long) base + 2 * scale));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    public static Unsafe getUnsafe() {
        Unsafe unsafe = null;
        try {
            Field field = Unsafe.class.getDeclaredField("theUnsafe");
            field.setAccessible(true);
            unsafe = (Unsafe) field.get(null);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return unsafe;
    }
}
```

**结果验证**：

```java
并发编程
```

通过对数组的元素的地址进行内存偏移，最后得到的结果为最后一个元素，并发编程。base + 2 * scale 表示字符串数组下标 i 在对象的内存中的偏移地址，偏移两个元素，得到最后一个元素。

## 6. 小结

本节内容主要对 Unsafe 类的常用方法的使用进行介绍，使学习者能够在使用 Unsafe 类操作对象和数组时，能够快速的使用课程中提供的实例思路。其实 Unsafe 类还可以操作内存地址，操作 CAS，对于初学者来说就比较晦涩了。

此处对操作对象和操作数组常用的方法的讲解，是本节的核心知识，掌握 Unsafe 类的常用方法的使用，非常重要。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
