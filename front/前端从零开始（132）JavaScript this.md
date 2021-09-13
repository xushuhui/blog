# JavaScript this

> 当前执行代码的环境对象，在非严格模式下，总是指向一个对象，在严格模式下可以是任意值。(MDN)

this 指向的是当前的代码上下文环境，所以不同情况下的 this 指向也不同。

## 1. 全局下的 this

在全局环境下，`this` 指向全局对象。

全局对象和宿主环境相关，在浏览器下，全局对象就是 `window` 对象，在 `node.js` 中，全局对象是 `global` 对象。

```javascript
window === this; // 输出：true
```

![图片描述](https://img.mukewang.com/wiki/5edc6469092f6da410140434.jpg)

> 新的标准提供了 `globalThis` 关键字来获取全局对象，这样就能抹平宿主的差异来操作处理全局对象了。

## 2. 函数中的 this

函数在不同情况下，其 `this` 的指向也不同。

### 2.1 对象下的方法

方法也是一个函数，如果通过对象调用一个函数，函数的 `this` 就会指向这个对象。

```javascript
var person = {
  age: 14,
  name: '鸽子王',
  skill: '放鸽子',
  say: function() {
    console.log('来一段自我介绍：');
    console.log('我是' + this.name);
    console.log('我今年' + this.age + '岁');
    console.log('我最擅长' + this.skill);
  },
};

person.say();
```

![图片描述](https://img.mukewang.com/wiki/5edc647809ab787c13780626.jpg)

`say函数`作为对象下的方法，在被调用后，其 `this` 指向的是他所在的对象，在这里就是 `person` 对象。

### 2.2 原型链上方法的 this

原型链上的方法，this 指向的也是调用该方法的对象。

```javascript
var __proto__ = {
  sum: function() {
    return this.number1 + this.number2;
  },
};

var object = Object.create(__proto__);

object.number1 = 1;
object.number2 = 2;

console.log(
  object.sum(),
); // 输出：3
```

![图片描述](https://img.mukewang.com/wiki/5edc653d098ca0f109580440.jpg)

`Object.create` 做就就是将参数作为原型，创建一个对象。

所以 `object` 的第一原型就是 `__proto__` 对象。

`number1` 和 `number2` 都是 `object` 变量的属性，但却可以被 `sum` 方法中的 `this` 访问到，所以在原型链的方法中，this 指向的就是调用该方法的对象。

### 2.3 getter / setter 下的 this

`getter` 和 `setter` 下的 this 也会指向调用该 `getter` 和 `setter` 的对象。

```javascript
var object = {
  _name: '鸽子王',

  get name() {
    return this._name;
  },

  set name(val) {
    console.log(val);
    this._name = val;
  }
};

console.log(object.name); // 输出：鸽子王

object.name = '鸽子天王'; // 输出：鸽子天王

console.log(object.name); // 输出：鸽子天王
```

![图片描述](https://img.mukewang.com/wiki/5edc655909aea5a511620708.jpg)

`getter` 和 `setter` 本质上也可以理解成两个函数，作为对象下的函数，在调用的时候 `this` 也会指向该对象。

### 2.4 作为 DOM 节点的事件处理器

作为 DOM 节点的事件处理器的时，函数的 `this` 会指向这个 DOM 对象。

```javascript
<div>
  <button>点击我</button>
</div>

<script>
  document.querySelector('button').addEventListener('click', function() {
    this.innerHTML = '被点击了！';
  });
</script>
```

![图片描述](https://img.mukewang.com/wiki/5edc65cd0a94d7b408880432.jpg)

### 2.5 作为一个内联的事件处理器

内联的事件处理器，其 `this` 指向的是 DOM 节点自身。

```javascript
<div>
  <button onclick="console.log(this); console.log(this === document.querySelector('button'))">点击我</button>
</div>
```

![图片描述](https://img.mukewang.com/wiki/5edc66190aeaa64318600456.jpg)

这个规则有局限性，只有最外层的 this 符合这个规则。

```javascript
<div>
  <button onclick="function test() { console.log(this) }; test();">点击我</button>
</div>
```

`test` 函数的 this 指向的是全局对象 `window`。

![图片描述](https://img.mukewang.com/wiki/5edc663709781ba820580402.jpg)

### 2.6 其他大部分情况下

排开上述的几个情况，剩下的函数大部分情况下在调用时，this 指向的是全局对象，在浏览器中就是 `window` 对象。

```javascript
function fn() {
  console.log(this);

  console.log(this === window);
}

fn();
```

这样调用函数，其 this 指向的就是 window 对象了。

有的时候可能会搞混以下情况：

```javascript
var object = {
  username: '咸鱼',
  fn: function() {
    console.log(this.username);

    function thisTest() {
      console.log(this.username);

      console.log(this === window);
    }

    thisTest();
  },
};

object.fn();
```

![图片描述](https://img.mukewang.com/wiki/5edc6aef09a30a8809380580.jpg)

这里 `thisTest` 方法输出的 `username` 就会是个 undefined，因为他的 this 指向的是 window，因为他不属于 `object` 对象的一个方法，所以 this 就指向了 window。

在回调函数中经常会碰到这个问题：

```javascript
var info = {
  account: '123',
  password: '456',
  login: function(cb) {
    setTimeout(function() {
      cb({
        account: this.account,
        password: this.password,
      });
    }, 1000);
  }
};

info.login(function(info) {
  console.log(info);
});
```

![图片描述](https://img.mukewang.com/wiki/5edc6b020ab84f5e18680538.jpg)

这里回调函数获取的账号和密码是 `undefined`，原因就是 this 的指向问题。

通常会使用保留上层 this 的方式解决这个问题。

```javascript
var info = {
  account: '123',
  password: '456',
  login: function(cb) {
    var _this = this;

    setTimeout(function() {
      cb({
        account: _this.account,
        password: _this.password,
      });
    }, 1000);
  }
};

info.login(function(info) {
  console.log(info);
});
```

![图片描述](https://img.mukewang.com/wiki/5edc6b110ad1385618680590.jpg)

这样就能解决这个问题。

另外一个情况也很容易混淆 this ：

```javascript
var object = {
  user: 'no.1',
  say: function() {
    console.log(this.user);
  },
};

var say = object.say;

object.say(); // 输出："no.1"
say(); // 输出：undefined
```

![图片描述](https://img.mukewang.com/wiki/5edc6b2f09fc5b4016760414.jpg)

这是因为把 `object` 下的 `say` 方法单独赋值给 say 变量的时候，其就作为了 window 下的一个方法，所以他的 this 指向的是 window。

在严格模式中，这种情况下的 `this` 会变成 `undefined`。

### 2.7 构造函数

在 JavaScript 构造函数也被成为 `对象构造器`，用于产生对象。

构造函数的声明和普通函数几乎没有区别：

```javascript
function Point(x, y) {
  this.x = x;
  this.y = y;
}

var point = new Point(1, 2);

console.log(point.x); // 输出：1
console.log(point.y); // 输出：2
```

构造函数使用 `new` 关键字来构造对象。所以当一个函数被使用 `new` 关键字调用时，这个函数就会作为一个构造函数。

在一个构造函数被调用后，其内部的 `this` 会指向一个对象，具体的内容可以参考 `构造函数` 章节。

## 3. 修改 this

### 3.1 call 方法和 apply 方法

函数具有 `call` 方法和 `apply` 方法，这两个方法可以在调用函数的时候指定函数的 this。

```javascript
var object = {
  user: 'no.1',
};

function say() {
  console.log(this.user);
}

say(); // 输出：undefined
say.call(object); // 输出："no.1"
say.apply(object); // 输出："no.1"
```

![图片描述](https://img.mukewang.com/wiki/5edc6b6709638f4409540464.jpg)

通过 `call` 和 `apply` 方法将 say 函数执行时候的 this 设置为 `object` 对象。

call 方法从第二个参数开始，表示是要传递给当前函数的参数。

```javascript
var object = {
  user: 'no.1',
};

function fn(arg1, arg2, arg3) {
  console.log(
    this,
    arg1,
    arg2,
    arg3,
  );
}

fn.call(object, 1, 2, 3);
```

![图片描述](https://img.mukewang.com/wiki/5f1d191e09fb00ab09060446.jpg)

apply 的第二个参数是个数组，数组里面的项会按数组的顺序作为参数传递给函数。

```javascript
var object = {
  user: 'no.1',
};

function fn() {
  console.log(
    this,
    arguments,
  );
}

fn.apply(object, [1, 2, 3]);
```

![图片描述](https://img.mukewang.com/wiki/5f1d194e091dc2e611400704.jpg)

通过 `arguments` 关键字就可以看到当前函数的参数，通常在需要修改 this ，又不确定参数的情况下，会使用 `apply` 来修改 this。

### 3.2 bind

bind 方法用于给一个函数永久绑定一个指定的 this，bind 不会修改原函数，会返回一个新的函数。

```javascript
var obj1 = { value: '今天打砖' };
var obj2 = { value: '明天打转' };

var fn = function() {
  console.log(this);
};

var bindFn1 = fn.bind(obj1)
var bindFn2 = bindFn1.bind(obj2);

bindFn1();
bindFn2();
```

![图片描述](https://img.mukewang.com/wiki/5edc6b9109601dcb08920442.jpg)

可以看到 `bindFn1` 被绑定了 `obj1` 作为 this，之后不论怎么操作，他的 this 都会是 `obj1`。

> bind 还有更多灵活的用法，参数也可以绑定，有关 bind、call、apply 这三个方法的更详细的信息可以查阅对应的文档。

## 4. 小结

理解好 this 的处理机制可以设计出更加完善的 JavaScript 应用程序。

this 在 ES6 的箭头函数中的表现也有所不同，可以查阅 ES6 中有关箭头函数的内容。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
