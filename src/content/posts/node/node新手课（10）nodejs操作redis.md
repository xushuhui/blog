---
title: Node 新手课（10）nodejs 操作 redis
published: 2020-03-27 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲了 nodejs 操作 mysql，今天我们来讲解 nodejs 操作 redis

## 安装 redis 包

在 sir-koa 目录下执行

```sh
$ npm i ioredis --save
```

## redis 配置和操作

> config/redis.js

```js
const Config = {
    host: "127.0.0.1", // Redis host
    port: 6379, // Redis port
    password: "",
    db: 0,
}
exports = {
    Config
}
```

> model/redis.js

```js
const ioRedis = require('ioredis');
const { Config } = require('../../config/redis') //redis 配置
class Redis {
    constructor() {
        this.redis = new ioRedis(Config);//创建 redis 连接
    }
    async get(key) {
        const data = await this.redis.get(key);
        return JSON.parse(data);
    }
    async set(key, value) {
        await this.redis.set(
            key,
            JSON.stringify(value)
        );
    }
    async destroy(key) {
        return await this.redis.del(key);
    }
}

module.exports = Redis;

```

现在我们可以在项目中使用 redis 做数据缓存了。当然只是针对简单的字符串操作，你也可以拓展，加入对其他数据结构的操作方法

## 总结

nodejs 操作 redis 就讲完了，你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解微信小程序登录，继续加油吧，Let's go！
