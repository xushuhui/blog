# border-radius 圆角

如果想要把元素边界变得圆润，不妨试试这个属性。

## 1. 官方定义

通过 CSS3，您能够创建圆角边框并且不需使用设计软件，比如 PhotoShop。

## 2. 慕课解释

通过给一个 html 元素标签的样式增加一条 `border-radius` 属性，让这个元素的边角由直角边变成圆弧。

## 3. 语法

它的用法遵循 css 通用的：_左上、 右上 、右下 、左下_ 的原则。

```javascript
border-radius:value;
```

|属性|描述|
|----|----|
|border-radius             |四个边角值  |
|border-top-left-radius    |左上角圆弧值|
|border-top-right-radius   |右上角圆弧值|
|border-bottom-right-radius|右下角圆弧值|
|border-bottom-left-radius |左下角圆弧值|

包含参数

|参数名称|参数类型|
|--------|--------|
|value|‘%’ | ‘px’ | ‘rem’|

1. 只有一个参数时：

```javascript
border-radius:value;
```

`value` 代表给这个元素 4 个方向增加一个的圆弧值。

1. 只有两个参数时，中间用空格分开：

```javascript
border-radius:value1 value2;
```

`value1` 代表 左上、右下，value2 代表 右上、左下角圆弧值。

1. 只有三个参数时，中间用空格分开：

```javascript
border-radius:value1 value2 value3;
```

`value1` 代表左上 `value2` 代表 右上 左下 value3 代表右下

4. 有四个参数时，中间用空格分开

```javascript
border-radius:value1 value2 value3 value3;
```

`value1` ~ `value4` 分别代表左上、 右上 、右下 、左下四个角的圆弧值

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9|all|all|all|all|all|all|all|

## 5. 示例

1. 给 demo 增加右上和左下的圆角

```javascript
<div class="demo"></div>
```

可以这样

```javascript
.demo{
    border-radius:0  8px 0 8px;
}
```

推荐第一种写法，但是也可以这样

```javascript
.demo{
    border-top-right-radius:8px;
    border-bottom-left-radius:8px;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d4aa09c0f23b01110113.jpg)

demo 增加右上和左下的圆角效果图

1. 制作一个带有圆角的卡片

```javascript
<div class="card">
    <div class="text">
        demo1
    </div>
</div>
```

```javascript
.card{
    background: red;
    width: 100px;
    height: 200px;
    line-height: 200px;
    text-align: center;
    border-radius: 6px;
}
.text{
    display: inline-block;
    width: 50px;
    height: 50px;
    line-height: 50px;
    background: #fff;
    border-radius: 50%;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d4ca094e73a101100209.jpg)

带有圆角的卡片效果图

1. 给一个 `table` 增加圆角左上和右上是 8px 右下和左下是直角

```javascript
<table>
    <tr>
        <td>姓名</td><td>年龄</td>
    </tr>
    <tr>
        <td>demo</td><td>19</td>
    </tr>
</table>
```

```javascript
table{
    background: red;
    border-radius: 8px 8px 0 0;
    font-size: 18px;
    color: #fff;
    border-collapse: collapse;
    overflow: hidden;
}
```

![](./img/border-radius-table.png)

`table` 增加圆角左上和右上是 8px 右下和左下是直角效果图

## 6. 经验分享

1. `border-radius：50%` 会让一个宽度和高度相等的块级元素变成一个圆。

```javascript
.demo{
    width:100px;
    height:100px;
    border-radius：50%
}
```

设置 50% 的好处就是不用再去计算他的宽高，例如上面这个例子 `border-radius:50px`同样可以让这个元素变成一个圆。

2. 如果圆角过大，记得要设定 `padding` ，这样可以避免里面的内容超出元素。

## 7. 小结

1. 不要让 `border-radius` 的 % 值大于 50，因为如果两个相邻的半径和超过了对应的盒子的边的长度，那么浏览器要重新计算以保证它们不重合。虽然表面上看没有问题但是这样会对性能有影响。
2. 如果用 `rem`、`em` 单位在移动端用`border-radius` 画圆，在部分机型里面是椭圆的问题，可以通过 `50%` 来解决 ，或者使用 `px` 配合缩放 `scale` 来实现 `rem` 单位的效果。
3. 在内联元素`span`这类标签使用这个属性的时候同样有效但是记得不要使用`%`这样会导致一些`span` 标签的圆角不一样因为 `%` 是参考长和宽计算的。
4. `border-radius` 圆角并不会隐藏标签内部元素的内容，如果有内容溢出的情况记得增加`overflow:hidden;`
5. 任何元素都可以使用这个属性，包括视频、音频标签等等。
6. 一般情况下不推荐 `border-top-left-radius` 这类的写法除非是需要在某种交互过程中需要改变其中一个的圆角值而其它的保持不变。因为这类的标签会影响浏览器渲染的性能。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
