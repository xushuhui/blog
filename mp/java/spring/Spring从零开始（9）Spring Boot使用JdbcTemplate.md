# Spring Boot 使用 JdbcTemplate

## 1. 前言

如果我们的项目非常简单，仅仅是对数据库几张表进行简单的增删改查操作，那么实际上直接使用 JDBC 操作数据库就可以了。

由于 JDBC 中有很多模板代码，每次都是加载驱动 - 建立数据库连接 - 查询或操作数据库 - 关闭数据库连接这样的模板代码， Spring 提供了 JdbcTemplate 对原生 JDBC 进行了简单的封装。

本篇文章，我们来实现一个完整的、基于 Spring Boot + JdbcTemplate + MySQL 的商品管理项目实例。

## 2. 技术选型

数据库使用 MySQL ，商品信息存储到商品表内即可。

后端项目使用 Spring Boot ，通过控制器暴露 RESTful 风格的接口供前端调用，通过 JdbcTemplate 实现对数据库的操作。

前端项目使用 Bootstrap 开发，通过 jQuery 提供的 $.ajax 方法访问后端接口。

## 3. 数据库模块实现

只需要一张商品表，保存商品相关的信息即可。我们使用 Navicat 新建数据库 shop ，并在其中新建数据表 goods 。

**实例：**

```java
CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一编号',
  `name` varchar(255) DEFAULT '' COMMENT '商品名称',
  `price` decimal(10,2) DEFAULT '0.00' COMMENT '商品价格',
  `pic` varchar(255) DEFAULT '' COMMENT '图片文件名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

> **Tips：** 默认值最好不要采用 NULL ， NULL 会影响索引的效率，而且在查询时需要用 is null 或 is not null 筛选，容易被忽略。

## 4. Spring Boot 后端实现

我们新建一个 Spring Boot 项目，通过 JdbcTemplate 访问数据库，同时接口依旧采用 RESTful 风格。

### 4.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 `com.imooc` ， Artifact 为 `spring-boot-jdbctemplate` ，生成项目后导入 Eclipse 开发环境。

### 4.2 引入项目依赖

我们引入 Web 项目依赖、热部署依赖。由于本项目需要访问数据库，所以引入 `spring-boot-starter-jdbc` 依赖和 `mysql-connector-java` 依赖。 pom.xml 文件中依赖项如下：

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
		<!-- jdbc -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-jdbc</artifactId>
		</dependency>
		<!-- myql驱动 -->
		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
		</dependency>
```

### 4.3 建立项目结构

依次新建以下类结构

* **GoodsDo**：商品类，对应 goods 商品表；
* **GoodsDao**：商品数据访问类，用于访问数据库；
* **GoodsService**：商品服务类，用于封装对商品的操作；
* **GoodsController**：商品控制器类，用于对外提供 HTTP 接口；
* **CorsConfig**：跨域配置类，允许前端页面跨域访问后端接口。

此时项目目录如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3pidobj60jh06lwf302)

### 4.4 开发商品类

开发商品类 GoodsDo ，代码如下：

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
	// 省略get set方法
```

### 4.5 开发商品数据访问类

商品数据访问类 GoodsDao 是本篇的重点，通过注入 JdbcTemplate 类型的组件，实现数据库操作。注入代码如下：

**实例：**

```java
/**
 * 商品数据库访问类
 */
@Repository // 标注数据访问类
public class GoodsDao {
	@Autowired
	private JdbcTemplate jdbcTemplate;
}
```

由于我们已经引入了 `spring-boot-starter-jdbc` 依赖，所以 Spring Boot 项目已经为我们自动配置了 JdbcTemplate 组件，我们拿来即用即可，这就是 Spring Boot 的强大之处！

此时我们启动应用，发现报错信息：

```java
***************************
APPLICATION FAILED TO START
***************************
Description:
Failed to configure a DataSource: 'url' attribute is not specified and no embedded datasource could be configured.
Reason: Failed to determine a suitable driver class
```

此处我们可以再度体会 Spring Boot 强大之处， Spring Boot 在为我们自动配置了 JdbcTemplate 之余，还在尝试自动为我们配置数据源 DataSource ，即 JdbcTemplate 要操作的真实数据库信息。报错信息已经提示我们，没有合适的数据库驱动、也没有合适的 URL 属性。

### 4.6 配置数据源信息

我们只需要通过配置文件指定数据源信息， Spring Boot 就可以识别配置，并加载到数据源组件中。 JdbcTemplate 也可以自动识别该数据源，从而实现对数据库的操作。配置文件信息如下：

**实例：**

```java
# 配置数据库驱动
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
# 配置数据库url
spring.datasource.url=jdbc:mysql://127.0.0.1:3306/shop?useUnicode=true&characterEncoding=utf-8&serverTimezone=GMT%2B8
# 配置数据库用户名
spring.datasource.username=root
# 配置数据库密码
spring.datasource.password=123456

