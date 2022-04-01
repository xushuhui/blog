Grid 与 Flex 布局有一定的相似性，但是功能更加强大，学习起来也有不少难度，不过相信下面的内容会帮你更快的掌握 Grid。

## 1. 官方定义

通过设置 `display: grid;` 可以定义一个 CSS 网格。然后使用 `grid-template-rows` 和 `grid-template-columns` 属性来定义网格的 `columns` 和 `rows`。

使用这些属性定义的网格被描述为 显式网格 (explicit grid)。

_参考文献：MDN_

## 2. 慕课解释

Grid 是一个二维网格布局，它有行 `grid-template-rows` （横排）、 列 `grid-template-columns`（竖排）, 内部的项目就分布在其中，而网格线就是行和列划分出来的。

基本属于解释：

容器：上面代码中，最外层的<div>元素`demo`就是容器。

项目：内层的三个<div>元素`item`就是项目。

行：把 `row` 即横向称为行，

列：把`column`即纵向称为列。

单元格：它们的交叉区域`cell` 也就是单元格。

网格线：`grid line`网格线就是由行和列划分出来的。

## 3. 语法

1. 块级的网格。

```javascript
.demo{
    display:grid
}
```

1. 内联级的网格。

```javascript
.demo{
    display:inline-grid;
}
```

容器包含属性如下

|属性名|值|说明|
|------|--|----|
|grid-template-columns|length                                                                      |列和每列宽度                   |
|grid-template-rows   |length                                                                      |行和每行的高度                 |
|grid-row-gap         |length                                                                      |行和行之间的距离               |
|grid-column-gap      |length                                                                      |列与列之间距离                 |
|grid-gap             |row column                                                                  |行、列间距的合并写法           |
|grid-template-areas  |string                                                                      |用来指定区域                   |
|grid-auto-flow       |row | column                                                                |默认是 row ，用来指定排列优先级|
|justify-items        |start | end | center | stretch                                              |水平方向内容的位置             |
|align-items          |start | end | center | stretch                                              |垂直方向内容的位置             |
|place-items          |align justify                                                               |垂直和水平位置合并写法         |
|justify-content      |start | end | center | stretch | space-around | space-between | space-evenly|水平方向整个内容区域的位置     |
|align-content        |start | end | center | stretch | space-around | space-between | space-evenly|垂直方向整个内容区域的位置     |
|place-content        |align justify                                                               |垂直和水平方向的合并写法       |
|grid-auto-columns    |length                                                                      |多于的网格列宽定义             |
|grid-auto-rows       |length                                                                      |多于的网格行高的定义           |

`grid-template` 是 `grid-template-columns` 、`grid-template-rows`、 `grid-template-areas` 缩写。

`grid``是 grid-template-rows`、g`rid-template-columns`、`grid-template-areas`、 `grid-auto-rows`、`grid-auto-columns`、`grid-auto-flow`的合并缩写。

**提示：gird 属性很复杂因此不推荐 `grid` 的缩写**

项目包含属性介绍

|属性名|值|说明|
|------|--|----|
|grid-column-start|number | areaName | span number|项目开始位置在左边框所在的第几根垂直网格线 |
|grid-column-end  |number | areaName | span number|项目开始位置在右边框所在的第几根垂直网格线 |
|grid-row-start   |number | areaName | span number|项目开始位置在上边框所在的第几根水平网格线 |
|grid-row-end     |number | areaName | span number|项目开始位置在下边框所在的第几根水平网格线 |
|grid-column      |number / number                |grid-column-start 和 grid-column-end 的合并|
|grid-area        |areaName                       |指定项目放在哪一个区域                     |
|justify-self     |start | end | center | stretch |单元格内容的水平方向位置                   |
|align-self       |start | end | center | stretch |单元格内容的垂直方向位置                   |
|place-self       |align-self justify-self        |单元格内容的垂直和水平位置缩写             |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|No|16+|52+|57+|10.1+|44+|10.3+|81|

## 5. 实例

本小节暂时不对父容器和子容器内的属性进行详细的实例使用展示，仅对 `display` 属性进行效果区分，可以从下一小节开始其他内容的学习。

1. 创建一个块级的 Gird 布局。

```javascript
<div class="demo">
    <div class="item">1</div>
    <div class="item">2</div>
    <div class="item">3</div>
    <div class="item">4</div>
</div>
```

通过下面的设置：

```javascript
.demo{
    display: grid;
    grid-template-columns:100px 100px;
    grid-template-rows:100px 100px;
    border:1px solid #eee
}
.item:nth-of-type(1){
    background: red;
}
.item:nth-of-type(2){
    background: green;
}
.item:nth-of-type(3){
    background: purple;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb553520960bdb903330232.jpg)

块级 Grid 布局效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo{
            display: grid;
            grid-template-columns:100px 100px;
            grid-template-rows:100px 100px;
            border:1px solid #eee
        }
        .item:nth-of-type(1){
            background: red;
        }
        .item:nth-of-type(2){
            background: green;
        }
        .item:nth-of-type(3){
            background: purple;
        }
    </style>
</head>
<body>
    <div class="demo">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    学习

</body>
</html>
```

1. 创建内联级的 Gird 布局。

```javascript
<div class="demo">
    <div class="item">1</div>
    <div class="item">2</div>
    <div class="item">3</div>
    <div class="item">4</div>
</div>
学习
```

```javascript
.demo{
    display: inline-grid;
    grid-template-columns:100px 100px;
    grid-template-rows:100px 100px;
    border:1px solid #eee
}
.item:nth-of-type(1){
    background: red;
}
.item:nth-of-type(2){
    background: green;
}
.item:nth-of-type(3){
    background: purple;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb5536109d3f2ca03030217.jpg)

内联 Grid 布局效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo{
            display: inline-grid;
            grid-template-columns:100px 100px;
            grid-template-rows:100px 100px;
            border:1px solid #eee
        }
        .item:nth-of-type(1){
            background: red;
        }
        .item:nth-of-type(2){
            background: green;
        }
        .item:nth-of-type(3){
            background: purple;
        }
    </style>
</head>
<body>
    <div class="demo">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    学习

</body>
</html>
```

## 6. 小结

1. Grid 布局是二维布局原因就是项目所在的单元格是由行和列产生的。
2. 网格线的开始位置在容器的最顶端和最左边。
3. 使用区域命名之后会影响网格线的名称会变成 `区域名-star`、`区域名-end`
4. 可以把 `columns` 理解为高度，`rows`理解为宽度这样便于理解。

1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
