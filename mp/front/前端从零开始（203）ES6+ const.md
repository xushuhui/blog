# ES6+ const

## 1. 前言

[上一节](https://www.imooc.com/wiki/ES6lesson/let.html)我们学习了使用 `let` 取代 `var` 声明变量，但是很多情况下，我们希望我们声明的变量不能被修改。在 ES5 中不能直接声明一个常量，如果想声明一个不可修改的变量需要借助 `defineProperty` 方法。ES6 为了弥补这方面的缺失，新增了 `const` 语句用于声明一个常量。本节我们还将学习到 `let`、 `const` 、 `var` 的区别。

## 2. 语句使用

### 2.1 基本使用

`const` 的使用类似于 `let`，不同的是 `const` 在声明时必须初始化一个值，而且这个值不能被改变。

```javascript
const PI = 3.1415;  // 定义一个圆周率常量 PI
```

上面的代码声明了一个常量 PI，如果声明时没有初始化值时，则会抛出异常。

```javascript
const PI;
// Uncaught SyntaxError: Missing initializer in const declaration
```

`const` 语句声明的是一个常量，并且这个常量是不能被更改的：

```javascript
const PI = 3.1415;  // 定义一个圆周率常量 PI
PI = 12             // Uncaught TypeError: Assignment to constant variable.
```

这里声明了一个圆周率常量，我们知道圆周率是固定的不会被改变的，如果对 PI 重新赋值，则会抛出不能给常量分配变量的错误。

但如果使用 `const` 声明的变量是一个对象类型的话，我们可以改变对象里的值，这是因为 `const` 存的变量只是一个地址的引用，所以只要不改变引用的值，就不会报错。如下面的例子：

```javascript
const obj = {};
obj.a = 12          // 12
const arr = [];
arr.push(12);       // 12
arr = {};           // Uncaught TypeError: Assignment to constant variable.
```

使用 `const` 声明了一个对象和一个数组，然后增加对象的属性和增加数组的值都不会报错，这是因为我们没有改变 obj 和 arr 的引用。如果对 arr 进行重赋值，则会报不能给常量分配变量的错误。

### 2.2 const 和 let 共有特性

由于 `const` 和 `let` 都是声明变量使用的，所以他们的使用方法基本相同，下面总结一下它们共有的内容，详情参考[上一节](https://www.imooc.com/wiki/ES6lesson/let.html)的 `let` 的使用：

1. 不能进行变量提升，在声明前不能被使用，否则会抛出异常；
2. 存在暂时性死区，在块中不能被重复声明。

## 3. ES5 模拟实现 const

在 ES6 之前是不能定义常量的，如果想定义常量的话，需要借助 ES5 中的 `defineProperty` 方法，这里我们写个示例：

```javascript
function setConst(key, value, obj) {
  Object.defineProperty(window, key, {
    get: function(){
      return value;
    },
    set: function(){
      console.error('Uncaught TypeError: Assignment to constant variable');
    },
  });
}
setConst('PI', 3.1415);
console.log(PI)     // 3.1415
PI = 3;             // Uncaught TypeError: Assignment to constant variable.
```

上面的代码是一个定义常量的函数，使用了 ES5 的 `Object.defineProperty` 方法，这个方法允许在一个对象上定义一个新的属性，具体使用详情可以参考 ES5 的相关文档说明。这里我们会在 window 对象上添加属性，也可以自己定义一个对象进行添加，可以实现局部作用域的效果。通过向 `setConst` 方法中传入指定的变量和值来声明一个常量，这样我们就在 ES5 中实现了常量的概念。由此可见，ES6 的 `const` 带来的好处。

## 4. 场景实例

### 4.1 let 及 const 常见问题

在工作中经常会遇到 `var`、`let` 及 `const` 以下几个问题：

* 什么是变量提升？
* 什么是暂时性死区？
* `var`、`let` 及 `const` 区别？

这些问题在上面的讲解中都有提到过，这里我们总结一下：

### 4.2 什么是变量提升？

变量还没有被声明，但是我们却可以使用这个未被声明的变量，这种情况就叫做提升，并且提升的是声明。

```javascript
console.log(a); // undefined
var a = 1
```

这个代码其实可以写出下面这样的方式：

```javascript
var a;
console.log(a);  // undefined
a = 1
```

其实变量提升就是，把变量名统一地提升到作用域的顶部进行率先定义，这也就是变量提升。不仅变量可以被提升，函数也可以被提升，并且函数的提升要优于变量的提升，函数提升会把整个函数挪到作用域顶部。

### 4.3 什么是暂时性死区？

暂时性死区主要是针对 let 和 const 而言的，因为它们不存在变量提升，所以在它们声明变量之前是不能使用的，这个时候如果使用了就会报错，这时候就形成了暂时性的死区，也就是不能被引用。

```javascript
{
  console.log(name);  // ReferenceError: name is not defined.
  let num = 100;
}
```

定义前被引用则会抛出异常。

### 4.4 var、let 及 const 区别？

上面两个问题解决了，再看它们的区别其实就是显而易见的，主要从以下几个方面来分析它们之间的区别：

* var 声明的变量是全局作用域下的，会污染全局变量；let、const 可以和 {} 连用，产生块作用域不会污染全局；
* var 存在提升，我们能在声明之前使用。let、const 因为暂时性死区的原因，不能在声明前使用；
* var 在同一作用域下，可以重复声明变量；let、const 不能重复声明变量，否则会报错；
* let 和 const 的用法基本一致，但是 const 声明的变量不能再次赋值。

### 4.5 实例

下面的实例主要考察作用域的问题，下面的代码输出的结果是什么？

```javascript
for (var i = 0; i < 3; i++) {
  setTimeout(() => {
    console.log(i);
  }, 1000*i);
}
// 3
// 3
// 3
```

**代码分析：** 这里由于 setTimeout 是异步回调函数，所以 for 循环运行完后才会执行 setTimeout 内的调用栈。使用 var 声明的变量是全局作用域的，循环完后 i 的值是 3，所以会间隔 1s 打印一个 3。想要 i 的值不被覆盖，这时可以借助 let 的块级作用域的特性，来解决这个问题：

```javascript
for (let i = 0; i < 3; i++) {
  setTimeout(() => {
    console.log(i);
  }, 1000 * i);
}
// 0
// 1
// 2
```

**代码分析：** 这里循环的变量 i 是 let 声明的，当前的 i 只在本轮循环有效，所以每一次循环的 i 其实都是一个新的变量。你可能会问，如果每一轮循环的变量 i 都是重新声明的，那它怎么知道上一轮循环的值，从而计算出本轮循环的值？这是因为在 JavaScript 引擎内部会记住上一轮循环的值，初始化本轮的变量 i 时，就在上一轮循环的基础上进行计算。

另外，for 循环还有一个特别之处，就是设置循环变量的那部分是一个父作用域，而循环体内部是一个单独的子作用域。这样每次定义的 i 都是局部作用域下的变量，所以在异步之后，i 的值是不会变的，所以依次打印 0 到 3 的结果。

## 5. 小结

本节我们学习了使用 `const` 来声明一个常量，这里需要注意以下几点：

1. 对于不可变的变量，尽可能地使用 `const` 来声明变量，如果需要更改值的时候再用 `let` 声明；
2. `let` 和 `const` 都只作用于块级作用域内；
3. `let` 和 `const` 声明的变量，都不能变量提升，都存在暂存死区；
4. `let` 和 `const` 声明的变量不允许重复声明，无论重复用 `var` 或者其他声明都不行。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
