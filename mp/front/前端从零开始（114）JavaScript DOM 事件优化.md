# 事件相关的优化

大部分的事件触发依赖于用户与浏览器的交互，但用户的行为是不可控的，许多交互设计上的缺陷与无法考虑到的因素会导致事件的频繁触发。

当事件处理器内部包含大量的操作，又不需要如此快速的响应事件时，就需要采用一些手段来限制事件处理器的执行。

事件的优化主要有两个目的：

1. 减少不必要的 HTTP 请求
2. 减少本机性能的消耗

## 1. 交互设计

通过交互的设计来优化事件是最常用到的方式。

如用户点击删除后将按钮禁止。

```javascript
<style>
  .list .item {display: flex;justify-content: space-between;border-bottom: 1px dashed #4caf50;padding: 8px 0;}
  .list .item .caption {font-weight: 700;}
  .list .item .operates .delete {border: 1px solid rgb(177, 107, 107);color: rgb(207, 72, 72);outline: none;cursor: pointer;}
</style>

<div class="list">
  <div class="item">
    <div class="content caption">
      今天要做的事情
    </div>
    <div class="operates caption">
      操作
    </div>
  </div>

  <div class="item">
    <div class="content">
      吃火锅
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>

  <div class="item">
    <div class="content">
      和小姐姐聊天
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>
</div>

<script>
  var listEle = document.querySelector('.list');
  var deleteEle = document.querySelectorAll('.delete');

  deleteEle.forEach(function(el) {
    el.addEventListener('click', function() {
      console.log('开始删除...');

      setTimeout(function() {
        var itemEl = el.parentNode.parentNode;

        listEle.removeChild(itemEl);

        console.log('删除成功');
      }, 1500);
    });
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6cc00a18b40818280756.jpg)

上述例子没有在用户第一次点击后，对按钮做一个禁止，或者采用一些上锁操作。

这种情况下用户可能会点击多次，删除操作通常会发送请求到服务端做处理，不对交互做优化可能会增加服务端的压力。

通过给予按钮状态就可以改善这个情况。

```javascript
<style>
  .list .item {display: flex;justify-content: space-between;border-bottom: 1px dashed #4caf50;padding: 8px 0;}
  .list .item .caption {font-weight: 700;}
  .list .item .operates .delete {border: 1px solid rgb(177, 107, 107);color: rgb(207, 72, 72);outline: none;cursor: pointer;}
</style>

<div class="list">
  <div class="item">
    <div class="content caption">
      今天要做的事情
    </div>
    <div class="operates caption">
      操作
    </div>
  </div>

  <div class="item">
    <div class="content">
      吃火锅
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>

  <div class="item">
    <div class="content">
      和小姐姐聊天
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>
</div>

<script>
  var listEle = document.querySelector('.list');
  var deleteEle = document.querySelectorAll('.delete');

  deleteEle.forEach(function(el) {
    el.addEventListener('click', function() {
      console.log('开始删除...');

      el.setAttribute('disabled', 'disabled');
      el.style.color = 'rgb(226, 174, 174)';
      el.style.borderColor = 'rgb(226, 174, 174)';
      el.style.cursor = 'wait';
      el.innerHTML = '处理中...';

      setTimeout(function() {
        var itemEl = el.parentNode.parentNode;

        listEle.removeChild(itemEl);

        console.log('删除成功');
      }, 1500);
    });
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6d0e0a1e014a18280756.jpg)

在用户第一次点击按钮后，就给予按钮禁止点击的状态，同时通过样式区分给予用户一个反馈，在提高用户体验的同时，优化了整个事件。

## 2. 事件委托（代理）

事件委托是利用事件冒泡的特性实现的，事件委托也被称为事件代理。

通过字面意思就可以理解，子节点的事件交给父节点来执行，一旦父节点发现子节点触发了对应的事件，就执行对应的事件处理器。

如：当点击按钮的时候，删除列表上的项

```javascript
<style>
  .list .item {display: flex;justify-content: space-between;border-bottom: 1px dashed #4caf50;padding: 8px 0;}
  .list .item .caption {font-weight: 700;}
  .list .item .operates .delete {border: 1px solid rgb(177, 107, 107);color: rgb(207, 72, 72);outline: none;cursor: pointer;}
</style>

<div class="list">
  <div class="item">
    <div class="content caption">
      今天要做的事情
    </div>
    <div class="operates caption">
      操作
    </div>
  </div>

  <div class="item">
    <div class="content">
      吃火锅
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>

  <div class="item">
    <div class="content">
      和小姐姐聊天
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>
</div>

<script>
  var listEle = document.querySelector('.list');

  listEle.addEventListener('click', function(e) {
    var el = e.target;

    if (el.className === 'delete') {
      console.log('开始删除...');

      el.setAttribute('disabled', 'disabled');
      el.style.color = 'rgb(226, 174, 174)';
      el.style.borderColor = 'rgb(226, 174, 174)';
      el.style.cursor = 'wait';
      el.innerHTML = '处理中...';

      setTimeout(function() {
        var itemEl = el.parentNode.parentNode;

        listEle.removeChild(itemEl);

        console.log('删除成功');
      }, 1500);
    }
  });
</script>
```

和上个小节对比效果，其实是一样的，但是这份示例代码中只在 `.list` 节点上绑定了事件，而上个小节则给每个按钮绑定了一个事件。

其关键的就是事件对象下的 `target` 属性，该属性表示当前事件流最终捕获到的元素。

很明显，根据 HTML 结构，删除按钮就是那一分支中能捕获到的最终节点。

当事件流到达捕获阶段后，则开始向上冒泡，进入冒泡阶段，在冒泡阶段会执行绑定在 `.list` 上的点击事件，在事件中对事件对象的 `target` 进行判定，如何条件就会执行真正想做的事情，这就是一个事件委托的流程。

`事件委托`非常适合列表相关的事件处理，假设有成千上万条的列表，这个时候每个列表的操作按钮都要绑定事件，这个消耗是非常巨大的，当列表增减还需要考虑给新列表绑定事件，给删除的列表注销事件，这个时候使用事件委托，只需要在列表之外的一个节点上绑定一个事件，其好处不言而喻。

```javascript
<style>
  .list .item {display: flex;justify-content: space-between;border-bottom: 1px dashed #4caf50;padding: 8px 0;}
  .list .item .caption {font-weight: 700;}
  .list .item .operates .delete {border: 1px solid rgb(177, 107, 107);color: rgb(207, 72, 72);outline: none;cursor: pointer;}

  .add { border: 1px dashed #4caf50; font-size: 14px; padding: 4px 8px; margin-top: 22px; outline: none; cursor: pointer; } .add:active { color: white; background: #4caf50; }
</style>

<div class="list">
  <div class="item">
    <div class="content caption">
      今天要做的事情
    </div>
    <div class="operates caption">
      操作
    </div>
  </div>

  <div class="item">
    <div class="content">
      吃火锅
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>

  <div class="item">
    <div class="content">
      和小姐姐聊天
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>
</div>

<button class="add">增加一项</button>

<script>
  var listEle = document.querySelector('.list');
  var deleteEle = document.querySelectorAll('.delete');

  deleteEle.forEach(function(el) {
    el.addEventListener('click', function() {
      console.log('开始删除...');

      el.setAttribute('disabled', 'disabled');
      el.style.color = 'rgb(226, 174, 174)';
      el.style.borderColor = 'rgb(226, 174, 174)';
      el.style.cursor = 'wait';
      el.innerHTML = '处理中...';

      setTimeout(function() {
        var itemEl = el.parentNode.parentNode;

        listEle.removeChild(itemEl);

        console.log('删除成功');
      }, 1500);
    });
  });

  document.querySelector('.add').addEventListener('click', function() {
    var el = document.createElement('div');

    el.className = 'item';

    el.innerHTML = [
      '<div class="content">',
        '学习',
      '</div>',
      '<div class="operates">',
        '<button class="delete">删除</button>',
      '</div>',
    ].join('');

    listEle.appendChild(el);
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6d550a518d1018280756.jpg)

稍微改写一下之前的例子，不采用事件委托的方式，这个列表中新增的项点击删除按钮是无用的，将这个例子改成事件委托的方式：

```javascript
<style>
  .list .item {display: flex;justify-content: space-between;border-bottom: 1px dashed #4caf50;padding: 8px 0;}
  .list .item .caption {font-weight: 700;}
  .list .item .operates .delete {border: 1px solid rgb(177, 107, 107);color: rgb(207, 72, 72);outline: none;cursor: pointer;}

  .add { border: 1px dashed #4caf50; font-size: 14px; padding: 4px 8px; margin-top: 22px; outline: none; cursor: pointer; } .add:active { color: white; background: #4caf50; }
</style>

<div class="list">
  <div class="item">
    <div class="content caption">
      今天要做的事情
    </div>
    <div class="operates caption">
      操作
    </div>
  </div>

  <div class="item">
    <div class="content">
      吃火锅
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>

  <div class="item">
    <div class="content">
      和小姐姐聊天
    </div>
    <div class="operates">
      <button class="delete">删除</button>
    </div>
  </div>
</div>

<button class="add">增加一项</button>

<script>
  var listEle = document.querySelector('.list');

  listEle.addEventListener('click', function(e) {
    if (e.target.className === 'delete') {
      var el = e.target;

      console.log('开始删除...');

      el.setAttribute('disabled', 'disabled');
      el.style.color = 'rgb(226, 174, 174)';
      el.style.borderColor = 'rgb(226, 174, 174)';
      el.style.cursor = 'wait';
      el.innerHTML = '处理中...';

      setTimeout(function() {
        var itemEl = el.parentNode.parentNode;

        listEle.removeChild(itemEl);

        console.log('删除成功');
      }, 1500);
    }
  });

  document.querySelector('.add').addEventListener('click', function() {
    var el = document.createElement('div');

    el.className = 'item';

    el.innerHTML = [
      '<div class="content">',
        '学习',
      '</div>',
      '<div class="operates">',
        '<button class="delete">删除</button>',
      '</div>',
    ].join('');

    listEle.appendChild(el);
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6d910a590f1818280756.jpg)

新增的项目是不需要再重新绑定事件的。

## 3. 事件节流

事件节流用于控制事件触发的最小间隔。

如一个事件 100 毫秒内只能触发一次。

如窗口缩放过程中对页面的元素大小重新调整，因为 `resize` 事件的触发是非常快的，用户虽然在频繁的变更窗口尺寸，但用户单位时间内能感知到的事情是有限的，也许一秒内执行 100 次尺寸计算和一秒钟内执行 10 次尺寸计算，感知上是没有太大区别的，而且事件内有太多的操作，又在频繁触发事件，这样很容易造成浏览器的卡顿。

```javascript
<style>
  .outer { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background:rgb(0, 0, 0); }
  .outer .text { position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); color: white; font-size: 100px; text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #FF1177, 0 0 35px #FF1177, 0 0 40px #FF1177, 0 0 50px #FF1177, 0 0 75px #FF1177; }
</style>

<div class="outer">
  <div class="text">100x200</div>
</div>

<script>
  var text = document.querySelector('.text');

  var resize = function() {
    var height = window.innerHeight;
    var width = window.innerWidth;

    text.innerText = width + 'x' + height;
  };

  window.addEventListener('resize', resize);

  resize();
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6dcc0ac51be917400762.jpg)

可以看到，resize 事件的响应是非常快的，与之类似的还有 `scroll` 事件，即滚动条滚动时触发的事件。

这种情况下就可以使用节流的方式来优化事件。

```javascript
<style>
  .outer { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background:rgb(0, 0, 0); }
  .outer .text { position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); color: white; font-size: 100px; text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #FF1177, 0 0 35px #FF1177, 0 0 40px #FF1177, 0 0 50px #FF1177, 0 0 75px #FF1177; }
</style>

<div class="outer">
  <div class="text"></div>
</div>

<script>
  var text = document.querySelector('.text');
  var timer = false;

  var resize = function() {
    if (timer) return; // 判断是不是上一次事件执行完300毫秒内

    var height = window.innerHeight;
    var width = window.innerWidth;

    text.innerText = width + 'x' + height;

    timer = setTimeout(function () {
      timer = null;
    }, 300);
  };

  window.addEventListener('resize', resize);

  resize();
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6f060a7733fb12180533.jpg)

对例子做了一个简单的修改，增加了 `timer` 变量，用于存放定时器的标志值（定时器的返回值），每当事件触发时，给 timer 赋值，这个时候事件就会处于一个锁住的状态，直到 300 毫秒后，timer 再次被设置为 null，表示可以触发事件。

根据需求，业务逻辑执行的时机是在定时器内还是定时器外可以自由调配。

## 4. 事件防抖

例如设定间隔时间为 200 毫秒，防抖则是在事件触发后 200 毫秒再执行事件处理器。

假设在这 200 毫秒内又触发了相同事件，则取消上一次的事件，不执行事件处理器，以最后一次触发事件的时机开始，等待 200 毫秒执行事件处理器。

这种情况常用于键盘输入，如表单验证，关键词联想。

搜索引擎基本都含有关键词联想功能，即在输入关键词的时候，根据关键词联想相关内容，为用户更好的命中搜索结果。

```javascript
<style>
  input { border: 1px solid #eee; padding: 4px 8px; min-width: 300px; font-size: 14px; height: 40px; display: block; margin: 0 auto; outline: none; }

  .result { text-align: center; }
</style>

<div>
  <input type="text">

  <div class="result"></div>
</div>

<script>
  var input = document.querySelector('input');
  var result = document.querySelector('.result');

  input.addEventListener('input', function(e) {
    var val = e.target.value;

    result.innerText = '联想内容：' + val;
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6e6d0a2c8e8917400432.jpg)

假设输出的就是服务端返回的联想内容。

其实可以发现，用户在输入关键词的时候，基本是按照词组的方式进行的，在某一个词组输入完成前去获取联想内容其实意义不大。

如搜索官方网站，用户可能会先连续打出``，而对于事件而言，就触发了许多次，通过防抖就可以剔除这些无意义的事件触发。

```javascript
<style>
  input { border: 1px solid #eee; padding: 4px 8px; min-width: 300px; font-size: 14px; height: 40px; display: block; margin: 0 auto; outline: none; }

  .result { text-align: center; }
</style>

<div>
  <input type="text">

  <div class="result"></div>
</div>

<script>
  var input = document.querySelector('input');
  var result = document.querySelector('.result');
  var timer = null;

  input.addEventListener('input', function(e) {
    clearTimeout(timer);

    timer = setTimeout(function() {
      var val = e.target.value;

      result.innerText = '联想内容：' + val;
    }, 300);
  });
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e9c6e8d0a24680717400432.jpg)

通过定时器，来延迟执行事件处理器，每次触发事件，就取消上一次事件处理器，这样就达到了防抖的效果。

## 5. 异步加载的事件处理器

这个方案目前使用的比较少，其就是在事件被触发的时候，去加载远端的事件处理器，加载完毕后再执行事件处理器。

以前因为缺少模块化规范，基本看不到这种优化方案，现在因为新标准`动态import`的出现，使其非常容易融合进业务代码中。

目前有许多构建工具支持动态的 `import` ，利用构建工具可以非常简单的实现异步加载事件处理器。

```javascript
// 这是一份伪代码

const el = document.querySelector('.delete');

el.addEventListener('click', async () => {
  try {
    const event = await import('./event/delete.js');

    // ...
  } catch (e) {
    // ...
  }
});
```

这么做其实优化的并不是事件本身，主要是为了减少首屏加载的代码体积。

## 6. 小结

事件的优化不一定要只从代码方面入手，还可以从其他方面，如交互上进行思考。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
