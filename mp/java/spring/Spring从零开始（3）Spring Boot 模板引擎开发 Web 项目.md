# Spring Boot 使用模板引擎开发 Web 项目

## 1. 前言

模板引擎这个词，咋听起来，有点高大上的意味。

实际上，模板引擎是非常平易近人的技术。譬如大家可能都比较熟悉的 JSP ，就是一种比较典型的模板引擎。

当浏览器将请求抛给控制器，控制器处理好数据后，就跳转 JSP 等模板引擎页面。注意在跳转的同时，还会将数据组装好，也交给模板引擎处理。

模板引擎会根据数据，和模板引擎的规则，动态生成 HTML 页面，最后返回给浏览器显示。

## 2. 模板引擎使用场景

我们使用 Spring Boot 开发 Web 项目，大体上有两种方式。

第一种方式，是后端服务化的方式，也是当前的主流方式。前端是静态的 HTML 页面，通过 Ajax 请求 Spring Boot 的后端接口。 Spring Boot 返回数据一般采用 JSON 格式，前端接收后将数据显示。

第二种方式，是采取模板引擎的方式。前端的请求，到达 Spring Boot 的控制器后，控制器处理请求，然后将返回数据交给模板引擎。模板引擎负责根据数据生成 HTML 页面，最后将 HTML 返回给浏览器。

我个人比较推荐第一种方式，说一下该方式的几个优点：

* **便于分工协作**：后端可以按自己的进度开发接口，前端可以开发页面，需要的时候直接调用后端 API ；
* **便于项目拓展**：比如前期是做的网站，后续要加一个 APP ，后端接口可以直接复用；
* **降低服务端压力**：后端只提供数据，一部分业务逻辑在前端处理了。服务端要做的事情少了，自然压力就小。

本篇是讲模板引擎，也得说说模板引擎的优点，王婆卖瓜不能光夸草莓啊。模板引擎开发的页面，对搜索引擎 SEO 比较友好；还有就是简单的页面，如果用模板引擎开发速度比较快，毕竟模板化的方法，目的就是减少重复提高效率。

## 3. Spring Boot 中常用的模板引擎

Spring Boot 支持的模板引擎种类很多，常见的有 FreeMarker 、 Thymeleaf 、 JSP 。

因为这些模板引擎使用的用户都不少，所以我们逐一介绍下其实现过程。

至于孰优孰劣，请各位看官自行评价。正所谓：尺有所短，寸有所长，各取所爱，万物生长！

## 4. 整体流程说明

本篇我们开发一个商品浏览项目实例。

此处说一个我个人的经验：在做一个项目或一个模块的时候，不要一开始就动手写代码，最好是谋定而后动。

我们作为程序员，实际上是整个程序世界的总指挥。应该先整体规划，再实现局部。这种总分型的开发方法便于我们理顺思路，提高编码效率！

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo38m20sj60ji0azac602)

编码如行军，先了解地形，设定整体作战计划，再派兵执行任务（图片来源于网络，版权归原作者所有）

好的，我们来思考下，实现商品浏览项目实例的整体流程：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo38zf55j60ji0irn0t02)

整体流程

可以看到，我们是先建立了控制器方法和页面，再去实现其中的具体细节。这样可以让我们的思维保持连贯性和整体性，在做一些页面和方法较多的项目时，会感觉更加顺畅。

## 5. 使用 FreeMarker

我们按整体流程，使用 FreeMarker 模板引擎，来实现商品浏览功能。

### 5.1 创建 Spring Boot 项目并导入开发环境

使用 Spring Initializr 创建项目，Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-freemarker` ，生成项目后导入 Eclipse 开发环境。

### 5.2 在 pom.xml 中引入相关依赖

引入 Web 项目及 FreeMarker 模板相关的依赖项，代码如下：

**实例：**

```java
	 	<!-- 引入web项目相关依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- freemarker -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-freemarker</artifactId>
		</dependency>
```

### 5.3 创建控制器方法，指向商品页面

创建控制器类，由于是商品相关的控制器，所以命名为 GoodsController ，代码如下：

**实例：**

```java
/**
 * 商品控制器
 */
