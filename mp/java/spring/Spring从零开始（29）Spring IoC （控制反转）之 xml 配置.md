# Spring IoC （控制反转）之 xml 配置

## 1. 前言

本小节目的在于带领大家熟练 **xml 文件配置**， 应用 xml 文件配置 IoC。

在第二节中我们通过一个入门工程简单地体验了一把 Spring 的使用。在第三节中梳理了一下 Spring 的工作流程。

大家有了一个初步认知，Spring 框架的工作脱离不了核心配置文件 `applicationContext.xml`。

在配置文件中我们目前只用到了一个 bean 标签，它的作用是用于描述 Java 的类，让框架启动加载配置文件实例化的。

**疑问导出**

那么我们知道描述一个类有几个要素，类名、属性、构造函数 set 和 get 方法对吧？而 bean 标签如何描述一个详细的类呢？

带着疑问… 开始本节内容。

## 2.bean 标签中的属性介绍

**核心配置文件回顾**

```java
<bean id="user" class="com.wyan.entity.User" ></bean>


```

在上面的代码中可以看到，在 bean 标签中有两个属性，一个是 id 一个是 class。那么在 bean 标签中都有哪些属性呢？

**属性列表**

|学号|姓名|
|----|----|
| id            | 定义的唯一标识  |
| name          | 同 id 的意义一致|
| class         | 类              |
| factory-bean  | 工厂对象        |
| factory-method| 工厂方法        |
| init-method   | 初始化执行的方法|
| destroy-method| 销毁执行的方法  |
| scope         | 对象的作用域    |
| lazy-init     | 懒加载          |
| autowire      | 依赖注入        |
| depends-on    | 依赖于某个实例  |

**疑问导出**

上述属性是配置 bean 标签中可以选择的属性，当然一般来讲，我们无需配置所有，可以根据自己的需求配置需要的属性信息，那么如何选择这些属性呢？

### 2.1 属性详细解释

#### **2.1.1 id 和 name 标签的使用**

我们目前已经知道所有被**实例化**后的对象都存在于 **Spirng 的容器**中，那么从容器中获取这些对象需要一个属性 id 对吧？那么 **name 和 id** 有什么关系呢？

查看官方文档得知 Spring 的容器会给初始化的每个 bean 都定义一个或多个**标识符**。这些标识符在容器内必须是**唯一**的。一个 bean 通常只有一个标识符。而 name 和 id 都可以起到标识符的作用。

所以在 **XML** 配置文件，我们一般使用 **id** 或者 **name** 属性，定义 bean 的唯一标识，

这样我们才能通过定义好的唯一标识，从 Spring 的容器中获取他们。

**代码实例**:

xml 的配置文件如下：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans.xsd">

    <bean id="user" name="user2" class="com.wyan.entity.User" ></bean>

</beans>

```

测试代码如下：

```java
	public static void main(String[] args) {
        ApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        System.out.println(context.getBean("user"));
        System.out.println(context.getBean("user2"));
    }

```

结果如图所示：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2qdb0kj611h0amaim02)

**结论证明**：

我们通过 bean 标签中的 id 属性 user， 或者使用 bean 标签中的 name 属性 user2， 都可以得到 Spring 容器中的 user 对象的示例，而且打印的地址是同一个。我们之前说过一句，默认在容器中的实例都是单例的，在这里也得到了证明。

#### **2.1.2 class 属性**

bean 标签的定义实质上是创建一个或多个对象的方法。当 **xml 文件被解析加载**的时候，使用该 bean 定义封装的配置数据来创建（或获取）实际对象，而创建获取的对象是谁呢？就是通过 **class 属性**中定义的类的全路径来指定 。

一般来讲 class 中的类实例化有**两种方式**：

​ 一种是**反射** ，相当于我们使用的 new 关键字。这种也是我们常用的方式。当然不要忘记提供无参数的构造方法（类中默认有无参构造，但是如果自定义了有参构造，默认的无参不会提供）

​ 一种是**工厂模式** ，需要借助于 factory-bean 和 factory-method 两个属性，这种方式不常用，我们可以了解下。

#### 2.1.3 factorybean 和 factorymethod 属性

这两个属性主要用于工厂模式实例化 bean 的时候使用，不是很常见。工厂模式有两种，这里分别做个实例，帮助大家理解。

**静态工厂**模式实例：

```java
<!--applicationContext的配置bean节点-->
<bean id="user" class="com.wyan.entity.User" factory-method="createUserInstance"/>

```

创建 bean 示例的 Java 工厂类：

```java
public class User {

    private static User user = new User();

    private User() {}

    public static User createInstance() {
        return user;
    }
}

