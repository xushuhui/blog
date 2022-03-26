# WebComponents

> Web Components 是一套不同的技术，允许您创建可重用的定制元素（它们的功能封装在您的代码之外）并且在您的 web 应用中使用它们。

`Web Components` 可以理解成自定义组件。

使用 `Web Components` 可以达到如原生的 `video` 标签的效果，`video` 标签在被浏览器解析后有许多子节点，如播放控制器部分的控制按钮，这些在使用 `video` 的时候是不需要关心的，也不用手动书写这些子节点，同时又能被多处复用。

这些子节点通常是看不见的，藏起来的，但是通过 `Chrome` 的设置可以显示出来。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f2f97950a18eec015390755.jpg)

这种看不见的子节点被称为 `Shadow DOM`，许多文献称为 `影子节点` 或者 `影子 DOM`。

`Web Components` 具有非常好的隔离性，特别在样式表现上，可以做到完全隔离。

## 1. 使用

一个 `Web Components` 继承自 `HTMLElement`，然后通过 `customElements.define()` 来注册这个 `Component`。

以下例子？会创建一个名为 `to-imooc` 的组件。

```javascript
class ToImooc extends HTMLElement { // Web Components 继承自 HTMLElement
  constructor() {
    super(); // 调用父类构造函数

    const link = document.createElement('a'); // 创建一个 a 元素

    link.href = 'javascript:;'; // 默认什么都不做
    link.innerText = 'GOGOGO! Imooc 出发!'; // 显示的文案

    link.addEventListener('click', () => { // 跳转事件
      alert('坐稳啦！要发车啦！！！');

      window.location.href = '//imooc.com';
    });

    this.append( // this 就是自定义的节点
      link,
      document.createElement('br'),
    );
  }
}

window.customElements.define('to-imooc', ToImooc);
```

这样就创建好了一个自定义的组件。

然后在 `HTML` 中使用就好了。

```javascript
<to-imooc></to-imooc>
<to-imooc></to-imooc>
<to-imooc></to-imooc>
<to-imooc></to-imooc>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f2f97ce0ae642c620080954.jpg)

这样就完成了一个简单的组件，可以重复使用。

## 2. 应用

根据 `Can I Use` 可以看出，该特性现在浏览器支持并不广泛，不太适合在生产环境使用，但基于该特性，还是衍射除了许多 “未来”。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f2f97dd096b8dc325861372.jpg)

### 2.1 微前端

`微前端` 尚处在发展时期，其核心概念和 `微服务` 相似。

现阶段较为常用的微前端框架为 `single-spa` 和 [qiankun](https://qiankun.umijs.org/zh/guide/#%E4%BB%80%E4%B9%88%E6%98%AF%E5%BE%AE%E5%89%8D%E7%AB%AF)，后者是基于前者实现的。

该技术能做到 `技术栈无关`，即一个应用，能由多个不同技术的子应用构成，同时做到子应用的相互隔离，这里的隔离就可以选择采用 `Web Components` 实现。

### 2.2 Omi

[Omi](https://tencent.github.io/omi/site/docs/cn.html) 是由腾讯开源的跨端框架，其就是基于 Web Components 来设计的。

## 3. 小结

虽然 `Web Components` 还存在兼容性问题，但可以作为预备知识进行了解，或者应用于一些内部对兼容性无要求的项目。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
