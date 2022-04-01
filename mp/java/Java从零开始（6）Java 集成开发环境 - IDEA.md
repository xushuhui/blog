---
title: Java从零开始（6）Java 集成开发环境 - IDEA
zhihu-url: https://zhuanlan.zhihu.com/p/399979538
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 集成开发环境 - IntelliJ IDEA


本小节我们将介绍如何在我们的电脑上安装并配置开发工具：`IntelliJ IDEA`

## 1. IDE 概述

### 1.1 IDE 是什么？

IDE 即 `Integrated Development Environment` 的缩写，中文意为集成开发环境，是用于提供程序开发环境的应用程序，一般包括代码编辑器、编译器、调试器和图形用户界面等工具。集成了代码编写功能、分析功能、编译功能、调试功能等一体化的开发软件服务套。

### 1.2 IntelliJ IDEA 简介

IntelliJ IDEA 是由 **JetBrains** 公司开发的高效智能的 Java 集成开发工具，在业界被公认为最好的 java 开发工具，JetBrains 针对个人开发者及企业组织提供不同的授权方式。由于其优越的开发体验，近些年来得到越来越多个人开发者及企业的青睐，将其作为首选的 Java 开发工具。

## 2. IntelliJ IDEA 下载

打开[Intellij IDEA 官网下载地址](https://www.jetbrains.com/idea/download)，网页会自动识别你的当前机器的操作系统。这里分为两个版本，左侧的 **Ultimate** 版为Web和企业开发提供，需要商业授权。右侧 **Community** 版为 JVM 和安卓开发提供，免费且开源。这里我们选择免费的社区版本下载即可。点击 Download 按钮开始下载。

![](https://xushuhui.gitee.io/image/imooc/5e8dca240971d9a607980361.jpg)

## 3. IntelliJ IDEA 安装和使用

### 3.1 安装到本机

打开我们下载好的安装包

如果是 Mac OS 平台，打开后出现如下的窗口，直接将`IntelliJ IDEA CE`拖动到 `Applications`目录即可完成安装。

![](https://xushuhui.gitee.io/image/imooc/5e8dca420953b9a304550318.jpg)

如果是 Windows 平台，与常用软件的安装步骤相同，这里不再赘述。

### 3.2 快速编写 Hello World 程序

打开 IntelliJ IDEA，完成一些初始配置。（建议初次使用的开发者，一直点击下一步采用默认配置即可）。完成初始配置后会出现如下的欢迎窗口。

![](https://xushuhui.gitee.io/image/imooc/5e8dca6309c23e6a06660482.jpg)

#### 第一步：新建一个 Java 工程

点击 Create New Project。选择项目的 SDK。这里会自动识别我们之前所安装的 JDK14。点击 `Next` 按钮。

![](https://xushuhui.gitee.io/image/imooc/5e8dca7f090b7f2b12120755.jpg)

选中 Create project from template -> Command Line App，表示会新建一个包含 `main()` 方法的简单 Java 应用。点击 `Next` 按钮继续。

![](https://xushuhui.gitee.io/image/imooc/5e8dcf460900665512120755.jpg)

设置项目名称（`Project name`）为 `hello`，`Project location` 为项目的存放目录，`Base package`是包名，自定义即可。点击 `Finish` 按钮。

![](https://xushuhui.gitee.io/image/imooc/5e8dca9a094acc5812120755.jpg)

#### 第二步：编写输出语句

新建项目成功后，IDE 会自动打开项目。点击左侧的 Project 按钮，即可查看项目的目录结构，`Main.java` 为 IDE 为我们自动创建的模板代码。

![](https://xushuhui.gitee.io/image/imooc/5e8dcab10940edbe05990311.jpg)

在 Main.java 的 main() 方法中，编写如下输出语句：

```java
System.out.println("Hello World!");
```

![](https://xushuhui.gitee.io/image/imooc/5e8dcacd092710a409610367.jpg)

#### 第三步：运行 Java 应用

点击 main() 方法左侧的绿色小箭头，会弹出一些选项，点击 `Run 'Main.main()'`。

![](https://xushuhui.gitee.io/image/imooc/5e8dcae409ce52c003890222.jpg)

IDE 会自动编译并执行 Java 应用，稍等片刻后，在下方的控制台中会输出 `Hello World!`。

![](https://xushuhui.gitee.io/image/imooc/5e8dcb0009a421b514400900.jpg)

想要执行源码，除了上述点击绿色小箭头，也可以在源代码文件中点击鼠标右键，选择`Run Main`来执行源代码：

![](https://xushuhui.gitee.io/image/imooc/5e8dcb2d09ef287c05450537.jpg)

## 4. 常用配置

### 4.1 调整字体、字号

打开 IntelliJ IDEA -> Preferences

![](https://xushuhui.gitee.io/image/imooc/5e8dcb860972e23202620242.jpg)

在左侧列表选择 Editor -> Font，右侧会出现设置字体和字号的选项，选择理想的字体、字号。点击 `OK` 按钮即可保存设置并关闭窗口。

![](https://xushuhui.gitee.io/image/imooc/5e8dcb4909c3984f09820722.jpg)

### 4.2 设置代码编辑区的字符编码

点击 IDE 右下角的 `UTF-8`即可弹出字符编码选项。选择对应的编码即可。

![](https://xushuhui.gitee.io/image/imooc/5e8dcba209d9db6a14400900.jpg)

## 5. 小结

本小节我们知道了什么是集成开发环境，利用好集成开发环境可以大大提高我们的工作效率。

我们对`IntelliJ IDEA`的下载和安装也做了详细介绍。当然，我们安装`IDE`的主要目的是为了学习 Java，推荐使用开源免费并且更加轻量的`Community`版本即可。

本小节我们使用`IDE`快速编写了一个 Java `Hello World`程序，从编码、编译到执行，相信你可以明显体会到`IDE`给我们带来的方便快捷。

我们也在最后介绍了`IntelliJ IDEA`的常用配置，如果你还想了解更多的自定义配置，可以上网查阅相关资料。

如果你想要安装使用`Ecplice IDE`，请阅读 [Java 集成开发环境 - Ecplice](https://www.imooc.com/wiki/javalesson/eclipse.html)小节。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
