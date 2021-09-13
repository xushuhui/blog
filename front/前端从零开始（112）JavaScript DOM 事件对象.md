# DOM 事件对象

> Event 对象代表事件的状态，比如事件在其中发生的元素、键盘按键的状态、鼠标的位置、鼠标按钮的状态。(W3C)

事件对象会在事件被触发时获得，对象包含了当前事件的一些信息，如点击事件可以获取到点击的位置，键盘输入事件可以获取到按下的键。

## 1. 获取事件对象

在给 DOM 节点绑定事件时，需要传递一个事件处理器，其本质上是个函数，在事件触发时被调用。

在事件处理器被调用时，默认就会传递一个参数，这个参数就是事件对象。

```javascript
<input type="text" class="input-here">

<div class="log"></div>

<script>
  var inputEle = document.querySelector('.input-here');
  var logEle = document.querySelector('.log');

  inputEle.onkeydown = function(event) {
    var ele = document.createElement('p');

    ele.innerText = '按下了' + event.keyCode;

    logEle.appendChild(ele);
  }

  inputEle.addEventListener('keyup', function(event) {
    var ele = document.createElement('p');

    ele.innerText = '弹起了' + event.keyCode;

    logEle.appendChild(ele);
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b29a0af269b811000432.jpg)

输入一个字符的动作包含`按下键`和`松开键`，对应的事件就是 `onkeydown` 和 `onkeyup`，如果使用二级 DOM 事件，则可以不加 `on` 前缀。

例子中的事件处理器接收了一个参数，这个参数就是事件对象，参数名是可以随意的，一般情况下开发者会选择 `e` 或者 `event` 作为参数名。

`onkeydown` 和 `onkeyup` 是键盘相关的事件，所以可以获取到按下的键是哪个，对应的就是事件对象下的 `keyCode` 属性。

`keyCode` 属性是按下键的 `ASCII` 码，如数字 1 对应的就是 49， 数字2对应的是 50。具体可以参阅 [ASCII](https://zh.wikipedia.org/wiki/ASCII) 映射表。

## 2. 常用的事件对象下的属性和方法

> 以下内容会涉及到事件流相关的内容，可以参阅 DOM 事件流章节。

在符合 DOM2 标准的浏览器中，事件对象都具有以下属性和方法。

### 2.1 属性

#### 2.1.1 target

target 表示当前事件最终捕获到的目标。

```javascript
<div class="a">
  我是第一个节点 a
  <div class="b">
    我是第二个节点 b
    <div class="c">
      我是第三个节点 c
      <div class="d">
        我是第四个节点 d
        <div class="e">
          我是第五个节点 e
          <div class="f">
            我是最里面的一个元素 f
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<div class="result" style="margin-top: 16px;"></div>

