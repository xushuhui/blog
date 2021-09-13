# DOM 和 JavaScript 的关系

## DOM 与 JavaScript 是分离的。

DOM 有自己的一套标准，JavaScript 也有自己的一套标准。

JavaScript 是对标准的实现，为一种语言，而 DOM 标准定义了一系列的接口，由此可以看出，他们两者是可以毫不相干的。

但是在 Web 开发中，页面内容的展示全部通过浏览器解析展现，JavaScript 想动态的修改页面，就必须由浏览器提供一些方法，交给开发者来操作页面上的元素，因为 JavaScript 本身是没有操作这些元素的能力的。

浏览器承担了实现与暴露 DOM 接口的工作，根据标准实现一系列方法，随后暴露给开发者使用。

如 `document` 对象，表示当前的页面，也可以理解成根节点，JavaScript 本身是没有这个全局对象的。

可以通过遍历 document 对象的属性，来观察一个 DOM 节点都有些什么属性和方法。

```javascript
for (var i in document) {
  console.log(i, document[i])
}
```

一个节点的属性非常多，包括许多事件、子节点、操作节点的方法等。

浏览器通过暴露这些 DOM 相关的内容给开发者，开发者通过 JavaScript 进行操作。

> 对浏览器而言，document 实际上不是最顶层的节点，再向上还有 window ，这一点可以在事件相关的章节体现。

![图片描述](https://img.mukewang.com/wiki/5e82de6c081861d414400768.jpg)

## 小结

JavaScript 和 DOM 本身是可以没有关系的，但是开发者需要操作 DOM ，浏览器实现了对应的方法，暴露给开发者，开发者使用 JavaScript 来调用以达到操作 DOM 的目的。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
