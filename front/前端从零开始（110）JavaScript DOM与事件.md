# DOM 与事件

> 事件是您在编程时系统内发生的动作或者发生的事情，系统通过它来告诉您在您愿意的情况下您可以以某种方式对它做出回应。例如：如果您在网页上单击一个按钮，您可能想通过显示一个信息框来响应这个动作。在这篇文章中我们将围绕事件讨论一些重要的概念，并且观察它们在浏览器上是怎么工作的。这篇文章并不做彻底的研究仅聚焦于您现阶段需要掌握的知识。—— MDN

事件即某个情况、某个事情。

* 当按钮被点击
* 视频播放、暂停
* 关闭浏览器
* 页面加载完毕
* 调整浏览器窗口大小

上述情况都是事件。

## 1. DOM 事件

> DOM 事件被发送用于通知代码相关的事情已经发生了。每个事件都是继承自 Event 类的对象，可以包括自定义的成员属性及函数用于获取事件发生时相关的更多信息。事件可以表示从基本用户交互到渲染模型中发生的事件的自动通知的所有内容。—— MDN

DOM 事件是指给 DOM 节点在触发某个条件下要做的事情，如：当按钮被点击的时候改变背景色。

```javascript
<style>
  .change-bg {
    border: 1px solid black;
    height: 40px;
    width: 120px;
    border-radius: 2px;
    margin-top: 16px;
    outline: none;
    cursor: pointer;
  }

  .change-bg:active {
    background: #efefef;
  }

  .box {
    width: 120px;
    height: 120px;
    background: #4caf50;
    border-radius: 60px;
  }
</style>
<div class="box"></div>

<button class="change-bg">戳这里改变背景色</button>
<script>
  var boxEle = document.querySelector('.box');
  var btnEle = document.querySelector('.change-bg');

  // 随机生成一个颜色 具体实现可以不管
  function getColor() {
    return '#' + ('00000' + (Math.random() * 0x1000000 << 0).toString(16)).slice(-6);
  }

  btnEle.onclick = function() {
    boxEle.style.backgroundColor = getColor();
  };
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e8c8df60a6d449117800432.jpg)

如上述例子中按钮的 `onclick` 属性，当他被赋值一个函数的时，这个函数就会在按钮被点击的时候触发。

`onclick` 属性是一种`事件处理器属性`，表示单击或点击事件，当想指定按钮在被点击的时候要做的事情时，就可以给这个属性赋值。

赋值的函数通常被称为`事件处理器`，即事件被触发时候时候执行的代码块，部分文献中会称为事件处理程序。

通常给 DOM 节点设置事件的操作，会被称为绑定事件，上述例子就是给一个按钮绑定了点击事件。

> 绝大部分事件处理器属性都是以 `on` 开头的。

## 2. JavaScript 与 DOM 事件

DOM 事件是由`DOM标准`提供规范，浏览器对其进行具体实现的。绑定事件需要借助 `DOM` 提供的接口，然后由 JavaScript 提供事件处理器。

使用 JavaScript 来给 DOM 节点绑定事件有多种方式，都是由 `DOM事件标准` 提供的，具体可以参阅事件绑定章节。

## 3. 小结

事件可以理解为某个情况，如点击的时候、关闭的时候。这些情况下要做的事情，就是事件处理器（事件处理程序）。

DOM 事件是由 DOM 标准预先定义好的接口，由 JavaScript 提供事件处理器，来决定当事件被触发时要做的事情。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
