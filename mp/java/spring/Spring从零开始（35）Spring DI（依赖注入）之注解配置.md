# Spring DI（依赖注入）之注解配置

## 1. 前言

上一节，我们通过 xml 文件的配置方式，实现了对多种依赖类型的注入，当然体会到了 xml 文件配置方式的弊端：有一点麻烦。

依赖注入是有两种方式，一种是 xml ，另外一种就是**注解的配置方式**。

本节，我们演示下通过注解配置这种方式来实现注入依赖。

来吧 ，直入主题，莫浪费大好光阴…

## 2. 工程实例

### 2.1 注解的介绍

在正式使用注解之前，我们首先介绍下注解语法以及它的作用。

* **@Autowired：** 此注解自动按照类型注入。从容器中寻找符合依赖类型的实例，当使用该注解注入属性时，set 方法可以省略。但是因为按照类型匹配，如果容器中有多个匹配的类型，会抛出异常，需要指定引入的实例 id。如果找不到匹配的实例，那么也会抛出异常；
* **@Qualifier：** 此注解不能单独使用，它的作用是在按照类型注入的基础之上，再按照 Bean 的 id 注入。所以如果是使用了 @Autowire 注解自动注入，但是容器中却有多个匹配的实例，可以搭配此注解，指定需要注入的实例 id；
* **@Resource** 此注解的作用是指定依赖按照 id 注入，还是按照类型注入。当只使用注解，但是不指定注入方式的时候，默认按照 id 注入，找不到再按照类型注入。

### 2.2 @Autowired 注解

1. 为了测试效果，我们创建 Service 和 Dao 两个类， Dao 作为 Service 的依赖。代码如下：

```java
//service实现类的代码
@Service
public class UserServiceImpl implements  UserService {

    @Autowired
    private UserDao userDao;

    public void saveUser() {

        System.out.println("执行service中的保存逻辑");
    }
}

//dao实现类的代码
@Repository
public class UserDaoImpl implements  UserDao {

    public void saveUser() {
        System.out.println("执行dao的保存方法");
    }
}
```

**代码解释：**

上面代码可以看到，两个类的实例化方式都是通过注解注入到容器， 并且在 service 实现类中的 userDao 属性上面加了注解 `@Autowired`。

我们首先测试下：能否通过这个注解，实现依赖注入，另外再测试下它是否是按照类型注入。

2. 配置文件的内容为注解实现 IoC。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo305b9pj60v70cjqdr02)

**配置文件解释：** 注解实现 IoC 的章节说过，需要通过组件扫描来实例化容器。

3. 编写测试代码

```java
public class SpringAnTest {

    public static void main(String[] args) {
        ApplicationContext context =
                new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        UserService userService = context.getBean(UserService.class);
        userService.saveUser();

    }
}
```

**测试结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo30kyrmj61110hq4ad02)

**结果解释**

可以看到 service 中的代码执行，并且通过 dao 的示例调用的方法也执行了，那么说明 `@Autowired` 注解实现了属性 userDao 的注入。

当然这种操作是小儿科，没有一个同学觉得他有什么。 我们验证下它的特点：set 方法我们是省略了，那么它是否按照类型注入的呢？

如果我们的实现类中有多个 userDao 接口的实现类呢，又该如何呢？

4. 添加 UserDaoImpl2 一样实现 userDao 的接口，代码如下：

```java
@Repository
public class UserDaoImpl1 implements  UserDao {

    public void saveUser() {
        System.out.println("执行dao1的保存方法");
    }
}
```

**测试结果**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo311pb9j61e70jgk3i02)

结果解释：

可以看到上面控制台打印的异常堆栈信息，清楚的提示错误原因，没有指定的 bean 实例 UserDao 类型的，期待单个 bean 匹配，但是找到了两个。

一个是 `userDaoImpl` 一个是 `userDaoImpl1`。看到这可以证明：`@Autowired` 注解是按照类型注入，如果匹配的类型多了就会报错。

**疑问导出：**

难道使用了 Spring 框架以后，我们的接口只能有一个实现类吗？ 当然不可能，毕竟我们看 Spring 的源码的时候 已经看到了，很多的接口对应一大堆的实现类。

那么，针对这种多个接口实例的情况，怎么解决的呢？继续我们注解的学习。

### 2.3 @Qualifier 注解

1. 此注解的作用，我们介绍过了，这里再看一下：它的作用是在按照类型注入的基础之上，再按照 Bean 的 id 注入，不能单独使用，搭配上面的 `@Autiwired` 注解。

在两个实现类的基础之上改造代码如下：

```java
@Service
public class UserServiceImpl implements  UserService {

    @Qualifier("userDaoImpl")
    @Autowired
    private UserDao userDao;

    public void saveUser() {

        System.out.println("执行service中的保存逻辑");
        userDao.saveUser();
    }
}
```

**代码解释**

在属性注入的地方，通过注解 `@Qualifier` 的参数，指定了注入的 bean 实例 id 为 userDaoImpl。

2. 测试方法继续执行查看结果。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo31g15xj60x809lwlp02)

**结果解释：**

那么可以看到，我们的方法正常执行，而且执行的就是 userDaoImpl 中的方法。

3. 继续改造 service 的代码如下，将 `@Qualifier` 注解中的值换成 userDaoImpl1 以后再看看结果。

```java
@Service
public class UserServiceImpl implements  UserService {

	@Qualifier("userDaoImpl1")
	@Autowired
	private UserDao userDao;

	public void saveUser() {

   	 	System.out.println("执行service中的保存逻辑");
    	userDao.saveUser();
	}
```

}

结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3236nrj60x80aa0zf02)

通过修改 `@Qualifier` 注解中 id 的属性值 ，可以分别注入不同的实现类，那么证明了 `@Qualifier` 注解的作用。

### 2.4 @Resources 注解

此注解的作用是指定依赖按照 id 注入，还是按照类型注入。当只使用注解，但是不指定注入方式的时候，默认按照 id 注入，找不到时再按照类型注入。

语法如下：

```java
@Resource   //默认按照 id 为 userDao的bean实例注入
@Resource(name="userDao")  //按照 id 为 userDao的bean实例注入
@Resource(type="UserDao")  //按照 类型 为 UserDao的bean实例注入

```

这里就只做个语法的介绍，注解的使用大同小异，大家按照上方步骤自行测试即可。

## 3. 小结

本节重点讲解注解注入依赖的使用，咱们做个总结：

1. 常用的注解有 3 种 ：`@Autowired``@Qualifier``@Resources`；

2. 注解注入的形式两种 ： 按照 bean 的 id 注入，或者按照 bean 的类型注入；

3. 哪种注解的目的，都是为了成功的注入使用的依赖，所以为了我们的开发服务，大家灵活使用即可。

不费力气就能得到的… 只有年龄。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
