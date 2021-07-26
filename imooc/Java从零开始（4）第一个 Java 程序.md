# 第一个 Java 程序

本节我们将以`Windows`操作系统为例，编写并执行第一个`Java`程序。在这之前，请确保你的操作系统上已经安装了[JDK](https://www.imooc.com/wiki/javalesson/installationwindows.html)

## 1. 编译程序

大家可能有个疑问，为什么需要编译程序呢？计算机不能直接执行我们编写的源代码吗？

这是由于计算机只能识别由 0 和 1 组成的二进制代码。需要通过编译将源代码转换为计算机认识的二进制代码。

### 1.1 Java 中的编译

在 Java 中，程序不是直接被编译为可执行文件，而是被编译为**字节码**文件， `JVM`（Java 虚拟机）在运行时执行字节码文件。当我们使用 `javac` 编译器时，Java 源代码文件被编译为字节码文件，字节码文件以扩展名 `.class` 的形式保存在磁盘上。当程序运行时，字节码文件将被转换为机器代码，并在内存中执行。

总的来说，Java 源代码需要被“**转换**”两次才能被计算机执行：

1. **Java 源代码被编译为字节码**：由 javac 前端编译器完成；
2. **字节码被编译为机器码**：由 JVM 的执行引擎完成。

下图描述了一个`Java`程序从**编写**到**编译**，再到**执行**的步骤：

![](//img.mukewang.com/wiki/5e8d8e700863af8113900730.jpg)

### 1.2 快速执行一个程序

#### 1.2.1 准备工作

为了确保我们人生中第一个 Java 程序能够顺利执行，请首先确认你已经完成了下面两件事情：

1. 已经在你的电脑上成功安装 `JDK`；
2. 有一个文本编辑器。（如 Windows 的记事本，Unix 的 vim）。

#### 1.2.2 编写源代码

新建一个文本，输入如下内容，将其命名为 `HelloWorld.java`并保存

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello World!");
    }
}
```

> **Tips**：
>
> * 文件名必须与**类名**同名（**类名**就是上面第一行代码中`class`后面的`HelloWorld`），并且扩展名为 `.java`，文件名和类名都区分大小写，所以请确保文件名和类名大小写保持一致；
> * 类的命名约定以大写字母开头；
> * 上述代码中的所有符号，请使用英文半角；
> * 别忽略了代码第三行结尾的**分号**`;`。

下面为实操中的截图。

打开记事本，输入源代码：

![](//img.mukewang.com/wiki/5e8d8ea109bc999f06920399.jpg)

点击**文件 -> 保存**，将文件保存到磁盘（此处保存在桌面上），注意：这里要将文件名命名为`HelloWorld.java`，保存类型选择**所有文件**，编码选择`UTF-8`：

![](//img.mukewang.com/wiki/5e8d8ed2097504ea16441082.jpg)

#### 1.2.3 编译后执行

现在，打开你常用的命令行工具。（如 Windows 上的**命令提示符**、MacOS 上的**终端**），并进入你刚刚保存 Java 源代码的目录。

在`Windows`搜索栏中输入`cmd`即可打开：

![](//img.mukewang.com/wiki/5e8d8eed098beb7b15581342.jpg)

使用`cd`命令，进入到我们刚刚保存源代码的目录（即进入到桌面目录`cd Desktop`），使用`dir`命令，查看目录下内容：

![](//img.mukewang.com/wiki/5e8d8f150964477e19561044.jpg)

使用`javac`命令编译源代码：

```java
javac HelloWorld.java
```

![](//img.mukewang.com/wiki/5e8d8f360926c43a08420610.jpg)

执行编译命令后，你会发现同一目录下会生成一个`HelloWorld.class`文件，这个`HelloWorld.class`就是我们说的**字节码**文件。接下来就是通过 `JVM` 执行字节码文件了，输入下面的命令：

```java
java HelloWorld
```

如果一切正常，你会看到屏幕上输出如下内容：

![](//img.mukewang.com/wiki/5e8d8f5b09f4e9f007120123.jpg)

祝贺你！成功执行了第一个 Java 程序！

> **Tips**：
>
> * `javac` 是编译器，而 `java` 是虚拟机，先使用 `javac` 编译器编译源代码，再使用`java`虚拟机执行字节码文件。这就是我们上面提到的 java 源代码被转换两次到执行的过程。
> * `java` 命令后面的**参数是类名**。换句话说，我们只需要给虚拟机**传递类名作为参数**即可，虚拟机会自动查找对应的以`.class`为扩展名的文件并且执行。

## 2. 理解程序

我们现在对上面的 `HelloWorld.java` 代码进行解析：

```java
public class HelloWorld {
		...
}
```

这段代码表示我们定义了一个**公开类**，类名为 `HelloWorld`，按照约定类名首字母要大写。

`public` 表示这个类是**公开的**（至于什么是公开的，我们后面再讲），`class`顾名思义，就是**类**， `public` 和 `class`都是 Java 中的**关键字**，必须小写。 花括号`{}` 中间为类的定义。

下面我们来看下花括号中间的 `main()`方法代码：

```java
    public static void main(String[] args) {
        ...
    }
```

这段代码表示我们定义了一个入口方法，注意：**入口方法**是 Java 程序执行的起点。

`public` 和 `static`分别表示方法是公开的、静态的，`void`是方法的返回类型，`main`后面的括号`()` 用来将方法的参数括起来，`String[]`是参数的类型，`args`是参数的名称。

当然，刚刚对于代码的解释略显生涩，作为初学者无需深究每一项的含义。但有一个点请务必牢记：Java 规定，某个类定义的`public static void main(String[] args)`是 Java 程序的固定入口方法。

在方法内部，有一行代码：

```java
        System.out.println("Hello World!");
```

这行代码的意义就是将`Hello World`这一串字符打印在屏幕上。你可以尝试修改这块内容，将代码中的`Hello World!`，替换为其他内容，例如：

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("我爱!");
    }
}
```

保存源代码，打开命令行，重新编译执行，屏幕上输出的内容将会被成功替换。

## 3. 小结

本小节我们一起完成了第一个 Java 程序的执行，在这个过程中，我们了解了编译相关的知识，知道了使用 `javac`命令，可以将一个`.java`扩展名的源代码文件编译成`.class`扩展名的字节码文件，使用 `java` 命令，可以执行字节码文件。我们也对程序的基础语法有了大致的了解。如果你对程序的理解还有困惑，也不必担心，在后面的章节中，我们将具体学习 Java 的基础语法，随着知识的累积，你也将更加游刃有余地阅读、编写程序。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
