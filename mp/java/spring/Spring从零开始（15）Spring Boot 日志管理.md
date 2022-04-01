# Spring Boot 日志管理

## 1. 前言

谁能保证开发的软件系统没有问题？恐怕任何一个有经验的程序员都不敢承诺吧！

在软件的设计、开发阶段，大家都是尽心尽力去做好各项工作，期望能有一个满意的效果。

但是一个投入生产环境、拥有众多用户的软件系统必然是一个复杂的系统工程，不经历现实的检验，没有人能准确地知道它到底会不会有问题。

所以，日志是重要的，不可或缺的。日志是软件系统出现故障时，分析问题的主要依据。就像飞机的黑匣子，平时感觉毫不起眼，到了关键时刻必须要依靠它！

## 2. Spring Boot 日志管理

### 2.1 默认日志配置

Spring Boot 默认已经集成了日志功能，使用的是 logback 开源日志系统。

我们新建一个项目，Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-log。生成项目后导入 Eclipse 开发环境，然后运行启动类，可以清楚地看到控制台打印的日志信息。Spring Boot 日志默认级别是 INFO ，下图也输出了几条 INFO 级别的日志。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo260gb4j60na06842d02)

Spring Boot 项目启动时控制台输出的内容

Spring Boot 默认的日志输出内容含义如下：

* **日期时间**：精确到毫秒。
* **日志级别**：打印 ERROR 、 WARN 、 INFO 、 DEBUG 、 TRACE 等级别日志信息。
* **进程 ID**：当前项目进程 ID 。
* **分隔符**：`---` 是分隔符，分隔符后面代表具体的日志内容。
* **线程名**：方括号中间的内容表示线程名称。
* **类名**：当前日志打印所属的类名称。
* **日志内容**：开发人员设定的日志具体内容。

### 2.2 日志级别控制

有时候，我们想指定打印的日志的级别，可以通过配置文件来设置。

**实例：**

```java
# 设置日志级别
logging.level.root=WARN
```

上面的配置表示项目日志的记录级别为 WARN ，所以会打印 WARN 及优先级更高的 ERROR 级别的日志。此时我们编写一个测试类，看看具体打印日志的情况。

**实例：**

```java
@SpringBootTest
class LogTest {
	private Logger logger = LoggerFactory.getLogger(this.getClass());

	@Test
	void testPrintLog() {
		logger.trace("trace log");
		logger.debug("debug log");
		logger.info("info log");
		logger.warn("warn log");
		logger.error("error log");
	}
}
```

运行测试类，控制台打印内容如下，说明我们指定的日志级别生效了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo268akuj60rf055dja02)

控制台打印指定级别日志

> **Tips**: `logging.level.root=WARN` 中的 root 可以改为指定包名或类名，表示设置指定包或者类的日志级别。

### 2.3 输出日志文件

控制台日志保存的内容十分有限，大多数情况下我们需要将日志写入文件，便于追溯。

可以通过配置文件指定日志文件，如下配置会将日志打印到 `C:\\logs\\spring-boot-log.log` 文件中。

**实例：**

```java
# 设置日志文件
logging.file=C:\\logs\\spring-boot-log.log
```

也可以指定日志文件输出的目录， Spring Boot 项目会在指定输出目录下新建 `spring.log` 文件，并在文件中写入日志。

**实例：**

```java
# 设置日志目录
logging.path=C:\\logs
```

> **Tips**：如果同时配置了 `logging.file` 和 `ogging.path` ，则只有 `logging.file` 生效。

### 2.4 使用 lombok 插件简化日志代码

在上面的示例中，如果要打印日志，需要添加一行代码 `private Logger logger = LoggerFactory.getLogger(this.getClass());` 还是比较麻烦的。我们可以安装 lombok 插件，使用一个注解代替这行代码。

#### 2.4.1 下载 lombok 插件

从 [lombok 下载链接](https://mvnrepository.com/artifact/org.projectlombok/lombok/1.18.12) 下载 lombok 插件。

#### 2.4.2 安装 lombok 插件

双击打开 lombok.jar ，点击 Specify Location 按钮，选择 eclipse.exe ，然后点击 Install 安装插件。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo26hxvxj60n40d60z302)