@Controller // 标注为控制器
public class GoodsController {
	/**
	 * 获取商品列表
	 */
	@RequestMapping("/goods") // 请求路径
	public String goods() {
		return "goods";// 跳转到goods.ftl页面
	}
}
```

我们具体解释下该类的作用。

* @Controller 注解标注在 GoodsController 类上，会为该类注册一个控制器组件，放入 Spring 容器中。该组件具备处理请求的能力，其中的方法可以响应 HTTP 请求；
* @RequestMapping ("/goods") 注解标注在方法 goods () 上，所以请求路径如果匹配 `/goods` ，则由该方法进行处理；
* 返回值是字符串 `"goods"` ，由于我们已经引入 FreeMarker 依赖，所以该返回值会跳转到 goods.ftl 页面。

> **Tips：** 注意需要在 application.properties 文件中设置模板文件的后缀，即： `spring.freemarker.suffix=.ftl` 。如果不添加该配置，直接 `return "goods.ftl";` 会报错。

### 5.4 创建商品页面

我们 `resource/templates` 目录下新建商品页面 `goods.ftl` ，先不必实现具体功能，代码如下：

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>商品列表</title>
</head>
<body>
商品列表
</body>
</html>
```

此时我们启动项目，然后访问 `http://127.0.0.1:8080/goods` ，即可显示对应页面内容。

### 5.5 在控制器方法中，调用服务方法获取商品信息，并将信息交给模板引擎处理

定义商品类 GoodsDo 用来描述商品信息，注意 Do 表示数据模型对象（Data Object），代码如下：

**实例：**

```java
/**
 * 商品数据对象
 */
public class GoodsDo {
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
  // 省略get set方法
}
```

然后我们编写服务类 GoodsService ，提供获取商品列表的方法。注意此处仅仅是演示模板引擎，并不需要访问数据库，直接返回一个指定内容的商品列表。

**实例：**

```java
/**
 * 商品服务
 */
@Service // 为GoodsService注册一个组件
public class GoodsService {
	public List<GoodsDo> getGoodsList() {
		List<GoodsDo> list = new ArrayList<GoodsDo>();
		GoodsDo goods = new GoodsDo();
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		list.add(goods);
		return list;
	}
}
```

此时，我们的控制器就可以注入 GoodsService 类型的组件，然后调用其方法了。

**实例：**

```java
@Controller
public class GoodsController {
	@Autowired
	private GoodsService goodsService;// 自动装配

	@RequestMapping("/goods") // 请求路径
	public String goods(Model model) {
		model.addAttribute("goodsList", goodsService.getGoodsList());// 交给模板引擎处理的数据
		return "goods";// 跳转到goods.ftl页面
	}
}
```

注意 `model.addAttribute("goodsList", goodsService.getGoodsList());` ，我们将商品列表相关的数据交给模板引擎去处理。

### 5.6 在商品页面通过模板引擎规则显示商品信息

此时我们可以根据 FreeMarker 模板引擎，按模板规则显示商品信息了。

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>商品列表</title>
</head>
<body>
<div>商品列表:</div>
<#list goodsList as item>
${item.name}--${item.price}--${item.pic}
</#list>
</body>
</html>
```

注意我们通过 FreeMarker 的模板语法，输出了商品列表信息。关于 FreeMarker 模板引擎更多的语法规则，感兴趣的同学可以后续查阅更多资料。

### 5.7 测试

启动项目，打开浏览器访问 `http://127.0.0.1:8080/goods` ，即可查看输出结果。

## 6. 使用 Thymeleaf

Thymeleaf 和 FreeMarker ，都是模板引擎，使用方法基本类似。此处我们仅仅是给出一个范例，不再做过多的解释。

### 6.1 创建 Spring Boot 项目并导入开发环境

使用 Spring Initializr 创建项目， Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-thymeleaf` ，生成项目后导入 Eclipse 开发环境。

### 6.2 在 pom.xml 中引入相关依赖

引入 Web 项目及 Thymeleaf 模板相关的依赖项。

**实例：**

```java
		<!-- 引入web项目相关依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- ThymeLeaf依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-thymeleaf</artifactId>
		</dependency>
