# Spring Boot 使用事务

## 1. 前言

工作中确实碰到过一些不知道使用事务的朋友，毫无疑问会给项目带来一些风险。

举个简单的例子吧，网购的时候需要扣减库存，同时生成订单。如果扣库存成功了，没生成订单，结果是库存不知道为何变少了；如果生成订单了，没扣库存，那就有可能卖出去的数量比库存还多。

这两种情况都是不能接受的，我们必须保证这两个对数据库的更新操作同时成功，或者同时失败。

事务就是这样一种机制，将对数据库的一系列操作视为一个执行单元，保证单元内的操作同时成功，或者当有一个操作失败时全部失败。

## 2. 实例场景

在 Spring Boot 中使用事务非常简单，本小节我们通过商品扣减库存、生成订单的实例，演示下 Spring Boot 中使用事务的具体流程。

## 3. 数据库模块实现

需要有一个商品表，保存商品的唯一标识、名称、库存数量，结构如下：

**实例：**

```java
CREATE TABLE `goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `num` bigint(255) DEFAULT NULL COMMENT '库存数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
```

购买商品后还需要生成订单，保存订单唯一标识、购买商品的 id 、购买数量。

**实例：**

```java
CREATE TABLE `order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `goods_id` bigint(20) DEFAULT NULL COMMENT '商品id',
  `count` bigint(20) DEFAULT NULL COMMENT '购买数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## 4. Spring Boot 后端实现

接下来，我们开始开发 Spring Boot 后端项目，并且使用事务实现扣减库存、生成订单功能。数据库访问部分使用比较流行的 MyBatis 框架。

### 4.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-transaction，生成项目后导入 Eclipse 开发环境。

### 4.2 引入项目依赖

我们引入热部署依赖、 Web 依赖、数据库访问相关依赖及测试相关依赖，具体如下：

**实例：**

```java
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter</artifactId>
		</dependency>
		<!-- 热部署 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-devtools</artifactId>
		</dependency>
		<!-- Web支持 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- JDBC -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-jdbc</artifactId>
		</dependency>
		<!-- MySQL驱动 -->
		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
		</dependency>
		<!-- 集成MyBatis -->
		<dependency>
			<groupId>org.mybatis.spring.boot</groupId>
			<artifactId>mybatis-spring-boot-starter</artifactId>
			<version>2.1.2</version>
		</dependency>
		<!-- junit -->
		<dependency>
			<groupId>junit</groupId>
			<artifactId>junit</artifactId>
			<scope>test</scope>
		</dependency>
		<!-- 测试 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-test</artifactId>
			<scope>test</scope>
			<exclusions>
				<exclusion>
					<groupId>org.junit.vintage</groupId>
					<artifactId>junit-vintage-engine</artifactId>
				</exclusion>
			</exclusions>
		</dependency>
```

### 4.3 数据源配置

修改 application.properties 文件，配置数据源信息。

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
```

### 4.4 开发数据对象类

开发 goods 表对应的数据对象类 GoodsDo ，代码如下：

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
    * 商品库存
    */
   private Long num;
   // 省略 get set
}
```

然后开发 order 表对应的数据对象类 OrderDo，代码如下：

**实例：**

```java
/**
 * 订单类
 */
public class OrderDo {
	/**
	 * 订单id
	 */
	private Long id;
	/**
	 * 商品id
	 */
	private Long goodsId;
	/**
	 * 购买数量
	 */
	private Long count;
	// 省略 get set
}
```

### 4.5 开发数据访问层

首先定义商品数据访问接口，实现查询剩余库存与扣减库存功能。

**实例：**

```java
/**
 * 商品数据库访问接口
 */
@Repository // 标注数据访问组件
public interface GoodsDao {
	/**
	 * 查询商品信息(根据id查询单个商品信息)
	 */
	public GoodsDo selectForUpdate(Long id);

