# 变量提升

变量提升是 `JavaScript` 在运行时的一种机制。

在代码被执行前，JavaScript 会做一些准备工作，其中会准备一个执行上下文，也就是代码的执行时的环境，如 绑定`this`、准备变量等。

变量提升这一特性就是在准备执行上下文时进行的，这一特性也是和执行上下文相关的最常在面试中出现的内容。

## 1. 表现

```javascript
console.log(number); // 输出：undefined

var number = 1;

console.log(number); // 输出：1
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eedc32809dcaafb10180232.jpg)

这段代码先输出了 `undefined` 再输出了 `1`。

说明 `number` 变量在第一行输出之前就已经存在了，只不过没有被赋值，因为如果变量不存在，访问变量会抛出异常：`ReferenceError: 变量 is not defined`。

可是在第一次使用 `number` 之前并没有声明过变量，却可以被访问到，出现这个表现就是因为变量提升的特性。

在生成执行上下文阶段，会把变量都收集起来，事先进行声明，需要注意的是，这个特性**只会声明变量，而不会初始化，即变量的值都会是 `undefined`**。

根据这个规则，上面这段代码在执行时可以理解成是这样的：

```javascript
var number;

console.log(number);

number = 1;

console.log(number);
```

需要注意的是，如果没有使用 `var` 关键字声明变量，是不会进行提升处理的。

```javascript
console.log(number);

number = 1;

console.log(number);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eedc339090cfa3f10640226.jpg)

这样就会抛出变量不存在的异常。

## 2. 函数提升

函数也会提升，函数的提升会把整个函数放到最前。

```javascript
fn('咕咕咕');

function fn(str) {
  console.log(str);
}
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eedc34509ec849810260200.jpg)

这段代码可以被正常执行，函数也能被正常调用，因为在生成执行上下文阶段，整个函数会提升到最前面。

这个规则只对函数声明的方式声明的函数有效，如果使用的是函数表达式，那就会走变量提升的规则。

```javascript
console.log(fn); // 输出：undefined
fn('咕咕咕'); // 抛出异常 TypeError: fn is not a function

var fn = function(str) {
  console.log(str);
};
```

可以看到 `fn` 能被访问到，已经声明了，但不能作为函数调用，这说明 `fn` 走了变量提升的机制。

在执行上下文生成的阶段，函数会比变量更早的进行提升，也就是说函数相比变量，更加靠前。

**函数在调用时也会生成函数级别的执行上下文，这也就意味着提升这个特性也会在函数执行前发生**。

## 3. 小结

> 在 ES6 中和提升相关的内容又有些许不同，let 和 const 这两个新关键字对提升的表现也和 var 不同，具体可以参阅 ES6 中的相关内容。

现在和提升相关的内容更多的出现在面试题里，由于代码检查工具的介入，一些由于提升特性造成的 `BUG` 出现的越来越少。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
