---
title: Java 从零开始（35）Java 输入输出流
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
---

# Java 输入输出流

本小节将会介绍基本输入输出的 Java 标准类，通过本小节的学习，你将了解到**什么是输入和输入**，**什么是流**；输入输出流的应用场景，`File`类的使用，什么是文件，Java 提供的输入输出流相关 API 等内容。

## 1. 什么是输入和输出（I / O）

### 1.1 基本概念

输入 / 输出这个概念，对于计算机相关专业的同学并不陌生，在计算中，**输入 / 输出**（`Input / Output`，缩写为 I / O）是信息处理系统（例如计算机））与外界（可能是人类或其他信息处理系统）之间的通信。输入是系统接收的信号或数据，输出是从系统发送的信号或数据。

那么在 Java 中，什么是输入和输出呢？要理解这个概念，可将 Java 平台视作一个系统。Java 平台是一个孤立的系统，系统之外的所有东西都是它的**环境**。系统与其环境之间的交互是一种双向对话。系统要么从其环境接收消息，要么将其消息传递给环境。当系统接收到消息时，将其称为输入，与之相反的是输出。

本小节将介绍 Java 的基本输入和输出，包括从键盘读取文本输入，将文本输出到屏幕以及从文件系统读取 / 写入文件。

Java 提供了两个用于 I / O 的包：较旧的`java.io`包（不支持符号链接）和较新的`java.nio`（“new io”）包，它对`java.nio.file`的异常处理进行了改进。

### 1.2 简单的 Java 输出——打印内容到屏幕

一直以来，我们都在向屏幕输出内容以验证我们编写的代码逻辑。向屏幕输出内容非常简单，可以由以下两种方式来完成：

```java
// 打印 Hello World，不换行
System.out.print("Hello World");
```

```java
// 打印 Hello Java，并换行
System.out.println("Hello Java");
```

### 1.3 简单的 Java 输入——从键盘输入

