# Spring 的事务管理器

## 1. 前言

各位同学大家好，又轮到我的时间了。

在上一小节咱们已经可以使用 Spring 集成了 Jdbc Template，并且实现了数据的基本操作基本使用，那么本小节我们在操作数据的基础之上，讲解一下 Spring 中的事务处理。

其实我们在前面的 AOP 相关的小节中，已经实现了对事务的控制。只不过呢，之前的事务控制是由我们手动创建的类来管理事务。嘿嘿嘿，代码略显简单，功能稍许单一。

而今天呢，我们使用 Spring 框架提供的事务管理器，并通过 AOP 配置切入点来管理事务。那么它能否给我们带来一些惊喜呢？来吧，大家。一起进入今天的课程…

**课程回顾**：

首先我们回顾一下事务控制的实现要求：

1. 提供一个类，作为切面用于处理事务的开启，提交和回滚。

2. 通过 xml 文件或者注解来配置 AOP，表述切入点和使用的切面类。

本小节带着大家分别使用 xml 文件的方式，和注解的方式实现 Spring 框架对于事务的控制。

## 2. 实例演示

### 2.1 思路介绍

**Spring 的事务管理器类**

Spring 框架本身已经充分考虑了对事物的支持，所以我们完全不必像之前一样自定义类来实现对事物的控制。Spring 已经抽象了一整套的事务机制，而作为开发人员根本不必了解底层的事务 API,

一样可以通过代码管理数据库的事务。顶层的事务管理器抽象就是 **PlatformTransactionManager**, 它为事务管理封装了一组独立于技术的方法。

而本示例就采用 Spring 提供的管理器实现类，来替换掉之前我们自己编写的事务控制工具类。

### 2.2 工程搭建

**1. 创建工程**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3fxb4xj60jh0frabc02)

**2. 引入依赖**

```java
<dependencies>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-context</artifactId>
            <version>5.0.2.RELEASE</version>
        </dependency>
		<!-- Spring jdbc 使用的依赖-->
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-jdbc</artifactId>
            <version>5.0.2.RELEASE</version>
        </dependency>

        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-tx</artifactId>
            <version>5.0.2.RELEASE</version>
        </dependency>

        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>5.1.6</version>
        </dependency>
     	<dependency>
            <groupId>org.aspectj</groupId>
            <artifactId>aspectjweaver</artifactId>
            <version>1.8.7</version>
   		 </dependency>
    </dependencies>
```

**3. 准备代码**

实体类代码

```java
/**
 * 账户的实体类
 */
public class Account implements Serializable {
    //数据id
    private Integer id;
    //账号编码
    private String accountNum;
    //账号金额
    private Float money;
}

```

持久层接口代码

```java
/**
 * 账户的持久层接口
 */
public interface IAccountDao {

    /**
     * 根据Id查询账户
     * @param accountId
     * @return
     */
    Account findAccountById(Integer accountId);

    /**
     * 保存账户
     * @param account
     */
    void saveAccount(Account account);

    /**
     * 更新账户
     * @param account
     */
    void updateAccount(Account account);


}
```

持久层实现类代码

```java
/**
 * 账户的持久层实现类
 */
@Repository
public class AccountDaoImpl implements IAccountDao {
    //jdbc模板类属性
    @Autowired
    private JdbcTemplate jdbcTemplate;

    //根据id查找
    public Account findAccountById(Integer accountId) {
        List<Account> accounts = jdbcTemplate.query("select * from account where id = ?",new BeanPropertyRowMapper<Account>(Account.class),accountId);
        return accounts.isEmpty()?null:accounts.get(0);
    }

    public void saveAccount(Account account) {
        jdbcTemplate.update("insert into account  values(?,?,?)",
                account.getId(),account.getAccountNum(),account.getMoney());
    }

    public void updateAccount(Account account) {
        jdbcTemplate.update("update account set accountnum=?,money=? where id=?",account.getAccountNum(),account.getMoney(),account.getId());
    }
}
```

业务层接口代码

```java
/**
 * @Auther: wyan
 */
public interface UserService {

    /**
     * 账户转账
     * @param fromId toId
     */
    public void transMoney(Integer fromId, Integer toId, Integer money);

}

```

业务层实现类代码

```java
/**
 * @Auther: wyan
 * @Description:
 */
@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private IAccountDao accountDao;

    public void transMoney(Integer fromId, Integer toId, Integer money) {
        Account fromAccount = accountDao.findAccountById(fromId);
        Account toAccount = accountDao.findAccountById(toId);
        //原始账号减钱
        fromAccount.setMoney(fromAccount.getMoney()-money);
        accountDao.updateAccount(fromAccount);
        //抛出异常
        int i=1/0;
        //转账账号加钱
        toAccount.setMoney(toAccount.getMoney()+money);
        accountDao.updateAccount(toAccount);
    }
}
```

