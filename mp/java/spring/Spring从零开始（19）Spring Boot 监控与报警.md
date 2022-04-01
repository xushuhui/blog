# Spring Boot 监控与报警

## 1. 前言

因为公司开发的项目多、为客户部署的项目实例多。工作中我们都会经常遇到，由于某个客户的项目突然无法访问，一堆研发、售后部门的同事火急火燎处理问题的场景。

所以我非常希望能够实现这样的功能：

* 能够有一个界面，监控所有关注的项目实例运行状态。
* 对于某个项目实例来说，可以监控该实例的各项运行参数，例如内存占用情况、磁盘使用情况、数据库连接情况。
* 项目实例因各种原因关闭时，可以自动报警。

在很长一段时间内，我感觉要实现这些功能比较复杂。后来我稍微研究了下，在 Spring Boot 中实现这种监控和报警的功能非常简单，那还等什么呢， Let’s Go ！

## 2. 可视化监控

可以直接利用 Spring Boot Admin 实现可视化监控，此时至少需要两个项目实例，一个是监控的管理端，一个是被监控的客户端。

### 2.1 构建监控管理端项目

打开 Spring Initializr ， Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-monitor-manager，生成项目后导入 Eclipse 开发环境。

### 2.2 引入管理端项目依赖

监控管理端需要使用网页展示监控信息，所以引入 Web 依赖，另外添加 Spring Boot Admin 管理端依赖项。

**实例：**

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- Spring Boot Admin 管理端依赖项 -->
		<dependency>
			<groupId>de.codecentric</groupId>
			<artifactId>spring-boot-admin-starter-server</artifactId>
			<version>2.2.3</version>
		</dependency>
```

### 2.3 开启监控管理端

在启动类上添加 @EnableAdminServer 注解开启 Spring Boot Admin 监控管理功能，代码如下：

**实例：**

```java
@SpringBootApplication
@EnableAdminServer // 开启监控管理
public class SpringBootMonitorManagerApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootMonitorManagerApplication.class, args);
	}
}
```

然后运行启动类，访问 `http://127.0.0.1:8080` 会发现界面上已经显示监控信息了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2bgbvjj60ra079q4d02)

Spring Boot Admin 监控管理页面

### 2.4 构建监控客户端项目

打开 Spring Initializr ， Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-monitor-client，生成项目后导入 Eclipse 开发环境。

### 2.5 引入客户端项目依赖

直接引入 Web 依赖和监控客户端依赖。

实例：

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- Spring Boot Admin监控客户端依赖 -->
		<dependency>
			<groupId>de.codecentric</groupId>
			<artifactId>spring-boot-admin-starter-client</artifactId>
			<version>2.2.3</version>
		</dependency>
```

### 2.6 修改客户端配置

修改客户端的配置文件 application.properties ，以便指定客户端指向的服务端的地址。由于刚刚服务端已经占用了 8080 端口，所以将客户端的端口设置为 8091 。

还有一个必要设置是客户端的名称，当我们监控的项目实例比较多时，需要通过客户端名称来区分。

**实例：**

```java
# 配置端口
server.port=8091
# 配置监控管理端地址
spring.boot.admin.client.url=http://127.0.0.1:8080
# 客户端的名称，用于区分不同的客户端
spring.boot.admin.client.instance.name=CLIENT1
```

> **TIps**：此处指定监控管理端地址使用的是 `spring.boot.admin.client.url` ，我个人认为应使用 `spring.boot.admin.server.url` 更加合理。当然大家不用纠结于此，此处只是特别提示。

### 2.7 测试监控效果

启动客户端程序，然后刷新服务端网页，会发现监控管理页面已经显示了客户端信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2btpdnj60rc09aabx02)

监控页面展示客户端信息

此时我们关闭客户端程序，然后稍等一会刷新下监控管理页面（注意服务端发现客户端离线是需要一定时间的），会发现监控管理页已经显示了离线项目实例信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2c84fbj60rd09babt02)

监控页面展示离线客户端信息

## 3. 监控实例运行参数

使用 Spring Boot Admin 后，默认显示的项目实例信息比较少。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2cg09ej60pj0au0vm02)

CLIENT1 实例默认显示信息

默认设置主要是为了保证项目实例的安全性，只展示了非常少的信息，我们可以通过配置文件指定展示哪些信息，如下。

**实例：**

```java
# 配置客户端展示哪些信息，*表示展示全部信息
management.endpoints.web.exposure.include=*
```

此时刷新监控管理页，会发现已经展示各类运行参数信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2crjzaj60q00min4h02)

全面展示项目实例运行参数信息

> **Tips**：不同项目需要展示的信息不同，可以通过 `management.endpoints.web.exposure.include` 配置项进行设置，感兴趣的同学可以自行查阅相关资料。

## 4. 自动报警

可视化监控提供了全面了解项目运行状况的手段，但是我们不可能 7*24 小时盯着界面看哪个应用离线了。

最妙的效果是，项目离线时自动通知售后、运维等相关技术人员。

Spring Boot Admin 也提供了自动报警的功能，简直太完美了，接下来我们来实现下。

### 4.1 引入依赖项

为监控服务端项目引入邮件依赖。

**实例：**

```java
		<!-- 邮件依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-mail</artifactId>
		</dependency>
```

### 4.2 配置邮件发送所需信息

只需要配置常规的邮件收发信息即可。

**实例：**

```java
# 网易邮箱发件服务器
spring.mail.host=smtp.163.com
# 网易邮箱发件端口
spring.mail.prot=25
# 发件人账号
spring.mail.username=taqsxxkj@163.com
# 发件授权密码，注意授权码是用于登录第三方邮件客户端的专用密码
spring.mail.password=123456
spring.mail.properties.mail.smtp.socketFactory.class=javax.net.ssl.SSLSocketFactory
# Spring Boot Admin 发件收件信息
spring.boot.admin.notify.mail.from=taqsxxkj@163.com
spring.boot.admin.notify.mail.to=taqsxxkj@163.com
spring.boot.admin.notify.mail.cc=taqsxxkj@163.com
```

此处特别注意发件授权密码不是普通邮箱的登录密码，而是授权密码，以网易邮箱为例在下图位置设置。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2d1iulj60km0j0jxh02)

网易邮箱授权密码设置

### 4.3 项目实例离线邮件报警

启动监控服务端和客户端，然后关闭客户端，稍等一会检查指定的报警接收邮箱，就会发现已收到报警邮件了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2de1t4j60kw086tbd02)

项目实例离线邮件报警内容

## 5. 视频演示

## 6. 小结

报警和监控机制可以为上线运行的项目提供额外的保障机制。

* 可视化监控可以提供一种总揽全局的监控视角，众多项目运行状况一目了然，做到心里有底。
* 对于重点项目，可以定期监控项目实例的详细运行参数。通过分析参数信息，预测性能瓶颈发生的时间节点，提前采取扩容等措施，防患于未然。
* 当发生严重错误，或者由于人为误操作导致项目离线时，指定邮箱能够及时收到报警信息。技术人员及时排查处理，作为项目运维方心里更有底。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
