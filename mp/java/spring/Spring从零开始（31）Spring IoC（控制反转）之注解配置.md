# Spring IoC（控制反转）之注解配置

## 1. 前言

上两节，我们学习了 Spring IoC 的 xml 配置实现，整理了 xml 方式的实现步骤，并且模拟了 Spring 的容器如何加载解析 xml 配置文件，那么我们发现一点现象：

对于 Spring 的 bean 管理而言，如果全部通过 xml 文件实现的话，配置文件的内容未免过于臃肿。因为对于一个类的实例化，就需要一个 bean 标签。

这样的话，一个大型工程下来，有那么几百个，几千个类，Spring 的 xml 文件维护起来，成本实在太高。

**疑问导出**：

Spring 能否有更方便的方式实现 IoC 呢？Spring 提出了两种 IoC 实现方式，一种是基于配置文件，一种是基于注解形式。

本节，我们学习下 Spring IoC 的注解形式是如何实现的。

## 2. 案例实现

### 2.1 步骤介绍

**回顾 Spring IoC 的 xml 实现步骤**：

​ 1. 使用 new 关键字对 `ClassPathXmlApplicationContext` 做初始化；

​ 2. 在初始化容器对象的构造传入 xml 配置文件的位置 ；

​ 3. 在配置文件中通过 bean 标签可以对类进行描述：类的路径、类的标识、类的构造参数等等。

**注解实现 IoC 的思路分析**:

​ 1.Spring 容器一样需要初始化；

​ 2. 一样需要传入 xml 配置文件 ----- 需要描述清楚 需要被实例化的类都有哪些；

​ 3.xml 文件中 不需要使用 bean 标签描述被实例化的类 ------ 使用注解实现 IoC 管理目的就是为了简化 bean 标签的配置。

**疑问导出：**

如果是 xml 文件方式实现 IoC ，加载 xml 文件的 bean 标签就已经知道，需要被实例化的对象，那么如果不使用 bean 标签描述，Spring 框架如何得知哪些类需要被容器管理呢？

**核心思想：**

开发人员无需使用 XML 来描述 bean ，而是将配置移入 Java 的类本身，通过 Spring 支持的组件扫描来实现。

看官稍等… 马上开始我们的案例实现。

### **2.2 工程实现**

**创建工程**：

为了区分 xml 工程，坐标名称换成 spring_an ，其实无所谓，大家自行创建即可。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2suv1wj60jv0g4myi02)

**导入依赖**：

依赖的坐标跟 xml 的工程坐标一致即可，无需导入多余的依赖。

```java
<dependencies>
    <dependency>
    <groupId>org.springframework</groupId>
    <artifactId>spring-context</artifactId>
    <version>5.0.2.RELEASE</version>
    </dependency>
</dependencies>
```

**项目代码**：

为了测试，在工程内部创建 `UserDao` 的接口和 `UserDao` 的实现类 `UserDaoImpl`。

`UserDao` 代码如下：

```java
public interface UserDao {

    public void saveUser();
}
```

`UserDaoImpl` 的实现类代码如下：

```java
@Repository
public class UserDaoImpl implements  UserDao {

    public void saveUser() {
        System.out.println("执行dao的保存方法");
    }
}
```

注意事项： 由于我们是基于注解的方式实现对 bean 的管理，所以在实现类上面需要添加一个注解 @Repository，此注解的作用是为了 Spring 的容器启动后，

需要要自动检测这些被注解的类并注册相应的 bean 实例到容器中。

**Spring 的核心配置文件**：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:context="http://www.springframework.org/schema/context"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
        https://www.springframework.org/schema/beans/spring-beans.xsd
        http://www.springframework.org/schema/context
        https://www.springframework.org/schema/context/spring-context.xsd">

    <context:component-scan base-package="com.wyan.dao"></context:component-scan>

</beans>
```

上面是本案例的配置文件，那么可以看出跟 xml 的配置文件有很大的区别：

​**配置节点**：`context-component-scan` 标签，这是 Spring 框架自定义的 xml 标签，通过 `base-package` 的属性，指明需要被自动扫描实例化的类所在位置。

如上图所示，我们在 com.wyan.dao 下的类是需要扫描自动注入容器的。

小细节：不是在 com.wyan.dao 下的所有类都会自动注入到容器，而是要搭配注解：比如我们的 @Repository 当然还有其余的注解，我们后面章节会详细讲解。

**测试类测试结果**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2t7ig3j61950mv19n02)

**代码解释**：

测试类其实跟 xml 的方式一模一样，我们本次测试的目的一样也是通过 Spring 容器管理注册的 bean 对象，只不过对象的实例化方式换成了注解，那么我们看到成功输出在控制台的测试语句，说明案例搭建完成。

## 3. 小结

本节带着大家使用注解的方式，实现了 Spring 对于 bean 的管理。

那么回顾下注解开发的步骤和注意点：

1. Spring 容器初始化一样需要 xml 文件，目前是 xml 文件搭配注解管理 bean 并不是纯注解开发；
2. Spring 的 xml 配置文件中使用 `context:component-scan` 标签指定注册 bean 的类所在目录位置；
3. 自定义编写的 Java 类，如果需要被自动扫描注入容器，必须搭配注解。

学习的苦只是一时之苦，学不到的苦是一世之苦，与君共勉…

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
