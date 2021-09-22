---
title: Java 从零开始（39）Java 数据库编程
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 数据库编程

本小节我们将学习如何使用 Java 语言结合数据库进行编程。注意，学习本小节需要你有一定的 SQL 基础，了解 MySQL 数据库的 基础 CRUD 操作，如果你还不了解 SQL ，推荐先去学习一个非常不错的 [wiki 教程](http://www.imooc.com/wiki/sqlbase)，只需掌握前几节的 SQL 初级知识即可。

本小节我们将选择开源免费的 `MySQL 5.7` 作为数据库，可以去官网下载并安装 `MySQL`，如果你不知如何下载安装，推荐按照 [这篇文章](http://www.imooc.com/wiki/mysqllesson/mysqlwindows.html) 来做。

通过本小节的学习，你将了解到什么是 JDBC，如何连接数据库，如何关闭数据库，JDBC 的新增、查询、更新和删除接口，如何执行批量等内容。

## 1. JDBC 概述

### 1.1 什么是 JDBC

JDBC 就是 `Java DataBase Connectivity` 的缩写，它是 Java 的标准库，定义了客户端访问数据库的 API。

市面上的数据库有很多种类，连接不同的数据库需要加载不同的**数据库驱动**。数据库驱动是由厂商提供的，需要我们引入。标准库编写了一套访问数据库的代码，因此不需要标准库代码的改动，只需加载不同的驱动，就可以访问不同的数据库。

### 1.2 JDBC 的作用

在 JDBC 出现之前，数据库驱动程序由不同的数据库厂商提供，程序员想要操作不同的数据库，就不得不学习各类不同驱动的用法，驱动的学习成本和代码的维护成本都非常高。

![](https://xushuhui.gitee.io/image/imooc/5ef2d2fd0972443312721120.jpg)

`Sun` 公司发现了这个问题，因此定义了一套标准的访问数据库的 API（即 JDBC），不同厂商按照这个 API 提供的统一接口，实现驱动，这保证了数据库操作的统一性。程序员也不需要再去学习不同厂商提供的五花八门的 API，只需要学习 JDBC 标准 API 即可。代码维护的成本也大大降低。

![](https://xushuhui.gitee.io/image/imooc/5ef2d39b09806ad210391056.jpg)

## 2. 连接数据库

### 2.1 建库建表

打开 `MySQL` 客户端，执行如下 sql 语句：

```java
-- 创建数据库并使用
CREATE DATABASE jdbcdemo;
USE jdbcdemo;

-- 创建数据表
CREATE TABLE `user` (
	`id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(20) NOT NULL comment "用户名",
  `nickname` varchar(20) NOT NULL comment "昵称",
  `password` CHAR(32) NOT NULL comment "密码"
);

-- 插入一些数据
insert into `user` values(null, "Colorful", "Colorful3", "123456");
insert into `user` values(null, "imooc", "小慕", "123456");
insert into `user` values(null, "Lillian", "小李", "123456");
```

执行截图如下：

![](https://xushuhui.gitee.io/image/imooc/5ef95092092bcf2206010730.jpg)

查询 `user` 表所有记录，结果如下：

![](https://xushuhui.gitee.io/image/imooc/5ef950d50951104e03420202.jpg)

### 2.2 引入驱动

去 [maven 中央仓库](https://mvnrepository.com/artifact/mysql/mysql-connector-java/5.1.49) 找到 `mysql-connector-java` 驱动。如果你熟悉 `Maven`，可直接引入 maven 依赖：

```java
<!-- https://mvnrepository.com/artifact/mysql/mysql-connector-java -->
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>5.1.49</version>
</dependency>
```

如果你还不熟悉 Maven，请跟着我来做如下步骤：

到 [maven 中央仓库](https://mvnrepository.com/artifact/mysql/mysql-connector-java/5.1.49) 下载 `jar` 包，鼠标左键单击 `jar`：

![](https://xushuhui.gitee.io/image/imooc/5ef951020935c25c08820630.jpg)

在工程目录下新建一个 `lib` 包，并将刚刚下载好的 `jar` 包复制到其中：

![](https://xushuhui.gitee.io/image/imooc/5ef951440966abe007170166.jpg)

在 `jar` 包上点击右键，选择 `Add as Library`：

![](https://xushuhui.gitee.io/image/imooc/5ef9518409fd4dce12811169.jpg)

如果有弹窗，单击确定即可。此时就可以在你的项目中引入驱动里的类了。

### 2.3 实例

我们下面通过实例代码来演示 JDBC 的简单使用，以下实例代码有这样几个步骤：

1. 加载数据库驱动；
2. 建立连接；
3. 创建 `Statement` 对象，用于向数据库发送 SQL 语句；
4. 获取 `ResultSet` 对象，取出数据，此对象代表结果集；
5. 释放资源，断开与数据库的连接。

具体实例如下：

```java
package com.imooc.jdbc;

import com.mysql.jdbc.Driver;

import java.sql.*;

public class JDBCDemo1 {

    public static void main(String[] args) throws SQLException {
        // 1. 加载数据库驱动
        DriverManager.registerDriver(new Driver());
        // 2. 建立连接
        final String url = "jdbc:mysql://localhost:3306/jdbcdemo";  // 数据库 url
        final String user = "root"; // 数据库用户名
        final String password = "123456"; // 数据库密码
        Connection connection = DriverManager.getConnection(url, user, password);
        // 3. 创建 Statement 对象，用于向数据库发送 SQL 语句
        String sql = "SELECT * FROM `user`";
        Statement statement = connection.createStatement();
        ResultSet resultSet = statement.executeQuery(sql);
        // 4. 获取 ResultSet 对象，取出数据，此对象代表结果集
        while (resultSet.next()) {
            int id = resultSet.getInt("id");
            String username = resultSet.getString("username");
            String nickname = resultSet.getString("nickname");
            String pwd = resultSet.getString("password");
            System.out.println("id=" + id + "; username=" + username + "; nickname=" + nickname + "; password=" + pwd + '\r');
        }
        // 5. 释放资源，断开与数据库的连接（调用 close() 方法）
        // 5.1 释放 ResultSet
        resultSet.close();
        // 5.2 释放 Statement
        statement.close();
        // 5.3 释放 Connection
        connection.close();
    }
}
```

运行结果：

```java
id=1; username=Colorful; nickname=Colorful3; password=123456
id=2; username=imooc; nickname=小慕；password=123456
id=3; username=Lillian; nickname=小李；password=123456
```

看了实例代码，你可能有些晕，这写类都是干嘛的呀？别担心，我们下面就来一一讲解。

## 4. JDBC 几个类的详解

在上面的实例程序中，我们用到了几个 JDBC 的类：`DriverManager`、`Collection`、`Statement` 和 `ResultSet`，下面我们将详细介绍这几个类。

### 4.1 DriverManager

DriverManager 是驱动管理类，此类用于注册驱动和获得连接。

在实际开发中，我们不是像实例这样注册驱动的，我们编写这样地代码会导致驱动注册两次：

```java
// 1. 加载数据库驱动
DriverManager.registerDriver(new Driver());
```

通过查看 `com.mysql.jdbc.Driver` 源码，我们发现在静态代码块处已经注册过了驱动：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef9522509b82e3e12750472.jpg)

那么如何改写呢，我们可以使用反射机制来注册驱动：

```java
Class.forName("com.mysql.jdbc.Driver");
```

加载了 `Driver` 类，其静态代码块就会执行，因此也就注册了驱动。

除了获得驱动，我们还可以调用 `getConnection(url, user, password)` 方法来获得连接，其中 `url` 这个参数不是很好理解：

```java
String url = "jdbc:mysql://localhost:3306/jdbcdemo";
```

其中 `jdbc` 是协议，`mysql` 是子协议，`localhost` 是主机名，`3306` 是端口号，`jdbcdemo` 是数据库名。/ 这里的协议是固定地写法，连接不同类型地数据库需要不同地协议。

### 4.2 Connection

Connection 是连接对象，它可以创建执行 SQL 语句的对象，还可以进行事务的管理。

下面列举了 `Connection` 类的常用实例方法：

* `Statement createStatement()`：创建执行 SQL 语句对象，又 SQL 注入风险；
* `PrepareStatement prepareStatement(String sql)`：预编译 SQL 语句，解决 SQL 注入的漏洞；
* `CallableStatement prepareCall(String sql)`：执行 SQL 中存储过程；
* `setAutoCommit(boolean autoCommit)`：设置事务是否自动提交；
* `commit()`：事务提交；
* `rollback()`：事务回滚。

### 4.3 Statement

Statement 是执行 SQL 语句的对象，下面列举了 `Statement` 类的常用实例方法：

* `boolean execute(String sql)`：执行 SQL 语句，如果返回的第一个结果为 `ResultSet` 对象，则返回 `true`，如果其为更新计数或者不存在任何结果，则返回 `false`。该方法不常用；
* `ResultSet executeQuery(String sql)`：执行 SQL 中的 `select` 语句；
* `int executeUpdate(String sql)`：执行 SQL 中的 `insert`、`update`、`delete` 语句，返回影响的行数。

Statement 还可以执行批量操作，关于批量操作，我们将在下面学习。

### 4.4 ResultSet

ResultSet 是结果集对象，它是 `select` 语句查询的结果的封装。下面列举了 `ResultSet` 类的常用实例方法：

* `boolean next()`：将光标从当前位置向前移一行，判断是否有下一行记录；
* `getString(String columnLable)`：以 Java 语言中 String 的形式获取此 ResultSet 对象的当前行中指定的值；
* `getInt(String columnLable)`：以 Java 语言中 int 的形式获取此 ResultSet 对象的当前行中指定的值；
* `getXXX()`：对于不同类型的数据，可以使用 `getXXX()` 来获取数据（例如 `getString()`，`getInt()`），另外还有一个通用的 `getObject()` 方法，用于获取所有 `Object` 类型的数据。

### 4.5 JDBC 资源的释放

JDBC 程序运行完成后，一定要记得释放程序在运行过程中，创建的那些与数据库进行交互的对象，这些对象通常是 `ResultSet`、`Statement` 和 `Connection` 对象。特别是 `Connection` 对象，它是非常稀有的资源，用完后必须马上释放，如果此对象不能及时、正确的关闭，极易导致系统的宕机。`Connection` 对象的使用原则是尽量晚创建，尽量早释放。

## 5. CRUD

### 5.1 新增数据

在执行新增数据的代码前，`user` 表中有如下数据：

```java
mysql> select * from user;
+----+--------------+--------------+----------+
| id | username     | nickname     | password |
+----+--------------+--------------+----------+
|  1 | Colorful     | Colorful3    | 123456   |
|  2 | imooc        | 小慕         | 123456   |
|  3 | Lillian      | 小李         | 123456   |
+----+--------------+--------------+----------+
3 rows in set (0.00 sec)
```

新增数据的实例代码如下：

```java
package com.imooc.jdbc;

import java.sql.*;

public class JDBCDemo2 {

    /**
     * 插入语句
     * @return 受影响的行数
     */
    public static int insert() {
        final String url = "jdbc:mysql://localhost:3306/jdbcdemo";  // 数据库 url
        final String user = "root"; // 数据库用户名
        final String password = "123456"; // 数据库密码

        Connection connection = null;
        Statement  statement = null;
        int result = 0;

        try {
            // 1. 加载数据库驱动
            Class.forName("com.mysql.jdbc.Driver");
            // 2. 建立连接
            connection = DriverManager.getConnection(url, user, password);
            // 3. 创建 Statement 对象，用于向数据库发送 SQL 语句
            String sql = "INSERT INTO `user` VALUES(null, \"testUsername\", \"testNickname\", \"123456\")";
            statement = connection.createStatement();
            result = statement.executeUpdate(sql);
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 4. 释放资源，断开与数据库的连接（调用 close() 方法）
            if (statement != null) {
                try {
                    statement.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                statement = null;
            }
            if (connection != null) {
                try {
                    connection.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                // 此处手动设置为 null，有益于垃圾回收机制更早地回收对象
                connection = null;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        int rows = JDBCDemo2.insert();
        System.out.println("受影响的行数为：" + rows);
    }

}
```

运行结果：

```java
受影响的行数为：1
```

此时，查询数据库中的记录，可以发现多了 1 条：

```java
mysql> select * from user;
+----+--------------+--------------+----------+
| id | username     | nickname     | password |
+----+--------------+--------------+----------+
|  1 | Colorful     | Colorful3    | 123456   |
|  2 | imooc        | 小慕         | 123456   |
|  3 | Lillian      | 小李         | 123456   |
|  4 | testUsername | testNickname | 123456   |
+----+--------------+--------------+----------+
4 rows in set (0.00 sec)
```

### 5.2 读取数据

读取数据的示例如下：

```java
package com.imooc.jdbc;

import java.sql.*;

public class JDBCDemo3 {

    /**
     * 根据 id 查询用户
     * @param id 用户的 id
     */
    public static void selectUserById(int id) {
        final String url = "jdbc:mysql://localhost:3306/jdbcdemo";  // 数据库 url
        final String user = "root"; // 数据库用户名
        final String password = "123456"; // 数据库密码
        Connection connection = null;
        Statement statement = null;
        ResultSet result = null;
        try {
            // 1. 加载数据库驱动
            Class.forName("com.mysql.jdbc.Driver");
            // 2. 建立连接
            connection = DriverManager.getConnection(url, user, password);
            // 3. 创建 Statement 对象，用于向数据库发送 SQL 语句
            String sql = "SELECT * FROM `user` WHERE id = " + id;
            statement = connection.createStatement();
            // 4. 获取 ResultSet 对象，取出数据
            result = statement.executeQuery(sql);
            while (result.next()) {
                String nickname = result.getString("nickname");
                String username = result.getString("username");
                String pwd = result.getString("password");
                System.out.println("id=" + id + "; username=" + username + "; nickname=" + nickname + "; password=" + pwd + '\r');
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 5. 释放资源，断开与数据库的连接（调用 close() 方法）
            if (result != null) {
                try {
                    result.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                result = null;
            }
            if (statement != null) {
                try {
                    statement.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                statement = null;
            }
            if (connection != null) {
                try {
                    connection.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                // 此处手动设置为 null，有益于垃圾回收机制更早地回收对象
                connection = null;
            }
        }
    }

    public static void main(String[] args) throws SQLException {
        JDBCDemo3.selectUserById(4);
    }

}
```

运行结果：

```java
id=3; username=Lillian; nickname=小李；password=123456
```

### 5.3 更新数据

在执行更新数据的代码前，`user` 表中有如下数据：

```java
mysql> select * from user;
+----+--------------+--------------+----------+
| id | username     | nickname     | password |
+----+--------------+--------------+----------+
|  1 | Colorful     | Colorful3    | 123456   |
|  2 | imooc        | 小慕         | 123456   |
|  3 | Lillian      | 小李         | 123456   |
|  4 | testUsername | testNickname | 123456   |
+----+--------------+--------------+----------+
4 rows in set (0.00 sec)
```

更新数据的实例代码如下：

```java
package com.imooc.jdbc;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

public class JDBCDemo4 {

    /**
     * 更新
     * @return 受影响的行数
     */
    public static int updateById(int id) {
        final String url = "jdbc:mysql://localhost:3306/jdbcdemo";  // 数据库 url
        final String user = "root"; // 数据库用户名
        final String password = "123456"; // 数据库密码
        Connection connection = null;
        Statement  statement = null;
        int result = 0;

        try {
            // 1. 加载数据库驱动
            Class.forName("com.mysql.jdbc.Driver");
            // 2. 建立连接
            connection = DriverManager.getConnection(url, user, password);
            // 3. 创建 Statement 对象，用于向数据库发送 SQL 语句
            String sql = "UPDATE `user` SET `nickname` = '更新后的 nickname' WHERE id = " + id;
            statement = connection.createStatement();
            result = statement.executeUpdate(sql);
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 4. 释放资源，断开与数据库的连接（调用 close() 方法）
            if (statement != null) {
                try {
                    statement.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                statement = null;
            }
            if (connection != null) {
                try {
                    connection.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                // 此处手动设置为 null，有益于垃圾回收机制更早地回收对象
                connection = null;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        int rows = JDBCDemo4.updateById(4);
        System.out.println("受影响的行数为：" + rows);
    }

}
```

运行结果：

```java
受影响的行数为：1
```

更新数据代码执行完成后，可以观察到 `id` 为 `4` 的记录 `nickname` 字段发生了改变：

```java
mysql> select * from user;
+----+--------------+----------------------+----------+
| id | username     | nickname             | password |
+----+--------------+----------------------+----------+
|  1 | Colorful     | Colorful3            | 123456   |
|  2 | imooc        | 小慕                 | 123456   |
|  3 | Lillian      | 小李                 | 123456   |
|  4 | testUsername | 更新后的 nickname     | 123456   |
+----+--------------+----------------------+----------+
4 rows in set (0.00 sec)
```

### 5.4 删除数据

在执行删除数据的代码前，`user` 表中有如下数据：

```java
mysql> select * from user;
+----+--------------+----------------------+----------+
| id | username     | nickname             | password |
+----+--------------+----------------------+----------+
|  1 | Colorful     | Colorful3            | 123456   |
|  2 | imooc        | 小慕                 | 123456   |
|  3 | Lillian      | 小李                 | 123456   |
|  4 | testUsername | 更新后的 nickname     | 123456   |
+----+--------------+----------------------+----------+
4 rows in set (0.00 sec)
```

删除数据的实例代码如下：

```java
package com.imooc.jdbc;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

public class JDBCDemo5 {

    /**
     * 删除
     * @return 受影响的行数
     */
    public static int deleteById(int id) {
        final String url = "jdbc:mysql://locahost:3306/jdbcdemo";  // 数据库 url
        final String user = "root"; // 数据库用户名
        final String password = "123456"; // 数据库密码

        Connection connection = null;
        Statement  statement = null;
        int result = 0;

        try {
            // 1. 加载数据库驱动
            Class.forName("com.mysql.jdbc.Driver");
            // 2. 建立连接
            connection = DriverManager.getConnection(url, user, password);
            // 3. 创建 Statement 对象，用于向数据库发送 SQL 语句
            String sql = "DELETE FROM `user` WHERE id = " + id;
            statement = connection.createStatement();
            result = statement.executeUpdate(sql);
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 4. 释放资源，断开与数据库的连接（调用 close() 方法）
            if (statement != null) {
                try {
                    statement.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                statement = null;
            }
            if (connection != null) {
                try {
                    connection.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
                // 此处手动设置为 null，有益于垃圾回收机制更早地回收对象
                connection = null;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        int rows = JDBCDemo5.deleteById(4);
        System.out.println("受影响的行数为：" + rows);
    }

}
```

运行结果：

```java
受影响的行数为：1
```

在执行删除数据的代码后，可观察到 `id` 为 4 的记录被成功删除，数据库只剩下了 3 行记录：

```java
mysql> select * from user;
+----+--------------+----------------------+----------+
| id | username     | nickname             | password |
+----+--------------+----------------------+----------+
|  1 | Colorful     | Colorful3            | 123456   |
|  2 | imooc        | 小慕                 | 123456   |
|  3 | Lillian      | 小李                 | 123456   |
+----+--------------+----------------------+----------+
3 rows in set (0.00 sec)
```

## 6. 批量操作

当我们需要向 `user` 表插入多条数据的时候，可以循环调用我们在上面 `JDBCDemo2` 实例代码中封装的 `insert()` 方法，但这样的效率是非常低的。

`Statement` 对象有如下常用的用于批量操作的方法：

* `void addBatch(String sql)`：将给定的 SQL 命令添加到此 `Statement` 对象的当前命令列表中；
* `int[] executeBatch()`：将一批命令提交给数据库来执行，如果全部命令执行成功，则返回更新计数组成的数组；
* `void clearBatch()`：清空此 `Statement` 对象的当前 SQL 命令列表。

## 7. 小结

通过本小节的学习，我们了解了 JDBC 定义了客户端访问数据库的 API，不同厂商通过实现统一的 JDBC 接口，降低了程序员的学习成本和维护成本。

DriverManager 类用于注册驱动和获得连接；

Connection 是连接对象，它可以创建执行 SQL 语句的对象，还可以进行事务的管理；

Statement 是执行 SQL 语句的对象、ResultSet 是结果集对象，它是 `select` 语句查询的结果的封装。

为了防止 SQL 注入，推荐使用 `PrepareStatement` 对象来预编译 SQL 语句，对于内容相同，参数不同的 SQL，推荐使用 JDBC 的 `batch` 操作，可大大提高执行效率。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
