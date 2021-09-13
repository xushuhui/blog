# HTML5 离线存储

本章介绍一下 HTML5 新增的离线存储特性 Localstorage，主要包括 Localstorage 的发展史、兼容性、优缺点以及使用场景。

说到 Localstorage，为什么要使用 Localstorage 呢？

因为开发程序当然就要存储数据了，但是 Web 开发场景比较特殊，数据正常情况下是要通过 HTTP 协议发送给服务器端，存储在服务器上，但是呢，如果所有数据都存到服务器一方面会浪费服务器资源，另一方面也会降低网页响应速度，所以设计网页时会抽取一些不太重要的或者临时性的数据使用离线存储方式放在浏览器上。

总的来说，Localstorage 是一个简单的离线存储技术，通过官方提供的增删改查的 API 可以方便的操作它，需要考虑的难点是每个浏览器的容量限制，使用时做好容错即可。

## 1. 离线存储发展史

在早期的互联网发展中，浏览器制定了不同的标准用于存储离线数据，其中比较出名的有微软 IE 浏览器的 userData（单个页面可存储 64 kb）、Adobe 的 flash6 中的 flash-cookies（允许存储 100kb）、flash8 中的 externalinterface、Google 的 gears，不幸的是这些技术没有统一的标准，而且只适用于单一的浏览器，不能跨平台，所以没有收录在 HTML 标准中。HTML5 之前，Cookie 是唯一在 HTML 标准中用于离线存储的技术，但是 Cookie 有一些不太友好的特征限制了它的应用场景：

* Cookie 会被附加在 HTTP 协议中，每次请求都会被发送到服务器端，增加了不必要的流量损耗
* Cookie 大小限制在 4kb 左右（不同的浏览器有一些区别），对于一些复杂的业务场景可能不够

这两个缺点在 Localstorage 中得到了有效的解决，下面我们就开始学习 Localstorage。

## 2. 兼容性

截止目前为止，已经有大部分浏览器已经支持 Localstorage，包括 IE8。

