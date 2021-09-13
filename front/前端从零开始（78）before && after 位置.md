# before && after

这两个伪类元素功能很相似，都是在元素内部插入新的内容。下面一起看下他们的区别和用法。

## 1. 官方定义

before：元素的内容之前插入新内容。

after：元素的内容之后插入新内容。

## 2. 慕课解释

`before` 和 `after` 的功能就是在元素的内部的原有内容之前，或者之后插入新的内容。

## 3. 语法

```javascript
.demo:before{

}
.demo:after{

}
```

解释：使用方法如上面，通过在元素选择器后面增加一个 `:` 来开始伪类的使用。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

```javascript
<div class="demo">慕课网</div>
```

1. 在元素内容之前插入文字：姓名。

```javascript
 .demo:before{
    content: '姓名：';
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5eb160e109bb346f01070030.jpg)

元素内容之前插入文字：姓名 效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo:before{
            content: '姓名：';
        }
    </style>
</head>
<body>
    <div class="demo">慕课网</div>
</body>
</html>
```

1. 在元素内容之后插入：很好。

```javascript
 .demo:after{
    content: '很好';
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5eb1616909c05ee701140032.jpg)

在元素内容之后插入：很好 效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo:after{
            content: '很好';
        }
    </style>
</head>
<body>
    <div class="demo">慕课网</div>
</body>
</html>
```

## 6. 经验分享

这两个伪类当然不是仅仅插入内容这么简单，它还有其他的妙用。

1. 使用伪类 after 清除元素内部浮动效果：

```javascript
 <div class="demo">
    <div class="item">慕</div>
    <div class="item">课</div>
</div>
<div class="">网</div>
```

```javascript
.demo:after{
    content: '';
    display: block;
    clear: both;
}
.item{
    float: left;
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5eb161900978367c01000089.jpg)

 使用伪类 after 清除浮动 效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
       .demo:after{
            content: '';
            display: block;
            clear: both;
        }
        .item{
            float: left;
        }
    </style>
</head>
<body>
    <div class="demo">
        <div class="item">慕</div>
        <div class="item">课</div>
    </div>
    <div class="">网</div>
</body>
</html>
```

说明：下面灰色部分是没有清除浮动的效果，上面是清除浮动的效果。因为清除了浮动所以 “网” 这个字换行了。

1. 在元素内容开始前插入图片。

```javascript
<div class="demo">慕课网</div>
```

```javascript
.demo:before{
    content: '';
    display:inline-block;
    width:12px;
    height:12px;
    font-size:12px;
    line-height:12px;
    background: url(https://img.mukewang.com/wiki/5eea2f6809a8d35e00400040.jpg) center  no-repeat;
    background-size: cover;
}
```

![图片描述](https://img.mukewang.com/wiki/5eb161bb09a1685b01020032.jpg)

元素内容开始前插入图片 效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
       .demo:before{
            content: '';
            display:inline-block;
            width:12px;
            height:12px;
            font-size:12px;
            line-height:12px;
            background: url(https://img.mukewang.com/wiki/5eea2f6809a8d35e00400040.jpg) center  no-repeat;
            background-size: cover;
        }
    </style>
</head>
<body>
    <div class="demo">慕课网</div>
</body>
</html>
```

## 7. 小结

1. 注意：对于 IE8 及更早版本中的`:before`、`:after`，必须声明 <!DOCTYPE>。
2. 在元素选择器后面这样写也可以：

```javascript
.demo::before{

}
.demo::after{

}
```

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
