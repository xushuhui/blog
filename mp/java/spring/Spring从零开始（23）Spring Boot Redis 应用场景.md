# Spring Boot Redis 应用场景

## 1. 前言

Redis 其实就是基于内存的键值型数据库，与 Oracle 、 SQL Server 、 MySQL 等传统关系型数据库相比，它最大的优势就是读写速度快。

到底有多快呢，我曾经使用 Windows 版本的 Redis 进行过真实测试，每秒读写次数均可以超过 1 万次。据了解 Redis 每秒的读写操作次数其实是可以达到 10 万多次的。

所以 Redis 非常适合作为热点数据的缓存，这个我们在上一节已经演示过了。本节通过其他两个实际场景来演示下 Spring Boot 中如何应用 Redis 。

* 网站的访问次数
* 热门商品排行榜

## 2. 网站的访问次数

大型网站访问次数的查询、更新非常频繁，如果通过关系数据库读写，无疑会耗费大量的性能，而使用 Redis 可以大幅提高速度并降低对关系数据库的消耗。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-redis，生成项目后导入 Eclipse 开发环境。

### 2.2 引入项目依赖

我们引入 Web 项目依赖与 Redis 依赖。

**实例：**

```java
		<!-- Web 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- Redis 依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-data-redis</artifactId>
		</dependency>
```

### 2.3 配置 Redis 数据库连接

修改 application.properties 配置文件内容如下。

**实例：**

```java
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

### 2.4 开发网站访问统计服务类

开发网站访问统计服务类，在第 1 次获取访问次数时初始化次数为 0 ，后续每次访问次数加 1 。

**实例：**

```java
/**
 * 网站访问统计服务类
 */
@Service
public class VisitService {
	// 设定访问次数Redis键名
	private final static String KEY = "visit_count";

	// 注入redisTemplate操作Redis
	@Autowired
	private RedisTemplate<String, String> redisTemplate;

	// 获取当前访问次数
	public String getCurrentCount() {
		String count = redisTemplate.opsForValue().get(KEY);
		if (count == null || "".equals(count)) {
			redisTemplate.opsForValue().set(KEY, "0");
			return "0";
		}
		return count;
	}

	// 访问次数加1
	public void addCount() {
		redisTemplate.opsForValue().increment(KEY, 1);
	}
}
```

### 2.5 并发访问测试

我们通过测试类发起并发访问测试，代码如下：

**实例：**

```java
/**
 * 访问统计服务测试
 */
@SpringBootTest
class VisitServiceTest {
	private Logger logger = LoggerFactory.getLogger(this.getClass());
	@Autowired
	private VisitService visitService;

	@Test
	void test() {
		logger.info("访问次数：{}", visitService.getCurrentCount());
		// 使用线程池快速发起10000次访问
		ExecutorService cachedThreadPool = Executors.newCachedThreadPool();
		for (int i = 0; i < 10000; i++) {
			cachedThreadPool.execute(new Runnable() {
				public void run() {
					visitService.addCount();
				}
			});
		}
	}
}
```

此时我们通过 Redis 客户端发现 visit_count 的值如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2k4bjej60hu01rmx602)

并发访问测试结果

> **Tips**：Redis 中的操作都是原子性的，要么执行，要么不执行，在高并发场景下依然可以准确的进行计数，关键是速度还非常之快！

## 3. 热门商品排行榜

如果是大型网站，时刻有很多用户在访问网页，对热门商品排行榜的访问频率是非常恐怖的。

我们可以通过定时器，定时从关系数据库中取出热门商品数据放入 Redis 缓存，用户访问网页时，直接从缓存中获取热门商品数据。这将大大提高响应速度，并降低对关系数据库的性能损耗。

### 3.1 定义商品类

我们简单的定义一个商品类，便于展现商品排行榜数据。

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
}
```

### 3.2 开发商品排行榜服务类

开发商品排行榜服务类，负责从数据库查询最新排行榜信息，并更新到 Redis ，以及从 Redis 中取出排行榜信息。

**实例：**

```java
/**
 * 商品排行榜服务类
 */
@Service
public class GoodsRankService {
	// 设定商品排行榜Redis键名
	private final static String KEY = "goods_rank";

	// 注入redisTemplate操作Redis
	@Autowired
	private RedisTemplate<String, String> redisTemplate;

	// 更新Redis缓存的排行榜
	public void updateRankList() throws JsonProcessingException {
		// 此处直接定义商品排行榜，真实场景应为从数据库获取
		List<GoodsDo> rankList = new ArrayList<GoodsDo>();
		GoodsDo goods = new GoodsDo();
		goods.setId(1L);
		goods.setName("鸡蛋" + new Date());// 添加时间信息，以便测试缓存更新了
		rankList.add(goods);
		// 将rankList序列化后写入Reidis
		ObjectMapper mapper = new ObjectMapper();
		redisTemplate.opsForValue().set(KEY, mapper.writeValueAsString(rankList));
	}

	// 获取Redis缓存的排行榜
	public List<GoodsDo> getRandkList() throws JsonMappingException, JsonProcessingException {
		ObjectMapper mapper = new ObjectMapper();
		return mapper.readValue(redisTemplate.opsForValue().get(KEY), List.class);
	}
}
```

### 3.3 通过定时器更新排行榜

为启动类添加 @EnableScheduling 注解，以便开启定时任务，然后编写 RankListUpdateTask 类定时刷新排行榜。

**实例：**

```java
/**
 * 排行榜更新任务
 */
@Component
public class RankListUpdateTask {
	@Autowired
	private GoodsRankService goodsRankService;

	/**
	 * 容器启动后马上执行，且每1秒执行1次
	 */
	@Scheduled(initialDelay = 0, fixedRate = 1000)
	public void execute() throws InterruptedException, JsonProcessingException {
		goodsRankService.updateRankList();
	}
}
```

### 3.4 开发控制器方法

我们还需要一个控制器方法，用于演示获取商品列表的结果。

**实例：**

```java
@RestController
public class GoodsRankController {
	@Autowired
	private GoodsRankService goodsRankService;

	@GetMapping("getRankList")
	public List getRankList() throws Exception {
		return goodsRankService.getRandkList();
	}
}
```

### 3.5 测试

运行启动类，然后访问 `http://127.0.0.1:8080/getRankList` ，结果如下：

```java
[{"id":1,"name":"鸡蛋Thu May 28 22:47:33 CST 2020","price":null,"pic":null}]
```

稍等会再次访问，结果如下：

```java
[{"id":1,"name":"鸡蛋Thu May 28 22:48:09 CST 2020","price":null,"pic":null}]
```

说明我们设计的缓存机制生效了。

## 4. 小结

开发的项目多了，越来越能体会，传统数据库访问速度是限制系统性能的最大瓶颈。

而 Redis 基于内存的特性，可以极大地提高读写效率，使用得当，往往使系统性能有质的提升。

Spring Boot 可以非常方便地集成 Redis ，当我们在项目开发中遇到访问频率非常高的热点数据时，可以优先考虑使用 Redis 进行存储操作。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
