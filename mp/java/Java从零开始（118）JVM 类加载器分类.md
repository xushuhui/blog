---
title: Java从零开始（118）JVM 类加载器分类
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# JVM 类加载器分类


## 1. 前言

我们之前对类加载子系统进行过简要的介绍，此处我们将会进行更加细致的讲解。本节主要知识点如下：

* 启动（Bootstrap）类加载器的作用及代码验证，为本节重点内容之一；
* 扩展（Extension）类加载器的作用及代码验证，为本节重点内容之一；
* 系统（System Application）类加载器的作用及代码验证，为本节重点内容之一。

通篇皆为重点内容，都是学习者需要重点掌握的。并且此节的内容也是后续内容的知识基础，为了更顺利的进行学习，次节内容需要重点掌握。

## 2. 类加载子系统知识回顾

我们在 JVM 总体架构的讲解过程中，提到过类加载子系统的工作流程分为三步：加载 ->链接 ->初始化。如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxq35l9j60jg070ta702)

本节我们所讨论的内容都是围绕第一步“加载（Loading）” 进行的。对于链接和初始化，我们会在后边的章节进行讲解。

我们将加载（Loading）这一步，再进行下细致的模块划分，如下图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnxqfgplj60jg0ezgom02)

从上图中我们可看到，加载（loading）这一步，里边包含了三个更加细粒度的模块，分别为 **BootStrap Class Loader**，**Extention Class Loader** 和 **Application Class Loader**，这三个 Class Loader 就是我们加载过程中必须要使用到的三大类加载器。

## 3. 启动（Bootstrap）类加载器

**定义**：启动（Bootstrap）类加载器也称为引导类加载器，该加载器是用本地代码实现的类加载器，它所加载的类库绝大多数都是出自 %JAVA_HOME%/lib 下面的核心类库，当然还有其他少部分所需类库。

由于引导类加载器涉及到虚拟机本地实现细节，开发者无法直接获取到启动类加载器的引用，所以不允许直接通过引用进行操作。

> **Tips**：从上述定义的描述中，我们可以看到一个特别需要关注的点：启动类加载器加载的绝对大多数是 %JAVA_HOME%/lib 下边的核心类库。这句话完完全全的体现出了启动（Bootstrap）类加载器存在的意义。对于其他少部分核心类的加载，我们在代码验证过程中来讲解。接下来，让我们通过示例代码进行下验证。

**示例**：通过编写一个 main 函数，打印出通过启动（Bootstrap）类加载器加载的所有的类库信息，以证实启动（Bootstrap）类加载器加载的是 %JAVA_HOME%/lib 下边的核心类库。

> **Tips**：注意下 main 函数代码的第二行代码 `URL[] urls = sun.misc.Launcher.getBootstrapClassPath().getURLs();` 这是通过 sun 公司提供的 Launcher 包获取 Bootstrap 类加载器下 ClassPath 下的所有的 URL。

```java
import java.net.URL;

public class LoaderDemo {
    public static void main(String[] args) {
        System.out.println("BootstrapClassLoader 的加载路径: ");
        URL[] urls = sun.misc.Launcher.getBootstrapClassPath().getURLs();
        for(URL url : urls)
            System.out.println(url);
     }
}
```

**结果验证**：运行 main 函数。

> **Tips**：此处运行结果所打印的类库的绝对路径为本人本机的安装路径，学习者应按照自己真实的 JDK 安装路径以及版本对号入座，此处仅为示例。

```java
BootstrapClassLoader 的加载路径:
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/resources.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/rt.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/sunrsasign.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jsse.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jce.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/charsets.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jfr.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/classes
```

**结果解析**：我们可以看到，运行结果中的前 7 个类库（不同的 JDK 版本会有差异，此处我们讨论的是 JDK 1.8 版本），都是出自 lib 下的核心类库。但是对于最后一条加载信息却不是 lib 下的类库。我们仔细看下最后这条信息的加载 file:/D:/Programs/Java/jdk1.8.0_111/jre/classes。

这就是前文我们所提到的其他少部分的核心类库加载，学习者可以根据自己真实的安装位置打开 /jre 文件夹，看看是否存在 /classes 路径。结果是 /classes 文件夹路径并不存在，除非我们进行特殊的参数创建才可以出现 /classes 路径。此处并非我们主要讨论的问题，我们关注的是 lib 文件夹下的核心类库加载，这里仅做了解即可。

