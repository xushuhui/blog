# AJAX

> Asynchronous JavaScript + XML（异步 JavaScript 和 XML）, 其本身不是一种新技术，而是一个在 2005 年被 Jesse James Garrett 提出的新术语，用来描述一种使用现有技术集合的‘新’方法。(MDN)

AJAX 是 2005 年提出的一种术语，并不代表某个特定的技术。

其译名 `异步JavaScript和XML` 描述出了核心，就是使用 `JavaScript` 发送异步 HTTP 请求，这样就摆脱了想要和服务端交互，必须刷新页面的痛点。

学习 AJAX 相关内容前，建议有一些简单的 HTTP 相关知识的储备，否则很难理解其工作流程。

## 1. XMLHttpRequest 对象

`XMLHttpRequest` 对象可以提供给前端开发人员使用 `JavaScript` 发起 HTTP 请求的能力。该对象会被简称为 `XHR` 对象。

```javascript
var xhr = new XMLHttpRequest();
```

这样就获得到了一个 `XHR` 对象的实例。接下来可以使用他发起请求。

```javascript
var xhr = new XMLHttpRequest();

xhr.onreadystatechange = function() { // 当 readyState 改变的时候
  if (xhr.readyState === 4 && xhr.status === 200) { // 判断当前请求的状态 与 请求的状态码
    console.log(xhr.responseText); // 输出服务端返回的内容
  }
}
xhr.open('GET', '/', true); // 设定 GET 请求，请求的路径是 /，并且请求是异步的
xhr.send(); // 发送！
```

![图片描述](https://img.mukewang.com/wiki/5eb2d8f40a37769217430926.jpg)

`onreadystatechange` 是一个事件处理器属性，每次 `readyState` 改变的时候都会触发。

如果 `readyState` 为 4，即请求已经完成，并且状态码是 200，表示请求结束并且服务端成功响应。

响应成功，通过 `responseText` 获取到服务端响应的内容。

通过 `open` 方法，设置请求的方法、路径等，例子中设置了 `/` 路径，如果当前站点的域名是 `imooc.com`，则请求地址就是 `imooc.com/`，拿到的数据应该会是网站首页的 HTML。

然后通过 `send` 方法发送请求，发送后 readyState 会在各个阶段发送改变，然后调用 `onreadystatechange`。

这是一个 AJAX 请求较为基本的流程。

## 2. ActiveXObject

IE6 之前是没有 XHR 对象的，需要使用 `ActiveXObject` 对象，这是 IE 特有的。

考虑到 IE6 之前版本浏览器的市场份额，个人觉得只需要做个了解，没太大必要再去学习使用 ActiveXObject 对象。

## 3. 小结

现代网页的构成几乎离不开 `AJAX` 技术，如果 JavaScript 无法发起 HTTP 请求，几乎所有的现代网站都会瘫痪，变得难用。

除了 `XHR` 对象之外，还有 [fetch API](https://developer.mozilla.org/zh-CN/docs/Web/API/Fetch_API/Using_Fetch) 这个后起之秀，也可以承担 `XHR` 对象的工作。但 fetch 还有部分缺陷，如无法监控进度，对状态码的处理逻辑不够符合业务，所以用户面较广的产品使用一般不会选择 fetch。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
