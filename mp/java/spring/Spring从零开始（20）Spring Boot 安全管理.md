# Spring Boot 安全管理

## 1. 前言

安全管理是软件系统必不可少的的功能。根据经典的“墨菲定律”——凡是可能，总会发生。如果系统存在安全隐患，最终必然会出现问题。

本节就来演示下，如何使用 Spring Boot + Spring Security 开发前后端分离的权限管理功能。

## 2. Spring Security 用法简介

作为一个知名的安全管理框架， Spring Security 对安全管理功能的封装已经非常完整了。

我们在使用 Spring Security 时，只需要从配置文件或者数据库中，把用户、权限相关的信息取出来。然后通过配置类方法告诉 Spring Security ， Spring Security 就能自动实现认证、授权等安全管理操作了。

* 系统初始化时，告诉 Spring Security 访问路径所需要的对应权限。
* 登录时，告诉 Spring Security 真实用户名和密码。
* 登录成功时，告诉 Spring Security 当前用户具备的权限。
* 用户访问接口时，Spring Security 已经知道用户具备的权限，也知道访问路径需要的对应权限，所以自动判断能否访问。

## 3. 数据库模块实现

### 3.1 定义表结构

需要 4 张表：

* **用户表 user**：保存用户名、密码，及用户拥有的角色 id 。
* **角色表 role** ：保存角色 id 与角色名称。
* **角色权限表 roleapi**：保存角色拥有的权限信息。
* **权限表 api**：保存权限信息，在前后端分离的项目中，权限指的是控制器中的开放接口。

具体表结构如下，需要注意的是 api 表中的 path 字段表示接口的访问路径，另外所有的 id 都是自增主键。

![](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2fcocnj60lp04j0ug02)

数据库表结构

### 3.2 构造测试数据

执行如下 SQL 语句插入测试数据，下面的语句指定了 admin 用户可以访问 viewGoods 和 addGoods 接口，而 guest 用户只能访问 viewGoods 接口。

**实例：**

```java
-- 用户
INSERT INTO `user` VALUES (1, 'admin', '$2a$10$D0OvhHj2Lh92rNey1EFmM.OqltxhH1vZA8mDpxz7jEofDEqLRplQy', 1);
INSERT INTO `user` VALUES (2, 'guest', '$2a$10$D0OvhHj2Lh92rNey1EFmM.OqltxhH1vZA8mDpxz7jEofDEqLRplQy', 2);
-- 角色
INSERT INTO `role` VALUES (1, '管理员');
INSERT INTO `role` VALUES (2, '游客');
-- 角色权限
INSERT INTO `roleapi` VALUES (1, 1, 1);
INSERT INTO `roleapi` VALUES (2, 1, 2);
INSERT INTO `roleapi` VALUES (3, 2, 1);
-- 权限
INSERT INTO `api` VALUES (1, 'viewGoods');
INSERT INTO `api` VALUES (2, 'addGoods');
```

> **Tips**：用户密码是 123 加密后的值，大家了解即可，稍后再进行解释。

## 4. Spring Boot 后端实现

我们新建一个 Spring Boot 项目，并利用 Spring Security 实现安全管理功能。

### 4.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-security，生成项目后导入 Eclipse 开发环境。

### 4.2 引入项目依赖

我们引入 Web 项目依赖、安全管理依赖，由于要访问数据库所以引入 JDBC 和 MySQL 依赖。

**实例：**

```java
		<!-- Web项目依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- 安全管理依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-security</artifactId>
		</dependency>
		<!-- JDBC -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-jdbc</artifactId>
		</dependency>
		<!-- MySQL -->
		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
		</dependency>
```

### 4.3 定义数据对象

安全管理，肯定需要从数据库中读取用户信息，以便判断用户登录名、密码是否正确，所以需要定义用户数据对象。

**实例：**

```java
public class UserDo {
	private Long id;
	private String username;
	private String password;
	private String roleId;
	// 省略 get set
}
```

### 4.4 开发数据访问类

系统初始化时，告诉 Spring Security 访问路径所需要的对应权限，所以我们开发从数据库获取权限列表的方法。

**实例：**

```java
@Repository
public class ApiDao {
	@Autowired
	private JdbcTemplate jdbcTemplate;

	/**
	 * 获取所有api
	 */
	public List<String> getApiPaths() {
		String sql = "select path from api";
		return jdbcTemplate.queryForList(sql, String.class);
	}
}
```

登录时，告诉 Spring Security 真实用户名和密码。 登录成功时，告诉 Spring Security 当前用户具备的权限。

所以我们开发根据用户名获取用户信息和根据用户名获取其可访问的 api 列表方法。

**实例：**

