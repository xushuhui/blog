---
title: Node 新手课（12）手机登录
published: 2020-04-22 07:57:34
tags: ["Node"]
categories: ["Node"]
---


上节课我们讲了微信小程序登录，今天我们来讲绑定手机号。

## 需求

用户微信登录后绑定自己的手机号。

## 功能流程

通过用户凭据 token 获取到用户 id，在数据库 user 表中找到用户记录，把前端传来的手机号更新到记录中。

## 中间件

sir-koa 目录下新建 middleware 目录，存放项目的中间件。

> middleware/auth.js

```js
const response = require('../app/core/response')
const Redis = require('../app/model/redis')
const redis = new Redis()
const resp = new response()
module.exports = async (ctx, next) => {
     let token = ctx.request.headers["token"];
    // 解码
    let user = await redis.get(token)
    if (!token || !user) {
        //过期
        ctx.body = resp.fail(10001, "请重新登录")
        return
    }

    ctx.state.userId = user.userId
    // 未过期
    await next();
}
```

> middleware/jsonHeader.js

```js
module.exports = async (ctx, next) => {
    ctx.response.type = 'application/json';
    await next();
  }
```

绑定中间件。

> app.js

```js
const jsonHeader = require('./middleware/jsonHeader')
app.use(jsonHeader);
// routes
```

> routes/user.js

```js
const auth = require('../middleware/auth')
// 手机号登录
router.post('/login', auth,function (ctx, next) {
  return userApi.login(ctx)
})
```

## 手机号登录

客户端传手机号给到服务端，服务端更新用户信息。

> model/usermodel.js

```js
const updateUser = async(userId,phone)=>{
    let sql = "update `user` set phone=? where id=?"
    return await mysql.exec(sql,[phone,userId])
}
```

> api/user.js

```js
const login = async(ctx) => {
    const phone = ctx.request.body.phone
    const userId = ctx.state.userId
    await userModel.updateUser(userId,phone)
    ctx.body = resp.succeed()
    return
}
```

## 运行

```sh
POST http://localhost:3000/user/login

{
    "phone":""
}
```


## 总结

绑定手机号就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解用户信息展示接口，继续加油吧，Let's go！
