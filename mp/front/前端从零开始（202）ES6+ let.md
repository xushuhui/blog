# ES6+ let

## 1. 前言

本节我们一起学习下 ES6 中的 `let`，在 ES5 中变量的方法只有一个 `var` ，但是使用 `var` 来定义的变量存在很多缺陷和弊端，ES6 引入了 `let` 语句来声明变量，同时引入了很多概念，比如块级作用域、暂存死区等等。限制了任意声明变量，提升了程序的健壮性。

## 2. 基本用法

`let` 的使用方法类似于 `var`，并可以代替 `var` 来声明变量。

```javascript
{
  let name = 'imooc';
}
```

`let` 允许你声明一个作用域被限制在块级中的变量、语句或者表达式。与 `var` 关键字不同的是，`var` 声明的变量只能在全局或者整个函数块中。 `var` 和 `let` 的不同之处在于 `let` 是在编译时才初始化，也就是在同一个块级下不能重复地声明一个变量，否则会报错。

`let` 不会在全局声明时创建 window 对象的属性，但是 `var` 会。

```javascript
{
  var name = 'imooc'  // imooc
  var name = 'iimooc' // iimooc
}
console.log(window.name)	// iimooc
{
  let age = 10        // 10
  let age = 18        // Uncaught SyntaxError: Identifier 'age' has already been declared
}
console.log(window.age)		// undefined
```

上面的代码中，在一个块中分别使用 `var` 和 `let` 来声明变量对比他们之间的差异，从上面的代码操作可以看出，我们可以使用 `var` 多次对 name 声明，但是使用 `let` 声明的 age，后面再使用 `let` 对其声明是会报错的。

`var` 是没有块的概念的，声明的变量会是 window 对象上的属性，在最外层的 window 上可以取到。而 `let` 存在块的概念，不会添加到 window 对象上，这些是 `let` 和 var 之间的区别。从这里我们可以了解到为什么使用 `let`。

## 3. 块级作用域

在深入了解 let 前，我们需要了解一下，在 JavaScript 中有哪些作用域：

1. 全局作用域
2. 函数作用域 / 局部作用域
3. 块级作用域

上面是 JavaScript 中的三种作用域，那什么是作用域呢？首先要明白的是：几乎所有的编程语言都存在在变量中储值的能力，存储完就需要使用这些值。所以，作用域就是一套规则，按照这套规则可以方便地去存储和访问变量。

在 ES5 中的作用域有全局作用域和函数作用域，而块级作用域是 ES6 的概念。

### 3.1 全局作用域

全局作用域顾名思义，就是在任何地方都能访问到它，在浏览器中能通过 window 对象拿到的变量就是全局作用域下声明的变量。

```javascript
var name = 'imooc';
console.log(window.name)   // imooc
```

使用 `var` 定义的变量，可以在 window 对象上拿到此变量。这里的 name 就是全局作用域下的变量。

### 3.2 函数作用域

函数作用域就是在函数内部定义的变量，也就是局部作用域，在函数的外部是不能使用这个变量的，也就是对外是封闭的，从外层是无法直接访问函数内部的作用域的。

```javascript
function bar() {
  var name = 'imooc';
}
console.log(name);  // undefined
```

在函数内部定义的 name 变量，在函数外部是访问不了的。要想在函数外部访问函数内部的变量可以通过 return 的方式返回出来。

```javascript
function bar(value) {
  var name = ' imooc';
  return value + name;
}
console.log(bar('hello'));  // hello imooc
```

借助 return 执行函数 bar 可以取到函数内部的变量 name 的值进行使用。

### 3.3 块级作用域

块级作用域是 ES6 的概念，它的产生是要有一定的条件的，在大括号（`{}`）中，使用 `let` 或 `const` 声明的变量，才会产生块级作用域。

这里需要注意的是，块级作用域的产生是 `let` 或 `const` 带来的，而不是大括号，大括号的作用是限制 `let` 或 `const` 的作用域范围。当不在大括号中声明时， `let` 或 `const` 的作用域范围是全局。

```javascript
let name = 10;
console.log(window.name)   // undefined
```

上面的代码可以看到，使用 `let` 方式声明的变量在 window 下是取不到的。

```javascript
var num = 10;
{
  var num = 20;
  console.log(num)  // 20
}
console.log(num)    // 20
```

在使用 `var` 声明的情况下，可以看出，外层的 num 会被 {} 中的 num 覆盖，所以没有块级作用域的概念，下面看下使用 `let` 方式声明：

```javascript
let num = 10;
{
  console.log(num); // Uncaught ReferenceError: Cannot access 'num' before initialization
  let num = 20;
  console.log(num)  // 20
}
console.log(num)    // 10
```

