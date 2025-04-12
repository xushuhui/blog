---
title: Node 新手课（15）图片上传
published: 2020-05-15 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了修改用户信息功能，今天我们来讲图片上传。

## 需求

用户在打卡功能中，选择图片上传。类似微信发朋友圈的功能。

## 功能流程

安装 koa-body 包，用来接收图片上传的文件流。然后保存到 public 目录。

## 代码

```sh
$ npm i koa-body --save
```

> app.js

```js
const koaBody = require('koa-body');

app.use(koaBody({
    multipart: true,
    formidable: {
        maxFileSize: 1200*1024*1024,	// 设置上传文件大小最大限制，默认2M
    }
}));
```

> routes/index.js

```js
router.post('/upload', async (ctx, next) => {
  return indexApi.upload(ctx)
})
```

> api/index.js

```js
  const file = ctx.request.files.file;
    const fileName = file.name;
    // 创建可读流
    const render = fs.createReadStream(file.path);
    let filePath = path.join('public/images',fileName);
    const fileDir = path.join('public/images');
    if (!fs.existsSync(fileDir)) {
      fs.mkdirSync(fileDir, err => {
        console.log('创建失败',err)
      });
      return
    }
    // 创建写入流
    const upStream = fs.createWriteStream(filePath);
    render.pipe(upStream);
    ctx.body = resp.succeed()
    return
```

## 运行

```sh
POST http://localhost:3000/upload

{
    "message": "OK",
    "data": {
        "url": "http://localhost:3000/public\\images\\f9a15976ab.jpg"
    }
}
```

## 总结

图片上传就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解打卡，继续加油吧，Let's go！
