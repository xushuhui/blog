# websocket

网页中的绝大多数请求使用的是 HTTP 协议，HTTP 是一个无状态的应用层协议，它有着即开即用的优点，每次请求都是相互独立的，这对于密集程度较低的网络请求来说是优点，因为无需创建请求的上下文条件，但是对于密集度或者实时性要求较高的网络请求（例如 IM 聊天）场景来说，可能 HTTP 会力不从心，因为每创建一个 HTTP 请求对服务器来说都是一个很大的资源开销。这时我们可以考虑一个相对性能较高的网络协议 Socket，他的网页版本被称为 Websocket。

## 1. 背景

近年来，随着 HTML5 和 w3c 的推广开来，WebSocket 协议被提出，它实现了浏览器与服务器的实时通信，使服务端也能主动向客户端发送数据。在 WebSocket 协议提出之前，开发人员若要实现这些实时性较强的功能，经常会使用一种替代性的解决方案——轮询。

**轮询**的原理是采用定时的方式不断的向服务端发送 HTTP 请求，频繁地请求数据。明显地，这种方法命中率较低，浪费服务器资源。伴随着 WebSocket 协议的推广，真正实现了 Web 的即时通信。

**WebSocket** 的原理是通过 JavaScript 向服务端发出建立 WebSocket 连接的请求，在 WebSocket 连接建立成功后，客户端和服务端可以实现一个长连接的网络管道。因为 WebSocket 本质上是 TCP 连接，它是一个长连接，除非断开连接否则无需重新创建连接，所以其开销相对 HTTP 节省了很多。

## 2. API

### 2.1 创建连接

通过使用新建一个 websocket 对象的方式创建一个新的连接，不过在创建之前需要检测一下浏览器是否支持 Websocket，因为只有支持 HTML5 的浏览器才能支持 Websocket，如下：

```javascript
if(typeof window.WebSocket == 'function'){
    var ws = new WebSocket('http://127.0.0.1:8003');//创建基于本地的8003端口的websocket连接
}else alert("您的浏览器不支持websocket");
```

上述代码会对本地的 8003 接口请求 Websocket 连接，前提是本地的服务器有进程监听 8003 端口，不然的话会连接失败。

### 2.2 创建成功

由于 JavaScript 的各种 IO 操作是基于事件回调的，所以 Websocket 也不例外，我们需要创建一个连接成功的回调函数来处理连接创建成功之后的业务处理，如下：

```javascript
ws.onopen = function(){//通过监听 open 时间来做创建成功的回调处理
    console.log('websocket连接创建成功')
    //进行业务处理
}
```

### 2.3 接收消息

我们辛辛苦苦创建了长连接就是为了发送或者接收网络数据，那么怎么接收呢，跟上边提到的意义，还是需要在回调函数里处理，一不小心就陷入了回调地狱了：

```javascript
ws.onmessage = function(event){
    var d = event.data;
    //接收到消息之后的业务处理
    switch(typeof d){//判断数据的类型格式
    case "String":
        break;
    case "blob":
        break;
    case "ArrayBuffer":
        break;
    default:
        return;
    }
}
```

上述实例通过监听 message 事件对 websocket 的消息进行一定的业务处理，这其中需要判断数据类型格式，因为 Websocket 是基于二进制流格式的，传输过来的消息可能不一定是基于 utf8 的字符串格式，因此需要对格式进行判断。

### 2.4 发送消息

客户端通过使用 send 函数向服务端发送数据，例如：

```javascript
ws.send("一段测试消息");
```

可以发送文本格式，也可以发送二进制格式，例如：

```javascript
var input  = document.getElementById("file");
input.onchange = function(){
    var file = this.files[0];
    if(!!file){
        //读取本地文件，以gbk编码方式输出
        var reader = new FileReader();
        reader.readAsBinaryString(file);
        reader.onload = function(){
            //读取完毕后发送消息
            ws.send(this.result);
        }
    }
}
```

### 2.5 监听错误信息

类似上述提到的如果创建实例失败的情况，系统会出现异常，但是我们并不能准确判断出异常的信息，这时需要通过监听错误事件来获取报错信息，例如：

```javascript
ws.onerror = function(event){
    //这里处理错误信息
}

```

### 2.6 关闭连接

当服务端或者客户端关闭 websocket 连接时，系统会触发一个关闭事件，例如：

