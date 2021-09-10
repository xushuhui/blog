# Spring Boot 定时任务

## 1. 前言

定时任务绝对是实际项目中的刚需。

* 我们想监控一个重点服务的运行状态，可以每隔 1 分钟调用下该服务的心跳接口，调用失败时即发出告警信息；
* 我们想每天凌晨的时候，将所有商品的库存置满，以免早上忘记添加库存影响销售；
* 我们想在每个周六的某个时段进行打折促销。

在以上的案例中，或者是指定时间间隔，或者是指定时间节点，按设定的任务进行某种操作，这就是定时任务了。

在 Spring Boot 中实现定时任务简单而灵活，本节我们来体验下。

## 2. Spring Task 定时任务

Spring Task 是 Spring Boot 内置的定时任务模块，可以满足大部分的定时任务场景需求。

通过为方法添加一个简单的注解，即可按设定的规则定时执行该方法。

下面就演示下 Spring Boot 中使用 Spring Task 的具体方法。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-task ，生成项目后导入 Eclipse 开发环境。

### 2.2 开启定时任务

在启动类上添加 `@EnableScheduling` 注解，开启定时任务功能。

**实例：**

```java
@SpringBootApplication
@EnableScheduling // 开启定时任务
public class SpringBootTaskApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootTaskApplication.class, args);
	}
}
```

### 2.3 通过注解设定定时任务

新建 MySpringTask 任务类，添加 `@Component` 注解注册 Spring 组件，定时任务方法需要在 Spring 组件类才能生效。

注意类中方法添加了 `@Scheduled` 注解，所以会按照 `@Scheduled` 注解参数指定的规则定时执行。

**实例：**

```java
/**
 * 任务类
 */
@Component
public class MySpringTask {
	/**
	 * 每2秒执行1次
	 */
	@Scheduled(fixedRate = 2000)
	public void fixedRateMethod() throws InterruptedException {
		System.out.println("fixedRateMethod:" + new Date());
		Thread.sleep(1000);
	}
}
```

上面例子执行情况如下，可见是每隔 2 秒执行 1 次。

```java
fixedRateMethod:Fri May 15 22:04:52 CST 2020
fixedRateMethod:Fri May 15 22:04:54 CST 2020
fixedRateMethod:Fri May 15 22:04:56 CST 2020
```

**实例：**

```java
/**
 * 任务类
 */
@Component
public class MySpringTask {
	/**
	 * 执行结束2秒后执行下次任务
	 */
	@Scheduled(fixedDelay = 2000)
	public void fixedDelayMethod() throws InterruptedException {
		System.out.println("fixedDelayMethod:" + new Date());
		Thread.sleep(1000);
	}
}
```

上面的例子执行情况如下，每次打印后先等待 1 秒，然后方法执行结束 2 秒后再次执行任务，所以是每 3 秒打印 1 行内容。

```java
fixedDelayMethod:Fri May 15 22:08:26 CST 2020
fixedDelayMethod:Fri May 15 22:08:29 CST 2020
fixedDelayMethod:Fri May 15 22:08:32 CST 2020
```

### 2.4 使用 Cron 表达式

`@Scheduled` 也支持使用 Cron 表达式， Cron 表达式可以非常灵活地设置定时任务的执行时间。以本节开头的两个需求为例：

* 我们想监控一个重点服务的运行状态，可以每隔 1 分钟调用下该服务的心跳接口，调用失败时即发出告警信息；
* 我们想在每天凌晨的时候，将所有商品的库存置满，以免早上忘记添加库存影响销售。

对应的定时任务实现如下：

**实例：**

```java
/**
 * 任务类
 */
@Component
public class MySpringTask {
	/**
	 * 在每分钟的00秒执行
	 */
	@Scheduled(cron = "0 * * * * ?")
	public void jump() throws InterruptedException {
		System.out.println("心跳检测:" + new Date());
	}
	/**
	 * 在每天的00:00:00执行
	 */
	@Scheduled(cron = "0 0 0 * * ?")
	public void stock() throws InterruptedException {
		System.out.println("置满库存:" + new Date());
	}
}
```

Cron 表达式并不难理解，从左到右一共 6 个位置，分别代表秒、时、分、日、月、星期，以秒为例：

