# Spring Boot AOP 应用场景

## 1. 前言

Spring 最重要的两个功能，就是依赖注入（DI）和面向切面编程 （AOP）。

AOP 为我们提供了处理问题的全局化视角，使用得当可以极大提高编程效率。

Spring Boot 中使用 AOP 与 Spring 中使用 AOP 几乎没有什么区别，只是建议尽量使用 Java 配置代替 XML 配置。

本节就来演示下 Spring Boot 中使用 AOP 的常见应用场景。

## 2. 构建项目

首先我们需要构建一个 Spring Boot 项目并引入 AOP 依赖，后续场景演示均是在这个项目上实现的。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-aop，生成项目后导入 Eclipse 开发环境。

### 2.2 引入项目依赖

我们引入 Web 项目依赖与 AOP 依赖。

**实例：**

```java
		<!-- Web项目依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- AOP -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-aop</artifactId>
		</dependency>
```

### 2.3 新建控制层、服务层、数据访问层

为了便于后续的演示，我们依次新建控制类、服务类、数据访问类，并将其放入对应的包中，项目结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2i5kt0j60kn052dgy02)

项目结构

各个类代码如下，注意此处仅仅是为了演示 AOP 的使用，并未真实访问数据库，而是直接返回了测试数据。

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
	 * 获取商品列表
	 */
	@GetMapping("/goods")
	public List getList() {
		return goodsService.getList();
	}
}
```

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
	 * 获取商品信息列表
	 */
	public List getList() {
		return goodsDao.getList();
	}
}
```

**实例：**

```java
/**
 * 商品数据库访问类
 */
@Repository // 标注数据访问类
public class GoodsDao {
	/**
	 * 查询商品列表
	 */
	public List getList() {
		return new ArrayList();
	}
}
```

## 3. 使用 AOP 记录日志

如果要记录对控制器接口的访问日志，可以定义一个切面，切入点即为控制器中的接口方法，然后通过前置通知来打印日志。

**实例：**

```java
/**
 * 日志切面
 */
@Component
@Aspect // 标注为切面
public class LogAspect {
	private Logger logger = LoggerFactory.getLogger(this.getClass());

	// 切入点表达式，表示切入点为控制器包中的所有方法
	@Pointcut("within(com.imooc.springbootaop.controller..*)")
	public void LogAspect() {
	}

	// 切入点之前执行
	@Before("LogAspect()")
	public void doBefore(JoinPoint joinPoint) {
		logger.info("访问时间：{}--访问接口:{}", new Date(), joinPoint.getSignature());
	}
}
```

启动项目后，访问控制器中的方法之前会先执行 doBefore 方法。控制台打印如下：

```java
2020-05-25 22:14:12.317  INFO 9992 --- [nio-8080-exec-2] com.imooc.springbootaop.LogAspect        :
访问时间：Mon May 25 22:14:12 CST 2020--访问接口:List com.imooc.springbootaop.controller.GoodsController.getList()
```

## 4. 使用 AOP 监控性能

在研发项目的性能测试阶段，或者项目部署后，我们会希望查看服务层方法执行的时间。以便精准的了解项目中哪些服务方法执行速度慢，后续可以针对性的进行性能优化。

此时我们就可以使用 AOP 的环绕通知，监控服务方法的执行时间。

**实例：**

```java
/**
 * 服务层方法切面
 */
@Component
@Aspect // 标注为切面
public class ServiceAspect {
	private Logger logger = LoggerFactory.getLogger(this.getClass());

	// 切入点表达式，表示切入点为服务层包中的所有方法
	@Pointcut("within(com.imooc.springbootaop.service..*)")
	public void ServiceAspect() {
	}

	@Around("ServiceAspect()") // 环绕通知
	public Object deAround(ProceedingJoinPoint joinPoint) throws Throwable {
		long startTime = System.currentTimeMillis();// 记录开始时间
		Object result = joinPoint.proceed();
		logger.info("服务层方法:{}--执行时间:{}毫秒", joinPoint.getSignature(), System.currentTimeMillis() - startTime);
		return result;
	}
}
```

