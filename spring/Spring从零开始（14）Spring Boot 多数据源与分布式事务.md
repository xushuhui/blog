# Spring Boot 多数据源与分布式事务

## 1. 前言

一个项目中使用多个数据源的需求，我们在日常工作中时常会遇到。

以商城系统为例，有一个 MySQL 的数据库负责存储交易数据。公司还有一套 ERP 企业信息化管理系统，要求订单信息同步录入 ERP 数据库，便于公司统一管理，而该 ERP 系统采用的数据库为 SQL Server 。

此时，就可以在 Spring Boot 项目中配置多个数据源。另外，使用多数据源后，需要采用分布式事务来保持数据的完整性。

## 2. 实例场景

本小节我们使用 Spring Boot 开发一个商城系统的订单生成功能，订单信息同时进入 MySQL 与 SQL Server 数据库。

## 3. 数据库模块实现

首先创建 MySQL 数据库 shop ，并新建订单表 order ，表结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo259cozj60m806yjsz02)

order 表结构

然后创建 SQL Server 数据库 erpshop ，并新建订单表 erp_order ，表结构如下。注意 id 是自增长的唯一标识，out_id 是对应订单在 MySQL 数据库中的唯一标识，以便在两个库中比对订单。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo25rb3uj60m303n3zq02)

erp_order 结构

## 4. Spring Boot 后端实现

接下来，我们开始实现 Spring Boot 后端项目，数据持久层采用 MyBatis 框架，同时访问两个数据源。

### 4.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-multidb，生成项目后导入 Eclipse 开发环境。

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
		<!-- SQL Server驱动 -->
		<dependency>
			<groupId>com.microsoft.sqlserver</groupId>
			<artifactId>mssql-jdbc</artifactId>
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

由于我们要同时访问两个数据库，所以需要在配置文件中添加两个数据源的配置信息。注意配置多数据源时， url 配置需要使用 `spring.datasource.db1.jdbc-url=xxx` 的形式。

**实例：**

```java
# 数据源1 MySQL
spring.datasource.db1.driver-class-name=com.mysql.jdbc.Driver
spring.datasource.db1.jdbc-url=jdbc:mysql://127.0.0.1:3306/shop?useUnicode=true&characterEncoding=utf-8&serverTimezone=UTC
spring.datasource.db1.username=root
spring.datasource.db1.password=Easy@0122
# 数据源2 SQL Server
spring.datasource.db2.driverClassName = com.microsoft.sqlserver.jdbc.SQLServerDriver
spring.datasource.db2.jdbc-url =jdbc:sqlserver://127.0.0.1:1433;DatabaseName=erpshop
spring.datasource.db2.username =sa
spring.datasource.db2.password =Easy@0122
```

### 4.4 注册数据源组件

多个数据源的情况下， 我们需要通过配置类，将数据源注册为组件放入 Spring 容器中。

**实例：**

```java
/**
 * 数据源配置类
 */
@Configuration//标注为配置类
public class DataSourceConfig {
	/**
	 * 数据源1
	 */
	@Bean//返回值注册为组件
	@ConfigurationProperties("spring.datasource.db1")//使用spring.datasource.db1作为前缀的配置
	public DataSource db1() {
		return DataSourceBuilder.create().build();
	}
	/**
	 * 数据源2
	 */
	@Bean//返回值注册为组件
	@ConfigurationProperties("spring.datasource.db2")//使用spring.datasource.db2作为前缀的配置
	public DataSource db2() {
		return DataSourceBuilder.create().build();
	}
}
```

通过这个配置类， Spring 容器中就有两个数据源组件，这两个组件分别采用 `spring.datasource.db1` 和 `spring.datasource.db2` 开头的配置信息。所以通过这两个组件，就能分别操作 MySQL 数据源 1 和 SQL Sever 数据源 2 。

### 4.5 MyBatis 配置

多数据源情况下， MyBatis 中的关键组件 SqlSessionFactory 和 SqlSessionTemplate 也需要单独配置，我们需要为两个数据源分别配置一套组件。

**实例：**

```java
/**
 * 数据源1 MyBatis配置
 */
@Configuration
@MapperScan(value = "com.imooc.springbootmultidb.mapper1", sqlSessionFactoryRef = "sqlSessionFactory1")
public class Db1MyBatisConfig {
	@Autowired // 自动装配
	@Qualifier("db1") // 指定注入名为db1的组件
	private DataSource db1;

	@Bean
	public SqlSessionFactory sqlSessionFactory1() throws Exception {
		SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
		sqlSessionFactoryBean.setDataSource(db1);// sqlSessionFactory1使用的数据源为db1
		sqlSessionFactoryBean
				.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mapper1/*.xml"));
		return sqlSessionFactoryBean.getObject();
	}

	@Bean
	public SqlSessionTemplate sqlSessionTemplate1() throws Exception {
		return new SqlSessionTemplate(sqlSessionFactory1());// sqlSessionTemplate1使用的数据源也是关联到db1
	}
}

```

