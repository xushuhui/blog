# HTML5 SSE 浏览器发送事件

在远古时代，网页大都是静态展示，服务器无需处理复杂且过多的请求，只需要静静地等待客户端的请求，将 HTML 代码通过 HTTP 的方式返回给客户端。因此服务器也没有主动推送数据给客户端的能力，毕竟 HTTP 是无状态的协议，即开即用。

后来随着互联网的发展，服务端有一些即时消息需要立即展示给客户端，早期的处理方式是通过客户端定时发起 HTTP 请求，这种方式命中率较低且浪费服务端资源。现在有了 HTML5 之后不需要那么麻烦了，可以使用 websocket 或者 SSE。SSE 全称 server-sent events 单项消息传递事件，相对于 websocket 这种双向协议，SSE 较为轻量，它只支持服务端向客户端推送消息。

## 1. 使用方式

### 1.1 创建实例

通过新建一个 sse 对象可以创建一个 SSE 实例，但是不要忘记检测浏览器的支持情况：

```javascript
if(typeof(EventSource)!=="undefined"){
    var source = new EventSource("http://127.0.0.1/test.php");
}

```

上述示例实现了一个创建 SSE 对象的功能，创建之前需要检测是否支持，目前 IE 之外的大部分浏览器都支持 SSE。sse 对象只有一个初始化参数，用于指定服务器的 url。

### 1.2 接收消息

创建实例成功之后，通过监听 message 事件来实时获取服务端的消息：

```javascript
source.onmessage = function (event){
    //处理业务请求
    console.log(event.data)
}

```

### 1.3 服务端支持

服务器端需要对客户端发起的 HTTP 请求做相应的回复，主要是将 HTTP 报文头的 content-type 字段设置成 text/event-stream，下边以 PHP 举例：

```javascript
header('content-type:text/event-stream');
while(true){
    sleep(30000);
    echo "message:".time();
    //每隔半分钟返回一个时间戳
}

```

### 1.4 其他事件

除了监听 message 事件用于获取服务端的数据之外，还有 open 事件用于监听连接打开的状态， error 事件用于监听错误信息。

## 2. 几种常用的客户端 - 服务器消息传递方式

* http 最常用的协议，用于客户端主动向服务器发送请求，单向传递；
* ajax HTTP 的扩展版，底层还是 HTTP 协议，只不过客户端是无刷新的；
* comet 也是基于 HTTP 封装的，使用 HTTP 长连接的方式，原理大致是将 HTTP 的 timeout 设置较长，服务器有数据变化时返回数据给客户端，同时断开连接，客户端处理完数据之后重新创建一个 HTTP 长连接，循环上述操作（这只是其中一种实现方式）；
* websocket 这是 HTML5 中的新标准，基于 socket 的方式实现客户端与服务端双向通信，需要浏览器支持 HTML5；
* Adobe Flash Socket 这个也是使用 socket 的方式，需要浏览器支持 flash 才行，为了兼容老版本的浏览器；
* ActiveX object 只适用于 IE 浏览器；

目前尚没有一种方式能兼容所有的浏览器，只能针对软件的目标客户人群做一定的兼容。
* sse 服务端单向推送。

## 3. 适用场景

并非所有场景都适合使用 sse 处理，在消息推送接收不频繁的情况下选用 ajax 轮询或者 sse 或者 websocket 其实差别不太大。sse 应该适用于服务端向客户端发送消息频繁而客户端几乎无需向服务端发送数据的场景下，例如：

* 新邮件通知；
* 订阅新闻通知；
* 天气变化；
* 服务器异常通知；
* 网站公告；
* 等等。

sse 的优缺点：

* SSE 使用 HTTP 协议，除 IE 外的大部分浏览器都支持；
* SSE 属于轻量级，使用简单；
* SSE 默认支持断线重连；
* SSE 一般只用来传送文本，二进制数据需要编码后传送；
* SSE 支持自定义发送的消息类型。

## 4. 总结

本章介绍了 websocket 的轻量级版本 sse 协议，简述了 sse 协议的使用方法，对比了其他网页中常用的消息推送方式以及他们的优缺点，这些协议涵盖了大部分的使用场景，选用适合的协议类型可以避免不必要的资源和性能消耗。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
