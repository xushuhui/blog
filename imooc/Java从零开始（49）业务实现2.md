# 实战 - 业务实现 2



这一小节，是 Java 基础教程的最后一节，很感谢大家能够坚持看到这里。本小节我将带领大家优化用户鉴权服务，并完成商品模块的实现。为了检验大家的学习成果，分类模块的实现将交给大家自行来完成。



## 1. 用户密码加密



上一小节的最后，我们提到用户鉴权服务是需要优化的。大家可以看到我们数据库存储的是明文密码，这是非常不推荐的，在实际的项目中，明文存储用户的密码是非常不安全的，也是不负责任的行为。我们在设计 `imooc_user`表时，给`password`设置的类型为固定长度类型`char(32)`，32 位正好是`MD5`算法加密后的长度。



本系统使用 `MD5` 算法对密码进行加密，下面在 `util`包下新建一个 `MD5Util`类并写入如下内容（可直接复制粘贴代码）：



```java
package com.colorful.util;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public class MD5Util {

    public static String md5(String source) {
        StringBuilder stringBuilder = new StringBuilder();
        try {
            MessageDigest messageDigest = MessageDigest.getInstance("MD5");
            // 将一个byte数组进行加密操作，返回的是一个加密的byte数组，二进制的哈西计算，md5加密的第一步
            byte[] digest = messageDigest.digest(source.getBytes());
            for (byte b : digest) {
                int result = b & 0xff;
                // 将得到的int类型的值转化为16进制的值
                String hexString = Integer.toHexString(result);
                if (hexString.length() < 2) {
                    //系统会自动把0省略，所以添加0
                    stringBuilder.append("0");
                }
                stringBuilder.append(hexString);
            }
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }
        return stringBuilder.toString();
    }

    public static void main(String[] args) {
        String password = "123456";
        String s = MD5Util.md5(password);
        System.out.println(s);
    }

}
```



在主方法中，我们编写了调用`md5()`加密方法的逻辑，运行代码，屏幕上得到`123456`加密后的字符串：



```java
e10adc3949ba59abbe56e057f20f883e
```



下面我们将`imooc_user`表中存储的明文密码，更新为上面的结果，大家可以使用`SQL`语句来进行更新：



```java
 UPDATE `imooc_user` SET `password` = 'e10adc3949ba59abbe56e057f20f883e' WHERE `id` = 1;
```



这里我直接通过 `MySQL` 客户端进行更新，如下是操作过程的截图：


