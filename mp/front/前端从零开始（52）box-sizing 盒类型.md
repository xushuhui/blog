# box-sizing 盒类型

它主要用来告诉浏览器怎么计算元素的展示宽高的。

## 1. 官方定义

box-sizing 属性允许您以特定的方式定义匹配某个区域的特定元素。

## 2. 慕课解释

在开发过程中，当有一个 固定宽高的元素带有 boder 或 padding 的时候，它在 IE 中展示和 Chrome 中是不同的，如果我们不设置这个属性，那么在 IE 浏览器中它的实际宽高是小于 Chrome 浏览器的这是为什么呢？

Chrome 浏览器使用的是标准盒模型 content-box，IE 盒模型是 border-box。

下面就是一个盒模型的结构图

![](https://xushuhui.gitee.io/image/imooc/5ea3da6808fbe08505000450.jpg)

## 3. 语法

```javascript
box-sizing: content-box | border-box
```

它接受一个参数 content-box 或 border-box。

![](./img/盒模型.jpg)

上图是一个盒模型结构

**content-box 的计算方式是：**

width = content width;

height = content height

**border-box 的计算方式是：**

width = border + padding + content width

heigth = border + padding + content heigth

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|10+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 不是设置 `box-sizing` 分别在不同浏览器中的展示。

```javascript
<div class="demo"> CSS3 学习分享</div>
```

```javascript
.demo{
    width:100px;
    height:100px;
    background:#000;
    color:#fff;
    padding:10px;
    border:5px solid red;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3da0609b50c5702600150.jpg)

不同浏览器中的展示效果图

说明：左侧是 Chorme 浏览器右侧是低版本 IE 浏览器。

1. 给上面 demo 中设置 `box-sizing` 为 `border-box`。

```javascript
.demo{
    width:100px;
    height:100px;
    background:#000;
    color:#fff;
    padding:10px;
    border:5px solid red;
    box-sizing:border-box;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3da1b0969371f01480144.jpg)

使用`border-box`效果图

1. 给上面 demo 中设置 `box-sizing` 为 `content-box`。

```javascript
.demo{
    width:100px;
    height:100px;
    background:#000;
    color:#fff;
    padding:10px;
    border:5px solid red;
    box-sizing:content-box;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3da290911f80501870164.jpg)

使用`content-box`效果图

## 6. 经验分享

推荐大家设置 `box-sizing` 为`border-box` 这样方便我们写样式不必在去减去 `padding` 也不会造成 IE 和 Chorme 这类浏览器展示不同的 bug 。

## 7. 小结

如果不设置 `box-sizing` 不同浏览器会有不同的计算方式。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
