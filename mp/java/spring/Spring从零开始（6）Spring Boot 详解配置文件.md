# Spring Boot 详解配置文件

## 1. 前言

Spring Boot 可以在零配置的情况下使用，但是不代表 Spring Boot 完全不需要配置文件。

举一个简单的例子， Spring Boot 开发的 Web 项目默认的启动端口是 8080 。如果我们想更换启动端口，通过配置文件去修改是比较好的。如果放到代码中，修改一个端口都要重新编译下程序，岂不烦哉？

配置文件不是必须的，但是如果想实现一些个性化的功能，还是需要用到配置文件的。本篇就讲下 Spring Boot 中使用配置文件的常用场景。

## 2. 构建演示 Web 项目

为了便于演示，我们先构建一个 Web 项目。

### 2.1 使用 Spring Initializr 构建一个 Spring Boot 应用

Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-profile` ，生成项目后导入 Eclipse 开发环境。

### 2.2 修改 pom.xml

引入 Web 项目依赖，同时开启热部署便于修改测试。

**实例：**

```java
		<!-- 热部署 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-devtools</artifactId>
		</dependency>
		<!-- web -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
```

### 2.3 编写控制器用于测试

**实例：**

```java
@RestController // 标注为控制器，且返回值序列化为json
public class HelloController {
	@GetMapping("/hello") // 响应get请求，匹配的请求路径为/hello
	public Map hello() {
		Map<String, String> map = new HashMap<String, String>();
		map.put("test", "content for test");
		return map;
	}
}
```

### 2.4 启动项目

访问 `http://127.0.0.1:8080/hello` ，效果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3m1ir4j60jf04djri02)

浏览器显示返回数据

## 3. 修改项目启动配置

我们运行启动类，启动 `spring-boot-profile` 应用，控制台会发现如下提示：

```java
Tomcat started on port(s): 8080 (http) with context path ''
```

可以看出， Spring Boot 应用默认启动端口是 8080 ，默认项目路径是空。

我们可以通过修改 `resources/application.properties` 来自定义项目启动配置：

**实例：**

```java
# 启动端口
server.port=8000
# 项目路径
server.servlet.context-path=/spring-boot-profile
```

再次启动应用，控制台提示变为：

```java
Tomcat started on port(s): 8000 (http) with context path '/spring-boot-profile'
```

此时项目对应的访问路径为： `http://127.0.0.1:8000/spring-boot-profile` , 使用浏览器访问效果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3mc7slj60jf06p3yo02)

浏览器显示返回数据

## 4. 配置文件格式

Spring Boot 支持两种格式的配置文件，即 `.properties` 文件和 `.yml` 配置文件。

上面的配置使用 `.yml` 则为：

**实例：**

```java
server:
 port: 8000
 servlet:
   context-path: /spring-boot-profile
```

`.properties` 配置使用顿号分割语义，而 `.yml` 配置使用缩进分割语义。这两种配置文件没有本质区别，只是格式不同。

## 5. 自定义配置项

我们还可以在配置文件中使用自定义配置，例如我们开发了一个微信公众号后台应用，需要在程序中配置公众号的 appid 和 secret 。

配置文件如下：

**实例：**

```java
# 公众号appid
wxmp.appid=111
# 公众号secret
wxmp.secret=222
```

我们定义一个组件，通过 `@Value` 注解注入配置项的值。

**实例：**

```java
/**
 * 微信公众号参数
 */
@Component//注册为组件
public class WxMpParam {
	@Value("${wxmp.appid}")//注入wxmp.appid配置项
	private String appid;
	@Value("${wxmp.secret}")//注入wxmp.secret配置项
	private String secret;
  //省略get set方法
}
```

通过控制器测试配置项是否注入成功。

**实例：**

```java
@RestController
public class HelloController {
	@Autowired
	private WxMpParam wxMpParam;
	@GetMapping("/hello")
	public Map hello() {
		Map<String, String> map = new HashMap<String, String>();
		map.put("appid",wxMpParam.getAppid());
		map.put("secret",wxMpParam.getSecret());
		return map;
	}
}
```

此时我们访问 `http://127.0.0.1:8000/spring-boot-profile/hello` ，浏览器显示如下，说明我们的配置注入成功。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3mnn2cj60jg049aa802)

浏览器显示返回数据

## 6. 配置项自动注入对象

如果参数很多，一一指定对象属性和配置项的关联非常麻烦。可以通过设定对象与配置项的对应关系，来实现配置项的自动注入。

**实例：**

```java
@Component // 注册为组件
@EnableConfigurationProperties // 启用配置自动注入功能
@ConfigurationProperties(prefix = "wxmp") // 指定类对应的配置项前缀
public class WxMpParam {
	private String appid;// 对应到wxmp.appid
	private String secret; // 对应到wxmp.secret
  //省略 get set
}
```

在上面的代码中，通过 `prefix = "wxmp"` 指定了关联配置的前缀，属性 appid 即关联到前缀 + 属性名为 wxmp.appid 的配置项。同理，属性 secret 关联到 wxmp.secret 配置项。

## 7. 在配置文件中使用随机数

配置文件中使用随机数也是比较常见的场景，尤其启动多个客户端时，希望指定一个启动端口的范围，例如 10 - 20 ，可配置如下：

**实例：**

```java
# 配置端口为1-20间的随机数
server.port=${random.int[10,20]}
```

这样我可以连续启动四个客户端，启动端口分别是 12 、 13 、 17 、 19 ，可见是随机的，而且在我指定的范围内波动。

## 8. 自定义配置文件

有时候参数太多，都放到一个配置文件中太乱了，我们会希望将配置分到不同文件中，然后每个文件保存不同配置。

例如上面微信公众号配置，我们单独建立一个 `wxmp.properties` 文件，内容如下：

**实例：**

```java
# wxmp.properties配置文件

# 公众号的appid
wxmp.appid=111
# 公众号的secret
wxmp.secret=222
```

WxMpParam 代码如下：

**实例：**

```java
/**
* 微信公众号参数
*/
@Component // 注册为组件
@PropertySource(value = "classpath:wxmp.properties", encoding = "utf-8") // 指定配置文件及编码
public class WxMpParam {
   @Value("${wxmp.appid}")
   private String appid;
   @Value("${wxmp.secret}")
   private String secret;
}
```

## 9. 配置项引用

Spring Boot 配置项是可以引用其他配置项的值的，这个稍微提一下，例如：

**实例：**

```java
# wxmp.properties

# 公众号的appid
wxmp.appid=111
# 公众号的secret,值为111222
wxmp.secret=${wxmp.appid}222
```

## 10. 小结

对一个 Spring Boot 应用而言。

* 如果配置项比较少，直接全部写在 `application.properties` 。

* 如果配置项很多，可以划分为若干配置文件。

* 如果很多自定义配置拥有相同的前缀，可以指定前缀，让配置项自动注入对象中。

* Spring Boot 提供了多变的配置文件使用机制，我们根据具体场景灵活使用即可。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
