# HTML5 语义化元素

本章节我们来介绍一个抽象的知识点 - 语义化。什么是语义化，浅显的来说就是使用**合适的语法**来实现相应的功能，这里说的合适并非是从性能、数据结构、算法等深度层面，而是从阅读和维护方式等层面。

编程过程中实现一个相同的功能，往往可以使用多种不同的方式，选择一个合适的方式需要综合考虑可维护性、扩展性、性能等几种不同的维度，而可维护性是其中比较重要的一个因素。可维护性就是指书写的代码是否通俗易懂方便阅读，当大家都遵守一种统一的书写标准时，团队的开发效率、协调能力就能得到很大的提升。

## 1. HTML 的语义化

HTML 语义化是指使用恰当语义的 html 标签、class 类名、ID、属性名称 等内容，让页面具有良好的结构与含义，从而让人和机器都能快速理解网页内容。语义化的 HTML 页面一方面可以让搜索引擎高效的理解和搜集网页内容；另一方面方便开发维护人员读懂代码。总结起来就是：正确的标签做正确的事情，页面内容结构化方便开发人员阅读，便于浏览器搜索引擎解析。

## 2. 常用语义化标签

* header 定义某一部分段落或者文本的头部信息
* nav 导航信息
* main 呈现网页的主体结构
* article 用于文本分段
* section 用于对主题相关的内容分组
* footer 定义网页底部
* h1-h6 定义标题栏
* div 定义块
* title 页面标题
* ul 无序列表
* ol 有序列表
* aside 表示与当前内容无关的内容
* small 小号字体
* em 斜体字体
* Mark 黄色突出字体
* figure 独立的流内容
* figcaption 定义 figure 元素的标题
* cite 表示文本是对参考文献的引用
* blockquote 定义块引用
* q 短引述
* time 定义合法的日期或时间格式
* dfn 定义术语元素
* abbr 简称或缩写
* del 定义删除的内容
* ins 添加的内容
* code 标记代码
* meter 定义标量测量
* progress 定义运行中的进度

上述罗列了包含明确语义内容的标签，实际项目中应当根据实际场景选择对应的语义标签。

### 2.1 small

small 标签属于 HTML 中的格式元素，用于显示较小的文本，例如：

```javascript
<small>用于定义小的文本</small>
```

### 2.2 em

em 用于显示斜体，它和 i 标签的效果类似， 不同的是 em 是语义化元素，用于强调斜体，例如：

```javascript
<em>用于显示斜体</em>
```

### 2.3 Mark

Mark 标签用于显示黄色背景的文本，例如：

```javascript
<mark>标记文本</mark>
```

### 2.4 figure

figure 标签用于在文档中插入图片、图标、照片、代码等流内容，例如：

```javascript
<figure>
    <img src="show.jpg" >
</figure>
```

### 2.5 figcaption

figcaption 标签用于 figure 标签的标题，它必须定义在 figure 内部，一个 figure 只能放一个 figcaption ，例如：

```javascript
<figure>
   <img src="actShare.png" alt="这是一张用于演示的图片">
   <figcaption>演示</figcaption>
</figure>
```

### 2.6 cite

cite 标签用于表示对某个文献引用的文本定义，例如书籍、杂志等内容，它所展示的是斜体文本，是一个典型的语义化标签，例如：

```javascript
<cite>语义化标签</cite>已经远远超过了改变它所包含的文本外观的作用，它使浏览器能够以各种实用的方式来向用户表达文档的内容
```

### 2.7 blockquote

blockquote 用于定义源于另一个块内容的引用，它的默认展示方式是左右两侧缩进，例如：

```javascript
<blockquote>
    <p>引用的段落1</p>
    <p>引用的段落2</p>
</blockquote>
```

### 2.8 q

q 标签用于定义短引用，浏览器默认会为它左右显示引号，例如：

```javascript
知识付费市场蓬勃兴起，但也存在<q>内容良莠不齐、服务跟不上等问题</q>，需要加以改进完善
```

### 2.9 time

