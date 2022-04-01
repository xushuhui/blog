# columns 字符分割

这个属性主要用来对字符进行横向分割排版，例如报纸的版面。

## 1. 官方定义

`columns`属性是一个简写属性，用于设置列宽和列数。

## 2. 慕课解释

`columns` 是 `column-width`每列宽度，`column-count` 每列的列数这两个属性的缩写，当列宽和列数的乘积大于元素的宽度时候就不会在分开自动合成一列。当他们的乘积小于元素的外宽的时候，每列的实际宽度可能大于`column-width` 设定的值。

## 3. 语法

使用 `columns` 时候

```javascript
.demo{
    columns: column-width column-count;
}
```

|值|描述|
|--|----|
|column-width|宽度 px|rem|em  |
|column-count|数字代表分的列数|

单独使用时候：

```javascript
.demo{
    column-count:number;
    column-width：value
}
```

另外 `colunms` 还有其他的补充属性：

```javascript
    column-gap:<length> | normal
```

设置列与列之间的距离：

```javascript
column-gap:<length> | normal
```

设置列与列之间的边线：

```javascript
column-rule：[ column-rule-width ]  [ column-rule-style ]  [ column-rule-color ]
```

内部元素是否允许横跨所有的列：

```javascript

column-span：none | all

```

|值|描述|
|--|----|
|none|不允许子元素单独一行|
|all |指定子元素单独一行  |

列的高度是否统一：

```javascript
column-fill	:auto | balance
```

|值|描述|
|--|----|
|auto   |自适应高度也就是它们有不同的高度。|
|balance|以最高的子元素为统一高度          |

这个属性兼容性极差，除了火狐支持外其它浏览器均不在支持了。

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9+|12+|28+|4+|6.1+|12.1+|7+|4.4|

## 5. 实例

1. 对一段文本分两列每列宽度不小于 200px。

```javascript
.demo{
    -webkit-columns:200px 2;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0cf1609dcb20412020324.jpg)

对一段文本分两列每列宽度不小于 200px 效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo{
    -webkit-columns:200px 2;
}
</style>

<div class="demo">

    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？

        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

1. 对一段文本分 3 列每列宽度不小于 200px。

```javascript
.demo{
    -webkit-columns:200px 3;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0cf830905d8c814340286.jpg)

对一段文本分 3 列每列宽度不小于 200px 效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo{
    -webkit-columns:200px 3;
}
</style>

<div class="demo">

    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？

        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

1. 对两段文本分 3 列。

```javascript
<div class="demo">
    <p>
        轻轻的我走了，正如我轻轻的来；
        我轻轻的招手，作别西天的云彩。 那河畔的金柳， 是夕阳中的新娘； 波光里的艳影， 在我的心头荡漾。软泥上的青荇，油油的在水底招摇； 在康河的柔波里，  我甘心做一条水草！  那榆荫下的一潭，不是清泉， 是天上虹； 揉碎在浮藻间，沉淀着彩虹似的梦。 寻梦？
    </p>
    <p>
        撑一支长篙， 向青草更青处漫溯；满载一船星辉， 在星辉斑斓里放歌。但我不能放歌，悄悄是别离的笙箫； 夏虫也为我沉默，沉默是今晚的康桥！悄悄的我走了，正如我悄悄的来； 我挥一挥衣袖，不带走一片云彩。
    </p>
</div>
```

```javascript
.demo
{
    -webkit-columns:200px 3;
}
p{
    margin: 0;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0d00a09ae6b6818000226.jpg)

对两段文本分 3 列效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo
{
    -webkit-columns:200px 3;
}
p{
    margin: 0;
}
</style>

<div class="demo">

    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？
    </p>
    <p>
        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

1. 修改两列间隔的距离。

```javascript
.demo{
    -webkit-columns:200px 2;
    column-gap:100px
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0d0320980d8e318040278.jpg)

修改两列间隔的距离效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo{
    -webkit-columns:200px 2;
    column-gap:100px
}
</style>

<div class="demo">

    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？
    </p>
    <p>
        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

1. 为每列直接增加边线。

```javascript
.demo{
    -webkit-columns:200px 2;
    column-gap:100px;
    column-rule:1px solid #ccc;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0d09809d2301021540238.jpg)

为每列直接增加边线效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo{
    -webkit-columns:200px 2;
    column-gap:100px;
    column-rule:1px solid #ccc;
}
</style>

<div class="demo">

    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？
    </p>
    <p>
        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

1. 让其内部 `class="head"`个子元素横跨所有列

```javascript
<div class="demo">
    <p class="head">再别康桥</p>
    <p>
        轻轻的我走了，正如我轻轻的来；
        我轻轻的招手，作别西天的云彩。 那河畔的金柳， 是夕阳中的新娘； 波光里的艳影， 在我的心头荡漾。软泥上的青荇，油油的在水底招摇； 在康河的柔波里，  我甘心做一条水草！  那榆荫下的一潭，不是清泉， 是天上虹； 揉碎在浮藻间，沉淀着彩虹似的梦。 寻梦？
    </p>
    <p>
        撑一支长篙， 向青草更青处漫溯；满载一船星辉， 在星辉斑斓里放歌。但我不能放歌，悄悄是别离的笙箫； 夏虫也为我沉默，沉默是今晚的康桥！悄悄的我走了，正如我悄悄的来； 我挥一挥衣袖，不带走一片云彩。
    </p>
</div>
```

```javascript
.demo{
    -webkit-columns:200px 2;
    column-gap:100px;
    column-rule:1px solid #ccc;

}
.head{
    column-span:all;
}
p{
    margin: 0;
    text
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb0d0c6093658d218300334.jpg)

让其内部 `class="head"`个子元素横跨所有列效果图

```javascript
<!DOCTYPE html>
<html>
<head>
<style>
.demo{
    -webkit-columns:200px 2;
    column-gap:100px;
    column-rule:1px solid #ccc;


}
.head{
    column-span:all;
    text-align: center;
}
p{
    margin: 0;
    background: #ccc;
    column-fill:balance;
}
</style>

<div class="demo">
    <p class="head">
        再别康桥
    </p>
    <p>
        轻轻的我走了，
    正如我轻轻的来；
    我轻轻的招手，
    作别西天的云彩。
    那河畔的金柳，
    是夕阳中的新娘；
    波光里的艳影，
    在我的心头荡漾。
    软泥上的青荇，
    油油的在水底招摇；
    在康河的柔波里，
    我甘心做一条水草！
    那榆荫下的一潭，
    不是清泉，
    是天上虹；
    揉碎在浮藻间，
    沉淀着彩虹似的梦。
    寻梦？
    </p>
    <p>
        撑一支长篙，
    向青草更青处漫溯；
    满载一船星辉，
    在星辉斑斓里放歌。
    但我不能放歌，
    悄悄是别离的笙箫；
    夏虫也为我沉默，
    沉默是今晚的康桥！
    悄悄的我走了，
    正如我悄悄的来；
    我挥一挥衣袖，
    不带走一片云彩。
    </p>

    </div>

</head>
<body>
```

## 6. 经验分享

使用 columns 可以快速的把元素内的字符分成我们想要的列数，如果想要自适应该怎么做呢？可以只设置列数这样在一定程度上可以不考虑元素的宽度，如下：

```javascript
.demo{
    -webkit-columns:2;
}
```

这样不管窗口怎么边它都是分成两列，其实任何自适应的原理也是如此。

## 7. 小结

1. 它分的列和子元素的个数无关。
2. 分的列数最好保证和内部子元素数量相等。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