	/**
	 * 修改商品信息(根据id修改其他属性值)
	 */
	public int update(GoodsDo Goods);
}
```

注意，在查询商品剩余库存时，我们采用面向对象的方法，将对应 id 的商品信息全部取出，更加方便点。采用 selectForUpdate 命名，表示该方法使用了 `select ... for update` 的 SQL 语句查询方式，以锁定数据库对应记录，规避高并发场景下库存修改错误问题。同样 update 方法也采用了面向对象的方式，根据 id 修改其他信息，方便复用。

然后定义订单数据访问接口，实现生成订单的功能。

**实例：**

```java
/**
* 订单数据库访问接口
*/
@Repository // 标注数据访问组件
public interface OrderDao {
   /**
    * 新增订单
    */
   public int insert(OrderDo order);
}
```

然后，我们修改 Spring Boot 配置类，添加 @MapperScan 注解，扫描数据访问接口所在的包。

**实例：**

```java
@SpringBootApplication
@MapperScan("com.imooc.springboottransaction") // 指定MyBatis扫描的包，以便将数据访问接口注册为Bean
public class SpringBootTransactionApplication {
	public static void main(String[] args) {
		SpringApplication.run(SpringBootTransactionApplication.class, args);
	}
}
```

### 4.6 添加 MyBatis 映射文件

编写 GoodsDao 、 OrderDao 对应的映射文件， 首先我们通过 application.properties 指定映射文件的位置：

**实例：**

```java
# 指定MyBatis配置文件位置
mybatis.mapper-locations=classpath:mapper/*.xml
```

然后在 resources/mapper 目录下新建 GoodsMapper.xml 文件，该文件就是 goods 表对应的映射文件，内容如下：

**实例：**

```java
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!-- 本映射文件对应GoodsDao接口 -->
<mapper namespace="com.imooc.springboottransaction.GoodsDao">
   <!-- 对应GoodsDao中的selectForUpdate方法 -->
   <select id="selectForUpdate" resultMap="resultMapBase" parameterType="java.lang.Long">
   	select <include refid="sqlBase" /> from goods where id = #{id} for update
   </select>
   <!-- 对应GoodsDao中的update方法 -->
   <update id="update" parameterType="com.imooc.springboottransaction.GoodsDo">
   	update goods set name=#{name},num=#{num} where id=#{id}
   </update>
   <!-- 可复用的sql模板 -->
   <sql id="sqlBase">
   	id,name,num
   </sql>
   <!-- 保存SQL语句查询结果与实体类属性的映射 -->
   <resultMap id="resultMapBase" type="com.imooc.springboottransaction.GoodsDo">
   	<id column="id" property="id" />
   	<result column="name" property="name" />
   	<result column="num" property="num" />
   </resultMap>
</mapper>
```

同样我们在 resources/mapper 目录下新建 OrderMapper.xml 文件，该文件是 order 表对应的映射文件，内容如下：

**实例：**

```java
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!-- 本映射文件对应OrderDao接口 -->
<mapper namespace="com.imooc.springboottransaction.OrderDao">
   <!-- 对应OrderDao中的insert方法 -->
   <insert id="insert" parameterType="com.imooc.springboottransaction.OrderDo">
   	insert into `order` (goods_id,count) values (#{goodsId},#{count})
   </insert>
</mapper>
```

### 4.7 编写服务方法

下单这个操作，可以封装为一个服务方法，不管是手机端下单还是电脑端下单都可以调用。

我们新建订单服务类 OrderService ，并在其中实现下单方法 createOrder ，代码如下：

**实例：**

```java
/**
 * 订单服务类
 */
@Service // 注册为服务类
public class OrderService {
	@Autowired
	private GoodsDao goodsDao;
	@Autowired
	private OrderDao orderDao;