time 标签用于表示 24 小时制时间或者公历日期，如果表示日期也可以包含时间和时区，这个标签用于搜索引擎的友好型，目前所有主流浏览器都不完全支持 time 标签，例如：

```javascript
<time datetime="YYYY-MM-DDThh:mm:ssTZD">今天是端午节</time>
```

### 2.10 dfn

dfn 标签用于首次定义术语，仅仅包含术语，不必包含术语的定义，再次出现术语时可 abbr 元素表示，例如：

```javascript
<p>The <dfn>GDO</dfn>is a device that allows off-world teams to open the iris.</p>
```

### 2.11 abbr

abbr 标签用于定义一个缩写内容，当鼠标停留在内容上时，浏览器会展示 title 属性的内容，例如：

```javascript
乘风破浪这个词，语出南朝名将<abbr>宗悫</abbr>
```

### 2.12 del

del 标签用于定义带有删除线（下划线）的文本，例如：

```javascript
<p>原价：<del>50</del>，促销价：10</p>
```

### 2.13 ins

ins 类似于 del，不同的是这个标签是用于插入新的内容，展现形式是文本下边加上下划线，例如：

```javascript
<p>一打有 <del>二十</del> <ins>十二</ins> 件。</p>
```

### 2.14 code

code 标签用于展示计算机编程代码或者伪代码，专为软件开发人员设计的，文本将用等宽、类似电传打字机样式的字体（Courier）显示出来，例如：

```javascript
<code>

    document.getELementById("test");

</code>
```

### 2.15 meter

meter 元素用于度量给定范围内的数据，例如：

```javascript
<meter value="20" min="0" max="100">百分之二十</meter>
```

### 2.16 progress

progress 标签是用于定义进度条，HTML5 之前的版本都是需要用 div 或者其他标签配合 css 以及 JavaScript 才能实现出来滚动条效果，现在只要定义一个标签就可以了，例如：

```javascript
<progress value="22" max="100"></progress>
```

max 属性用于表示滚动条的最大长度，value 值表示当前完成了多少。

## 3. 语义化延伸

实际项目中应尽量按照如下标准，做到代码易扩展、易维护：

* 尽量做到标签语义化，少量使用没有语义的标签，例如 div、span；
* 熟悉每个标签的属性规范，属性值命名应当浅显易懂；
* 网页尽量使用结构化，例如区分头部、内容、尾部；
* 样式与内容分离，样式应尽力放到 css 文件中；
* 脚本 JavaScript 尽量与内容分离，包含到 JavaScript 引用中；
* 复杂的代码需要使用注释；
* 尽量使用 w3c 定义的标准语法，避免出现浏览器兼容性问题

## 4. 举例

如下 div 布局及结构标签布局两个例子，在网页中展现一模一样。明显结构标签布局语义行更强，便于开发者理解和阅读：

```javascript
<html lang="zh-cn">
    <head>
        <title>Insert a title</title>
        <meta charset="utf-8">
    </head>
    <body>
        <div id="header">顶部</div>
        <div id="nav">导航</div>
        <div id="banner">内容</div>
        <div id="main">
            <div id="left_side">左边栏</div>
            页面主体
        </div>
        <div id="footer">页面底部</div>
    </body></html>
```

```javascript

<html lang="zh-cn">
    <head>
        <title>Insert a title</title>
        <meta charset="utf-8">
    </head>
    <body>
        <header>顶部</header>
        <nav>导航</nav>
        <div>内容</div>
        <section>
            <aside>左边栏</aside>
            页面主体
        </section>
        <footer>页面底部</footer>
    </body></html>
```

虽然使用 div 通过使用 css 样式可以实现大部分标签的效果，但是并不建议这样使用

## 5. 小结

本章介绍了 HTML 的语义化概念，罗列出了一部分有明确定义的语义化标签，总结了在项目开发中使用语义化有什么优点，以及通过语义化的概念进一步延伸到了实际项目开发中需要注意到的扩展性、可维护性等问题，最后通过实际举例来对比语义化和非语义化的实际代码差别。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