通过上面的配置类， `com.imooc.springbootmultidb.mapper1` 包中的 DAO 数据访问接口会自动调用 sqlSessionTemplate1 组件实现具体数据库操作，而 sqlSessionTemplate1 操作的数据源已经通过配置类设置为 db1 。同时， DAO 数据访问接口对应的映射文件已经指定到 `classpath:mapper1/` 目录去寻找。这样数据源 – DAO 数据访问接口 – 映射文件三者的对应关系就建立起来了。

数据源 2 的配置方法是一样的，`com.imooc.springbootmultidb.mapper2` 包中的 DAO 数据访问接口会自动调用 sqlSessionTemplate2 组件，其操作的数据源即为 db2 ，其对应的映射文件指定到 `classpath:mapper2/` 目录去寻找。

**实例：**

```java
/**
 * 数据源2 MyBatis配置
 */
@Configuration
@MapperScan(value = "com.imooc.springbootmultidb.mapper2", sqlSessionFactoryRef = "sqlSessionFactory2")
public class Db2MyBatisConfig {
	@Autowired // 自动装配
	@Qualifier("db2") // 指定注入名为db1的组件
	private DataSource db2;

	@Bean
	public SqlSessionFactory sqlSessionFactory2() throws Exception {
		SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
		sqlSessionFactoryBean.setDataSource(db2);// sqlSessionFactory2使用的数据源为db2
		sqlSessionFactoryBean
				.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mapper2/*.xml"));
		return sqlSessionFactoryBean.getObject();
	}

	@Bean
	public SqlSessionTemplate sqlSessionTemplate2() throws Exception {
		return new SqlSessionTemplate(sqlSessionFactory2());// sqlSessionTemplate2使用的数据源也是关联到db2
	}
}

```

### 4.6 数据访问接口实现

数据访问接口的位置已经在配置类指定，首先在 `com.imooc.springbootmultidb.mapper1` 创建 OrderDao ，操作的是数据源 1 中的 order 表。

**实例：**

```java
/**
 * 数据访问接口
 */
@Repository
public interface OrderDao {
	public int insert(OrderDo order);
}
```

然后在 `com.imooc.springbootmultidb.mapper2` 创建 ErpOrderDao ，操作的是数据源 2 中的 erporder 表。

**实例：**

```java
/**
 * 数据访问接口
 */
@Repository
public interface ErpOrderDao {
	public int insert(ErpOrderDo erpOrder);
}
```

这两个接口中使用的数据对象比较简单，代码如下：

**实例：**

