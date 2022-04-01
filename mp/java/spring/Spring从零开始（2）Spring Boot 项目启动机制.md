# Spring Boot 项目启动机制

## 1. 前言

很多同学，学了很久的 Spring ，也用了很久的 Spring ，却还是不知道 Spring 是什么？Spring 中 XML / 注解 / Java 类三种配置方式，有什么区别和联系。

上面两个问题，正是理解 Spring Boot 的关键！

Spring 本质上是一个容器，里面存放的是 Java 对象，放入容器的 Java 对象被称为 Spring 组件（Bean）。

而 XML / 注解 / Java 类三种配置方式，只是形式不同，目的都是在容器中注册 Bean 。三种方式可以同时使用，只是需要注意， Bean 命名不要发生冲突。

当我们使用 Spring Boot 时会有变化吗？实际上，容器还是那个容器，配置也还是那三种配置。当然 Spring Boot 本身就是为了简化配置，所以基本不再使用 XML 配置方式了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2re8nsj60ji0azdhb02)

大海是鱼类的容器，地球是人类的容器，容器化的思想是普通存在的（图片来源于网络，版权归原作者所有）

我们打个比方， Spring 是插钥匙启动的轿车，而 Spring Boot 是无钥匙启动的轿车。功能和原理是几乎一样的， Spring Boot 更加简单方便而已。

## 2. Spring Boot 是如何启动的

Spring Boot 项目生成后，只有简简单单一个类，简单优雅，赏心悦目！

**实例：**

```java
@SpringBootApplication
public class SpringBootHelloApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootHelloApplication.class, args);
	}
}
```

我们来分析下这段代码， `public static void main` 是普通的 main 方法，是程序执行的入口。

`SpringApplication.run` 看字面意思就知道，这是 Spring 应用的启动方法，运行该行代码后， Spring 应用就跑起来了。

这个方法有两个参数， args 是命令行参数，此处没啥作用；另一个参数是 `SpringBootHelloApplication.class` ，包含类的信息。

这个类有啥信息啊？放眼看去，除了一个类名、一个静态方法外，并无其他。凭这些信息就能启动 Spring 应用？

等等，好像还有一个注解 `@SpringBootApplication` ，该注解是标注在类上的，属于类的信息。嗯，看来 Spring Boot 启动的秘密就在这个注解上了。

## 3. 神奇的 @SpringBootApplication 注解

我们来看看这个注解到底是何方神圣！在 Eclipse 中选中该注解，按 F3 即可查看其定义。

**实例：**

```java
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan(excludeFilters = { @Filter(type = FilterType.CUSTOM, classes = TypeExcludeFilter.class),
		@Filter(type = FilterType.CUSTOM, classes = AutoConfigurationExcludeFilter.class) })
public @interface SpringBootApplication {
}
```

看起来很复杂，其实就是一个组合注解，包含了多个注解的功能，咱们来分析一下。

首先是 `@SpringBootConfiguration` 注解，它继承自 `@Configuration` 注解，功能也跟 `@Configuration` 一样。它会将当前类标注为配置类了，我们在启动类中配置 Bean 就可以生效了。

其次是 `@ComponentScan` 注解，用来指定我们要扫描的包，以便发现 Bean 。注意在默认情况下， SpringBoot 扫描该注解标注类所在包及其子包。当我们的控制器、服务类等 Bean 放到不同的包中时，就需要通过 `@ComponentScan` 注解指定这些包，以便发现 Bean 。

最重要的是 `@EnableAutoConfiguration` 注解，用来启动自动配置。开启自动配置后， Spring Boot 会扫描项目中所有的配置类，然后根据配置信息启动 Spring 容器。

拥有了 `@SpringBootConfiguration` ，我们就拥有了一个可以拿来即用的 Spring 容器环境了。

## 4. 视频演示

## 5. 小结

Spring Boot 实际上就是 Spring 应用的快速开发版本，可以一键启动整个 Spring 容器供我们使用。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2roqonj60ji0d1mxt02)

 Spring Boot 准备好了，让我们开始启动吧！（图片来源于网络，版权归原作者所有）

我们运行下启动类，可以看到一个显眼的图案，它是 Spring Boot 的启动标志。接下来我们就可以，使用 Spring Boot ，来开发我们的应用了！

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2s4d5aj60jg05qq6302)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