```

需要注意的是，我们在 URL 配置中指定了编码方式，这样可以防止出现数据库中文乱码情况。同时指定了时区为北京时间所在的东八区（GMT%2B8），避免因时区问题导致错误。

此时再次启动 Spring Boot 应用，正常运行，说明我们的数据源配置生效了。

### 4.7 通过 JdbcTemplate 操作数据库

通过 JdbcTemplate 进行增删改查操作非常简洁， Spring 官方封装了原生 JDBC 中冗余的模板代码，使数据库访问操作更加简洁，代码如下：

**实例：**

```java
/**
 * 商品数据库访问类
 */
@Repository // 标注数据访问类
public class GoodsDao {
	@Autowired
	private JdbcTemplate jdbcTemplate;

	/**
	 * 新增
	 */
	public void insert(GoodsDo goods) {
		jdbcTemplate.update("insert into goods(name,price,pic)values(?,?,?)", goods.getName(), goods.getPrice(),
				goods.getPic());
	}

	/**
	 * 删除
	 */
	public void delete(Long id) {
		jdbcTemplate.update("delete from goods where id =?", id);
	}

	/**
	 * 更新
	 */
	public void update(GoodsDo goods) {
		jdbcTemplate.update("update goods set name=?,price=?,pic=? where id=?", goods.getName(), goods.getPrice(),
				goods.getPic(), goods.getId());
	}

	/**
	 * 按id查询
	 */
	public GoodsDo getById(Long id) {
		return jdbcTemplate.queryForObject("select * from goods where id=?", new RowMapper<GoodsDo>() {
			@Override
			public GoodsDo mapRow(ResultSet rs, int rowNum) throws SQLException {
				GoodsDo goods = new GoodsDo();
				goods.setId(rs.getLong("id"));
				goods.setName(rs.getString("name"));
				goods.setPrice(rs.getString("price"));
				goods.setPic(rs.getString("pic"));
				return goods;
			}
		}, id);
	}

	/**
	 * 查询商品列表
	 */
	public List<GoodsDo> getList() {
		return jdbcTemplate.query("select * from goods", new RowMapper<GoodsDo>() {
			@Override
			public GoodsDo mapRow(ResultSet rs, int rowNum) throws SQLException {
				GoodsDo goods = new GoodsDo();
				goods.setId(rs.getLong("id"));
				goods.setName(rs.getString("name"));
				goods.setPrice(rs.getString("price"));
				goods.setPic(rs.getString("pic"));
				return goods;
			}
		});
	}
}
```

getById 和 getList 方法中使用了匿名内部类，如果不了解的可以先去学习下相关知识。

### 4.8 开发商品服务类

商品服务类比较简单，直接调用 GoodsDao 完成商品服务方法封装即可。

**实例：**

```java
/**
 * 商品服务类
 */
 @Service
public class GoodsService {
	@Autowired
	private GoodsDao goodsDao;

	/**
	 * 新增商品
	 */
	public void add(GoodsDo goods) {
		goodsDao.insert(goods);
	}

	/**
	 * 删除商品
	 */
	 public void remove(Long id) {
        goodsDao.delete(id);
     }

	/**
	 * 编辑商品信息
	 */
	public void edit(GoodsDo goods) {
		goodsDao.update(goods);
	}

	/**
	 * 按id获取商品信息
	 */
	public GoodsDo getById(Long id) {
		return goodsDao.getById(id);
	}

	/**
	 * 获取商品信息列表
	 */
	public List<GoodsDo> getList() {
		return goodsDao.getList();
	}
}
```

### 4.9 开发商品控制器类

我们还是遵循之前的 RESTful 风格，制定后端访问接口如下：

|动词|接口含义|接口地址|
|----|--------|--------|
|GET   |查询商品(id=1)信息|[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|
|GET   |查询商品列表信息  |[http://127.0.0.1:8080/goods](http://127.0.0.1:8080/goods)    |
|POST  |新增商品          |[http://127.0.0.1:8080/goods](http://127.0.0.1:8080/goods)    |
|PUT   |修改商品(id=1)信息|[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|
|DELETE|删除商品(id=1)    |[http://127.0.0.1:8080/goods/1](http://127.0.0.1:8080/goods/1)|

我们根据上面的接口列表，实现控制器类代码如下：

**实例：**

```java
/**
 * 商品控制器类
 */
