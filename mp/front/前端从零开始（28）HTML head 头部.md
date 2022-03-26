# head 头部

本章节讲解 HTML 中的 head 标签，以及 head 内包含关于 HTTP 协议的标签 -meta。

`<head>` 定义文档的头部，它是所有头部元素的容器。`<head>` 中的元素可以引用脚本、指示浏览器在哪里找到样式表、提供元信息等等。

## 1. head

head 标签是 HTML 标准框架中的头部分，前边章节我们讲到 HTML 的标准框架中提到 一个标准的 HTML 中包含 HTML 标签、head 标签、 body 标签以及 w3c 文档标准头。所以说 head 标签是必须的标签，它本身没有视觉展示效果，仅仅是作为容器标签，其中可以包含的标签包含：

* base
* link
* meta
* script
* style
* title

例如：

```javascript
<html>
<head>
  <title>标题</title>
</head>
<body>
  内容... ...
</body>
</html>
```

所有主流的浏览器都支持 head 标签。

## 2. meta

### 2.1 meta 标签的作用

* 优化搜索引擎
* 定义页面使用语言
* 控制页面缓存
* 网页定义评价
* 控制页面显示窗口
* ……

例如

```javascript
<meta name="keywords" content="HTML,PHP,SQL"> <!-- 定义文档关键词，用于爬虫搜索引擎 -->
<meta http-equiv="charset" content="iso-8859-1"> <!-- 定义文档的字符集 -->
<meta http-equiv="expires" content="31 Dec 2020"> <!-- 定义文档的缓存过期时间 -->
```

### 2.2 meta 的属性

* name 描述网页
* content 方便搜索引擎查找和分类
* http-equiv http 文件头设置

## 3. header

header 标签定义文档的页眉，与 head 类似，它也是仅仅起到容器作用，不同的是 header 非网页必须标签，而且是 HTML5 的新增标准，放到 header 标签的内容大都是一些文档的介绍信息，例如：

```javascript
<header>
<h1>本章介绍HTML头</h1>
</header>
```

## 3. 其他

head 内还可以包括 link、script 等标签，用于引用 css JavaScript 文件等作用，例如：

```javascript
<head>
<link href='/css/1.css' type="text/css" /> <!-- 定义层叠样式表 -->
<script src="/script/1.js" type='text/javascript'></script><!-- 定义JavaScript脚本 -->
</head>
```

在实际项目开发中，为了更好的扩展性和可维护性，一般会把 head 标签以及其中的内容放到一个全局 include 文件，因为这个文件一般改动不太频繁且每个文件都必须引用，所以由框架在加载的时候自动引用，这样的设计方式符合 MVC 以及面向对象开发的思想。

## 4. 小结

回顾本节，主要介绍了 HTML 中 head 标签的语法和作用，以及其中包含的其他标签元素的用法。

```javascript

```

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
