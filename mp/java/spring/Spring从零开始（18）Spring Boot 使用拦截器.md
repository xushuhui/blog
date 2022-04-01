# Spring Boot 使用拦截器

## 1. 前言

拦截器这个名词定义的非常形象，就像导弹要攻击目标的时候，可能会被先进的反导系统拦截，此处的反导系统就是一种拦截器。

我们开发的应用，对外暴露的是控制器中定义的 API 方法，我们可以在 API 方法的外围放置拦截器，所有对 API 的访问都可以通过拦截器进行过滤。

OK，那么这样的拦截有什么意义吗，其实已经很明显了，反导系统可以保护目标的安全并识别对目标的攻击行为。同理，拦截器可以跟踪对应用的访问行为，对合法访问行为予以放行，对非法访问行为予以拒绝。怎么样，是不是很牛，接下来咱们就在 Spring Boot 项目中具体实现下。

## 2. 跟踪访问行为

要想实现对访问的拦截，首先要能跟踪访问行为，我们在 Spring Boot 中引入拦截器来实现下。

### 2.1 使用 Spring Initializr 创建项目

Spring Boot 版本选择 2.2.5 ，Group 为 com.imooc ， Artifact 为 spring-boot-interceptor ，生成项目后导入 Eclipse 开发环境。

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

### 2.3 创建拦截器

创建的类实现 HandlerInterceptor 接口，即可成为拦截器类。

**实例：**

```java
/**
 * 自定义拦截器类
 */
public class MyInterceptor implements HandlerInterceptor {// 实现HandlerInterceptor接口
	/**
	 * 访问控制器方法前执行
	 */
	@Override
	public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
			throws Exception {
		System.out.println(new Date() + "--preHandle:" + request.getRequestURL());
		return true;
	}

	/**
	 * 访问控制器方法后执行
	 */
	@Override
	public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
			ModelAndView modelAndView) throws Exception {
		System.out.println(new Date() + "--postHandle:" + request.getRequestURL());
	}

	/**
	 * postHandle方法执行完成后执行，一般用于释放资源
	 */
	@Override
	public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex)
			throws Exception {
		System.out.println(new Date() + "--afterCompletion:" + request.getRequestURL());
	}
}
```

在上面的实例中，我们定义了一个拦截器类 MyInterceptor ，通过实现 HandlerInterceptor 接口，该类具备了拦截器的功能。

MyInterceptor 中的方法执行顺序为 preHandle – Controller 方法 – postHandle – afterCompletion ，所以拦截器实际上可以对 Controller 方法执行前后进行拦截监控。

最后还有一个非常重要的注意点， preHandle 需要返回布尔类型的值。 preHandle 返回 true 时，对控制器方法的请求才能到达控制器，继而到达 postHandle 和 afterCompletion 方法；如果 preHandle 返回 false ，后面的方法都不会执行。

### 2.4 配置拦截器

上一步我们开发了配置器类，如果想让配置器生效，还需要通过配置类进行相应配置。

**实例：**

```java
/**
 * Web配置类
 */
@Configuration
public class WebConfig implements WebMvcConfigurer {
	/**
	 * 添加Web项目的拦截器
	 */
	@Override
	public void addInterceptors(InterceptorRegistry registry) {
		// 对所有访问路径，都通过MyInterceptor类型的拦截器进行拦截
		registry.addInterceptor(new MyInterceptor()).addPathPatterns("/**");
	}
}
```

### 2.5 创建控制器

我们建立一个简单的控制器，实现登录方法，以便检验拦截器的效果。

**实例：**

```java
/**
* 登录控制器
*/
@RestController
public class LoginController {
   @RequestMapping("/login")
   public boolean login(String username, String password) {
   	System.out.println(new Date() + " 某用户尝试登录，用户名:" + username + " 密码：" + password);
   	return true;
   }
}
```

### 2.6 跟踪访问行为

运行启动类，访问 `http://127.0.0.1:8080/login?username=imooc&password=123`，控制台输出如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2abt33j60kk03htcb02)

控制台输出内容

可见我们已经完整的跟踪了一次对 `http://127.0.0.1:8080/login` 接口的访问。

## 3. 实现访问控制

区分合法、非法访问，最常见的就是根据用户的登录状态、角色判断。接下来我们就演示下，对未登录用户非法访问请求的拦截。

### 3.1 修改控制器方法

修改登录方法，当用户输入的用户名和密码正确时，通过 Session 记录登录人信息。

然后开发获取登录人员信息方法，返回 Session 中记录的登录人信息。

**实例：**

```java
/**
 * 登录控制器
 */
@RestController
public class LoginController {
	/**
	 * 登录方法
	 */
	@RequestMapping("/login")
	public boolean login(HttpServletRequest request, String username, String password) {
		if ("imooc".equals(username) && "123".equals(password)) {
			// 登录成功，则添加Session并存储登录用户名
			request.getSession().setAttribute("LOGIN_NAME", username);
			return true;
		}
		return false;
	}

	/**
	 * 获取登录人员信息
	 */
	@RequestMapping("/info")
	public String info(HttpServletRequest request) {
		return "您就是传说中的：" + request.getSession().getAttribute("LOGIN_NAME");
	}
}
```

### 3.2 修改拦截器方法

由于用户在登录之前还没有设置 Session ，所以登录方法不应该拦截，可以让用户自由请求。但是只有登录成功后的用户，也就是说具备 Session 的用户才能访问 info 方法。

**实例：**

```java
/**
 * 自定义拦截器类
 */
public class MyInterceptor implements HandlerInterceptor {// 实现HandlerInterceptor接口
	/**
	 * 访问控制器方法前执行
	 */
	@Override
	public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
			throws Exception {
		if (request.getRequestURI().contains("/login") == true) {// 登录方法直接放行
			return true;
		} else {// 其他方法需要先检验是否存在Session
			if (request.getSession().getAttribute("LOGIN_NAME") == null) {//未登录的不允许访问
				return false;
			} else {
				return true;
			}
		}
	}
}
```

### 3.3 测试

首先直接请求 `http://127.0.0.1:8080/info` ，由于此时未登录，所以请求被拦截，网页输出如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2aoq4cj60kl032q3h02)

访问被拦截

如果先请求登录方法 `http://127.0.0.1:8080/login?username=imooc&password=123` ，然后访问 `http://127.0.0.1:8080/info` ，则网页输出：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2b2rp8j60kn034mxs02)

登录成功后，访问正常通过拦截器

## 4. 小结

Spring Boot 的拦截器能够管理对控制器方法的访问请求，通过使用拦截器，可以实现访问控制，加强项目的安全性。

当然对于更加复杂的安全管理续期， Spring Boot 也可以快速的整合 Spring Security 或 Shiro ，以构建企业级的安全管理体系，在后续章节再进一步介绍吧。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
