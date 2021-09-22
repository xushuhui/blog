# DOM 事件流

DOM 事件流描述了 DOM 时间响应的阶段、路径。

DOM 事件流也会被称为 DOM 事件模型。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e989c8e092494ef06400690.jpg)

## 1. 事件流阶段

事件流有三个阶段：

1. `捕获阶段` 从 window 开始，寻找触发事件最深层的节点，过程中如果有节点绑定了对应事件，则触发事件
2. `目标阶段` 找到事件触及的最深节点
3. `冒泡阶段` 从最深节点按照捕获的路径进行返回，过程中如果有节点绑定了对应事件，则触发事件

现代浏览器默认都会在冒泡阶段触发事件。

通过一个例子来简单的感受一下。

```javascript
<style>
  .box {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .size-100 {
    width: 100px;
    height: 100px;
    background: #4caf50;
  }

  .size-200 {
    width: 200px;
    height: 200px;
    background: chocolate;
  }

  .size-300 {
    width: 300px;
    height: 300px;
    background: wheat;
  }
</style>

<div class="box size-300">
  <div class="box size-200">
    <div class="box size-100">
    </div>
  </div>
</div>

<div class="result"></div>

<script>
  var boxes = document.querySelectorAll('.box');
  var result = document.querySelector('.result');

  boxes.forEach(function(box) {
    box.addEventListener('click', function() {
      var el = document.createElement('p');

      el.innerText = '现在触发点击事件的是' + this.className;

      result.appendChild(el);
    });
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9897f00a6d04e919360834.jpg)

点击后，观察输出可以发现，事件是点击到的最深层次的节点开始向上执行的。

即从 `size-100` 到 `size-200` 到 `size-300`，这就是冒泡的过程。

如果想让事件在捕获阶段就执行，可以传递 `addEventListener` 方法第三个参数。

## 2. addEventListener 的第三个参数

addEventListener 的第三个参数用来决定事件在冒泡阶段触发还是在捕获阶段触发，其为一个布尔值，传递 `false` 则事件会在冒泡阶段触发，传递 `true` 则会在捕获阶段触发。

```javascript
<style>
  .box {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .ele1 {
    background: wheat;
    width: 200px;
    height: 200px;
  }

  .ele2 {
    background: yellowgreen;
    width: 100px;
    height: 100px;
  }
</style>

<div class="box ele1">
  <div class="box ele2"></div>
</div>

<div class="result"></div>

<script>
var ele1 = document.querySelector('.ele1');
var ele2 = document.querySelector('.ele2');
var result = document.querySelector('.result');

function getElement(content) {
  var el = document.createElement('p');

  el.innerText = content;

  return el;
}

ele1.addEventListener('click', function() {
  result.appendChild(getElement('我是元素ele1'));
});

ele2.addEventListener('click', function() {
  result.appendChild(getElement('我是元素ele2'));
});
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9898030aebacb818920574.jpg)

根据默认浏览器事件是在冒泡阶段触发的规则，上述例子会先触发子节点 `.ele2` 的事件，再触发 `.ele1` 的事件。

如果想让 `.ele1` 在捕获阶段就触发事件，则在绑定事件的时候传递第三个参数为 `true` 即可。

```javascript
<style>
  .box {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .ele1 {
    background: wheat;
    width: 200px;
    height: 200px;
  }

  .ele2 {
    background: yellowgreen;
    width: 100px;
    height: 100px;
  }
</style>

<div class="box ele1">
  <div class="box ele2"></div>
</div>

<div class="result"></div>

<script>
var ele1 = document.querySelector('.ele1');
var ele2 = document.querySelector('.ele2');
var result = document.querySelector('.result');

function getElement(content) {
  var el = document.createElement('p');

  el.innerText = content;

  return el;
}

ele1.addEventListener('click', function() {
  result.appendChild(getElement('我是元素ele1'));
}, true);

ele2.addEventListener('click', function() {
  result.appendChild(getElement('我是元素ele2'));
});
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9898100acf9b9218920574.jpg)

这样 `.ele1` 的事件就会在捕获阶段触发。

## 3. 不符合 W3C 标准的事件流

早期的 IE 和 Netscape Navigator 是不符合标准的。

前者是使用事件冒泡流，后者使用事件捕获流。

前面的章节有提到过 `0级DOM事件` ，其提供的绑定事件的方式是不能指定事件触发的阶段的，其原因是在那个阶段下，还没有现在制定的 `DOM 事件流`。

当时并没有统一的标准，`0级DOM事件`也并不是一套官方出台的标准，所有相关内功全部由浏览器厂商决定。

后来 W3C 很好的整合了这两种模型，便有了现在的 DOM 事件流。

## 4. 冒泡的终点元素

这个问题其实经常会在面试中被问到，通常题目会是这样的：

> 请描述一下事件捕获和冒泡的具体流程

其实问的是事件从那个节点开始捕获，然后到目标节点，最后又在哪个节点冒泡结束。

大部分面试者会回答 `document`，其实根据事件对象的 `path` 属性就可以得到答案。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e989c3d0a6fda7719800804.jpg)

path 属性会返回事件冒泡的路径，其最后是到 `window` 对象才停止的。

其实这点在[标准](https://www.w3.org/TR/DOM-Level-3-Events/#dom-event-architecture)中也有描述。

> 注意：path 属性有兼容性问题，可以通过 can i use 确定。可以用标准中的 composedPath 代替。

## 5. 小结

开发过程中很少会取改变事件触发的阶段。但是事件流的概念依然重要，因为很多时候要阻止事件冒泡。

理解了事件流，可以理解事件委托的原理，事件委托相关的内容可以参阅事件相关的性能优化。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
