---
title: Node 新手课（3）nodejs 和 javascript 区别
published: 2020-01-27 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们成功用开启了一个 http 服务，写了个 helloworld，今天我们讲解下 nodejs 和 javascript 区别。

### ECMAScript

* 定义了语法，写 javascript 和 nodejs 必须遵守它的语法规范
* 变量定义、循环、判断、函数
* 原型和原型链、作用域和闭包、异步
* 不能操作 DOM，不能监听 click 事件，不能发送 ajax 请求
* 不能处理 http 请求，不能操作文件
* 如果只有 ECMAScript，几乎做不了任何实际的项目

### Javascript

* Javascript 由 ECMAScript 语法规范 +Web Api 组成，两者结合即可完成浏览器端的操作
* Javascript 可以执行 DOM 操作，BOM 操作，事件绑定，Ajax 等

### Nodejs

*  ECMAScript 语法规范 +Nodejs Api 组成，两者结合即可完成服务器端的操作
* Nodejs 可以处理 http，处理文件等，具体参考 http://nodejs.cn/api/

### 总结

* ECMAScript 和 JavaScript 的关系是，前者是后者的规格，后者是前者的一种实现。

按类和对象比喻，ECMAScrpt 是类，Javascript 是对象

### 前端和服务端的区别

客户端可能有多个不同的版本，统一从服务端获取数据。


#### 稳定性

* 当个客户端可以意外挂掉，服务端不能，因为服务端是给无数个客户端提供数据，一旦宕机，影响非常大。
* nodejs 一般使用 PM2 管理进程，断开的进程自动重启

#### 内存和 CPU

* 客户端独占一个浏览器，无需考虑内存和 CPU 占用情况
* 服务端需要承载很多客户端请求，处理由客户端的大批量数据，内存和 CPU 比较稀缺
* 课程会使用 stream 写入日志，使用 redis 存储 session

#### 日志记录

* 客户端会参与日志写入，但只是发起方，不考虑后续影响和结果
* 服务端详细地记录日志，存储日志，分析日志
* 课程会解决多种日志记录方式，以及如何分析日志

#### 安全

* 服务端可能会遭受各种恶意攻击和误操作，比如越权操作，SQL 注入等，前端就很少遇到
* 课程会讲解登录验证，预防 xss 攻击和 SQL 注入。

#### 集群和服务拆分

* 产品发展速度快，流量可能会迅速增加。我们要通过扩展服务器集群和服务拆分承载更大流量。
* 课程虽然是单台机器部署，但是支持服务拆分，后续也会增加这部分内容。

### 相关文档

* https://es6.ruanyifeng.com/

### 总结

你了解了 nodejs 和 javascript 区别，有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解 nodejs 处理 http 请求，继续加油吧，Let's go！
