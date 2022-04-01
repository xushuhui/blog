# debugger

> debugger 语句调用任何可用的调试功能，例如设置断点。 如果没有调试功能可用，则此语句不起作用。(MDN)

`debugger` 通常用于调试，主要是为了设置一个断点。

如果浏览器支持 `debugger`，那碰到 debugger 就会暂停程序的执行，提供调试功能，如单步调试、跳出当前函数、结束调试等。

## 1. 使用 debugger

```javascript
debugger; // 设置断点
```

在需要设置断点的地方写上 debugger 即可。

```javascript
console.log(1);

var str = '在这里暂停';

debugger; // 设置断点

console.log(str);

console.log(1 + 1);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ec00c3a0aac534914020837.jpg)

断点设置好之后可以在`开发者工具`的 `Sources` 面板进行调试。

## 2. 其他设置断点的方式

假设对其他网站的某个实现细节很感兴趣，但又不能直接窥探出原理，也可以借助断点来进行调试。

这种情况下需要在 `开发者工具` 的 `Sources` 面板找到对应的源码，打上断点。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ec00c460a4b294f21540910.jpg)

在源码的对应行号出点击，即可设置上断点，如果是已经执行过的代码，则需要刷新才会在断点处暂停程序。

很多情况下，都会利用事件来定位源码位置。

一个节点上的事件，可以通过 `Elements` 面板的 `Event Listeners` 来查看定位。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ec00e260a10ae5e20460865.jpg)

## 3. 小结

debugger 用于设置断点，调试非常有用。

如果没有特殊需求，**一定要确保线上 `debugger` 不会被执行！一定要确保线上 `debugger` 不会被执行！一定要确保线上 `debugger` 不会被执行！**

这一点非常关键，带上线了直接影响用户体验，可能公司第二天就倒了。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
