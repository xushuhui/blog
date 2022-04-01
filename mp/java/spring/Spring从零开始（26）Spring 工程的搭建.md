# Spring 工程的搭建

## 1. 前言

**“Spring 的工程如何创建？”**

在上一章节中我们通过 Spring 的简介，了解了 Spring 的概念、体系结构、与它的核心功能。那么本章带你体验一下 Spring 的项目开发和我们之前搭建过的开发项目有哪些不同。

## 2. 课程进入

### 2.1 Spring 框架版本介绍与依赖引入

**版本历史**

Spring 诞生到现在经历太多的历史版本，有的已经消失在历史长河中了… 我们选择最新的版本给大家进行案例讲解。

* 5.2.x 是最新的生产线（通常于 2019 年 9 月下旬提供）；
* 5.1.x 是之前的生产线（自 2018 年 9 月以来），一直得到积极支持，直到 2020 年底；
* 5.0.x 于 2019 年第二季度进入 EOL 阶段。出于对 5.0.x 用户的礼貌，我们将在 2020 年 1 月之前提供主动维护，并在 2020 年底之前提供安全补丁（如果需要）；
* 4.3.x 是第四代的最后一个功能分支。它具有延长的维护期限，直到 2020 年 1 月，并且安全补丁甚至超过了这一点。4.3.x 将于 2020 年 12 月 31 日达到其正式停产（停产）；
* 截至 2016 年 12 月 31 日，3.2.x 属于产品停产（寿命终止）。该产品线中没有计划进一步的维护版本和安全补丁。请尽快迁移到 4.3 或 5.x。

我们建议从 Maven Central 升级到最新的 Spring Framework 5.2.x 版本。

以上是官网列出 Spring 的历史版本介绍，我们采用的是 5.2.2 版本，对应的 jdk 最少是 jdk 1.8 ，我相信大家的 jdk 一般来讲都是满足要求的。

### 2.2 Spring 框架源码下载

**下载方式**：

1. 下载源码文件 。

Spring 的源码下载地址 ：

[https://github.com/spring-projects/spring-framework/releases](https://github.com/spring-projects/spring-framework/releases)
2.  第二种是使用 maven 的坐标方式 。

maven 的 pom 文件坐标。

```java
<dependency>
<groupId>org.springframework</groupId>
<artifactId>spring-context</artifactId>
<version>5.2.2.RELEASE</version>
</dependency>
```

## 3. 工程创建

**准备好依赖之后 废话不多说，我们开始撸代码** 。

### 3.1 使用 IDEA 创建 Web 工程

开发工具选择 idea ，创建 Maven 的 jar 工程即可。因为涉及不到浏览器的请求，所以无需创建 web 工程。

创建 Maven 工程 。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2m19hoj60jv0g411202)

补全坐标信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ma8foj60jv0g4wfu02)

继续下一步 finish 完成创建即可 这里不做截图演示了。

### 3.2 引入项目使用的坐标依赖

将准备好的坐标信息粘贴到工程下 `pom` 文件中 。cv 大法总会吧？找准位置一下就贴上了。 看下图：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2mkilrj61370lsds102)

### 3.3. 编写 Spring 框架使用的配置文件

坐标有了之后，说明我们的工程中已经引入了 Spring 框架的依赖。小伙伴可以检查下是否正确，点开左侧的 External Libraries 查看一下 。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2n73w0j60gg0bidjt02)

那么看到上面的 jar 包列表，表示 Spring 框架中的基本依赖我们已经成功引入。接下来：既然我们使用的是框架，框架是一个半成品，已经封装好了很多功能提供我们使用，而我们如何让他们工作呢？ 这里需要一个和 Spirng 框架通信的桥梁 —Spring 框架的核心配置文件。

**小提示**：

文件的名称你们可以随便起，我习惯使用 `applicationContext.xml`。

文件的位置放在哪里呢？ maven 工程需要放在 `src` 下面的 `resources` 下面，如下图：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ngnw7j60wy0ffwkf02)

那么配置文件是空的，不要着急。到底应该配置什么，不是自己臆想猜测的。

如果你已经下载了源码，那么解压缩它，打开 **docs\spring-framework-reference** 目录，打开 **core.html** 查看官方文档，

已经给了说明书你不看，你赖谁？ 不知道怎么看？下图告诉你：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2nq4k6j613z0fvgus02)

将上面的实例配置信息拷贝到我们的配置文件中，它只是给了最基本的配置头信息，内容部分 针对 bean 做初始化的部分 需要我们自行填充 。

## 4. 编写代码测试

准备好工程后，编写我们的代码。

### 4.1. 编写接口 和接口的实现类

**代码如下：**

```java
//接口的代码
public interface UserService {

    public void saveUser();
}
//实现类的代码
public class UserServiceImpl implements UserService {

    public void saveUser() {
        System.out.println("service的save方法执行了");
    }
}

```

### 4.2. 补充 Spring 的配置文件

配置文件的目的是将我们自定义的实现类交给 Spring 的容器管理。因为 Spring 框架核心功能之一就是 IoC 控制反转，目的是将对象实例化的动作交给容器。还记得第一节介绍的吗？不记得了？走你，剩下的我们继续。最终 Spring 的配置文件如下：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:util="http://www.springframework.org/schema/util"
       xsi:schemaLocation="
        http://www.springframework.org/schema/beans https://www.springframework.org/schema/beans/spring-beans.xsd
        http://www.springframework.org/schema/util https://www.springframework.org/schema/util/spring-util.xsd">

    <!-- 此标签的作用 是实例化UserServiceImpl类的实例 交给 Spring 容器 -->
    <bean id="userService" class="com.wyan.service.impl.UserServiceImpl"></bean>
</beans>

```

### 4.3 测试结果

从容器中获取对象实例，调用提供的方法

```java
public class DemoTest {

    public static void main(String[] args) {
        ApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        UserService service = (UserService) context.getBean("userService");
        service.saveUser();
    }
}

```

解释：

1. `ApplicationContext` 是 Spring 框架提供的一个接口，目前只需要知道它是作为存储实例化 bean 对象的容器即可。下一节我们会细讲。
2. `context.getBean ()` 方法是通过配置文件中声明的 bean 标签 id 属性获取容器内的实例。

最终结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2o3c78j60ku07tq4p02)

可以看到控制台打印输出 证明确实从容器中获取到了 userService 的实例。入门就是如此简单…

## 5. 小结

技术之路很简单 一是思路步骤清晰，二就是代码的熟练度。

先理清入门示例的步骤 ：

1. 创建 Maven 工程；
2. 导入 Spring 的依赖；
3. 编写 Spring 的配置文件；
4. 编写测试的代码。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
