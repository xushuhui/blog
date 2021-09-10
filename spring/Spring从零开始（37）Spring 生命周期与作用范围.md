# Spring 生命周期与作用范围

## 1. 前言

​ 上一节，我们多学习了一种初始化 Spring 容器的方式，那么不管是何种初始化容器的方式，目的都是对容器中的 bean 实例做管理的。

本节我们就学习 Spring 的容器如何管理对象的实例的。主要在于两个方向：

1. 对象的生命周期
2. 对象的作用范围

## 2. 对象的生命周期

### 2.1. 生命周期的概念

​ 生命周期，通俗的理解就是从出生到死亡的过程，那么对于对象而言，就是实例在 Spring 容器中创建到销毁的过程。

### 2.2 生命周期流程概要

简单地来说，一个 Bean 的生命周期分为四个阶段：

(1) 实例化（Instantiation）

(2) 属性设置（populate）

(3) 初始化（Initialization）

(4) 销毁（Destruction）

流程图如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo35bcgnj606k0fpgmi02)

Spring 中 bean 的实例化过程：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo35m6q7j60j90aedi702)

Bean 实例生命周期的执行过程如下：

* Spring 对 bean 进行实例化，默认 bean 是单例；

* Spring 对 bean 进行依赖注入；

* 如果 bean 实现了 `BeanNameAware` 接口，Spring 将 bean 的 id 传给 `setBeanName()` 方法；

* 如果 bean 实现了 `BeanFactoryAware` 接口，Spring 将调用 `setBeanFactory` 方法，将 BeanFactory 实例传进来；

* 如果 bean 实现了 `ApplicationContextAware` 接口，它的 `setApplicationContext()` 方法将被调用，将应用上下文的引用传入到 bean 中；

* 如果 bean 实现了 `BeanPostProcessor` 接口，它的 `postProcessBeforeInitialization` 方法将被调用；

* 如果 bean 实现了 `InitializingBean` 接口，Spring 将调用它的 `afterPropertiesSet` 接口方法，类似地如果 bean 使用了 init-method 属性声明了初始化方法，该方法也会被调用；

* 如果 bean 实现了 `BeanPostProcessor` 接口，它的 `postProcessAfterInitialization` 接口方法将被调用；

* 此时 bean 已经准备就绪，可以被应用程序使用了，他们将一直驻留在应用上下文中，直到该应用上下文被销毁；

* 若 bean 实现了 `DisposableBean` 接口，Spring 将调用它的 `distroy()` 接口方法。同样地，如果 bean 使用了 destroy-method 属性声明了销毁方法，则该方法被调用；

## 3. 对象的作用范围

### 3.1 介绍

作用范围四个字很好理解，其实就是对象起作用的范围在哪，那么到底有哪些作用范围呢？

### 3.2 作用范围举例

singleton — 单例模式

prototype — 原型模式

request — 请求范围

session — 会话范围

application — 应用范围

### 3.3 作用范围详解

**单例模式**

单例大家都不陌生，Java 基础我们学过如何创建单例，而在 Spring 的容器中，bean 的作用范围如果是 singleton，那么表示容器中仅管理一个共享实例，该单个实例存储在 Spring 容器的缓存 map 集合中，所有的对象对于该实例的引用，都是从缓存的 map 中返回该对象的实例。

结构图如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo35yzybj60lh08y0wa02)

图片解释：可以看到左侧的三个，是通过 bean 标签实例化的类，而在每个类中都有个 property 引入 accountDao 的依赖，右侧表示在内存中仅存在一个 dao 的实例，

这也是默认的实例模式 —— 单例模式。

**原型模式**

原型模式 听起来比较模糊，解释的意思就是 —— 如果 bean 的作用范围如果是 prototype，那么表示 Spring 容器针对当前对象获取实例的时候，每次都会重新 new 一个对象，返回给调用者。

结构图如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo36d5bij60lj0adjvm02)

图片解释：可以看到左侧的三个，是通过 bean 标签实例化的类，在每个类中都有个 property 引入 accountDao 的依赖。

右侧存在一个 dao 的 bean 标签配置。而在配置的属性中，有一个 scope =prototype ，即是原型模式的作用范围。 每个类中已引入的依赖，Spring 的容器都是通过 class 重新 new 一个实例分别返回给调用者。

**其余模式**

`request`，`session`，`application`，等等作用范围工作使用中并不常见，这里不做描述。大家可自行查询资料。

## 4. 小结

Spring 的容器中初始化的实例，可以存在多种作用范围，而最常用的就是 默认的单例模式，至于 prototype 原型也有使用场景。

那么到底是单例模式，还是原型模式都是根据实际需求来决定，如果使用的实例中存在多种共享属性，又要求数据状态不被改变，

那么就必须是原型 prototype 模式，否则使用单例更好。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