* 如果该位置上是 `0` ，表示在第 0 秒执行；
* 如果该位置上是 `*` ，表示每秒都会执行；
* 如果该位置上是 `?` ，表示该位置的取值不影响定时任务，由于月份中的日和星期可能会发生意义冲突，所以日、 星期中需要有一个配置为 `?` 。

按照上面的理解，`cron = "0 * * * * ?"` 表示在每分钟的 00 秒执行、`cron = "0 0 0 * * ?"` 表示在每天的 00:00:00 执行。

> **Tips**：Cron 表达式的描述能力很强，此处只是简单提及，感兴趣的同学可以自行查阅相关资料了解更多信息。

## 3. Quartz 定时任务

Spring Task 已经可以满足绝大多数项目对定时任务的需求，但是在企业级应用这个领域，还有更加强大灵活的 Quartz 框架可供选择。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo29yck3j60l506ndj502)

Quartz 官网介绍：企业级的任务调度框架

举个例子，当我们想根据数据库中的配置，动态地指定商品打折的时间区间时，就可以利用 Quartz 框架来实现。 OK ，接下来我们就来具体完整实现下。

### 3.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-quartz ，生成项目后导入 Eclipse 开发环境。

### 3.2 引入项目依赖

需要引入 Quartz 框架相关依赖。

**实例：**

```java
		<!-- Quartz -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-quartz</artifactId>
		</dependency>
```

### 3.3 开启定时任务

同样需要，在启动类上添加 `@EnableScheduling` 注解，开启定时任务功能。

**实例：**

```java
@SpringBootApplication
@EnableScheduling // 开启定时任务
public class SpringBootQuartzApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootQuartzApplication.class, args);
	}
}
```

### 3.4 Quartz 定时任务开发

Quartz 定时任务需要通过 Job 、 Trigger 、 JobDetail 来设置。

* **Job**：具体任务操作类
* ** Trigger**：触发器，设定执行任务的时间
* ** JobDetail**：指定触发器执行的具体任务类及方法

我们先开发一个 Job 组件：

**实例：**

```java
/**
 * 打折任务
 */
@Component // 注册到容器中
public class DiscountJob {
	/**
	 * 执行打折
	 */
	public void execute() {
		System.out.println("更新数据库中商品价格，统一打5折");
	}
}
```

然后在配置类中设定 Trigger 及 JobDetail 。

**实例：**

```java
/**
 * 定时任务配置
 */
@Configuration
public class QuartzConfig {
	/**
	 * 配置JobDetail工厂组件，生成的JobDetail指向discountJob的execute()方法
	 */
	@Bean
	MethodInvokingJobDetailFactoryBean jobFactoryBean() {
		MethodInvokingJobDetailFactoryBean bean = new MethodInvokingJobDetailFactoryBean();
		bean.setTargetBeanName("discountJob");
		bean.setTargetMethod("execute");
		return bean;
	}
	/**
	 * 触发器工厂
	 */
	@Bean
	CronTriggerFactoryBean cronTrigger() {
		CronTriggerFactoryBean bean = new CronTriggerFactoryBean();
		// Corn表达式设定执行时间规则
		bean.setCronExpression("0 0 8 ? * 7");
		// 执行JobDetail
		bean.setJobDetail(jobFactoryBean().getObject());
		return bean;
	}
}
```

具体分析下上面的代码：

1. 触发器设定的 Corn 表达式为 `0 0 8 ? * 7` ，表示每周六的 08:00:00 执行 1 次；
2. 触发器指定的 JobDetail 为 jobFactoryBean 工厂的一个对象，而 jobFactoryBean 指定的对象及方法为 discountJob 与 execute () ；
3. 所以每周六的 8 点，就会运行 discountJob 组件的 execute () 方法 1 次；
4. Corn 表达式和执行任务、方法均以参数形式存在，这就意味着我们完全可以根据文件或数据库配置动态地调整执行时间和执行的任务；
5. 最后，周六 8 点的时候，商品都打了 5 折，别忘了促销结束的时候恢复价格啊。

## 4. 小结

Spring Boot 可以利用一个简单的注解，快速实现定时任务的功能。

说实话我第一次使用 `@Scheduled` 注解时，完全被这种开箱即用型的简洁震撼了，我的感受是：似乎不能更加简洁了。

如果感觉 Spring Task 提供的定时任务机制还不足以满足需求，Spring Boot 还可以方便地集成 Quartz 框架来帮忙。

开箱即用满足不了，还可以即插即用，确实够人性化的。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
