# Spring Boot 缓存应用场景

## 1. 前言

缓存是性能提升的大杀器！

要知道，内存的读写速度是硬盘的几十倍到上百倍。缓存实际上就是利用内存的高速读写特性，提高热点数据的操作速度。

Spring Boot 中使用缓存非常简单，并且支持多种缓存实现。

本篇介绍比较常用的几种缓存实现方式，及其对应的应用场景。

## 2. Spring Boot 默认缓存

Spring Boot 默认缓存是基于 ConcurrenMapCacheManager 缓存管理器实现的，从这个类名就能发现它本质上应该是一个 Map 集合容器。

ConcurrenMapCacheManager 结构比较简单，一般用于比较轻量级的缓存使用场景。也就是缓存的数据量比较小，缓存操作不是特别频繁的场景。

接下来就具体演示下， Spring Boot 默认缓存实现过程。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-cache ，生成项目后导入 Eclipse 开发环境。

### 2.2 引入项目依赖

引入 Web 项目依赖和缓存依赖。

**实例：**

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- 缓存依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-cache</artifactId>
		</dependency>
```

### 2.3 开启缓存

在启动类上添加注解 @EnableCaching 开启缓存功能。

**实例：**

```java

@SpringBootApplication
@EnableCaching // 开启缓存
public class SpringBootCacheApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootCacheApplication.class, args);
	}
}
```

### 2.4 定义服务层方法

正常服务层方法会调用数据访问层方法访问数据库，此处我们只需要演示缓存的作用，所以打印日志代替数据库访问方法。

**实例：**

```java
/**
 * 商品服务类
 */
@Service
@CacheConfig(cacheNames = "GoodsCache")
public class GoodsService {

	private Logger logger = LoggerFactory.getLogger(this.getClass());

	/**
	 * 按id获取商品信息
	 */
	@Cacheable
	public GoodsDo getById(Long id) {
		logger.info("getById({})", id);
		GoodsDo goods = new GoodsDo();
		goods.setId(id);
		goods.setName("goods-" + id);
		return goods;
	}

	/**
	 * 删除商品
	 */
	@CacheEvict(key = "#id")
	public void remove(Long id) {
		logger.info("remove({})", id);
	}

	/**
	 * 编辑商品信息
	 */
	@CachePut(key = "#goods.id")
	public GoodsDo edit(GoodsDo goods) {
		logger.info("edit id:{}", goods.getId());
		return goods;
	}
}
```

对于使用缓存的 GoodsService 服务类，我们需要具体解释下：

1. `@CacheConfig` 注解用于指定本类中方法使用的缓存名称，该类使用的缓存名称为 GoodsCache ，与其他缓存区域是隔离的。
2. `@Cacheable` 用于开启方法缓存，缓存的键是方法的参数，缓存的值是方法的返回值。如果多次调用该方法时参数 id 值相同，则第一次会执行方法体，并将返回值放入缓存；后续方法不会再执行方法体，直接将缓存的值返回。
3. `@CachePut` 可以更新缓存，`key = "#id"` 表示采用参数中的 id 属性作为键。当缓存中该键的值不存在时，则将返回值放入缓存；当缓存中该键的值已存在时，会更新缓存的内容。
4. `@CacheEvict` 可以移除缓存，当调用该方法时，会移除 goods 中 id 属性对应的缓存内容。

### 2.5 测试

为了充分理解缓存的含义，我们通过测试类发起测试。

**实例：**

```java
@SpringBootTest
class SpringBootCacheApplicationTests {
	private Logger logger = LoggerFactory.getLogger(this.getClass());
	@Autowired
	private CacheManager cacheManager;

	@Autowired
	private GoodsService goodsService;

	// 显示当前使用的缓存管理器类型
	@Test
	void showCacheManager() {
		// 输出：org.springframework.cache.concurrent.ConcurrentMapCacheManager
		logger.info(cacheManager.getClass().toString());
	}

	// 缓存测试
	@Test
	void cacheTest() {
		// 第一次执行，没有缓存，执行方法体
		goodsService.getById(1L);
		// 再次执行，直接取出缓存，不执行方法体
		goodsService.getById(1L);
		// 移除缓存
		goodsService.remove(1L);
		// 再次执行，已经没有对应缓存，所以执行方法体
		GoodsDo oldGoods = goodsService.getById(1L);
		// 打印缓存内容
		logger.info("old goods id:{} name:{}", oldGoods.getId(), oldGoods.getName());
		// 更新缓存
		GoodsDo temp = new GoodsDo();
		temp.setId(1L);
		temp.setName("新的商品");
		goodsService.edit(temp);
		// 查询并打印已更新的缓存内容
		GoodsDo newGoods = goodsService.getById(1L);
		logger.info("new goods id:{} name:{}", newGoods.getId(), newGoods.getName());
	}
}
```

我们查看下控制台输出如下，验证了我们设计的缓存机制。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ikfl4j60rw03rn1402)

使用 Spring Boot 默认缓存时控制台输出内容

## 3. 使用 Ehcache 缓存

Spring Boot 默认的缓存实现比较简单，功能也十分有限。如果是企业级的中大型应用，需要寻求更加稳定、可靠的缓存框架。

Ehcache 是 Java 编程领域非常著名的缓存框架，具备两级缓存数据——内存和磁盘，因此不必担心内存容量问题。另外 Ehcache 缓存的数据会在 JVM 重启时自动加载，不必担心断电丢失缓存的问题。

总之 Ehcache 的功能完整性和运行稳定性远远强于 Spring Boot 默认的缓存实现方式，而且 Spring Boot 使用 Ehcache 非常便捷，接下来我们就来实现下。

### 3.1 添加 Ehcache 依赖

我们在 spring-boot-cache 项目的基础上添加 Ehcache 依赖。

**实例：**

```java
		<!-- Ehcache 依赖 -->
		<dependency>
			<groupId>org.ehcache</groupId>
			<artifactId>ehcache</artifactId>
		</dependency>
		<!-- cache-api 依赖 -->
		<dependency>
			<groupId>javax.cache</groupId>
			<artifactId>cache-api</artifactId>
		</dependency>
