# Java 集成开发环境 - Eclipse

本小节我们将介绍如何在我们的电脑上安装并配置开发工具：`Eclipse IDE`

如果你想查看如何安装配置 `IntelliJ IDEA`，请点击[此处](https://www.imooc.com/wiki/javalesson/idea.html)查看

## 1. IDE 概述

### 1.1 IDE 是什么？

IDE 即 `Integrated Development Environment` 的缩写，中文意为集成开发环境，是用于提供程序开发环境的应用程序，一般包括代码编辑器、编译器、调试器和图形用户界面等工具。集成了代码编写功能、分析功能、编译功能、调试功能等一体化的开发软件服务套件。

### 1.2 Ecplise IDE 简介

`Eclipse IDE` 是一款由 **Eclipse 基金会**开发的跨平台开源集成开发环境，该软件可以用来 Java 程序的集成开发。可大大提高我们的开发效率。

## 2. Eclipse IDE 下载

打开 [Eclipse 官网下载地址](https://www.eclipse.org/downloads/) ，点击 **Download Packages**。

![](https://xushuhui.gitee.io/image/imooc/5e987bb509ddbacc03240267.jpg)

找到 [Eclipse IDE for Java Developers](https://www.eclipse.org/downloads/packages/release/2019-12/r/eclipse-ide-java-developers)，这个包是任何 Java 开发者都可以使用的一个**基本工具**。点击右侧对应平台的下载链接即可开始下载。

![](https://xushuhui.gitee.io/image/imooc/5e987bce09d3422408140130.jpg)

## 3. Eclipse IDE 安装和使用

### 3.1 安装到本机

下载完成后，打开安装包。

如果是 Mac OS 平台，打开后会出现下面一个目录。直接将 `Eclipse`拖动到旁边的 `Applications` 目录即可完成安装。

![](https://xushuhui.gitee.io/image/imooc/5e987bec090f25f707700436.jpg)

如果是 Windows 平台，与常用软件的安装步骤相同，这里不再赘述。

### 3.2 快速编写 Hello World 程序

打开 IDE，会提示我们选择**工作目录**，**工作目录用于存放开发者的配置和开发工件**。选择好自己的工作目录后直接点击 `Launch`按钮即可打开 IDE。

![](https://xushuhui.gitee.io/image/imooc/5e987c0609ebe2f606050244.jpg)

打开后会显示一个 ecplise 的欢迎页面。点击 **Create a Hello World application**。按照向导步骤即可创建并运行我们的 Hello World 程序。

![](https://xushuhui.gitee.io/image/imooc/5e987c20099db5af10240768.jpg)

右侧区域为创建一个`Hello World`程序的整体步骤。

![](https://xushuhui.gitee.io/image/imooc/5e987c3709ffd79b14400900.jpg)

下面我们一起动手开始吧。

#### 第一步：新建一个 Java 工程

点击 New -> Java Project

![](https://xushuhui.gitee.io/image/imooc/5e987dbe0985ca9b08090559.jpg)

下图中`Project name`为项目名称。将项目命名为 `hello`（项目名的命名规范：全小写英文），并选择 JRE 为 JavaSE-14，点击 Finish 按钮。

![](https://xushuhui.gitee.io/image/imooc/5e987c740991a55708120721.jpg)

这时会出现一个是否创建模块信息的弹窗，点击 `Don't Create` 按钮。

![](https://xushuhui.gitee.io/image/imooc/5e987c8b099ed11d05870361.jpg)

#### 第二步：在工程下新建一个类

在刚刚创建好的工程下的 src 目录上，点击右键 -> New -> Class

![](https://xushuhui.gitee.io/image/imooc/5e987ca00988fe6605600656.jpg)

将类名命名为 HelloWorld，并且选择**创建 main 方法**的复选框。点击 `Finish` 按钮

![](https://xushuhui.gitee.io/image/imooc/5e987cb7095acee105900632.jpg)

#### 第三步：编写打印语句

现在，ecplise 已经为我们在 HelloWorld.java 文件中自动创建了 main 方法，在 main 方法下增加如下打印语句：

```java
System.out.println("Hello World");
```

![](https://xushuhui.gitee.io/image/imooc/5e987cdb09308ab204180173.jpg)

#### 第四步：运行 Java 应用

在源代码文件 HelloWorld.java 上点击右键（或直接在源代码中点击右键），选择 Run As -> Java Application

![](https://xushuhui.gitee.io/image/imooc/5e987cf809ebc7f407330750.jpg)

我们可以看到控制台已成功打印 Hello World

![](https://xushuhui.gitee.io/image/imooc/5e987e100988a73804800208.jpg)

## 4. 常用配置

### 4.1 调整字体、字号

打开 Eclipse -> Preferences

![](https://xushuhui.gitee.io/image/imooc/5e987d0f098930b902920242.jpg)

在左侧列表选择 General -> Appearance -> Colors and Fonts

![](https://xushuhui.gitee.io/image/imooc/5e987d270943799c06300587.jpg)

在右侧选择 Basic -> Text Font

![](https://xushuhui.gitee.io/image/imooc/5e987d38092f127806300587.jpg)

点击 `Edit` 按钮，可选择理想的字体、字号

![](https://xushuhui.gitee.io/image/imooc/5e987d4b0945ec9b04450270.jpg)

最后，点击 `Preferences` 窗口中的 `Apply and Close` 按钮即可保存配置。

### 4.2 设置代码编辑区的字符编码

打开 Eclipse -> Preferences

![](https://xushuhui.gitee.io/image/imooc/5e987d0f098930b902920242.jpg)

在左侧列表中点击 General -> Workspace， 在右侧的红框中的 Text file encoding 点击 Other 单选框，在其下拉列表可选择对应的字符编码。

![](https://xushuhui.gitee.io/image/imooc/5e987d7f0959221807490698.jpg)

点击 `Apply and Close` 按钮即可完成保存并关闭窗口。

## 5. 小结

本小节我们知道了什么是集成开发环境，利用好集成开发环境可以大大提高我们的工作效率。

我们对`Eclipse IDE`的下载和安装也做了详细介绍，并且使用它快速编写了一个 Java `Hello World`程序，从编码、编译到执行，相信你可以明显体会到`IDE`给我们带来的方便快捷。

我们也在最后介绍了`Eclipse`的常用配置，如果你还想了解更多的自定义配置，可以上网查阅相关资料。

如果你想要了解另外一款常用`IDE`:`IntelliJ IDEA`，请阅读[Java 集成开发环境 - IntelliJ IDEA](https://www.imooc.com/wiki/javalesson/idea.html)小节。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
