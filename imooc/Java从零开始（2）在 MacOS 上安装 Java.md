# 在 MacOS 上安装 Java

本小节我们将介绍如何在 MacOS 平台安装 Java 。

如果你想在其他平台安装 Java，请查看对应平台的安装教程：

* [在 Windows 上安装 Java](https://www.imooc.com/wiki/javalesson/installationwindows.html)

* [在 Linux 上安装 Java](https://www.imooc.com/wiki/javalesson/installationlinux.html)

## 1. 下载安装包

我们首先打开[Oracle官网的 JDK 下载地址](https://www.oracle.com/java/technologies/javase-downloads.html#JDK15)，找到 Java SE 15 版块，点击 `JDK Download` 按钮。

![](//img.mukewang.com/wiki/5f8fd5be092514ce15290797.jpg)

点击 `JDK Download`按钮后我们会跳转到 JDK 下载详情页面。

我们在下载详情页面可以找到如下图这样的一个表格，在最右侧 Download 一列中找到`jdk-15.0.1_osx-x64_bin.dmg`一项，单击鼠标左键。

![](//img.mukewang.com/wiki/5f8ff85c09096d1028801682.jpg)

此时网页上会弹出如下对话框，提示如果你想要下载必须遵守其协议，**先勾选上复选框**，**再点击下载按钮**即可开始下载。

![](//img.mukewang.com/wiki/5f8fd5e7099310df15510501.jpg)

## 2. 安装到本机

下载好安装包后，打开安装包。

![](//img.mukewang.com/wiki/5f8fff1209d2470d12400876.jpg)

MacOS 平台的安装流程非常简单，点击**继续 -> 安装 -> 输入本机密码** ，等待安装成功。

![](//img.mukewang.com/wiki/5f8fff53098bcab512400876.jpg)

点击安装后，会提示输入用户密码：

![](//img.mukewang.com/wiki/5f8fff940976984b08860458.jpg)

安装成功，点击关闭即可：

![](//img.mukewang.com/wiki/5f8fffbc090f525512400876.jpg)

如下视频演示了整个安装过程：

## 3. 配置环境变量

按照上面的操作，我们已经在 MacOS 上成功安装了 JDK 15 ，接下来我们需要配置一个 `JAVA_HOME`环境变量，来指向 Java 的安装目录，并且将`JAVA_HOME`的`bin`目录附加到系统变量的`PATH`上， 其目的是为了我们在任何目录位置都可以执行`java` 命令。

### 3.1 打开终端

打开 Mac 终端应用，在终端输入如下命令可以查看`Java 15`的安装目录：

```java
$ /usr/libexec/java_home -v 15
/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home
```

请记住这里的安装目录：`/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home`，下面将会用到。

### 3.2 编辑启动脚本

MacOS 默认的 shell 是`bash`，**启动脚本**是 `~/.bash_profile`, 如果你的 shell 和我一样是`zsh`，那么启动脚本就是 `~/.zshrc`, 以`bash`为例，使用`vim`编辑器编辑启动脚本：

```java
vim ~/.bash_profile
```

输入字母`i`切换到输入模式，在启动脚本下添加如下两行命令：

```java
export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home
export PATH=$JAVA_HOME/bin:$PATH
```

第一行命令是设置一个名为 `JAVA_HOME`的环境变量，它指向 Java 的安装目录。

第二行命令是将 `JAVA_HOME` 的`bin`目录附加到系统变量的 `PATH`上，这样，bin 目录下的很多可执行文件就被系统加载了。

最后保存并退出启动脚本，切换到底线命令模式（敲击`esc`按键 ，输入 `:`），输入`wq`，敲击回车按键。

### 3.3 加载环境变量

为了让我们刚刚添加的环境变量生效，使用`source`命令加载环境变量：

```java
source ~/.bash_profile
```

### 3.4 验证环境变量

那么如何验证上述一系列操作是否成功呢？

打开终端，键入 `java -version` 命令，看到如下输入，即证明你已经成功配置好了环境变量。

![](//img.mukewang.com/wiki/5f9000aa096e7bf314060230.jpg)

## 4. 小结

本节我们在`MacOS`系统上完成了`Java`的下载、安装以及环境变量的配置。这个过程中，还要再次强调几点：

* 在官网下载页面，要先勾选上`接受许可协议`，再点击下载按钮才能开始下载。
* MacOS 是类 `UNIX`的操作系统，了解`vim`编辑器的基本使用是很有必要的。
* 配置环境变量是为了我们在任何目录位置都可以执行`java` 命令。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