lombok 插件安装

#### 2.4.3 引入 lombok 依赖

lombok 安装后还需要引入依赖项，在 pom.xml 中添加如下依赖即可。

**实例：**

```java
		<!-- lombok -->
		<dependency>
			<groupId>org.projectlombok</groupId>
			<artifactId>lombok</artifactId>
			<version>1.18.12</version>
			<scope>provided</scope>
		</dependency>
```

#### 2.4.4 使用注解输出日志

此时，可以直接给类添加注解，然后就能直接输出日志了。

**实例：**

```java
@SpringBootTest
@Slf4j // 添加日志输出注解
class LogTest {
	// 不再需要定义 logger
	// private Logger logger = LoggerFactory.getLogger(this.getClass());

	@Test
	void testPrintLog() {
		// 直接使用log输出日志
		log.trace("trace log");
		log.debug("debug log");
		log.info("info log");
		log.warn("warn log");
		log.error("error log");
	}
}
```

> **Tips**：lombok 插件的功能比较强大，不仅可以简化日志模板代码，还可以自动生成常用的 getter /setter/toString 等模板代码，感兴趣的同学可以查阅相关资料。

## 3. 自定义日志配置

Spring Boot 也支持自定义日志配置，可以直接采用指定日志系统的配置文件，如 logback 、 log4j 。以 logback 为例，可以直接在 application.properties 文件中指定 logback 配置文件。

**实例：**

```java
# 指定logback配置文件，位于resources目录下
logging.config=classpath:logback-spring.xml
```

> **Tips**：使用 logback 日志系统后，日志级别与日志文件等信息都可以使用 logback-spring.xml 文件设置，不再需要从 properties 文件中设置了。

在生产环境，我们希望指定日志保存的位置，另外日志不能无限制一直保存，一般情况下保存最近 30 天左右的日志即可。这些都可以在 logback-spring.xml 文件中指定，此处给出一个完整实例供大家参考。

**实例：**

```java
<?xml version="1.0" encoding="UTF-8"?>
<!-- logback 配置 -->
<configuration>
	<!-- 输出到控制台 -->
	<appender name="STDOUT"
		class="ch.qos.logback.core.ConsoleAppender">
		<encoder
			class="ch.qos.logback.classic.encoder.PatternLayoutEncoder">
			<!--格式化输出:%d表示日期;%thread表示线程名;%-5level:左对齐并固定显示5个字符;%msg:日志消息;%n:换行符; -->
			<pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger{50} -
				%msg%n</pattern>
		</encoder>
	</appender>
	<!-- 输出到文件 -->
	<appender name="FILE"
		class="ch.qos.logback.core.rolling.RollingFileAppender">
		<!-- 正在打印的日志文件 -->
		<File>C:/logs/spring-boot-log.log</File>
		<encoder>
			<!--格式化输出:%d表示日期;%thread表示线程名;%-5level:左对齐并固定显示5个字符;%msg:日志消息;%n:换行符; -->
			<pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger{50} -
				%msg%n
			</pattern>
		</encoder>
		<!-- 日志文件的滚动策略 -->
		<rollingPolicy
			class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
			<!-- 日志归档 -->
			<fileNamePattern>C:/logs/spring-boot-log-%d{yyyy-MM-dd}.log
			</fileNamePattern>
			<!-- 保留30天日志 -->
			<maxHistory>30</maxHistory>
		</rollingPolicy>
	</appender>
	<!-- 指定日志输出的级别 -->
	<root level="INFO">
		<appender-ref ref="STDOUT" />
		<appender-ref ref="FILE" />
	</root>
</configuration>
```

logback 日志系统的功能比较全面，网上可以查询到的资料也非常多，大家可以自行查阅以做进一步的了解。

## 4. 小结

Spring Boot 项目可以使用简单的几个配置，实现日志的打印，并设置相应的级别、日志文件等信息。

如果想要对日志的方方面面进行设定，也可以快速地集成常见的日志系统如 logback 、log4j 。

日志系统对生产环境项目来说是不可或缺的，大家可以选择使用 Spring Boot 集成一种自己用起来顺手的日志系统。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
