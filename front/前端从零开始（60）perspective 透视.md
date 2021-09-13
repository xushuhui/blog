# perspective 透视

**透视距离**和**透视位置**可以更好地观察拥有 3D 效果的元素。

## 1. 官方解释

perspective 属性定义 3D 元素距视图的距离，以像素计算。该属性允许您改变 3D 元素查看 3D 元素的视图。

当为元素定义 perspective 属性时，其子元素会获得透视效果，而不是元素本身。

perspective-origin 属性定义 3D 元素所基于的 X 轴和 Y 轴。该属性允许您改变 3D 元素的底部位置。

当为元素定义 perspective-origin 属性时，其子元素会获得透视效果，而不是元素本身。

## 2. 慕课解释

通过在父级元素设置这两个属性，可以简单的理解为设置一个观察者的位置，也就是我们的眼睛 perspective 的大小代表眼睛距离元素的位置。

perspective-origin，代表眼睛所在的坐标点，我们可以设置 x 轴和 y 轴，这两个属性其实就间接的组成了 （x,y,z）空间坐标组，要注意的是，这是设置都是在父元素上进行的。

## 3. 语法

```javascript
div
{
	perspective: 500px;
	perspective-origin:50% ,50%;
}
```

## 4. 兼容性

目前浏览器都不支持 perspective 属性。

Chrome 和 Safari 支持替代的 -webkit-perspective 属性。

## 5. 实例

1. 增加一个 500px 的透视效果

```javascript
<div class="demo">
	<div class="cell"></div>
</div>
```

```javascript
.demo{
	perspective: 500px;
	background: #f2f2f2;
}
.cell{
	width: 100px;
	height: 100px;
	background: #000;
	transform: translate3d(1px,-1px,-200px) rotateY(70deg);
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea79586099b089b01000100.jpg)

无透视

![图片描述](https://img.mukewang.com/wiki/5ea795bc09d02f3001000101.jpg)

有透视效果图

解释：加了 500px 的透视效果。

1. 修改观察点的位置为 50% 50% 。

```javascript
.demo{
	perspective: 500px;
	background: #f2f2f2;
	perspective-origin:50% 50%;
}
.cell{
	width: 100px;
	height: 100px;
	background: #000;
	transform: translate3d(1px,-1px,-200px) rotateY(70deg);
}
```

效果图：

![图片描述](https://img.mukewang.com/wiki/5ea795ec089e096006440192.jpg)

设置透视的 x 轴和 y 轴。

## 6. 经验分享

`perspective-origin` 通常使用 `%` 代表在观察父元素，观察点的坐标。

## 7. 小结

1. 推荐设置 `none` 而不是 0 ，内部的子元素不会透视。
2. 该属性的作用范围是针对子元素让其具有透视效果。
3. 不可以使用 % 数作为透视距离。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
