# Spring 工程执行过程

## 1. 前言

**Spring 框架是如何工作的？**

本节目的在于帮助大家理解 Spring 框架底层干了什么事情。

在上一节中我们通过一个入门工程简单地体验了一把 Spring 的使用。

我们发现，通过构造一个 `ClassPathXmlApplicationContext` 对象，加载项目的 `applicationContext.xml` 文件，确实可以实例化对象。

**疑问导出**

而脑海中不禁有一个想法… Spring 如何初始化对象的实例的？我们又如何从容器中获取得到对象的实例的呢？

带着疑问… 开启本节的源码和原理之旅。

## 2. 容器初始化

回顾代码：

```java
public static void main(String[] args) {
    ApplicationContext context =
            new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
    UserService service = (UserService) context.getBean("userService");
    service.saveUser();
}

```

在上面的代码中可以得知 Spring 的容器是 `ApplicationContext`，那么它到底是什么东西呢？先跟我一起追踪一下它的角色。

**官方文档**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ob9qij613y08j7cd02)

**慕课解释**

简单翻译过来就是 `ApplicationContext` 是一个 接口，是 `BeanFactory` 这个接口的子接口，它扩展了 `BeanFactory` 这个接口，提供了额外附加的功能。

而 `BeanFactory` 是管理 bean 对象的容器的根接口，大家了解下就好，我们是针对它的子接口 `ClassPathXmlApplicationContext` 做的实例化，目的是加载项目中的 Spring 的配置文件，使 Spring 来管理我们定义的 bean 对象。

**疑问导出**

那么我们的问题是…`ClassPathXmlApplicationContext` 对象实例化之后，干了哪些事情呢？

### 2.1 容器初始化执行动作

**`applicationContext` 实例化执行代码逻辑** 。

我们追踪下源码，发现 `ClassPathXmlApplicationContext` 初始化的时候，它做了一系列的事情。源码如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2olwhbj60uq0drwk302)

**代码解释：**

1. 是初始化 `ClassPathXmlApplicationContext` 对象执行的有参构造；
2. 加载项目下的 xml 配置文件；
3. 调用 refresh 刷新容器的方法 bean 的实例化就在这个方法中。

**继续跟踪：**

### 2.2 容器初始化 bean 对象动作

**下面是从源码中粘贴的部分代码**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2ou8lsj60ld0phgy102)

**步骤阐述：**

对于我们而言 这些英文看起来很吃力… 放轻松大家，我们只关注对我们理解流程有用的代码：

1. 1 的位置：是准备刷新，那么 Spring 只是设置刷新的标记，加载了外部的 `properties` 属性文件；
2. 2 的位置：是准备 bean 工厂对象；
3. 3 的位置：这一步骤就加载了配置文件中的所有 bean 标签，但是并没有对他们进行实例化；
4. 4 的位置：完成此上下文的 bean 工厂的初始化，初始化所有剩余的单例 bean。（Spring 中默认加载的 bean 就是单例模式后面生命周期会讲）
5. 最后的位置：完成容器的刷新，也就是所有的 bean 初始化完成。

```java
		//这里粘贴一部分初始化代码的逻辑 帮助大家理解
		// Instantiate all remaining (non-lazy-init) singletons.
		beanFactory.preInstantiateSingletons();
		// Trigger initialization of all non-lazy singleton beans...
		//所有非懒加载的单例bean的触发器初始化。。。
		for (String beanName : beanNames) {
		  ...//省略循环的代码
		}

```

OK 上面就是加载配置文件后 Spring 框架做的所有事情，当然实际底层涉及的东西 更多，但是我们没有必要深究，毕竟我们是理解过程，不是追求实现。

**疑问导出：**

我们整理了 Spring 初始化 bean 对象的过程，那么如果容器中确实存在了 bean 的实例，我们是如何获取得到的呢？

## 3. 容器中获取对象的过程

还是先看下我们获取容器对象的代码：

```java
public static void main(String[] args) {
    ApplicationContext context =
            new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
    UserService service = (UserService) context.getBean("userService");
    service.saveUser();
}

```

**代码分析**：

`context.getBean` 的方法是通过 **bean** 标签里的 **id** 来从容器中获取，那么我们看下源码 ：

在父类 **`AbstractApplicationContext`** 中有对 getBean 方法的实现。

```java
	@Override
	public Object getBean(String name) throws BeansException {
		assertBeanFactoryActive();
		return getBeanFactory().getBean(name);
	}

```

**追踪父类方法**

最终通过我们层层追踪，我们在 `AbstractAutowireCapableBeanFactory` 中发现这样的一段代码：

```java
protected Object doCreateBean(final String beanName, final RootBeanDefinition mbd, final @Nullable Object[] args)
			throws BeanCreationException {
		//...
		//省略大量方法内部代码
        //...
		// Initialize the bean instance.
		Object exposedObject = bean;
		try {
            //给实例中的属性赋值
			populateBean(beanName, mbd, instanceWrapper);
            //真实实例化对象
			exposedObject = initializeBean(beanName, exposedObject, mbd);
		}
		//...
        //继续省略大量方法
        //...
		// Register bean as disposable.
		try {
            //将实例化后的对象放入容器中
			registerDisposableBeanIfNecessary(beanName, bean, mbd);
		}
		catch (BeanDefinitionValidationException ex) {
			throw new BeanCreationException(
					mbd.getResourceDescription(), beanName, "Invalid destruction signature", ex);
		}
		//返回实例化后的对象实例
		return exposedObject;
	}

```

上面源码中我们可以看到： 对象实例的获取好像是在获取的时候执行的 `doCreateBean`，那么之前记载的 `xml` 文件不是实例过了吗？稍微解释下：加载文件时候的实例化操作，其实是实例化了一个 Spring 框架提供的对象，作用是对于我们 bean 对象做描述，这里才是真实的实例化动作。我们再看看 `registerDisposableBeanIfNecessary` 这个方法做的是什么。

```java
public void registerDisposableBean(String beanName, DisposableBean bean) {
		synchronized (this.disposableBeans) {
			this.disposableBeans.put(beanName, bean);
		}
	}

```

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2p3xtij60r90ac7bm02)

**结论**

一切真相大白。它其实就是一个 map 集合 ，这个 map 集合的 key 就是我们定义的 bean 的 id 或者 bean 的 name ，那么值就是对象的实例。

## 4. 小结

本章节 带着大家梳理了一下 Spring 初始化 bean 和获取 bean 的流程：

1. Spring 框架通过 ResourceLoader 加载项目的 xml 配置文件；
2. 读取 xml 的配置信息 变成对象存储，但未实例化；
3. 通过 bean 工厂处理器对 bean 做实例化，存储到一个 map 集合中默认是单例；
4. 获取对象 通过 xml 文件中 bean 的 id 从 map 集合中通过 get (key) 获取。

罗马不是一天建成的 ，书山有路勤为径…

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