```

### 6.3 创建控制器方法，指向商品页面

创建控制器类， GoodsController ， Thymeleaf 直接使用 HTML 作为模板页面，故代码如下：

**实例：**

```java
/**
 * 商品控制器
 */
@Controller // 标注为控制器
public class GoodsController {
	/**
	 * 获取商品列表
	 */
	@RequestMapping("/goods") // 请求路径
	public String goods() {
		return "goods.html";// 跳转到goods.html页面
	}
}
```

### 6.4 创建商品页面

我们在 `resource/templates` 目录下新建商品页面 `goods.html` ，先不必实现具体功能，代码如下：

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>商品列表</title>
</head>
<body>
商品列表
</body>
</html>
```

此时我们启动项目，然后访问 `http://127.0.0.1:8080/goods` ，即可显示对应页面内容。

### 6.5 在控制器方法中，调用服务方法获取商品信息，并将信息交给模板引擎处理

商品类 GoodsDo ，服务类 GoodsService ，这两个类与上面没有区别直接放出代码。

**实例：**

```java
/**
 * 商品数据对象
 */
public class GoodsDo {
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
  // 省略get set方法
}
```

**实例：**

```java
/**
 * 商品服务
 */
@Service // 为GoodsService注册一个组件
public class GoodsService {
	public List<GoodsDo> getGoodsList() {
		List<GoodsDo> list = new ArrayList<GoodsDo>();
		GoodsDo goods = new GoodsDo();
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		list.add(goods);
		return list;
	}
}
```

好的，此时我们的控制器就可以注入 GoodsService 类型的组件，然后调用其方法了。

**实例：**

```java
@Controller
public class GoodsController {
	@Autowired
	private GoodsService goodsService;// 自动装配

	@RequestMapping("/goods") // 请求路径
	public String goods(Model model) {
		model.addAttribute("goodsList", goodsService.getGoodsList());// 交给模板引擎处理的数据
		return "goods.html";// 跳转到goods.html页面
	}
}
```

### 6.6 在商品页面通过模板引擎规则显示商品信息

此时我们可以根据 Thymeleaf 模板引擎，按模板规则显示商品信息了。

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>商品列表</title>
</head>
<body>
	<div>商品列表:</div>
	<div th:each="item:${goodsList}">
		<span th:text="${item.name}"></span>
		<span th:text="${item.price}"></span>
		<span th:text="${item.pic}"></span>
	</div>
</body>
</html>
```

注意我们通过 Thymeleaf 的模板语法，输出了商品列表信息。关于 Thymeleaf 模板引擎更多的语法规则，感兴趣的同学可以后续查阅更多资料。

### 6.7 测试

启动项目，打开浏览器访问 `http://127.0.0.1:8080/goods` ，即可查看输出结果。

到此，大家基本上也能发现，这两种方式除了模板页面文件内容不同，其他地方基本都是一模一样的。

也就是说，模板引擎主要负责通过一些模板标签，将控制器返回的数据解析为网页。

## 7. 使用 JSP

注意 Spring Boot 官方已经不推荐使用 JSP 了，确实操作起来也比较麻烦。但是由于 JSP 用户体量还是比较大的，所以此处还是简单演示下，开发步骤与 FreeMarker / Thymeleaf 基本一致。

### 7.1 创建 Spring Boot 项目并导入开发环境

使用 Spring Initializr 创建项目， Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-jsp` ，生成项目后导入 Eclipse 开发环境。

### 7.2 在 pom.xml 中引入相关依赖

引入 Web 项目及 JSP 模板相关的依赖项。

**实例：**

```java
		<!-- 添加web开发功能 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!--内嵌的tomcat支持模块 -->
		<dependency>
			<groupId>org.apache.tomcat.embed</groupId>
			<artifactId>tomcat-embed-jasper</artifactId>
			<scope>provided</scope>
		</dependency>
		<!-- 对jstl的支持 -->
		<dependency>
			<groupId>javax.servlet</groupId>
			<artifactId>jstl</artifactId>
		</dependency>
