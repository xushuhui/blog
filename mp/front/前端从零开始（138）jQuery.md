# jQuery

> jQuery is a fast, small, and feature-rich JavaScript library. It makes things like HTML document traversal and manipulation, event handling, animation, and Ajax much simpler with an easy-to-use API that works across a multitude of browsers. With a combination of versatility and extensibility, jQuery has changed the way that millions of people write JavaScript.(jQuery 官方介绍）

## 1. 什么是 jQuery

jQuery 是一个使用 `JavaScript` 编写的库，可以让开发者用更少的代码来完成业务逻辑。

许多年前前端的技术没有现在这么丰富，jQuery 和 JavaScript 也会被经常称为两个技术，因为使用 jQuery 完全可以替代掉使用原生的 JavaScript 操作 DOM、处理动画、处理 AJAX 等，这让两者之间的概念变得模糊。

可以对比一下删除一个节点的操作：

```javascript
// 使用JavaScript
var el = document.getElementById('element');

el.parentNode.removeChild(el);

// 使用 jQuery
$('#element').remove();
```

两者的区别一比较就出来了，jQuery 封装一层 DOM 操作，将原生的 DOM 方法向上层抽象，提供了一套更简洁的 API 来操作 DOM，同时也针对各个浏览器做了兼容性处理，如事件对象、事件的绑定方式等。

## 2. 引入 jQuery

jQuery 可以直接从官网下载，也可以用 `npm` 安装，也可以使用 `bower` 等这些包管理工具，本篇幅采用 CDN 的形式引入，本身 jQuery 就是一个 `.js` 文件，只需引入就能使用。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
```

引入之后就可以在全局下通过 `jQuery` 或者 `$` 调用 jQuery 了。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
<script>
  console.log($);
  console.log(jQuery);

  console.log($ === jQuery); // 输出：true
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef62c6009848b6d10260168.jpg)

## 3. 使用 jQuery

jQuery 使用 `$` 或者 `jQuery` 来生成一个 `jQuery` 对象，这里统一使用 $。

```javascript
<div class="element">

</div>

<script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
<script>
  $('.element').html('<p>这里是用 jQuery 插入的 HTML</p>');
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef62c8c09d2d36f23300494.jpg)

$ 可以接受一个 `CSS` 规范的选择器，用来选择元素，`html` 方法相当于设置 `DOM` 节点的 `innerHTML` 属性。

在 DOM 相关章节有提到，如果使用 `querySelector` 来选择节点，碰到节点不存在的情况下，会返回 null，这样就需要一层判断， jQuery 已经处理好了这些情况。

```javascript
<div>DOM节点</div>
<div class="element">

</div>

<script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
<script>
  $('.ele').html('<p>这里是用 jQuery 插入的 HTML</p>');

  console.log('不会影响正常程序执行');
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef62c7b090c59b213960486.jpg)

其可以接受的参数不仅仅是 `CSS 选择器`，也可以是一个原生 DOM 节点，一段 HTML 字符串等。

> jQuery 选择 $ 作为作为入口名称，一部分是因为简单，原生 DOM 提供的选择 DOM 节点的方法都是一长串，另一个原因是 $ 本身的发音 `dollar` 和 `DOM` 的发音接近。

## 4. 小结

jQuery 提供了一系列的方法使得操作 DOM 变得更简单，更具体的内容可以参考官方[官方文档](https://api.jquery.com/)。

现代框架的涌现，使得 jQuery 被使用到的场景也在变少，也有许多文献有 `不应将 jQuery 与现代框架一起使用` 的说法。其实 jQuery 在压缩后非常小，如果有需要完全可以考虑引入。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
