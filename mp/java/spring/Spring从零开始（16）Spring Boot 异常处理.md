# Spring Boot 异常处理

## 1. 前言

程序中出现异常是普遍现象， Java 程序员想必早已习惯，根据控制台输出的异常信息，分析异常产生的原因，然后进行针对性处理的过程。

Spring Boot 项目中，数据持久层、服务层到控制器层都可能抛出异常。如果我们在各层都进行异常处理，程序代码会显得支离破碎，难以理解。

实际上，异常可以从内层向外层不断抛出，最后在控制器层进行统一处理。 Spring Boot 提供了全局性的异常处理机制，本节我们就分别演示下，默认情况、控制器返回视图、控制器返回 JSON 数据三种情况的异常处理方法。

## 2. Spring Boot 默认异常处理机制

Spring Boot 开发的 Web 项目具备默认的异常处理机制，无须编写异常处理相关代码，即可提供默认异常机制，下面具体演示下。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-exception-default ，生成项目后导入 Eclipse 开发环境。

### 2.2 引入项目依赖

引入 Web 项目依赖即可。

**实例：**

```java
		<!-- web项目依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
```

### 2.3 Spring Boot 默认异常处理

我们在启动项目， Spring Boot Web 项目默认启动端口为 8080 ，所以直接访问 `http://127.0.0.1:8080` ，显示如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo26t2t9j60kf04ugnr02)

Spring Boot 默认异常信息提示页面

如上图所示，Spring Boot 默认的异常处理机制生效，当出现异常时会自动转向 `/error` 路径。

## 3. 控制器返回视图时的异常处理

在使用模板引擎开发 Spring Boot Web 项目时，控制器会返回视图页面。我们使用 Thymeleaf 演示控制器返回视图时的异常处理方式，其他模板引擎处理方式也是相似的。

### 3.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-exception-controller，生成项目后导入 Eclipse 开发环境。

### 3.2 引入项目依赖

引入 Web 项目依赖、热部署依赖。此处使用 Thymeleaf 演示控制器返回视图时的异常处理方式，所以引入 Thymeleaf 依赖。

**实例：**

```java
		<!-- web项目依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- 热部署 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-devtools</artifactId>
		</dependency>
		<!-- ThymeLeaf依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-thymeleaf</artifactId>
		</dependency>
```

### 3.3 定义异常类

在异常处理之前，我们应该根据业务场景具体情况，定义一系列的异常类，习惯性的还会为各种异常分配错误码，如下图为支付宝开放平台的公共错误码信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo275vd2j60w30jl10z02)

支付宝开放平台错误码

本节我们为了演示，简单的定义 2 个异常类，包含错误码及错误提示信息。

**实例：**

```java
/**
 * 自定义异常
 */
public class BaseException extends Exception {
	/**
	 * 错误码
	 */
	private int code;
	/**
	 * 错误提示信息
	 */
	private String msg;

	public BaseException(int code, String msg) {
		super();
		this.code = code;
		this.msg = msg;
	}
	// 省略get set
}
```

**实例：**

```java
/**
 * 密码错误异常
 */
public class PasswordException extends BaseException {
	public PasswordException() {
		super(10001, "密码错误");
	}
}
```

**实例：**

```java
/**
 * 验证码错误异常
 */
public class VerificationCodeException extends BaseException {
	public VerificationCodeException() {
		super(10002, "验证码错误");
	}
}
```

### 3.4 控制器抛出异常

定义控制器 GoodsController ，然后使用注解 @Controller 标注该类，类中方法的返回值即为视图文件名。

在 GoodsController 类定义 4 个方法，分别用于正常访问、抛出密码错误异常、抛出验证码错误异常、抛出未自定义的异常，代码如下。

**实例：**