当服务层方法被调用时，控制台输入日志如下：

```java
2020-05-25 22:25:56.830  INFO 4800 --- [nio-8080-exec-1] com.imooc.springbootaop.ServiceAspect    :
服务层方法:List com.imooc.springbootaop.service.GoodsService.getList()--执行时间:3毫秒
```

> **Tips**：正常情况下，用户查看页面或进行更新操作时，耗时超过 1.5 秒，就会感觉到明显的迟滞感。由于前后端交互也需要耗时，按正态分布的话，大部分交互耗时在 0.4 秒 左右。所以在我参与的项目中，会对耗时超过 1.1 秒的服务层方法进行跟踪分析，通过优化 SQL 语句、优化算法、添加缓存等方式缩短方法执行时间。上面的数值均为我个人的经验参考值，还要视乎具体的服务器、网络、应用场景来确定合理的监控临界值。

## 5. 使用 AOP 统一后端返回值格式

前后端分离的项目结构中，前端通过 Ajax 请求后端接口，此时最好使用统一的返回值格式供前端处理。此处就可以借助 AOP 来实现正常情况、异常情况返回值的格式统一。

### 5.1 定义返回值类

首先定义返回值类，它属于业务逻辑对象 （Bussiness Object），所以此处命名为 ResultBo ，代码如下：

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
	// 省略 get set
}
```

### 5.2 修改控制层返回值类型

对所有的控制层方法进行修改，保证返回值均通过 ResultBo 包装，另外我们再定义一个方法，模拟抛出异常的控制层方法。

**实例：**

```java
	/**
	 * 获取商品列表
	 */
	@GetMapping("/goods")
	public ResultBo getList() {
		return new ResultBo(goodsService.getList());
	}
	/**
	 * 模拟抛出异常的方法
	 */
	@GetMapping("/test")
	public ResultBo test() {
		int a = 1 / 0;
		return new ResultBo(goodsService.getList());
	}
```

### 5.3 定义切面处理异常返回值

正常控制层方法都返回 ResultBo 类型对象，然后我们需要定义切面，处理控制层抛出的异常。当发生异常时，同样返回 ResultBo 类型的对象，并且对象中包含异常信息。

**实例：**

```java
/**
 * 返回值切面
 */
@Component
@Aspect
public class ResultAspect {
	// 切入点表达式，表示切入点为返回类型ResultBo的所有方法
	@Pointcut("execution(public com.imooc.springbootaop.ResultBo *(..))")
	public void ResultAspect() {
	}

	// 环绕通知
	@Around("ResultAspect()")
	public Object deAround(ProceedingJoinPoint joinPoint) throws Throwable {
		try {
			return joinPoint.proceed();// 返回正常结果
		} catch (Exception ex) {
			return new ResultBo<>(ex);// 被切入的方法执行异常时，返回ResultBo
		}
	}
}
```

### 5.4 测试

启动项目，访问 `http://127.0.0.1:8080/goods` 返回数据如下：

**实例：**

```java
{"code":0,"msg":"操作成功","data":[]}
```

然后访问 `http://127.0.0.1:8080/test` ，返回数据如下：

**实例：**

```java
{"code":99999,"msg":"/ by zero","data":null}
```

这样，前端可以根据返回值的 code， 来判断后端是否正常响应。如果 code 为 0 ，则进行正常业务逻辑操作；如果 code 非 0 ，则可以弹窗显示 msg 提示信息。

## 6. 小结

AOP 之所以如此重要，在于它提供了解决问题的新视角。通过将业务逻辑抽象出切面，功能代码可以切入指定位置，从而消除重复的模板代码。

使用 AOP 有一种掌握全局的快感，发现业务逻辑中的切面颇有一番趣味，希望大家都能多多体会，编程且快乐着应该是我辈的追求。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
