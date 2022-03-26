# Spring AOP 的代理模式

## 1. 前言

大家好，我们学习了 AOP 的概念，理解了 AOP 的作用，明白了什么是 AOP 面向切面编程，目的是进行解耦，可以自由地对某个功能方法做增强，而它的设计模式就是代理模式。

那么本小节我们就学习 AOP 设计模式之代理模式，明白代理模式的含义，掌握代理模式的语法，同时明白它的应用场景。

## 2. 代理模式介绍

### 3.1 概念解释

代理名词解释为：以被代理人名义，在授权范围内与第三方实施行为。而在软件行业中代理模式是一种非常常用的设计模式，跟现实生活中的逻辑一致。

在开发中代理模式的表现为：我们创建带有现有对象的代理对象，以便向外界提供功能接口。代理对象可以在客户端和目标对象之间起到中介的作用。

中介什么作用？ 可以为被代理对象执行一些附带的，增加的额外功能。

**代理模式图例：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo36rfxbj627c0soaot02)

**图例解释**

* 数字 1 处表示发起请求的客户端；
* 数字 2 处表示需要调用的接口；
* 数字 3 处表示真实接口的实现类，执行的代码逻辑在此；
* 数字 4 处表示代理对象，同样实现了接口，覆写了需要执行的代码 Request 方法；
* 数字 5 处是代理类中多余的代码逻辑，对原有方法的增强；
* 数字 6 处是代理类中的执行的被代理类的原始逻辑。

### 3.2 代理模式分类

上面已经解释了代理模式的概念和原理，在开发中，实现代理模式可以分为两种。

* **静态代理**：若代理类在程序运行前就已经存在，那么这种代理方式被成为静态代理 ，这种情况下的代理类通常都是我们在 Java 代码中定义的， 静态代理中的代理类和委托类会实现同一接口；
* **动态代理**：代理类在程序运行时创建的代理方式被称为动态代理。 也就是说，这种情况下，代理类并不是在 Java 代码中定义的，而是在运行时根据我们在 Java 代码中的 “指示” 动态生成的。

下面，我们就针对这两种模式演示效果… 各位看官随我来。

## 3. 代码模式测试

### 3.1 静态代理实现

**1.userService 接口代码**

```java
public interface UserService {

    public void saveUser();
}
```

**2.userServiceImpl 实现类代码**

```java
@Service
public class UserServiceImpl implements  UserService {

    public void saveUser() {

        System.out.println("执行service中的保存逻辑");
    }
}
```

**3.userServiceProxy 代理类代码**

```java
public class UserServiceProxy implements  UserService {
	//被代理类实现接口
    private UserService userService;

    public UserServiceProxy(UserService userService) {
        this.userService = userService;
    }
	//覆写的方法
    @Override
    public void saveUser() {
        System.out.println("原始功能执行之前的逻辑代码");
        userService.saveUser();;
        System.out.println("原始功能执行之后的逻辑代码");
    }
}
```

**代码解释**：

`userService` 接口和 `userServiceImpl` 实现类代码不做赘述，已经用过多次。

重点关注于在 `userServiceProxy` 代理类代码，其中属性为被代理类的接口，目的是传入进来被代理类实例，对它做功能增强。

下面的 `saveUser` 方法是代理类执行的逻辑，在方法内部有增强的代码逻辑，也保留了原始实例的代码功能。

**4. 测试代码**

```java
public static void main(String[] args) {
    AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext();
    context.register(SpringConfig.class);
    context.refresh();
    //获取接口实例
    UserService service = context.getBean(UserService.class);
    //创建实例的代理
    UserServiceProxy proxy = new UserServiceProxy(service);
    //执行方法
    proxy.saveUser();
}
```

**5. 测试结果**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo37m5haj60qm0a1jwp02)

可以看到，执行结果中即包含了被代理对象的原始保存方法的逻辑，也有代理类中对原始方法的两个增强代码。

### 3.2 动态代理实现

**1. 创建动态处理器**

```java
public class DynamicProxy implements InvocationHandler {


    private Object object;

    public DynamicProxy(final Object object) {
        this.object = object;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                  System.out.println("执行前逻辑");
                    Object result = method.invoke(object, args);
                  System.out.println("执行后逻辑");
                   return result;
    }
}
```

**2. 测试类代码**

```java
public class SpringTest {

    public static void main(String[] args) {
        AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext();
        context.register(SpringConfig.class);
        context.refresh();
        //获取接口实例
        UserService service = context.getBean(UserService.class);
        //动态创建实例的代理
        UserService proxy = (UserService)Proxy.newProxyInstance(UserService.class.getClassLoader(),
                            new Class[]{UserService.class}, new DynamicProxy(service));
        //proxy执行方法
        proxy.saveUser();

    }
}
```

**代码解释**

`Proxy.newProxyInstance` 是 JDK 提供的一个用于动态创建代理实例的方法，JDK 1.7 的 API 有如下说明：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo380oxmj60ub0bhgp002)

红线框的内部是简易方式创建代理对象的方式，也就是我们示例中代码的实现方式，参数解释如下：

* **ClassLoader loader**: 指定当前目标对象使用的类加载器，获取加载器的方法是固定的
* ** Class<?>[] interfaces**: 指定目标对象实现的接口的类型，使用泛型方式确认类型
* ** InvocationHandler**: 指定动态处理器，执行目标对象的方法时，会触发事件处理器的方法

**3. 测试结果**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo38czryj60qm0a1tdc02)

## 4. 小结

本小节学习了代理模式以及它的的实现，那么要求大家掌握的内容如下：

**1. 实现要素**

被代理类最少实现一个接口。

**2. 实现步骤**

1. `newProxyInstance` 方法获取代理类实例；
2. 编写代理实例方法；
3. 添加代理实例方法内逻辑；
4. invoke 方法中调用委托类，也就是被代理类的方法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
