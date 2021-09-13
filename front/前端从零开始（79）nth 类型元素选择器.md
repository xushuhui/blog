# nth 元素选择

当我们要一组 `class` 同名，或者连续的一组元素的其中一个，或者某种规律的元素添加单独样式的时候，不妨看看这类的元素选择器。

## 1. 官方定义

* `nth-child(n)` 选择器匹配属于其父元素的第 N 个子元素；
* `nth-last-child(n)` 选择器匹配属于其元素的第 N 个子元素的每个元素，从最后一个子元素开始计数；
* `nth-of-type(n)` 选择器匹配属于父元素的**特定类型**的第 N 个子元素的每个元素。

## 2. 慕课解释

`nth-child(n)`、 `nth-last-child(n)` 、`nth-of-type(n)` 都是用来匹配父元素内部子元素的。不过也有些区别：

`nth-child` 按照个数来算；

`nth-of-type` 按照类型来计算；

`nth-last-child(n)` 从最后一个子元素往前开始计算。

## 3. 语法

```javascript
.item:nth-child(2n+1){

}
.item:nth-of-type(n){

}
.item:nth-last-child(2n){

}

n 从 0 开始计数的正整数。
```

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

选择 demo 内第 3 个子元素背景为红色。

1. 使用 `nth-child`。

```javascript
.item{
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    border: 1px solid #ccc;
    background: #f2f2f2;
}
.item:nth-child(3){
    background: red;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5eb17681094b2e2701110421.jpg)

第三个背景变红效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .item{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-child(3){
            background: red;
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

1. 使用 `nth-last-child`。

```javascript
.item{
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    border: 1px solid #ccc;
    background: #f2f2f2;
}
.item:nth-last-child(2){
    background: red;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5eb17681094b2e2701110421.jpg)

第三个背景变红效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .item{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-last-child(2){
            background: red;
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

1. 使用`nth-of-type`。

```javascript
.item{
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    border: 1px solid #ccc;
    background: #f2f2f2;
}
.item:nth-of-type(3){
    background: red;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5eb17681094b2e2701110421.jpg)

第三个背景变红效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .item{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-of-type(3){
            background: red;
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

## 6. 经验分享

1. 在实例中我们看到 `nth-of-type` 和 `nth-child` 同样都使用的是 （3）， 那么它们的不同是什么呢？下面这个例子我们一起看下：

```javascript
<div class="demo">
    <p class="item">我是 p 标签</p>
    <div class="item">1</div>
    <div class="item">2</div>
    <div class="item">3</div>
    <div class="item">4</div>
</div>
<div class="demo">
    <p class="item-2">我是 p 标签</p>
    <div class="item-2">1</div>
    <div class="item-2">2</div>
    <div class="item-2">3</div>
    <div class="item-2">4</div>
</div>
```

```javascript
   .demo{
           float: left;
       }
       .item,.item-2{
           width: 100px;
           height: 100px;
           text-align: center;
           line-height: 100px;
           border: 1px solid #ccc;
           background: #f2f2f2;
       }
       .item:nth-of-type(3){
           background: red;
       }
       .item-2:nth-child(3){
           background: red;
       }
```

效果图

![图片描述](https://img.mukewang.com/wiki/5eb176a809a3eb7402170545.jpg)

`nth-of-type` 和 `nth-child` 效果图

通过效果图我们就清楚的明白他们的差异了。

简述实例展现效果，通过实例分析他们两个的区别

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo{
            float: left;
        }
        .item,.item-2{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-of-type(3){
            background: red;
        }
        .item-2:nth-child(3){
            background: red;
        }
    </style>
</head>
<body>
    <div class="demo">
        <p class="item">我是 p 标签</p>
        <div class="item">1</div>
        <div class="item">2</div>
        <div class="item">3</div>
        <div class="item">4</div>
    </div>
    <div class="demo">
        <p class="item-2">我是 p 标签</p>
        <div class="item-2">1</div>
        <div class="item-2">2</div>
        <div class="item-2">3</div>
        <div class="item-2">4</div>
    </div>
</body>
</html>
```

下面是让所有偶数的背景变红。

```javascript
 .item{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-of-type(2n){
            background: red;
        }
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5eb176ee0958ef1001150415.jpg)

偶数的背景变红 效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .item{
            width: 100px;
            height: 100px;
            text-align: center;
            line-height: 100px;
            border: 1px solid #ccc;
            background: #f2f2f2;
        }
        .item:nth-of-type(2n){
            background: red;
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

1. 使用 `nth-of-type(3n+1)` 起作用，而 `nth-of-type(1+3n)` 不起作用，所以 `n` 一定要放在最前面。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
