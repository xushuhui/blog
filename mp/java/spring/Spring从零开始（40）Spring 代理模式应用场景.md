# Spring 代理模式应用场景

## 1. 前言

大家好，我们学习了代理模式的概念，也知道了代理模式可以在程序的运行过程中，实现对某个方法的增强。那么，在我们程序的编写过程中，

什么样的场景，能使用代理模式呢？

本节，我们模拟一个实际应用场景，目的是观察日常程序中可能发生的问题，以及代理模式如何解决问题，这样可以更加深刻地理解代理模式的意义。

## 2. 案例实战

### 2.1 转账工程的搭建

我们模拟一个实际生活中常见的情景，就是账号的转账。 假设有两个用户 A 和 用户 B，我们通过程序，从 A 账号中转成指定的 money 到 B 账号中。

那么，针对正常和异常的程序执行，我们来分析下问题以及它的解决方案。

#### 2.1.1 工程准备

**创建 maven 工程**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo39hwdzj60jp0g0gn002)

**引入 pom 文件的依赖 jar 包坐标信息**

```java
<dependencies>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-context</artifactId>
        <version>5.0.2.RELEASE</version>
    </dependency>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-test</artifactId>
        <version>5.0.2.RELEASE</version>
    </dependency>
    <dependency>
        <groupId>commons-dbutils</groupId>
        <artifactId>commons-dbutils</artifactId>
        <version>1.4</version>
    </dependency>

    <dependency>
        <groupId>mysql</groupId>
        <artifactId>mysql-connector-java</artifactId>
        <version>5.1.6</version>
    </dependency>

    <dependency>
        <groupId>c3p0</groupId>
        <artifactId>c3p0</artifactId>
        <version>0.9.1.2</version>
    </dependency>

    <dependency>
        <groupId>junit</groupId>
        <artifactId>junit</artifactId>
        <version>4.12</version>
    </dependency>
</dependencies>
```

**Spring 框架的配置文件编写**

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
        <property name="runner" ref="queryRunner"></property>
        <!-- 注入ConnectionUtils -->
        <property name="connectionUtils" ref="connectionUtils"></property>
    </bean>
    <!--配置QueryRunner-->
    <bean id="queryRunner" class="org.apache.commons.dbutils.QueryRunner" scope="prototype"></bean>
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
</beans>
```

配置文件说明：

1. connectionUtils 是获取数据库连接的工具类；
2. dataSource 采用 c3p0 数据源，大家一定要注意数据库的名称与账号名和密码；
3. queryRunner 是 dbutils 第三方框架提供用于执行 SQL 语句，操作数据库的一个工具类；
4. accountDao 和 accountService 是我们自定义的业务层实现类和持久层实现类。

**项目使用数据库环境**

CREATE TABLE `account` (

`id` int(11) NOT NULL auto_increment,

`accountNum` varchar(20) default NULL,

`money` int(8) default NULL,

PRIMARY KEY (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo39tua6j60no040wgi02)

#### 2.1.2 代码编写

**实体类代码**

```java
public class Account implements Serializable {
    //数据id
    private Integer id;
    //账号编码
    private String accountNum;
    //账号金额
    private Float money;
    //省略 get 和 set 的方法

}
```

**持久层接口**

```java
//接口代码
public interface IAccountDao {

    /**
     * 更新
     * @param account
     */
    void updateAccount(Account account);

    /**
     * 根据编号查询账户
     * @param accountNum
     * @return  如果有唯一的一个结果就返回，如果没有结果就返回null
     *          如果结果集超过一个就抛异常
     */
    Account findAccountByNum(String accountNum);
}
```

**持久层实现类**

```java
public class AccountDaoImpl implements IAccountDao {
	//数据库查询工具类
    private QueryRunner runner;
	//数据库连接工具类
    private ConnectionUtils connectionUtils;
    //省略 get 和 set 的方法