```java
@Repository
public class UserDao {
	@Autowired
	private JdbcTemplate jdbcTemplate;
	/**
	 * 根据用户名获取用户信息
	 */
	public List<UserDo> getUsersByUsername(String username) {
		String sql = "select id, username, password from user where username = ?";
		return jdbcTemplate.query(sql, new String[] { username }, new BeanPropertyRowMapper<>(UserDo.class));
	}
	/**
	 * 根据用户名获取其可访问的api列表
	 */
	public List<String> getApisByUsername(String username) {
		String sql = "select path from user left join roleapi on user.roleId=roleapi.roleId left join api on roleapi.apiId=api.id where username = ?";
		return jdbcTemplate.queryForList(sql, new String[] { username }, String.class);
	}
}

```

### 4.5 开发服务类

开发 SecurityService 类，保存安全管理相关的业务方法。

**实例：**

```java
@Service
public class SecurityService {
	@Autowired
	private UserDao userDao;
	@Autowired
	private ApiDao apiDao;

	public List<UserDo> getUserByUsername(String username) {
		return userDao.getUsersByUsername(username);
	}

	public List<String> getApisByUsername(String username) {
		return userDao.getApisByUsername(username);
	}

	public List<String> getApiPaths() {
		return apiDao.getApiPaths();
	}
}
```

### 4.6 开发控制器类

开发控制器类，其中 notLogin 方法是用户未登录时调用的方法，其他方法与权限表中的 api 一一对应。

**实例：**

```java
@RestController
public class TestController {
	/**
	 * 未登录时调用该方法
	 */
	@RequestMapping("/notLogin")
	public ResultBo notLogin() {
		return new ResultBo(new Exception("未登录"));
	}

	/**
	 * 查看商品
	 */
	@RequestMapping("/viewGoods")
	public ResultBo viewGoods() {
		return new ResultBo<>("viewGoods is ok");
	}

	/**
	 * 添加商品
	 */
	@RequestMapping("/addGoods")
	public ResultBo addGoods() {
		return new ResultBo<>("addGoods is ok");
	}
}
```

由于是前后端分离的项目，为了便于前端统一处理，我们封装了返回数据业务逻辑对象 ResultBo 。

**实例：**

```java
public class ResultBo<T> {
	/**
	 * 错误码 0表示没有错误(异常) 其他数字代表具体错误码
	 */
	private int code;
	/**
	 * 后端返回消息
	 */
	private String msg;
	/**
	 * 后端返回的数据
	 */
	private T data;
	/**
	 * 无参数构造函数
	 */
	public ResultBo() {
		this.code = 0;
		this.msg = "操作成功";
	}
	/**
	 * 带数据data构造函数
	 */
	public ResultBo(T data) {
		this();
		this.data = data;
	}
	/**
	 * 存在异常的构造函数
	 */
	public ResultBo(Exception ex) {
		this.code = 99999;// 其他未定义异常
		this.msg = ex.getMessage();
	}
}
```

### 4.7 开发 Spring Security 配置类

现在，我们就需要将用户、权限等信息通过配置类告知 Spring Security 了。

#### 4.7.1 定义配置类

定义 Spring Security 配置类，通过注解 @EnableWebSecurity 开启安全管理功能。

**实例：**

```java
@Configuration
@EnableWebSecurity // 开启安全管理
public class SecurityConfig {
	@Autowired
	private SecurityService securityService;
}
```

#### 4.7.2 注册密码加密组件

Spring Security 提供了很多种密码加密组件，我们使用官方推荐的 BCryptPasswordEncoder ，直接注册为 Bean 即可。

我们之前在数据库中预定义的密码字符串即为 123 加密后的结果。 Spring Security 在验证密码时，会自动调用注册的加密组件，将用户输入的密码加密后再与数据库密码比对。

**实例：**

```java
	@Bean
	PasswordEncoder passwordEncoder() {
		return new BCryptPasswordEncoder();
	}
    public static void main(String[] args) {
		//输出 $2a$10$kLQpA8S1z0KdgR3Cr6jJJ.R.QsIT7nrCdAfsF4Of84ZBX2lvgtbE.
		System.out.println(new BCryptPasswordEncoder().encode("123"));
	}
```

#### 4.7.3 将用户密码及权限告知 Spring Security

通过注册 UserDetailsService 类型的组件，组件中设置用户密码及权限信息即可。

**实例：**

```java
	@Bean
	public UserDetailsService userDetailsService() {
		return username -> {
			List<UserDo> users = securityService.getUserByUsername(username);
			if (users == null || users.size() == 0) {
				throw new UsernameNotFoundException("用户名错误");
			}
			String password = users.get(0).getPassword();
			List<String> apis = securityService.getApisByUsername(username);
			// 将用户名username、密码password、对应权限列表apis放入组件
			return User.withUsername(username).password(password).authorities(apis.toArray(new String[apis.size()]))
					.build();
		};
	}
```

