# letter-spacing 字间距

当我们需要扩大或者减少字符之间的距离时候，会用到这个属性。如果我们需要调整字符间的距离可以用到这个属性。

## 1. 官方定义

letter-spacing 属性增加或减少字符间的空白（字符间距）。

## 2. 慕课解释

该属性定义了在文本字符框之间插入多少空间。由于字符字形通常比其他字符框要窄，指定长度值时，会调整字母之间通常的间隔。因此，normal 就相当于值为 0。

这个属性常用来修改文字之间的距离，它允许为负值，默认的字符间的距离为 0。如果数值小于 0 字符会紧凑，大于 0 时越大越松散。

## 3. 语法

```javascript
.demo{
    letter-spacing: value;
}
```

值说明

|参数名称|参数类型|
|--------|--------|
|value|‘px’ | ‘rem’ | ‘em’|

通过上面图片，我们可以看到从左到右每个字符的右侧都增加了 10px。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

1. 字符之间增加 5px 的间距。

```javascript
.demo{
    letter-spacing:5px;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e55e089cbef402260074.jpg)

 字符之间增加 5px 的间距效果图

1. 设置字符之间距离为 -2px。

```javascript
.demo{
    letter-spacing:-2px;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e56c09a2a59902130036.jpg)

 设置字符之间距离为 -2px 效果图

1. 通过使用 letter-spacing 清除我们编辑 html 代码时候元素设定 inline-block 之后出现的空格。

```javascript
<div class="demo">
    <span>慕课网</span>
    <span>学习</span>
</div>
```

```javascript
.demo{
    letter-spacing:-5px;
}
.demo>span{
    display: inline-block;
    width: 50px;
    background: red;
    letter-spacing:normal
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e578096aff2501200085.jpg)

 使用 letter-spacing 清除空格效果图

上图是 span 只设定了 `display: inline-block;`，下图我们通过 在父级元素上设置`letter-spacing`去除了`span` 之间空格的距离。而在`span`中又重置了`letter-spacing`让文字恢复了它们之间的距离，让其不在拥挤。**不过我们不推荐这种方式，建议内联块级元素不要分行。**

## 6. 经验分享

与 word-spacing 的区别：

`word-spacing` 这个属性只能作用英文，意思是两个英文单词之间的距离，这里要注意是‘英文单词’而不是‘字符’，而 `letter-spacing` 没有任何限制可以作用于‘任何字符’。

## 7. 小结

1. 这个属性仅仅对字符起作用，不能作用于元素标签上面。

2. 这个属性不可以是 % 这样的计量单位，因为它不是一个距离，没有相对点，浏览器不知道该如何解释。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
