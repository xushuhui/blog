# 作用域

作用域即代码片段的有效范围，这里的代码片段可以是一个函数、一个变量等。

在 JavaScript 中，通常被拿来讨论的是 `全局作用域` 和 `函数作用域`。

## 1. 全局作用域

在全局环境下定义的变量、函数，都属于全局作用域的范围，也就是这些变量、函数在任何地方都能被访问。

```javascript
var number = 1;

var fn = function() {
  console.log('我是一个全局下的函数');
  console.log('访问全局下的 number: ', number);
};

fn();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee4718e0925603808600322.jpg)

全局的作用域很好理解，即全局下的变量、函数在任何地方都能被访问到。

> 在 ES6 的模块化规范中，一个模块就是一个 js 文件，所以如果在 ES6 的模块的最顶层声明变量和函数，是不属于全局的。

## 2. 函数作用域

函数拥有自己的作用域，也就是说函数内声明的变量和函数，只在函数内有效。

```javascript
var fn = function() {
  var innerFn = function() {
    console.log('我是函数内的函数');
  };

  var str = '我是函数内的变量';

  innerFn();
  console.log(str);
};

fn();

console.log(str); // 输出：ReferenceError: str is not defined
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee4719b09882f6714280548.jpg)

函数内的变量 `str` 在函数外部无法访问到，因为其所在的作用域是函数 `fn` 的作用域，所以只在 `fn` 函数内能被访问。

## 3. eval 的作用域

eval 根据调用的方式，其作用域也会发生变化。

如果直接调用 `eval` 则其作用域就是当前上下文中的作用域。如果间接性质的调用，比如通过一个 `eval` 的引用来调用，那作用域就是全局的。

```javascript
var storeEval = eval;

(function() {
  storeEval('var number1 = 1;');
  eval('var number2 = 2');

  console.log('自执行匿名函数内输出：', number2);
})();

console.log(number1); // 输出：1
console.log(number2); // 输出：ReferenceError: number2 is not defined
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee471af09fd7bf112620466.jpg)

本身 eval 用到的情况就少，所以这种情况下做个了解即可。

## 4. 作用域链

理解了作用域，作用域链就很好理解了。

通常情况下，内层作用域拥有访问上层作用域的能力，而外层无法访问到内层的作用域。

```javascript
var number = 1;

var fn = function() {
  console.log(number);

  var str = '字符串';
};

fn();

console.log(str); // 输出：ReferenceError: str is not defined
```

由此可见作用域从内往外的。

```javascript
var number1 = 1;
var fn1 = function() {
  var number2 = 2;

  var fn2 = function() {
    var number3 = 3;

    var fn3 = function() {
      console.log('fn3 内的输出：');
      console.log(number1);
      console.log(number2);
      console.log(number3);
    }

    fn3();
  };

  fn2();
}

fn1();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee471ca09ae9f0b11740758.jpg)

例子中的 `fn3` 就具有访问 `fn2作用域`、`fn1作用域`、`全局作用域`的能力。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee471d209e69ed119120730.jpg)

这样从内往外就形成了一条作用域链。

## 5. 利用函数作用域进行封装

函数作用域最常用的场景之一就是隔离作用域。

因为函数有自己的作用域，所以很多库、框架在实现的时候都会把内容写在一个函数中。

```javascript
(function() {
  var config = {};

  var fn = function() {
    // ...
  };

  window.$ = fn;
  window.jQuery = fn;
})();
```

这样就不会污染到全局，只对外暴露想要暴露的部分。

## 6. 小结

有关作用域有更深入的内容，本篇探讨的是最容易理解的部分。

理解作用域可以更好的组织代码结构，减少各个上下文的污染。

在 ES6 中引入了块及作用域的概念，这是在之前都没有的，可以查阅 ES6 中对应的内容进行了解。



### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