> **Tips：** 我们再给原始账号减掉钱后，执行保存。然后在这里会出现个异常，就是为了测试事务的特性，所以手动加了个除 0 的代码。

**4. 配置文件**

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:aop="http://www.springframework.org/schema/aop"
       xmlns:tx="http://www.springframework.org/schema/tx"
       xmlns:context="http://www.springframework.org/schema/context"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans.xsd
        http://www.springframework.org/schema/aop
        http://www.springframework.org/schema/aop/spring-aop.xsd
        http://www.springframework.org/schema/tx
        http://www.springframework.org/schema/tx/spring-tx.xsd
        http://www.springframework.org/schema/context http://www.springframework.org/schema/context/spring-context.xsd">

    <!--配置JdbcTemplate-->
    <bean id="jdbcTemplate" class="org.springframework.jdbc.core.JdbcTemplate">
        <property name="dataSource" ref="dataSource"></property>
    </bean>

    <!-- 配置数据源-->
    <bean id="dataSource" class="org.springframework.jdbc.datasource.DriverManagerDataSource">
        <property name="driverClassName" value="com.mysql.jdbc.Driver"></property>
        <property name="url" value="jdbc:mysql:///transmoney"></property>
        <property name="username" value="root"></property>
        <property name="password" value="root"></property>
    </bean>
    <!--包路径扫描-->
    <context:component-scan base-package="com.offcn"></context:component-scan>
    <!--事务管理器-->
    <bean id="transactionManager" class="org.springframework.jdbc.datasource.DataSourceTransactionManager">
        <property name="dataSource" ref="dataSource"></property>
    </bean>
    <!--事务的通知-->
    <tx:advice id="txAdvice" transaction-manager="transactionManager">
        <tx:attributes>
            <tx:method name="save*" propagation="REQUIRED"/>
            <tx:method name="del*" propagation="REQUIRED"/>
            <tx:method name="update*" propagation="REQUIRED"/>
            <tx:method name="find*" read-only="true"/>
        </tx:attributes>
    </tx:advice>
    <!--切面的配置-->
    <aop:config>
        <aop:pointcut id="pt" expression="execution(* com.offcn.service.*.*(..))"/>
        <aop:advisor pointcut-ref="pt" advice-ref="txAdvice" />
    </aop:config>

</beans>
```

此处需要注意：

​**context:component-scan**：扫描的节点路径为包含 service 和 dao 两个子目录的父级目录；

​**transactionManager**： 此节点作用就是初始化 Spring 框架提供的事务管理器的实现类；

​**tx:advice:** 此节点的作用是配置切面的通知，因为之前我们的切面类是自定义的，这里使用的是 Spring 提供的事务管理器类作为切面，那么针对什么方法需要做增强呢，在此节点配置，可以看得出来：以 save、del、update 开头的方法都会支持事务，而 find 开头的方法，指定的是只读。

​**aop:config:** 此节点就是 AOP 的相关配置节点了，将切入点和通知整合到一起，同以前的项目差别不大。这里可以看到：切入点规则是针对 service 下面的所有类所有方法任意参数做增强。通知使用的就是我们上面配置过的 tx:advice 节点。

**5. 测试代码**

```java
public class AccountServiceTest {

    public static void main(String[] args) {
        //1.获取容器
        ApplicationContext ac = new ClassPathXmlApplicationContext("applicationContext.xml");
        //2.获取业务对象
        UserService userService = ac.getBean(UserService.class);
        //3.从id为1的账号转成1000到2账号
        userService.transMoney(1,2,1000);
        System.out.println("转账完成..");
    }
}
```

**6. 测试结果：**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3g9of7j61ay0hc7s502)

ok, 大家，控制台如愿以偿打印了异常的堆栈信息，但是这个不是目的，哈哈哈。目的是在程序执行发生异常的情况下，数据的数据不会错乱。我们可以看见数据库数据并没有发生改变。

## 3. 总结

Spring 对于事务的控制，我们今天就到这里。通过本小节，我们也能体会到，使用 Spring 对事务控制还是非常简单的。无非以下三个注意事项：

1. 配置 Spring 框架提供的事务管理器；
2. 配置控制事务使用的通知；
3. 配置切面将通知与切入点结合即可。

没有比人更高的山… 没有比脚更长的路… 继续加油哦！

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