@RestController
public class GoodsController {
	@Autowired
	private GoodsService goodsService;
	/**
	 * 按id获取商品信息
	 */
	@GetMapping("/goods/{id}")
	public GoodsDo getOne(@PathVariable("id") long id) {
		return goodsService.getById(id);
	}
	/**
	 * 获取商品列表
	 */
	@GetMapping("/goods")
	public List<GoodsDo> getList() {
		return goodsService.getList();
	}
	/**
	 * 新增商品
	 */
	@PostMapping("/goods")
	public void add(@RequestBody GoodsDo goods) {
		goodsService.add(goods);
	}
	/**
	 * 编辑商品
	 */
	@PutMapping("/goods/{id}")
	public void update(@PathVariable("id") long id, @RequestBody GoodsDo goods) {
		// 修改指定id的博客信息
		goods.setId(id);
		goodsService.edit(goods);
	}
	/**
	 * 移除商品
	 */
	@DeleteMapping("/goods/{id}")
	public void delete(@PathVariable("id") long id) {
		goodsService.remove(id);
	}
}
```

### 4.10 开发跨域配置类

由于我们是前后端分离的项目开发方式，所以需要为 Spring Boot 添加跨域配置类：

**实例：**

```java
/**
 * 跨域配置类
 */
@Configuration
public class CorsConfig {
	@Bean
	public WebMvcConfigurer corsConfigurer() {
		return new WebMvcConfigurer() {
			@Override
			public void addCorsMappings(CorsRegistry registry) {
				registry.addMapping("/**")// 对所有请求路径
						.allowedOrigins("*")// 允许所有域名
						.allowCredentials(true)// 允许cookie等凭证
						.allowedMethods("GET", "POST", "DELETE", "PUT", "PATCH")// 允许所有方法
						.maxAge(3600);
			}
		};
	}
}
```

## 5. 前端页面开发

本节主要介绍 Spring Boot 中 JdbcTemplate 的用法，所以前端页面仅给出代码和注释，不再进行详细介绍了。

前端只有一个页面，使用 Bootstrap 的样式和插件，通过 jQuery 的 $.ajax 方法访问后端接口，逻辑并不复杂。

此处简单展示下浏览商品部分的前端代码，感兴趣的同学可以从 [Git仓库](https://codechina.csdn.net/woshisangsang/spring-boot-wikis) 查看完整代码。

**实例：**

```java
    //浏览商品
    function viewGoods() {
      var row = "";
      //先清空表格
      $('#GoodsTable').find("tr:gt(0)").remove();
      $.ajax({
        type: "GET",
        url: "http://127.0.0.1:8080/goods",
        dataType: "json",
        contentType: "application/json; charset=utf-8",
        success: function (res) {
          console.log(res);
          $.each(res, function (i, v) {
            row = "<tr>";
            row += "<td>" + v.id + "</td>";
            row += "<td>" + v.name + "</td>";
            row += "<td>" + v.price + "</td>";
            row += "<td>" + v.pic + "</td>";
            row +=
              "<td><a class='btn btn-primary btn-sm' href='javascript:editGoods(" + v.id + ")' >编辑</a>";
            row +=
              "<a class='btn btn-danger btn-sm' href='javascript:removeGoods(" + v.id + ")' >删除</a></td>";
            row += "</tr>";
            console.log(row);
            $("#GoodsTable").append(row);
          });
        },
        error: function (err) {
          console.log(err);
        }
      });
    }

```

## 6. 项目效果

直接使用浏览器打开前端页面，效果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3pqh8bj60lh06waah02)

浏览商品

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3q06f1j60lf0bg3z002)

新增商品

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3qba5lj60lh0de3z402)

编辑商品

## 7. 小结

本篇重点演示了 Spring Boot 中使用 JdbcTemplate 的方法。

基于 Spring Boot 自动装配的功能，我们只需要引入相应的依赖，编写必要的数据库参数配置，即可直接使用 JdbcTemplate 。所谓开箱即用，就是只需必要的操作，就可以直接到达可用的境界。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