```java
/**
 * 商品控制器
 */
@Controller
public class GoodsController {
	/**
	 * 正常方法
	 */
	@RequestMapping("/goods")
	public String goods() {
		return "goods";// 跳转到resource/templates/goods.html页面
	}

	/**
	 * 抛出密码错误异常的方法
	 */
	@RequestMapping("/checkPassword")
	public String checkPassword() throws PasswordException {
		if (true) {
			throw new PasswordException();// 模拟抛出异常，便于测试
		}
		return "goods";
	}

	/**
	 * 抛出验证码错误异常的方法
	 */
	@RequestMapping("/checkVerification")
	public String checkVerification() throws VerificationCodeException {
		if (true) {
			throw new VerificationCodeException();// 模拟抛出异常，便于测试
		}
		return "goods";
	}

	/**
	 * 抛出未自定义的异常
	 */
	@RequestMapping("/other")
	public String other() throws Exception {
		int a = 1 / 0;// 模拟异常
		return "goods";
	}
}

```

### 3.5 开发基于 @ControllerAdvice 的全局异常类

@ControllerAdvice 注解标注的类可以处理 @Controller 标注的控制器类抛出的异常，然后进行统一处理。

**实例：**

```java
/**
 * 控制器异常处理类
 */
@ControllerAdvice(annotations = Controller.class) // 全局异常处理
public class ControllerExceptionHandler {
	@ExceptionHandler({ BaseException.class }) // 当发生BaseException类(及其子类)的异常时，进入该方法
	public ModelAndView baseExceptionHandler(BaseException e) {
		ModelAndView mv = new ModelAndView();
		mv.addObject("code", e.getCode());
		mv.addObject("message", e.getMessage());
		mv.setViewName("myerror");// 跳转到resource/templates/myerror.html页面
		return mv;
	}

	@ExceptionHandler({ Exception.class }) // 当发生Exception类的异常时，进入该方法
	public ModelAndView exceptionHandler(Exception e) {
		ModelAndView mv = new ModelAndView();
		mv.addObject("code", 99999);// 其他异常统一编码为99999
		mv.addObject("message", e.getMessage());
		mv.setViewName("myerror");// 跳转到resource/templates/myerror.html页面
		return mv;
	}
}

```

按照 ControllerExceptionHandler 类的处理逻辑，当发生 BaseException 类型的异常时，会跳转到 myerror.html 页面，并显示相应的错误码和错误信息；当发生其他类型的异常时，错误码为 99999 ，错误信息为相关的异常信息。

### 3.6 开发前端页面

在 resource/templates 下分别新建 goods.html 和 myerror.html 页面，作为正常访问及发生异常时跳转的视图页面。

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>goods.html页面</title>
</head>
<body>
	<div>商品信息页面</div>
</body>
</html>
```

**实例：**

```java
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>myerror.html页面</title>
</head>
<body>
	错误码：
	<span th:text="${code}"></span>
	错误信息：
	<span th:text="${message}"></span>
</body>
</html>
```

### 3.7 测试

启动项目，分别访问控制器中的 4 个方法，结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo27hg2wj60km046wf402)

访问正常方法 /goods

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo27tlq2j60km04575802)

访问抛出自定义异常的方法 /checkPassword

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo283qhkj60km047jsd02)

访问抛出自定义异常的方法 /checkVerification

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo28fw2sj60ko0480tm02)

访问抛出未自定义异常的方法 /other

可见，当控制器方法抛出异常时，会按照全局异常类设定的逻辑统一处理。

## 4. 控制器返回 JSON 数据时的异常处理

在控制器类上添加 @RestController 注解，控制器方法处理完毕后会返回 JSON 格式的数据。

此时，可以使用 @RestControllerAdvice 注解标注的类 ，来捕获 @RestController 标注的控制器抛出的异常。

### 4.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-exception-restcontroller，生成项目后导入 Eclipse 开发环境。

### 4.2 引入项目依赖

引入 Web 项目依赖、热部署依赖即可。

**实例：**

```java
		<!-- web项目依赖 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<!-- 热部署 -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-devtools</artifactId>
		</dependency>
