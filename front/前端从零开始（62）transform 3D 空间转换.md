# transform 3D 空间转换

transform 这个属性的强大之处在于它可以把一个二维的空间转化成一个三维的空间，给视觉设计师更多的发挥空间，也给用户带来更好的视觉体验。

## 1. 官方定义

transform 属性向元素应用 3D 转换。属性允许我们对元素进行旋转、缩放、移动或倾斜。

## 2. 慕课解释

当给元素使用 transform 之后，它就可以在它原来所在的位置变成一个向任意空间变换的元素，这里可以通过在 Z 轴上的设置，让他在空间上呈现 3D 效果。

## 3. 语法

```javascript
transform: none|transform-functions;
```

3D 空间坐标轴

![坐标轴](https://xushuhui.gitee.io/image/imooc/5ea79e8008dcde4804200440.jpg)

坐标轴值说明

|值|描述|
|--|----|
|translate3d(x,y,z)   |定义 3D 转换。                                        |
|translateX(x)        |定义转换，只是用 X 轴的值。                           |
|translateY(y)        |定义转换，只是用 Y 轴的值。                           |
|translateZ(z)        |定义 3D 转换，只是用 Z 轴的值。                       |
|scale3d(x,y,z)       |定义 3D 缩放转换。                                    |
|scaleX(x)            |通过设置 X 轴的值来定义缩放转换。                     |
|scaleY(y)            |通过设置 Y 轴的值来定义缩放转换。                     |
|scaleZ(z)            |通过设置 Z 轴的值来定义 3D 缩放转换。                 |
|rotate3d(x,y,z,angle)|定义 3D 旋转。                                        |
|rotateX(angle)       |定义沿着 X 轴的 3D 旋转。                             |
|rotateY(angle)       |定义沿着 Y 轴的 3D 旋转。                             |
|rotateZ(angle)       |定义沿着 Z 轴的 3D 旋转。                             |
|transform-style      |在空间内如何呈现 `flat` 2D 呈现，`preserve-3d` 3D 呈现|

我们在 transform2D 中已经对平面属性做了详细的介绍，本章节主要是其 3D 属性，这些属性的使用需要在父级设置 `perspective` 和 `transform-style`

让父级有了透视效果以及设置父级在内部空间的呈现方式。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

通用 html ：

```javascript
<div class="common demo">transfrom3d</div>
<div class="common demo-3d">transfrom3d</div>
```

通用 style ：

```javascript
body{
    perspective: 500px;
}
.common{
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
```

1. demo-3d 在 z 坐标轴向内延伸 100px。

```javascript
.demo-3d{
    transform:translate3d(0 ,0 ,-100px);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea79e1c089a217a05280260.jpg)

demo 在 z 坐标轴向内延伸 100px ，效果图

**说明：** 红色背景是 demo-3d 原来的位置，我们通过图片看到它的表现是水平向右移动且缩小了，其实他是进行了 3D 空间的移动。

1. demo-3d 在 z 轴空间上缩放。

```javascript
.demo-3d{
    transform:scale3d(1 ,1 ,0);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea79e3b085cfe5402440236.jpg)

demo 在 z 轴空间上缩放效果图

**说明：** scale3d 这个属性可以拆成 scaleX() 、scaleY() 、 scaleZ() 。我们发现 scaleZ() 在 3D 空间变化上，它的区间 0~1 是不起作用的，只有 0 代表缩小到 0（消失）， 1 （不变）。

1. demo-3d 在 z 轴上旋转。

```javascript
.demo-3d{
    transform: rotate3d(0,0,1,45deg);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea79e4b080baf2b03020288.jpg)

demo-3d 在 z 轴上旋转效果图

**说明：** 不推荐使用 rotate3d() 这个属性，因为它只能通过 0 或 1 去选择是否需要旋转，第 4 个参数给 1 个旋转角度，这种方式很不灵活，不过它的特性就是可以同时控制 x，y，z 方向上的旋转角度。

其实我们从 1～3 这 3 个例子中看到只设定了其中一项，接下来我们全方位的变化。

1. 在 x，y，z 上应用 translate3d 和 rotateZ

```javascript
.demo-3d{
    transform:translate3d(100px  ,100px ,-100px) rotateZ(45deg);
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea79e590864104f06240404.jpg)

在 x，y，z 上应用 translate3d 和 rotateZ 效果图

5. 写一个墙角效果

```javascript
<div class="cude">
        <div class="common left">left</div>
        <div class="common right">right</div>
        <div class="common bottom">bottom</div>
</div>
```

```javascript
.cude {
    perspective: 1500px;
    width:200px;
    height:200px;
    position: relative;
    margin: 100px auto;
    transform-style: preserve-3d;
    transform: rotateX(-14deg) rotateY(-45deg);
}
.common {
    position: absolute;
    top: 0;
    left: 0;
    width: 200px;
    height: 200px;
    background:#666;
    opacity: 0.8;
    font-size:20px;
    text-align: center;
    line-height:200px;
    font-weight: bold;
    color:#fff;
    border:1px solid #fff;
}
.right {
    transform: rotateY(180deg) translateZ(101px);
    background: rosybrown;

}
.left {
    transform: rotateY(-90deg) translateZ(101px);
    background: rosybrown;
}
.bottom {
    transform: rotateX(90deg) translateZ(-100px);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea79e6e086fab5107340590.jpg)

墙角效果图

**说明：** 写这个其实没有什么技巧，首先设置 `transform-style: preserve-3d;` 然后在理解每个面相对角度的基础上去设置 `translateZ 和 rotate3d`.

## 6. 经验分享

1. 我们如果是初学 transform，这里介绍一个区分旋转角度方向的方法，也就是左手法则，我们左手指向设置旋转坐标的正向，例如 z 轴，左手只向屏幕外，手指自然弯曲的方向就是旋转的方向。

2. 我们通常旋转 都是以中心点为起点开始旋转的，这往往不是我们想要的方式，可以通过设置 `transform-origin` 这个属性改变旋转起始点的位置。

```javascript
transform-origin: 50% 50% 0;
```

上面设置代表在元素的水平面的中心位置。

```javascript
transform-origin: 0 50% 0;
```

上面这个设置代表在元素 `top` 的中心位置。

```javascript
transform-origin: 50% 0  0;
```

上面这个设置代表在元素 `left` 的中心位置。

1. 如果我们在实际工作中遇到改变元素的位置，例如拖拽这些使用 `transform:translate3D(x,y,z)` 可以提高浏览器的性能，而且它的位置变化不会改变页面中其它元素的位置。

## 7. 小结

`rotateX`、`rotateY`这些是 3D 空间的变化，不可以出席在 2D 空间上面。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
