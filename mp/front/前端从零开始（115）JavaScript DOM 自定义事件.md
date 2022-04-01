# 自定义事件

自定义事件主要会被用于框架、组件设计与实现中。

自定义的事件有许多的创建方式，但实际的业务场景中几乎不会被用到，网络上的文献记载其具体的使用场景也相对较少。

## 1. 使用 Event 构造函数

使用 `Event` 构造函数就可以创建一个自定义事件。

```javascript
<style>
  .btn { border: 1px solid #4caf50; padding: 8px 12px; min-width: 200px; color: #4caf50; background: white; outline: none; }
  .btn:active { background: #4caf50; color: white; }
</style>

<div>
  <button class="btn">点击我</button>
</div>

<script>
  var afterClick = new Event('afterclick');

  var btnEle = document.querySelector('.btn');

  btnEle.addEventListener('afterclick', function() {
    alert('我是点击事件执行完后做的事情');
  });

  btnEle.onclick = function() {
    alert('我被点击了');

    this.dispatchEvent(afterClick);
  };
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9de13e0a5a0f0a17400432.jpg)

使用 `new Event` 可以创建一个自定义事件，事件名就是构造函数的第一个参数 `afterclick`，表示点击事件完成后做的事情。

创建一个自定义事件后需要给 DOM 元素绑定这个事件，只有绑定后才能触发，使用 `addEventListener` 来绑定事件。

随后再给按钮绑定点击事件，在事件末尾，即事情做完后，使用 `dispatchEvent` 触发这个自定义事件。

**自定义事件是需要手动触发的！**

`Event` 构造函数还支持第二个参数，其接受一个对象，可以传递三个属性，都为布尔值：

* `bubbles` 默认 false ，表示是否会冒泡
* `cancelable` 默认 false ， 表示事件是否可以被取消
* `composed` 默认 composed， 表示事件是否会在 [Shadow DOM](https://developer.mozilla.org/zh-CN/docs/Web/Web_Components/%E5%BD%B1%E5%AD%90_DOM) 根节点之外触发。

## 2. 使用 CustomEvent 构造函数

上面使用 `Event` 构造函数的例子，将其替换成 `CustomEvent` 构造函数也是一样可以执行的。

```javascript
<style>
  .btn { border: 1px solid #4caf50; padding: 8px 12px; min-width: 200px; color: #4caf50; background: white; outline: none; }
  .btn:active { background: #4caf50; color: white; }
</style>

<div>
  <button class="btn">点击我</button>
</div>

<script>
  var afterClick = new CustomEvent('afterclick');

  var btnEle = document.querySelector('.btn');

  btnEle.addEventListener('afterclick', function() {
    alert('我是点击事件执行完后做的事情，我被改成了 CustomEvent');
  });

  btnEle.onclick = function() {
    alert('我被点击了');

    this.dispatchEvent(afterClick);
  };
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9de1610a3eaa3a17400432.jpg)

两个例子效果是一样的。

其主要的区别在参数和工作环境上，CustomEvent 是可以在 WebWorker 中被使用的，而 Event 不行。

CustomEvent 可以在构造函数的第二个参数中传递 `detail` 属性，在事件触发时，事件对象中就会携带这个 detail 属性。

假设现在想完成一个键盘的点击事件，即键盘上某个键按下并弹起后做的事情。

```javascript
<style>
  input {width: 200px;padding: 8px;font-size: 16px;outline: none;border: 1px dashed #4caf50;}
  input:focus {border: 1px solid #4caf50;}
</style>

<div>
  <input type="text">
</div>

<script>
  var inputEle = document.querySelector('input');

  var onKeyClick = function(e) {
    console.log(e);
    alert('现在输入框内容是：' + e.detail.value + '，触发的键是：' + e.detail.keyCode);
  };

  inputEle.addEventListener('keyup', (e) => {
    console.log('键盘按键弹起了');

    var keyClick = new CustomEvent('keyclick', {
      detail: {
        value: e.target.value,
        keyCode: e.keyCode,
      },
    });

    inputEle.addEventListener('keyclick', onKeyClick);

    inputEle.dispatchEvent(keyClick);

    inputEle.removeEventListener('keyclick', onKeyClick);
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9de1990a5793ee16160450.jpg)

这里通过 `keyup` 事件，在事件处理器的最末尾增加了一个 `keyclick` 事件。这里并没有结合 `keydown` 来判断按键的落下和弹起，因为一个按键要弹起，必定得先落下，所以只需要监听 `keyup`。

其实可以看出这段代码比较奇怪，真正的业务场景并不会这样写，会选择直接调用 `onKeyClick` 函数。

```javascript
<style>
  input {width: 200px;padding: 8px;font-size: 16px;outline: none;border: 1px dashed #4caf50;}
  input:focus {border: 1px solid #4caf50;}
</style>

<div>
  <input type="text">
</div>

<script>
  var inputEle = document.querySelector('input');

  var onKeyClick = function(value, keyCode) {
    alert('现在输入框内容是：' + value + '，触发的键是：' + keyCode);
  };

  inputEle.addEventListener('keyup', (e) => {
    console.log('键盘按键弹起了');

    onKeyClick(e.target.value, e.keyCode);
  });
</script>
```

这段代码的执行结果和采用 `CustomEvent` 的效果是一样的。

这就是为什么自定义事件更常用于框架或者库，因为暴露事件有时候比单纯的提供回调配置项更好理解和解耦。

## 3. document.createEvent

使用 document.createEvent 也可以用来创建自定义事件，但其由于许多配套属性、方法都已经从标准中移除，MDN 也不再建议开发者使用，所以这里不再深入探讨，只做了解。

## 4. 小结

自定义事件不常用，主要被应用于框架级别的设计上，日常开发很少有使用场景，许多场景下还会让代码变得冗余复杂。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