    //修改账号的方法
    public void updateAccount(Account account) {
        try{
            runner.update(connectionUtils.getThreadConnection(),
                          "update account set accountNum=?,money=? where id=?",account.getAccountNum(),account.getMoney(),account.getId());
        }catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
	//根据账号查询 Account 对象的方法
    public Account findAccountByNum(String accountNum) {
        try{
            List<Account> accounts = runner.query(connectionUtils.getThreadConnection(),
                                                  "select * from account where accountNum = ? ",new BeanListHandler<Account>(Account.class),accountNum);
            if(accounts == null || accounts.size() == 0){
                return null;
            }
            if(accounts.size() > 1){
                throw new RuntimeException("结果集不唯一，数据有问题");
            }
            return accounts.get(0);
        }catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
```

**业务层接口**

```java
public interface IAccountService {
    /**
     * 转账
     * @param sourceAccount        转出账户名称
     * @param targetAccount        转入账户名称
     * @param money             转账金额
     */
    void transfer(String sourceAccount, String targetAccount, Integer money);

}
```

**业务层实现类**

```java
public class AccountServiceImpl implements IAccountService {
	//持久层对象
    private IAccountDao accountDao;
    //省略 set 和 get 方法

    //转账的方法
    public void transfer(String sourceAccount, String targetAccount, Integer money) {
		//查询原始账户
        Account source = accountDao.findAccountByNum(sourceAccount);
        //查询目标账户
        Account target = accountDao.findAccountByNum(targetAccount);
        //原始账号减钱
        source.setMoney(source.getMoney()-money);
        //目标账号加钱
        target.setMoney(target.getMoney()+money);
        //更新原始账号
        accountDao.updateAccount(source);
        //更新目标账号
        accountDao.updateAccount(target);
        System.out.println("转账完毕");

    }
}
```

**测试运行类代码**

```java
@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations = "classpath:bean.xml")
public class AccountServiceTest {

    @Autowired
    @Qualifier("proxyAccountService")
    private  IAccountService as;

    @Test
    public  void testTransfer(){
        as.transfer("622200009999","622200001111",100);
    }
}
```

**测试结果**

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3a2oghj611m0gmtqz02)

代码执行完毕，可以看到输出打印转账 ok 了。那么数据库的数据有没有改变呢？我们再看一眼：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3ab8omj607h040myc02)

可以看到：两个账号的数据已经发生了改变，证明转账的动作，确实完成了。那这样看来，我们的代码也没有问题啊，代理模式有什么用呢？

接下来我们改造下工程，模拟程序发生异常时候，执行以后的结果如何。

#### 2.1.3 改造业务类代码

在业务层的代码加入一行异常代码，看看结果是否还会转账成功呢？

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3alvezj60z50e3dmo02)

执行结果：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3aytsqj611b0hc1gk02)

当然了，其实提前也能想得到，肯定会执行失败的啦，哈哈哈哈，我们手动加了运算会出现异常的代码嘛！但是转账的动作是不是也失败了呢？我们再来看一下数据库：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3bdccuj607h040my902)

问题来了： id 为 1 的账号 money 的列值由原来的 900 变成了 800，说明存款确实减少了 100，但是由于在代码执行的过程中，出现了异常，导致原始账号减少 100 的金钱后保存成功， 而 id 为 2 的账号并没有增加 100。这就出现了数据的事务问题，破坏了数据的原子性和一致性。

那么如何解决呢？ 思路就是将我们的数据操作代码，使用事务控制起来。由于本小节篇幅有限，我们留待下一小节解决。

## 3. 小结

本小节模拟了一个现实生活中转账的业务场景，其目的是为了引出我们的程序在执行过程中可能会产生的问题。而如何解决，并且对于原始代码侵入性更小，耦合性更低，是我们需要思考的事情。

使用知识点：

1. JDBC 的基础，本案例用到了 JDBC 连接数据库的工具类 dbutils 知识；
2. 数据库基础，本案例的数据操作，涉及到了事务的四个特性：原子性，一致性，隔离性，持久性。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
