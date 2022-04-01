# flex-wrap 换行

`flex-wrap` 主要通过在外层容器中设置它里面的子项目是否可以换行。默认情况下项目是不换行的。

## 1. 官方定义

flex-wrap 属性规定 flex 容器是单行或者多行，同时横轴的方向决定了新行堆叠的方向。

## 2. 慕课解释

默认情况下，设置了 `display:flex` 的容器是不会换行的，这时候如果我们希望它换行就可以通过 `flex-wrap`设置超出宽度换行，也可以设置它如何换行，既换行之后的排列的方向。

## 3. 语法

```javascript
flex-wrap: nowrap|wrap|wrap-reverse|initial|inherit;
```

属性值

|值|描述|
|--|----|
|nowrap      |默认值。规定灵活的项目不拆行或不拆列。                  |
|wrap        |规定灵活的项目在必要的时候拆行或拆列。                  |
|wrap-reverse|规定灵活的项目在必要的时候拆行或拆列，但是以相反的顺序。|
|initial     |设置该属性为它的默认值。请参阅 initial。                |
|inherit     |从父元素继承该属性。请参阅 inherit。                    |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|10+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 设置一个容器，当内部的内容超过容器的宽度时候向下换行。

```javascript
.demo{
    display: flex;
    flex-wrap: wrap;
}
.item{
    width: 200px;
    height: 100px;
    line-height: 100px;
    background: #ccc;
    border-right: 1px solid #fff;
    text-align: center;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb13aa8099506c307380211.jpg)

换行效果图

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
        }
        .item{
            width: 200px;
            height: 100px;
            line-height: 100px;
            background: #ccc;
            border-right: 1px solid #fff;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="demo">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
</body>
</html>
```

1. 设置一个容器当内部的项目超过容器的宽度时候反向向下换行。

```javascript
.demo{
    display: flex;
    flex-wrap:  wrap-reverse;
}
.item{
    width: 200px;
    height: 100px;
    line-height: 100px;
    background: #ccc;
    border-right: 1px solid #fff;
    text-align: center;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb13ac809db0d0f07480213.jpg)

换行反向效果图

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
            flex-wrap: wrap-reverse;
        }
        .item{
            width: 200px;
            height: 100px;
            line-height: 100px;
            background: #ccc;
            border-right: 1px solid #fff;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="demo">
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
</body>
</html>
```

## 6. 小结

`flex` 弹性盒模型默认是不换行的既 `nowrap`

1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
