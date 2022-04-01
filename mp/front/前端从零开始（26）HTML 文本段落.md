# 文本段落列表

本章介绍文本段落等控制文字章节显示的标签。这些标签通常只起到展示排版作用（当然还有 SEO 的作用），其他的用途不多。

## 1. 标题

用于标题的 HTML 标签包括`<h1> - <h6>` 来定义。标题类似于 Word 中的标题，其作用是为了对文章进行排版，而不是只为了放大字号。良好的标题排版对搜索引擎比较友好。

```javascript
<h1>一级标题</h1>
<h2>二级标题</h2>
<h3>三级标题</h3>
<h4>四级标题</h4>
<h5>五级标题</h5>
<h6>六级标题</h6>
```

展示效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04b0b0946b29e02990347.jpg)

标题标签元素在 HTML5 使用较少，一般使用 css 样式控制文字大小。

## 2. 文档标题

title 用于定义文档的标题，通常嵌入到 header 头里边，其作用是：

* 定义在浏览器中显示的文档标题
* 当网页被添加到收藏夹时，显示在收藏夹的标题
* 显示在搜索引擎中的标题

```javascript
<title>HTML 页面的标题</title>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04b3e09a4a08b08280168.jpg)

title 标题显示在浏览器左上角

## 3. 段落

p 标签用于定义段落，它是一个块级元素，浏览器会自动在其前后换行

```javascript
<p>麻叶层层苘叶光，谁家煮茧一村香。</p>
<p>隔篱娇语络丝娘。</p>
<p>垂白杖藜抬醉眼，捋青捣麨软饥肠。</p>
<p>问言豆叶几时黄。</p>
```

上述代码段使用 p 段落标签分隔了一首宋词浣溪沙，使用 div+css 可以实现相同效果

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04b660961e02003860190.jpg)

## 4. 列表

### 4.1 有序列表

有序列表类似于 Word 中的有序列表，使用 ol 标签定义， li 标签定义列表项

```javascript
<h2>HTML手册</h2>
<ol> <!-- 定义一个有序列表 -->
<li>第一章.语法介绍</li> <!-- 列表项1 -->
<li>第二章.标签</li> <!-- 列表项2 -->
<li>第三章.属性</li>
</ol>
```

在浏览器中展示如下

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04b8e09e0c32702670161.jpg)

其中有序列表中可以定义 start 属性，用于设置列表序号的起始值

### 4.2 无序列表

ul 无序列表不同于有序列表的是 - 列表中的编号使用粗体原点表示，而不是数字其可以通过定义 compact 和 type 来设置编号的样式，type 的可选值有 disc、square、circle。但是一般建议通过 css 统一控制样式。

```javascript
<h2>世界有几大洲：</h2>
<ul>
    <li>亚洲</li>
    <li>欧洲</li>
    <li>北美洲</li>
    <li>南美洲</li>
    <li>非洲</li>
    <li>大洋洲</li>
    <li>南极洲</li>
</ul>
```

上述代码展现的列表：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04bce09062f8502630282.jpg)

### 4.3 自定义列表

自定义列表通过 dl 标签定义，使用 dt 定义标题， dd 定义内容。自定义列表可以实现类似 table 的效果。

```javascript
<h3>自定义列表</h3>
        <dl>
            <dt>一级</dt>
                <dd>二级</dd>
                <dd>二级</dd>
            <dt>一级</dt>
                <dd>二级</dd>
                <dd>二级</dd>
        </dl>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04bf5096e87e602500250.jpg)

# 5. 总结

本章主要介绍了几种文本段落列表的控制标签以及属性，这些标签大部分用于文本排版。其实在在 HTML5 中，通过 css3+div 可以实现绝大多数效果，只不过需要对 css 非常熟悉才行，所以在不是非常熟悉 css 的情况下可以使用一些快捷的标签来实现相同的效果

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
