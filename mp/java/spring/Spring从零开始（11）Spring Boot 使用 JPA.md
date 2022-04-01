# Spring Boot 使用 JPA

## 1. 前言

使用 JDBC ，或者 JdbcTemplate 操作数据库，需要编写大量的 SQL 语句。SQL 语句基本都是些模板代码，实际上是可以通过分析模板代码的规则自动生成的。

JPA 就是简化 Java 持久层数据操作的技术标准，是一种方案和规范。最开始是 Sun 公司提出的， Sun 公司就是开发出 Java 的公司，一度非常厉害，结果被 Oracle 收购了。Sun 公司虽然提出了 JPA 标准，但是并没有具体实现。JPA 的实现里面比较出名的就是 Hibernate 了，所以本篇我们也是以 Hibernate 实现为基础进行 Spring Boot + JPA 的实例讲解。

本篇演示一个 Spring Boot 商品管理项目实例，其中数据持久层操作采用 JPA ，以体会 JPA 的优雅与高效。

## 2. JPA 基本原理

在开始实例之前，还是有必要聊聊 JPA 是如何实现的，便于大家理解。

首先是 ORM 映射，通过注解或 XML 描述对象和表直接的映射关系。例如 GoodsDo 商品类对应数据库中的 goods 商品表，商品类里面的属性和商品表里面的列一一对应，商品类的一个对象就对应商品表中的一行数据。

然后就是对数据库进行 CRUD （增删改查）操作了，由于已经配置了对象和表的映射关系，所以可以自动生成对应的 SQL 语句，然后执行语句即可。

## 3. 开发流程

光说不练那是假把式，我们来使用 Spring Boot + JPA 开发一个完整实例。

### 3.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 `com.imooc` ， Artifact 为 `spring-boot-jpa` ，生成项目后导入 Eclipse 开发环境。

### 3.2 引入项目依赖

我们引入 Web 项目依赖、热部署依赖。由于本项目需要使用 JPA 访问数据库，所以引入 `spring-boot-starter-jdbc` 、 `mysql-connector-java` 和 `spring-boot-starter-data-jpa` 依赖。 pom.xml 文件中依赖项如下：

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
		<!-- jpa -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-data-jpa</artifactId>
		</dependency>
		<!-- myql驱动 -->
		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
		</dependency>
```

### 3.3 修改配置文件

在 `application.properties` 中添加以下配置：

**实例：**

```java
# 配置数据库驱动
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
# 配置数据库url
spring.datasource.url=jdbc:mysql://127.0.0.1:3306/shop?useUnicode=true&characterEncoding=utf-8&serverTimezone=UTC
# 配置数据库用户名
spring.datasource.username=root
# 配置数据库密码
spring.datasource.password=Easy@0122

# 启动时更新表结构，保留数据
spring.jpa.hibernate.ddl-auto=update
```

此处需要注意的是 `spring.jpa.hibernate.ddl-auto=update` 。可以理解项目启动时，根据实体类结构更新数据库表结构，且保留数据库中的数据。

### 3.4 开发商品类

开发商品类 GoodsDo ，并通过注解实现类结构与数据表结构的映射。

**实例：**

```java

/**
 * 商品类
 */
@Entity // 表示这是一个数据对象类
@Table(name = "goods") // 对应数据库中的goods表
public class GoodsDo {
	/**
	 * 商品id
	 */
	@Id // 该字段对应数据库中的列为主键
	@GeneratedValue(strategy = GenerationType.IDENTITY) // 主键自增长
	@Column(name = "id") // 对应goods表中的id列
	private Long id;
	/**
	 * 商品名称
	 */
	@Column(name = "name") // 对应goods表中的name列
	private String name;
	/**
	 * 商品价格
	 */
	@Column(name = "price") // 对应goods表中的price列
	private String price;
	/**
	 * 商品图片
	 */
	@Column(name = "pic") // 对应goods表中的pic列
	private String pic;
	// 省略get set方法
}
```

### 3.5 开发数据操作接口

开发商品数据接口，代码如下：

**实例：**

```java
/**
 * 商品数据操作接口
 */
@Repository
public interface IGoodsDao extends CrudRepository<GoodsDo, Long> {
}
```

解释下，`@Repository` 将接口标注为数据访问层组件，该接口通过继承 `CrudRepository` 实现 CRUD 操作。泛型参数分别为实体类及主键的数据类型。

注意此时已经可以通过 IGoodsDao 对数据库 goods 表进行增删改查操作了。

### 3.6 开发服务层

开发 Goods Service ，注入 IGoodsDao 类型组件实现服务方法。

**实例：**

```java
/**
 * 商品服务类
 */
@Service
public class GoodsService {
	@Autowired
	private IGoodsDao goodsDao;

