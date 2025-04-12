---
title: Node 新手课（11）微信小程序登录
published: 2020-04-07 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了 nodejs 操作 redis，今天我们来讲解微信小程序登录。

## 代码

微信服务端接口文档地址

https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

> 简而言之，小程序前端调用 wx.login 获取 code, 传给服务端。服务端携带 code 和 appid，secret 发给微信服务器，获取到用户 openid 和 sessionkey。

新增小程序的配置文件

> routes/user.js

```js
router.post('/wechatlogin',function (ctx, next) {
  return userApi.wechatLogin(ctx)
})
```

> config/miniapp.js

```js
const appid = ''
const secret = ''

module.exports={
    appid,secret
}
```

api/user.js 文件，里面写用户 api 的相关方法。model/usermodel.js 文件，里面写用户模型相关方法。

> api/user.js

```js
const miniappConfig = require('../../config/miniapp')
const koa2Req = require('koa2-request')
const response = require('../core/response')
const userModel = require('../model/usermodel')
const request = require('koa2-request')
const { md5 } = require('../util/cryp')
const Redis = require('../model/redis')
const redis = new Redis()
const resp = new response()

const wechatLogin = async (ctx) => {

    const code = ctx.request.body.code
    const body = await wxUser(code)

    if (body.errcode > 0) {
        ctx.body = resp.fail(10001, body.errmsg)
        return
    }
    let userId = await userModel.getUserByOpenid(body.openid)
    if(userId == 0){
        userId = await userModel.createUser(body.openid)
    }
    const token = md5(body.openid)
    redis.set(token,{ "openid": body.openid,"userId":userId })
    ctx.body = resp.setData({ "token": token })
    return
}
const wxUser = async (code) => {
    const appid = miniappConfig.appid
    const secret = miniappConfig.secret
    const loginUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code"
    const wxresult = await koa2Req(loginUrl).catch((err) => {
        console.log("err", err);
    });
    return JSON.parse(wxresult.body)
}
```

> model/usermodel.js

```js
const mysql = require("./mysql")

const getUserByOpenid = async(openid)=>{
    let sql = "select id from user where openid=?"
    return await mysql.exec(sql,[openid])
}
const createUser= async(openid)=>{
    let sql = "insert into `user` (openid)values(?) "
    const insertData = await mysql.exec(sql,[openid])
    return insertData.insertId
}
module.exports = {
    getUserByOpenid,createUser
}
```

## 运行

```sh
POST http://localhost:3000/user/wechatlogin

{
    "code":""
}
```

token 是客户端调用用户相关接口需要携带的凭据，服务端通过 token 识别用户身份信息，相当于 session。

## 总结

微信小程序登录就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解手机号登录，继续加油吧，Let's go！
