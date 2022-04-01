## flex 弹性盒子布局介绍

flex 布局可以说是目前为止最好用的布局方式，但是目前还稍微有一点受到兼容性的影响，它对 IE9 不兼容，但是在未来随着 IE9 逐渐被淘汰，我相信，它一定会在布局这块大放异彩，因为它实现了太多我们曾经不能实现的布局效果，而且只要简单的几个属性就可以搞定！

## 1. 官方解释

一种弹性盒模型布局方式。

## 2. 慕课解释

flex 布局也叫弹性布局，它的特点是可以实现子元素的自适应屏幕大小，可以自由的分配每个 box 需要占用的空间比例。我们把父元素称作为：容器。子元素称作为：项目。容器默认存在两个轴：水平主轴（mian axis）、垂直交叉轴（cross axis）。左侧是主轴的开始点，右侧是主轴的结束点，垂直方向上顶部是交叉轴的开始位置，底部是交叉轴的结束位置。

## 3. 语法

通过下面两种形式都可以实现弹性盒模型“容器”的初始化。

1. 块级弹性模块。

```javascript
div{
    display:flex;
}
```

1. 内联弹性模块。

```javascript
div{
    display:inline-flex;
}
```

容器包含属性 点击查看详细

|参数名称|参数|解释|
|--------|----|----|
|flex-direction |row | row-reverse | column | column-revers                             |定义主轴上项目的的方向              |
|flex-wrap      |nowrap | wrap | wrap-reverse                                           |定义项目如何换行                    |
|flex-flow      |< flex-direction > | < flex-wrap >                                     |前两个属性的简写                    |
|justify-content|flex-start | flex-end | \center | space-between | space-around         |定义主轴（水平）上项目的对齐方式      |
|align-items    |flex-start | flex-end | center | baseline | stretch                    |定义交叉（垂直）方向上项目的对齐方式|
|align-content  |flex-start | flex-end | center | space-between | space-around | stretch|多轴（多行）下项目的（水平）对齐方式  |

项目包含属性 点击查看详细

|参数名称|参数|解释|
|--------|----|----|
|flex-grow  |number                                                    |            |
|flex-shrink|number                                                    |            |
|flex-basis |像素                                                      |            |
|flex       |||                                                        |            |
|order      |number                                                    |            |
|align-self |auto | flex-start | flex-end | center | baseline | stretch|修改单个项目|

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|10+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 创建一个弹性盒模型，容器为块级，项目自适应。

```javascript
.demo{
    display:flex
}

```

1. 创建一个行内盒模型。

```javascript
.demo{
    display:inline-flex
}
```

## 6. 小结

1. 需要父元素首先设置成 `dislpay:flex` 这样子元素才能起作用，而子元素的 `float` 、 `clear` 、 `vertical-align` 属性都失去作用。

2. 子元素可以使用 `position` 来脱离 flex 布局。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