`java.util`包下的`Scanner`类可用于获取用户从键盘输入的内容，我们在[Java Scanner 类](http://www.imooc.com/wiki/javalesson/scanner.html)这一小节已经介绍过具体使用，实例如下：

```java
import java.util.Scanner;

/**
 * @author colorful@TaleLin
 */
public class ScannerDemo {
    public static void main(String[] args) {
        // 创建扫描器对象
        Scanner scanner = new Scanner(System.in);
        System.out.println("请输入您的姓名：");
        // 可以将用户输入的内容扫描为字符串
        String name = scanner.nextLine();
        // 打印输出
        System.out.println("你好 ".concat(name).concat(" ，欢迎来到！"));
        // 关闭扫描器
        scanner.close();
    }
}
```

运行结果：

```java
请输入您的姓名：
Colorful
你好 Colorful ，欢迎来到！
```

## 2. 什么是流（Stream）

Java 中最基本的输入 / 输入是使用流来完成的。

流是代表数据源和数据目标的对象，怎么理解这句话呢？简单来说，可以读取作为数据源的流，也可以写入作为数据目标的流。Java 中的流是长度不确定的有序字节序列，**它是一连串流动的字符**，是以先进先出的方式发送信息的通道。

## 3. 输入输出流的应用场景

上面我们已经了解了输入输出流的基本概念，那么它具体是做什么用的呢？

在`web`产品的开发中，最常开发的功能就是**上传文件到服务器**了，这个文件的读写过程就要用到输入输出流。对于计算机中文件的读写、复制和删除等操作也都要用到输入输出流。输入输出流可以说是无处不在，下面我们将会介绍 Java 中输入输出流相关的 `API`。

## 4. File 类

在 Java 中，提供了`java.io.File`类对文件和目录进行操作。

File 意思为文件，文件在计算机中非常重要，我们编写的 word 文档、PPT 演示文稿、运行游戏的`.exe`可执行文件以及我们编写的 Java 源代码等等都是文件。

### 4.1 实例化

要实例化`File`对象，需要传入一个文件或目录的路径。

File 类提供了如下 4 个构造方法：

1. `File(File parent, String child)`：从父抽象路径名和子路径名字符串创建新的文件实例；
2. `File(String pathName)`：通过将给定的路径名字符串转换为抽象路径名，创建一个新的文件实例（最常用）；
3. `File(String parent, String child)`：从父路径名字符串和子路径名字符串创建新的文件实例；
4. `File(URI uri)`：通过将给定的文件：URI 转换为抽象路径名，创建一个新的文件实例。

以`Windows`系统为例，在桌面下有一个`imooc`目录，该目录下有一个`Hello.java`文件和一个空的`images`目录，截图如下：

![](https://xushuhui.gitee.io/image/imooc/5edda58b09d6ed8721451333.jpg)

我们可以单击`Windows`的路径栏，来获取`imooc`目录的绝对路径：

![](https://xushuhui.gitee.io/image/imooc/5edda5d509f567de21451333.jpg)

有了目录和文件以及路径。我们分别实例化两个`File`对象，实例如下：

```java
import java.io.File;

public class FileDemo1 {
    public static void main(String[] args) {
        // 传入目录绝对路径
        File dir = new File("C:\\Users\\Colorful\\Desktop\\imooc\\images");
        // 传入文件绝对路径
        File file = new File("C:\\Users\\Colorful\\Desktop\\imooc\\Hello.java");
        // 打印两个File对象
        System.out.println(dir);
        System.out.println(file);
    }
}
```

我们可以直接打印`File`对象，`File`类重写了`toString()`方法，查看 Java 源码，`toString()`方法直接返回了`getPath()`实例方法，此方法返回构造方法传入的路径字符串：

![](https://xushuhui.gitee.io/image/imooc/5ed9e88309b127ca14220513.jpg)

运行结果：

```java
C:\Users\Colorful\Desktop\imooc\images
C:\Users\Colorful\Desktop\imooc\Hello.java
```

上面代码中，使用`\\`表示`Windows`下的路径分隔符`\`，`Linux`和`MacOS`下使用正斜杠`/`作为路径分隔符。假设是同样的目录结构，在`MacOS`和`Linux`下是这样表示的：

```java
File dir = new File("/Users/Colorful/Desktop/imooc/images");
```

因为`Windows`平台和其他平台路径分隔符不同，使用不同平台的开发者就难以保证路径分隔符的统一。

为了保证代码更好的兼容性，`File`类下提供了一个静态变量`separator`，用于表示当前平台的系统分隔符：

```java
// 根据当前平台输出 / 获取 \
System.out.println(File.separator);
```

### 4.2 绝对路径和相对路径

在实例化`File`对象时，既可以传入绝对路径，也可以传入相对路径。

绝对路径是以根目录开头的完整的全路径，上面代码实例中传入的是绝对路径，我们再来看看什么是相对路径，以及如何传入相对路径。

相对路径指的是当前文件所在的路径引起的跟其它文件（或文件夹）的路径关系。听起来有点绕，我们举例来说明一下，在`imooc`目录下新建一个`FileDemo2.java`文件，此时`imooc`目录的文件目录树结构如下：

```java
└── imoooc
    ├── FileDemo2.java
    ├── Hello.java
    └── images
```

内容如下：

```java
import java.io.File;
import java.io.IOException;

public class FileDemo2 {
    public static void main(String[] args) throws IOException {
        // 传入目录相对路径
        File dir = new File(".\\images");
        File imoocDir = new File("..\\imooc");
        // 传入文件相对路径
        File file = new File(".\\Hello.java");
    }
}
```

上面代码的`File`构造方法中传入的就是相对路径，代码中的`.`表示当前目录，`..`表示上级目录。

> **Tips：** 我们在实例化 File 对象时，不会产生对磁盘的操作，因此即使传入的文件或目录不存在，代码也不会抛出异常。只有当调用 File 对象下的一些方法时，才会对磁盘进行操作。

File 对象下有 3 个表示路径的实例方法：

1. `String getPath()`：将抽象路径名转换为路径名字符串；
2. `String getAbsolute()`：返回此抽象路径名的绝对路径名字符串；
3. `String getCanonicalPath()`：返回此抽象路径名的**规范路径名**字符串。

我们可以调用这 3 个方法并打印结果，实例如下：

```java
import java.io.File;
import java.io.IOException;

public class FileDemo2 {
    public static void main(String[] args) throws IOException {
        // 传入目录相对路径
        File imagesDir = new File(".\\images");
        File imoocDir = new File("..\\imooc");
        // 传入文件相对路径
        File file = new File(".\\Hello.java");

        System.out.println("-- imagesDir ---");
        System.out.println(imagesDir.getPath());
        System.out.println(imagesDir.getAbsolutePath());
        System.out.println(imagesDir.getCanonicalPath());

        System.out.println("-- imoocDir ---");
        System.out.println(imoocDir.getPath());
        System.out.println(imoocDir.getAbsolutePath());
        System.out.println(imoocDir.getCanonicalPath());

        System.out.println("-- file ---");
        System.out.println(file.getPath());
        System.out.println(file.getAbsolutePath());
        System.out.println(file.getCanonicalPath());
    }
}
```

运行结果：

```java
-- imagesDir ---
.\images
C:\Users\Colorful\Desktop\imooc\.\images
C:\Users\Colorful\Desktop\imooc\images
-- imoocDir ---
..\imooc
C:\Users\Colorful\Desktop\imooc\..\imooc
C:\Users\Colorful\Desktop\imooc
-- file ---
.\Hello.java
C:\Users\Colorful\Desktop\imooc\.\Hello.java
C:\Users\Colorful\Desktop\imooc\Hello.java
```

通过运行结果可以看出，规范路径名就是把`.`和`..`转换为标准的绝对路径。

### 4.3 判断对象是文件还是目录

我们可以通过如下两个方法判断 File 对象是文件还是目录：

1. `boolean isFile()`：测试此抽象路径名表示的文件是否为普通文件；
2. `boolean isDirectory()`：测试此抽象路径名表示的文件是否为目录。

实例如下：

```java
import java.io.File;

public class FileDemo3 {

    public static void printResult(File file) {
        // 调用isFile()方法并接收布尔类型结果
        boolean isFile = file.isFile();
        String result1 = isFile ? "是已存在文件" : "不是已存在文件";
        // 掉用isDirectory()方法并接收布尔类型而己过
        boolean directory = file.isDirectory();
        String result2 = directory ? "是已存在目录" : "不是已存在目录";
        // 打印该file对象是否是已存在文件/目录的字符串结果
        System.out.print(file);
        System.out.print('\t' + result1 + '\t');
        System.out.println(result2);
    }

    public static void main(String[] args) {
        // 传入目录绝对路径
        File dir = new File("C:\\Users\\Colorful\\Desktop\\imooc\\images");
        // 传入文件绝对路径
        File file = new File("C:\\Users\\Colorful\\Desktop\\imooc\\test.java");
        FileDemo3.printResult(dir);
        FileDemo3.printResult(file);
    }
}
```

运行结果：

```java
C:\Users\Colorful\Desktop\imooc\images	不是已存在文件	是已存在目录
C:\Users\Colorful\Desktop\imooc\test.java	不是已存在文件	不是已存在目录
```

代码中我们封装了一个静态方法`printResult()`，此方法打印 File 对象是否是文件 / 目录。值得注意的是，我们的磁盘中不存在`C:\Users\Colorful\Desktop\imooc\test.java`，因此无论调用`isFile()`方法还是`isDirectory()`方法，其返回结果都为`false`。

### 4.4 创建和删除目录

#### 4.4.1 创建目录

对于一个不存在的目录，我们可以使用`boolean mkdir()`方法创建一个目录。例如，我们想在`imooc`目录下创建一个`codes`目录，就可以使用该方法编写一段创建目录的代码。

实例如下：

```java
import java.io.File;

public class FileDemo4 {

    public static void main(String[] args) {
        // 传入目录绝对路径
        File dir = new File("C:\\Users\\Colorful\\Desktop\\imooc\\codes");
        if (!dir.exists()) {
            // 调用 mkdir() 方法
            boolean result = dir.mkdir();
            if (result) {
                System.out.println("目录创建成功");
            }
        }
    }
}
```

代码中我们调用了`File`对象的`boolean exists()`方法，此方法用于测试由此抽象路径名表示的文件或目录是否存在。当不存在时，我们才去创建目录。

运行代码前，`imooc`文件目录树结构如下：

```java
└── imoooc
    ├── FileDemo2.java
    ├── Hello.java
    └── images
```

运行结果：

```java
目录创建成功
```

运行代码后，`imooc`目录下多了一个`codes`目录，树结构如下：

```java
└── imoooc
    ├── FileDemo2.java
    ├── Hello.java
    ├── images
    └── codes
```

另外，File 类也提供了一个`boolean mkdirs()`方法，用来创建由这个抽象路径名命名的目录，包括任何必要但不存在的父目录。实际上是在递归执行`mkdir()`方法。

#### 4.4.2 删除目录

如果我们想要删除刚刚创建的`codes`目录，可以调用`boolean delete()`方法，实例如下：

```java
import java.io.File;

public class FileDemo5 {
    public static void main(String[] args) {
        // 传入目录绝对路径
        File dir = new File("C:\\Users\\Colorful\\Desktop\\imooc\\codes");
        if (dir.exists()) {
            // 调用 delete() 方法
            boolean deleted = dir.delete();
            if (deleted) {
                System.out.println("删除目录成功");
            }
        }
    }
}
```

运行代码前，`imooc`文件目录树结构如下：

```java
└── imoooc
    ├── FileDemo2.java
    ├── Hello.java
    ├── images
    └── codes
```

运行结果：

```java
删除目录成功
```

运行代码后，树结构如下：

```java
└── imoooc
    ├── FileDemo2.java
    ├── Hello.java
    └── images
```

### 4.5 创建和删除文件

对于文件类型的`File`对象，可以通过`boolean createNewFile()`方法创建一个新文件，使用`boolean delete()`方法删除文件。其调用方法和创建 / 删除目录相同，此处不再赘述。

关于更多`File`对象的操作，可翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/io/File.html)。

## 5. InputStream 抽象类

### 5.1 概述

`java.io.InputStream`抽象类是 Java 提供的最基本的输入流，它是所有输入流的父类。其最常用的抽象方法`int read()`签名如下：

```java
public abstract int read() throws IOException;
```

这个方法用于读取输入流的下一个字节，返回的`int`如果为`-1`，则表示已经读取到文件末尾。

`InputStream`与其子类的 UML 图如下所示：

![](https://xushuhui.gitee.io/image/imooc/5eddad4e09265fd007180648.jpg)

### 5.2 FileInputStream 实现类

我们将以最常用的`FileInputStream`实现类为例进行学习。其他实现类大同小异，如有需要可翻阅官方文档。

`FileInputStream`就是从文件流中读取数据，我们在`imooc`目录下新建一个文本文档`Hello.txt`，并输入如下内容：

![](https://xushuhui.gitee.io/image/imooc/5edda6e1091ba8c121451333.jpg)

读取`Hello.txt`文件中数据的实例代码如下：

```java
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

public class FileInputStreamDemo1 {

    public static void main(String[] args) throws IOException {
        // 实例化文件流
        FileInputStream fileInputStream = new FileInputStream("C:\\Users\\Colorful\\Desktop\\imooc\\Hello.txt");
        for (;;) {
            int n = fileInputStream.read();
            if (n == -1) {
                // read() 方法返回-1 则跳出循环
                break;
            }
            // 将n强制转换为 char 类型
            System.out.print((char) n);
        }
        // 关闭文件流
        fileInputStream.close();
    }
}
```

运行结果：

```java
Hello imooc!
```

如果我们打开了一个文件并进行操作，不要忘记使用`close()`方法来及时关闭。这样可以让系统释放资源。

## 6. OutputStream 抽象类

### 6.1 概述

`OutPutStream`抽象类是与`InputStream`对应的最基本的输出流，它是所有输出流的父类。其最常用的抽象方法`void write(int b)`签名如下：

```java
public abstract void write(int b) throws IOException;
```

这个方法用于写入一个字节到输出流。

`OutputStream`与其子类的 UML 图如下所示：

![](https://xushuhui.gitee.io/image/imooc/5eddad86091b120807520466.jpg)

### 6.2 FileOutputStream 实现类

我们同样以最常用的`FileOutputStream`实现类为例进行学习。其他实现类大同小异，如有需要可翻阅官方文档。

`FileOutputStream`就是向文件流中写入数据，下面我们向`imooc`目录下的文本文档`Hello.txt`输入一段字符串`HHH`。完整实例如下：

```java
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

public class FileOutputStreamDemo1 {
    public static void main(String[] args) throws IOException {
        FileOutputStream fileOutputStream = new FileOutputStream("C:\\Users\\Colorful\\Desktop\\imooc\\Hello.txt");
        // 写入 3 个H字符
        fileOutputStream.write(72);
        fileOutputStream.write(72);
        fileOutputStream.write(72);
        fileOutputStream.close();
    }
}
```

运行代码后，`Hello.txt`后面成功写入了 3 个字符`H`。

## 7. 小结

通过本小节的学习，我们知道了什么是输入输出流的概念，输入输出流经常用于上传文件到服务器的场景。想要通过 Java 操作文件和目录，要学会使用`java.io.File`类，`InputStream`和`OutputStream`分别是所有输入流和所有输出流的父类，`FileInputStream`实现了文件流的输入，`FileOutputStream`实现了文件流的输出。还有很多其它实现类我们没有介绍到，但使用方法大同小异，希望同学可以在用到时自行查阅文档来学习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
