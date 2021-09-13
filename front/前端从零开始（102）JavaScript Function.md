# JavaScript Function

> Function 构造函数创建一个新的 Function 对象。直接调用此构造函数可用动态创建函数，但会遭遇来自 eval 的安全问题和相对较小的性能问题。—— MDN

`Function` 可以用来创建函数，JavaScript 中的所有函数，都是 `Function对象`。

## 1. 使用 Function 创建函数

Function 在被当作构造函数调用的时候，可以用来创建函数。

```javascript
var fn = new Function(函数参数1, 函数参数2, ..., 函数参数n, 函数体);
```

函数体是一个字符串，字符串的内容是就是函数调用时候被执行的语句。

```javascript
var fn = new Function('a', 'b', 'return a + b');

var result = fn(1, 3);

console.log(result); // 输出：4
```

![图片描述](https://img.mukewang.com/wiki/5e7a45b80aace13814600432.jpg)

将上面创建函数的方式“翻译”成常用的创建方式可以是：

```javascript
var fn = function(a, b) {
  return a + b;
};

var result = fn(1, 3);

console.log(result); // 输出：4
```

对比一下就很好理解使用 `new Function` 时候所传递的参数的含义了。

## 2. 与常规方式创建函数的区别

这里指的常规方式是指`函数声明`、`函数表达式`或 ES6 中的`箭头函数`。

使用 Function 创建的函数，最后一个参数，即函数体内在执行的时候作用域是在全局的。

```javascript
var num = 20;

function fn() {
  var num = 10;

  var func = new Function('console.log(num)');

  console.log(num);

  func();
}

fn();

// 输出：
//  10
//  20
```

这个例子在执行后，依次输出了 `10` ，`20` ，根据结果可以知道 `new Function` 创建的函数，在执行过程中，上层作用域是顶级的全局作用域，在浏览器下即为 `window` 。

## 3. 使用场景

使用 Function 来创建函数是比较麻烦的，照道理讲不会有小伙伴会喜欢用这种方式创建函数。

Function 的使用主要场景与 `eval` 类似，用于动态的运行代码。

如在线的代码解析器就可以配合 Function 完成。

```javascript
var run = function(code, callback) {
  window._callback = callback;

  var fn = Function(code + ';_callback()');

  fn();
};


run('console.log("动态执行的代码");', function() {
  console.log('代码执行后做的事');
});
```

这样就可以实现简单的动态运行代码。

> 注意：真正需要完成这个功能需要大量的细节处理，如处理输出，处理异步，绝非这么简单。

还有一些代码编译工具会将编译后的代码，使用 `new Function` 进行包裹，如以下代码：

```javascript
var number = 1;
var flag = false;

console.log(number, flag);
```

上面这份代码在经过编译后可能会变成：

```javascript
(new Function('console.log(1,!1)'))()
```

这样做是为了缩短代码，另外就是让格式化工具无法很好的格式化代码。

## 4. 小结

开发者通常不会通过内建对象 `Function` 来创建函数，更多的是利用 `Function` 的特性来动态执行代码。

通常情况下 `Function` 创建的函数，在执行过程中其 this 是指向最顶层的。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