## 4. 扩展（Extension）类加载器

**定义**：扩展类加载器是由 Sun 公司提供的 ExtClassLoader（sun.misc.Launcher$ExtClassLoader）实现的，它负责将 %JAVA_HOME%/lib/ext 或者少数由系统变量 -Djava.ext.dir 指定位置中的类库加载到内存中。开发者可以直接使用标准扩展类加载器。

> **Tips**：此处我们依旧对大多数的核心类库加载位置进行讨论，即 %JAVA_HOME%/lib/ext 文件夹下的扩展核心类库。对于系统变量指定的类库，稍作了解即可。下边进行示例代码验证

**示例**：

```java
import java.net.URL;
import java.net.URLClassLoader;
public class LoaderDemo {
    public static void main(String[] args) {
        //取得扩展类加载器
        URLClassLoader extClassLoader = (URLClassLoader)ClassLoader.getSystemClassLoader().getParent();
        System.out.println(extClassLoader);
        System.out.println("扩展类加载器 的加载路径: ");
        URL[] urls = extClassLoader.getURLs();
        for(URL url : urls)
            System.out.println(url);
     }
}
```

**结果验证**：运行 main 函数。

```java
扩展类加载器 的加载路径:
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/access-bridge-64.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/cldrdata.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/dnsns.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/jaccess.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/jfxrt.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/localedata.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/nashorn.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunec.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunjce_provider.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunmscapi.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunpkcs11.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/zipfs.jar
```

**结果解析**：我们可以看到，运行结果中所有的核心类库均来自 %JAVA_HOME%/lib/ext 的文件夹。

## 5. 系统（System Application）类加载器

**定义**：系统类加载器是由 Sun 公司提供的 AppClassLoader（sun.misc.Launcher$AppClassLoader）实现的，它负责将 用户类路径（java -classpath 或 -Djava.class.path 变量所指的目录，即当前类所在路径及其引用的第三方类库的路径）下的类库加载到内存中。开发者可以直接使用系统类加载器。

> **Tips**：系统（System Application）类加载器加载的核心类库类型比较多，也会加载 lib 下的未被 BootStrap 类加载器加载的类库，还会加载 ext 文件夹下的未被 Extension 类加载器加载的类库，以及其他类库。总而言之一句话，加载除了 BootStrap 类加载器和 Extension 类加载器所加载的其余的所有的核心类库。

**示例**：

```java
import java.net.URL;
import java.net.URLClassLoader;
public class LoaderDemo {
    public static void main(String[] args) {
        //取得应用(系统)类加载器
        URLClassLoader appClassLoader = (URLClassLoader)ClassLoader.getSystemClassLoader();
        System.out.println(appClassLoader);
        System.out.println("应用(系统)类加载器 的加载路径: ");
        URL[] urls = appClassLoader.getURLs();
        for(URL url : urls)
            System.out.println(url);
     }
}
```

**结果验证**：运行 main 函数。

```java
应用(系统)类加载器 的加载路径:
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/charsets.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/deploy.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/access-bridge-64.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/cldrdata.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/dnsns.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/jaccess.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/jfxrt.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/localedata.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/nashorn.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunec.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunjce_provider.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunmscapi.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/sunpkcs11.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/ext/zipfs.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/javaws.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jce.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jfr.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jfxswt.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/jsse.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/management-agent.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/plugin.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/resources.jar
file:/D:/Programs/Java/jdk1.8.0_111/jre/lib/rt.jar
file:/E:/IdeaWorkspace/LeeCode/target/classes/
file:/D:/Programs/IntelliJ%20IDEA%20Educational%20Edition%202019.3.1/lib/idea_rt.jar
```

**结果解析**：我们可以看到， 系统（System Application）类加载器加载的类库种类很多，除了之前两种类加载器加载的类库，其余必须的核心类库，都由系统类加载器加载。

## 6. 小结

对于类加载器中的第一步加载（Loading），我们主要讲解了 3 种类加载器。并且对不同的类加载器所加载的类库进行了讲解以及代码验证。通篇皆为重点知识，需要学习者用心学习。

对于加载（Loading）这一步，我们还未讲解完，下节课程会讲解加载（Loading）这一步所遵循的双亲委派模型，本节作为下一节的知识基础，更需要着重理解、掌握。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
