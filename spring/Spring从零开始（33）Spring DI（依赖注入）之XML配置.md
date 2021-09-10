# Spring DI（依赖注入）之 XML 配置

## 1. 前言

在第二个整个大章节，我们详细讲解了 控制反转，也就是对 bean 做实例化的部分。而我们知道 ，Spring 的核心功能是两个：控制反转 和 依赖注入。

那么控制反转我们已经讲过，而依赖注入是什么呢？

各位看官，随小可一起来…

## 2. 依赖注入案例

#### 2.1 概念介绍

**知识回顾**

对于依赖注入，我们在第一章第一节已经介绍过，我们回顾一下

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2wcpk6j60lb0auadn02)

**概念解释**

上面是我们之前对于依赖注入的一个通俗解释。那么这里再着重强调一下 **IOC** 控制反转与 **DI** 依赖注入的关系：

IOC 控制反转是将对象实例化的动作交由了 Spring 框架， 它的作用是降低了程序的耦合，不需要我们手动的创建对象，但是程序的耦合性还是存在。

对象中肯定会有一些其余对象的引用，那么这种引用就称呼为对象的依赖，而 DI 依赖注入其实 是 IOC 设计思想的一种表现形式。

对于 这种属性依赖，我们无需手动赋予，也是讲赋值的动作交给 Spring ，那么这种操作就是 **依赖注入**。

**依赖注入方式**：

* 第一种方式是通过 xml 配置的方式实现；
* 第二种方式是在属性或者方法上使用注解的方式实现。

那么，本章节先带大家体验下 xml 方式实现依赖注入。

#### 2.2 工程实现：

**搭建动作介绍**

1. 创建一个 maven 工程
2. 导入 Spring 使用的依赖
3. 编写业务层的 Service 和持久层的 Dao java 类
4. 编写 Spring 的配置文件

**创建工程 导入依赖 省略**

可以参考之前创建过的 IOC 工程

**java 代码**

创建 Servcie 的接口和接口的实现类，代码如下：

```java
//接口代码
public interface UserService {

    public void deleteById(Integer id);
}
//实现类代码
public class UserServiceImpl implements UserService {

    private UserDao userDao;

    public UserDao getUserDao() {
        return userDao;
    }

    public void setUserDao(UserDao userDao) {
        this.userDao = userDao;
    }

    public void deleteById(Integer id) {

        System.out.println("删除的方法执行");
    }

}
```

UserDao 接口和实现类代码：

```java
//dao接口代码
public interface UserDao {

}
//dao实现类代码
public class UserDaoImpl implements UserDao {
}

```

代码解释： dao 的接口和实现类中并没有方法，只是为了测试 作为 service 中的属性依赖，可以实现由 Spring 完成动态注入。

**重点来了：spring 的核心配置文件：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2wydjcj60th0eu7dl02)

配置解释：

在上面的配置文件中：

* bean 标签是描述一个被实例化的类 而 property 则表示一类中的属性
* property 标签中的属性 name 一般我们写成类中的属性名称， 实际上，起决定作用的并不是属性名，下面示例再展示
* ref 表示当前的属性 是一个引用对象，而引用的是谁呢？ ref 中的值 必须是在容器中已经实例化的一个引用对象的唯一标识。
* value 当前的属性可以直接赋值，所以通过 value 中，填写要赋予的数值即可

**测试结果**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2xkzvgj60ye0kztju02)

**代码解释**

可以看到 我们得到了 service 中的类属性 Userdao 的实例，并且也 得到了 字符串属性 userName 的值 zs

#### 2.3 property 注入属性的解释

刚刚我们在上面的示例中 展示了 xml 依赖属性的注入，也是比较好理解。

这里我们强调一下使用的注意事项：

如果是 property 属性标签实现属性注入，那么类中必须由配置在 property 标签中 name 属性的 set 方法

下面我们测试一下 set 方法的改变：

**先讲 service 中 dao 的 set 方法改造如下：**

```java
public void setDao(UserDao userDao) {
    System.out.println("执行了set方法 给dao属性赋值");
    this.userDao = userDao;
}
```

这时候代码中的 set 方法变成了 setDao 配置文件不变，依然是

```java
<property name="userDao" ref="userDao"></property>
```

我们看看会产生什么问题

```java

Caused by: org.springframework.beans.NotWritablePropertyException: Invalid property 'userDao' of bean class [com.wyan.service.UserServiceImpl]: Bean property 'userDao' is not writable or has an invalid setter method. Does the parameter type of the setter match the return type of the getter?
	at org.springframework.beans.BeanWrapperImpl.createNotWritablePropertyException(BeanWrapperImpl.java:247)
	at org.springframework.beans.AbstractNestablePropertyAccessor.processLocalProperty(AbstractNestablePropertyAccessor.java:426)
	at org.springframework.beans.AbstractNestablePropertyAccessor.setPropertyValue(AbstractNestablePropertyAccessor.java:278)
	at org.springframework.beans.AbstractNestablePropertyAccessor.setPropertyValue(AbstractNestablePropertyAccessor.java:266)
	at org.springframework.beans.AbstractPropertyAccessor.setPropertyValues(AbstractPropertyAccessor.java:97)
	at org.springframework.beans.AbstractPropertyAccessor.setPropertyValues(AbstractPropertyAccessor.java:77)
	at org.springframework.beans.factory.support.AbstractAutowireCapableBeanFactory.applyPropertyValues(AbstractAutowireCapableBeanFactory.java:1646)
```

可以看到异常的堆栈信息 无效的 userDao 属性， userDao 不可以 或者 没有有效的 setter 方法提供。

**更改 xml 文件中的 property 标签的 name 属性 为 dao**

```java
<property name="dao" ref="userDao"></property>
```

测试结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2y9dm0j60v009uahn02)

所以我们说 property 中的 name 属性不一定要跟 Java 类中的属性名保持一致 而是必须跟 setter 方法的名称一致

## 3. 总结：

本章节重点依赖注入的 xml 实现

1. 依赖注入 实际上是 IOC 设计思想的一种具体实现
2. 依赖注入 可以通过 xml 配置实现 ，可以通过注解实现
3. xml 的依赖注入 是依托于类中的 set 方法实现的。
4.

不积跬步无以至千里，不积小流无以成江海…继续努力

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
