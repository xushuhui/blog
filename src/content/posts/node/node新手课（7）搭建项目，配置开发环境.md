---
title: Node 新手课（7）搭建项目，配置开发环境
published: 2020-02-25 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲解了项目技术架构和项目展示，这节课我们来搭建项目架构，koa2 和 mysql

### koa2

#### 介绍

为什么使用 koa2，框架给我们节省很多时间，制定了流程和标准，降低了我们开发难度。使我们遵循标准，专心开发我们的业务，
写更少的代码完成更多的功能。

前端同学应该都经历过从 jquery 到 vue/react 的过程，vue/react 就是为了减少前端开发的工作量，让开发专注业务和实现功能。

框架封装了基本的工具和 API（解析请求参数，路由，session/cookie，日志，操作数据库，中间件），以及一些插件。

#### 安装

```sh
$ npm install koa-generator -g
$ Koa2 sir-koa
$ cd sir-koa && npm install

```

目录结构
~~~
sir-koa
├─bin                   可执行目录
│  ├─www                执行命令
│
├─node_modules          node 依赖包目录
│
├─routes                路由定义目录
│  ├─index.js           路由定义文件
│  └─...                更多
│
├─public                静态资源目录
│  ├─images             图片目录
│  ├─javascripts        js 文件目录
│  └─stylesheets        css 文件目录
│
├─views                 视图文件目录
│
├─app.js                应用启动文件

~~~
我们先安装 nodemon 包

```sh
$ npm i -g nodemon
```

现在我们启动项目

```sh
$ npm run dev
```

打开浏览器，输入`` `http://localhost:3000/` ``，项目就启动了。

### mysql

地址：https://dev.mysql.com/downloads/mysql/

#### windows 安装

点击 https://dev.mysql.com/downloads/windows/installer/，直接点击 download


然后下一步又下一步安装就可以了

#### linux 安装（推荐）

我们在 Centos7 系统下使用 yum 命令安装 MySQL

#### 安装

```sh

$ wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
$ rpm -ivh mysql-community-release-el7-5.noarch.rpm
$ yum update
$ yum install mysql-server
```

#### 设置权限和初始化

```sh
$ chown mysql:mysql -R /var/lib/mysql
$ mysqld --initialize
```

#### 启动 mysql

```sh
$ systemctl start mysqld
```

#### 连接 mysql

```sh
$ mysql

mysql> SHOW DATABASES;
+----------+
| Database |
+----------+
| mysql    |
| test     |
+----------+
2 rows in set (0.13 sec)

```

以上命令执行后会输出 mysql > 提示符，这说明你已经成功连接到 Mysql 服务器上，你可以在 mysql> 提示符执行 SQL 命令。

### redis

#### 安装和配置

以 centos 为例，我们直接通过 yum 安装 redis。

```sh
$ yum install redis
$ vim /etc/redis.conf
//把 daemonize no 改成 daemonize yes
$ redis-server /etc/redis.conf
```

现在 redis 就成功启动了

## 总结

搭建项目架构就已经完成了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解 koa 的路由，继续加油吧，Let's go！
