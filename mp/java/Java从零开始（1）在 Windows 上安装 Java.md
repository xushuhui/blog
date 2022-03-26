---
title: Java从零开始（1）在 Windows 上安装 Java
zhihu-url: https://zhuanlan.zhihu.com/p/393139333
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 在 Windows 上安装 Java


本小节我们将介绍如何在 Windows 平台安装 Java 。由于微软已正式终止对 Win7 操作系统的支持，作为新时代的程序员，我们使用 Win10 操作系统来进行安装演示。




## 1. 下载安装包

我们首先打开[Oracle官网的 JDK 下载地址](https://www.oracle.com/java/technologies/javase-downloads.html#JDK15)，找到 Java SE 15 版块，点击 `JDK Download` 按钮。

![](https://xushuhui.gitee.io/image/imooc/5f8fd5be092514ce15290797.jpg)

点击 `JDK Download`按钮后，我们会跳转到 JDK 下载详情页面。

我们在下载详情页面可以找到如下图这样的一个表格，在最右侧 Download 一列中找到`jdk-15.0.1_windows-x64_bin.exe`一项，单击鼠标左键。

![](https://xushuhui.gitee.io/image/imooc/5f8fc99809eb520f29341720.jpg)

此时网页上会弹出如下对话框，提示如果你想要下载必须遵守其协议，**先勾选上复选框**，**再点击下载按钮**即可开始下载。

![](https://xushuhui.gitee.io/image/imooc/5f8fd5e7099310df15510501.jpg)

## 2. 安装到本机

安装包下载完成后，打开安装包，开始安装。

与常用软件安装的流程基本相同。点击**下一步 -> （选择安装目录）下一步 -> 等待安装完成**即可完成安装。

![](https://xushuhui.gitee.io/image/imooc/5f8fd73209f01c1410010759.jpg)

选择安装目录，这里我们采用默认的 `C:\Program Files\Java\jdk-15.0.1` ，直接点击下一步。**请记住这里的安装目录，我们将在配置环境变量时用到**。

![](https://xushuhui.gitee.io/image/imooc/5f8fd7c5093a8d0b10010759.jpg)

点击下一步后开始安装，等待安装完成。

![](https://xushuhui.gitee.io/image/imooc/5f8fd98a09b2516610010759.jpg)

安装完成，点击关闭按钮。

![](https://xushuhui.gitee.io/image/imooc/5f8fd9cc09e4de8510010759.jpg)

## 3. 配置环境变量

完成了安装包的下载和本机安装，我们还有一个重要步骤，就是配置环境变量。其**目的**是为了让我们在任何目录位置都可以执行 java 命令。

### 3.1 打开高级系统设置

首先找到我的电脑，右键点击**属性**。再点击**高级系统设置**

![](https://xushuhui.gitee.io/image/imooc/5e8d48ec099cd29e04410499.jpg)

### 3.2 打开环境变量窗口

点击系统设置后，继续点击下面的**环境变量**按钮

![](https://xushuhui.gitee.io/image/imooc/5e8d490a0977da3c09501080.jpg)

### 3.3 添加环境变量

找到系统变量的**新建按钮**，点击添加一个新的环境变量。

![](https://xushuhui.gitee.io/image/imooc/5e8d492d0987f5f912361203.jpg)

为其变量名设置为 `JAVA_HOME`，变量值设置为我们刚刚的 Java 安装目录：`C:\Program Files\Java\jdk-15.0.1`，点击确定保存环境变量。

![](https://xushuhui.gitee.io/image/imooc/5f8fda0409e2062f13060335.jpg)

### 3.4 编辑系统变量 Path

最后，将我们刚刚添加的 `JAVA_HOME` 的 `bin` 目录附加到系统变量 PATH 上。

首先找到系统变量列表中的 `Path` 变量，选中后点击编辑按钮。

![](https://xushuhui.gitee.io/image/imooc/5e8d49900989529a12361203.jpg)

新建一个值为 `%JAVA_HOME%\bin`的变量，这里 `%JAVA_HOME%` 的写法表示调用我们刚刚定义的环境变量`JAVA_HOME`。也就是说，这里我们新建的变量实际上就是`C:\Program Files\Java\jdk-15.0.1\bin`目录，这样写是为了更直观，容易维护，如果你需要修改 Java 的安装目录，直接编辑`JAVA_HOME`变量即可完成修改。

![](https://xushuhui.gitee.io/image/imooc/5e8d49c1098cefdf10541029.jpg)

### 3.5 验证环境变量

完成了上述一系列配置，我们如何验证刚刚的配置是否成功呢？

首先，在 Windows 搜索栏处输入`cmd`找到命令提示符应用，再点击打开：

![](https://xushuhui.gitee.io/image/imooc/5e8d49ed097c8a0115611362.jpg)

输入`java -version`命令。显示如下输出，即证明你已经成功完成了环境变量的配置。

![](https://xushuhui.gitee.io/image/imooc/5f8fda25097af9b519581050.jpg)

## 4. 小结

本节我们在`Windows10`系统上完成了`Java`的下载、安装以及环境变量的配置。这个过程中，还要再次强调几点：

* 在官网下载页面，要先勾选上`接受许可协议`，再点击下载地址才能开始下载。
* 安装过程很简单，但请不要忘记`Java`的安装目录。
* 明确配置环境变量的目的。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