![图片描述](//img.mukewang.com/wiki/5f30bd8c09d5756328121236.jpg)

数据库存储的密码更新后，我们就无法直接通过原本的验证逻辑来验证密码了，需要修改用户鉴权逻辑 —— 将用户输入的密码加密后，再与数据库的密码进行对比。那么这段逻辑要写在`service`层还是`dao`层呢？答案肯定是`service`层，此时`service`层用于处理业务的特性得到了体现，修改`UserService`下的`login`方法，将参数`password`加密：



```java
public User login(String username, String password) {
    String md5Password = MD5Util.md5(password);
    return userDAO.selectByUserNameAndPassword(username, md5Password);
}
```



再次启动应用程序，验证改写的逻辑是否正确：



![](//img.mukewang.com/wiki/5f2a5af00909cade10280894.jpg)

至此，我们就完成了对用户鉴权服务的优化。



## 2. 控制台（仪表盘）



用户登录成功后，应该显示控制台面板，我们下面称之为仪表盘，它主要包含 3 个选项，分别是管理商品、管理分类以及退出登录。下面我们编写一个`dashboard()`方法，该方法用来打印仪表盘的相关操作提示，以及根据用户的输入来执行相应的操作。如下是部分代码：



```java
/**
 * 主流程方法
 */
public static void run() {
    // ... 已省略前面的鉴权代码
    // 登录成功后，跳转到仪表盘页面
    dashboard();
}

/**
 * 仪表盘操作
 */
private static void dashboard() {
    Scanner scanner = new Scanner(System.in);
    int code1 = 0, code2 = 0;
    while (true) {
        printDashboardTips();
        code1 = scanner.nextInt();
        if (code1 == 0) {
            System.out.println("您已退出登录");
            break;
        }
        switch (code1) {
            case 1:
                System.out.println("正在查询商品列表...");
                // TODO 实现商品模块
                break;
            case 2:
                System.out.println("正在查询分类列表...");
                // TODO 实现分类模块
                break;
            default:
                System.out.println("不存在您输入的选项，请重新输入");
                break;
        }
    }
}

/**
 * 输出仪表盘操作提示
 */
private static void printDashboardTips() {
    System.out.println("请输入对应数字以进行操作：");
    System.out.println("（1. 管理商品 | 2. 管理分类 | 0. 退出登录）");
}
```



我们把向控制台输出的操作提示，封装成了一个方法`printDashboardTips()`，这样使代码更简洁易读。



在`dashboard()`方法内部，实例化了一个`Scanner`类，初始化的`code1`变量接收用户的输入，根据输入的数值用来操作仪表盘，关于`code2`变量，我们将在实现商品模块代码的时候使用。紧接着有一个`while`循环，其条件始终为`true`，当用户输入的`code`登录 `0` 的时候，就跳出循环，也就是退出了应用程序。



完成上面的代码编写后，我们启动应用程序，来验证一下：



![](//img.mukewang.com/wiki/5f2a5b16099bfc3409000640.jpg)

至此，我们已实现展示仪表盘以及退出登录的代码编写。



## 3. 商品模块实现



### 3.1 商品管理主流程



当用户输入的`code1`变量为数字 `1` 的时候，就要显示商品管理相关的操作。我们再封装一个`printGoodsListTips()`方法，用于打印商品管理模块的相关操作提示。方法的代码如下：



```java
/**
 * 输出商品列表页操作提示
 */
private static void printGoodsListTips() {
    System.out.println("请输入对应数字以进行操作：");
    System.out.println("（1. 新增商品 | 2. 编辑商品 | 3. 查看商品详情 | 4. 删除商品 | 5. 搜索商品 | 6. 按分类查询商品 | 0. 返回上一级菜单）");
}
```



向屏幕打印这些提示后，下面还是一个条件始终为`true`的`while`循环，当用户输入的`code`登录 `0` 的时候，就跳出当前层循环，也就是返回上一级仪表盘的菜单。



已知了商品管理模块的所有操作，下面我们在`switch(code1)`的`case 1`条件分支加入如下逻辑代码（部分伪代码）：



```java
case 1:
    while (true) {
        System.out.println("正在查询商品列表...");
        // TODO 查询并显示商品列表
        printGoodsListTips();
        code2 = scanner.nextInt();
        if (code2 == 0) {
            // 返回上一级，即跳出本层循环
            System.out.println("返回上一级");
            break;
        }
        switch (code2) {
            case 1:
                System.out.println("新增商品");
                break;
            case 2:
                System.out.println("编辑商品");
                break;
            case 3:
                System.out.println("商品详情");
                break;
            case 4:
                System.out.println("删除商品");
                break;
            case 5:
                System.out.println("搜索商品");
                break;
            case 6:
                System.out.println("按分类查询");
                break;
            default:
                System.out.println("不存在您输入的选项，请重新输入");
        }
    }
    break;
```



上面我们提到，`code2`变量用于接收用户对于管理商品操作的输入，此处又是一个`switch case`结构，每一个条件分支，都对应到用户输入的数字，如果用户输入的数字找不到对应的分支，那么就重复执行循环体中的代码。



接下来我们就要实现这些操作。



### 3.2 查询商品列表



在`dao`包下新建一个`GoodsDAO`类，并写入一下内容：



```java
package com.colorful.dao;

import com.colorful.model.Goods;
import com.colorful.util.JDBCUtil;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.util.ArrayList;
import java.util.List;

public class GoodsDAO {

    private Connection connection = null;
    private PreparedStatement preparedStatement = null;
    private ResultSet resultSet = null;
    boolean executeResult;
    
    public List<Goods> selectGoodsList() {
        List<Goods> goodsList = new ArrayList<>();
        try {
            // 获得链接
            connection = JDBCUtil.getConnection();
            // 编写 SQL 语句
            String sql = "SELECT `id`, `name`, `price` FROM `imooc_goods` where `delete_time` is null";
            // 预编译 SQL
            preparedStatement = connection.prepareStatement(sql);
            resultSet = preparedStatement.executeQuery();
            while (resultSet.next()) {
                Goods goods = new Goods();
                goods.setId(resultSet.getInt("id"));
                goods.setName(resultSet.getString("name"));
                goods.setPrice(resultSet.getDouble("price"));
                goodsList.add(goods);
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 释放资源
            JDBCUtil.release(resultSet, preparedStatement, connection);
        }
        return goodsList;
    }

}
```



`selectGoodsList()`方法就用于查询商品列表（由于数据量不大，此处我没有对列表数据进行分页查询，大家也可以自行加入）。



在`service`包下新建`GoodsService`，并调用`dao`层下封装好的方法：



```java
package com.colorful.service;

import com.colorful.dao.GoodsDAO;
import com.colorful.model.Goods;

import java.util.List;

public class GoodsService {

    private final GoodsDAO goodsDAO = new GoodsDAO();

    /**
     * 获取商品列表
     * @return 商品列表
     */
    public List<Goods> getGoodsList() {
        return goodsDAO.selectGoodsList();
    }

}
```



这样，我们就完成了查询商品列表的服务层代码编写。



### 3.3 删除商品



新增商品、删除商品、查看商品详情等功能都是简单的`SQL`语句，这里不再具体写出实现，大家可以参考[源码](https://github.com/fujiale33/goods-cms)自行实现。但关于删除商品，我要特殊说明一下。对于实际的项目，往往不用对数据执行`DELETE`操作，对于数据的删除往往是更新操作，这也是我们设置了一个公用字段`delete_time`的意义，当这个`delete_time`字段不为`null`的时候，才会被查询出来。在`GoodsDAO`类下，新增如下方法：



```java
public boolean deleteGoodsById(Integer id) {
    try {
        // 获得链接
        connection = JDBCUtil.getConnection();
        // 编写 SQL 语句
        String sql = "UPDATE `imooc_goods` set `delete_time` = ? WHERE id = ?";
        // 预编译 SQL
        preparedStatement = connection.prepareStatement(sql);
        preparedStatement.setTimestamp(1, new Timestamp(System.currentTimeMillis()));
        preparedStatement.setInt(2, id);
        executeResult = preparedStatement.execute();
    } catch (Exception e) {
        e.printStackTrace();
    } finally {
        // 释放资源
        JDBCUtil.release(preparedStatement, connection);
    }
    return executeResult;
}
```



大家可以看到，我们的代码实现没有使用`DELETE`语句，而是使用了`UPDATE`语句，更新了指定`id`记录的`delete_time`字段为系统当前时间。



`dao`层方法编写完成后，就可以在`service`层调用该方法了：



```java
/**
 * 删除商品
 * @param id 商品id
 */
public void removeGoodsById(Integer id) {
    goodsDAO.deleteGoodsById(id);
}
```



### 3.4 搜索商品



除了删除商品的实现，搜索商品的实现我们也要特殊讲解一下。上面我们提到，由于商品的数据量不大，在查询商品列表时，没有使用`LIMIT`关键字进行分页查询。正是由于数据量不大的原因，对于搜索商品，我们没有使用`LIKE`关键字进行模糊查询，而是使用`Stream API`直接对商品列表进行过滤，希望通过这里的实现来协助让大家理解`Stream API`，直接在`GoodsService`下添加如下方法：



```java
/**
 * 根据商品名称搜索商品
 * @param name 商品名称
 * @return 商品列表
 */
public List<Goods> searchGoodsByName(String name) {
    List<Goods> goodsList = this.getGoodsList();
    return goodsList.stream().filter(
            goods -> goods.getName().contains(name)
    ).collect(Collectors.toList());
}
```



该方法先是调用了`getGoodsList()`方法获取了商品列表，然后使用`Stream API`中的`filter()`中间操作，对商品进行过滤，`filter()`接收一个**断言型接口**，由于是一个函数式接口，我们可通过`lambda`表达式来进行表示。最后调用`collect()`终止操作，将流转化为列表。



服务层的接口完成后，大家就可以在对应的`case`分支编写的具体的逻辑了，每个分支的逻辑大体相同，主要是接收用户的输入，以及服务层方法的调用。大家可参考[`github`仓库的源码](https://github.com/fujiale33/goods-cms)来补全自己的代码。



## 4. 作业 - 分类模块实现



上面，我们已经实现了较为复杂的商品模块，对于分类模块的实现也大同小异，甚至更加简单，剩下的功能 ——分类的增删改查就交由同学们自行实现。希望大家能够按照我们项目的架构，将合理的代码写到合适的位置，对每个功能点都要将细节考虑周全，这将有助于降低大家后续对框架学习的上手成本。



## 5. 小结



通过实战阶段的学习，我们知道了数据表中的密码字段，是不能够明文存储的，通常使用一些加密算法进行加密，也复习了`switch case`条件结构的使用，对于商品模糊查询，我们使用了 Java 8 中的 `Stream API`。



中间还讲解了项目的分层技术、MySQL 的增删改查操作、JDBC API 的封装与使用以及`Scanner`类的使用等知识，实际的项目基本不会使用`Scanner`来与用户进行交互，都是通过优美的前端界面与用户进行交互的，建议大家可以去看看[Lin CMS](http://face.cms.talelin.com)的示例`demo`，它是一个能够达到企业级应用标准的内容管理系统开发框架。



当然，想要上手使用`Lin CMS`，大家还有很长的一段路要走，但是请记住，莫要浮空建高楼，Java 的基础知识在任何时候都是不能忽视的，希望大家反复学习。Java 基础的学习到此也就结束了，再次感谢大家能够坚持看完！






### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)