# 浏览器的多线程和单线程

学习过 JavaScript 的可能会了解，JavaScript 的宿主浏览器只有一个线程运行 JavaScript，除了 JavaScript 的线程，浏览器中单个页面还有一些其他线程，例如：UI 线程负责处理渲染 DOM 元素；GUI 线程用于处理与用户交互的逻辑；网络线程用于发送接收 HTTP 请求；file 线程用于读取文件；定时器线程处理定时任务等等。

## 1. 单线程原因

为什么不能像很多高级语言一样支持多线程呢？假定 JavaScript 同时有两个线程，一个线程在 HTML 中创建了一个标签元素，另一个线程删除了这个标签，这时浏览器应该执行什么操作？浏览器中 JavaScript 的主要用途是操作 DOM 。这决定了它只能是单线程，否则会带来很复杂的同步问题。为了避免复杂性，大部分主流浏览器的 JavaScript 运行环境只支持单线程。

## 2. JavaScript 的事件驱动

既然 JavaScript 只支持单线程，那么有人可能会好奇为什么浏览器中的 JavaScript 可以同时发送多个网络请求或者执行多个事件回调函数呢？

这是因为 JavaScript 是基于事件驱动，当需要进行网络请求时，JavaScript 线程会把请求发送给 network 线程执行，并等待执行结果；当进行文件读取时则调用 file 线程，然后等待结果。然后 JavaScript 会一直轮询事件库 event loop，直到有事件完成，这时浏览器会驱动 JavaScript 去执行事件的回调函数。这就是 JavaScript 的事件驱动模型。

![](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1590642785006&di=37bbd5830df62f50e1dbca9d847293d2&imgtype=0&src=http%3A%2F%2Fwww.uml.org.cn%2Fzjjs%2Fimages%2F2016092627.png)

## 3. web worker 诞生

单线程的最大问题是不能利用多核 CPU 的优点，HTML5 推出的 Web Worker 标准，允许 JavaScript 创建多线程，但是子线程受主线程约束，且不得操作 DOM 。所以，这个新标准不会产生多线程同步的问题。

## 4. 适用场景

Web Worker 能解决传统的 JavaScript 单线程出现的执行阻塞问题，因而适合以下几种业务场景：

* 并行计算；
* ajax 轮询；
* 耗时的函数执行；
* 数据预处理 / 加载。

## 5. 函数介绍

### 5.1 创建

初始化一个 Web Worker，由于不是所有的浏览器都支持 Web Worker，所以需要判断一下浏览器是否支持：

```javascript

if (window.Worker) {//判断浏览器是否支持web worker
    var worker = new Worker('test.js');//创建一个线程，参数为需要执行的JavaScript文件
}
```

### 5.2 向线程传递参数

新的线程的上下文环境跟原宿主环境相对独立的，所以变量作用域不同，如果需要互相读取变量的话需要通过消息发送的方式传输变量，例如：

```javascript
worker.postMessage('test'); //数据类型可以是字符串
worker.postMessage({method: 'echo', args: ['Work']});//数据类型可以是对象

```

### 5.3 主线程接受消息

跟上述场景类似，主线程也需要通过监听的方式获取辅线程的消息：

```javascript
worker.onmessage = function (event) {
  console.log('接收到消息： ' + event.data);
}

```

### 5.4 线程加载脚本

子线程内部也可以通过函数加载其他脚本：

```javascript
importScripts('script1.js','script2.js');

```

### 5.5 关闭线程

```javascript

// 主线程中关闭子线程
worker.terminate();
// 子线程关闭自身
self.close();
```

## 6. 使用 JavaScript 多线程实现非阻塞全排列

### 6.1 什么是全排列

**从 n 个不同元素中任取 m（m≤n）个元素，按照一定的顺序排列起来，叫做从 n 个不同元素中取出 m 个元素的一个排列。当 m=n 时所有的排列情况叫全排列。**

### 6.2 为什么使用多线程处理

这里并非突出使用 JavaScript 实现全排列的优势，而是在实际项目中类似这种**科学运算相关的算法可能会消耗一定的 CPU**，由于 JavaScript 是解释型语言，运算性能是它的弱项，而且浏览器中运行的 JavaScript 又是单线程的，所以一旦出现性能问题可能会导致线程阻塞，阻塞之后会导致页面卡顿，非常影响用户体验。使用 webworker 的多线程功能将这个运算函数单独 fork 出一个子线程去运行，运行完成之后发送结果给主线程，可以有效的避免性能问题。

### 6.3 代码示例

```javascript
<html>
<head>
<meta http-equiv="Content-Type" content="text/html;
charset=UTF-8">
<title>JavaScript实现全排列</title>
<script type="text/JavaScript">
function combine() {//点击按钮向webworker线程发送请求
    var worker = new Worker('http://wiki-code.oss-cn-beijing.aliyuncs.com/html5/js/worker.js');
    worker.postMessage(document.getElementById("str").value);
        worker.onmessage= function (event) {
	        document.getElementById("result").innerHTML  =   event.data ; //监听JavaScript线程的结果
	    };
}
</script>
</head>
<body>
    <input type="text" id="str" />
    <button onclick="combine()">全排列</button>
    结果是：<div id="result" style="width:500px;height:500px;word-break: break-all;"></div>
</body>
</html>
```

worker.js 代码如下：

```javascript

function getGroup(data, index = 0, group = []) {//生成全排列
    var need_apply = new Array();
    need_apply.push(data[index]);
    for(var i = 0; i < group.length; i++) {
        need_apply.push(group[i] + data[index]);
    }
    group.push.apply(group, need_apply);
    if(index + 1 >= data.length) return group;
    else return getGroup(data, index + 1, group);
}
onmessage = function(message){//监听主线程的数据请求
    var msg = message.data;
    if(msg == "") postMessage("请输入正确的字符串");
    else {
        var data = msg.split("");//将字符串转数组
        postMessage(getGroup(data));
    }
}
```

上述代码实现了一个使用 JavaScript 的 Web Worker 实现的全排列的功能。上半部分是主线程的代码，主要实现了创建子线程、发送数据给子线程、接收子线程的消息这几个功能；下半部分是子线程，子线程主要负责运算，并将运算结果发送给主线程。

## 7. 总结

早期的 JavaScript 由于考虑操作 DOM 的一致性问题，以及当时的网页没有过多的交互所以不需要大量的计算，所以只支持单线程。这在多核 CPU 时代的劣势愈发明显，所以 HTML5 中推出多线程解决这个问题。回顾本章主要介绍了 Web Worker 的使用方式以及其适用场景。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
