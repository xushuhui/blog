# 认识自定义文本内容标签 span 标签

我们在上一章节已经学过了段落标签 P 标签，而且我们也学习了 P 标签是可以嵌套其他标签的。当我们定义的段落里，有一些文字或者一句话有单独的样式时，则我们需要单独的处理这些内容，最好的方式则是用 span 标签将这些内容包裹起来。例如：

![](https://xushuhui.gitee.io/image/imooc/5f07be6909f86af106710209.jpg)

## 1. SPAN 标签的作用

span 标签用来定义文本内容，可以是单独的一句话，一些内容，也可以是段落里面的内容。如果不对 span 应用样式，那么 span 元素中的文本与其他文本不会任何视觉上的差异。尽管如此，上例中的 span 元素仍然为 p 元素增加了额外的结构。span 标签是不带任何默认样式的，所以它极大的提高了我们可以对文本内容自定义样式的需求。

## 2. SPAN 标签的用法

span 标签为双标签，所以必须有首尾标签。文本的内容既为 span 标签的内容。例如：

```javascript
<span>
  我是一个文本内容
</span>
```

SPAN 标签没有默认样式，所以上述内容会在页面上呈现以下效果：

![](https://xushuhui.gitee.io/image/imooc/5f07be880965550c01860050.jpg)

## 3 .SPAN 标签的特点

1. span 标签为行内元素，行内元素和块级元素（p 标签）不同，默认是在同一排排列，如我们写两个 span 标签，会呈现以下效果：

![](https://xushuhui.gitee.io/image/imooc/5f07be990934605903150050.jpg)

2. span 不能使用 CSS 为其设置宽高，即使**设置了宽高也会无效**。

3. span 标签里也可以嵌套其他标签。例如：

```javascript
  <span>
    <p>
      我是 span 标签里的 p 标签
    </p>
  </span>

```

4. span 标签的应用场景多数为为某些内容单独设置样式，我们可以用 span 标签将这些包裹起来，这样既单独设置了样式，也不会影响其他内容。当然，如果你想内容在同一排排列，也可以使用 span 标签包裹这些内容。

## 4. 经验分享

1. 如果我们这样书写代码，每段代码都换行，例如：

```javascript
   我是其他内容
  <span>我是第一个 span 标签</span>
```

那么两个 span 标签之间左右会有间距，如下图所示：

![](https://xushuhui.gitee.io/image/imooc/5f07beb109abb33204150037.jpg)

如何解决这个问题呢，我们只需要讲所有代码写在同一行即可，如：

```javascript
我是其他内容<span>我是第一个 span 标签</span><span>我是第二个 span 标签</span>
```

那么问题就解决了，如下图所示：

![](https://xushuhui.gitee.io/image/imooc/5f07beda0923edd904000032.jpg)

2. span 标签为行内元素，不能对其设置宽高属性，如果既想保持行内元素在同一行排列的特性，又想为其设置宽高，我们可以利用 CSS 的 display 属性将其设置为行内块元素即可。

## 5. 真实案例分享

慕课手记技术文章。

前端开发规范（节选）：

```javascript
<p>
	<span>关注</span>
</p>

<div>
  <span> JAVA 开发工程师 </span>
</div>
<div>
	<a>
		<span></span> 篇手记
	</a>
	<a>
		贡献 <span></span> 字
	</a>
</div>
```

新浪新闻：

```javascript
<div>
	 <span >2020 年 06 月 14 日 13:00</span>
</div>
<div>
	<span> 缩小字体 </span>
	<span> 放大字体 </span>
	<span> 收藏 </span>
	<span> 微博 </span>
	<span> 微信 </span>
	<span> 分享 </span>
	<span >0 </span>
<div>
	<span> 腾讯 QQ </span>
	<span> QQ 空间 </span>
</div>
```

## 6. 小结

1. span 标签为双标签，它总是成对出现的，需要首尾标签；

2. span 用于对文档中的行内元素进行组合；

3. span 标签没有固定的格式表现。当对它应用样式时，它才会产生视觉上的变化。如果不对 SPAN 应用样式，那么 SPAN 元素中的文本与其他文本不会任何视觉上的差异；

4. span 标签提供了一种将文本的一部分或者文档的一部分独立出来的方式；

5. span 标签不会自动换行，他们会在同一行显示，但是左右会有间隙。如想解决此问题，把代码书写在一行即可。

![](https://xushuhui.gitee.io/image/imooc/5f62fed109c3cb1210370732.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
