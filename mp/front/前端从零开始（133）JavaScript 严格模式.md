# 严格模式

> JavaScript 的严格模式是使用受限制的 JavaScript 的一种方式，从而隐式地退出“草率模式”。严格模式不仅仅是一个子集：这种模式有意地与普通情形下的代码有所区别。(MDN)

严格模式为 `JavaScript` 提供了一个更严格的运行环境。

开启严格模式后，部分特性会发生改变，如 `this` 指向 `window` 的函数不再指向 `window`，而是变成了 `undefined`。

```javascript
function Test() {
  'use strict';
  console.log('this:', this);
}

Test(); // 输出：this: undefined
```

## 1. 开启严格模式

### 1.1 对单个 script 标签或者 js 文件开启严格模式

单个 js 文件或者 script 标签的严格模式，可以通过在所有代码执行前加上 `'use strict'` 字符串开启。

```javascript
'use strict';

function Test() {
  console.log('this:', this);
}

Test(); // 输出：this: undefined
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfba6809d3db4f07460212.jpg)

### 1.2 对一个函数开启严格模式

在函数顶端协商 `'use strict'` 字符串，就可以打开整个函数的严格模式。

```javascript
function testWith() {
  'use strict';

  var person = {
    name: '鸽子天王',
    age: 12,
  };

  var age = 11;

  with (person) {
    console.log(name);
    console.log(age);
  }
}

testWith();
```

在严格模式下是不提供 `with` 语句的调用的，所以这里会爆 `Strict mode code may not include a with statement` 错误。

## 2. 严格模式的规范

### 2.1 禁止使用 with

在严格模式下是禁止使用 with 语句的。

```javascript
'use strict';
var person = {
  name: '鸽子巨星',
};

with (person) {
  console.log(name);
}
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfba5809e5c1df10080300.jpg)

### 2.2 变量必须被声明

在严格模式下，变量必须被声明才能使用，否则会报错。

```javascript
// 非严格模式下
number = 1;

console.log(number); // 输出：1
```

```javascript
// 严格模式下
'use strict';

number = 1; // 报错：ReferenceError: number is not defined
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfba900966596510120172.jpg)

### 2.3 eval 会创建自己的作用域

非严格模式下的 eval 作用域是其本身所在的作用域，而严格模式下，eval 执行过程中会创建一个新的作用域，并在结束后销毁。

```javascript
// 非严格模式下

var number = 1;

eval('var number = 3; console.log(number)'); // 输出：3

console.log(number); // 输出：3
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbaca0931ceba08800232.jpg)

```javascript
// 严格模式下
'use strict';

var number = 1;

eval('var number = 3; console.log(number)'); // 输出：3

console.log(number); // 输出：1
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbad609fd69ee08660286.jpg)

### 2.4 函数的 arguments 不能被修改

在非严格模式下，函数的 `arguments` 可以被重新赋值，在严格模式下，做赋值操作会报错。

```javascript
function fn() {
  console.log(arguments);

  arguments = 1;

  console.log(arguments);
}

fn(1, 2, 3);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbae2095aa8dd09860344.jpg)

```javascript
'use strict';
function fn() {
  console.log(arguments);

  arguments = 1;

  console.log(arguments);
}

fn(1, 2, 3);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbaee094e582b10120340.jpg)

### 2.5 函数的 this 规则有些变化

在 `this` 章节中讨论了不同情况的指向，其中有一种情况函数的 `this` 是指向 window 的。

在严格模式中，这种情况下的 `this` 会变成 `undefined`。

```javascript
function testThis() {
  console.log(this);
}

testThis();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb0109ab669f09400284.jpg)

```javascript
// 严格模式下
'use strict';
function testThis() {
  console.log(this);
}

testThis();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb0c09b39aa709880286.jpg)

## 2.6 caller 与 arguments.callee 被禁用

`arguments.caller` 可以获取到调用当前函数的函数的引用（该属性已经被标准废弃，不再使用了）。

`arguments.callee` 则可以获取到当前函数的引用。

这两个属性在严格模式下都被禁用。

```javascript
function fn1() {
  console.log(arguments.callee === fn1);
}

fn1();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb1a092ace3508920190.jpg)

```javascript
'use strict';
function fn1() {
  console.log(arguments.callee === fn1);
}

fn1();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb2209af302e10120342.jpg)

### 2.7 删除 configurable 为 false 的属性时报错

在非严格模式下，这种情况会直接忽略。

```javascript
var obj = {};

Object.defineProperty(obj, 'prop', {
  configurable: false,
  value: 1,
});

console.log(obj);

delete obj.prop;

console.log(obj);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb2e0937e94608920436.jpg)

```javascript
'use strict';
var obj = {};

Object.defineProperty(obj, 'prop', {
  configurable: false,
  value: 1,
});

console.log(obj);

delete obj.prop;

console.log(obj);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb3a09ba277209520504.jpg)

### 2.8 修改 writable 为 false 的属性时会报错

在非严格模式下，这种情况会直接忽略。

```javascript
var obj = {};

Object.defineProperty(obj, 'prop', {
  writable: false,
  value: 1,
});

console.log(obj.prop);

obj.prop = 2;

console.log(obj.prop);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb5c09b65dae09540424.jpg)

```javascript
'use strict';
var obj = {};

Object.defineProperty(obj, 'prop', {
  writable: false,
  value: 1,
});

console.log(obj.prop);

obj.prop = 2;

console.log(obj.prop);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb3a09ba277209520504.jpg)

### 2.9 禁止书写八进制的字面量

八进制的表示是在数字前面加一个 `0`，但其实 `ECMAScript` 标准中并没有这种表示法。

在 ES6 中提供了八进制数的表示方式，即在数字前加上 `0o` 前缀。

在严格模式下是禁止使用 `0` 前缀表示的八进制字面量的。

```javascript
// 非严格模式中

var num = 010;

console.log(num); // 输出：8
```

```javascript
// 严格模式下

'use strict';
var num = 010;

console.log(num); // 输出：8
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb8509e3ebd609420202.jpg)

### 2.10 新增了一些不能作为变量的关键字

许多关键字在非严格模式下是可以当作变量名的。

```javascript
var yield = 1;

console.log(yield);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbb9409f7ea7209100136.jpg)

```javascript
'use strict';

var yield = 1;

console.log(yield);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5edfbba1091c2a3209040188.jpg)

根据 `MDN` 提供的内容，保留字有：

* implements
* interface
* let
* package
* private
* protected
* public
* static
* yield

## 3. 小结

对严格模式大多数需要注意的是他带来的一些语法上的改变，特别是有些写法，在之前是不会报错的，开启严格模式后就会报错，如果没有进行错误处理，就会导致程序的中断。

以上对严格模式下的改变不一定是全部，可以参考 [MDN](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Strict_mode) 的文档。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