```

解释：在定义使用静态工厂方法创建的 bean 时，class 属性指定的是被创建的类，包含静态的方法，并使用 factory-method 属性来指定工厂方法本身名称。

**普通工厂模式**：

```java
<!--spring实例化工厂对象 用于创建java实例 -->
<bean id="beanFactory" class="com.wyan.factory.BeanFactory"></bean>
<!-- 被工厂创建的对象实例 -->
<bean id="user1" factory-bean="beanFactory" factory-method="createUser1"/>

```

工厂类代码：

```java
public class BeanFactory {

    private static User1 user1 = new User1();

    private static User2 user2 = new User2();

    public User1 createUser1() {
        return user1;
    }

    public User2 createUser2() {
        return user2;
    }
}

```

解释：先实例化先创建各个对象示例的工厂对象到容器中，自身的 bean 标签将 `class` 属性保留为空，并在 `factory-bean` 属性中指定当前容器中的工厂 Bean 名称，再使用 `factory-method` 属性设置创建示例的方法名称。

#### 2.1.4 init-method 和 destroy-method 属性的使用

这两个属性比较好理解 init-method 就是 bean 被初始化后执行的方法，destory-method 就是 bean 被销毁执行的代码。

**我们来个测试类**：

```java
public class User {

    public User(){
        System.out.println("我被spring实例化了");
    }

    public void initMethod(){
        System.out.println("user类实例化时候执行的代码");
    }
    public void destoryMethod(){
        System.out.println("user类实例被销毁时候执行的代码");
    }
}

```

**配置文件**：

```java
 <bean id="user" name="user2" class="com.wyan.entity.User" init-method="initMethod" destroy-method="destoryMethod" ></bean>**

```

**测试代码**:

```java
	public static void main(String[] args) {
        ApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");

    }

```

加载 Spring 的配置文件控制台打印如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2qn0h3j61370acgul02)

有个小疑问：销毁语句没打印呢？那是因为并没有调用容器的销毁方法。

**改造测试**代码如下：

```java
	public static void main(String[] args) {
        AbstractApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        context.close();
    }

```

解释：ApplicationContext 没有 close 方法使用它的子类

**运行结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2qw5b0j611n0bgak302)

#### 2.1.5 其余属性作用

**scope** ：指定示例的作用范围，后续章节详细讲解；

**lazy-init** ：表示是否为懒加载；

**autowire** ：指定属性注入方式，后续章节详解；

**depends-on**： 表示是否有依赖的 bean 对象，后续依赖注入章节详细解释。

### 2.2 构造函数的使用

刚刚我们详细解释了 bean 标签内部的属性，经过几个小实例以后不禁也有个问题：

如果我们定义的类中有一些初始化的参数，并且定义好了有参数的构造，通过 xml 配置文件如何体现呢？

实现起来非常简单，跟我来进行一个小实例：

**改造 User 类**：

这是一个普通的 Java 类对象，包含两个属性及其 get 和 set 方法，并且提供了空参构造和有参构造，为了测试方便再覆写一个 toString 方法。

```java
public class User {

    private Integer id;
    private String name;

    public User() {
    }
    public User(Integer id, String name) {
        this.id = id;
        this.name = name;
    }
    public Integer getId() {
        return id;
    }
    public void setId(Integer id) {
        this.id = id;
    }
    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", name='" + name + '\'' +
                '}';
    }
}

```

**xml 配置文件方式**：

```java
    <bean id="user"  class="com.wyan.entity.User"  >
        <constructor-arg name="id" value="1"></constructor-arg>
        <constructor-arg name="name" value="zs"></constructor-arg>
    </bean>

```

**测试结果**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2r4sphj611p0ac46a02)

其实对于有参构造实例化对象而言，使用一个标签 constructor-arg 即可，表示构造的参数，如果有多个，可以继续添加，这里不多做演示。

**疑问导出**：

可能有同学会想，那么如果以后我们的属性需要动态更改呢？或者我们的属性不是基本类型而是另外的对象呢？ 后续在第三章依赖注入多种属性的小节给大家讲解 。

## 3. 小结

本章节带着大家详细解释了 bean 标签的使用，那么通过本章节我们收获了哪些呢？

1. 容器内部命名唯一标识可以通过 id 也可以通过 name；

2. 实例化对象有两种方式 反射模式和工厂模式；

3. 如果是反射模式，那么必须配置 class 属性，因为需要用 class 属性中类的全路径来实例化 bean 对象；

4. 如果需要在类实例化初始化参数，可以使用 init 方法也可以使用有参构造。

持之以恒的学习是成功的最快捷径… 切记眼高手低。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