<script>
  var resultEle = document.querySelector('.result');

  document.querySelector('.a').addEventListener('click', function(e) {
    resultEle.innerText = '捕获到的元素类名是' + e.target.className;
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b2b90a0f104c10680432.jpg)

可以看到事件绑定在类名为 `a` 的节点上，点击其子节点的时候，子节点就是最终捕获到的元素。

#### 2.1.2 currentTarget

通过 `currentTarget` 可以获取到目前触发事件的元素。

```javascript
<div class="a">
  我是第一个节点 a
  <div class="b">
    我是第二个节点 b
    <div class="c">
      我是第三个节点 c
      <div class="d">
        我是第四个节点 d
        <div class="e">
          我是第五个节点 e
          <div class="f">
            我是最里面的一个元素 f
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<div class="result" style="margin-top: 16px;"></div>

<script>
  var resultEle = document.querySelector('.result');

  document.querySelector('.a').addEventListener('click', function(e) {
    resultEle.innerText = [
      '当前触发事件的元素的类名是：',
      e.currentTarget.className,
      '。当前捕获到的元素类名是：',
      e.target.className,
    ].join('');
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b2d50a6e384210680432.jpg)

不论点击的是哪个子节点，currentTarget 都是表示当前触发事件的节点。

### 2.2 方法

#### 2.2.1 stopPropagation

调用此方法就会阻止事件的冒泡，使用到的场景大多为某个父元素和元素本身绑定了相同事件时。

```javascript
<style>
  .list {
    width: 300px;
    margin: 0 auto;
  }

  .list .item {
    width: 100%;
    border: 1px dashed #4caf50;
    border-bottom: 0;
    border-radius: 2px;
    padding: 16px;
  }

  .list .item:last-child {
    border-bottom: 1px dashed #4caf50;
  }

  .list .item button {
    border-radius: 2px;
    font-size: 14px;
    color: #666;
    outline: none;
  }

  .list .item button:active {
    background: #eee;
  }
</style>

<div class="list">
  <div class="item">
    <p>一句简单的介绍。</p>
    <button>点击我删除这条</button>
  </div>

  <div class="item">
    <p>两句简单的介绍。两句简单的介绍。</p>
    <button>点击我删除这条</button>
  </div>
</div>

<script>
  var items = document.querySelectorAll('.list .item');
  var btns = document.querySelectorAll('.list .item button');

  items.forEach(function(item) {
    item.addEventListener('click', function() {
      alert('马上进入到详情');
    });
  });

  btns.forEach(function(btn) {
    btn.addEventListener('click', function() {
      var parent = btn.parentNode;

      parent.parentNode.removeChild(parent);
    });
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b2e80a97192f18080468.jpg)

上述例子，在点击按钮的时候，虽然完成了删除操作，但还是会弹出一个框，触发到了父级的事件，这是冒泡特性导致的，所以需要阻止向上冒泡，

```javascript
<style>
  .list {
    width: 300px;
    margin: 0 auto;
  }

  .list .item {
    width: 100%;
    border: 1px dashed #4caf50;
    border-bottom: 0;
    border-radius: 2px;
    padding: 16px;
  }

  .list .item:last-child {
    border-bottom: 1px dashed #4caf50;
  }

  .list .item button {
    border-radius: 2px;
    font-size: 14px;
    color: #666;
    outline: none;
  }

  .list .item button:active {
    background: #eee;
  }
</style>

<div class="list">
  <div class="item">
    <p>一句简单的介绍。</p>
    <button>点击我删除这条</button>
  </div>

  <div class="item">
    <p>两句简单的介绍。两句简单的介绍。</p>
    <button>点击我删除这条</button>
  </div>
</div>

<script>
  var items = document.querySelectorAll('.list .item');
  var btns = document.querySelectorAll('.list .item button');

  items.forEach(function(item) {
    item.addEventListener('click', function() {
      alert('马上进入到详情');
    });
  });

  btns.forEach(function(btn) {
    btn.addEventListener('click', function(e) {
      // 阻止冒泡
      e.stopPropagation();

      var parent = btn.parentNode;

      parent.parentNode.removeChild(parent);
    });
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b2f50a5887f218080468.jpg)

#### 2.2.2 preventDefault

此方法可以取消事件的默认行为，比如超链接的点击，会发生跳转，跳转动作就是默认行为。

给超链接绑定点击事件，调用事件对象下的 `preventDefault` 属性，默认行为就会取消，即不会发生跳转。

```javascript
<a href="https://imooc.com">冲鸭！！前往慕课网！！</a>

<script>
  var aTag = document.querySelector('a');

  aTag.onclick = function(e) {
    e.preventDefault();

    alert('跳转被终止！');
  };
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b3090a02094015960432.jpg)

## 3. 兼容性问题

早期 IE 下的事件模型与 DOM 标准提供的有些不同。

如事件对象，在 IE8 之前并不是通过传递给事件处理器获取的，而是要通过 `window.event` 获取。

```javascript
<div id="ele">
  点我
</div>
<script>
  var ele = document.getElementById('ele');

  ele.onclick = function(e) {
    alert(e); // undefined
    alert(window.event);
  }
</script>
```

![图片描述](https://img.mukewang.com/wiki/5e92b3340a41fd5810080448.jpg)

以下代码在 IE8 中，第一个 alert 将会返回 undefined，第二个才会是事件对象。

部分事件属性也不同，如标准中的 `target` 属性，在早期 IE 下需要用 `srcElement` 替代。

建议对兼容性相关内容做个了解即可，框架通常会处理好兼容性问题。

## 4. 小结

事件对象包含了事件相关的信息，有事件对象，才能对各个事件做更深层次的交互优化。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
