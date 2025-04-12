---
title: Node 新手课（4）nodejs 处理 http 请求
published: 2020-02-01 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们成功用开启了一个 http 服务，写了个 helloworld，今天我们讲解下 nodejs 是如何处理 http 的 get 和 post 请求。

我们先新建个 http-test 目录，该目录主要是用做测试 nodejs 代码，现在切换到这个目录

```js
$ mkdir http-test
$ cd http-test/
```

### get 请求

* 客户端向服务端获取数据，如查询用户信息
* 通过 querystring 传递数据，如 a.html?id=1&status=1
* 浏览器输入网址直接访问，就是发送 get 请求

#### 代码示例

我们新建个 get.js 文件，编辑以下内容。

```js
const http = require("http")
const querystring = require('querystring')

http.createServer((request,response)=>{
    const url =request.url // 获取完整 url
    console.log("method ",request.method) //GET
    console.log("url ",url)
    const query = querystring.parse(url.split("?")[1])// 解析 querystring
    response.end(JSON.stringify(query));// 把 querystring 返回到页面
}).listen(3000)

console.log('Server running at http://127.0.0.1:3000/');
```

启动 get.js

```js
$ node get.js
Server running at http://127.0.0.1:3000/
```

打开浏览器输入 http://127.0.0.1:3000/user/list?i=1&word=abc


控制台也打印出了 method 和 url

### post 请求

* 客户端向服务端发送数据，如用户注册
* 浏览器无法直接模拟 post 请求，需要写 js 代码或通过接口测试工具，比如 postman

我们新建个 post.js 文件，编辑以下内容。

#### 代码示例

```js
const http = require("http")

http.createServer((request,response)=>{
    if(request.method === "POST"){
        // 数据格式
        console.log('content-type',request.headers['content-type']);
        // 接收数据
        let postData = ""
        // 客户端请求传输数据的时候触发
        request.on('data',chunk=>{
            postData += chunk.toString()
        })
        // 客户端请求结束的时候触发
        request.on('end',()=>{
            console.log(postData)
            response.end("post end ...")// 异步方法，在这里返回
        })

    }

}).listen(3000)

console.log('Server running at http://127.0.0.1:3000/');
```

启动 post.js

```js
$ node post.js
Server running at http://127.0.0.1:3000/
```

用 postman 模拟发送 post 请求
结果如下


### http 请求

现在综合处理 get 和 post 请求，新建个 http.js，编辑以下内容。

```js
const http = require("http")
const querystring = require('querystring')

http.createServer((request,response)=>{
    const url =request.url // 获取完整 url
    const path = url.split('?')[0]
    const query = querystring.parse(url.split("?")[1])// 解析 querystring
    const method = request.method

    // 设置返回格式为 json
    response.setHeader('content-type','application/json')
    // 返回的数据
    const resData = {
        method,url,path,query
    }
    if(method === 'GET'){
        response.end(JSON.stringify(resData))
    }
    if(method === 'POST'){
        let postData = ''
        request.on('data',chunk=>{
            postData += chunk.toString()
        })
        request.on('end',()=>{
            resData.postData = postData
            response.end(JSON.stringify(resData))
        })
    }
}).listen(3000)

console.log('Server running at http://127.0.0.1:3000/');
```

启动 http.js

```sh
$ node post.js
Server running at http://127.0.0.1:3000/
```

现在不管是 get 请求还是 post 请求，我们都可以处理了。

## 总结

今天的内容你掌握了吗？有问题欢迎到群里和志同道合的小伙伴一起交流。

下节课我们开始讲解 callback，promise，async 和 await，继续加油吧，Let's go！
