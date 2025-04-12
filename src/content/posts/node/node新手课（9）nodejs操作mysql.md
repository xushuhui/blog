---
title: Node 新手课（9）nodejs 操作 mysql
published: 2020-03-17 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了 koa2 的路由，今天我们来讲解 nodejs 操作 mysql

## 导入 sql 文件

sql 文件地址
https://github.com/guojiangclub/sir-node/blob/master/sir.sql

安装 navicat，打开 navicat, 选择连接 - 新建连接，输入本地数据库地址，用户名和密码，成功连接到数据库。

右键点击连接，选择新建数据库 sir

右键点击 sir 数据库，选择运行数据库，导入 sql 文件。

## mysql 测试

新建 mysql-test 目录，新建 index.js

```js
const mysql = require('mysql')

// 创建链接对象
 const con = mysql.createConnection({
    host: 'localhost',
    user: '数据库用户名',
    password: '数据库密码',
    port: '数据库端口，默认 3306',
    database: 'sir'
})

// 开始连接
 con.connect()

// 执行 sql 语句
const sql = `insert into test (title, content, author) values (' 标题 C', ' 内容 C','zhangsan');`
con.query(sql, (err, result) => {
    if (err) {
        console.error(err)
        return
    }
    console.log(result)
})

// 关闭连接
 con.end()
```

执行如下操作

```sh
$ cd mysql-test/
$ npm init -y
$ npm i mysql --save
$ node index.js
```

打开 sir 数据库，查看 test 表，现在已经新增了一条记录。

## 项目添加 mysql 包

在 sir-koa 目录执行

```sh
$ npm i mysql --save
```

在 sir-koa 目录下新建 config 目录，用来存放所有项目配置文件。

> config/mysql.js

```js
const Config = {
    host     : 'localhost',
    user     : '数据库用户名',
    password : '数据库密码',
    port:3306,
    database : 'sir',

}
module.exports= {
    Config
}
```

model 目录中新建 mysql.js，用来在项目中调用 mysql

> model/mysql.js

```js
const mysql = require('mysql')
const { Config } = require('../../config/mysql')

// 创建链接对象
 const con = mysql.createConnection(Config)

// 开始链接
 con.connect()

// 统一执行 sql 的函数
 function exec(sql,params) {
    const promise = new Promise((resolve, reject) => {
        con.query(sql,params, (err, result) => {
            if (err) {
                reject(err)
                return
            }
            resolve(result)
        })
    })
    return promise
}

module.exports = {
    exec,
    escape: mysql.escape
}
```

## 总结

nodejs 操作 mysql 就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解微信小程序注册和登录，继续加油吧，Let's go！
