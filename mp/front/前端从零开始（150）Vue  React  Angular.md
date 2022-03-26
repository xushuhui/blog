# Vue、React、Angular

`Vue`、`React`、`Angular` 常被一起称作三大框架、现代框架。

三大框架是目前驱动前端项目底层的最常用的框架。随着前端行业从业人员的增加，更易上手的 `Vue` 和 `React` 占据了更大部分市场。

> 本章节不会探讨这些框架的具体用法

## 1. 前端框架改变了什么

随着 `AJAX` 的普及以及浏览器性能的提升，前端的交互越来越复杂，前端工程师的工作职责也在变广。

其中最容易让代码变得复杂的业务逻辑就是 `DOM` 操作。

在没有任何框架的情况下，给一个按钮切换文案可能是这样的：

```javascript
var btn = document.querySelector('.btn');

btn.addEventListener('click', function() {
  var txt = btn.innerText;

  if (txt === '开') {
    btn.innerText = '关';
  } else {
    btn. innerText = '开';
  }
});
```

如果要往里面插入各种逻辑，如发起请求，请求后对应界面上的某个 DOM 的复杂改变，代码就会变得越来越难维护。

如果有维护过老项目，对这方面的印象会更深刻。老项目可能会充斥着各种字符串拼接 HTML，代码可读性差，逻辑难以被后人扩充维护，小模块的重构又怕影响到项目根基，这些问题会随着时间慢慢暴露出来。

再就是花了太多时间在 `DOM` 操作上，为了取某个父级会经历多次 `.parentNode`，导致经常要去数数等这些问题。

不管是性能还是可维护性，总归来讲就是在 `DOM` 操作上吃了太多亏，这一点也是出现这些前端框架的出要原因：**UI 与 数据的同步太费事儿**。

对于新人，刚学习前端框架感到最震撼的点通常都是框架对 `DOM` 操作的解放，以 `Vue2.x` 为例：

```javascript
<template>
	<button @click="toggle">
    {{ text }}
  </button>
</template>

<script>
	export default {
    data() {
      return {
        text: '开',
      }
    },
    method: {
      toggle() {
        this.text = (this.text === '开') ? '关' : '开';
      },
    },
  };
</script>
```

以数据来驱动视图，特别是在列表渲染上，这个特性的优点就能被放的很大，其具体实现原理可以学习对应框架的底层细节。

所以前端框架带来的最大改变，就是解放了大量的操作 `DOM` 的工作，让开发者更注重逻辑上的表现。

其他的改变，还有组件化、工程化等，具体开发就能体会到。

## 2. 小结

前端框架是目前必备技能，可以选其一针对学习研究，再扩展其他。


### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
