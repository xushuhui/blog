# Spring IoC（控制反转）之常用注解

## 1. 前言

上一节，我们通过注解的方式，实现了 Spring 对于 bean 的管理，那么如何实现的，是否还记得呢，我们回顾一下

两个重要点：

1. 注解实例化的类上，需要使用一个注解 `@Repository`；

2.Spring 的配置文件中，需要使用组件扫描 `<context:component-scan>` 。

**疑问导出**：

组件扫描的作用我们清楚，是为了扫描路径之下带有注解的类，但是为什么类上面的注解是 `@Repository` 呢？或者说，是否还有其余的注解可以实现呢？

本节，我们一起来学习下 Spring IoC 的常用注解。

## 2. 注解的详解

在我们详细讲解注解之前，首先明确一点：

注解配置和 xml 配置实现的功能都是一样的，只不过实现的方式不同，那么也就是说，xml 文件可以实现的，通过注解都可以完全办得到。比如实例化对象，设置属性，设置作用范围，生命周期的方法执行等等…

### 2.1 注解分类介绍

**按功能划分**：

1. **创建对象：** 对应的就是在 xml 文件中配置的一个 bean 标签，可以定义 id、name、class 等属性；
2. **注入数据：** 对应的就是在 bean 标签下，使用 property 标签给类中的依赖属性赋值；
3. **作用范围：** 对应的就是设置 bean 标签的 scope 属性，不设置默认也是单例；
4. **生命周期：** 对应的就是设置 bean 标签的 init-method 和 destroy-method 方法。

#### 2.1.1 创建对象的注解介绍

从 Spring 的官网得知一段话：

> `@Repository` 注释是针对满足的存储库（也被称为数据访问对象或 DAO）的作用，或者固定型的任何类的标记。

也就是说，我们上一节中使用的注解，一般用于 dao 层使用。那么，我们都知道，JAVAEE 体系结构，一般开发分为三个层级：

1. **表现层**： 主要作用为处理数据生成静态的页面响应给浏览器展示 ；
2. **业务层**： 主要作用为业务逻辑代码编写，数据的获取，数据的封装返回等等操作都在这里；
3. **持久层**： 主要作用为跟数据库打交道，对于数据的持久化操作等。

那么，如果是创建的表现层或者业务层代码，应该使用什么注解呢？

好了，看一下创建对象注解的划分：

1. **@Component** ：一般用于通用组件的类上使用的注解；
2. **@Service** ： 一般用于业务层类上使用的注解；
3. **@Controller** ： 一般用于控制层类上使用的注解；
4. **@Repository** ：一般用于持久层类上使用的注解。

**官网解释**：

> Spring 提供进一步典型化注解：`@Component`，`@Service`，和 `@Dao`。
>
>
> `@Component` 是任何 Spring 托管组件的通用构造型。
>
>
> `@Repository`，`@Service` 和 `@Controller` 是 `@Component` 针对更特定用例的专业化（分别在持久性，服务和表示层）。

**慕课解释：**

`@Component` 注解是 Spring 框架中通用的一个注解，用于组件扫描实例化对象使用， 那么其余的三个注解 `@Controller` ，`@Service`，`@Repository` 都是 `@Component` 注解的衍生注解，作用跟 `@Componet` 注解的作用一致。

那么意义在于， 三个注解，对应的是三个开发层级 ，一般来讲我们将 `@Controller` 作为表现层的使用，`@Service` 作为业务层的注解，`@Repository` 作为持久层使用的注解。我们下面通过案例演示一下。

### 2.2 创建对象的注解

**实例说明**

四种注解的测试，本节重点讲解创建对象使用的注解，而作用范围 scope 和生命周期的两个注解，我们放在后续对应的小节进行讲解测试。

置于注入数据的注解，是比较重要的一个内容， 我们放在依赖注入这节详细讲解。

各位同学… 稍安勿躁，我们一起慢慢成长。

**创建工程省略**

我们继续使用上一节的注解工程实例即可，那么为了演示三个注解，我们分别创建三个层级对应的代码：

* 表现层的 `UserController`
* 业务层的 `UserService`
* 实现类 `UserServiceImpl`

持久层 dao 代码已经创建过了，这里不多解释。创建好的所有代码如下：

**UserController 代码：**

```java
@Controller
public class UserController {

    public void saveUser(){
        System.out.println("这是controller的执行保存..");
    }
}

```

**UserService 和实现类代码**：

```java
public interface UserService {

    public void saveUser();
}

@Service
public class UserServiceImpl implements  UserService {

    public void saveUser() {

        System.out.println("执行service中的保存逻辑");
    }
}
```

**项目结构如下：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2tpjd6j60vt0hcjy102)

上面是本案例的工程以及代码结构：

​ 类虽然看起来很多，实际没有业务逻辑代码，只不过在各个层级使用了三个注解来注入到容器，目的是测试当 Spring 的配置文件加载扫描后，是否可以从容器中获取三种注解（`@Controller``@Service``@Repository`）注入的 bean 对象。

> **Tips：** Spring 的配置文件 `context:component-scan` 标签的扫描层级 需要包含三个包路径，例如我的工程实例代码如下：

```java
<context:component-scan base-package="com.wyan"></context:component-scan>
```

**测试类与测试结果**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2u4fshj61270lhk7y02)

**结论：**

可以三个注解都可以将对象注入到 Spring 的容器，那么以后开发时候按照规范或者习惯，分层开发，使用对应的注解。但它并不是必须这么做，你使用任意一种都可以，只不过，代码的可读性会差。

所以，我们一般表现层使用 `@controller` ，业务层使用 `@service`， 持久层使用 `@Repository`。

至于 `@Component` 如果有其余的类，不属于三个层级，可以采用 `@Component` 作为通用组件扫描注入容器。

### 2.3 注解注入规则

刚刚通过三个注解都可以完成了 bean 的实例化注入，通过测试代码也获取到了容器中的三个对象实例，那么这里不知道大家是否发现一个问题：

我们知道，Spring 这个容器本质是个 map 集合来存储实例化后的对象。既然是个 map 集合，就应该对应的有 key 和 value。

我们都知道 value 肯定是实例化后的 bean ，那么 key 是什么呢？

**注入规则**：

1. 四种注解都支持 value 的属性作为自定义的 bean id ;

2. 如果 value 属性没有指定，那么默认以类的简单名称（类名首字母小写）作为 bean 对象的 id。

所以我们可以看到：

当我们只使用注解没有自定义 id 的时候可以通过，每个类的首字母小写来获取对象实例，那么如果有了自定义的 id，上述代码是否继续可用呢？

**自定义 id 获取实例**：

改造类上面的注解，设置自定的 id，更改的注解如下：

```java
@Controll("uc")
@Service("us")
@Repository("ud")
```

**测试结果**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2vswstj619j0lw1kx02)

**测试结果**：

为了区分测试结果，我在测试代码中，只修改了 controller 的获取方式，将 id 改成了 uc 。service 和 dao 并没有修改。

从控制台打印可以看到，只有 controller 对象可以成功获取，service 和 dao 都失败了，因为我们已经使用了自定义的 id，所以容器中没有默认的以类名作为 id 的 bean 对象实例。

## 3. 小结

本章节重点讲解注解的使用：

1. Spring 支持的注解有四种分类；
2. Spring 创建对象的注解四种分类；
3. Spring 创建对象注入容器的规则。

学习的苦只是一时之苦，学不到的苦是一世之苦，与君共勉…

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
