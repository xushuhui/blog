---
title: Node 新手课（5）callback-promise 和 async-await
published: 2020-02-10 07:57:34
tags: ["Node"]
categories: ["Node"]
---



上节课我们讲解了 nodejs 处理 http 请求，今天我们讲下 callback, promise 和 async, await

### callback

js 代码处理逻辑都是异步，会造成 callback 层层嵌套。简称：回调地狱，非常不好维护，我们举例说明下
我们在 sir-node 中新建个目录 callback-test, 里面新建个 callback.js

```sh
$ mkdir callback-test && cd callback-test/
$ touch callback.js
```

callback.js

```js
const fs = require('fs')
const path  = require('path')
//callback 方式获取文件内容
function getFile(fileName,callback){
    const fullName = path.resolve(__dirname,"",fileName)
    fs.readFile(fullName,(err,data)=>{
        if(err){
            console.error(err)
            return
        }
        callback(JSON.parse(data.toString()))
    })
}
// 测试
getFile('a.json',aData=>{
    console.log('aData',aData)
    getFile(aData.next,bData=>{
        console.log('bData',bData)
        getFile(bData.next,cData=>{
            console.log('cData',cData)
        })
    })
})
```

再新建 3 个文件 a.json, b.json, c.json，内容如下

```json
// a.json

{"next":"b.json"}

// b.json

{"next":"c.json"}

// c.json

{"next":"test"}
```

```sh
$ node callback.js
aData { next: 'b.json' }
bData { next: 'c.json' }
cData { next: 'test' }

```

由此可见，callback 层层嵌套。当然现在只有 3 层，实际项目中可能会有很多层，可想而知维护的难度，接下来就是解决方案。

### promise

新建个 promise.js 文件，代码如下

```js
const fs = require('fs')
const path  = require('path')
//promise 方式获取文件内容
function getFile(fileName){
    const promise = new Promise((resolve,reject)=>{
        const fullName = path.resolve(__dirname,"",fileName)
        fs.readFile(fullName,(err,data)=>{
            if(err){
                reject(err)
                return
            }
            resolve(JSON.parse(data.toString()))
        })
    })
    return promise
}
// 测试
getFile('a.json').then(aData=>{
    console.log('aData',aData)
    return getFile(aData.next)
}).then(bData=>{
    console.log('bData',bData)
    return getFile(bData.next)
}).then(cData=>{
    console.log('cData',cData)
})
```

表面看上去代码略微繁琐，但是使用起来很简洁，promise 把 callback 层层嵌套的方式变成了很多个 then 平铺，当然 promise 无法解决把异步代码写成同步的风格，我们来讲下异步写法的终极解决方案，async 和 await

### async 和 await

新建个 async.js 文件，代码如下

```js
const fs = require('fs')
const path  = require('path')
//promise 方式获取文件内容
function getFile(fileName){
    const promise = new Promise((resolve,reject)=>{
        const fullName = path.resolve(__dirname,"",fileName)
        fs.readFile(fullName,(err,data)=>{
            if(err){
                reject(err)
                return
            }
            resolve(JSON.parse(data.toString()))
        })
    })
    return promise
}
// 测试
async function getFileContent(){
    aData = await getFile('a.json')
    console.log('aData',aData)
    bData = await getFile(aData.next)
    console.log('bData',bData)
    cData = await getFile(bData.next)
    console.log('cData',cData)
}
getFileContent()
```

async 和 await 也是基于 promise 的，代码写法和同步一样了。

## 总结

今天我们讲了 callback, promise 和 async, await。你掌握了吗？在项目中我们使用的 koa2 框架完全支持 async 和 await。
有问题欢迎到群里和志同道合的小伙伴一起交流。
下节课我们介绍项目功能和技术架构，继续加油吧，Let's go！
