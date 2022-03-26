# Spring DI（依赖注入）之多种类型属性

## 1. 前言

上一节，我们演示了如何使用 xml 文件配置，实现属性的依赖注入。但是，注入的依赖类型只是作为演示使用的两种，

而一个类中的属性，可能会多种多样。那么，xml 配置文件如何实现其余类型的属性注入呢？

我们进入本节的学习内容。

## 2. 多种类型依赖注入

### 2.1 属性类型分类

* 基本类型包装类，比如 Integer、Double、Boolean；
* 字符串类型，比如 String；
* 类类型，比如 其余定义的 java 类；
* 集合类型，比如 map、set、list 。

对于基本类型和字符串类型，在 xml 的配置文件中，通过 value 属性即可以复制，我们上个案例已经测试过，这里不做赘述，主要演示集合类型的属性注入测试。

### 2.2 工程搭建：

搭建工程，引入依赖，配置文件步骤省略，参考上一个章节的工程自行实现

**1. 编写一个 java 类，属性为多种类型的集合**

代码如下：

```java
public class User {

private Integer id;
private String name;
private Object [] array;
private List    list;
private Map  map;
//省略get和set方法

}
```

可以看到上面在同一个类中，我们定义了多个属性， array 数组、list 和 map 集合。

**2. 编写配置文件 属性注入数组的依赖属性**

配置文件如下：

```java
	<!-- 数组的属性注入 -->
	<bean id="user" class="com.offcn.entity.User">
		<property name="array">
			<array>
				<value>tom</value>
				<value>jerry</value>
			</array>
		</property>
	</bean>
```

**配置解释：**

在上面的配置文件中：

* property 中的 name 是 Java 类中数组的属性名称，用于 set 方法提供注入；
* array 标签是固定的，不能变化，表示属性是一个数组，所以加在了 property 的属性内部；
* value 表示数组中的值，因为数组可以存储多个值，所以每一个数组的值，通过一个 value 标签声明。

**测试结果**：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2yubn6j61500h17hs02)

**代码解释**

可以看到： 我们得到了 user 类中的 array 数组中在 xml 文件中配置的数组值，定义好的 tom 和 jerry。那么数组的依赖注入完成。

**3. 集合的属性注入**

改造 xml 的配置文件，实现 list 集合的属性配置。

```java
	<!-- 集合的属性注入 -->
	<bean id="user" class="com.offcn.entity.User">
		<property name="list">
			<list>
				<value>笑傲江湖</value>
				<value>侠客行</value>
				<value>连城诀</value>
			</list>
		</property>
	</bean>
```

**配置解释：**

在上面的配置文件中：

* property 中的 name 是 java 类中数组的属性名称 用于 set 方法提供注入；
* list 标签是固定的，不能变化，表示属性是一个 list 集合 ，所以加在了 property 的属性内部；
* value 表示集合中的值，因为集合可以存储多个值，所以每一个集合中的值，通过一个 value 标签声明。

测试结果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2zd8dqj611v0f5ali02)

同理：测试结果可以看到，list 集合中的三个数据都打印出来，那么 list 集合的 xml 依赖注入也搞定。

**4.map 集合的注入实现**

继续更改 xml 文件的配置 ，如下：

```java
	<bean id="user" class="com.offcn.entity.User">
		<property name="map">
			<map>
				<entry key="小亮" value="小路"></entry>
				<entry key="文同学" value="伊利姐"></entry>
			</map>
		</property>
	</bean>
```

**配置解释：**

在上面的配置文件中：

* property 中的 name 是 Java 类中 map 的属性名称，用于 set 方法提供注入；
* map 标签是固定的，不能变化，表示属性是一个 map 集合 ，所以加在了 property 的属性内部；
* entry 标签固定表示 map 中的一对键值对，map 也可以存储多对，所以 key 表示键值对的键，value 表示键值对的值。

**测试结果如下：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2zpms9j614q0f5gxy02)

map 集合注入的效果我们也看到了…当然还有很多种类型，同理操作即可。

## 3. 小结

本节重点依赖注入的 xml 实现多种属性的注入。其实对于本节而言，重点还是理解 xml 文件依赖注入的方式，是通过 property 标签，搭配各个标签节点实现。

至于开发中使用方式，其实 xml 文件的配置未免显得麻烦了一点，所以下一节我们会讲解通过注解方式进行依赖注入。

才须学也。非学无以广才，非志无以成学。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
