---
title: Java 从零开始（24）Java 包
zhihu-url: https://zhuanlan.zhihu.com/p/408514025
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
---

# Java 包

当我们的程序规模越来越大，类的数量也会随之增多，数量繁多的类会造成项目的混乱，不易于维护管理。本小节所介绍的包就是为了**将类分类**而产生的，我们可以使用包让程序结构更加清晰且易于管理。本小节将会学习到**什么是包**，**如何声明包**，**包作用域**以及**包的命名规范**等知识点。

## 1. 概述

### 1.1 什么是包

包是一个命名空间，它可以将类和接口进行分组。

从概念上讲，我们可以将包看作是计算机上的不同文件夹。例如，我们可以将图片保存在一个文件夹中，将视频保存在另一个文件夹中。

### 1.2 为什么需要包

正如我们在前言中所提到的，当程序规模越来越大，类的数量也会随之增多。使用包将这些类分组后，可以让类更易于**查找**和**使用**，也可以**避免命名冲突**，还可以**控制访问**。

易于查找和使用很容易理解，试想我们电脑的某一个文件夹中，存放着成百上千的各类文件，即使使用搜索功能都不能迅速定位到我们想要的文件，而将这些文件按照功能或特点分类存入到不同的文件夹下，就可以大大提高我们查找和使用的效率。对于 Java 语言的类也是同样的道理。

对于命名冲突，我们来看实际开发中的一个例子：一个 Java 工程由 3 个程序员协同开发，程序员 `A` 写了一个工具类 `Util`，程序员 `B` 也写了一个工具类 `Util`，程序员 `C` 既想使用 `A` 的 `Util`，又想使用 `B` 的 `Util`，这个时候可以分别将 `A`、`B` 两个人的 `Util` 类放入两个不同的包下，就可以解决命名冲突了。

对于控制访问，我们将在本小节的包作用域部分举例讲解。

## 2. 包声明

### 2.1 语法

在一个类中，使用关键字 `package` 在类的顶部声明一个包：

```java
package service;

public class DemoService {
    ...
}
```

上面的代码就表示 `DemoService` 类放到了 `service` 包下，我们也可以说：`DemoService` 在 `service` 包中声明。我们的类如果在包中声明，那么类就必须放置在源码目录下以包名称命名的子目录中。这样编译器才能在系统中找到 `Java` 文件。

### 2.2 实例

下面我们使用 `IntelliJ IDEA` 来演示一个简单的建包到建类的过程。

新建一个名为 `imooc` 的工程，之前我们已经学过如何在 IDE 中新建工程，因此这里不再截图演示。查看下图，`imooc` 是工程名，`src` 是源码目录。

