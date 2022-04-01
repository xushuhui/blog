# box-shadow 阴影

使用这个属性可以让页面更有立体感，给元素添加一个阴影，使得元素看起来是悬浮在原来的位置，下面就看看它的用法吧。

## 1. 官方定义

`box-shadow` 属性向框添加一个或多个阴影。

## 2. 慕课解释

通过 `box-shadow` 可以给任意 H5 元素添加阴影，可以是一个，如果用`,`号隔开可以添加多个。

## 3. 语法

```javascript
    box-shadow:h-shadow v-shadow blur color;
```

```javascript
.demo{
    box-shadow:1px 1px 5px #ccc;
}
```

属性值

|值|说明|
|--|----|
|h-shadow|可选。水平方向阴影位置，以文字的中心为起点 >0 是往右，<0 时候偏左。|
|v-shadow|可选。竖直方向阴影位置，用法同上。                                |
|blur    |可选。模糊的大小。                                                |
|color   |可选。阴影的颜色                                                  |

## 5. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

IE9+、Firefox 4、Chrome、Opera 以及 Safari 5.1.1 支持 box-shadow 属性。

## 6. 实例

```javascript
<div class="demo">

</div>
```

1. 给 demo 添加一个阴影。

```javascript
.demo{
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    box-shadow: 1px 1px 5px #ccc;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d65609580af901260120.jpg)

添加一个阴影效果图

1. 给 demo 添加多个阴影。

```javascript
.demo{
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    box-shadow: 1px 1px 5px #ccc,2px 2px 5px rgba(255, 25, 25,.5),3px 3px 1px rgba(5, 206, 89, 0.5) ;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d662094edae501310124.jpg)

添加多个阴影效果图

## 7. 经验分享

它的一把用来给元素添加一个阴影，交互设计师们热衷于在鼠标覆盖到元素时候，给元素增加一个悬浮效果，使它变得不同，例如：

```javascript
.demo{
    width:100px;
    height:100px;
    transition: box-shadow 1s;
}
.demo:hover{
     box-shadow: 1px 1px 5px #ccc；
}
```

而网上常见的 css3 下雨效果，就是利用 box-shadow 颜色叠加的特性制作出来的。

```javascript
<div class="demo"></div>
```

```javascript
.demo{
    width: 15px;
    height: 15px;
    border-radius: 50%;
    box-shadow: 100px 100px 4px #dedede, 30px 40px 4px #dedede,70px 20px 4px #dedede,80px 60px 4px #dedede;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d67b09c9d6bc01290137.jpg)

下雨效果效果图

如果画得密集些就更像了，当然这些就需要专业的设计师去做一个方案了。

最后这里介绍一下阴影效果使用的窍门。

```javascript
.demo{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;;
            box-shadow:  15px 0 15px -15px #000, -15px 0 15px -15px #000;
        }
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3d68709e38a8201260126.jpg)

给元素增加悬浮效果效果图

看完这个例子，我们应该已经明白 `box-shadow` 形成的阴影效果可能是一个组合。

## 8. 小结

1. 如果给一个元素添加多个阴影，那么后面的颜色层级高于前面的，也就是说如果前 3 个参数一样，后面的颜色会覆盖前面的颜色。
2. 不要给不规则的图形添加阴影，因为这只会添加到块级元素中，不会验证图片的路径添加阴影，可以使用滤镜来达到这样的效果。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
