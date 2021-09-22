## flex order 排序

一般情况下浏览器会把元素从左到右或者从上到下排列，如果我们想要更改它们的排列顺序该如何做呢？使用`order`就可以轻松的修改。数字越大越往后，数字越小越在前。

## 1. 官方定义

order 属性设置或检索弹性盒模型对象的子元素出现的順序。

## 2. 慕课解释

子元素可以通过设置 order 数值的大小来设定在页面中出现的顺序，数值小的在前，数值大的在后。

## 3. 语法

```javascript
.item-child{
    order:1;
}
```

属性说明

|参数名称|参数类型|解释|
|--------|--------|----|
|order|number|数值越小排位越靠前|

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|10+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 子元素 child-1 在右侧 child-2 在左侧。

```javascript
<div class="demo">
    <div class="child-1">
        1
    </div>
    <div class="child-2">
        2
    </div>
</div>
```

```javascript
.demo{
    display: flex;
}
.child-1{
    flex:auto;
    order:2;
    background: #000;
}
.child-2{
    flex:auto;
    order:1;
    background: rgb(255, 2, 2);
}

```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0ea02095d2d2a02120113.jpg)

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style>
       .demo{
    display: flex;
}
.child-1{
    flex:auto;
    order:2;
    background: #000;
}
.child-2{
    flex:auto;
    order:1;
    background: rgb(255, 2, 2);
}
    </style>
</head>
<body>
<div class="demo">
    <div class="child-1">
        1
    </div>
    <div class="child-2">
        2
    </div>
</div>

</body>
</html>
```

子元素 child-1 在右侧 child-2 在左侧效果图

1. 子元素 child-1 在下 child-2 在上。

```javascript
<div class="demo">
    <div class="child-1">
        1
    </div>
    <div class="child-2">
        2
    </div>
</div>
```

```javascript
.demo{
    display: flex;
    flex-direction: column ;
}
.child-1{
    flex:auto;
    order:2;
    background: #000;
}
.child-2{
    flex:auto;
    order:1;
    background: rgb(255, 2, 2);
}

```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0eab709b707c801150211.jpg)

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style>
       .demo{
    display: flex;
    flex-direction: column ;
}
.child-1{
    flex:auto;
    order:2;
    background: #000;
}
.child-2{
    flex:auto;
    order:1;
    background: rgb(255, 2, 2);
}
    </style>
</head>
<body>
<div class="demo">
    <div class="child-1">
        1
    </div>
    <div class="child-2">
        2
    </div>
</div>
</body>
</html>
```

 子元素 child-1 在下 child-2 在上

## 6. 经验分享

通过使用 `order` 属性可以实现拖动排序，当 JS 脚本运行之后，只要确定元素拖动到指定的位置通过修改对应的 `order` 就可以轻松完成顺序的改变。

## 7. 小结

只有在弹性盒模型下起作用。

1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
