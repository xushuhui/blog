# Spring Boot RabbitMQ 应用场景

## 1. 前言

消息队列是一个容器，可以对程序产生的消息进行存储。消息队列的主要用途是削峰、异步、解耦，我们用一个实际场景来解释下。

有一家果汁生产企业，张三是采购员，负责采购水果；李四、赵五是配送员，分别负责将苹果、香蕉配送到生产车间。

### 1.1 削峰

传统模式下，张三采购完成，回到公司后，联系李四、赵五配送采购的水果。但是随着公司业务量大增，张三一次性采购的水果，李四、赵五得需要几天才能配送完。所以需要一个仓库，张三采购完成直接放到仓库里，李四、赵五慢慢从仓库取出配送。

此处的仓库就是消息队列，张三是采购消息的生产者，李四、赵五是消费者。当生产的消息太多时，可以使用队列`削峰`，这样消费者可以慢慢处理消息。

### 1.2 异步

传统模式下，张三采购完成后，需要等待李四、赵五来取，实际上极大浪费了张三的时间。如果直接放入仓库，可以不必等待，直接进行下面的工作。也就是说，张三与李四、赵五的工作是`异步`的，减少了等待时间。

### 1.3 解耦

之前张三采购完成后，有责任通知李四、赵五来取。万一李四、赵五忘带手机，张三还得联系领导协调处理，说实话张三就是个大老粗，整天为这些破事烦得不行。

如果直接放入仓库，张三根本不用管李四、赵五的事情，感觉愉快极了。张三与李四、赵五的工作不再互相依赖，都变得更加简单了，这就是`解耦`。

## 2. RabbitMQ 简介

RabbitMQ 是非常出名的消息中间件，遵循 AMQP 协议，可以跨平台、跨语言使用。 RabbitMQ 具备低时延、高可用的特点，还有简洁易用的可视化管理界面，所以本节我们使用 RabbitMQ 来进行消息队列技术的演示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ke9yuj60hw09776702)

RabbitMQ 可视化管理界面

## 3. Spring Boot 实现

我们就针对上面的场景，使用 Spring Boot ，结合 RabbitMQ 来具体实现下水果采购、配送的管理。

### 3.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-rabbitmq，生成项目后导入 Eclipse 开发环境。

### 3.2 引入项目依赖

我们引入 Web 项目依赖与 AMQP 消息队列依赖。

**实例：**

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- AMQP 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-amqp</artifactId>
		</dependency>
```

### 3.3 配置 RabbitMQ 连接信息

项目创建后，通过 applicaiton.properties 配置 RabbitMQ 的链接信息。

**实例：**

```java
#地址
spring.rabbitmq.host=127.0.0.1
#端口 默认5672
spring.rabbitmq.port=5672
#用户名
spring.rabbitmq.username=guest
#密码
sprng.rabbitmq.password=guest
```

### 3.4 配置队列

首先配置两个队列，存储苹果采购消息、香蕉采购消息。

**实例：**

```java
/**
 * 消息队列配置类
 */
@Configuration
public class RabbitConfig {
	/**
	 * 苹果采购消息队列
	 */
	@Bean
	public Queue appleQueue() {
		return new Queue("apple-queue");
	}

	/**
	 * 香蕉采购消息队列
	 */
	@Bean
	public Queue bananaQueue() {
		return new Queue("banana-queue");
	}
}
```

### 3.5 配置交换机和绑定

如果消息直接发到队列的话，不够灵活， RabbitMQ 提供了交换机与绑定机制。

消息发送给交换机，交换机可以灵活地与队列进行绑定，这样消息就可以通过多种方式进入队列了。

**实例：**

```java
	/**
	 * 配置交换机
	 */
	@Bean
	TopicExchange exchangeTopic() {
		return new TopicExchange("exchange-topic");
	}

	/**
	 * 交换机绑定苹果采购消息队列
	 */
	@Bean
	Binding bindAppleQueue() {
		return BindingBuilder.bind(appleQueue()).to(exchangeTopic()).with("#.apple.#");
	}

	/**
	 * 交换机绑定香蕉采购消息队列
	 */
	@Bean
	Binding bindBananaQueue() {
		return BindingBuilder.bind(bananaQueue()).to(exchangeTopic()).with("#.banana.#");
	}
