---
title: Node 新手课（16）打卡
published: 2020-05-25 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了图片上传，今天我们来讲打卡。

## 需求

用户填写文字内容，上传图片，发送后首页列表可以展示出来。

## 功能流程

上传图片后前端获得图片地址，再把用户输入的文字内容 content 和 image 发到后端，后端插入记录到 card 表。

## 代码

> routes/index.js

```js
//打卡
router.post('/card',auth, async (ctx, next) => {
  return indexApi.addCard(ctx)
})
```

> api/index.js

```js
const addCard = async(ctx) => {
    const userId = ctx.state.userId
    const content =ctx.request.body.content
    const image = ctx.request.body.image
    await cardModel.addCard(userId,content,image)
    ctx.body = resp.succeed()
    return
}
```

> model/cardmodel.js

```js
const addCard = async(userId,content,image)=>{
    let sql = "insert into `card` (user_id,content,image)values(?,?,?) "
    const insertData = await mysql.exec(sql,[userId,content,image])
    return insertData.insertId
}
```

## 运行

```sh
POST http://localhost:3000/card

{
    "content":"test",
    "image":"https://cdn.guojiang.club/Fo6qftIGMQq8uM7ANRdWs9Kskow5"
}
```


## 总结

打卡功能就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解点赞，继续加油吧，Let's go！
