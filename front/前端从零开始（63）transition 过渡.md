# transition 过渡

如果想做出细腻的过渡效果，那么这个属性可能会满足你的需求。这个属性简单的来说就是用来模拟需要变化的属性，从开始到结束数值之间的过渡。

## 1. 官方定义

transition 属性是一个简写属性，用于设置四个过渡属性：

* transition-property
* transition-duration
* transition-timing-function
* transition-delay

## 2. 慕课解释

`transition` 用来设置一个属性状态从开始到结束中间这个过程的变化。它是 `transition-property`、`transition-duration`、`transition-timing-function`、`transition-delay`、这四个属性的缩写。它们分别代表了：要使用过度动画的属性、过渡动画的时间、过渡动画的加速度函数即数值变化的快慢过程、过渡动画的延迟时间。而我们通常使用过渡属性完成元素过渡的这个过程一般使用 `transition` 。

## 3. 语法

```javascript
.demo{
    transition: property duration timing-function delay;
}
```

属性值说明：

|属性值|描述|
|------|----|
| transition-property      | 规定设置过渡效果的 CSS 属性的名称。|
|transition-duration       | 规定完成过渡效果需要多少秒或毫秒。 |
|transition-timing-function| 规定速度效果的速度曲线。           |
|transition-delay          | 定义过渡效果何时开始。             |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 当鼠标移动到元素上，使用过渡属性来让元素的高度变化，从而实现一个过渡效果。

```javascript
<div class="demo"></div>
```

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #000;
    transition: height 1s;
}
.demo:hover{
    height: 150px;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea7a07c0a988dcc01350232.jpg)

 `hover` 之后高度变化效果图

1. 当鼠标移动上去改变元素的宽高值，让它们都实现过渡动画。

写法一：

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #000;
    transition: height 1s,width 1s;
}
.demo:hover{
    width: 150px;
    height: 150px;
}
```

写法二：

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #000;
    transition: all 1s;
}
.demo:hover{
    width: 150px;
    height: 150px;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea7a08c0aede55402220213.jpg)

 `hover` 宽高变化效果图

说明：这两种方式都可以实现我们所要的过渡方式。不过这里慕课推荐使用第一种方式。

1. 改变上面过渡完成的速度。

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #000;
    transition: height 1s ease-in,width 1s ease-out;
}
.demo:hover{
    width: 150px;
    height: 150px;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea7a0a30a025d9402190221.jpg)

 改变过渡完成的速度效果图

说明：在 `transition` 第三个值使用了动画函数，改变了过渡过程中完成的速度，我们可以很清楚的看到他们的变化速度。

1. 当鼠标移动上去 1s 之后开始动画。

```javascript
.demo{
    width: 100px;
    height: 100px;
    background: #000;
    transition: height 1s ease-in 1s,width 1s ease-out 1s;
}
.demo:hover{
    width: 150px;
    height: 150px;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea7a0af0a034a9102100204.jpg)

 时间设置效果图

说明：我们可以看到鼠标放到元素上 1s 之后开始动画，而离开元素之后 1s 之后开始动画。

## 6. Tips

通过上面的实例可以知道 `transition` 的属性值配置很灵活，但是我们要遵循一定的规律，这不单是增加了代码的可读性，也符合浏览器解析规则的规律。

`hover` 到按钮上改变按钮的位置和背景颜色。

```javascript
<button class="demo">慕课</button>
```

```javascript
.demo{
    width: 100px;
    height: 30px;
    line-height: 30px;
    border-radius: 4px;
    background: #000;
    color:#fff;
    border:none;
    transition: background .4s,transform .4s;
}
.demo:hover{
   background: red;
   transform: translateY(-5px);
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea7a0bc0a4be8bb01590104.jpg)

 改变按钮的位置和背景颜色效果图

## 7. 小结

1. 尽量不要使用 `all` 来驱动过渡的属性，这会使得浏览器卡顿。
2. 尽量不要使用 `margin-left` 这类的属性，很可能会打乱页面元素的位置。
3. 推荐使用 `transform` 来改变元素的位置和大小。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