![](https://xushuhui.gitee.io/image/imooc/5eb50cca0985442904450084.jpg)

首先在源码目录新建一个包，命名为 `service`。在 `src` 目录上点击鼠标右键 -> New -> Package

![](https://xushuhui.gitee.io/image/imooc/5eb50cf409285b4511910377.jpg)

填入包名 `service`

![](https://xushuhui.gitee.io/image/imooc/5eb50d0a095396d906640108.jpg)

输入 `Enter` 完成包的新建，此时工程结构如下：

![](https://xushuhui.gitee.io/image/imooc/5eb50d2709798ab404560116.jpg)

接下来在包下面新建一个 `DemoService` 类。在 `service` 点击右键 -> New -> Java Class，输入类名称 `DemoService` 完成新建，此时工程结构如下：

![](https://xushuhui.gitee.io/image/imooc/5eb50d40094b85b805070165.jpg)

而 IDE 也会在类的顶部自动为我们声明包：

![](https://xushuhui.gitee.io/image/imooc/5eb50d54096e47b208020322.jpg)

这样就完成了一个包到类的新建。

重复以上步骤，我们接着在源码目录下新建一个 `util` 包，在 `util` 包中新建一个 `DemoUtil` 类。

此时，`imooc` 工程在物理硬盘上的文件结构为：

```java
├── imooc
│   └── src
│       ├── service
│       │   └── DemoService.java
│       └── util
│           └── DemoUtil.java
```

通过上述一系列演示，我们验证了这个结论：**所有的 Java 文件对应的目录层级要和包的层级保持一致**。

另外，包下面也可以包含子包，子包中也可以声明类，例如，可以在上面的 `service` 包下新建一个 `demo` 包，在 `demo` 包下新建一个 `DemoClass` 类。如下代码可以看到类顶部包的声明，**要使用`.` 分隔多个包名**：

```java
package service.demo;

public class DemoClass {
}
```

## 3. 导入包

### 3.1 完整类名导入

在类名前面加上包名（我们称为完整类名），就可以使用指定包中的类。例如，对与上面示例的项目结构，我们想在 `DemoService` 类中使用 `DemoUtil` 类，可以这样写：

```java
package service;

public class DemoService {
    util.DemoUtil demoUtil = new util.DemoUtil();
}
```

如果使用同一包中的类，则不必在类前面指定包名。

### 3.2 import 关键字导入

上面的语法比较冗长，我们可以使用 `import` 关键字在文件顶部导入指定包中类，在代码中就不必指定包名来使用类了。例如，上面的代码可以改成写：

```java
package service;

import util.DemoUtil;

public class DemoService {
    DemoUtil demoUtil = new DemoUtil();
}
```

> **Tips**：如果有两个或多个类名相同，那么就不能使用 `import` 关键字同时导入了，此时建议使用完整类名的方式使用类。当然，实际的编码中，我们也不推荐命名多个同名类。

### 3.3 通配符导入

在使用 `import` 关键字时，可以使用通配符 `*` 来导入指定包下面的所有类。例如：

```java
package service;

import util.*;

public class DemoService {
    DemoUtil demoUtil = new DemoUtil();
}
```

尽管它看起来非常方便，但是我们不推荐这种写法。一方面降低了代码的可读性，另外也可能导致一些问题。

还有一点需要特别提醒，`java.lang` 包中的所有类会都被编译器隐式导入，所以在使用诸如 `System`、`String` 等类的时候，我们不需要手动导入。

## 4. 包作用域

所谓包作用域，就是没有使用访问修饰符修饰的字段和方法。同一个包中的类，可以访问**包作用域**的字段和方法。

例如，在 `com.imooc` 包下面有一个 `Student` 类：

```java
package com.imooc;

public class Student {

    // 包作用域的属性
    String name;

    // 包作用域的方法
    void sayHello() {
        System.out.println("你好！" + name);
    }

}
```

在 `com.imooc` 包下面的 `Test` 类，访问 `Student` 类中的字段和方法：

```java
package com.imooc;

public class Test {
    public static void main(String[] args) {
        Student student = new Student();
        // 操作包作用域下属性
        student.name = "Colorful";
        // 访问包作用域下的方法
        student.sayHello();
    }
}
```

由于 `Student` 与 `Test` 都属于 `com.imooc` 包，所以可以访问包作用域内的字段和方法。

## 5. 包命名规范

包名应该使用全小写字母命名，这样易于将包名与类名区分开。

为了避免名称冲突，通常建议以公司域名的倒置来确保唯一性。例如，域名为 `imooc.com`，项目的包命就应该为 `com.imooc`，其具体的子包可根据功能业务命名。

再举几个例子，大名鼎鼎的 Apache 软件基金会的包命名为 `org.apache`，Java 领域最知名开源框架 `Spring` 包名为 `org.springframework`。

## 6. 小结

本小节我们学习了 Java 包的定义与使用，Java 提供包的机制主要是为了类避免命名冲突。我们也知道了如何声明包、如何导入包以及包作用域的概念，实际项目的包名推荐使用域名倒置。

包的使用，除了在我们自己编写的项目中非常常见，你也可以翻阅 JDK 的源码，Java 平台按照功能将类放入了不同的包中。例如：基础的类在 `java.lang` 包中，用于读写的类在 `java.io` 包中，一些工具类放在 `java.util` 包中。Java 平台由成百上千的类组成，把这些类放入包中，可以使程序结构井井有条。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
