# border-image 边框图片

这个属性的兼容性不是很好，所以在正常的前端开发工作中很难用到，但是在不考虑兼容性的情况下，它无疑的一个很强大的属性，因为它可以自定义漂亮的边框，而不在是单调的线条。

## 1. 官方解释

CSS 属性允许在元素的边框上绘制图像。这使得绘制复杂的外观组件更加简单，也不用在某些情况下使用九宫格了。使用 `border-image` 时，其将会替换掉 `border-style` 属性所设置的边框样式。虽然规范要求使用 `border-image` 时边框样式必须存在，但一些浏览器可能没有实现这一点。

## 2. 慕课解释

通过 `border-image` 属性可以给元素添加自定义得而边框样式，而不单单是系统提供的那几种。换句话说就是我们可以自定义一个图片来充当元素的边框对它进行环绕。

## 3. 语法

```javascript
border-image:source slice repeat;
```

包含属性

|属性|描述|
|----|----|
|border-image-source|背景图片源                                                                |
|border-image-slice |需要展示出来图片的尺寸，如果这个量等于图片的尺寸就都相当于整个图片展示出来|
|border-image-width |图片边框的宽度。                                                          |
|border-image-outset|边框图像区域超出边框的量。                                                |
|border-image-repeat|图片的填充形式                                                            |

## 4. 实例

1. 使用 `border-image` 为元素自定义一个图片边框。

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #ccc;
    border-width: 50px;
    border-style: solid;
    border-image: url(./../img/border-image.jpg);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3ab8509528c6202140215.jpg)

 使用 `border-image` 为元素自定义一个图片边框效果图

我们这使用的是 `bordr-image` 这个属性，并定义了图片路径 其它的不设定使用默认值。

1. 使用 `border-image-source` 为元素设定一个边框。

```javascript
.demo2{
    width: 100px;
    height: 100px;
    background: #ccc;
    border-width: 50px;
    border-style: solid;
    border-image-source: url(./../img/border-image.jpg);
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3ab8509528c6202140215.jpg)

 使用 `border-image-source` 为元素设定一个边框效果图

我们可以看到 例 1 和 例 2 两个图是一样的，因为我们仅仅使用了 `border-image-source` 增加了图片路径而已。

1. 我们在 demo2 上增加 `border-image-slice` 。

```javascript
.demo2{
    width: 100px;
    height: 100px;
    background: #ccc;
    border-width: 50px;
    border-style: solid;
    border-image-source: url(./../img/border-image.jpg);
    border-image-slice: 70;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3abbc090f814601560156.jpg)

 增加 `border-image-slice`效果图

通过给 `slice` 设定一个 70 我们得到了一个漂亮的边框，下面重点说下这个属性值是如何作用在边框图片上的。

4. 继续在 demo2 的基础上增加 `border-image-outset` 。

```javascript
.demo2{
    width: 100px;
    height: 100px;
    background: #ccc;
    border-width: 20px;
    border-style: solid;
    border-image-source: url(./../img/border-image.jpg);
    border-image-slice:20;
    border-image-outset:22px;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3abce0920b25102180200.jpg)

 增加 `border-image-outset`效果图

从效果图中的红色箭头我们可以看到，边框图片和灰色元素之间有一条 2px 的白线，这是因为我设置了 `image-outset` 向外偏移 了 22px 的原因。

1. 使用 `border-image-repeat` 来为 demo2 设定图片的填充形式

```javascript
.demo2{
    width: 100px;
    height: 100px;
    background: #ccc;
    border-width: 20px;
    border-style: solid;
    border-image-source: url(./../img/border-image.jpg);
    border-image-slice:20;
    border-image-outset:22px;
    border-image-repeat: repeat;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3abdf09e69bbd01990197.jpg)

 使用 `border-image-repeat` 来为 demo2 设定图片的填充形式效果图通过效果图我们可以看到图片是以平铺重复的方式来填充的，而这个属性默认是 stretch 拉伸来填充图片的。这个属性还有以下值

|属性值|描述|
|------|----|
|stretch|默认值，拉伸图片来填充区域。                                                   |
|repeat |平铺并重复图像来填充区域。                                                     |
|round  |类似 repeat 值。如果无法完整平铺所有图像，则对图像进行缩放以适应区域。         |
|space  |不拉伸图片，而是让图片成四周环绕即左上右上右下左下。                           |
|initial|关键字用于设置 CSS 属性为它的默认值 。可以用于任何 HTML 元素上的任何 CSS 属性。|
|inherit|继承父级的设定                                                                 |

## 5. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|11|12+|50+|5+|9.1+|43+|9.3+|2.1+|

## 6. 经验分享

1. 这个属性使用并不太难，不过其中的`border-image-slice`属性在使用的时候有一个小技巧，就是当想要其中 `image` 不失真即不拉伸，`slice` 要和 `width` 一样。

例如：

```javascript
.demo{
    border-width:20px;
    border-image-slice:20;
}
```

还有一点要注意的是 `slice` 不需要挂单位。

2. 如果想填充中心可以加上`-webkit-border-image`就像下面这样

```javascript
.demo{
    -webkit-border-image: url(./../img/border-image.jpg) 20 20  stretch;
}
```

这是 `border-imgae` 的连写方式，第一个和第二个 20 都是 `slice`，认出它的最后方法就是 它们都不带单位。

## 7. 小结

1. `border-image` 使用时候一定要设定 `border-style` 虽然这个属性没什么用，但是如果不设定它 `border-image` 就不生效。
2. `border-image-slice` 如果设定数值当 `px` 用时候直接写数字就可以 如果加上 `px` 反而会不生效。
3. `border-width` 用来设定边框的宽度，它决定图片边框展示的厚度（即围绕宽度）。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
