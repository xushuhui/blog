# JVM 参数：跟踪类的加载与卸载

## 1. 前言

本节内容主要是学习， JVM 跟踪类的加载与卸载的常用参数配置，这是工作中跟踪类的加载与卸载情况时 JVM 中最常用的参数配置。本节主要知识点如下：

* 理解并掌握跟踪类的加载与卸载的参数 -XX:+TraceClassLoading，为本节重点内容；
* 理解并掌握跟踪类的加载与卸载的参数 -XX:+TraceClassUnloading，为本节了解内容，非重点知识；
* 理解并掌握跟踪类的加载与卸载的参数 -XX:+PrintClassHistogram，为本节重点内容；

JVM 跟踪类的加载与卸载的常用参数是使用 JVM 所必须的知识点，通篇皆为重点掌握内容，需要在理解的基础上并掌握参数的使用方法。

## 2. 示例代码准备

此处的示例代码，与上一节的示例代码相似，但是有重要的区别。详细区别请看 Tips 的内容。

**实例**：准备测试代码，创建一个 String 类型的 ArrayList，并在 list 中添加三个元素，分别是 “Hello”，“World”，“！！！”。

> **Tips**：注意，此处的示例代码，并没有执行 gc 操作。上一节的内容是为了跟踪垃圾回收，所以需要手动调用 gc 方法而达到垃圾回收的效果。而此处我们讨论的是类的加载与卸载，此处无需进行手动垃圾回收。

```java
public class TracingClassParamsDemo {
    public static void main(String[] args) {
        ArrayList<String> list = new ArrayList<String>();
        list.add("Hello");
        list.add("World");
        list.add("!!!");
    }
}
```

## 3. -XX:+TraceClassLoading 参数

**参数作用**：-XX:+TraceClassLoading 参数是为了跟踪类的加载。

为了更好的理解并掌握 -XX:+TraceClassLoading 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+TraceClassLoading 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：由于追踪的结果日志非常庞大，此处仅展示具有代表性的类的加载。全部的类加载日志，请学习者自行执行代码进行验证。

```java
[Opened C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.lang.Object from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.util.ArrayList$SubList from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.util.ListIterator from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.util.ArrayList$SubList$1 from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded DemoMain.TracingClassParamsDemo from file:/D:/GIT-Repositories/GitLab/Demo/out/production/Demo/]
[Loaded java.lang.Class$MethodArray from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.lang.Void from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.lang.Shutdown from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
[Loaded java.lang.Shutdown$Lock from C:\Program Files\Java\jdk1.8.0_152\jre\lib\rt.jar]
```

**结果分析**：我们来对类的加载日志进行分析。

* **第一行**：Opened rt.jar。打开 rt.jar，rt.jar 全称是 Runtime，该 jar 包含了所有支持 Java 运行的核心类库，是类加载的第一步；
* **第二行**：加载 java.lang.Object。Object 是所有对象的父类，是首要加载的类；
* **第三、四、五行**：加载了 ArrayList 的相关类，我们的示例代码中使用到了 ArrayList，因此需要对该类进行加载；
* **第六行**：加载我们的测试类 TracingClassParamsDemo ；
* **第七行**：加载 java.lang.Class 类，并加载类方法 MethodArray；
* **第八行**：加载 java.lang.Void 类，因为我们的 main 函数是 void 的返回值类型，所以需要加载此类；
* **第九、十行**：加载 java.lang.Shutdown 类， JVM 结束运行后，关闭 JVM 虚拟机。

从以上对日志的分析来看，JVM 对类的加载，不仅仅加载我们代码中使用的类，还需要加载各种支持 Java 运行的核心类。类加载的日志量非常庞大，此处仅仅对重点类的加载进行日志的解读，全部的类加载日志，请学习者自行执行代码进行验证。

## 4. -XX:+TraceClassUnloading 参数

**参数作用**：-XX:+TraceClassUnloading 参数是为了跟踪类的卸载。由于系统类加载器加载的类不会被卸载，并且只加载一次，所以普通项目很难获取到类卸载的日志。

此处我们先来看看，通过系统类加载器加载的类是否会被卸载。

为了更好的理解并掌握 -XX:+TraceClassUnloading 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+TraceClassUnloading 并保存；
* **步骤 2**：运行示例代码，观察执行结果。

**结果验证**：未打印日志，未发生类的卸载。

**引出问题**：为什么看不到跟踪类卸载的日志呢？

上文提到了，由系统类加载器加载的类不能够被卸载。所以想要看到跟踪类卸载的日志，我们需要使用自定义的类加载器。通过自定义的类加载器加载的类，在类不可达的时候，会发生垃圾回收，并卸载该类。

一般情况下，开发过程中很少实现自定义的类加载器，除非有特殊的需求场景需要通过自定义的类加载器进行类的加载，因此此处对 -XX:+TraceClassUnloading 稍作了解即可。

## 5. -XX:+PrintClassHistogram 参数

**参数作用**：-XX:+PrintClassHistogram 参数是打印、查看系统中类的分布情况。

为了更好的理解并掌握 -XX:+PrintClassHistogram 参数，我们通过如下步骤进行操作。

* **步骤 1**：在 VM Options 中配置参数 -XX:+PrintClassHistogram 并保存；
* **步骤 2**：修改示例代码，在代码最后添加代码 `Thread.sleep(99999999999L)`，确保 main 函数长时间不能结束执行（当然了，也可以使用 while (true) 语句，进行无限长时间循环来创造这种场景，可自行选择），以便于观察类的分布情况；

```java
public class TracingClassParamsDemo {
    public static void main(String[] args) {
        ArrayList<String> list = new ArrayList<String>();
        list.add("Hello");
        list.add("World");
        list.add("!!!");
        try {
            Thread.sleep(99999999999L);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
```

* **步骤 3**：运行程序，观察日志输出；
* **步骤 4**：不中断 main 函数的运行，将鼠标指针移动到日志输出的 console 界面并单击鼠标左键，确保鼠标的实时位置在 console 界面。按下键盘 Ctrl+Break 键，观察日志输出。

**结果验证**：我们执行步骤 3 时，没有观察到日志的输出。当我们尝试步骤 4 时，获取到了日志输出如下图所示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxjrdshj60jf0asjx802)

**结果分析**： 这是系统中类的分布情况，那我们来看下日志中每列的表头部分代表的意思：

* **num**：自增的序列号，只是为了标注行数，没有特殊的意义；
* **instances**：实例数量，即类的数量；
* **bytes**：实例所占子节数，即占据的内存空间大小；
* **class name**：具体的实例。

我们取出第 3 条日志来进行下翻译：系统当前状态下，java.lang.String 类型的实例共有 2700 个，共占用空间大小为 64800 bytes。

## 6. 小结

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxk3g66j60jg09ftat02)

本小节的重点内容即，我们所讲述的三个常用的跟踪类的加载与卸载参数，学习者需要对这三个常用参数的意义，以及使用方式进行掌握。

需要特别注意第二个参数 -XX:+TraceClassUnloading，在后续讲解类加载器的时候，会实现自定义的类加载器，并使用该参数演示类的卸载。通篇皆为重点内容，需要认真对待，掌握本节要点内容。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
