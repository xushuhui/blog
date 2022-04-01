# Spring Boot 开发 RESTful 风格 Web 项目

## 1. 前言

很多研发团队，可能都会有一个体会，当多人研发一个项目时，并不能达到 1+1>2 的效果。有时候还会出现 1+1<1 ，即 2 个人还不如 1 个人干得快，甚是悲哀。

就比如我们要开发一个 Web 项目，由前端工程师和后端工程师共同完成。

前端工程师懂 HTML / CSS / JavaScript 和 Bootstrap 等前端技术与框架，但是几乎不懂后端 Java 语言。

后端工程师懂 Spring Boot 开发，略懂 HTML / CSS / JavaScript ，但是没用过前端框架。

这种情况，如果使用 FreeMarker / Thymeleaf / JSP 等模板引擎，将面临不小的困难。前端工程师得先去学习模板引擎的规则，后端人员需要跟前端沟通控制器交给模板引擎处理的数据。

怎么办呢，不要担心，大佬们早就解决这个问题了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3h9hqqj60ji0c7q4m02)

互联网江湖圈，遇到问题先莫慌，早有大佬为我们开好路！（图片来源于网络，版权归原作者所有）

## 2. 前后端分离

前后端分离这种概念和技术，早就流行多年了。

具体点说，前端编写 HTML 页面，然后通过 Ajax 请求后端接口；后端把接口封装成 API ，返回 JSON 格式的数据；前端接收到 JSON 返回数据后渲染到页面。

前端工程师根本不需要懂后端，调用后端接口就行。后端使用 Spring Boot 控制器返回 JSON 十分简单，给方法添加个注解，就能将返回值序列化为 JSON 。

前端干前端的活，后端干后端的活，职责分明，界限明确。这就是前后端分离的好处啊！

## 3. RESTful 风格后端接口

前后端分离时，后端接口可不能太随意，目前后端接口编写大多遵循 RESTful 风格。

做后端接口的公司这么多，如果大家都定义自己的规范，是不利于公司之间的合作的。如果大家都能遵循一个规范来开发接口，很明显相关人员都能省心不少。

RESTful 就是一种非常流行的 HTTP 接口规范，简单明了，使用的人也多，用它准没错。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3hjt6qj60ji0azq4x02)

规范的意义，就是提供标准，提高效率。汽车行业标准化程度已经很高了，软件行业还需努力！（图片来源于网络，版权归原作者所有）

## 4. 整体流程说明

本节实现一个基于 RESTful 风格的 Spring Boot 商品浏览 API 实例。

做事之前，先定整体流程。凡事预则立，不预则废，老祖宗的智慧太厉害了，我们争取发扬光大。确定流程如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3hs4rgj60ji08yjsz02)

整体流程

## 5. 开发阶段

### 5.1 根据需求制定 RESTful 风格的接口文档

既然是要做商品浏览页面，将商品增删改查都实现了就是了。 RESTful 风格接口并不麻烦，一般情况下需要项目团队一起商量制定。此处我们指定如下：

