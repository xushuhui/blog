# 什么是 BOM

BOM 即浏览器对象模型（browser object model），其提供了一系列接口供开发者使用 `JavaScript` 与浏览器窗口进行交互。

BOM 不像 ECMAScript 和 DOM 有一套自己的标准，BOM 是没有公共组织制定标准的。

神奇的是所有现代浏览器在 BOM 的相关内容上几乎一致，所以 BOM 也足够通用，所有浏览器的实现几乎一致。

## 1. BOM 的作用

BOM 提供的是与浏览器窗口交互的能力，其包含了一些处理窗口的方法，如打开新窗口，控制新窗口大小，也提供了窗口相关的属性，如窗口尺寸。

BOM 的操作入口可以理解为 `window` 对象，即浏览器下的全局对象。

```javascript
window.location.href = '//imooc.com'; // 跳转

window.onload = function() {
  // 当前窗口的页面加载完毕做的事情
};

window.open('https://immoc.com'); // 打开新的窗口

window.navigator.userAgent; // 获取 UA
```

## 2. 理解 BOM、DOM、JS 的关系

部分读者可能会进入一个难以理解的怪圈。

全局的对象，如 `String`、`Boolean`、`Function`，或者一些方法 `parseInt`，`isNaN`，甚至是 DOM 的入口 `document`，这些都是被放在 `window` 下供开发者访问的。

而 BOM 没有自己的标准，也基本和这些东西没有半毛关系，为什么这些内容会被放在作为 BOM 操作的入口的 `window` 对象下呢？

之前也有提过，JavaScript 有他自己独立的标准，本身是不具备和浏览器交互的特性的，交互的接口都由浏览器来提供。

在 ECMAScript 的标准中，这些全局的对象都是被放在 `Global` 上的，而标准又没指出如何直接访问 `Global` 对象，所以在浏览器上这个所谓的 `Global` 就被 `window` 对象所替代，所有 `Global` 下的内容全部被放到了 `window` 下。

`window` 对象就这样作为了全局对象，并且 DOM 相关的内容，与窗口交互的方式都放在了 `window` 对象下。

> 扩展：globalThis
>
> `globalThis` 指向当前环境的全局的 this，也可以理解为指向全局对象。
>
>
> 这个属性在浏览器中指向的就会是 `window`，而在 `node` 环境下指向的就是 `global`。
>
>
> 这个属性最大的好处就是开发通用的库不必要再使用 window 或 global 这样的标志性对象来判断当前的宿主环境。
>
>
> 需要注意的是 globalThis 还是试验性的内容，chrome 和 新版本的 node 已经支持。

## 3. window.window.window.window…window

window 对象下还有一个名为 `window` 的属性，其指向 window 对象。

也就是说可以像套娃一样无限套。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea467d20a14bdf014480570.jpg)

这样设计也是有原因的，假设没有 `window.window` 这个属性，那就访问不到全局的 `window` 了，这就等于不能直接访问到全局对象，如果想要访问直接访问 window 对象，则必须自建一个变量用于放置 window 对象。

```javascript
var window = this; // 在代码最外层

window.location.href = 'https://imooc.com';
```

## 4. 小结

BOM 没有自己的标准，在浏览器中 window 对象就是 BOM 相关内容的入口。

window 对象是 BOM 相关内容的入口。

在浏览器环境中，window 对象就是全局对象，所有 DOM 相关内容与 ECMAScript 描述的全局对象等内容也被放在 window 对象下。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
