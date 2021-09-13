# word-break 文本打断

这个属性主要用来处理英文单词，在超出一行之后如何换到下一行的规则。

## 1. 官方定义

word-break 属性规定自动换行的处理方法。

## 2. 慕课解释

一段英文段落，在其文本所在的元素边缘通常都会把整个单词另起一行，而这个属性可以打破这种排版方式，让这个段落的英文单词都是分开的，同汉字一样，在元素的边缘只是最后一个字母换行。

## 3. 语法

```javascript
word-break: normal|break-all|keep-all;
```

```javascript
.demo{
    word-break:break-all;
}
```

属性值

|值|说明|
|--|----|
|normal   |就是按照浏览器自己的排版规则，不设置就是默认。                              |
|break-all|其意义就同英文直接翻译一样，打破所有的英文单词，可以在任意的字母处另起一行。|
|keep-all |只能在半角空格或连字符处换行。                                              |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

1. 对超出元素区域的的文本换行。

```javascript
<div class="demo">
    class.imooc.com class.imooc.com
</div>
<div class="demo demo-1">
    class.imooc.com class.imooc.com
</div>
```

```javascript
.demo{
    background: #ccc;
    width: 100px;
    height: 100px;
    margin-bottom: 10px;
}
.demo-1{
     word-break:break-all;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea3e41509a121db01580223.jpg)

对超出元素区域的的文本换行效果图

说明： 上图是没有使用换行属性的效果。下图使用了换行属性

1. 仅对段落中的半角空格和连字符进行换行。

```javascript
<div class="demo-2">
    class imooc-com class imooc-com classimooccom
</div>
```

```javascript
.demo-2{
    background: #ccc;
    width: 100px;
    height: 100px;
    word-break: keep-all;
}

```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea3e42209d37f2d01960225.jpg)

对段落中的半角空格和连字符进行换行效果图

说明： 如图第一行的结尾使用了连字符所以后面的英文字符换行了，第二行使用了空格所以后面的也换行了，而第三行没有空格或连字符因此没有换行。

## 6. 经验分享

这个属性用来处理当我们不想让 一个英文单词直接下一行，而是从中间断开，断开的地方换行例如我们使用连字符的时候。

## 7. 小结

1. 这个和 `word-wrap`有区别， `word-wrap`必须要是连续的英文字符，而它没有限制，所以不要记混。
2. 这个属性对英文字符、半角空格、连字符都起作用。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