```

我们来详细解释下交换机与绑定的运行机制。

1. 我们配置了一个交换机 exchangeTopic ，它可以接收消息。
2. 交换机 exchangeTopic 绑定了两个队列，分别是 appleQueue 和 bananaQueue ，说明这两个队列在关注该交换机收到的消息。
3. 那么交换机 exchangeTopic 收到的消息到底会进入哪个队列呢，我们发现交换机的类型是 `TopicExchange` ，说明该交换机是`话题`交换机，队列应该是获取其感兴趣的话题相关的消息。
4. 当 appleQueue 队列绑定到交换机时，`with("#.apple.#")` 就表示 appleQueue 关心的是 apple 相关的话题；而 bananaQueue 关心的是 banana 相关的话题。
5. 所以可以推断出，消息在发送时，可以指定话题相关的信息，以便消息能被关注该话题的队列接收。

经过上面的分析，我们就知道了消息发送时通过携带话题信息，交换机会将该消息路由到关心该话题的队列中。

### 3.6 创建消费者

接下来，我们就可以定义消息的消费者李四、赵五了。他俩分别关心苹果采购消息和香蕉采购消息。也就是监听苹果消息队列和香蕉消息队列。

**实例：**

```java
/**
 * 消息队列接收
 */
@Component
public class RabbitReceiver {
	/**
	 * lisi负责监听apple-queue
	 */
	@RabbitListener(queues = "apple-queue")
	public void lisi(String msg) {
		System.out.println("李四知道:" + msg);
	}

	/**
	 * zhaowu负责监听banana-queue
	 */
	@RabbitListener(queues = "banana-queue")
	public void zhaowu(String msg) {
		System.out.println("赵五知道:" + msg);
	}
}
```

### 3.7 测试

运行启动类，从 RabbitMQ 管理界面可以看到已生成指定名称的队列了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2l96epj60hn0bsq6802)

RabbitMQ 已生成队列

此时我们定义一个控制器用于发起测试，直接使用 rabbitTemplate 发送消息即可。

**实例：**

```java
@RestController
public class TestController {
	@Autowired
	private RabbitTemplate rabbitTemplate;

	@GetMapping("/test")
	public void test() {
		// 发送消息 参数分别为：交换机名称、路由键、消息内容
		rabbitTemplate.convertAndSend("exchange-topic", "apple", "苹果来了10斤");
		rabbitTemplate.convertAndSend("exchange-topic", "banana", "香蕉来了5斤");
		rabbitTemplate.convertAndSend("exchange-topic", "apple.banana", "苹果来了8斤;香蕉来了20斤");
	}
}

```

convertAndSend() 方法的第 1 个参数表示交换机，第 2 个参数表示路由键（消息的话题），第 3 个是消息内容。

所以第 1 个消息会被 apple-queue 接收，第 2 个消息会被 banana-queue 接收，第 3 个消息会被两个队列接收。

我们启动项目，然后访问 `http://127.0.0.1:8080/test` ，控制台输出如下，验证成功。

```java
赵五知道:香蕉来了5斤
李四知道:苹果来了10斤
赵五知道:苹果来了8斤;香蕉来了20斤
李四知道:苹果来了8斤;香蕉来了20斤
```

## 4. 小结

本小节通过一个实际应用场景，演示了 Spring Boot 中使用 RabbitMQ 消息队列的方法。

至此， Spring Boot 的内容就全部结束了。纸上得来终觉浅，绝知此事要躬行。任何的实用技能都需要在不断练习与使用中感悟、完善、提升， Spring Boot 也不例外。

所以还没有使用 Spring Boot 的朋友，抓紧上手吧！

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
