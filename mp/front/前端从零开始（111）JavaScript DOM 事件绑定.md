# DOM 事件绑定

DOM 提供了许多事件供开发者进行绑定，以响应各种操作，丰富页面交互。

想要触发事件，就得先给 DOM 节点绑定事件，提供事件处理器。

## 1. 直接在 HTML 上提供事件

这种方式是将事件内联在 HTML 代码中。

```javascript
<style>
  .box {
    width: 100px;
    height: 100px;
    background: green;
  }
</style>

<div class="box" onclick="alert('耶耶耶耶耶！')"></div>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e8e12360a2de61f23320432.jpg)

通过设置 HTML 的 `onclick` 属性，可以绑定点击事件，属性内可以写简单的 JavaScript 代码。

但是可以看到这种方式是有局限性的，写大量的 JavaScript 代码在里面肯定是不合理的。

过去网页的交互相对较少，会在 body 标签上绑定 `onload` 事件，即页面加载完毕后做的事情。

```javascript
<script>
  function load() {
    alert('页面加载完毕啦！');
  }
</script>
<body onload="load()">
  <div>我是一段滥竽充数文案。</div>
</body>
```

`onload="load()"` 即表示在页面加载完成后，执行 `load` 函数。现在翻阅部分文献，依然能够看到这种写法。

## 2. 设置事件处理器属性

通过`事件处理器属性设置`绑定事件，需要先获取到 DOM 节点。

```javascript
<style>
  .btn {
    border: 1px solid #4caf50;
    font-size: 22px;
    padding: 8px 12px;
    color: #4caf50;
    outline: none;
  }
  .btn:active {
    background: #eee;
  }
</style>

<button class="btn">
  我是一个可以改变一切的按钮
</button>

<script>
  var btn = document.querySelector('.btn');

  btn.onclick = function() {
    var text = btn.innerText;

    btn.innerText = text.replace('一切', '世界');
  };
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e8e12ac0a08c75a23320432.jpg)

通过 `onclick` 就可以设置按钮的点击事件。

如果需要清除事件，可以重新将属性重新设置为 `null`。

```javascript
<style>
  .btn {
    border: 1px solid #4caf50;
    font-size: 14px;
    padding: 8px 12px;
    color: #4caf50;
    outline: none;
  }
  .btn:active {
    background: #eee;
  }

  p {
    font-size: 22px;
  }
</style>


<button class="btn">
  我是一个可以改变一切的按钮
</button>

<p>1</p>

<script>
  var btn = document.querySelector('.btn');
  var p = document.querySelector('p');
  var total = 1;

  btn.onclick = function() {
    p.innerText = ++total;

    if (total >= 5) {
      btn.onclick = null;
    }
  };
</script>
```

这种绑定事件的方式依然常用，各种文献的 `demo` 也会采用这种方式，但缺点是同一个事件在没有其他辅助条件下，只能绑定一个事件处理器。

## 3. 使用 element.addEventListener 绑定事件

使用 DOM 节点的 `addEventListener` 方法也可以绑定事件。

```javascript
<style>
  .btn {
    border: 1px solid #4caf50;
    font-size: 14px;
    padding: 8px 12px;
    color: #4caf50;
    outline: none;
  }
  .btn:active {
    background: #eee;
  }

  input {
    padding: 10px 12px;
    font-size: 14px;
    outline: none;
  }

  p {
    font-size: 22px;
  }
</style>


<input type="text">

<button class="btn">
  按钮
</button>

<p>1</p>

<script>
  var btn = document.querySelector('.btn');
  var input = document.querySelector('input');
  var p = document.querySelector('p');

  var total = 1;

  btn.addEventListener('click', function() {
    input.value = input.value.split('').reverse().join('');
    btn.value = input.value;
    p.innerText = input.value;
  });

  input.addEventListener('keyup', function() {
    btn.innerText = input.value;
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e8e13560a6d77cd17520432.jpg)

`addEventListener` 的第一个参数是事件类型，和使用`事件处理器属性`与`HTML内联事件`不同，这个方法的事件类型不需要 `on` 前缀。

第二个参数就是事件处理器，即触发事件时要做的事情。

使用 `addEventListener` 可以绑定多个事件处理器。

```javascript
<style>
  .btn {
    border: 1px solid #4caf50;
    font-size: 14px;
    padding: 8px 12px;
    color: #4caf50;
    outline: none;
  }
  .btn:active {
    background: #eee;
  }
</style>

<button class="btn">
  按钮
</button>
<p></p>

<script>
  var btn = document.querySelector('.btn');
  var p = document.querySelector('p');
  var total = 1;

  btn.addEventListener('click', function() {
    btn.innerText = ++total;
  });

  btn.addEventListener('click', function() {
    p.innerText = total;
  });
</script>
```

一个事件绑定多个事件处理器，其执行顺序是按照绑定顺序来的。

如果需要去除事件，可以使用 `removeEventListener` 方法。

```javascript
<style>
  .btn {
    border: 1px solid #4caf50;
    font-size: 14px;
    padding: 8px 12px;
    color: #4caf50;
    outline: none;
  }
  .btn:active {
    background: #eee;
  }
</style>

<button class="btn">
  按钮
</button>

<script>
  var btn = document.querySelector('.btn');
  var total = 1;

  function fn() {
    btn.innerText = ++total;

    if (total > 5) {
      btn.removeEventListener('click', fn);
    }
  }

  btn.addEventListener('click', fn);
</script>
```

使用 `removeEventListener` 去除事件，需要传递指定的事件，且事件处理器必须与绑定事件传入的一样，为同一个。

addEventListener 与 removeEventListener 还具有第三个参数，会在事件流章节进行讨论。

> IE8 及以下都不支持该方法，需要使用 IE 提供的 attachEvent 与 detachEvent 来处理事件

## 4. DOM 事件级别

目前常讨论的 DOM 级别有 4 级，为 `DOM 0级` 至 `DOM 3级`，这里的`级`可以理解成标准的版本。

这 4 级内容中， [DOM 1级](https://www.w3.org/TR/2000/WD-DOM-Level-1-20000929/)没有提供事件相关的内容，所以不会拿来讨论，因此 DOM 事件就被分为了常说的 `DOM 0级事件`、`DOM 2级事件`、`DOM 3级事件`。

* `DOM 0级` 提供的事件绑定方式为在`HTML 上内联事件`与`事件处理器属性`。
* `DOM 2级` 提供了 `addEventListener` 与 `removeEventListener` 方法。
* `DOM 3级` 则是在 2 级的基础上进行了扩充，添加了更多的事件类型。

事实上 DOM 0 级事件并不是标准中的，在 W3C 制定 DOM 标准前，部分浏览器已经有了 DOM 模型，也有自己相应的事件，这些出现在标准第一个版本之前的，就被称为了 `DOM 0`。

## 5. 小结

绑定 DOM 事件可以通过以下三种方式：

1. 将事件内联在 HTML 中
2. 使用 事件处理器属性 绑定事件
3. 使用 `addEventListener` 与 `removeEventListener` 处理事件

前两种方式不能绑定多个事件处理器，三种绑定方式由不同版本的 DOM 标准提供，通常使用 DOM 级别进行区分。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
