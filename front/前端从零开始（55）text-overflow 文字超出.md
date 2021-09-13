# text-overflow 文字超出

该属性主要用来定义当文字超出元素限定范围内之后，为了防止页面混乱，对超出元素范围的文字做的一个处理。

## 1. 官网定义

text-overflow 属性规定当文本溢出包含元素时发生的事情。

## 2. 慕课解释

text-overflow 用于设置当文字内容超过所在元素设定的范围时候的展示效果。

## 3. 语法

```javascript
.demo{
    text-overflow: clip|ellipsis|string;
}
```

属性值说明

|值|说明|
|--|----|
|clip    |超出内容后裁剪                      |
|ellipsis|文字溢出后使用在最后的结尾使用省略号|
|string  |使用给定的字符串来代表被修剪的文本  |

## 4. 兼容

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|all|all|all|all|all|all|all|all|

## 5. 实例

1. 文字超出后裁剪内容

```javascript
<div class="demo">
    慕课网css3属性教学 text-overflow。
</div>
```

```javascript
.demo{
    height: 30px;
    width: 100px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow:clip ;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea3e2ab09cae9a801350036.jpg)

文字超出后裁剪内容效果图

1. 文字超出元素后使用省略号。

```javascript
<div class="demo">
    慕课网css3属性教学 text-overflow。
</div>
```

```javascript
.demo{
    height: 30px;
    width: 100px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow:ellipsis ;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea3e2b70922f6bd01200031.jpg)

文字超出元素后使用省略号效果图

1. 文字超出后使用 `！`。

```javascript
<div class="demo">
    慕课网css3属性教学 text-overflow。
</div>
```

```javascript
.demo{
    height: 30px;
    width: 100px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow:"!" ;
}
```

效果图

![图片描述](https://img.mukewang.com/wiki/5ea3e2c30923179901190032.jpg)

文字超出后使用 `！`效果图

说明：只在火狐浏览器兼容。

## 6. 经验分享

`text-overflow`这个属性常常用来限制文字超出后我们怎么去处理超出的这部分文字，通常是隐藏掉，这样才不会破坏页面的视觉效果，不会把文字堆积重叠到一起。

## 7. 小结

1. `text-overflow`一定要和`overflow: hidden;`、`white-space: nowrap;` 一起使用，不能单独用。
2. 这个属性通常是在有设置宽度和高度的元素使用。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
