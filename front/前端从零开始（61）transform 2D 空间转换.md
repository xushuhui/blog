# transform 2D 平面转换

在以前我们改变元素的位置需要设置 `left` 、 `right` 这类的属性，它对其它元素有很大的影响，现在通过 `transform` 就可以实现任意空间的改变了。

## 1. 官方解释

CSS `transform` 属性允许你旋转，缩放，倾斜或平移给定元素。这是通过修改 CSS 视觉格式化模型的坐标空间来实现的。

## 2. 慕课解释

`transfrom` 这个属性可以改变一个目标元素在页面中的位置，例如相对原来元素所在的位置平移，相对原来的尺寸放大或者缩小，也可以旋转或者斜切。

## 3. 语法

通用坐标轴说明：

x 代表横轴，y 代表纵轴。

![图片描述](https://img.mukewang.com/wiki/5ea797b9090fb90202020188.jpg)

坐标轴效果图

包含参数：

|值|描述|
|--|----|
|translate(x,y)       |可以改变元素的位置，而不会对相邻元素由影响。|
|translateX(x)        |只改变元素的水平位置。                      |
|translateY(y)        |只改变元素在竖直方向的位置。                |
|scale(x,y)           |元素缩放，x 代表水平方向，y 代表竖直方向。  |
|scaleX(x)            |仅对元素 x 方向上缩放。                     |
|scaleY(y)            |仅对元素 y 方向上缩放。                     |
|skew(x-angle,y-angle)|定义沿着 X 和 Y 轴的 2D 倾斜转换。          |
|skewX(angle)         |定义沿着 X 轴的 2D 倾斜转换。               |
|skewY(angle)         |定义沿着 Y 轴的 2D 倾斜转换。               |
|rotate(angle)        |在平面上旋转一个角度                        |

## 4. 实例

接下来我们都是对 demo 这个元素进行操作。

```javascript
<div class="demo"></div>
```

1. 使用 `translate` 让元素位移。

```javascript
.demo{
     transform: translate(40px,40px);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea797df09c8f28101520150.jpg)

`translate` 让元素位移效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        body{
            perspective: 500px;
            transform-style: preserve-3d;
        }
        .commn{
            width:100px;
            height:100px;
            text-align: center;
            line-height: 100px;
            background:#f2f2f2;
            border:1px solid #ccc;
            position: absolute;
            top: 0;
            left: 0;

        }
        .demo{
            z-index: 1;
            opacity: .5;
            background: red;
        }
        .demo-3d{
            transform:translate3d(100px  ,100px ,-100px) rotateZ(45deg);

        }
    </style>
</head>
<body>
    <div class="commn demo">transfrom3d</div>
    <div class="commn demo-3d">transfrom3d</div>
</body>
</html>
```

1. 使用 `translateX` 让元素水平位移：

```javascript
.demo{
     transform: translateX(80px);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea797ec0918753d02020115.jpg)

`translateX` 让元素水平位移效果图

1. 使用 `translateY` 让元素在竖直方向上位移：

```javascript
.demo{
     transform: translateY(40px);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea797fc0933d4cd01250155.jpg)

`translateY` 让元素水平位移效果图

1. 使用 scale 对元素缩放：

```javascript
.demo{
     transform: scale(.8,.8);
}
```

scale 接受一个倍数大于 1 时候放大，小于 1 时候缩小。当 x，y 参数的值一样时，可以如下面这样写：

```javascript
.demo{
     transform: scale(.8);
}
```

如果只需要对水平方向缩放，可以向下面这样写，竖直方向同理：

```javascript
.demo{
     transform: scaleX(.8);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea79810093882eb01230119.jpg)

缩放效果图

1. 使用 `skew` 对元素倾斜。

```javascript
.demo{
    transform:skew(30deg,0deg);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea7981d09ee456101860122.jpg)

斜切效果图

```javascript
.demo{
    transform:skew(0deg,30deg);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea7982e09a952d201300169.jpg)

斜切效果图

如果只是对一个方向斜切可以如下：

```javascript
.demo{
    transform:skewX(30deg);
}
```

![图片描述](https://img.mukewang.com/wiki/5ea7981d09ee456101860122.jpg)

斜切效果图

1. `rotate` 使元素旋转一个角度。在 2D 效果中它只接受一个参数角度，并沿着顺指针方向开始。

```javascript
.demo{
    transform:rotate(30deg);
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea7984e09fe698b01610151.jpg)

旋转效果图

## 5. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 6. 场景

1. 在需要动画渲染改变元素位置时候。
2. 需要对场景进行缩放或者旋转时候。

## 7. 小结

1. 要分清 `transform` 和 `transition`，后者是过渡；
2. `transform` 可以使得元素位置改变，而不会影响其他围绕元素，所以可以使用 `transform` 尽量使用，可以提高浏览器的渲染效率；
3. `transform` 中斜切的效果，我们拿 X 水平坐标轴为例，其实就是底部向右移动一个角度，这个角度就是竖直方向偏移的角度。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
