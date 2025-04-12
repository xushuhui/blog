---
title: Node 新手课（2）Hello World
published: 2020-01-23 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲解了 node 环境的搭建，按照编程语言的惯例，我们首先开始个 HelloWorld 程序。

### 创建项目

创建 sir-node 文件夹，这个就是我们要开发的项目

```sh 

$ mkdir sir-node
$ cd sir-node

```

### 新建 hello.js

新建 hello.js，输入以下内容

```js
// 引入 Http 模块：我们可以使用 require 指令来载入 Http 模块。
const http = require("http")

// 创建服务器：服务器可以监听客户端的请求，类似于 Apache 、Nginx 等 HTTP 服务器。
http.createServer((request,response)=>{
    // 设置响应的 http 头部信息
    response.writeHead(200, {'Content-Type': 'text/plain'});
    // 设置响应的 http 数据
    response.end('Hello World\n');
}).listen(3000)

console.log('Server running at http://127.0.0.1:3000/');
```

### 启动 hello.js

```js
$ node hello.js
Server running at http://127.0.0.1:3000/
```

客户端可以使用浏览器或终端发送 HTTP 请求，服务器接收请求后返回响应数据


## 总结

现在我们成功开启了一个 http 服务，写了个 helloworld。

这是一个良好的开始，你学得怎么样呢？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们讲解 nodejs 和 javascript 区别，继续加油吧，Let's go！