	/**
	 * 下单
	 *
	 * @param goodsId 购买商品id
	 * @param count   购买商品数量
	 * @return 生成订单数
	 */
	@Transactional // 实现事务
	public int createOrder(Long goodsId, Long count) {
		// 锁定商品库存
		GoodsDo goods = goodsDao.selectForUpdate(goodsId);
		// 扣减库存
		Long newNum = goods.getNum() - count;
		goods.setNum(newNum);
		goodsDao.update(goods);
		// 生成订单
		OrderDo order = new OrderDo();
		order.setGoodsId(goodsId);
		order.setCount(count);
		int affectRows = orderDao.insert(order);
		return affectRows;
	}
}
```

我们在 createOrder 方法上添加了 `@Transactional` 注解，该注解为 createOrder 方法开启了事务，当方法结束时提交事务。这样保证了 createOrder 内方法全部执行成功，或者全部失败。

## 5. 测试

### 5.1 构造测试数据

在数据库中构造一条测试数据如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo24gv1pj60jj039gmg02)

测试数据

### 5.2 正常测试

编写测试方法发起测试：

**实例：**

```java
/**
 * 订单测试
 */
@SpringBootTest
class OrderTest {

	@Autowired
	private OrderService orderService;

	/**
	 * 新增一个商品
	 */
	@Test
	void testCreateOrder() {
		// 购买id为1的商品1份
		int affectRows = orderService.createOrder(1L, 1L);
		assertEquals(1, affectRows);
	}
}
```

运行测试方法后，手机的库存变为 19 ，且生成一条订单记录，测试通过，具体结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo24oz93j60kp04hmym02)

正常测试结果

### 5.3 模拟异常测试

修改下单方法，在扣减库存后抛出异常，看看事务能否回滚到修改全部未发生的状态。为了便于测试我们将库存重新设为 20 ，然后将下单方法修改如下：

**实例：**

```java
	@Transactional // 实现事务
	public int createOrder(Long goodsId, Long count) {
		// 锁定商品库存
		GoodsDo goods = goodsDao.selectForUpdate(goodsId);
		// 扣减库存
		Long newNum = goods.getNum() - count;
		goods.setNum(newNum);
		goodsDao.update(goods);
		// 模拟异常
		int a=1/0;
		// 生成订单
		OrderDo order = new OrderDo();
		order.setGoodsId(goodsId);
		order.setCount(count);
		int affectRows = orderDao.insert(order);
		return affectRows;
	}
```

运行测试方法后，抛出异常，查看数据库发现，库存还是 20 ，说明 `goodsDao.update(goods);` 的修改没有提交到数据库，具体结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo24xm7jj60k604b3zy02)

模拟异常测试结果

## 6. 使用注意事项

Spring 事务在一些情况下不能生效，需要特别注意。

### 6.1 抛出检查型异常时事务失效

首先了解下异常类型：

* **Exception 受检查的异常**：在程序中必须使用 try…catch 进行处理，遇到这种异常不处理，编译器会报错。例如 IOException 。
* **RuntimeException 非受检查的异常**：可以不使用 try…catch 进行处理。例如常见的 NullPointerException 。

在大多数人潜意识中，只要发生异常，事务就应该回滚，实际上使用 @Transactional 时，默认只对非受检查异常回滚。例如：

**实例：**

```java
	@Transactional // 实现事务
	public int createOrder(Long goodsId, Long count) {
		// 锁定商品库存
		GoodsDo goods = goodsDao.selectForUpdate(goodsId);
		// 扣减库存
		Long newNum = goods.getNum() - count;
		goods.setNum(newNum);
		goodsDao.update(goods);
		if (count > goods.getNum()) {
			// 非受检查异常抛出时，会回滚
			throw new RuntimeException();
		}
		// 生成订单
		OrderDo order = new OrderDo();
		order.setGoodsId(goodsId);
		order.setCount(count);
		int affectRows = orderDao.insert(order);
		return affectRows;
	}
```

**实例：**

```java
   @Transactional // 实现事务
   public int createOrder(Long goodsId, Long count) throws Exception {
   	// 锁定商品库存
   	GoodsDo goods = goodsDao.selectForUpdate(goodsId);
   	// 扣减库存
   	Long newNum = goods.getNum() - count;
   	goods.setNum(newNum);
   	goodsDao.update(goods);
   	if (count > goods.getNum()) {
   		//注意！此处为受检查的异常，就算抛出也不会回滚
   		throw new Exception();
   	}
   	// 生成订单
   	OrderDo order = new OrderDo();
   	order.setGoodsId(goodsId);
   	order.setCount(count);
   	int affectRows = orderDao.insert(order);
   	return affectRows;
   }
