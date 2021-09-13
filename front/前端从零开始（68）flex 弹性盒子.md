## flex 弹性盒子

`display:flex` 和接下来我们介绍的这个 flex 是有区别的，前者是修改`display`实现弹性和模型的，而 flex 仅仅是弹性盒模型下 flex-grow、flex-shrink 和 flex-basis 三个的缩写属性，用来定义和模型内部的子元素在浏览器重的展示形式。 下面我们主要讲这三个属性。

## 1\. 官方定义

属性用于设置或检索弹性盒模型对象的子元素如何分配空间。
flex 属性是 `flex-grow`、`flex-shrink` 和 `flex-basis` 属性的简写属性。

## 2\. 慕课解释

fl 父元素设置成 `dispaly：flex` 之后子元素的空间分配通过 `flex` 设置，其特点为弹性，即内部分配空间如果按照比例分配则其不会随着父元尺寸变化而变化。

## 3\. 语法

子元素

```css
{
    flex: flex-grow flex-shrink flex-basis|auto|initial|inherit|none;
}
```

|参数名称|参数类型|解释|
|--------|--------|----|
|flex-grow|number|其它子元素的比例关系默认为 0 ，存在剩余空间不扩大|
|flex-shrink|number|默认为 1 空间不足时候缩小|
|flex-basis|\| ‘auto’|设定一个长度或者自动填充|

## 4\. 兼容性

flex:

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|----|----|---|---|---|---|---|---|
|–|–|63-74|84-85|-|-|-|-|

flex-grow| flex-shrink|flex-basis:

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|----|----|---|---|---|---|---|---|
|10+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 给一个块级元素添加 flex 属性 ，让其子元素平均分配空间。

```html
<div class="demo">
    <div class="item">1</div>
    <div class="item">2</div>
    <div class="item">3</div>
</div>
```

```css
.demo{
    display:flex;
    width:200px;
    height:60x;
    line-height:60px;
    border: 1px solid #ccc;
    border-right: none;
}
div>.item{
    width:100px;
    border-right:1px solid #ccc;
    text-align: center;
    flex:1;
}
```

效果图

![块级元素平均分配空间](https://img.mukewang.com/wiki/5eb0fdc1096fc75402220076.jpg)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style>
    .demo{
	    display:flex;
	    width:200px;
	    height:60x;
	    line-height:60px;
	    border: 1px solid #ccc;
	    border-right: none;
	}
	div>.item{
	    width:100px;
	    border-right:1px solid #ccc;
	    text-align: center;
	    flex:1;
	}
    </style>
    <body>
    <div class="demo">
	    <div class="item">1</div>
	    <div class="item">2</div>
	    <div class="item">3</div>
	</div>
    </body>
</html>
```

解释：容器 demo 设置了 flex 总宽度为 200px，项目 item 设置宽度 100px；如果正常情况下会超出容器，我们通过设置 `flex:1` 让项目自适应容器，并等分了空间。

2. 给一个块级元素添加 inline-flex 属性，让其变成行内元素，子元素平均分配

```css
.demo{
    display:inline-flex;
    width:200px;
    height:60x;
    line-height:60px;
    border: 1px solid #ccc;
    border-right: none;
}
div>.item{
    width:100px;
    border-right:1px solid #ccc;
    text-align: center;
    flex:1;
}
```

效果图

![内联元素平分空间效果图](https://img.mukewang.com/wiki/5eb0febe09a6006602510074.jpg)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
	.demo{
    display:inline-flex;
    width:200px;
    height:60x;
    line-height:60px;
    border: 1px solid #ccc;
    border-right: none;
}
div>.item{
    width:100px;
    border-right:1px solid #ccc;
    text-align: center;
    flex:1;
}
	</style>
</head>
<body>
     <div class="demo">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
    </div>
    慕课
</body>
</html>
```

demo 和文字在一行，demo 变成内联元素了。

3. 一个左侧 100px，右侧自适应的，左右布局

```html
    <div class="demo-2">
        <div class="item-left">1</div>
        <div class="item-right">2</div>
    </div>
```

```css
.demo-2{
    display:flex;
}
.item-left{
    flex-basis: 100px;
}
.item-right{
    flex-grow:1;
}
```

效果图

![左侧100px，右侧自适应的，左右布局效果图](https://img.mukewang.com/wiki/5eb0ffa509d1c40a06340113.jpg)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
    .demo-2{
    display:flex;
}
.item-left{
    flex-basis: 100px;
}
.item-right{
    flex-grow:1;
}
    </style>
</head>
<body>
    <div class="demo-2">
        <div class="item-left">1</div>
        <div class="item-right">2</div>
    </div>
</body>
</html>
```

4. 一个左侧为 100px，右侧最大为 600px 的左右布局

```css
.demo-2{
    display:flex;
}
.item-left{
    flex-basis: 100px;
    background: red;
    flex-shrink:0;
}
.item-right{
    flex-basis: 600px;
    background: yellow;
}
```

![左右布局](https://img.mukewang.com/wiki/5eb1044a0958c0e108060109.jpg)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
		.demo-2{
    display:flex;
}
.item-left{
    flex-basis: 100px;
    background: red;
    flex-shrink:0;
}
.item-right{
    flex-basis: 600px;
    background: yellow;
}
	</style>
</head>
<body>
    <div class="demo-2">
        <div class="item-left">
1
        </div>
        <div class="item-right">
2
        </div>
    </div>
</body>
</html>
```

解释：右侧最大宽度为 600，如果小于 600 右侧将随屏幕尺寸缩小。

## 6\. 经验分享

现在的很多前端开发都会接到一些设计稿，要求在各种终端都可以适应，那么用好 `flex` 是一个关键。 `flex:1` 是其中最常见的设置，它等价于：

```css
.demo{
    flex-grow:1;
    flex-shrink:1;
    flex-basis:auto
}
```

其意思就是剩余空间就扩大，而剩余空间不足就缩小，就像弹簧一样。那么这部分就可以自适应各种屏幕大小了。

## 7\. Tips

1. `flex-basis` 和 `flex-grow` 同时使用时候 `flex-basis` 不起作用。

2. `flex` 的属性 默认是 0 1 auto，它们的顺序是 flex-grow flex-shrink 和 flex-basis 即三不：有剩余空间不扩大、当空间不足时缩小、不限制尺寸。

3. `flex` 属性有两个快捷值 即 auto（ 1 1 auto）和 none（0 0 auto）。

4. 尽量不要使用缩小，因为它的兼容性不是很好。
