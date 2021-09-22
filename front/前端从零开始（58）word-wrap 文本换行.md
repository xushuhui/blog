# word-wrap 换行

在了解这个属性前，首先要说的是它只用于英文。

## 1. 官方定义

word-wrap 属性允许长单词或 URL 地址换行到下一行。

## 2. 慕课解释

当一个英文单词，或者一段很长且中间没有空格的英文字母的文本超出文本所在元素边缘时候，直接将超出的部分换行，而不是把这个连续的文本直接全部另起一行。

## 3. 语法

```javascript
word-wrap: normal|break-word;
```

```javascript
.demo{
    word-wrap:break-word;
}
```

属性值

|值|说明|
|--|----|
|normal    |就是按照浏览器自己的排版规则，不设置就是默认。                              |
|break-word|当连续的英文字符超过元素的宽度时候直接折行，而不是把整个连续的英文单词换行。|

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
    word-wrap: break-word;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e4a7094efd9f01880228.jpg)

对超出元素区域的的文本换行效果图

说明： 上图是没有使用换行属性的效果。下图使用了换行属性。

## 6. 经验分享

这个属性有一个近似属性`wrod-break`, 我们通过表面意思进行区分。`wrod-wrap`主要是换行，`wrod-break`是用来英文单子怎么在内断开的。

最后我们在一次对比下这两个属性

```javascript
<div class="demo">
    class.imooc.com class.imooc.com
</div>
<div class="demo demo-1">
    class.imooc.com class.imooc.com
</div>
<div class="demo demo-2">
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
    word-wrap: break-word;
}
.demo-2{
    word-break:break-all;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e4b309c91cd704120112.jpg)

`wrod-wrap`对比 `wrod-break` 效果图

说明 从左往右，第一张图什么都不设定，第二张图设置`word-wrap: break-word`, 它在空格处换行了。而第三章图设置`word-break:break-all;`在空格处没有换行而是打破了连续的单词。

## 7. 小结

1. 必须要是连续的英文字符。
2. 这个属性主要是对英文起作用，如果中英文混杂，则优先把中文和英文先换行，然后再打破连续的英文单词。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
