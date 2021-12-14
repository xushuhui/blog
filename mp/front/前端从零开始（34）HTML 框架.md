# HTML 框架

框架可以理解为在网页中插入了一个独立的网页窗口元素，比较常用的框架元素是 iframe，iframe 有自己独立的窗口 window 以及上下文，iframe 元素自己内部的逻辑操作可以独立，当然在同域条件下，iframe 元素的窗口可以访问父级窗口，父级窗口也可以访问 iframe 窗口。

iframe 在网站应用刀耕火种的开发年代是非常常见的，现在前端应用嵌入 iframe 的场景越来越少了，但是在大型的网站中也会经常利用 iframe 嵌入多个网页到一套网站系统中，方便用户在一个系统中去进行业务操作，而不需要在几个不同的系统中来回切换。

## 1. iframe

### 1.1 定义方式

在网页中通过定义一个 iframe 标签即可引用另一个网页进来，例如：

```javascript

<iframe  src="http://www.baidu.com" height=500 width=500 /><!-- 通过src设置iframe的地址 -->
```

### 1.2 height width 属性

通过设置 height 和 width 可以控制 iframe 的高和宽，但是实际项目中有时你并不知道所引用的 iframe 的宽高，或者说 iframe 的宽和高是动态不固定的，这时需要先获取到引用的 iframe 的宽高，然后将值设置给其属性，例如：

```javascript
<iframe id="iframe" src="http://www.imooc.com/wiki/html5/forms.html?preview=d2269b100476b4b77b9ddc11242e85ab" frameborder="1"></iframe><!--src的域名需要和当前网页域名一直，否则跨域不能获取iframe的contentWindow-->
 <script>
var iframe = document.getElementById("iframe");//获取iframe的dom节点
iframe.onload = function(){
    var doc     =  this.contentWindow.document;
    this.height =  doc.documentElement.scrollHeight;//高度动态设置
    this.width  =  doc.documentElement.scrollWidth;//宽度动态设置
}
</script>
```

> **注意：** 上述代码需要有一些 JavaScript 基础才能理解。

### 1.3 frameborder 属性

该属性表示是否显示 iframe 边框，默认为 1 表示显示边框，0 表示隐藏，例如：

```javascript
<iframe src="https://www.baidu.com" frameborder=0></iframe><!--无边框-->
<iframe src="https://www.baidu.com" ></iframe><!--默认有边框（很细的一个框，仔细看）-->
```

### 1.4 scrolling 属性

该属性表示是否显示滚动条，默认根据 iframe 的尺寸来判断是否需要滚动条，设置为 yes 时将始终显示滚动条，设置为 no 时表示完全不显示滚动条：

* auto 默认值
* yes 有滚动条
* no 无滚动条

默认值 auto 表示，当内容超过 iframe 高或宽时，自动出现滚动条。

```javascript
<iframe src="https://www.baidu.com"></iframe><!--默认情下，在需要的情况下出现滚动条-->
<iframe src="https://www.baidu.com" scrolling='yes'></iframe><!--强制显示滚动条-->
<iframe src="https://www.baidu.com" scrolling='no'></iframe><!--强制不显示滚动条-->
```

### 1.5 srcdoc 属性

该属性定义在 iframe 展示的 HTML 描述代码，该属性不支持 IE 浏览器，例如：

```javascript
<iframe src="https://www.baidu.com" srcdoc='<p>使用百度搜索</p>'></iframe>
```

如果浏览器支持 srcdoc 属性，会只显示 srcdoc 中的内容，不支持的时候 ( IE ) 就显示 src 指定的内容。

### 1.6 sandbox 属性

该属性用于表示对 iframe 的权限限制，可选值有：

* 空，表示限制所有权限
* allow-same-origin 允许相同源
* allow-top-navigation 允许包含导航栏
* allow-forms 允许表单提交
* allow-scripts 允许脚本执行

```javascript
<iframe src="https://www.baidu.com" sandbox="" width=600 height=600>
</iframe>
```

执行上述代码，会发现加载之后的百度首页虽然显示，但是并不能执行搜索，这是因为 sandbox 属性设置**为空字符串**时，iframe 会限制框内网页运行 JavaScript 脚本。

### 1.7 优缺点及建议

页面引用 iframe 元素，相当于引用一个完整的 HTML 网页。

**这样做的好处是：**

* 代码可复用性，相同的页面无需重复实现，只需要引用即可；
* iframe 是一个封闭的运行环境，环境变量完全独立、隔离，不会污染宿主环境；
* iframe 可以用于创建新的宿主环境，用于隔离或者访问原始接口及对象，提升网站的安全性

**缺点是：**

* 被引用的 iframe 如果过多的话，可能会产生过量的 HTTP 请求；
* 跨域问题；
* 样式不容易适配

基于 iframe 的优缺点来看，在实际项目开发中，一般用来加载广告、播放器、富文本编辑器等非核心的或者需要格里运行的网页代码。

### 1.8 iframe 通信及跨域问题

#### 1.8.1 iframe 通信

在使用 iframe 时难免会碰到需要在父窗口中使用 iframe 中的变量、或者在 iframe 框架中使用父窗口的变量，在 iframe 的域名和父窗口的域名完全一致的情况下，可以实现调用：

**在父窗口中调用 iframe 元素的变量：**

可以使用 contentWindow 的方式调用：

```javascript
<iframe src='index.html' id='test' />
<script>
    //父窗口调用 iframe 的window对象
	var obj = document.getElementById("test").contentWindow;
</script>
```

上述代码在父窗口中调用 iframe 元素的变量，以下是在 iframe 中调用父窗口的变量的方式：

**在 iframe 中调用父窗口的变量：**

```javascript
<script>
	var dom = window.top.document.getElementById("父窗口的元素ID");
</script>
```

**兄弟 iframe 间相互调用变量：**

```javascript
<iframe src='index1.html' id='test1' />
<iframe src='index2.html' id='test2' />
<script>
	var dom2 = window.top.document.getElementById("test2").contentWindow.getElementById("");//这里是在test1调用test2中的某个dom
</script>
```

#### 1.8.2 跨域问题

但是，JavaScript 出于安全方面的考虑，不允许跨域调用其他页面的对象。这样在安全限制的同时也给 iframe 元素上带来了不少麻烦，导致一个网页中如果出现与当前域名不相同的 iframe，就无法通过 JavaScript 调用 iframe 中的 DOM 结点了 。

> 扩展知识：什么是跨域？简单地理解就是因为 JavaScript 同源策略的限制，[a.com](http://a.com) 域名下的 JavaScript 无法操作 [b.com](http://b.com) 或是 [c.com](http://c.com) 域名下的对象。

## 2. 小结

框是 HTML 中引用外部网页的方式，本章介绍了 iframe 元素的使用方式，以及使用框的优缺点。在实际项目开发中需要结合网页设计中的网络加载时间、性能、代码可用性等因素考虑，决定是否需要使用 iframe。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
