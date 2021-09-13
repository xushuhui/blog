# HTML 结构化元素标签

网页是由一个一个的区域组成，就像房子中的房间一样，每个区域内可以放置 HTML 元素，然后可以对区域设置样式从而将其与别的区域区分开来。这个区域在 HTML 中由结构元素实现，又可以将其称之为容器元素。通过搭配 ID 属性使用，结构元素将会帮助我们构建结构化条理分明的网站逻辑。本章我们介绍一下 结构元素的使用方式和场景。

## 1. div

### 1.1 适用场景

div 元素在 HTML 中用于定义一个区块，用于对区块内的元素统一布局或者隔离，它本身并不呈现视觉样式，也没有明确的语义，可适用于没有明确语义化的场景，例如：

```javascript
<div class="AppHeader-inner">
    <ul role="navigation" class="Tabs AppHeader-Tabs">
        <li role="tab" class="Tabs-item AppHeader-Tab Tabs-item--noMeta">
        <a class="Tabs-link AppHeader-TabsLink is-active" href="//www.zhihu.com/" data-za-not-track-link="true">首页</a>
        </li>
        <li role="tab" class="Tabs-item AppHeader-Tab Tabs-item--noMeta">
        <a class="Tabs-link AppHeader-TabsLink" href="//www.zhihu.com/explore" data-za-not-track-link="true">发现</a>
        </li>
        <li role="tab" class="Tabs-item AppHeader-Tab Tabs-item--noMeta">
        <a class="Tabs-link AppHeader-TabsLink" href="//www.zhihu.com/question/waiting" data-za-not-track-link="true">等你来答</a>
        </li>
    </ul>
</div>
```

代码解释：上述代码将网页头部使用 div 区块定义，这样定义的好处在于，一旦网页需要隐藏头部，或者给头部设置一个特殊的样式（例如颜色、字体）只需要通过 JavaScript 或者 css 操控 div 即可。

这样对 div 设置文字颜色，div 中包含的元素统一生效。div 内部可以包含任意元素，理论上可以配合 css 模拟成任何样式的元素，**所以在 HTML5 之后原本跟样式相关的元素属性基本上不建议使用了，都可以用 css 替代。**

### 1.2 避免滥用 div

div 作为一个没有特殊意义的标签，相比其它有明确用途的标签，它没有特殊属性，也没有特殊效果，实在想不出还没有什么可以掌握的。但这并不妨碍 div 被滥用。一个标签什么时候应该使用，是由这个标签本身的语义，以及使用者想要表现的内容决定的。

**div 没有具体的语义**

div 的语义是 division，一个很广泛的语义，在使用 div 之前应该先想想，还有没有更具语义的标签，尽量使用语义明确的标签。比如用 div 来标记文段（p）和列表（li）显然是不合适的。

事实上，为了更好地使用 div，你需要全面了解 HTML 标准提供了哪些标签，它们的语义是什么。如果你不知道 p、span、h1、h2、h3 的存在，只会用 div，必定写出糟糕的 HTML 代码。

```javascript
<div class="container" id="header">
    <div class="header header-main"></div>
    <div class="site-navigation">
        <a href="/">Home</a>
        <a href="/about">About</a>
        <a href="/archive">Archive</a>
    </div>
</div>
<div class="container" id="main">
    <div class="article-header-level-1">    </div>
    <div class="article-content">
        <div class="article-section">
            <div class="article-header-level-2">            </div>
        </div>
        <div class="article-section">
            <div class="article-header-level-2">            </div>
        </div>
    </div>
</div>
<div class="container" id="footer">
    <div class="contact-info">
        <p class="email">
            <a href="mailto:us@example.com"></a>
        </p>
        <div class="street-address">        </div>
    </div>
</div>
```

代码解释：例如上述代码，完全使用 div 来定义网页区域，这样做虽然可以实现对应的效果，但是对于搜索引擎、爬虫，甚至网站维护人员来说是一个糟糕的代码，因为搜索引擎识别不出哪部分是头部哪部分是尾部。

```javascript
<header>
    <h1>Super duper best blog ever</h1>
    <nav>
        <a href="/">Home</a>
        <a href="/about">About</a>
        <a href="/archive">Archive</a>
    </nav>
</header>
<main>
    <article>
    <header>
        <h1>Why you should buy more cheeses than you currently do</h1>
    </header>
    <section>
        <header>
            <h2>Part 1: Variety is spicy</h2>
        </header>
        <!-- cheesy content -->
    </section>
    <section>
        <header>
            <h2>Part 2: Cows are great</h2>
        </header>
        <!-- more cheesy content -->
    </section>
    </article>
</main>
<footer>
    <section class="contact" vocab="http://schema.org/" typeof="LocalBusiness">
        <h2>Contact us!</h2>
        <address property="email">
            <a href="mailto:us@example.com"></a>
        </address>
        <address property="address" typeof="PostalAddress">
            <p property="streetAddress"></p>
            <p property="addressCountry"></p>
        </address>
    </section>
</footer>
```