```

### 7.3 创建控制器方法，指向商品页面

创建控制器类， GoodsController ，代码如下：

**实例：**

```java
/**
 * 商品控制器
 */
@Controller // 标注为控制器
public class GoodsController {
	/**
	 * 获取商品列表
	 */
	@RequestMapping("/goods") // 请求路径
	public String goods() {
		return "goods";// 跳转到goods.jsp页面
	}
}
```

### 7.4 创建商品页面

手工添加 `src/main/webapp` 及子目录如下，同时目录下放一个 goods.jsp 用于测试。注意该目录是一个 Source Folder 源代码目录，不是普通文件夹目录。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo397n9rj60mt07bq3d02)

spring-boot-jsp 项目结构

**实例：**

```java
<%@ page language="java" contentType="text/html; charset=UTF-8"
	pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>商品列表</title>
</head>
<body>商品列表
</body>
</html>
```

注意，我们还需要添加一个视图解析器，实现 JSP 页面往指定目录跳转。

**实例：**

```java
@SpringBootApplication
public class SpringBootJspApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootJspApplication.class, args);
	}
	@Bean // 注册视图解析器
	public InternalResourceViewResolver setupViewResolver() {
		InternalResourceViewResolver resolver = new InternalResourceViewResolver();
		resolver.setPrefix("/WEB-INF/jsp/");// 自动添加前缀
		resolver.setSuffix(".jsp");// 自动添加后缀
		return resolver;
	}
}
```

此时我们启动项目，然后访问 `http://127.0.0.1:8080/goods` ，即可显示对应页面内容。

### 7.5 在控制器方法中，调用服务方法获取商品信息，并将信息交给模板引擎处理

商品类 GoodsDo ，服务类 GoodsService ，这两个类与上面没有区别直接放出代码。

**实例：**

```java
/**
 * 商品数据对象
 */
public class GoodsDo {
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
  // 省略get set方法
}
```

**实例：**

```java
/**
 * 商品服务
 */
@Service // 为GoodsService注册一个组件
public class GoodsService {
	public List<GoodsDo> getGoodsList() {
		List<GoodsDo> list = new ArrayList<GoodsDo>();
		GoodsDo goods = new GoodsDo();
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		list.add(goods);
		return list;
	}
}
```

好的，此时我们的控制器就可以注入 GoodsService 类型的组件，然后调用其方法了。

**实例：**

```java
@Controller
public class GoodsController {
	@Autowired
	private GoodsService goodsService;// 自动装配

	@RequestMapping("/goods") // 请求路径
	public String goods(Model model) {
		model.addAttribute("goodsList", goodsService.getGoodsList());// 交给模板引擎处理的数据
		return "goods";// 跳转到goods.jsp
	}
}
```

### 7.6 在商品页面通过模板引擎规则显示商品信息

此时我们可以根据 JSP 模板引擎，按模板规则显示商品信息了。

**实例：**

```java
<%@ page language="java" contentType="text/html; charset=UTF-8"
	pageEncoding="UTF-8"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>商品列表</title>
</head>
<body>
	<div>商品列表:</div>
	<c:forEach var="item" items="${goodsList}">
         ${item.name}--${item.price}--${item.pic}
	</c:forEach>
</body>
</html>
```

注意我们通过 JSP 的模板语法，输出了商品列表信息。关于 JSP 模板引擎更多的语法规则，感兴趣的同学可以后续查阅更多资料。

### 7.7 测试

启动项目，打开浏览器访问 `http://127.0.0.1:8080/goods` ，即可查看输出结果。

## 8. 小结

最后大家应该也发现了， FreeMarker 和 Thymeleaf 的用法几乎是一模一样的，而 JSP 还需要手工添加一些目录和配置。

三种方式各有优劣， FreeMarker 模板语法比较简洁， Thymeleaf 可以直接使用 HTML 作为模板文件， JSP 用户群体广泛。

但是三种方式，都是一种模板引擎而已，将控制器返回数据转化为 HTML 页面显示，本质上没啥区别，大家对模板引擎有一个了解即可。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
