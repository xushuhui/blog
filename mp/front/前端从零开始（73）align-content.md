## align-content 多轴对齐

这个属性可以改变项目在容器中的对齐方式。

## 1. 官方定义

`align-content` 属性在弹性容器内的各项没有占用交叉轴上所有可用的空间时对齐容器内的各项（垂直）。

## 2. 慕课解释

`align-content` 是当容器内部的元素换行之后，我们如何设置他们所有在水平方向上排列的，这里要说的是它是一个多轴的统一设置。

## 3. 语法

```javascript
align-content: stretch|center|flex-start|flex-end|space-between|space-around|initial|inherit;
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

|IE|Edge|Firefox|Chrome|Safari|Opera|
|--|----|-------|------|------|-----|
|10+|12+|28+|4+|6.1+|12.1+|

## 5. 实例

想改变对齐方式只要给 `align-content` 使用不同的属性值，我们看下不同的值带来的效果。

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
            flex-wrap: wrap;
            height: 100px;
            border:1px solid #ccc;
        }
        .item{
            width: 400px;
            height: 30px;
            line-height: 30px;
            text-align: center;
            background: #ccc;
            border-right: 1px solid #fff;
        }
        .demo-2{
           align-content: flex-end;
           justify-content:flex-end;
        }
        .demo-3{
           align-content: center;
        }
        .demo-4{
           align-content: space-between;
        }
        .demo-5{
           align-content: space-around	;
        }
    </style>
</head>
<body>
    <p>
        flex-start: 默认值。项目位于容器的开头。
    </p>
    <div class="demo demo-1">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>
        flex-end 项目位于容器的结尾。
    </p>
    <div class="demo demo-2">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>
        flex-end 项目位于容器的中心。
    </p>
    <div class="demo demo-3">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>space-between 项目位于各行之间留有空白的容器内。</p>
    <div class="demo demo-4">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <p>space-around 项目在容器的前后留白并</p>
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

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb15157098a4dd108790926.jpg)

各种值设置的效果

## 6. 小结

1. 使用 justify-content 属性对齐主轴上的各项（水平），它和 `align-content` 并不冲突
2. 器内必须有多行的项目，该属性才能渲染出效果。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
