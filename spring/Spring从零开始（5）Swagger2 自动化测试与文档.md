# Swagger2 自动化测试与文档

## 1. 前言

使用 Spring Boot 后，开发人员心里美美的，再也不需要写一大堆的配置文件了。

每天都能早早地下班，回家可以多打两把王者荣耀啦。

但是每次开发完后端接口，使用 Postman 测试比较麻烦。差不多的接口地址，差不多的参数，每次测试都要输入一遍，挺烦心。

另外前端那些家伙，完全不懂后端技术，天天要文档。就这么简简单单几个接口，还得给前端写。

咦，能不能自动生成接口文档，然后自动生成测试界面呢。

百度一搜 "Spring Boot 接口自动化测试"，发现很多文章推荐使用 Swagger2 。哈哈，站在巨人的肩膀上果然好办事。

## 2. Swagger2 功能

Swagger2 可以识别控制器中的方法，然后自动生成可视化的测试界面。

后端开发人员编写完 Spring Boot 后端接口后，直接可视化测试就行了。无需借助 Postman 等工具，也无需编写测试类和测试方法，更无需联系前端开发确认接口是否正常。

如果给控制器方法添加注解，还能自动生成在线 API 文档，简直太省心了。

## 3. Spring Boot 中使用 Swagger2 流程

还是先思考整体流程，再动手落地。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3ke13fj60ji0b8tan02)

整体流程

我们继续使用上一篇文章中的 `spring-boot-restful` 项目，为其添加 Swagger2 相关功能。

### 3.1 引入 Swagger2 相关依赖

修改 pom.xml 文件，引入 Swagger2 相关依赖。

**实例：**

```java
	<!-- 添加swagger2相关功能 -->
   	<dependency>
   		<groupId>io.springfox</groupId>
   		<artifactId>springfox-swagger2</artifactId>
   		<version>2.9.2</version>
   	</dependency>
   	<!-- 添加swagger-ui相关功能 -->
   	<dependency>
   		<groupId>io.springfox</groupId>
   		<artifactId>springfox-swagger-ui</artifactId>
   		<version>2.9.2</version>
   	</dependency>
```

### 3.2 启用并配置 Swagger2 功能

我们添加一个配置类，专门用于配置 Swagger2 相关功能，这样比较清晰点。通过 `@EnableSwagger2` 注解开启 Swagger2 功能，通过 `@Bean` 标注的方法将对 Swagger2 功能的设置放入容器。

**实例：**

```java
@Configuration // 告诉Spring容器，这个类是一个配置类
@EnableSwagger2 // 启用Swagger2功能
public class Swagger2Config {
	/**
	 * 配置Swagger2相关的bean
	 */
	@Bean
	public Docket createRestApi() {
		return new Docket(DocumentationType.SWAGGER_2)
				.apiInfo(apiInfo())
				.select()
				.apis(RequestHandlerSelectors.basePackage("com"))// com包下所有API都交给Swagger2管理
				.paths(PathSelectors.any()).build();
	}

	/**
	 * 此处主要是API文档页面显示信息
	 */
	private ApiInfo apiInfo() {
		return new ApiInfoBuilder()
				.title("演示项目API") // 标题
				.description("学习Swagger2的演示项目") // 描述
				.termsOfServiceUrl("http://www.imooc.com") // 服务网址，一般写公司地址
				.version("1.0") // 版本
				.build();
	}
}
```

### 3.3 使用 Swagger2 进行接口测试

此时我们启动项目，然后访问 `http://127.0.0.1:8080/swagger-ui.html` ，即可打开自动生成的可视化测试页面，如下图。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3ksvf1j60ji0lk76n02)

Swagger2 自动生成可视化测试界面

嗯，感觉这个页面简单整洁，直接给测试人员使用都很方便。我们以 update 方法为例演示下如何测试。先看看该方法的代码：

**实例：**

```java
	/**
	 * 修改商品
	 */
	@PutMapping("/goods/{id}")
	public void update(@PathVariable("id") long id, @RequestBody GoodsDo goods) {
		// 修改指定id的商品信息
		goods.setId(id);
		goodsService.editGoods(goods);
	}
```

测试时先选中对应的方法 update , 然后点击 Try it out 开始测试。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3l0z72j60jb02laa002)

Swagger2 生成的测试方法

在参数区域输入 id 和 goods 的值。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3l98vsj60jd096dg502)

Swagger2 可视化测试参数输入

点击 Execute 后，返回 Code 为 200 表示 http 请求成功！

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3lhqb6j60j207qdgc02)

Swagger2 可视化测试结果输出

由此可见， Swagger2 将接口以可视化的方式呈现出来，开发人员不必手输接口地址、参数名称，就可以发起测试并查看结果，确实非常方便。

后端人员在开发完成后，可以自己使用 Swagger2 测试下接口可行性。而前端人员也可以打开 Swagger2 网页直接验证接口功能。

### 3.4 使用 Swagger2 生成在线 API 文档

使用 Swagger2 生成在线文档比较简单，直接在控制器方法上添加注解即可。如下：

**实例：**

```java
@Api(tags = "商品API") // 类文档显示内容
@RestController
public class GoodsController {
	@Autowired
	private GoodsService goodsService;
	@ApiOperation(value = "根据id获取商品信息") // 接口文档显示内容
	@GetMapping("/goods/{id}")
	public GoodsDo getOne(@PathVariable("id") long id) {
		return goodsService.getGoodsById(id);
	}
	@ApiOperation(value = "获取商品列表") // 接口文档显示内容
	@GetMapping("/goods")
	public List<GoodsDo> getList() {
		return goodsService.getGoodsList();
	}
	@ApiOperation(value = "新增商品") // 接口文档显示内容
	@PostMapping("/goods")
	public void add(@RequestBody GoodsDo goods) {
		goodsService.addGoods(goods);
	}
	@ApiOperation(value = "根据id修改商品信息") // 接口文档显示内容
	@PutMapping("/goods/{id}")
	public void update(@PathVariable("id") long id, @RequestBody GoodsDo goods) {
		goods.setId(id);
		goodsService.editGoods(goods);
	}
	@ApiOperation(value = "根据id删除商品") // 接口文档显示内容
	@DeleteMapping("/goods/{id}")
	public void delete(@PathVariable("id") long id) {
		goodsService.removeGoods(id);
	}
}
```

此时再次打开 `http://127.0.0.1:8080/swagger-ui.htm` ，会发现相关接口都已经有文字描述了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3lqbaaj60ji09ymxx02)

Swagger2 生成在线 API 文档

## 4. 视频演示

## 5. 小结

一般公司在开发时，可以使用 Swagger2 快速实现测试，并生成在线文档。

由于开发阶段经常会修改接口，所以编写纸质文档实在是劳民伤财。如果使用 Swagger2 ，重启 Spring Boot 后刷新页面，就能看到最新 API 文档。

还有一个小技巧，我们可以开启热部署。当我们的代码发生变化时， Spring Boot 会自动重启，这样就省得每次都要手工重启 Spring Boot 。

**实例：**

```java
	    <!-- 引入该依赖即可开启热部署 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-devtools</artifactId>
		</dependency>
```

通过引入依赖，简单配置，就能实现一个非常棒的功能，这就是 Spring Boot 开箱即用的优点体现。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
