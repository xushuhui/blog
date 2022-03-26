---
title: Java 从零开始（47）数据库设计
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# 实战 - 数据库设计

经过上一小节的需求分析，我们将系统分为了鉴权模块、商品模块和分类模块，本节中，我们将围绕功能模块，进行数据库设计。你将学习到实际开发中的一些数据库设计技巧。请确保在你的开发环境下，已经准备好了一个 `MySQL`数据库。

## 1. 创建库

首先，我们先给商品管理项目创建一个数据库，命名为`imooc_goods_cms`，`cms`为`Content Management System`（内容管理系统）的缩写。

链接数据库，执行以下 SQL：

```java
CREATE DATABASE `imooc_goods_cms` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';
```

![](https://xushuhui.gitee.io/image/imooc/5f17cc230a41dda319401034.jpg)

## 2. 创建表

### 2.1 通用字段

有一定数据库设计经验的同学都知道，一个数据库中的数据表都会有一些通用字段。关于通用字段有哪些、如何命名以及如何选定字段类型都有一定的套路，不仅受个人开发习惯的影响，也受到团队开发规范的约束。

该项目，所有的数据表都包含以下字段：

|字段名称|字段类型|字段长度|允许为 NULL|默认值|注释|
|--------|--------|--------|-----------|------|----|
|id         |int      |0|否|无               |表主键，自增长|
|create_time|timestamp|-|否|CURRENT_TIMESTAMP|创建时间      |
|update_time|timestamp|-|否|CURRENT_TIMESTAMP|更新时间      |
|delete_time|timestamp|-|是|无               |删除时间      |

另外，除了通用字段，该项目中的所有表都以`imooc`开头。

### 2.2 用户表

由于该系统鉴权模块使用用户名和密码进行鉴权，需要设计一个用户表`imooc_user`，用户表包含如下字段（已省略通用字段）：

|字段名称|字段类型|字段长度|允许为 NULL|默认值|注释|
|--------|--------|--------|-----------|------|----|
|nickname|varchar|50|是|空字符串|昵称  |
|username|varchar|50|否|无      |用户名|
|password|char   |32|否|无      |密码  |

> **Tips**：`password`字段的类型为`char`，这是因为我们在实现代码时，将使用`MD5`算法进行加密，加密后的密码长度为`32`，因此使用定长字符串更加节省存储空间。

链接`MySQL`并选择`imooc_goods_cms`数据库，执行以下 SQL：

```java
-- ----------------------------
-- Table structure for imooc_user
-- ----------------------------
DROP TABLE IF EXISTS `imooc_user`;
CREATE TABLE `imooc_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT "昵称",
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT "用户名",
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT "密码",
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of imooc_user
-- ----------------------------
INSERT INTO `imooc_user` VALUES (1, '小慕', 'admin', '123456', '2020-07-20 16:53:19', '2020-07-20 16:53:19', NULL);
```

![](https://xushuhui.gitee.io/image/imooc/5f20d5270a3f719519401034.jpg)

### 2.3 商品表

商品模块需要一个商品表`imooc_goods`，商品表包含如下字段（已省略通用字段）：

|字段名称|字段类型|字段长度|允许为 NULL|默认值|注释|
|--------|--------|--------|-----------|------|----|
|name       |varchar|100  |否|无      |商品名|
|description|varchar|255  |是|空字符串|简介  |
|category_id|int    |11   |是|0       |分类 id|
|price      |dicimal|10, 2|否|无      |价格  |
|stock      |int    |11   |是|0       |库存  |

链接`MySQL`并选择`imooc_goods_cms`数据库，执行以下 SQL：

```java
-- ----------------------------
-- Table structure for imooc_goods
-- ----------------------------
DROP TABLE IF EXISTS `imooc_goods`;
CREATE TABLE `imooc_goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `category_id` int(11) NULL DEFAULT 0,
  `price` decimal(10, 2) NOT NULL,
  `stock` int(11) NULL DEFAULT 0,
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of imooc_goods
-- ----------------------------
INSERT INTO `imooc_goods` VALUES (1, '测试商品 1', '', 0, 12.30, 3, '2020-07-20 16:53:19', '2020-07-20 16:53:19', NULL);
INSERT INTO `imooc_goods` VALUES (2, '测试商品 2', '', 0, 33.20, 10, '2020-07-20 17:17:53', '2020-07-20 17:17:53', NULL);
INSERT INTO `imooc_goods` VALUES (3, '测试商品 3', '', 0, 20.00, 50, '2020-07-20 17:18:09', '2020-07-20 17:18:09', NULL);
```

### 2.4 分类表

分类模块需要一个商品表`imooc_category`，分类表包含如下字段（已省略通用字段）：

|字段名称|字段类型|字段长度|允许为 NULL|默认值|注释|
|--------|--------|--------|-----------|------|----|
|name       |varchar|100|否|无      |商品名|
|description|varchar|255|是|空字符串|简介  |

链接`MySQL`并选择`imooc_goods_cms`数据库，执行以下 SQL：

```java
-- ----------------------------
-- Table structure for imooc_category
-- ----------------------------
DROP TABLE IF EXISTS `imooc_category`;
CREATE TABLE `imooc_category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
```

![](https://xushuhui.gitee.io/image/imooc/5f20d54c09d1a5fb19030466.jpg)

至此，我们的数据库和表已经建立完成。

## 3. 小结

本小节，我们一起创建了数据库和表，表结构比较简单，大家可以去我的 [代码仓库](https://github.com/colorful3/goods-cms/blob/master/src/main/resources/imooc_goods_cms.sql) 找到`SQL`文件，并直接在你的数据库中执行，下一小节，我们将着手业务代码的实现。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