```

### 4.3 定义异常类

还是使用上文中定义的异常类即可。

### 4.4 统一控制器返回数据格式

这时候，我们就需要思考一个问题了。前端请求后端控制器接口后，怎么区分后端接口是正常返回结果，还是发生了异常？

不论后端接口是正常执行，还是中间发生了异常，最好给前端返回统一的数据格式，便于前端统一分析处理。

OK，此时我们就可以封装后端接口返回的业务逻辑对象 ResultBo ，代码如下：

**实例：**

```java
/**
 * 后端接口返回的统一业务逻辑对象
 */
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
		if (ex instanceof BaseException) {
			this.code = ((BaseException) ex).getCode();
			this.msg = ex.getMessage();
		} else {
			this.code = 99999;// 其他未定义异常
			this.msg = ex.getMessage();
		}
	}
	// 省略 get set
}
```

### 4.5 控制器抛出异常

定义控制器 RestGoodsController ，并使用 @RestController 注解标注。在其中定义 4 个方法，然后分别用于正常访问、抛出密码错误异常、抛出验证码错误异常，以及抛出不属于自定义异常类的异常。

**实例：**

```java
/**
 * Rest商品控制器
 */
@RestController
public class RestGoodsController {
	/**
	 * 正常方法
	 */
	@RequestMapping("/goods")
	public ResultBo goods() {
		return new ResultBo<>(new ArrayList());// 正常情况下应该返回商品列表
	}

	/**
	 * 抛出密码错误异常的方法
	 */
	@RequestMapping("/checkPassword")
	public ResultBo checkPassword() throws PasswordException {
		if (true) {
			throw new PasswordException();// 模拟抛出异常，便于测试
		}
		return new ResultBo<>(true);// 正常情况下应该返回检查密码的结果true或false
	}

	/**
	 * 抛出验证码错误异常的方法
	 */
	@RequestMapping("/checkVerification")
	public ResultBo checkVerification() throws VerificationCodeException {
		if (true) {
			throw new VerificationCodeException();// 模拟抛出异常，便于测试
		}
		return new ResultBo<>(true);// 正常情况下应该返回检查验证码的结果true或false
	}

	/**
	 * 抛出未自定义的异常
	 */
	@RequestMapping("/other")
	public ResultBo other() throws Exception {
		int a = 1 / 0;// 模拟异常
		return new ResultBo<>(true);
	}
}
```

### 4.6 开发基于 @RestControllerAdvice 的全局异常类

@RestControllerAdvice 注解标注的类可以处理 RestController 控制器类抛出的异常，然后进行统一处理。

**实例：**

```java
/**
 * Rest控制器异常处理类
 */
@RestControllerAdvice(annotations = RestController.class) // 全局异常处理
public class RestControllerExceptionHandler {
	/**
	 * 处理BaseException类(及其子类)的异常
	 */
	@ExceptionHandler({ BaseException.class })
	public ResultBo baseExceptionHandler(BaseException e) {
		return new ResultBo(e);
	}

	/**
	 * 处理Exception类的异常
	 */
	@ExceptionHandler({ Exception.class })
	public ResultBo exceptionHandler(Exception e) {
		return new ResultBo(e);
	}
}
```

### 4.7 测试

启动项目，分别尝试访问控制器中的 4 个接口，结果如下。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo28q48ej60kl0473z502)

访问正常方法 /goods

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo290nczj60kj048q3s02)

访问抛出异常的方法 /checkPassword

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo29anq1j60kk042js602)

访问抛出异常的方法 /checkVerification

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo29pbvuj60kh047q3l02)

访问抛出异常的方法 /other

## 5. 小结

Spring Boot 的默认异常处理机制，实际上只能做到提醒开发者 “这个后端接口不存在” 的作用，作用非常有限。

所以我们在开发 Spring Boot 项目时，需要根据项目的实际情况，定义各类异常，并站在全局的角度统一处理异常。

不管项目有多少层次，所有异常都可以向外抛出，直到控制器层进行集中处理。

* 对于返回视图的控制器，如果没发生异常就跳转正常页面，如果发生异常可以自定义错误信息页面。
* 对于返回 JSON 数据的控制器，最好是定义统一的数据返回格式，便于前端根据返回信息进行正常或者异常情况的处理。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
