# Spring Boot 集成 Druid 数据源

## 1. 前言

首先要理解数据源的作用，数据源实际上是一个接口 `javax.sql.DataSource` 。 Spring Boot 在操作数据库时，是通过数据源类型的组件实现的。

数据源的类型有很多，但是都实现了 `javax.sql.DataSource` 接口，所以 Spring Boot 可以尽情地更换各类数据源以实现不同的需求。

其中 Druid 数据源就是数据源中比较出类拔萃的存在，而且是阿里开发的。国人出品的东西，咱们能支持的必须得支持啊，当然这是建立在人家做的确实好的基础上。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo22pjsvj60xd0l7tc202)

 Druid 是阿里巴巴数据库事业部出品

本篇我们使用 Druid 替换默认的数据源，然后做一下性能对比测试。网上有很多文章写 Druid 性能如何如何强悍，但是很多并没有事实依据。我们做程序员还是要严谨，相信实践是检验真理的唯一标准。所以本篇的内容就是研究下 Druid 如何使用，及其性能到底是否足够优异。

## 2. 使用默认数据源（HikariDataSource）

Spring Boot 2.2.5 版本使用的默认数据源是 HikariDataSource ，该数据源号称拥有全世界最快的数据库连接池，嗯，我们来试试它的深浅。

### 2.1 准备数据库

还是使用之前的商城数据库（shop）及商品信息数据表（goods），表结构如下：

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

### 2.2 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-hikari` ，生成项目后导入 Eclipse 开发环境。

### 2.3 引入项目依赖

我们引入 Web 项目依赖、热部署依赖。由于本项目需要访问数据库，所以引入 `spring-boot-starter-jdbc` 依赖和 `mysql-connector-java` 依赖。

pom.xml 文件中依赖项如下：

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

### 2.4 构建商品类和商品数据访问类

定义商品类，对应商品表：

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

定义商品数据库访问类：

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

### 2.5 配置数据源信息

通过配置文件，设置数据源信息。

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

### 2.6 测试

通过测试类发起测试，此处我们简单执行 1000 次插入，看看执行时间。

需要注意的是，Spring Boot 进行测试时，需要添加注解 `@SpringBootTest` 。添加注解后该类可以直接通过 `@Test` 标注的方法发起单元测试，容器环境都已准备好，非常方便。

**实例：**

```java
@SpringBootTest // 通过该注解，开启测试类功能，当测试方法启动时，启动了Spring容器
class SpringBootHikariApplicationTests {
	@Autowired
	private DataSource dataSource;// 自动注入数据源
	@Autowired
	private GoodsDao goodsDao;

	/**
	 * 打印数据源信息
	 */
	@Test // 测试方法
	void printDataSource() {
		System.out.println(dataSource);
	}

	/**
	 * 批量插入测试
	 */
	@Test
	void insertBatch() {
		// 开始时间
		long startTime = System.currentTimeMillis();
		// 执行1000次插入
		GoodsDo goods = new GoodsDo();
		goods.setName("测试");
		goods.setPic("测试图片");
		goods.setPrice("1.0");
		for (int i = 0; i < 1000; i++) {
			goodsDao.insert(goods);
		}
		// 输出操作时间
		System.out.println("use time:" + (System.currentTimeMillis() - startTime)+"ms");
	}
}
```

输出结果如下，可见默认数据源类型为 `HikariDataSource` ，插入 1000 条数据的时间大概为 1500ms （注意时间可能跟电脑性能等很多因素相关，此处只是进行简单的对比测试）。

```java
use time:1518ms
com.zaxxer.hikari.HikariDataSource
```

## 3. 使用 Druid 数据源

接下来我们使用 Druid 数据源进行对比测试。

### 3.1 准备数据库

与上面的商城数据库（shop）及商品信息数据表（goods）一致。

### 3.2 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ， Group 为 `com.imooc` ， Artifact 为 `spring-boot-druid` ，生成项目后导入 Eclipse 开发环境。

### 3.3 引入项目依赖

我们引入 Web 项目依赖、热部署依赖。由于本项目需要访问数据库，所以引入 `spring-boot-starter-jdbc` 依赖和 `mysql-connector-java` 依赖，由于使用 Druid ，所以还需要添加 Druid 相关依赖。

