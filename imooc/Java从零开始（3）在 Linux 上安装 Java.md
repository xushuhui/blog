# 在 Linux 上安装 Java

本小节我们将介绍如何在 Linux 平台安装 Java 。我们选用最常用的 Linux 发行版本 `CentOS` 来进行（注：版本号为`CentOS 7.6`）。

如果你想在其他平台安装 Java，请查看对应平台的安装教程：

* [在 Windows 上安装 Java](https://www.imooc.com/wiki/javalesson/installationwindows.html)

* [在 MacOS 上安装 Java](https://www.imooc.com/wiki/javalesson/installationmacos.html)

## 1. 下载安装包

我们首先打开[Oracle官网的 JDK 下载地址](https://www.oracle.com/java/technologies/javase-downloads.html#JDK15)，找到 Java SE 15 版块，点击 `JDK Download` 按钮。

![](//img.mukewang.com/wiki/5f8fd5be092514ce15290797.jpg)

点击 `JDK Download`按钮后，我们会跳转到 JDK 下载详情页面。

在下载详情页面可以找到如下图这样的一个表格，在最右侧 Download 一列中找到`jdk-15.0.1_linux-x64_bin.rpm`一项，单击鼠标左键。

此时网页上会弹出如下对话框，提示如果你想要下载必须遵守其协议，**先勾选上复选框**，**再使用鼠标右键点击下载按钮，复制链接地址**。

![](//img.mukewang.com/wiki/5e8d70fb0a85d59e25681144.jpg)

> **Tips：**
>
> 本小节的操作动图中，使用的 Java 版本为 14，由于安装 Java15 和 安装 Java14 的操作流程完全相同，我并没有进行统一替换。

登录至我们要安装 Java 的 Linux 主机，使用 `wget` 命令来下载我们刚刚复制的 JDK 链接地址。执行如下命令，将 JDK 下载至服务器：

```java
wget --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" https://download.oracle.com/otn-pub/java/jdk/14+36/076bab302c7b4508975440c56f6cc26a/jdk-14_linux-x64_bin.rpm
```

> **Tips:** 如果你的主机没有安装 `wget` 命令，执行上述命令会报错：
>
> ```java
> -bash: wget: command not found
> ```
>
> 这是因为主机还没有安装`wget`命令，执行 `yum install wget` 命令来进行安装。成功安装后再执行下载命令。

下载过程如下：

![](//img.mukewang.com/wiki/5e8d8a070a593af415621426.jpg)

使用 `ls` 命令查看当前目录下的内容，`jdk-14_linux-x64_bin.rpm`就是我们刚刚下载好的安装包。

```java
[root@Colorful ~]# ls
jdk-14_linux-x64_bin.rpm
```

## 2. 安装到本机

使用 `rpm` 命令安装刚刚下载好的安装包。执行命令：

```java
rpm -ivh jdk-14_linux-x64_bin.rpm
```

安装过程如下：

![](//img.mukewang.com/wiki/5e8d8a3a0aa5049915621426.jpg)

安装成功后，输入 `java -version`来验证是否安装成功：

```java
[root@Colorful ~]# java -version
java version "14" 2020-03-17
Java(TM) SE Runtime Environment (build 14+36-1461)
Java HotSpot(TM) 64-Bit Server VM (build 14+36-1461, mixed mode, sharing)
```

屏幕输出了如上内容，表示我们已经成功在 Linux 主机上安装了 Java。

## 3. 配置环境变量

按照上面的操作，我们已经在 Linux 上成功安装了 JDK 14 ，接下来我们需要配置一个 `JAVA_HOME`环境变量，来指向 Java 的安装目录，并且将`JAVA_HOME`的`bin`目录附加到系统变量的`PATH`上， 其目的是为了我们在任何目录位置都可以执行 java 命令。

Java 的默认安装目录为 `/usr/java/jdk-14`，编辑启动脚本 `~/.bash_profile`，在启动脚本下添加如下两行命令

```java
export JAVA_HOME=/usr/java/jdk-14
export PATH=$JAVA_HOME/bin:$PATH
```

为了让刚刚在启动脚本添加的环境变量生效，执行 `source` 命令：

```java
source ~/.bash_profile
```

最后，打印一下 PATH 系统变量，查看环境变量是否正确添加：

```java
[root@Colorful ~]# echo $PATH
/usr/java/jdk-14/bin:/usr/java/jdk-13.0.2/bin:/usr/local/node/8.11.1/bin:/usr/local/node/8.9.3/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin:/root/bin
```

## 4. 卸载 Java

`Linux`上的卸载与其他操作系统有所不同。有必要单独说明一下。

执行如下命令，可查看操作系统中的 `JDK`版本

```java
rpm -qa | grep jdk
```

![](//img.mukewang.com/wiki/5e8d8a97091efd5507430133.jpg)

可以看到我们的系统上存在两个版本的 Java，执行命令删除对应版本：

```java
yum -y remove java jdk-13.0.2-13.0.2-ga.x86_64
```

![](//img.mukewang.com/wiki/5e8d8a73098dab7315721460.jpg)

## 5. 小结

本小节中，我们在`Linux`操作系统上的完成了`Java`的下载、安装、环境变量的配置以及卸载。当然还要再次提醒几个需要注意的点：

* 复制下载链接地址前，请记住要勾选`接受许可协议`，否则无法复制正确的地址。
* 建议直接复制粘贴对应命令以确保准确性。
* 如果你想在其他`Linux`发行版本上安装 Java，只要掌握步骤要领，执行对应安装命令即可。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
