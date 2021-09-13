# 认识 textarea 文本域标签

文本域既表示一个文本的区域，通俗来说就是可以一个区域内，可以输入多行文本，普通的输入框只能输入单行文本。文本区中可容纳无限数量的文本，其中的文本的默认字体是等宽字体。比如网站中的评论、留言等功能都可以使用 textarea 文本域标签来实现。

## 1. textarea 文本域标签的使用

我们直接书写 textarea 标签既代表文本域标签，代码如下：

```javascript
<textarea></textarea>
```

效果如下：

![20200709154411323](https://img.mukewang.com/wiki/5f07cede097177db02960114.jpg)

我们可以通过可以通过 COLS 和 ROWS 属性来规定 textarea 的尺寸。cols 设置文本域的宽度，rows 设置文本域的高度。代码如下：

```javascript
<textarea cols="30" rows="10"></textarea>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cef1098548e504410266.jpg)

需要注意的是，文本域的右下角有一个可以拖拽的区域，可以通过拖拽来改变文本的宽高，如果想要取消这个功能，需要通过 CSS 样式来取消。

文本域也可以设置 `placeholder` 属性来实现占位符的效果，用法和作用和 input 的 `placeholder` 属性一样。

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cf0609e1f03304310278.jpg)

如果在文本域标签当中写入了内容，那么这些内容会显示文本域的区域内。代码如下：

```javascript
<textarea cols="30" rows="10" placeholder="请输入内容">123456</textarea>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cf1609b3bffd04710262.jpg)

## 2. 注意事项

1. 文本域可以设置 cols 和 rows 属性来改变文本域的宽高，不过更好的办法是使用 CSS 的 height 和 width 属性；
2. 文本域默认右下角有一个可以拖拽区域，通常情况下我们不需要这个功能，所以需要借助 CSS 来取消这个功能。

## 3. 真实案例分享

简书

```javascript
<textarea placeholder="写下你的评论..."></textarea>
```

慕课网

```javascript
<textarea placeholder="写下你对评价的回复..."></textarea>
```

## 4. 小结

1. 文本域为双标签，如果在文本域的标签当中写内容，则内容会出现在文本域中。
2. 文本域可以通过设置 cols 和 rows 来改变文本域的宽高。
3. 文本域可以设置`placeholder` 属性来添加占位符。

![图片描述](https://img.mukewang.com/wiki/5f6305ea09e4b5b214070662.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