|动词|接口含义|接口地址|
|----|--------|--------|
| GET  | 查询商品 (id=1) 信息 |[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|
|GET   | 查询商品列表信息     |[http://127.0.0.1:8080/goods](http://127.0.0.1:8080/goods)    |
|POST  | 新增商品             |[http://127.0.0.1:8080/goods](http://127.0.0.1:8080/goods)    |
|PUT   | 修改商品 (id=1) 信息 |[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|
|DELETE| 删除商品 (id=1)      |[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|

> **Tips：** RESTful 风格通过 HTTP 动词（ GET / POST / PUT / DELETE ）区分操作类型， URL 格式比较固定，仅此而已，非常简单。

### 5.2 按文档开发后端 API 接口

相比于使用模板引擎，用 Spring Boot 开发后端接口简直太轻松了。通过给控制器添加注解，就能将控制器方法返回值序列化为 JSON 。程序员最爱，就是轻松又愉快。

#### 5.2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-restful` ，生成项目后导入 Eclipse 开发环境。

像这种老套的重复操作，我还是拿出来讲讲，为何，因为要继续呵护我们的初学者们。万一有朋友因为一个小地方看不明白，丧失了学习编程的动力，那就是罪过了。

#### 5.2.2 引入项目依赖

RESTful 项目其实就是标准的 Web 项目，引入 Web 项目依赖即可。

**实例：**

```java
	  <!-- 引入web项目相关依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
```

#### 5.2.3 创建商品类与商品服务类

创建商品类与商品服务类，以完成对商品的增删改查操作。由于本章我们的重点是演示 RESTful 后端接口，所以此处没有操作真实的数据库。

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
  //省略get set方法
}
```

**实例：**

```java
/**
 * 商品服务
 */
@Service // 注册为服务类
public class GoodsService {
	/**
	 * 获取商品列表
	 */
	public List<GoodsDo> getGoodsList() {
		List<GoodsDo> goodsList = new ArrayList<GoodsDo>();
		GoodsDo goods = new GoodsDo();
		goods.setId(1L);
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		goodsList.add(goods);
		return goodsList;
	}
	/**
	 * 按id获取商品信息，模拟返回对应商品信息
	 */
	public GoodsDo getGoodsById(Long id) {
		GoodsDo goods = new GoodsDo();
		goods.setId(1L);
		goods.setName("苹果");
		goods.setPic("apple.jpg");
		goods.setPrice("3.5");
		return goods;
	}
	/**
	 * 新增商品，模拟返回数据库影响行数
	 */
	public int addGoods(GoodsDo goods) {
		return 1;
	}
	/**
	 * 根据商品id更新商品信息，模拟返回数据库影响行数
	 */
	public int editGoods(GoodsDo goods) {
		return 1;
	}
	/**
	 * 根据商品id删除对应商品，模拟返回数据库影响行数
	 */
	public int removeGoods(Long id) {
		return 1;
	}
}
```

> **Tips：** 服务层方法，建议不要使用 select /insert/update /delete 命名，因为服务层处理的逻辑往往不止于关系数据库表的增删改查。此处采用的是 get /add/edit /remove 。

#### 5.2.4 根据 API 文档实现控制器方法

此处需要解释的地方我都写在注释中了。

**实例：**

```java
@RestController // 通过该注解，第一是将GoodsController注册为控制器，可以响应Http请求；第二是可以将控制器中的方法返回值序列化为json格式。
public class GoodsController {
	@Autowired // 自动装配goodsService
	private GoodsService goodsService;
	/**
	 * 查询商品信息
	 * 1、@GetMapping表示可以使用get方法请求该api
	 * 2、"/goods/{id}"表示请求路径为/goods/{id}的形式，其中{id}为占位符
	 * 3、@PathVariable("id")表示将占位符{id}的值传递给id
	 * 4、也就是说/goods/123请求的话，会将123传递给参数id
	 */
	@GetMapping("/goods/{id}")
	public GoodsDo getOne(@PathVariable("id") long id) {
		return goodsService.getGoodsById(id);
	}
	/**
	 * 查询商品列表，使用get方法
	 */
	@GetMapping("/goods")
	public List<GoodsDo> getList() {
		return goodsService.getGoodsList();
	}
	/**
	 * 新增商品
	 * 1、@PostMapping表示使用post方法
	 * 2、@RequestBody表示将请求中的json信息转换为GoodsDo类型的对象信息，该转换也是由SpringMVC自动完成的
	 */
	@PostMapping("/goods")
	public void add(@RequestBody GoodsDo goods) {
		goodsService.addGoods(goods);
	}
	/**
	 * 修改商品
	 */
	@PutMapping("/goods/{id}")
	public void update(@PathVariable("id") long id, @RequestBody GoodsDo goods) {
		// 修改指定id的商品信息
		goods.setId(id);
		goodsService.editGoods(goods);
	}
	/**
	 * 删除商品
	 */
	@DeleteMapping("/goods/{id}")
	public void delete(@PathVariable("id") long id) {
		goodsService.removeGoods(id);
	}
}
```

### 5.3 使用 Postman 测试 API 接口可用

后端开发完 API 接口后，需要先进行下简单测试以保证接口是正确可用的。

我们可以使用 Postman 进行简单测试。启动 Spring Boot 项目后，使用 Postman 进行可视化测试，此处结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3i41u3j60rz0demyk02)

测试查询商品信息

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3ie4cej60s30e9abh02)

测试查询商品列表

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3iozazj60s10dwtaf02)

测试新增商品

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3iz9n7j60s30dsdhi02)

测试修改商品

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3jaowgj60s60b70tx02)

测试删除商品

## 6. 视频演示

## 7. 小结

使用 Spring Boot 开发 RESTful 风格的 Web 项目，是当前主流的 Web 项目开发模式。

这种开发模式的特点就是简单、明确。简单主要是依赖 Spring Boot 实现，明确就是靠 RESTful 风格来规范。

简单的目的就是提高开发效率；明确的目的是统一规范，从而降低沟通成本，最终也是为了提高开发效率。

Spring Boot 和 RESTful ，绝配！

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
