# border 边框

有时候在页面中需要做一些分割来区分不同的区域，这个属性不但可以用来给元素添加一个边框，也可以作为不同区域的分割线。

## 1. 官方解释

CSS 的 `border` 属性是一个用于设置各种单独的边界属性的简写属性。 `border` 可以用于设置一个或多个以下属性的值： `border-width`、`border-style`、`border-color`。

## 2. 慕课解释

任何一个 h5 标签通过添加一个 `border` 属性，可以给它设置边框的宽度、展示出来的样子（实线、虚线、圆点等）和颜色。

## 3. 语法

```javascript
border: [border-width ||border-style ||border-color |inherit] ;
```

属性值：

|参数名称|参数类型|解释|
|--------|--------|----|
|border-width|[‘px’|‘rem’|’%’]|控制边框的宽度                          |
|border-style|`string`              |控制边框的样式                          |
|border-color|`string`              |控制边框的颜色                          |
|inherit     |[’’|’’]           |控制边框展示在元素宽高尺寸的外部还是内部|

`border-width`、`border-style`、`border-color`它们的用法遵循 css 的：_左上、 右上 、右下 、左下_ 的原则。最多可以添加 4 个参数。

相关属性：

|参数名称|参数类型|
|--------|--------|
|border-top   |[border-width |border-style |border-color |inherit]|
|border-bottom|[border-width |border-style |border-color |inherit]|
|border-right |[border-width |border-style |border-color |inherit]|
|border-left  |[border-width |border-style |border-color |inherit]|

`boder`可以直接设置元素的宽度、边框样式、颜色，不需要再去单独通过`border-width`、`border-style`、`border-color`去设置。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

接下来我们通过一个 div 元素，来说明 border 的使用方法。

```javascript
 <div class="demo"></div>
```

1. 为 demo 增加边框

```javascript
.demo{
        width:100px;
        height: 100px;
        border:1px solid #ddd;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea2b8d009e3bb3e01110112.jpg)

为 demo 增加边框效果图

解释：通过`border` 为`div`四周增加一个宽度为 1px、填充样式为`solid`（实线）、颜色为`#ddd`的边框。

1. 只给 demo 的顶部增加一个边框：

```javascript
.demo{
    width:100px;
    height: 100px;
    border-top:1px solid #000;
    background: #f2f2f2;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea2b8ed098e23b201110113.jpg)

只给 demo 的顶部增加一个边框效果图

解释：通过`border-top`可以只给 div 的顶部增加边框而不会影响其它部分的样式。

1. 修改 button 默认的边框样式

```javascript
<button class="btn"></button>
```

```javascript
.btn{
        border-width: 4px;
    }
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea2b8ff0971ee5901090035.jpg)

修改 button 默认的边框样式效果图

解释：左边是默认的 button 效果，右边图片是我们修改后的效果。

1. 个性化 demo 每个边框的颜色。

```javascript
.demo{
    width:100px;
    height: 100px;
    border:2px solid;
    border-top-color:red;
    border-right-color:green;
    border-bottom-color:black;
    border-left-color:orange;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea2b910090d82b201120113.jpg)

个性化 demo 每个边框的颜色效果图

这里千万不能写成 `border-right:green;`

## 6. 经验分享

我们有时候在元素内可能会使用`margin-top`这样的属性来让子元素和父元素的顶部有一个间隔，但是经常不天遂人愿，发现它并没有达到我们想要的效果，反而变成了父元素距离上一个元素增加了一个距离。如果我们在父元素设置一个看不见的`border`就可以解决这个问题。

## 7. 小结

1. border-width 属性会有影响设定元素的尺寸。
2. 在 table 中使用 border，要使用`border-xx`这样的属性，为的是去掉一边避免重叠。

```javascript
td{
    border:1px solid #ccc;
    border-bottom:none;
}
```

1. border-color 如果不设置那么它会使用元素中字体的颜色。

2.

```javascript
div{
widht:100px;
height:100px
border-width:2px;
}
```

上面这样的设置 div 在页面中实际站位是 104px，如果不注意很容易造成页面错乱。这是因为我们大部分情况下盒模型使用的是 W3C 标准的’box-sizing: content-box;’，它在页面中实际宽度 = `width`+`border`( 该公式仅针对上面例子）。

1. 如果我们需要给 button 设置一个颜色，那么它就会失去浏览器自带的交互效果。

2. border-top 的使用和 border 的使用方法是一样的，如果要个性化一个边的颜色，可以这样设置： `border-top-color:red;`也可以这样设置： `border-top:1px solid red;`。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
