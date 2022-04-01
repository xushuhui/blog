---
title: Java 从零开始（48）业务实现 1
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 实战 - 业务实现 1

上一小节我们完成了数据库的设计和创建，也向数据表中插入了一些初始数据，本小节我们将开始具体业务代码的实现，如果大家还没有完成上一小节的任务，请务必先完成再来学习本节内容。

## 1. 准备工作

在开始正式编码之前，我们要做一些准备工作，主要是环境的搭建和工具类的引入。

### 1.1 创建 Maven 工程

打开 `idea`，点击`Create new Project`按钮：

![](https://xushuhui.gitee.io/image/imooc/5f2a5d49098e98f815500970.jpg)

在左侧栏选择`Maven`，`Project SDK`选择`14`，勾选`Create from archetype`复选框，再选择`maven-archetype-quickstart`，表示创建一个简单 Java 应用，点击`next`按钮：

![](https://xushuhui.gitee.io/image/imooc/5f2b7b4409b13e4924911158.jpg)

输入项目名称`goods`，将项目路径设置为本地桌面，`GroupId`可根据实际情况自定义，此处我设置为`com.colorful`，其余输入框无需修改，采用默认即可，设置完成后，点击`next`按钮：

![](https://xushuhui.gitee.io/image/imooc/5f2b7e1d09d0fddd24991162.jpg)

这一步来到`Maven`配置，`idea`自带了`Maven`，我们使用默认的即可，直接点击`Finish`按钮完成项目创建：

![](https://xushuhui.gitee.io/image/imooc/5f2b7dc609ae540f24931159.jpg)

此时，`Maven`会进行一些初始化配置，右下角对话框选择`Enable Auto-import`按钮，表示允许自动导入依赖：

![](https://xushuhui.gitee.io/image/imooc/5f2b7e4f090ccb4227301740.jpg)

稍等片刻，待看到左侧项目的目录结构已经生成好了，及表示已完成项目的初始化工作：

![](https://xushuhui.gitee.io/image/imooc/5f2b7e6c09ff91e527321739.jpg)

### 1.2 引入 MySQL 驱动

接下来引入`mysql-connector-java`驱动，由于我本地安装的`MySQL`版本为`8.0.21`，因此`mysql-connector-java`的版本号也选择`8.0.21`，大家根据自己实际情况选择对应版本。

打开`pom.xml`文件，在`<dependencies></dependencies>`节点内插入如下`xml`：

```java
<!-- https://mvnrepository.com/artifact/mysql/mysql-connector-java -->
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.21</version>
</dependency>
```

![](https://xushuhui.gitee.io/image/imooc/5f2b7e9209fd6bf015650802.jpg)

由于我们已经配置了允许自动导入依赖，稍等片刻，`mysql-connector-java 8.0.21`就会被成功导入。可在`idea`右侧点击`Maven`按钮查看项目的依赖关系：

![](https://xushuhui.gitee.io/image/imooc/5f2b7eaf09eab22809080494.jpg)

### 1.3 引入 JDBC 工具类

JDBC 相关操作是本项目的最常用的操作，我封装了一个 JDBC 的工具类，主要通过 Java 的 JDBC API 去访问数据库，提供了加载配置、注册驱动、获得资源以及释放资源等接口。

大家可以到我的 [`Github` 仓库】(https://github.com/colorful3/goods-cms/blob/master/src/main/java/com/colorful/util/JDBCUtil.java) 下载这个 `JDBCUtil`类；也可以直接复制下面的代码：

```java
package com.colorful.util;

import java.io.IOException;
import java.io.InputStream;
import java.sql.*;
import java.util.Properties;

/**
 * @author colorful@TaleLin
 */
public class JDBCUtil {

    private static final String driverClass;
    private static final String url;
    private static final String username;
    private static final String password;

    static {
        // 加载属性文件并解析
        Properties props = new Properties();
        // 使用类的加载器的方式进行获取配置
        InputStream inputStream = JDBCUtil.class.getClassLoader().getResourceAsStream("jdbc.properties");
        try {
            assert inputStream != null;
            props.load(inputStream);
        } catch (IOException e) {
            e.printStackTrace();
        }

        driverClass = props.getProperty("driverClass");
        url = props.getProperty("url");
        username = props.getProperty("username");
        password = props.getProperty("password");
    }

    /**
     * 注册驱动
     */
    public static void loadDriver() throws ClassNotFoundException{
        Class.forName(driverClass);
    }

    /**
     * 获得连接
     */
    public static Connection getConnection() throws Exception{
        loadDriver();
        return DriverManager.getConnection(url, username, password);
    }

    /**
     * 资源释放
     */
    public static void release(PreparedStatement statement, Connection connection){
        if(statement != null){
            try {
                statement.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
            statement = null;
        }
        if(connection != null){
            try {
                connection.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
            connection = null;
        }
    }

    /**
     * 释放资源 重载方法
     */
    public static void release(ResultSet rs, PreparedStatement stmt, Connection conn){
        if(rs!= null){
            try {
                rs.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
            rs = null;
        }
        if(stmt != null){
            try {
                stmt.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
            stmt = null;
        }
        if(conn != null){
            try {
                conn.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
            conn = null;
        }
    }

}
```

我本地将这个类放在了 `com.colorful.util`包下，大家可根据自身情况随意放置。另外，由于该类在静态代码块中加载了配置文件`jdbc.properties`，需要在`resource`下面新建一个 `jdbc.properties`文件，并写入一下内容：

```java
driverClass=com.mysql.cj.jdbc.Driver
url=jdbc:mysql:///imooc_goods_cms?serverTimezone=Asia/Shanghai&characterEncoding=UTF8
username=root
password=123456
```

我将数据放到了本地系统中，并且启动端口是默认 3306，大家根据自己的`MySQL`实际配置自行修改。

### 1.4 测试代码

为了测试我们的数据库配置以及 `JDBCUtil` 类是否成功引入，现在到 test 目录下，新建一个 `JDBCTest` 类：

```java
package com.colorful;

import com.colorful.util.JDBCUtil;
import org.junit.Test;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Timestamp;

public class JDBCTest {

    @Test
    public void testJDBC() {
        Connection connection = null;
        PreparedStatement preparedStatement = null;
        ResultSet resultSet = null;
        try {
            // 获得链接
            connection = JDBCUtil.getConnection();
            // 编写 SQL 语句
            String sql = "SELECT * FROM `imooc_user` where `id` = ?";
            // 预编译 SQL
            preparedStatement = connection.prepareStatement(sql);
            // 设置参数
            preparedStatement.setInt(1, 1);
            resultSet = preparedStatement.executeQuery();
            if (resultSet.next()) {
                int id = resultSet.getInt("id");
                String nickname = resultSet.getString("nickname");
                Timestamp createTime = resultSet.getTimestamp("create_time");
                System.out.println("id=" + id);
                System.out.println("nickname=" + nickname);
                System.out.println("createTime=" + createTime);
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 释放资源
            JDBCUtil.release(resultSet, preparedStatement, connection);
        }
    }

}
```

如果配置成功，运行单元测试，将得到如下运行结果：

```java
id=1
nickname=小慕
createTime=2020-07-20 16:53:19.0
```

下面为运行截图：

![](https://xushuhui.gitee.io/image/imooc/5f2b7ed409f0582d17161432.jpg)

## 2. 系统架构

本商品管理系统的包结构如下：

```java
src
├── main
│   ├── java    # 源码目录
│   │   └── com
│   │       └── colorful
│   │           ├── App.java    # 入口文件
│   │           ├── dao         # 数据访问对象（Data Access Object，提供数据库操作的一些方法）
│   │           ├── model       # 实体类（类字段和数据表字段一一对应）
│   │           ├── service     # 服务层（提供业务逻辑层服务）
│   │           └── util        # 一些帮助类
│   └── resources
│       ├── imooc_goods_cms.sql # 建表的 SQL 文件
│       └── jdbc.properties     # jdbc 配置文件
└── test       # 单元测试目录
    └── java
        └── com
            └── colorful
                ├── AppTest.java
                └── JDBCTest.java
```

大家可以提前熟悉一下本项目的项目结构，下面我们会一一讲解。

## 3. 实体类

实体类的作用是存储数据并提供对这些数据的访问。在我们这个项目中，实体类统一被放到了`model`包下，通常情况下，实体类中的属性与我们的数据表字段一一对应。当我们编写这些实体类的时候，建议对照着数据表的字段以防疏漏。

### 3.1 BaseModel

在我们数据表中，有几个公共的字段，可以提取出一个实体类的父类 `BaseModel` ，并提供 `getter` 和 `setter`，源码如下：

```java
package com.colorful.model;

import java.sql.Timestamp;

public class BaseModel {

    private Integer id;

    private Timestamp createTime;

    private Timestamp updateTime;

    private Timestamp deleteTime;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public Timestamp getCreateTime() {
        return createTime;
    }

    public void setCreateTime(Timestamp createTime) {
        this.createTime = createTime;
    }

    public Timestamp getUpdateTime() {
        return updateTime;
    }

    public void setUpdateTime(Timestamp updateTime) {
        this.updateTime = updateTime;
    }

    public Timestamp getDeleteTime() {
        return deleteTime;
    }

    public void setDeleteTime(Timestamp deleteTime) {
        this.deleteTime = deleteTime;
    }

}
```

值得注意的是，`Timestamp`是`java.sql`下的类。

### 3.2 实体类编写

接下来，再在`model`包下新建 3 个类：`User`、`Goods` 和 `Category`，并提供`getter` 和 `setter` 。如下是每个类的代码：

```java
package com.colorful.model;

public class User extends BaseModel {

    private String userName;

    private String nickName;

    private String password;

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getNickName() {
        return nickName;
    }

    public void setNickName(String nickName) {
        this.nickName = nickName;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

}
```

```java
package com.colorful.model;

public class Goods extends BaseModel {

    private String name;

    private String description;

    private Integer categoryId;

    private Double price;

    private Integer stock;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Integer getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(Integer categoryId) {
        this.categoryId = categoryId;
    }

    public Double getPrice() {
        return price;
    }

    public void setPrice(Double price) {
        this.price = price;
    }

    public Integer getStock() {
        return stock;
    }

    public void setStock(Integer stock) {
        this.stock = stock;
    }

}
```

```java
package com.colorful.model;

public class Category extends BaseModel {

    private String name;

    private String description;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

}
```

## 4. 实现用户鉴权

### 4.1 登录方式

想要使用系统进行商品管理，第一步要做的就是登录。

我们的系统使用用户名和密码进行登录校验，上一小节我们已经建立了`imooc_user`表，并向表中插入了一个用户 `admin`，其密码为 `123456` 。显然，通过如下`SQL`就可以查询到该用户：

```java
SELET * FROM `imooc_user` WHERE `username` = 'admin' AND password = '123456';
```

如果查询到这个用户，就表示用户名密码通过校验，用户可执行后续操作，如果没有查到，就要提示用户重新输入账号和密码。

### 4.2 数据访问对象

我们先不管用户是如何输入账号密码的，接下来要编写的业务代码就是根据用户名和密码去查询用户。那么涉及到数据库查询的代码应该放到哪里呢？参考上面的系统架构图，`DAO`是数据访问对象，我们可以在`dao`包下面新建一个`UserDAO`，并写入如下代码：

```java
package com.colorful.dao;

import com.colorful.model.User;
import com.colorful.util.JDBCUtil;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;

public class UserDAO {

    public User selectByUserNameAndPassword(String username, String password) {
        Connection connection = null;
        PreparedStatement preparedStatement = null;
        ResultSet resultSet = null;
        User user = new User();
        try {
            // 获得链接
            connection = JDBCUtil.getConnection();
            // 编写 SQL 语句
            String sql = "SELECT * FROM `imooc_user` where `username` = ? AND `password` = ? AND `delete_time` is null ";
            // 预编译 SQL
            preparedStatement = connection.prepareStatement(sql);
            // 设置参数
            preparedStatement.setString(1, username);
            preparedStatement.setString(2, password);
            resultSet = preparedStatement.executeQuery();
            if (resultSet.next()) {
                user.setId(resultSet.getInt("id"));
                String nickname = resultSet.getString("nickname");
                if (nickname.equals("")) {
                    nickname = "匿名";
                }
                user.setNickName(nickname);
                user.setUserName(resultSet.getString("username"));
            } else {
                user = null;
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 释放资源
            JDBCUtil.release(resultSet, preparedStatement, connection);
        }
        return user;
    }

}
```

`UserDAO` 类下面有一个 `selectByUserNameAndPassword()`方法， 接收两个参数 `username` 和 `password`，返回值类型是实体类 `User`，如果没有查询到，返回的是一个 `null`。

完成了 `UserDAO` 的编写，我们需要到服务层 `service`包下，新建一个 `UserService` ，并写入如下代码：

```java
package com.colorful.service;

import com.colorful.dao.UserDAO;
import com.colorful.model.User;

public class UserService {

    private final UserDAO userDAO = new UserDAO();

  	// 登陆
    public User login(String username, String password) {
        return userDAO.selectByUserNameAndPassword(username, password);
    }

}
```

到这里大家可能有些疑问，这个类下面的`login()`方法，直接调用了我们刚刚编写的 `DAO` 下面的 `selectByUserNameAndPassword()` 方法，为什么还要嵌套这么一层么？这不是多此一举么？

要讨论 `service` 层的封装是不是过度设计，就要充分理解设计服务层的概念和意义，服务层主要是对业务逻辑的封装，对于更为复杂的项目，用户登录会有更多的方式，因此在服务层，会封装更多的业务逻辑。如果没有服务层，这些复杂的逻辑不得不都写在数据访问层，显然这是不合理的。我们现在这个项目没有使用任何框架，等到后面大家学习了`Spring`这种框架，一定会对这样的分层的好处有所体会。

### 4.3 使用 Scanner 类与用户交互

完成了上面一系列的封装，就剩下我们和用户的交互了，本项目中，我们使用 `Scanner` 类来接收用户的输入，并使用`print()`方法向屏幕输出。

打开 `App.java` 入口文件，创建`UserService`实例，编写一个主流程方法 `run()`，并在入口方法 `main()`中调用该方法：

```java
package com.colorful;

import com.colorful.model.User;
import com.colorful.service.UserService;

import java.util.Scanner;

/**
 * @author colorful@TaleLin
 * Imooc Goods
 */
public class App {

    private static final UserService userService = new UserService();

    /**
     * 主流程方法
     */
    public static void run() {
        User user = null;
        System.out.println("欢迎使用商品管理系统，请输入用户名和密码：");
        do {
            Scanner scanner = new Scanner(System.in);
            // 登录
            System.out.println("用户名：");
            String username = scanner.nextLine();
            System.out.println("密码：");
            String password = scanner.nextLine();
            user = userService.login(username, password);
            if (user == null) {
                System.out.println("用户名密码校验失败，请重新输入！");
            }
        } while (user == null);
        System.out.println("欢迎您！" + user.getNickName());
        // TODO 登录成功，编写后续逻辑
    }

    public static void main( String[] args )
    {
        run();
    }
}
```

`run()`方法中有一个 `do ... while`循环，循环的条件是 `user` 对象为 `null`。

我们知道，`do... while`循环会首先执行 `do` 中的循环体，循环体中创建了一个 `Scanner` 类的实例，获取到用户的输入后，我们会调用用户服务层的`login()`方法，该方法返回实体类对象`User`，如果其为 `null`表示用户名密码校验失败，需要用户重新输入， `user == null`，满足循环的条件，会一直执行循环体中的代码。直到循环体中的 `user`不为 `null` （也就是用户登录校验成功后）才终止循环。

下面运行`App.java`的`main()`方法，如下为登录失败的截图：

![](https://xushuhui.gitee.io/image/imooc/5f2b7f0109e6b06612540864.jpg)

如果用户名密码检验错误，就要反复输入用户名密码重新登录。

如下为登录成功的截图：

![](https://xushuhui.gitee.io/image/imooc/5f2b7f17091116f510500712.jpg)

## 5. 小结

在本小节，我们成功搭建了项目工程，通过实现一个用户鉴权模块，介绍了整体的系统架构。我们在编写实体类的同时，复习了面向对象的继承性；在数据访问层，也复习了 `JDBC API` 的使用；在编写程序入口文件的同时，也复习了 `Scanner` 类的使用和循环的使用。

关于系统鉴权，这里还有一个待优化的地方，大家下去之后可以思考一下，在下一小节的开头，我将带领大家一起来优化。下一小节也将主要讲解最后剩余的商品模块和分类模块的实现，也会复习到很多其他方面的基础知识。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
