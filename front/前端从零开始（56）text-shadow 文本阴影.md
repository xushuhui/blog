# text-shadow 文本阴影

它可以给任意的字符设置一个或多个阴影。

## 1. 官方定义

text-shadow 属性向文本设置阴影。

## 2. 慕课解释

text-shadow 一共接受 4 个参数，前两个是阴影的位置通过 x，y 坐标系来设定，第三个参数设定模糊的大小，最后一个参数设定阴影的颜色。

## 3. 语法

```javascript
.demo{
    text-shadow: h-shadow v-shadow blur color;
}
```

属性值

|值|说明|
|--|----|
|h-shadow|可选。水平方向阴影位置，以文字的中心为起点》0 是往右，<0 时候偏左。|
|v-shadow|可选。竖直方向阴影位置，用法同上。                             |
|blur    |可选。模糊的大小。                                             |
|color   |可选。阴影的颜色                                               |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

1. 为文字添加阴影。

```javascript
<div class="demo">慕课网</div>
```

```javascript
.demo{
     text-shadow:5px 5px 5px red;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e35f09d6f69700780033.jpg)

为文字添加阴影效果图

1. 制作一个文字发光效果。

```javascript
html,body{
    background: #000;
}
.demo{
    color: #fff;
    text-shadow:5px 5px 20px #fff,-5px -5px 20px #fff,5px -5px 20px #fff,-5px 5px 20px #fff;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e36c09ec72ca00710052.jpg)

 制作一个文字发光效果效果图

说明：其实就是在各个方向上都增加一个白色的阴影，在黑色的背景下就显得有发光的效果了。

1. 通过投影直至化制作一个 3D 的文字效果。

```javascript
.demo{
    font-size: 30px;
    color: #fff;
    text-shadow:1px 1px hsl(0,0%,85%),
            2px 2px hsl(0,0%,80%),
            3px 3px hsl(0,0%,75%),
            4px 4px hsl(0,0%,70%),
            5px 5px hsl(0,0%,65%),
            5px 5px 10px black;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e3760985aa8901100051.jpg)

 制作一个 3D 的文字效果效果图

说明：这个效果也是利用各种色组叠加来实现的。

## 6. 经验分享

首先通过上面的例子我们可以了解到这个属性是可以无限制的增加一个颜色组，通过通过这个可以制作出很多有意思的效果。例如下雨的圆形光晕等等。

## 7. 小结

在以前也有`text-shadow:#000 5px 5px 5px`这样的写法，不过不推荐。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
