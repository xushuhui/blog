# Spring 代理模式解决事务

## 1. 前言

大家好，本小节，我们学习代理模式解决转账过程中产生的事务问题，如有中途而来的童鞋，请先移步上一小节学习下问题的场景。

## 2. 实战案例

### 2.1 实现思路介绍

1. 创建一个工具类，目的是用于管理数据库的事务，提供事务的开启，提交，回滚等操作；

2. 创建一个代理处理器类，目的是生成转账实现类的代理对象，对转账的业务方法提供增强，主要是在数据操作之前，和操作之后干点事，嘿嘿嘿；

3. 在 Spring 的配置文件中，通过 xml 文件的标签实例化管理事务的工具类和生成代理对象的处理器类。

### 2.2 代码实现

**1. 创建事务管理器类**

```java
package com.offcn.transaction;

/**
 * @Auther: wyan
 * @Date: 2020-05-26 21:20
 * @Description:
 */

import com.offcn.utils.ConnectionUtils;

/**
 * 和事务管理相关的工具类，它包含了，开启事务，提交事务，回滚事务和释放连接
 */
public class TransactionManager {
	//获取数据库连接的工具类
    private ConnectionUtils connectionUtils;

    public void setConnectionUtils(ConnectionUtils connectionUtils) {
        this.connectionUtils = connectionUtils;
    }

    /**
     * 开启事务
     */
    public  void beginTransaction(){
        try {
            connectionUtils.getThreadConnection().setAutoCommit(false);
        }catch (Exception e){
            e.printStackTrace();
        }
    }

    /**
     * 提交事务
     */
    public  void commit(){
        try {
            connectionUtils.getThreadConnection().commit();
        }catch (Exception e){
            e.printStackTrace();
        }
    }

    /**
     * 回滚事务
     */
    public  void rollback(){
        try {
            connectionUtils.getThreadConnection().rollback();
        }catch (Exception e){
            e.printStackTrace();
        }
    }


    /**
     * 释放连接
     */
    public  void release(){
        try {
            connectionUtils.getThreadConnection().close();//还回连接池中
            connectionUtils.removeConnection();
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}

```

代码解释：此工具类主要作用是对数据库连接实现事务的开启，提交以及回滚。至于何时开启事务，何时提交事务，何时回滚事务，那就根据业务场景需要调用该类的方法即可。

**2. 创建动态处理器**

```java
package com.offcn.utils;

import com.offcn.service.IAccountService;
import com.offcn.transaction.TransactionManager;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

/**
 * @Auther: wyan
 * @Date: 2020-05-26 21:08
 * @Description:
 */
public class TransactionProxyFactory {
	//被代理的业务类接口
    private IAccountService accountService;
	//提供事务管理的工具类
    private TransactionManager txManager;

    public void setTxManager(TransactionManager txManager) {
        this.txManager = txManager;
    }

    public final void setAccountService(IAccountService accountService) {
        this.accountService = accountService;
    }
    /**
     * 获取Service代理对象
     * @return
     */
    public IAccountService getAccountService() {
        return (IAccountService) Proxy.newProxyInstance(accountService.getClass().getClassLoader(),
                accountService.getClass().getInterfaces(),
                new InvocationHandler() {
                    /**
                     * 添加事务的支持
                     *
                     * @param proxy
                     * @param method
                     * @param args
                     * @return
                     * @throws Throwable
                     */
                    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {

                       //
                        Object rtValue = null;
                        try {
                            //1.开启事务
                            txManager.beginTransaction();
                            //2.执行操作
                            rtValue = method.invoke(accountService, args);
                            //3.提交事务
                            txManager.commit();
                            //4.返回结果
                            return rtValue;
                        } catch (Exception e) {
                            //5.回滚操作
                            txManager.rollback();
                            throw new RuntimeException(e);
                        } finally {
                            //6.释放连接
                            txManager.release();
                        }
                    }
                });

    }
}
```

**代码解释：**

此类的核心代码就是 `getAccountService` 方法，该方法返回代理业务类示例，而在代理对象的 invoke 方法内部，实现对原始被代理对象的增强。

方法的参数解释如下：

1. **proxy：** 该参数就是被代理的对象实例本身；
2. **method**: 该参数是被代理对象正在执行的方法对象；
3. **args：** 该参数是正在访问的方法参数对象。