```

### 3.2 添加 Ehcache 配置文件

首先在 application.properties 中指定配置文件的位置。

**实例：**

```java
spring.cache.jcache.config=classpath:ehcache.xml
spring.cache.type=jcache
```

然后在 resource 文件夹中添加 ehcache.xml 配置文件，内容如下：

**实例：**

```java
<?xml version="1.0" encoding="UTF-8"?>
<config xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
	xmlns='http://www.ehcache.org/v3'
	xsi:schemaLocation="http://www.ehcache.org/v3 http://www.ehcache.org/schema/ehcache-core.xsd">
	<!-- 持久化路径 -->
	<persistence directory="C://ehcache" />
	<!--缓存模板 -->
	<cache-template name="CacheTemplate">
		<expiry>
			<!--存活时间 -->
			<tti>60</tti>
		</expiry>
		<resources>
			<!--堆空间 -->
			<heap unit="entries">2000</heap>
			<!-- 堆外空间 -->
			<offheap unit="MB">500</offheap>
		</resources>
	</cache-template>

	<!--缓存对象 -->
	<cache alias="GoodsCache" uses-template="CacheTemplate">
	</cache>
</config>
```

> **Tips**：Ehcache 的配置比较复杂，此处只是给出简单的示例，感兴趣的同学可以查阅更多资料。

### 3.3 测试

由于之前已经在启动类添加 @EnableCaching ，我们再次运行测试类，输出结果如下。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2j00m3j60p9046n1w02)

使用 Ehcache 时控制台输出内容

注意控制台出现了 EhcacheManager 的字样，说明我们此时使用的缓存是 Ehcache 。

## 4. 使用 Redis 缓存

Ehcache 依然是 Java 进程内的缓存框架，受限于 JVM 整体的内存分配策略。

如果是大型系统，缓存的数据量特别大，且性能要求很高，可以考虑直接使用 Redis 作为缓存。

Redis 可以采用单机、主备、集群等模式，视乎具体项目需求决定即可。目前各大云计算厂商均提供商用版的 Redis 缓存服务，性能卓越且接入简单快速。

本节简单地演示 Spring Boot 中使用 Redis 单机缓存的方法，真实生产环境中建议至少使用主备类型的 Redis 实例。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2je0w8j60jt075dhj02)

华为云提供的缓存服务

### 4.1 修改缓存依赖

因为需要使用 Redis 缓存，所以将引入的依赖项修改如下：

**实例：**

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- 缓存依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-cache</artifactId>
		</dependency>
		<!-- Redis 相关依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-data-redis</artifactId>
		</dependency>
		<dependency>
			<groupId>io.lettuce</groupId>
			<artifactId>lettuce-core</artifactId>
		</dependency>
		<dependency>
			<groupId>redis.clients</groupId>
			<artifactId>jedis</artifactId>
		</dependency>
```

### 4.2 修改缓存配置

修改 application.properties 配置文件，将 Redis 配置及缓存配置设置如下：

**实例：**

```java
# 过期时间
spring.cache.redis.time-to-live=6000s

# Redis库的编号
spring.redis.database=0
# Redis实例地址
spring.redis.host=127.0.0.1
# Redis实例端口号，默认6379
spring.redis.port=6379
# Redis登录密码
spring.redis.password=Easy@0122
# Redis连接池最大连接数
spring.redis.jedis.pool.max-active=10
# Redis连接池最大空闲连接数
spring.redis.jedis.pool.max-idle=10
# Redis连接池最小空闲连接数
spring.redis.jedis.pool.min-idle=0
```

### 4.3 测试

由于之前已经通过注解 @EnableCaching 开启了缓存功能，此时我们直接运行测试类进行测试，输出结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2jorc1j60n903iadl02)

使用 Redis 缓存时控制台输出内容

从上图输出结果可以看出，已经成功使用了 Redis 缓存管理器。

另外我们可以直接使用 Redis 客户端查看生成的缓存信息，如下图已经有名为 `GoodsCache::1` 的缓存键存在了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2jxxivj60qm03pjsu02)

Redis 客户端查看缓存信息

## 5. 小结

Spring Boot 支持多种缓存实现方式，可以根据项目需求灵活选择。

* 缓存数据量较小的项目，可以使用 Spring Boot 默认缓存。
* 缓存数据量较大的项目，可以考虑使用 Ehcache 缓存框架。
* 如果是大型系统，对缓存的依赖性比较高，还是建议采用独立的缓存组件 Redis ，通过主备、集群等形式提高缓存服务的性能和稳定性。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