	/**
	 * 新增商品
	 */
	public void add(GoodsDo goods) {
		goodsDao.save(goods);
	}

	/**
	 * 删除商品
	 */
	public void remove(Long id) {
		goodsDao.deleteById(id);
	}

	/**
	 * 编辑商品信息
	 */
	public void edit(GoodsDo goods) {
		goodsDao.save(goods);
	}

	/**
	 * 按id获取商品信息
	 */
	public Optional<GoodsDo> getById(Long id) {
		return goodsDao.findById(id);
	}

	/**
	 * 获取商品信息列表
	 */
	public Iterable<GoodsDo> getList() {
		return goodsDao.findAll();
	}
}
```

此处需要解释下 `Optional` 类，它是一个包装类。它的内容是空或者包含的对象，所以可以避免空指针问题。此处稍作了解即可。

### 3.7 开发控制器

我们还是遵循 RESTful 风格，开发控制器类。

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
	public Optional<GoodsDo> getOne(@PathVariable("id") long id) {
		return goodsService.getById(id);
	}
	/**
	 * 获取商品列表
	 */
	@GetMapping("/goods")
	public Iterable<GoodsDo> getList() {
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

## 4. 测试

我们主要是测试 JPA 模块正确可用，所以直接在测试类发起对 IGoodsDao 方法的测试即可。

### 4.1 新增测试

首先我们建立数据库 shop ，数据库中不必有表 goods ，如果有 goods 表的话可以将它删除。因为我们设置了 `spring.jpa.hibernate.ddl-auto=update` ， JPA 会在项目启动时自动建立表结构。

**实例：**

```java
@RunWith(SpringRunner.class)
@SpringBootTest
public class JpaAddTest {
	@Autowired
	private IGoodsDao goodsDao;

	/**
	 * 新增测试
	 */
	@Test
	public void testAdd() {
		GoodsDo goods = new GoodsDo();
		goods.setName("梨张");
		goods.setPic("梨图片");
		goods.setPrice("2.0");
		GoodsDo result = goodsDao.save(goods);
		System.out.println("新增商品id：" + result.getId());
		assertNotNull(result);
	}
}
```

运行测试类，控制台输出`新增商品id：1`，说明插入一条数据成功，且插入数据 id 为 1 。

同时查看数据库，发现已经自动构建表结构：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo23s8kij60ji06rwex02)

MySQL 数据库已自动构建表结构

### 4.2 修改测试

当调用 save 方法，如果给参数中 id 属性赋值，则会进行数据更新操作。

**实例：**

```java
@RunWith(SpringRunner.class)
@SpringBootTest
public class JpaEditTest {
   @Autowired
   private IGoodsDao goodsDao;

   /**
    * 修改测试
    */
   @Test
   public void testEdit() {
   	GoodsDo goods = new GoodsDo();
   	goods.setId(1L);
   	goods.setName("梨张");
   	goods.setPic("梨图片");
   	goods.setPrice("100.0");
   	GoodsDo result = goodsDao.save(goods);
   	assertNotNull(result);
   }
}
```

此时查看数据库中数据，发现金额已修改成功。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2422drj60j403m0sx02)

MySQL 数据库中金额修改成功

### 4.3 查询测试

我们进行按 id 查询、查询所有操作，并打印查询结果。

**实例：**

```java
@RunWith(SpringRunner.class)
@SpringBootTest
public class JpaQueryTest {
   @Autowired
   private IGoodsDao goodsDao;

   /**
    * 按id查询
    */
   @Test
   public void testQueryById() {
   	Optional<GoodsDo> goodsOptional = goodsDao.findById(1L);
   	GoodsDo goods = goodsOptional.get();
   	System.out.println(goods.getId() + "-" + goods.getName() + "-" + goods.getPic() + "-" + goods.getPrice());
   }

   /**
    * 查询全部
    */
   @Test
   public void testQueryAll() {
   	Iterable<GoodsDo> goodsIt = goodsDao.findAll();
   	for (GoodsDo goods : goodsIt) {
   		System.out.println(goods.getId() + "-" + goods.getName() + "-" + goods.getPic() + "-" + goods.getPrice());
   	}
   }
}
```

### 4.4 删除测试

指定删除 id 为 1 的商品。

**实例：**

```java
@RunWith(SpringRunner.class)
@SpringBootTest
public class JpaRemoveTest {
   @Autowired
   private IGoodsDao goodsDao;

   /**
    * 删除测试
    */
   @Test
   public void testRemove() {
   	goodsDao.deleteById(1L);
   }
}
```

运行后，数据库中商品信息被删除，大功告成！

## 5. 小结

使用 JPA 后，最大的好处就是不用写 SQL 了，完全面向对象编程，简洁又省心，何乐而不为。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
