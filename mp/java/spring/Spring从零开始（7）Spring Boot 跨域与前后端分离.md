# Spring Boot 跨域与前后端分离

## 1. 前言

目前，使用 Spring Boot 进行前后端分离项目开发，应该是主流做法了。这种方式，在开发、测试阶段，都比较方便。

开发阶段，项目组定义好接口规范后，前端按规范开发前端页面，后端按规范编写后端接口，职责分明。

测试阶段，后端是独立项目，可以进行单元测试。前端可以随时使用最新版本的后端程序进行实际测试。

前后端分离的模式，有着很多的优越性，所以造就了它的流行。

## 2. 技术选型

本篇我们通过商品浏览项目实例，展现前后端分离项目的开发、测试全流程。

技术选型方面，后端毫无疑问选择 Spring Boot ，接口风格采用 RESTful 标准。前端则使用简单的 HTML + Bootstrap + jQuery ，并通过 jQuery 的 $.ajax 方法访问后端接口。

## 3. 后端开发流程

### 3.1 使用 Spring Initializr 构建一个 Spring Boot 应用

Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-cors` ，生成项目后导入 Eclipse 开发环境。

### 3.2 修改 pom.xml

引入 Web 项目依赖，开启热部署便于修改测试。

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

### 3.3 定义商品类、商品服务类、商品控制器类

在控制器类中编写获取商品列表的接口供前端调用。

**实例：**

```java
/**
 * 商品类
 */
public class GoodsDo {
	/**
	 * 商品id
	 */
	private Long id;
	/**
	 * 商品名称
	 */
	private String name;
	/**
	 * 商品价格
	 */
	private String price;
	/**
	 * 商品图片
	 */
	private String pic;
  //省略get set
}
```

**实例：**

```java
/**
 * 商品服务类
 */
@Service // 注册为服务类
public class GoodsService {
	/**
	 * 获取商品列表
	 */
	public List<GoodsDo> getGoodsList() {
		List<GoodsDo> goodsList = new ArrayList<GoodsDo>();//模拟从数据库查询出的结果返回
		GoodsDo goods = new GoodsDo();
		goods.setId(1L);
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		goodsList.add(goods);
		return goodsList;
	}
}
```

**实例：**

```java
/**
 * 商品控制器类
 */
@RestController
public class GoodsController {
	@Autowired
	private GoodsService goodsService;

	@GetMapping("/goods")//遵循Restful规范的接口
	public List<GoodsDo> getList() {
		return goodsService.getGoodsList();
	}
}
```

### 3.4 启动应用

访问 `http://127.0.0.1:8080/goods` ，返回内容如下，可见后端接口已经可用。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3n0xh1j60je05swep02)

浏览器显示返回数据

## 4. 前端开发流程

前后端分离开发，实际上前端工作就简化了。我们直接新建项目文件夹 shop-front （商城前端项目文件夹），然后将前端页面放到该文件夹即可。

注意该页面不需要放到 Spring Boot 项目目录下，随便找个目录放置即可。实际开发过程中，后端和前端的项目可能都不在一台计算机上。

前端核心业务代码如下，由于前端技术不是本节介绍的重点，所以不再详细解释，感兴趣的同学可以从 [Git仓库](https://codechina.csdn.net/woshisangsang/spring-boot-wikis) 查看完整代码 。

**实例：**

```java
  //初始化方法
  $(function () {
    var row = "";
    $.ajax({
      type: "GET",
      url: "http://127.0.0.1:8080/goods", //后端接口地址
      dataType: "json",
      contentType: "application/json; charset=utf-8",
      success: function (res) {
        $.each(res, function (i, v) {
          row = "<tr>";
          row += "<td>" + v.id + "</td>";
          row += "<td>" + v.name + "</td>";
          row += "<td>" + v.price + "</td>";
          row += "<td>" + v.pic + "</td>";
          row += "</tr>";
          $("#goodsTable").append(row);
        });
      },
      error: function (err) {
        console.log(err);
      }
    });
  });
```

开发完该页面后，直接使用浏览器双击打开，查看控制台发现有错误信息提示。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3nc3ixj60j90akq3m02)

浏览器控制台返回错误信息

考验英文水平的时候到了！关键是 `has been blocked by CORS policy` ，意味着被 CORS 策略阻塞了。我们的前端页面请求被 CORS 阻塞了，所以没成功获取到后端接口返回的数据。

## 5. CORS 跨域介绍

跨域实际上源自浏览器的同源策略，所谓同源，指的是协议、域名、端口都相同的源（域）。浏览器会阻止一个域的 JavaScript 脚本向另一个不同的域发出的请求，这也是为了保护浏览器的安全。

在上面的例子中，发起请求的网页与请求资源的 URL 协议、域名、端口均不同，所以该请求就被浏览器阻止了。

CORS 的意思就是`跨域资源共享`，是一种允许跨域 HTTP 请求的机制，在这种情况下我们就要想办法实现 CORS 跨域了。

## 6. Spring Boot 跨域的实现

跨域的方法有很多种，我们此处演示一种常用的跨域方法。我们添加一个配置类，代码如下：

**实例：**

```java
@Configuration//配置类
public class CorsConfig {
	@Bean
	public WebMvcConfigurer corsConfigurer() {
		return new WebMvcConfigurer() {
			@Override
			public void addCorsMappings(CorsRegistry registry) {
				registry.addMapping("/**")//对所有请求路径
				          .allowedOrigins("*")//允许所有域名
				          .allowCredentials(true)//允许cookie等凭证
				          .allowedMethods("GET", "POST", "DELETE", "PUT","PATCH")//允许所有方法
				          .maxAge(3600);
			}
		};
	}
}
```

通过上面的配置类，实现了允许所有对该 Spring Boot 的请求跨域。

此时再次打开网页，被跨域策略阻塞的提示消失，界面显示如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3novhuj60jg05tq2x02)

浏览器正常显示商品信息

## 7. 小结

一般在项目开发过程中，会先根据功能需求规划页面和后端 API 接口， API 接口往往会形成详细的 API 文档说明。

后端会快速地开发 API 接口，前期并不会将 API 功能完全实现，而是仅仅返回一些测试值。就像本篇文章中后端接口返回的博客列表，并不是真实从数据库查询出来的，而是构造的测试数据。

同期，前端可以根据 API 文档，编写前端界面，并调用后端 API 进行测试。

也就是说，前后端实际上互相不必直接沟通，他们之间的交流是通过 API 文档完成的。由于后端接口往往采用 RESTful 标准规范，所以在理解上并不会存在很多问题。

越大型的项目，越需要规范，越需要职责分明。 Spring Boot 对后端服务化的支持实际上是相当到位的，使用几个简单的注解，就能轻松构建独立的后端服务项目。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
