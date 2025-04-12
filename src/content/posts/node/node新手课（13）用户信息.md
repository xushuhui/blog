---
title: Node 新手课（13）用户信息
published: 2020-05-01 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了手机号登录，今天我们来讲展示用户信息接口。

## 需求

用户查看自己的信息，如头像，昵称，手机号等。

## 功能流程

通过用户凭据 token 获取到用户 id，在数据库 user 表中找到用户记录，返回前端用户的昵称和手机号。

## 代码

> routes/user.js

```js
router.get('/me', auth,function (ctx, next) {
  return userApi.me(ctx)
})
```

> api/user.js

```js
const me = async(ctx) => {
    const userId = ctx.state.userId
    const user = await userModel.getUserById(userId)
    ctx.body = resp.setData(user)
    return
}
```

> model/usermodel.js

```js
const getUserById = async(id)=>{
    let sql = "select nickname,phone from user where id=?"
    let user = await mysql.queryOne(sql,[id])
    return user ? user : []
}
```

## 运行

```sh
GET http://localhost:3000/user/me

```

响应数据

```json
{
    "message":"OK",
    "data":{
        "nickname":"昵称","phone":"13012341236"
    }
}
```

## 总结

用户信息展示接口就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解修改用户信息，继续加油吧，Let's go！
