# only-child & only-of-type

在前端开发页面的过程中需要对一些特定类型的元素赋予特殊的样式，通常我们不会在 HTML 标签上一个个去增加 `class` 去设置特殊的样式，这时候通过元素选择伪类就能解决这类问题。本章主要介绍 `only-child` 和 `only-of-type` 这两个伪类。

## 1. 官方定义

`only-child` 匹配属于父元素中唯一子元素。

`only-of-type` 匹配属于父元素的特定类型的唯一子元素。

## 2. 慕课解释

`only-child` 当元素添加这个伪类的时候，它在所属的父元素之内，有且仅有它自己时伪类生效。

## 3. 语法

```javascript
.demo:only-child{

}
.demo:only-of-type{

}
```

说明：通过 `:` 后面加伪类进行元素选择。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

### only-child

**1. 当页面中只有一个。demo 标签时候背景变成红色：**

```javascript
<body>
    <div class="demo"></div>
</body>
```

```javascript
.demo:only-child{
    color:#fff;
    background: red;
    padding:10px;
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb1563709bdf1cf00800054.jpg)

一个标签时候背景变成红色效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
    .demo:only-child{
    color:#fff;
    background: red;
    padding:10px;
    }
</style>
</head>
<body>
    <body>
	    <div class="demo"></div>
	</body>
</body>
</html>
```

**2. 当页面有两个 demo class 时候不再有任何效果：**

```javascript
<body>
    <div class="demo"></div>
     <div class="demo"></div>
</body>
```

```javascript
.demo:only-child{
    color:#fff;
    background: red;
    padding:10px;
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb156c109377b6101340045.jpg)

无效果效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo:only-child{
            color:#fff;
            background: red;
            padding:10px;
        }
    </style>
</head>
<body>
        <div class="demo"></div>
         <div class="demo"></div>
</body>
</html>
```

说明：`body` 下面有两个 demo 不是唯一子元素，这时候伪类就不再起作用。

注意：当 demo 元素内部包含 demo 元素还是起作用的，因为 `body` 下面的子元素只有 1 个。

```javascript
<body>
	<div class="demo"> 
	    <div class="demo">   </div>
	    <div class="demo">   </div>
	</div>
</body>
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb1575209a60dc700870092.jpg)

一个元素效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
        .demo:only-child{
            color:#fff;
            background: red;
            padding:10px;
        }
    </style>
</head>
<body>
      <div class="demo"> 
          <div class="demo">   </div>
          <div class="demo">   </div>
      </div>
</body>
</html>
```

如果我们希望在 demo 内部只有**一个** demo 时候 ，内部的 demo 变成红色怎么做呢？

```javascript
<div class="demo"> 
    <div class="demo">   </div>
</div>
```

```javascript
.demo>.demo:only-child{
    color:#fff;
    background: red;
    padding:10px;
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb157b2094034d601370072.jpg)

 demo 内部只有 一个 demo 时候 内部的 demo 变成红色效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
     .demo>.demo:only-child{
        color:#fff;
        background: red;
        padding:10px;
    }
    </style>
</head>
<body>
    <body>
        <div class="demo"> 
            <div class="demo">   </div>
        </div>
    </body>
</body>
</html>
```

### only-of-type

**1. 给类名为 demo 的元素增加红色背景**

```javascript
<body>
    <div class="demo">  </div>
</body>
```

```javascript
 .demo:only-of-type{
    color:#fff;
    background: red;
    padding:10px;
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb1563709bdf1cf00800054.jpg)

 demo 变红效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
    .demo:only-of-type{
        color:#fff;
        background: red;
        padding:10px;
    }
    </style>
</head>

<body>
    <div class="demo">  </div>
</body>

</html>
```

说明：这里发现它和 `only-child` 的功能类似，但其实是不一样的我们看下面这个例子：

```javascript
<body>
    <div class="demo">  </div>
    <p class="demo">   </p>
</body>
```

```javascript
 .demo:only-of-type{
    color:#fff;
    background: red;
    padding:10px;
}
```

效果图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb158d00981f13702160104.jpg)

 变红效果图

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=`, initial-scale=1.0">
    <title>Document</title>
    <style>
    .demo:only-of-type{
        color:#fff;
        background: red;
        padding:10px;
    }
    </style>
</head>

<body>
    <div class="demo">  </div>
    <p class="demo">   </p>
</body>

</html>
```

解释：我们发现同样都变红了。这是因为 两个 demo 并不是唯一的。因为其中一个是 `div` 而另一个是 `p`，这时候 `only-child` 是不能分辨的，这也是它们的区别。

## 6. Tips

这两个伪类功能很类似，我们不容易区分但是这里有个小技巧 `:only-child` 就像 JS 中的 `id` 一样，只能是唯一的。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