```javascript
ws.onclose = function (event){
    //这里处理关闭之后的业务
}

```

### 2.7 连接的状态

通过 websocket 对象的 readyState 属性可以获取到当前连接的状态，其中常用的有 4 种，通过 websocket 对象的几种定义常量对比判断：

```javascript
switch (ws.readyState){
    case WebSocket.CONNECTING:break;//处于正在连接中的状态
    case WebSocket.OPEN:break;//表示已经连接成功
    case WebSocket.CLOSING:break;//表示连接正在关闭
    case WebSocket.CLOSE:break;//表示连接已经关闭，或者创建连接失败
    default:break;
}
```

## 3. websocket 实例

```javascript
<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8">
    <style>
        p {
            text-align: left;
            padding-left: 20px;
        }
    </style>
</head>
<body>
<div style="width: 700px;height: 500px;margin: 30px auto;text-align: center">
    <h1>聊天室实战</h1>
    <div style="width: 700px;border: 1px solid gray;height: 300px;">
        <div style="width: 200px;height: 300px;float: left;text-align: left;">
            <p><span>当前在线:</span><span id="user_num">0</span></p>
            <div id="user_list" style="overflow: auto;">

            </div>
        </div>
        <div id="msg_list" style="width: 598px;border:  1px solid gray; height: 300px;overflow: scroll;float: left;">
        </div>
    </div>
    <br>
    <textarea id="msg_box" rows="6" cols="50" onkeydown="confirm(event)"></textarea><br>
    <input type="button" value="发送" onclick="send()">
</div>
</body>
</html>

<script type="text/javascript">
    var uname = window.prompt('请输入用户名', 'user' + uuid(8, 16));
    var ws = new WebSocket("ws://127.0.0.1:8081");
    ws.onopen = function () {
        var data = "系统消息：连接成功";
        listMsg(data);
    };
    ws.onmessage = function (e) {
        var msg = JSON.parse(e.data);
        var data =  msg.content;
        listMsg(data);
    };

    ws.onerror = function () {
        var data = "系统消息 : 出错了,请退出重试.";
        listMsg(data);
    };

    function confirm(event) {
        var key_num = event.keyCode;
        if (13 == key_num) {
            send();
        } else {
            return false;
        }
    }

    /**
     * 发送并清空消息输入框内的消息
     */
    function send() {
        var msg_box = document.getElementById("msg_box");
        var content = msg_box.value;
        var reg = new RegExp("\r\n", "g");
        content = content.replace(reg, "");
        var msg = {'content': content.trim(), 'type': 'user'};
        sendMsg(msg);
        msg_box.value = '';
    }

    /**
     * 将消息内容添加到输出框中,并将滚动条滚动到最下方
     */
    function listMsg(data) {
        var msg_list = document.getElementById("msg_list");
        var msg = document.createElement("p");

        msg.innerHTML = data;
        msg_list.appendChild(msg);
        msg_list.scrollTop = msg_list.scrollHeight;
    }

    /**
     * 将数据转为json并发送
     * @param msg
     */
    function sendMsg(msg) {
        var data = JSON.stringify(msg);
        ws.send(data);
    }
</script>
```

上述实例通过使用 websocket 实现了一个简单的聊天室功能，功能上只实现了接受和发送消息的功能，在登录认证和安全性等问题上并没有做过多的处理，只是为了给大家连贯的展示一下 websocket 在实际项目中的使用。

## 4. 注意事项

实际项目中使用 websocket 需要注意一些问题 :

* websocket 创建之前需要使用 HTTP 协议进行一次握手请求，服务端正确回复相应的请求之后才能创建 websocket 连接；
* 创建 websocket 时需要进行一些类似 token 之类的登录认证，不然任何客户端都可以向服务器进行 websocket 连接；
* websocket 是明文传输，敏感的数据需要进行加密处理；
* 由于 websocket 是长连接，当出现异常时连接会断开，服务端的进程也会丢失，所以服务端最好有守护进程进行监控重启；
* 服务器监听的端口最好使用非系统性且不常使用的端口，不然可能会导致端口冲突

## 5. 小结

本章介绍了 websocket 的前世今生，详细说明其对应的 API 的调用方式，最后使用了一个简单的聊天室的例子来对其函数串通了一下，最后延伸了一下实际项目中使用 websocket 需要注意的地方，希望大家在实际开发中针对其优缺点来选择合适的使用场景。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
