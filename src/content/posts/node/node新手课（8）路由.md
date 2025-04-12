---
title: Node 新手课（8）路由
published: 2020-03-10 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们来搭建项目架构，koa2 和 mysql。这节课我们讲解 koa2 的路由

### 路由

路由就是客户端的请求发送到服务端的的 url。我们打开浏览器输入 http://www.baidu.com 进入百度，这个网址就是路由。

> 根据功能划分路由，把路由和数据处理分离，符合单一职责的设计原则

上节课我们搭建的项目的 routes 目录就是用来存放所有路由文件。

我们在 sir-koa 目录下新建 app 目录，在 app 目录中新建 api 和 model 目录，分别作为控制层和模型层。

路由调用控制层的方法。控制层接收前端传过来的参数，调用模型层的方法进行数据库操作，处理后返回结果给前端。

入口文件 app.js 中调用了路由模块

```js
// routes
app.use(index.routes(), index.allowedMethods())
app.use(users.routes(), users.allowedMethods())
```

## 功能模块

功能分为用户模块和卡片模块，我们把列表每条信息当成一个卡片。

> routes/user.js

```js
const router = require('koa-router')()
const userApi = require('../app/api/user')

router.prefix('/user')
//微信登录
router.post('/wechatlogin', function (ctx, next) {
  userApi.wechatLogin(ctx)
})
// 手机号登录
router.post('/login', function (ctx, next) {
  userApi.login(ctx)
})
//更新个人信息
router.put('/info', function (ctx, next) {
  userApi.info(ctx)
})
//查询个人信息
router.get('/me', function (ctx, next) {
  userApi.me(ctx)
})

module.exports = router
```

> routes/index.js

```js

const indexApi = require('../app/api/index')
//首页列表
router.get('/card', async (ctx, next) => {
  indexApi.getCard(ctx)
})
//打卡
router.post('/card', async (ctx, next) => {
  indexApi.addCard(ctx)
})
//点赞
router.post('/praise', async (ctx, next) => {
  indexApi.praise(ctx)
})
```

### 接口文档

https://note.youdao.com/ynoteshare1/index.html?id=ae97d2c2bc96012409533ab7679061d2&type=note

## 总结

koa2 的路由就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解 nodejs 操作 mysql，继续加油吧，Let's go！