```

如果想实现只要抛出异常就回滚，可以通过添加注解 `@Transactional(rollbackFor=Exception.class)` 实现。

**实例：**

```java
	@Transactional(rollbackFor = Exception.class) // 抛出异常即回滚
	public int createOrder(Long goodsId, Long count) throws Exception {
		// 锁定商品库存
		GoodsDo goods = goodsDao.selectForUpdate(goodsId);
		// 扣减库存
		Long newNum = goods.getNum() - count;
		goods.setNum(newNum);
		goodsDao.update(goods);
		if (count > goods.getNum()) {
			throw new Exception();
		}
		// 生成订单
		OrderDo order = new OrderDo();
		order.setGoodsId(goodsId);
		order.setCount(count);
		int affectRows = orderDao.insert(order);
		return affectRows;
	}
```

OK，我们将在测试类中，将购买数量设为大于库存数量的 100 ，然后一次测试上面三种情况，就能验证上面的说法了。

**实例：**

```java
/**
 * 订单测试
 */
@SpringBootTest
class OrderTest {

	@Autowired
	private OrderService orderService;

	/**
	 * 创建订单测试
	 */
	@Test
	void testCreateOrder() throws Exception {
		// 购买id为1的商品1份
		int affectRows = orderService.createOrder(1L, 100L);
		assertEquals(1, affectRows);
	}
}
```

### 6.2 一个事务方法调用另一个事务方法时失效

先看下面的实例，我们修改下 OrderService 类，通过一个事务方法调用 createOrder 方法。

**实例：**

```java
/**
 * 订单服务类
 */
@Service // 注册为服务类
public class OrderService {
	@Autowired
	private GoodsDao goodsDao;
	@Autowired
	private OrderDao orderDao;

	@Transactional // 开启事务
	public int startCreateOrder(Long goodsId, Long count) throws Exception {
		return this.createOrder(goodsId, count);
	}

	/**
	 * 下单
	 *
	 * @param goodsId 购买商品id
	 * @param count   购买商品数量
	 * @return 生成订单数
	 */
	@Transactional(rollbackFor = Exception.class) // 抛出异常即回滚
	public int createOrder(Long goodsId, Long count) throws Exception {
		// 锁定商品库存
		GoodsDo goods = goodsDao.selectForUpdate(goodsId);
		// 扣减库存
		Long newNum = goods.getNum() - count;
		goods.setNum(newNum);
		goodsDao.update(goods);
		if (count > goods.getNum()) {
			// 非受检查异常抛出时，会回滚
			throw new Exception();
		}
		// 生成订单
		OrderDo order = new OrderDo();
		order.setGoodsId(goodsId);
		order.setCount(count);
		int affectRows = orderDao.insert(order);
		return affectRows;
	}
}
```

此时我们在测试类中通过 startCreateOrder 方法再去调用 createOrder 方法，代码如下：

**实例：**

```java
/**
 * 订单测试
 */
@SpringBootTest
class OrderTest {

	@Autowired
	private OrderService orderService;

	/**
	 * 创建订单测试
	 */
	@Test
	void testCreateOrder() throws Exception {
		// 购买id为1的商品1份
		int affectRows = orderService.startCreateOrder(1L, 100L);
		assertEquals(1, affectRows);
	}
}
```

startCreateOrder 和 createOrder 方法都是事务方法，且这两个方法事务特性不同 （一个没有 rollbackFor=Exception.class)，如果我们调用 startTransaction 方法，则 createOrder 中的事务并不会生效。

也就是说，如果在同一个类中，一个事务方法调用另一个事务方法，可能会导致被调用的事务方法的事务失效！

这是因为 Spring 的声明式事务使用了代理，具体机制此处不再探讨，但是一定要注意规避这种事务失效的场景。

## 7. 小结

Spring Boot 中的事务使用非常简单，是因为进行了高度的封装。正是由于封装的很彻底，所以我们一般接触不到其具体原理和实现方式，这就需要我们注意一些事务可能失效的情况，避免因事务失效带来风险和损失。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