#### 4.7.4 设置访问路径需要的权限信息

同样，我们通过注册组件，将访问路径需要的权限信息告知 Spring Security 。

**实例：**

```java
	@Bean
	public WebSecurityConfigurerAdapter webSecurityConfigurerAdapter() {
		return new WebSecurityConfigurerAdapter() {
			@Override
			public void configure(HttpSecurity httpSecurity) throws Exception {
				// 开启跨域支持
				httpSecurity.cors();
				// 从数据库中获取权限列表
				List<String> paths = securityService.getApiPaths();
				for (String path : paths) {
					/* 对/xxx/**路径的访问，需要具备xxx权限
					例如访问 /addGoods，需要具备addGoods权限 */
					httpSecurity.authorizeRequests().antMatchers("/" + path + "/**").hasAuthority(path);
				}
				// 未登录时自动跳转/notLogin
				httpSecurity.authorizeRequests().and().formLogin().loginPage("/notLogin")
						// 登录处理路径、用户名、密码
						.loginProcessingUrl("/login").usernameParameter("username").passwordParameter("password")
						.permitAll()
						// 登录成功处理
						.successHandler(new AuthenticationSuccessHandler() {
							@Override
							public void onAuthenticationSuccess(HttpServletRequest httpServletRequest,
									HttpServletResponse httpServletResponse, Authentication authentication)
									throws IOException, ServletException {
								httpServletResponse.setContentType("application/json;charset=utf-8");
								ResultBo result = new ResultBo<>();
								ObjectMapper mapper = new ObjectMapper();
								PrintWriter out = httpServletResponse.getWriter();
								out.write(mapper.writeValueAsString(result));
								out.close();
							}
						})
						// 登录失败处理
						.failureHandler(new AuthenticationFailureHandler() {
							@Override
							public void onAuthenticationFailure(HttpServletRequest httpServletRequest,
									HttpServletResponse httpServletResponse, AuthenticationException e)
									throws IOException, ServletException {
								httpServletResponse.setContentType("application/json;charset=utf-8");
								ResultBo result = new ResultBo<>(new Exception("登录失败"));
								ObjectMapper mapper = new ObjectMapper();
								PrintWriter out = httpServletResponse.getWriter();
								out.write(mapper.writeValueAsString(result));
								out.flush();
								out.close();
							}
						});
				// 禁用csrf(跨站请求伪造)
				httpSecurity.authorizeRequests().and().csrf().disable();
			}
		};
	}
```

按上面的设计，当用户发起访问时：

* 未登录的访问会自动跳转到`/notLogin` 访问路径。
* 通过 `/login` 访问路径可以发起登录请求，用户名和密码参数名分别为 username 和 password 。
* 登录成功或失败会返回 ResultBo 序列化后的 JSON 字符串，包含登录成功或失败信息。
* 访问 `/xxx` 形式的路径，需要具备 `xxx` 权限。用户所具备的权限已经通过上面的 UserDetailsService 组件告知 Spring Security 了。

## 5. 测试

启动项目后，我们使用 PostMan 进行验证测试。

### 5.1 未登录测试

在未登录时，直接访问控制器方法，会自动跳转 `/notLogin` 访问路径，返回`未登录`提示信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2flqwgj60km09uq4902)

未登录测试

### 5.2 错误登录密码测试

调用登录接口，当密码不对时，返回`登录失败`提示信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2fvvjxj60ko0bi75u02)

错误登录密码测试

### 5.3 以 guest 用户登录

使用 guest 用户及正确命名登录，返回`操作成功`提示信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2g8et4j60km0bf0ub02)

以 guest 用户登录

### 5.4 guest 用户访问授权接口

按照数据库中定义的规则， guest 用户可以访问 viewGoods 接口方法。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2gjm17j60ko09sq4b02)

guest 用户访问授权接口

### 5.5 guest 用户访问未授权接口

按照数据库中定义的规则， guest 没有访问 addGoods 接口方法的权限。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2gs6acj60kn0amwge02)

guest 用户访问未授权接口

### 5.6 admin 用户登录及访问授权接口

按照数据库中定义的规则， admin 用户登录后可以访问 viewGoods 和 addGoods 两个接口方法。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2h5h3dj60kl0bmdhg02)

admin 用户登录

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2hhnnyj60kk09twft02)

admin 用户访问授权接口

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2hsv49j60ko09uabf02)

admin 用户访问授权接口

## 6. 小结

Spring Boot 整合 Spring Security ，实际上大部分工作都在安全管理配置类上。

我们通过安全管理配置类，将用户、密码及其对应的权限信息放入容器，同时将访问路径所需要的权限信息放入容器， Spring Security 就会按照`用户访问路径--判断所需权限--用户是否具备该权限--允许或拒绝访问`这样的逻辑实施权限管理了。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