代码解释：这段代码使用 header 、nav、footer、section 等标签，详细且准确的根据语义化来定义网页的布局。

## 2. 其他语义化相关机构元素

### 2.1 span

和 div 的作用类似，span 是通常用来作为文本容器，它没有明确定义的语义，在没有其他合适的元素标签时，可以用 span 来包含文本，它是一个行内元素。例如：

```javascript
<p><span>一些文字</span></p><!-- 定义一个段落 -->
```

### 2.2 header

header 元素通常用于包含头部介绍性相关的元素，是 HTML5 中新增元素。它可以使用全局属性。一个网页中可以使用多个 header 元素，也可以在每个独立的内容机构中使用 header 元素，例如：

```javascript
<header>
  <h1>网站标题</h1>
  <img src="logo-sm.png" alt="logo">
</header>
```

> **注意：** header 是 HTML5 新增的元素，可能不兼容低版本的 IE 浏览器

### 2.3 nav

nav 是 Navigation 的简称，通常用于网页中的导航栏，使用 nav 将网页中导航栏相关的标签归拢在一个区域，这样结构更清晰明了，例如：

```javascript
<nav><!-- 定义一个导航栏 -->
    <ul>
        <li><a href="#">首页</a></li>
        <li><a href="#">新闻资讯</a></li>
        <li><a href="#">常见问题</a></li>
        <li><a href="#">更新日志</a></li>
        <li><a href="#">论坛</a></li>
    </ul>
</nav>
```

> **注意：** nav 是 HTML5 新增的元素，可能不兼容低版本的 IE 浏览器

### 2.4 article

article 用于定义网页中文档、章节、段落相关的文本结构，一个网页中可以定义多个 article 标签。与 span 不同的是，article 具有明确定义的语义。每个 article 中通常包含 header、H1-H6、address 等和文章标题、作者等相关信息的标签。

```javascript
<article>　<!-- 定义一篇文章 -->
    <header>
        <h1>article元素使用方法</h1>
        <p>发表日期：<time pubdate="pubdate">2017/2/9</time></p>
    </header>
    <p>此标签里显示的是article整个文章的主要内容</p>
</article>
```

> **注意：** article 是 HTML5 新增的元素，可能不兼容低版本的 IE 浏览器。

### 2.5 aside

aside 元素用来定义当前页面或者文章的复数信息部分，它主要包含相关的引用、侧边栏、广告、导航条等其他类似的有别于主要内容的部分，通常结合 article 标签使用。

### 2.6 section

section 元素用于对网页中的内容分块，与 div 不同的是它具有明确的分块语义，但是 HTML 没有对它定义更加细化的语义，所以说如果适用场景有比较明确的语义场景的话应该使用更加符合条件的标签，例如 nav、article、header 等。如果需要对块使用自定义的样式，请尽量使用 div 标签。使用 section 通常用于将一块内容细分成多个小块区域，例如：

```javascript
<article>
    <h1>编程语言</h1>
    <p>编程语言常用的有 asp,asp.net,php,jsp </p>
    <section>
        <h2>asp</h2>
        <p>asp全称Active Server Page</p>
    </section>
    <section>
        <h2>asp.net</h2>
        <p>asp升级版</p>
    </section>
    <section>
        <h2>php</h2>
        <p>简单、入门门槛低</p>
    </section>
</article>
```

> **注意：** section 是 HTML5 新增的元素，可能不兼容低版本的 IE 浏览器。

### 2.7 footer

footer 元素用于定义网页的页尾，与 header 元素类似，除了可以定义网页的页尾，还可以用于 article 元素的尾部，例如：

```javascript
<article>
    <header>文章目录</header>
    <p>文章内容</p>
    <footer> 文章信息 </footer>
</article>
<footer> 页面底部 </footer>
```

> **注意：** footer 是 HTML5 新增的元素，可能不兼容低版本的 IE 浏览器。

## 3. 小结

本章介绍 HTML 中特殊的标签 - 结构化标签，由于结构化标签在默认状态下不能呈现相关的视觉样式，所以通常不太起眼，但是想要写出一个结构分明、易扩展、易维护的网页代码，它是必不可少的。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
