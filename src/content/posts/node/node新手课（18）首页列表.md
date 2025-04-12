---
title: Node 新手课（18）首页列表
published: 2020-06-17 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了点赞，今天我们来讲查询首页列表。

## 需求

用户打卡小程序的首页就以列表的形式展示所有用户打卡发布的图片和内容。用户可以点赞某一条打卡。

## 功能流程

查询数据库中 card 表的所有记录，后期可以扩展加入分页或分类、标签等功能。

## 代码

> routes/index.js

```js
router.get('/card', async (ctx, next) => {
  return indexApi.getCard(ctx)
})
```

> api/index.js

```js
const getCard = async(ctx) => {
    const list = await cardModel.getCards()
    ctx.body = resp.setData(list)
    return
}
```

> model/cardmodel.js

```js
const getCards = async()=>{
    let sql = "select id,content,image from `card` "
    const res = await mysql.exec(sql)
    return res
}
```

## 运行

```sh
POST http://localhost:3000/card

{
    "content":"content"
}
```

响应数据

```json
{
    "message": "OK",
    "data": [
        {
            "id": 1,
            "content": "111",
            "image": ""
        },
        {
            "id": 2,
            "content": "13012341234",
            "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1593516228269&di=b24a61eb023e7b2d324c77921960d1f7&imgtype=0&src=http%3A%2F%2Fa3.att.hudong.com%2F14%2F75%2F01300000164186121366756803686.jpg"
        }
    ]
}
```

## 总结

查询首页列表就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解项目部署上线，继续加油吧，Let's go！
