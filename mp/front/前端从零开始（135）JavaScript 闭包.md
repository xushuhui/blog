# 闭包

> 函数和对其周围状态（lexical environment，词法环境）的引用捆绑在一起构成闭包（closure）。也就是说，闭包可以让你从内部函数访问外部函数作用域。在 JavaScript 中，每当函数被创建，就会在函数生成时生成闭包。

**由于闭包的概念比较抽象，所以本篇幅会有较多的主观理解。**

在作用域相关的内容中可以知道，全局下的作用域想访问一个函数内部的作用域是办不到的，但是 `闭包` 的特性可以突破这一限制。

每个函数都会形成一个闭包。

## 1. 什么是闭包

闭包可以理解成，保留了函数作用域链的一个环境。

```javascript
var fn = function() {
  var number = 0;
};

fn();

console.log(number);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee5c7aa09f26f7411460290.jpg)

这个例子是访问不到 `number` 的，想访问到就可以借助闭包的特性。

```javascript
var fn = function() {
  var number = 0;

  return function() {
    number++;

    console.log(number);
  };
};

var increment = fn();

increment();
increment();
increment();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee5c7b60993310c11720550.jpg)

这里的 `fn` 函数返回了一个函数，在这个返回的函数所形成的闭包环境就拥有访问上一层作用域的能力，所以每次在调用 `fn` 返回的函数时，就可以累加 `number`。

借助一个函数形成的闭包环境作为跳板，来访问另一个函数的作用域，就是闭包最常见的使用场景。

网络上有许多文献把闭包称为`能访问其他函数内部变量的函数`，这样理解可能更容易一些。

函数用到了上层作用域的变量，所以这些变量会在内存中被保留，不会被释放，一些旧的浏览器在内存管理上没有现代浏览器完善，大量的闭包可能会导致页面卡顿，不过通常业务开发，会先考虑效果，再考虑性能。

## 2. 闭包的应用

### 2.1 模拟私有属性

在 `JavaScript` 中是没有私有属性特性的，利用闭包来隐藏变量，就可以模拟出私有属性的效果。

```javascript
var counter = (function() {
  var count = 0;

  return {
    increment: function() {
      count++;
      return count;
    },
    zero: function() {
      count = 0;
      return count;
    },
    get value() {
      return count;
    },
  };
})();

counter.increment();
console.log(counter.value); // 输出：1
counter.increment();
console.log(counter.value); // 输出：2

console.log(counter.count); // 输出：undefined
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee5c7d009a098ca11720818.jpg)

这里的自执行匿名函数返回一个对象，对象中的方法就具有访问上层函数中的变量的能力，所以他们都能访问 count。

因为 count 不会被释放，所以可以当作一个属性来使用。

### 2.2 回调函数几乎都用到了闭包的特性

回调函数通常会用到上层作用域的变量，然后在某一情况下进行调用。

```javascript
var fn = function(cb) {
  console.log('异步操作开始');
  setTimeout(function() {
    console.log('异步操作结束');
    cb();
  }, 1000);
};

var obj = {
  flag: false,
};

fn(function() {
  obj.flag = true;

  console.log(obj);
});
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee5c8130a4340b211720684.jpg)

很明显，这里的回调函数就是用到了闭包的特性。

所以闭包其实很常用，结合日常的这些场景能更好的理解闭包。

## 3. 小结

每个函数都有闭包，闭包可以访问到这个函数所在的上层作用域，利用这一特性，就能访问到一个函数作用域下的变量。

大量的闭包可能会造成性能问题，不过现在的计算机处理器、内存已经让开发者不太需要关注这方面的问题，但在设计一个会被大量应用的库和框架时，应当做这方面的考虑，因为用户的环境千变万化。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