在方法内部，`method.invoke()` 的方法调用，即表示被代理业务类的方法执行，我们调用 `txManager` 的开启事务方法。在 `method.invoke()` 方法执行之后，调用提交事务的方法。

一旦执行过程出现异常，在 `catch` 代码块中调用事务回滚的方法。这样就保证了事务的原子性，执行的任务，要么全部成功，要么全部失败。

最终在 `finally` 的代码块中，调用释放连接的方法。

**3. 配置文件的修改：**

添加事务管理的相关配置，完整配置文件如下：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans.xsd">

    <!-- 配置Service -->
    <bean id="accountService" class="com.offcn.service.impl.AccountServiceImpl">
        <!-- 注入dao -->
        <property name="accountDao" ref="accountDao"></property>
    </bean>
    <!--配置Dao对象-->
    <bean id="accountDao" class="com.offcn.dao.impl.AccountDaoImpl">
        <!-- 注入QueryRunner -->
        <property name="runner" ref="runner"></property>
        <!-- 注入ConnectionUtils -->
        <property name="connectionUtils" ref="connectionUtils"></property>
    </bean>
    <!--配置QueryRunner-->
    <bean id="runner" class="org.apache.commons.dbutils.QueryRunner" scope="prototype"></bean>
    <!-- 配置数据源 -->
    <bean id="dataSource" class="com.mchange.v2.c3p0.ComboPooledDataSource">
        <!--连接数据库的必备信息-->
        <property name="driverClass" value="com.mysql.jdbc.Driver"></property>
        <property name="jdbcUrl" value="jdbc:mysql://localhost:3306/transmoney"></property>
        <property name="user" value="root"></property>
        <property name="password" value="root"></property>
    </bean>
    <!-- 配置Connection的工具类 ConnectionUtils -->
    <bean id="connectionUtils" class="com.offcn.utils.ConnectionUtils">
        <!-- 注入数据源-->
        <property name="dataSource" ref="dataSource"></property>
    </bean>

    <!-- 配置事务管理器-->
    <bean id="txManager" class="com.offcn.transaction.TransactionManager">
        <!-- 注入ConnectionUtils -->
        <property name="connectionUtils" ref="connectionUtils"></property>
    </bean>
    <!--配置beanfactory-->
    <bean id="beanFactory" class="com.offcn.utils.TransactionProxyFactory">
        <!-- 注入service -->
        <property name="accountService" ref="accountService"></property>
        <!-- 注入事务管理器 -->
        <property name="txManager" ref="txManager"></property>
    </bean>
    <!--配置代理的service-->
    <bean id="proxyAccountService" factory-bean="beanFactory" factory-method="getAccountService"></bean>

</beans>
```

**4. 测试类代码**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3brsghj61580don5602)

**代码解释：**

本测试代码发生一个小变化，第 23 行的位置，多了一个注解 `@Qualifier` 。此注解的作用不知各位是否还记得，如果在 Spring 的容器中，出现多种同类型的 bean ，可以通过此注解指定引入的

实例，所以这里的 注解内的字符串 `proxyAccountService` 表示本 `IAccountService` 接口引入的实例为代理对象。那么为什么要引入代理对象呢？因为代理对象的方法内部已经做了增强逻辑，通过 TransactionManager 类实现对事务的开启，提交和回滚。

**5. 测试结果：**

为了测试效果更明显，我们先把数据库的数据还原为每人各 1000，如图：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3bztfwj607g04075b02)

执行代码后结果：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3cbt7mj61510f9wxw02)

当然还会继续报错，但是数据库呢？上次是一个账号减去了 100 块钱，另外一个账号却没有增加钱，这次我们来看看：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3d3927j607j040my702)

可以看到：账号的金钱依然是原样，这就说明事务的控制已经生效了，保证了数据的一致性。

## 3. 小结

本小节学习了代理模式实现对事务的控制，加深了代理模式的优点及作用：

1. **职责清晰：** 代理类与被代理类各司其职，互不干扰；
2. **高扩展性：** 代码耦合性低，可以更加方便对方法做增强；
3. **符合开闭原则：** 系统具有较好的灵活性和可扩展性。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