pom.xml 文件中依赖项如下：

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
		<!-- springboot druid -->
		<dependency>
			<groupId>com.alibaba</groupId>
			<artifactId>druid-spring-boot-starter</artifactId>
			<version>1.1.22</version>
		</dependency>
```

### 3.4 构建商品类和商品数据访问类

与 `spring-boot-hikari` 项目一致。

### 3.5 配置数据源信息

通过配置文件，设置数据源信息。由于我们不再使用默认数据源，所以此处需要指定数据源类型为 DruidDataSource 。

**实例：**

```java
# 指定数据源类型
spring.datasource.type=com.alibaba.druid.pool.DruidDataSource
# 配置数据库驱动
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
# 配置数据库url
spring.datasource.url=jdbc:mysql://127.0.0.1:3306/shop?useUnicode=true&characterEncoding=utf-8&serverTimezone=UTC
# 配置数据库用户名
spring.datasource.username=root
# 配置数据库密码
spring.datasource.password=Easy@0122
```

### 3.6 测试

测试类代码同 `spring-boot-hikari` 一致，运行测试类后，结果如下：

```java
use time:1428ms
com.alibaba.druid.spring.boot.autoconfigure.DruidDataSourceWrapper
```

## 4. 对比结果分析

其实只能得出一个结论，在某些场景下 Druid 的速度不比 Hikari 慢，甚至还略胜一筹。

当然我们只是对两种数据源的默认配置、单一线程情况进行了简单测试，大家感兴趣的话可以研究两种数据源的配置方式，然后通过多线程进行全面测试。

## 5. Druid 监控

看到这个结果，大家可能对本篇文章不满了，说了半天，也没看出 Druid 好在哪儿啊，为啥还费劲将默认的 Hikari 更换掉呢。

不要着急，我们仔细看下官方介绍：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2320bcj60wl0i0dii02)

 Druid 在阿里巴巴开源项目官网的描述

可以看到， Druid 是为监控而生，说明 Druid 最强大的功能实际上是监控，接下来我们就演示下如何实现 Druid 监控。

添加监控相关的配置类，需要注意的是我们设定了监控功能的账号和密码。

**实例：**

```java
/**
 * Druid配置
 */
@Configuration
public class DruidConfig {
	/**
	 * 注册servletRegistrationBean
	 */
	@Bean
	public ServletRegistrationBean servletRegistrationBean() {
		ServletRegistrationBean servletRegistrationBean = new ServletRegistrationBean(new StatViewServlet(),
				"/druid/*");
		servletRegistrationBean.addInitParameter("allow", "");
		// 账号密码
		servletRegistrationBean.addInitParameter("loginUsername", "imooc");
		servletRegistrationBean.addInitParameter("loginPassword", "123456");
		servletRegistrationBean.addInitParameter("resetEnable", "true");
		return servletRegistrationBean;
	}

	/**
	 * 注册filterRegistrationBean
	 */
	@Bean
	public FilterRegistrationBean filterRegistrationBean() {
		FilterRegistrationBean filterRegistrationBean = new FilterRegistrationBean(new WebStatFilter());
		// 添加过滤规则.
		filterRegistrationBean.addUrlPatterns("/*");
		filterRegistrationBean.addInitParameter("exclusions", "*.js,*.gif,*.jpg,*.png,*.css,*.ico,/druid/*");
		return filterRegistrationBean;
	}
}
```

此时打开网址 `http://127.0.0.1:8080/druid` 即可显示 Druid 登录页面：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo23anerj60ji07nt8r02)

 Druid 登录页面

我们使用指定的用户名 imooc 密码 123456 登录后，即可查看各类监控信息，内容还是非常全面的，此处就不再展开介绍了。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo23j1itj60ut0b8t9p02)

 Druid 监控页面

## 6. 小结

在实际研发与生产测试过程中，使用 Druid 的情况还是非常多的， Druid 非常稳定、性能也表现相当优异，更重要的是它提供了全面直观的监控手段，所以现阶段还是推荐大家使用 Druid 。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
