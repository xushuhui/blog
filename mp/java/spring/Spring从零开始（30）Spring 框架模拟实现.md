# Spring 框架模拟实现

## 1. 前言

通过几个章节的学习，大家对于 Spring 已经有了初步的认知，我们通过案例练习，或者源码追踪，可以粗略的看到 Spring 框架初始化 bean 对象的过程，那么这个章节，我们模拟 Spring 框架的思路，来写一个类似 Spring 加载对象的案例，加深大家的印象。

## 2. 案例实现思路

### 2.1 步骤介绍

**思路分析：**

我们通过写过的案例可以知道：

1. Spring 框架的容器 是一个接口 `ApplicationContext` 和接口的实现类 `ClassPathXmlApplicationContext` 来初始化的；
2. 在初始化容器对象的时候需要传递 xml 配置文件的位置；
3. xml 的配置文件中主要是通过 bean 标签可以对 Java 的类进行描述：类的路径 类的标识 类的构造参数等等；
4. 容器初始化以后需要解析 xml 配置文件的各个 bean 标签；
5. 实例化的对象如果有参数或者构造方法，那么也需要给参数赋值；

**开发准备**：

为了方便理解测试 ，我们来自定义容器的接口和实现类。

名称改为 `SpringContext` 和 `XmlSpringContext` 区别于框架的接口和实现类。

接口定义方法 getBean 用于获取容器内的示例，实现类定义有参构造用于接受初始化时候的配置文件路径。

**接口代码如下**：

```java
public interface SpringContext {
    public Object getBean(String beanName);
}
```

**实现类代码如下**:

```java
public class XmlSpringContext  implements SpringContext  {

    Map<String,Object> map = new HashMap<String,Object>();

    public XmlSpringContext (String filename){

    }

    public Object getBean(String beanName){
    	return map.get(beanName);
    }

}
```

**代码解释**：

1. map 用于存储实例化的 bean 对象 ；
2. 有参构造方法逻辑暂时为空，下面会做实现，加载文件实例化对象在方法内部；
3. getBean 的方法用于通过 key 获取 map 中存储的实例。

为了测试对象的实例化，我们自定义 `UserService` 和 `UserServiceImpl` 作为测试的接口对象和实现类。

**接口代码如下**：

```java
public interface UserService {

	public void deleteById(Integer id);

}
```

**接口的实现类代码如下**：

```java
public class UserServiceImpl implements UserService {
	//持久层的dao属性
    private UserDao userDao;

    public UserDao getUserDao() {
        return userDao;
    }

    public void setUserDao(UserDao userDao) {
        this.userDao = userDao;
    }
	//实现接口的方法
    public void deleteById(Integer id) {

        System.out.println("删除的方法执行");
    }

}
```

代码解释：dao 的属性其实是为了模拟属性赋值，后面依赖注入章节会详细讲解。

**自定义一个 xml 文件 作为模拟框架的配置文件** ：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans>
     <bean name="userDao" class="com.wyan.dao.UserDaoImpl"></bean>
     <bean name="userService" class="com.wyan.service.UserServiceImpl">
         <property name="userDao" ref="userDao"></property>
     </bean>
</beans>
```

代码解释：userDao 的 bean 需要实例化 是因为 service 用到了它的引用，所以这里多个属性 property。

**编写测试类加载文件测试**：

```java
public class TestSpring {

	@Test
	public void test() {
		//初始化容器（读取配置文件 构建工厂）
		SpringContext context =
                new XmlSpringContext("applicationContext.xml");
        UserServiceImpl userService = (UserServiceImpl) context.getBean("userService");
        userService.deleteById(1);
        System.out.println(userService.getUserDao());
	}

}
```

代码解释：这里的目的只是测试能否获取对象调用方法，如果控制台打印证明案例成功

### 2.2 容器对象的实现类构造函数具体代码

**思路分析**：

1. 读取初始化时候传递的文件路径；

2. 通过 SAXReader 解析 xml 文件的节点得到 beans 节点下对应多个 bean 节点集合；

3. 每一个 bean 表示一个对象，都需要被初始化，所以需要循环遍历集合；

4. 在循环遍历的过程中获取 id 属性和 class 属性，id 属性作为存入 map 的 key，class 属性用于反射实例化对象，并存储 map 的 value；

5. 继续解析子节点，如果有参数，反射获取 method 执行参数赋值。

**完整代码**：

```java
public XmlSpringContext(String filename){

		// xml文件的解析器
		SAXReader  sr = new SAXReader();
		try {
			//构建一个直接通向我们配置文件路径 的输入流
			InputStream inputStream = this.getClass().getClassLoader().getResourceAsStream(filename);
			//文档模型对象
			Document doc = sr.read(inputStream);
			//获取根标签
			Element root = doc.getRootElement();
			//获取当前根标签的子标签
			List<Element>  beans = root.elements("bean");
			for(Element bean:beans){
				String key = bean.attributeValue("name");
				String value = bean.attributeValue("class");
				Class<?> myclass = Class.forName(value);
				//当前对象
				Object obj = myclass.newInstance();
				map.put(key, obj);
				List<Element> elements = bean.elements("property");
				if(elements.size()>0){
					for(Element pro: elements){
						 String av = pro.attributeValue("name");//dao--->setDao
						 //方法名
						 String methodName="set"+(av.charAt(0)+"").toUpperCase()+av.substring(1,av.length());
						 //方法参数
						 String refvalue = pro.attributeValue("ref");
						 Object refobj = map.get(refvalue);
						 //根据方法名称获取方法对象Method
						 Method method = myclass.getMethod(methodName,refobj.getClass().getInterfaces()[0]);
						 method.invoke(obj, refobj);
					}
				}
			}

		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}
```

**测试结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2simtkj60xu0k8k1b02)

## 3. 小结

本章节带着大家模拟一下 Spirng 加载文件的过程和实例化对象的过程，当然这个过程只是模拟 Spring 的框架的思路，而并不是真正的 Spring 框架源码，实际源码远比这个要复杂的多，

那么通过本章节我们收获哪些知识呢？

1. Spring 容器类的使用
2.  xml 配置文件的作用
3. 反射技术的应用

我不相信不劳而获，如果有谁告诉你，他可以做到 XX 速成，请远离他，他连基本的诚信都没有…

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