![](https://img4.sycdn.imooc.com/5bab321d0001972205000121.jpg)

具体浏览器是否支持 Localstorage 可以通过简单的 JavaScript 代码判断。

```javascript
function testLocalstorage(){
    if ( typeof window.localStorage == "object" ) return true;//判断localstorage对象是否定义
    else return false;//未定义返回false
}
```

## 3. API 接口

Localstorage 是一个简单的 key/value 形式的数据库，以键值对的方式存储，所以提供的接口主要是基于 k/v 的操作。基于提供的接口只能存储简单的一维数组，但是有些业务场景可能会牵涉到多维数据甚至对象的存储，怎么办？

1. 建议使用 `JSON.stringify()` 将数据转化成字符串方式再存储；
2. 使用复杂的前端数据库，例如 indexDB，具体不做深入讨论。

### 3.1 存储数据

```javascript
window.localStorage.setItem("test",1)//设置key=test的值为1
localStorage.setItem("test",1)//设置key=test的值为1，localstorage可以作为全局对象处理
localStorage.test = 1//可以通过属性值的方式直接操作localstorage的key
```

运行下面案例代码，试一试：

```javascript
<!DOCTYPE html>
<html>
<body>
<p id="demo"></p>
<script>

window.localStorage.setItem("test",1)//设置key=test的值为1
localStorage.setItem("test",1)//设置key=test的值为1，localstorage可以作为全局对象处理
localStorage.test = 1//可以通过属性值的方式直接操作localstorage的key

document.getElementById("demo").innerHTML =
"localStorage.test 的值是" + localStorage.test + "。";

</script>

</body>
</html>
```

### 3.2 读取数据 - 按键值

**getItem**

```javascript

var a = window.localStorage.getItem("test")//获取key=test的值
var a = localStorage.test//可以直接通过对象属性的方式操作
```

如果获取一个不存在的 key 返回 null ，下同。

### 3.3 读取数据 - 按位置

```javascript

var a = window.localStorage.key(0)//可以根据key在localstorage的位置的方式操作，类似操作JavaScript的array的方式
```

### 3.4 删除数据

```javascript
window.localStorage.removeItem("test")//删除key=test的值
window.localStorage.test = ''//可以通过赋空值的方式等价操作

```

### 3.5 整体清空

```javascript

window.localStorage.clear()//clear函数清空整个localstorage
```

### 3.6 存储事件监听

当 localstorage 发生改变时，可以通过监听 storage 事件作出相应的业务处理。

```javascript

if (window.addEventListener) {   //通过addEventListener方式监听事件，为了兼容IE
    window.addEventListener("storage", function(e){//监听storage事件
        //业务处理
    }, false);
} else {
    window.attachEvent("onstorage", function(e){//通过attachEvent方式监听事件
        //业务处理
    });
}
```

## 4. 适用场景及局限性

### 4.1 局限性

前边提到 Localstorage 相比较 Cookie 的优势是容量大和节省 HTTP 带宽，但是它还是有自身的缺点，下边罗列了它的缺点

* 5M 容量依然小，用过数据库的同学应该知道，MySQL 随便一个表加上索引很容易超过 5M
* 不能跨域名访问，同一个网站可能会牵涉到子域名
* 不能存储关系型数据
* 不能搜索

### 4.2 适用场景

那么以上缺点有没有解决方案，肯定是有的，例如 HTML 的 webSql 或者 indexDB，那肯定有人问了，为什么不直接用最复杂的数据库，跳过 Localstorage 呢？原因是技术没有最好的，只有最适合的，不同的业务场景应该选择最匹配的而且成本最小的解决方案。例如你在存储简单的业务场景中的临时数据时完全可以使用 Localstorage 甚至 Cookie 搞定，假如使用 indexDB 的话系统的开发成本以及维护成本会翻番，得不偿失。

所以说总结下来 Localstorage 的适用业务场景是：

* **数据关系简单明了**
* **数据量小**
* **数据无需持久化存储且不需要考虑安全性**
* **无需跟服务器交互**

## 5. 业务实战

### 5.1 项目使用场景

之前开发一个场馆管理系统时有一个功能是根据用户输入的关键字搜索场馆，业务方的需求是需要临时保留搜索关键词的历史记录，考虑到是临时保存，而且只保存关键字不需要复杂的数据结构存储且只保存 10 条最新的数据，项目组商量下来决定使用 Localstorage 保存，搜索成功之后添加到历史记录。

```javascript

function chooseClubItem(e) {
        let mid = e.currentTarget.dataset.id
        let findAddressInfo = this.data.markers.find(item => item.id === mid)
        const historyList = window.localStorage.getItem('historyList') || []//获取localstorage需要操作的键值
        if (findAddressInfo) {
            const index = historyList.findIndex(history => history.id == findAddressInfo.id)
            if (index !== -1) {
                historyList.splice(index, 1)
            }
            if(historyList.length>=10)historyList.pop();//超过最大历史数目，删除最后一个
            historyList.unshift(findAddressInfo)//加入到历史存储队列中
            window.localStorage.setItem('historyList', historyList)//设置离线存储
        }
    },
```

清空历史记录

```javascript

function delHistory() {
        let that = this
        showModal({//弹出确认对话框
            title: '',
            content: '您确定要清空搜索历史吗',
            showCancel: true,
            cancelText: '取消',
            cancelColor: '#000000',
            confirmText: '确定',
            confirmColor: '#3CC51F',
            success: result => {
                if (result.confirm) {
                   window.localStorage.removeItem("historyList")//情况历史队列
                }
            }
        })
    },
```

### 5.2 使用第三方库

现实中考虑到浏览器对 Localstorage 毕竟不是百分之百兼容，而且 Localstorage 本身提供的 API 比较简单，所以实际项目中可以考虑使用第三方封装库操作，比如 store.js。

store.js 优先选择 localStorage 来进行存储，在 IE6 和 IE7 下降级使用 userData 来达到目的。 没有使用 flash ，不会减慢你的页面加载速度。也没有使用 Cookies ，不会使你的网络请求变得臃肿。store.js 依赖 JSON 来序列化数据进行存储。

## 6. 总结

以上介绍了 Localstorage 和传统离线存储的优缺点对比，Localstorage 的使用方法以及项目实战分析。总的来说 Localstorage 适用于业务简单的轻量级存储中，通过简单的 API 操作增删改查存储键值对，而且可以通过事件监听的方式获取 Localstorage 的操作事件，无需发送 HTTP 请求，真正实现了离线存储

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
