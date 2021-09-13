## align-items 竖直方向对齐方式

`align-items` 属性可以改变项目在容器中的对齐方式。

## 1. 官方定义

`align-items` 属性定义`flex` 子项在 `flex` 容器的当前行的侧轴（纵轴）方向上的对齐方式。

## 2. 慕课解释

`align-items` 主要用来设置一行内，当项目大小不一致时候的对齐方式。

**提示：** 子项目含有一个 `align-self` 属性可重写父级容器 `align-items` 属性，可以对单个项目对齐方式进行单一控制。

## 3. 语法

```javascript
align-items: stretch|center|flex-start|flex-end|baseline|initial|inherit;
```

属性值

|值|描述|
|--|----|
|stretch   |默认值。元素被拉伸以适应容器。          |
|center    |元素位于容器的中心。                    |
|flex-start|元素位于容器的开头。                    |
|flex-end  |元素位于容器的结尾。                    |
|baseline  |元素位于容器的基线上。                  |
|initial   |设置该属性为它的默认值。请参阅 initial。|
|inherit   |从父元素继承该属性。请参阅 inherit。    |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|10+|12+|28+|4+|6.1+|12.1+|7+||

## 5. 实例

想要改变对齐方式可以用过 `align-items`设置不同的属性值，我们看下不同的值带来的效果。

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo{
            display: flex;
        }
        .item{
            width: 100px;
            height: 100px;
            line-height: 100px;
            text-align: center;
            background: #ccc;
            border-right: 1px solid #fff;
        }
        .item:first-child{
            height: 120px;
        }
        .item:nth-of-type(3){
            height: 160px;
        }
        .demo-2{
            align-items: center;
        }
        .demo-3{
            align-items:flex-start;
        }
        .demo-4{
            align-items:flex-end;
        }
        .demo-5{
            align-items:baseline;
        }
    </style>
</head>
<body>
    <p>
        stretch 默认值。元素被拉伸以适应容器。
    </p>
    <div class="demo demo-1">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>
        center 项目位于容器的中心。
    </p>
    <div class="demo demo-2">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>
        flex-start 项目位于容器的头部。
    </p>
    <div class="demo demo-3">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p> flex-end 项目位于容器的低部。</p>
    <div class="demo demo-4">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>baseline 元素位于容器的基线上。默认情况和 flex-star 一样。</p>
    <div class="demo demo-5">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
</body>
</html>
```

效果图

![图片描述](https://img.mukewang.com/wiki/5eb14ad109f5342704441080.jpg)

 各种对齐方式的效果图

## 6. 小结

在可以使用 flex 的开发环境中，我们可以使用这种方式去对齐文字和图片，文字和 `input`这样的对齐方式简单快捷，远胜于其他的方式。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
