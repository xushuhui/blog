# Spring 多种方式初始化

## 1. 前言

通过之前的学习，我们可以熟练掌握 Spring 容器初始化的方法。常用的方法：一种是纯 xml 文件的方式，第二种是使用群体最多的一种，就是 xml 文件搭配类上面的注解，来进行初始化容器。

我们今天讲解一种全新的方法，也是目前最为流行的一种方法。是基于 JavaConfig 的方式来实现。通俗地说也叫基于注解的方式。

**疑问导出：**

我们学完了那么多种 Spring 的使用，其实完全可以胜任开发的需求，还有必要学习一种新的初始化方式吗？ 它又有什么优点呢？

伴随疑问：我们开始新的小节学习。

## 2. 注解介绍

### 2.1 注解列举

* **@Bean**
* **@Configuration**
* **@ComponentScan**
* **@Component @Service @Controller @Repository**

上述注解是我们在案例中需要使用到的。当然对于我们来讲，其实最后面的几个注解已经是熟客了，也就不多介绍。所以我把他们放到了一起。

上面的三个我们是没有用过的，所以下面我分别对三个注解做个解释说明。

### 2.2 注解作用阐述

**1.@Bean 注解的解释：**

见名知意，此注解作用于方法上，表示方法返回的实例初始化到由 Spring 管理的容器中。其实对现在的我们而言，非常好理解。因为我们知道 Spring 的 IoC 其实就是控制反转，实例化 bean 实例，管理 bean 实例。

相当于使用 xml 文件方式的 **<bean>** 标签，一般来说，此注解出现在被 `@Configuration` 注释的类中。

那么…@Configuration 注解的作用又是什么呢？

**2.@Configuration 注解的解释：**

其实看此注解的名称，也能大致猜到它的作用。没错，它的作用就是配置，也就是相当于我们之前编写过很多次的 Spring 配置文件 ——`Applicationcontext.xml`。

之前写过的案例中，Spring 容器的初始化必须加载这个配置文件，而在配置文件中，就包含了很多的 标签，大家没忘记吧？不要提了… 不要关了电脑就忘记之前写过的代码。

那么我们上面说 `@Bean` 标签，一般出现在被此注解注释的类中，就是这个原因。

**3.@ComponentScan 注解的解释：**

根据上面两个注解的推断，我们能猜出它的作用就是组件扫描。相当于之前我们在 Spring 的配置文件中用过的标签 **[context:component-scan]()**,

现在讲述使用 JavaConfig 是基于注解的方式，当然也避免不了组件的扫描。此注解也是配置类的组成部分。

搭建项目最基本的三个注解给大家介绍过了，下面我们就进入正题。

## 3. 工程实例

**工程步骤：**

1. 创建 maven 工程
2. 导入依赖
3. 编写配置类
4. 测试代码

**步骤解释：**

第一步，第二步省略，实在演示太多次了。谁不会拖出去枪毙十分钟… 剩下的咱们继续。

**配置类：**

在 src 下的 `com.wyan.config` 目录之下创建配置类 SpringConfig，代码与目录结构如下：

![](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo32ld7bj618j0j6wo002)

**代码解释：**

SpringConfig 的类上面的注解，表示当前的类是个配置类，容器中的 bean 实例都从这个配置类中来。

`@Bean` 注解所在的方法会返回 UserServiceImpl 的实例，而且被 Spring 的容器管理。

**测试类：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo32yzdzj618j0dvqbv02)

**代码解释：**

1. 16 行的代码，我们看到一个新的类叫做 `AnnotationConfigApplicationContext`，此类是注解配置应用容器。
2. 17 行代码，通过调用方法 register () 将配置类 SpringConfig 作为参数传入，容器就成功加载了配置类。
3. 18 行刷新方法调用大家应该不陌生吧，在源码解释的小节，就看到它内部就是初始化容器的详细步骤。

**测试结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo33wx6sj60vl0aetdq02)

上面可以看到，servcie 的实例确实从容器中获取得到，那么说明，我们基于 JavaConfig 的配置方式已经实现完成。但可能大家会有一点疑问：

配置类中的方法是返回 bean 对象的实例，放入容器。那么一个方法返回一个 bean ，如果我们的项目中存在非常多的类，难道需要创建很多个方法吗？

下面使用最后一个注解，也就是 `@ComponentScan` 的使用。

**更改配置类：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo34i619j60vp0d07ab02)

代码解释：

1. 在类上面加入注解 `@ComponentScan`，而注解中的值，就是需要被扫描的组件类所在目录；
2. 因为通过扫描加载组件，所以之前创建 bean 返回实例的方法不需要；
3. 扫描的 servcie 目录下的类上一定需要搭配 `@Service``@Controller``@Repository``@Component`。

**测试结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo34wqo4j60vl0aetdq02)

## 4. 小结

本节，主要讲解 Spring 容器的基于 JavaConfig 的实现方式。核心思想是借助于注解加载配置，加载 bean。与其他方式相比各有利弊，大家不要执着于形式，根据自己的开发习惯，公司规范来选择就可以。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
