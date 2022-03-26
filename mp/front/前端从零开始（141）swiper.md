# swiper

> Swiper 常用于移动端网站的内容触摸滑动
>
>
> Swiper 是纯 JavaScript 打造的滑动特效插件，面向手机、平板电脑等移动终端。

`swiper.js` 在国内使用面非常广，可以用于实现轮播、图片预览、可滚动应用等，发挥想象可是实现诸多需求。

本篇幅采用 `swiper5`，所有版本的 `api` 都很相似，主要区别可以参考[官方的提供的说明](https://www.swiper.com.cn/api/index.html)。

## 1. 使用

```javascript
<style>
  .container {height: 100px;overflow: hidden;}
  .slide-item {display: flex;justify-content: center;align-items: center;color: white;font-size: 42px;}
  .item1 {background-color: cornflowerblue;}
  .item2 {background-color: darkslateblue;}
  .item3 {background-color: darkorange;}
</style>

<div class="container">
  <div class="swiper-wrapper">
    <div class="swiper-slide slide-item item1">第一屏</div>
    <div class="swiper-slide slide-item item2">第二屏</div>
    <div class="swiper-slide slide-item item3">第三屏</div>
  </div>
</div>
<link href="https://cdn.bootcdn.net/ajax/libs/Swiper/5.4.5/css/swiper.min.css" rel="stylesheet">
<script src="https://cdn.bootcdn.net/ajax/libs/Swiper/5.4.5/js/swiper.min.js"></script>
<script>
  var mySwiper = new Swiper('.container', {
    autoplay: true, // 自动切页
  });
</script>
```

`swiper` 需要引入两个资源，通常和 `UI` 相关的框架都会有两个及以上资源，需要额外引入样式。

轮播是非常常见的需求，只需一些简单的 `DOM` 结构就可以完成。

`.swiper-wrapper` 和 `.swiper-slide` 两个类是和 `swiper` 联动的，swiper 在初始化的时候会在 `swiper-wrapper` 下 `swiper-slide` 作为每一项，这些类名都是可通过配置修改的。

## 2. 使用 swiper 实现移动端的图片预览

移动端产品的图片查看几乎都是全屏预览，可以作用滑动切换图，支持缩放手势等，swiper 天然支持这些功能，同时又可以深度定制，适合制作业务组建嵌入项目。

分析一下需求：点击图片查看大图，图片可以手势缩放，同时支持左右切换。

其实这就是一个不会自动切换的轮播，通过 swiper 就能实现。

可以设计一个方法，方法接收 当前图片和所有图片列表，然后每个图片为一页，生成一个轮播，显示在页面的最上层。

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
</head>
<body>
  <style>
    .container {position: fixed;top: 0;bottom: 0;right: 0;left: 0;background-color: rgba(0, 0, 0, .7);}
    .container .swiper-pagination {display: none;}
    .slide-item {overflow: hidden;}
    .slide-item img {width: 100%;}
    .count {position: absolute;left: 50%;transform: translateX(-50%);top: 16px;color: white;}
  </style>

  <link href="https://cdn.bootcdn.net/ajax/libs/Swiper/5.4.5/css/swiper.min.css" rel="stylesheet">
  <script src="https://cdn.bootcdn.net/ajax/libs/Swiper/5.4.5/js/swiper.min.js"></script>
  <script>
  function previewImage(current, list) {
    if (!list) list = [current]; // 如果没传，默认以初始图为列表

    if (list.length === 0) list = [current]; // 如果数组为空 则以初始图为列表

    var idx = 0; // 寻找初始图在列表的位置
    var html = list.map(function(item, index) {
      if (item === current) { // 如果两个图 url 相等，则说明初始图就是在这个位置
        idx = index + 1; // 记录下位置
      }

      // 拼一个 swiper-slide
      return [
        '<div class="swiper-slide slide-item">',
          '<div class="swiper-zoom-container">', // 应用缩放配置项要提供这个节点
            '<img src="' + item + '" />',
          '</div>',
        '</div>',
      ].join('');
    });

    var wrapper = document.createElement('div'); // 创建一个 swiper-wrapper

    wrapper.className = 'swiper-wrapper';
    wrapper.innerHTML = html.join(''); // 把所有 swiper-slide 塞进去

    var container = document.createElement('div'); // 创建跟节点

    container.className = 'container';

    // 把所有 html 拼起来
    container.innerHTML = [
      '<div class="count">',
        '<span class="current">' + (idx || 1) + '</span>',
        '/',
        '<span class="total">' + list.length + '</span>',
      '</div>',
      wrapper.outerHTML,
      '<div class="swiper-pagination"></div>',
    ].join('');

    // 添加到 DOM 中
    document.body.appendChild(container);

    // 实例化一个 swiper
    new Swiper(container, {
      zoom: true, // 缩放开启
      loop: list.length > 1, // 如果图片只有一张，则不开启循环
      pagination: { // 分页配置
        el: '.swiper-pagination',
      },
      on: { // 事件监听
        paginationUpdate: function(e) { // 当分页发生变化的时候
          var idx = e.realIndex; // 拿到当前页索引

          // 赋值给分页计数器
          container.querySelector('.current').innerText = idx + 1;
        },
      },
    }).slideTo(idx, 0); // 默认展示初始图
  }

  previewImage(
  '5ef94c8e000109e118720764.jpg',
  [
    '5f057a6a0001f4f918720764.jpg',
    '5ef94c8e000109e118720764.jpg',
    '5ef15e4e00010b0018720764.jpg',
    '5f0561160001630718720764.jpg',
  ]);
  </script>
</body>
</html>
```

源码没有跳着走的逻辑，都加上了注释，相对好理解。

这个图片查看方法利用了 `swiper` 提供的滚动、手势缩放、手势拖动、分页的能力，实现相对简单，如果需要自己去实现相应功能，就需要花费大量的经历。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f0de9880aa761f702830614.jpg)

## 3. 小结

swiper 本身的定位并不是轮播，轮播是其应用场景之一，发挥想象，可以用 swiper 做许多事情。



### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