这里可以看出 {} 内外是互不干涉和影响的，如果在声明 num 的前面进行打印的话，还会报错，这个时候，num 处于暂存死区，是不能被使用的，下面我们会具体说明。

在低版本浏览器中不支持 ES6 语法，通常需要把 ES6 语法转换成 ES5，使用 babel 把上面的代码转换后得到如下结果：

```javascript
var num = 10;
{
  console.log(_num); // num is not defined
  var _num = 20;
  console.log(_num); // 20
}
console.log(num);    // 10
```

从上面的代码中可以看出，虽然在 ES6 语法使用的是相同的变量名字，但是底层 JS 进行编译时会认为他们是不同的变量。也就是说即使大括号中声明的变量和外面的变量是相同的名字，但是在编译过程它们是没有关系的。

块级作用域可以任意嵌套，如下实例：

```javascript
{{
  let x = 'Hello imooc'
  {
    console.log(x); // Hello imooc
    let y = 'Hello World'
  }
  console.log(y);   // 引用错误 ReferenceError: y is not defined.
}};
```

上方代码每一层都是一个单独的作用域，内层作用域可以读取外层的变量，所以第一个会打印 `Hello imooc`, 而外层无法读取内层的变量，所以会报错。

## 4. 不能变量提升

对应 `var` 我们知道可以变量提升的，提升到作用域的最顶部，作用域是全局，使得声明变量前也可以使用，但值为 `undefined`。

```javascript
{
  console.log(bar); // 输出undefined，没有值但不会报错
  var bar = 1;
}
```

一般变量都应该先声明再使用，所以 `let` 和 `const` 语法规定必须声明后使用，否则报错。

```javascript
{
  console.log(name); // 引用错误 ReferenceError: name is not defined.
  let name = 'imooc';
}
```

上面代码中，都是在声明前的时候使用变量的，这时候由于 `let` 不能进行变量提升所以会报引用错误。

## 5. 暂时性死区

上面讲到在变量声明前使用这个变量，就会报错。在代码块内，使用 `let` 命令声明变量之前，该变量都是不可用的。这在语法上，称为 “暂时性死区”（temporal dead zone，简称 TDZ）。

```javascript
{
  console.log(name);  // ReferenceError: name is not defined.
  let name = 'imooc';
  console.log(name);  // imooc
}
```

上面代码中，在块中使用了 `let` 声明了 name 变量，在使用前对 name 进行了声明，name 则会处于暂存死区，不能被使用，如果引用则会引用错误。

> **Tips**：注意对于 `typeof` 也是会报错的。

```javascript
{
  console.log(typeof name);  // Uncaught ReferenceError: Cannot access 'name' before initialization
  let name = 'imooc';
}
```

上面的代码中，name 引用错误：无法在初始化之前访问 name，因为 name 在这个块的下面进行了声明，name 就是一个死区，不能被引用了。因此，`typeof` 运行时就会抛出一个 `ReferenceError` 的参数错误。

## 6. 重复声明报错

`let` 不允许在同一个函数或块作用域中重复声明同一个变量，否则会引起语法错误（SyntaxError）。

```javascript
{
  let x = 10;
  let x = 11;
}
// Uncaught SyntaxError: Identifier 'x' has already been declared
```

在上面的代码中报错，所以，同一个变量名不可以在同一个作用域内重复声明。

```javascript
{
  let x = 10;
  var x = 1;
}
```

即使使用 `var` 去声明也是不可以的，我们知道当使用 `let` 声明的时候 x 已经是一个死区了，不可以被重复声明了。

> **Tips**：注意在 `switch` 语句中只有一个块级作用域，所以下面这种情况也是会报错的。

```javascript
let x = 1;
switch(x) {
  case 0:
    let num;
    break;

  case 1:
    let num;//重复声明了
    break;
}
// 报错
```

如果把 `case` 后面的语句放到块作用域中则不会报错。

```javascript
let x = 1;
switch(x) {
  case 0: {//块
    let num;
    break;
  }
  case 1: {//块
    let num;//这里就没有关系了，可以正常声明
    break;
  }
}
```

上方代码，`case` 后面的语句 `let` 变量声明在放到块中，是单独的作用域，所以就不会报错。

## 7. 小结

本节讲解了 let 语句的使用，还有作用域的概念，需要注意以下几点：

1. `let` 只作用于块级作用域内；
2. `let` 声明的变量不能进行变量提升，存在暂存死区；
3. `let` 声明的变量不允许重复声明，无论重复用 `var` 或者其他声明都不行；
4. 尽量使用 `let` 去代替 `var` 来声明变量。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
