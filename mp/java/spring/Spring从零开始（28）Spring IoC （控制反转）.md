# Spring IoC（控制反转）

## 1. 前言

通过第一章第二小节，我们已经可以使用 **Spring** 框架实现对自定义的 Java 对象管理，由 Spring 框架加载对象，实例化对象，放入容器。其实这就是 **Spirng** 的核心功能之 IoC，那么什么是 IoC 呢？什么又是容器呢？

跟我来，一步步揭开他们的神秘面纱。

## 2. 什么是 IoC？

**来自百度百科的解释 —— 控制反转（IoC）：**

（Inversion of Control，缩写为 IoC），是面向对象编程中的一种设计原则，可以用来降低计算机代码之间的耦合度。其中最常见的方式叫做依赖注入（Dependency Injection，简称 DI），还有一种方式叫 “依赖查找”（Dependency Lookup）。通过控制反转，对象在被创建的时候，由一个调控系统内所有对象的外界实体将其所依赖的对象的引用传递给它。也可以说，依赖被注入到对象中。

**慕课解释**

如何理解好 IoC 呢？上一个小节中，我们使用简单的语言对它做了一个描述 —— IoC 是一种设计模式。将实例化对象的控制权，由手动的 **new** 变成了 **Spring** 框架通过反射机制实例化。

那我们来深入分析一下为什么使用 **IoC** 做控制反转，它到底能帮助我们做什么。

我们假设一个场景：

我们在学习 **Web** 阶段的过程中，一定实现过数据的查询功能，那么这里我就举一个实例：

我们有这样几个类：

* UserServlet
* UserService 接口
*  UserServiceImpl 接口的实现类
*  UserDao

**代码如下：**

```java

/*
UserServlet  作为控制器 接收浏览器的请求
*/
public class UserServlet extends HttpServletRequest {
 //用户的业务类 提供逻辑处理 用户相关的方法实现
 private UserService userService;

 public void service(HttpServletRequest request,HttpServletResponse response){
    //手动实例化UserService接口的实现类
    userService = new UserServiceImpl();
    List<User> list =  userService.findAll();
    //省略结果的跳转代码

 }
}
/*
用户的业务接口UserService
*/
public interface UserService{
 public List<User> findAll();
}
/*
UserServiceImpl 作为用户的业务实现类 实现类UserService的接口
*/
public class UserServiceImpl implements UserService{
 //用户的Dao
 private UserDao userDao;

 public List<User> findAll(){
     //手动实例化Dao
     userDao = new UserDao();
     return userDao.findAll();
 }
}

```

**问题分析：**

上面的代码有什么问题吗？ 按照我们学习过的知识… 答案是没有。因为 Dao 只要数据源编写代码正确， 完全可以实现数据的增删改查 ，对吗？

但是分析分析它我们发现：

1. 代码耦合性太强 不利于程序的测试：

因为 `userServlet` 依赖于 `userService` ，而 `userService` 依赖于 `userDao` ， 那么只要是被依赖的对象，一定要实例化才行。所以我们采取在程序中硬编码，使用 `new` 关键字对对象做实例化。 不利于测试，因为你不能确保所有使用的依赖对象都被成功地初始化了。有的朋友很奇怪，对象实例化有什么问题吗？ 如果构造参数不满足要求，或者你的构造进行了逻辑处理，那么就有可能实例化失败；

2. 代码也不利于扩展：

假设一下，我们花了九牛二虎外加一只鸡的力气，整理好了所有的类使用的依赖，确保不会产生问题，那么一旦后续我们的方法进行扩充，改造了构造函数，或者判断逻辑，那么是不是所有手动 new 对象的地方都需要更改？ 很明显这就不是一个优雅的设计。

**解决方式：**

Spring 的 IoC 完美的解决了这一点， 对象的实例化由 Spring 框架加载实现，放到 Spring 的容器中管理，避免了我们手动的 new 对象，有需要用到对象实例依赖，直接向 Spring 容器要即可，而一旦涉及到对象的实例修改，那么 只需更改 Spring 加载实例化对象的地方，程序代码无需改动，降低了耦合，提升了扩展性。

## 3. 容器的使用

刚刚我们解释了 IoC 的作用，是对象的实例化由主动的创建变成了 Spring 的创建，并放入容器管理，那么这个容器是什么？

**概念理解：**

日常生活中有很多的容器，例如：水桶、茶杯、酒瓶，那么他们都有一个特点，就是装东西。而 Spring 的容器，就是装对象的实例的。

### 3.1 IoC 容器的体系结构

Spring 的容器有两个：

1. BeanFactory
2. ApplicationContext

他们两个都是接口，那么有什么区别呢？见图如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2pfc3rj60tf0odtgj02)

`BeanFactory` 才是 Spring 容器中的顶层接口。 `ApplicationContext` 是它的子接口。

简而言之，`BeanFactory` 提供了配置框架和基本功能，并在 `ApplicationContext` 中增加了更多针对企业的功能。

`BeanFactory` 和 `ApplicationContext` 的区别： 创建对象的时间点不一样。

`ApplicationContext`：只要一读取配置文件，默认情况下就会创建对象。

`BeanFactory`：什么时候使用，什么时候创建对象。

### 3.2 IoC 容器实例化的方式

上面已经知道 Spring 的容器是通过一个接口 `org.springframework.context.ApplicationContext` 表示，并负责实例化，配置和组装 Bean 对象。容器通过读取 xml 文件中的配置信息来获取关于实例化对象，配置属性等命令。

而 `ApplicationContext` 只是一个接口，我们通常创建 `ClassPathXmlApplicationContext` 的实例或者 `FileSystemXmlApplicationContext` 的实例。前者是从类路径中获取上下文定义文件，后者是从文件系统或 URL 中获取上下文定义文件 。例如：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2q5w7uj619j0d2gup02)

**代码解释：**

15 行注释掉的代码是通过加载类路径下的配置文件，一般来说 Java 工程放在 `src` 目录下。我们使用的是 Maven 工程放在 `resources` 目录下。

18 行代码是通过加载本地 D 盘目录下的文件来初始化容器， 实例化 bean 对象。

**结论**

通过上面的两种方式测试，发现都可以成功初始化容器， 获取测试的 bean 对象实例。

也证明了容器的初始化可以创建 `ClassPathXmlApplicationContext` 也可以创建 `FileSystemXmlApplicationContext` 的实例。

### 3.3 IoC 容器的使用实例

我们知道了加载配置文件初始化容器的方式，现在了解下容器的使用。其实对于我们而言，已经不陌生了，在第一章第二小节中也已经成功的从容器中获取了对象实例。

这里我们就回顾一下：

1. 容器的初始化必须先配置 xml 文件，代码回顾如下：

```java
<bean id="user"  class="com.wyan.entity.User" ></beans>

```

1. 加载配置文件

```java
ApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");

```

1. 调用方法

```java
context.getBean("user")

```

## 4. 小结

本小节对 IoC 概念做了一个详解，同时介绍了 IoC 解决的问题，演示了 IoC 的使用实例，对于初学者来说搞清楚概念，理解作用，实践出结果，就是出色的完成了任务。

技术没有捷径可走，多学习可以增加我们的知识，勤练习可以增加我们的经验，善于总结思路可以提升我们的能力，一起加油。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
