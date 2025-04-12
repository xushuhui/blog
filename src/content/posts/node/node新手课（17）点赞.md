---
title: Node 新手课（17）点赞
published: 2020-06-08 07:57:34
tags: ["Node"]
categories: ["Node"]
---


上节课我们讲了讲和打卡，今天我们来讲点赞。

## 需求

用户对信息列表中的信息进行点赞操作。

## 功能流程

前端传到后端信息 card_id，后端根据 card_id 和 token 解析出来的 user_id，查询 card_praise 表，如果记录存在说明已经点赞过了，返回给前端提示。否则插入 card_praise，再更新 card 表的 praise 字段的值。

## 代码

> routes/index.js

```js
//点赞
router.post('/praise', auth,async (ctx, next) => {
  return indexApi.praise(ctx)
})
```

> api/index.js

```js
const praise = async(ctx) => {
    const cardId = ctx.request.body.card_id
    const userId = ctx.state.userId
    let id =  await cardModel.getCardPraise(cardId,userId)
    if(id > 0){
        ctx.body = resp.fail(20001,"不要重复点赞")
        return
    }
    await cardModel.createCardPraise(cardId,userId)
    await cardModel.incrCardPraise(cardId)
    ctx.body = resp.succeed()
    return
}
```

> model/cardmodel.js

```js
const createCardPraise = async(cardId,userId)=>{
    let sql = "insert into `card_praise` (card_id,user_id)values(?,?) "
    const insertData = await mysql.exec(sql,[cardId,userId])
    return insertData.insertId
}
const getCardPraise = async(cardId,userId)=>{
    let sql = "select id from `card_praise` where card_id=? and user_id=?"
    const res = await mysql.queryOne(sql,[cardId,userId])
    return res.id
}
const incrCardPraise = async(cardId)=>{
    let sql = "update `card` set praise=praise+1 where id=?"
    const res = await mysql.exec(sql,[cardId])
    return res.affectedRows
}
```

## 运行

```bash
POST http://localhost:3000/praise

{
    "card_id":1
}
```


## 总结

点赞功能就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解查询首页列表，继续加油吧，Let's go！
