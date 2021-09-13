# HTML 文件基本结构

HTML 文件和普通文本文件不一样的地方还在于，我们如果想在 HTML 文件当中编写网页的内容，我们不仅要遵循 HTML 的语法，我们还必须按照 HTML 文件的结构来编写我们的代码，只有我们按照规定的结构来编写代码，才能达到我们想要的效果。

## 1. HTML 结构代码展示

```javascript
<!DOCTYPE HTML> <!-- HTML5标准网页声明 -->
<HTML> <!-- HTML为根标签，代表整个网页 -->

<head> <!-- head为头部标签，一般用来描述文档的各种属性和信息， 包括标题等-->
  <meta charset="UTF-8"> <!-- 设置字符集为utf-8 -->
  <title>my HTML</title> <!-- 设置浏览器的标题 -->
</head>

<!-- 网页所有的内容都写在body标签内 -->
<body>
  我的第一个HTML网页
</body>

</HTML>
```

> **Tips**：`<!-- -->`为 HTML 文件的注释， 注释的内容写在 `<!-- -->` 内，但不会在页面中显示。

## 2. HTML 文件结构详解

1. `<!DOCTYPE HTML>` 标签：

为文档类型声明，表示该文件为 HTML5 文件。 <!DOCTYPE> 声明必须是 HTML 文档的第一行，位于 `<HTML>`标签之前。

2. `<HTML></HTML>`标签对：

`<HTML>` 标签位于 HTML 文档的最前面，用来标识 HTML 文档的开始； `</HTML>` 标签位于 HTML 文档的最后面，用来标识 HTML 文档的结束；这两个标签对成对存在，中间的部分是文档的**头部**和**主题**。

3. `<head></head>` 标签对：

标签包含有关 HTML 文档的信息，可以包含一些辅助性标签。如 `<title></title>` ，`<link /><meta />` ， `<style></style>` ， `<script></script>` 等，**但是浏览器除了会在标题栏显示 `<title>` 元素的内容外，不会向用户显示 `head` 元素内的其他任何内容**。

4. `<body></body>` 标签对：

它是 HTML 文档的主体部分，在这个标签中可以包含 `<p><h1><br>` 等众多标签，`<body>` 标签出现在 `</head>` 标签之后，且必须在闭标签 `</HTML>` 之前闭合。

## 3. 小结

1. HTML 文件必须按照规定结构来熟悉；
2. `<!DOCTYPE HTML>` 代表 `HTML5` 标准网页声明；
3. `HTML` 标签代表网页根标签；
4. `head` 代表头部标签，一般用来放描述文档的各种属性和信息；
5. `meta` 标签用来设置当前文件的编码集；
6. `title` 标签用来设置网页的标题；
7. `body` 用来放网页的主体内容。

代码展示：

```javascript
<!DOCTYPE HTML> <!-- HTML5标准网页声明，放在HTML文件的最顶部 -->
<HTML> <!-- HTML为根标签，代表整个网页 -->
  <head> <!-- head为头部标签，一般用来描述文档的各种属性和信息， 包括标题等-->
    <meta charset="UTF-8"> <!-- 设置字符集为utf-8 -->
    <title>my HTML</title> <!-- 设置浏览器的标题 -->
  </head>
  <!-- 网页所有的内容都写在body标签内 -->
  <body>
    我的第一个HTML网页
  </body>
</HTML>
```

![图片描述](https://img.mukewang.com/wiki/5f62fa8f0961c9ad13450542.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