```java
/**
 * 订单数据类
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
/**
 * ERP订单数据类
 */
public class ErpOrderDo {
	/**
	 * 订单id
	 */
	private Long id;
	/**
	 * 商城系统订单id
	 */
	private Long outId;
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

### 4.7 编写映射文件

分别针对 OrderDao 、 ErpOrderDao 编写对应的映射文件，然后按照配置类指定的位置，两个文件分别放到 `resources/mapper1` 和 `resources/mapper2` 目录下。

**实例：**

```java
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!-- 本映射文件对应OrderDao接口 -->
<mapper namespace="com.imooc.springbootmultidb.mapper1.OrderDao">
	<!-- 对应OrderDao中的insert方法 -->
	<insert id="insert"
		parameterType="com.imooc.springbootmultidb.mapper1.OrderDo"
		useGeneratedKeys="true" keyProperty="id">
		insert into `order`
		(goods_id,count) values (#{goodsId},#{count})
	</insert>
</mapper>
```

**实例：**

```java
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!-- 本映射文件对应ErpOrderDao接口 -->
<mapper
	namespace="com.imooc.springbootmultidb.mapper2.ErpOrderDao">
	<!-- 对应ErpOrderDao中的insert方法 -->
	<insert id="insert"
		parameterType="com.imooc.springbootmultidb.mapper2.ErpOrderDo">
		insert into erp_order (out_id,goods_id,count) values
		(#{outId},#{goodsId},#{count})
	</insert>
</mapper>
```

### 4.8 多数据源测试

数据操作接口与对应的映射文件均已编写完毕，现在可以通过测试类进行多数据源测试了，我们在测试类中同时向两个库插入记录。

**实例：**

```java
/**
 * 多数据源测试
 */
@SpringBootTest
class MultidbTest {
	@Autowired
	private OrderDao orderDao;// 对应数据源1
	@Autowired
	private ErpOrderDao erpOrderDao;// 对应数据源2

	/**
	 * 插入测试
	 */
	@Test
	void testInsert() {
		// 数据源1插入数据
		OrderDo order = new OrderDo();
		order.setCount(1L);
		order.setGoodsId(1L);
		int affectRows1 = orderDao.insert(order);
		// 数据源2插入数据
		ErpOrderDo erpOrder = new ErpOrderDo();
		erpOrder.setCount(order.getCount());
		erpOrder.setGoodsId(order.getGoodsId());
		erpOrder.setOutId(order.getId());
		int affectRows2 = erpOrderDao.insert(erpOrder);
		assertEquals(1, affectRows1);
		assertEquals(1, affectRows2);
	}
}
```

运行测试方法后，两个数据库表中均新增数据成功，这样我们就成功的使用 Spring Boot 同时操作了两个数据源。

## 5. 分布式事务

采用多数据源之后，事务的实现方式也随之发生变化。当某个数据源操作出现异常时，该数据源和其他数据源的事务都需要回滚。这种涉及多个数据源的事务，称为分布式事务，接来下我们就来具体实现一下。

### 5.1 引入分布式事务依赖

在 pom.xml 引入 Atomikos 事务管理器相关的依赖项， Atomikos 是一个开源的事务管理器，支持分布式事务。

**实例：**

```java
		<!--分布式事务 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-jta-atomikos</artifactId>
		</dependency>
```

### 5.2 更换数据源组件

需要将默认的数据源更换为支持分布式事务的数据源， MySQL 对应的数据源为 MysqlXADataSource ， SQL Server 对应的数据源为 SQLServerXADataSource 。

**实例：**

```java

/**
 * 数据源配置类
 */
@Configuration // 标注为配置类
public class DataSourceConfig {
	// 注入数据源1配置项
	@Value("${spring.datasource.db1.jdbc-url}")
	private String db1_url;
	@Value("${spring.datasource.db1.username}")
	private String db1_username;
	@Value("${spring.datasource.db1.password}")
	private String db1_password;
	// 注入数据源2配置项
	@Value("${spring.datasource.db2.jdbc-url}")
	private String db2_url;
	@Value("${spring.datasource.db2.username}")
	private String db2_username;
	@Value("${spring.datasource.db2.password}")
	private String db2_password;

	/**
	 * 数据源1
	 */
	@Bean // 返回值注册为组件
	public DataSource db1() throws SQLException {
		MysqlXADataSource dataSource = new MysqlXADataSource();
		dataSource.setUrl(db1_url);
		dataSource.setUser(db1_username);
		dataSource.setPassword(db1_password);
		AtomikosDataSourceBean atomikosDataSourceBean = new AtomikosDataSourceBean();
		atomikosDataSourceBean.setXaDataSource(dataSource);
		atomikosDataSourceBean.setUniqueResourceName("db1");
		return atomikosDataSourceBean;

	}

	/**
	 * 数据源2
	 */
	@Bean // 返回值注册为组件
	public DataSource db2() {
		SQLServerXADataSource dataSource = new SQLServerXADataSource();
		dataSource.setURL(db2_url);
		dataSource.setUser(db2_username);
		dataSource.setPassword(db2_password);
		AtomikosDataSourceBean atomikosDataSourceBean = new AtomikosDataSourceBean();
		atomikosDataSourceBean.setXaDataSource(dataSource);
		atomikosDataSourceBean.setUniqueResourceName("db2");
		return atomikosDataSourceBean;
	}
}

```

### 5.3 添加分布式事务管理器组件

继续修改 DataSourceConfig 类，在其中配置分布式事务管理器组件。当项目中使用事务时，会通过配置的分布式事务管理器管理分布式事务操作。

**实例：**

```java
	/**
	 * 分布式事务管理器
	 */
	@Bean(name = "jtaTransactionManager")
	public JtaTransactionManager jtaTransactionManager() {
		UserTransactionManager userTransactionManager = new UserTransactionManager();
		UserTransaction userTransaction = new UserTransactionImp();
		return new JtaTransactionManager(userTransaction, userTransactionManager);
	}
```

### 5.4 测试分布式事务

在测试方法上添加 `@Transactional` 开启事务，然后在两个数据源操作中间模拟抛出异常。

**实例：**

```java
	/**
	 * 插入测试
	 */
	@Test
	@Transactional // 开启事务
	void testInsert() {
		// 数据源1插入数据
		OrderDo order = new OrderDo();
		order.setCount(1L);
		order.setGoodsId(1L);
		int affectRows1 = orderDao.insert(order);
		// 模拟抛出异常
		int a = 1 / 0;
		// 数据源2插入数据
		ErpOrderDo erpOrder = new ErpOrderDo();
		erpOrder.setCount(order.getCount());
		erpOrder.setGoodsId(order.getGoodsId());
		erpOrder.setOutId(order.getId());
		int affectRows2 = erpOrderDao.insert(erpOrder);
		assertEquals(1, affectRows1);
		assertEquals(1, affectRows2);
	}
```

此时运行测试类，可以发现数据源 1 的事务已回滚，验证成功！

> **Tips**：如果运行测试类报错 `master..xp_sqljdbc_xa_init_ex` 相关信息，是 SQL Server 默认配置不支持分布式事务问题，可查询相关资料解决该问题。

## 6. 小结

在开发 Spring Boot 项目时，如果默认配置满足不了我们的需求，可以通过手工配置组件实现我们需要的功能。这些组件可能是各个公司提供的，我们根据相应文档，为其配置各个属性即可。

* 配置多个数据源时，通过配置类，逐一配置数据源组件及其参数。
* 配置 MyBatis 时，手工配置 MyBatis 相关组件，并指定相应扫描的 DAO 类及映射文件。
* 使用分布式事务时，使用支持分布式事务的数据源组件，并配置分布式事务管理器。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
