# for…of

## 1. 前言

在编程中最常见的就是对数据的遍历操作，ES5 有针对数组和对象的遍历方法，但这些方法或多或少地都会存在一定的问题。为了统一解决这些问题，ES6 给出了终极的解决方案 ——`for...of`。

`for...of` 对于可迭代的对象（包括：内置的 String、Array、类似数组的对象（arguments 或 NodeList）、TypedArray、Map、Set，和用户定义的可迭代对象等）上创建一个迭代循环，它不局限于数组下的循环，只要是可迭代的对象就可以被 `for...of` 进行循环。

## 2. 基本语法

### 2.1 语法使用

```javascript
for (const iterator of iterable) {
  // 执行语句
}
```

**参数解释：**

|参数|描述|
|----|----|
|`iterator`|在每次迭代中，将不同属性的值分配给变量，用于循环中语句的使用；|
|`iterable`|被迭代枚举的对象。                                            |

### 2.2 迭代 Array

`for...of` 最常用的场景就是对数组的迭代，也是取代 `for`、`forEach` 的最好选择。

```javascript
let arr = [10, 20, 30];

for (let value of arr) {
    value += 1;
    console.log(value);
}
// 11
// 21
// 31
```

上面的代码中对 value 值进行加 1 操作，如果 value 值不能被修改，也可以使用 `const` 来定义 value。

### 2.3 迭代字符串

`for...of` 可以迭代字符串，迭代后的结果是把字符进行分割，得到每个单个字符。

```javascript
let str = '慕课';

for (let value of str) {
    console.log(value);
}
// 慕
// 课
```

### 2.4 迭代 TypedArray

```javascript
let iterable = new Uint8Array([0x00, 0xff]);

for (let value of iterable) {
  console.log(value);
}
// 0
// 255
```

### 2.5 迭代 Set 和 Map

在 Set 和 Map 章节中我们就说到了，Set 和 Map 可以使用 `for...of` 来进行循环，主要因为 Set 和 Map 具有可迭代属性。

```javascript
let setArr = new Set([1, 1, 2, 2, 3, 3]);

for (let value of setArr) {
  console.log(value);
}
// 1
// 2
// 3
```

上面的代码需要注意的是，迭代的是 `new Set()` 后的结果，`new Set()` 会对数组进行去重操作，所以得到以上结果。

```javascript
let map = new Map([["a", 1], ["b", 2], ["c", 3]]);

for (let value of map) {
  console.log(value);
}
// ["a", 1]
// ["b", 2]
// ["c", 3]
```

上面的代码中使用 `new Map()` 传入一个二维数组，这里需要注意的是，迭代的结果是一个带有 key 和 value 的数组，所以也可以用数组解构的方式 把 key 和 value 的值取出来，直接使用：

```javascript
for (let [key, value] of map) {
  console.log(key, value);
}
// a 1
// b 2
// c 3
```

### 2.6 迭代类数组对象

**1. 迭代 argument 对象**

我们知道在函数中可以使用 `Argument` 对象拿到在调用函数时拿到传递的参数，因为 `arguments` 不是一个 `Array`，它属于类数组，可以借助 `call` 来得到一个数组。`[].slice.call(arguments)`， 而使用 `for...of` 可以直接对 `arguments` 循环，得到的结果也只是传入的参数。这个可以很方便地去循环类数组对象。

```javascript
function argfn() {
  for (let argument of arguments) {
    console.log(argument);
  }
}
argfn(1,2,3)
// 1
// 2
// 3
```

上面的代码可以看出来，打印的结果只有 `1、2、3` 没有类数组上的其他属性值。

**2. 迭代 DOM 集合**

其实最常见的数组对象是得到网页上的 DOM 元素的集合，它也是一个类数组对象。比如一个 NodeList 对象：下面的例子演示给每一个 p 标签添加一个 “read” 类。

```javascript
//注意：这只能在实现了NodeList.prototype[Symbol.iterator]的平台上运行
let prags = document.querySelectorAll("p");

for (let value of prags) {
  value.classList.add("read");
}
```

上面的代码，需要在在带有 p 的标签的 html 文件中运行。

## 3 知识对比

ES5 中提供了很多遍历的方法，下面我们与之一一对比看看 `for...of` 有什么优势。

### 3.1 对比 for

最原始的语法是 `for` 循环语句，但是这种写法比较麻烦，每个步骤的信息都需要手动地去处理。

```javascript
const fib = [1,1,2,3,5,8,13...];  // 斐波那切数列
for (let index = 0; index < fib.length; index++) {
  console.log(fib[index]);
}
```

### 3.2 对比 forEach

数组中内置了 `forEach` 方法，这个方法的致命缺点就是不能跳出循环，`break` 命令和 `return` 命令都不能奏效。

```javascript
fib.forEach((value) => {
  console.log(value);
});
```

### 3.3 对比 for…in

* `for...in` 用以遍历对象的属性，`for...of` 用以遍历数据，就像数组中的值一样；
* `for...in` 主要是针对对象循环而设计的，对于数组，键就是数字，但是在 `for...in` 循环中是以字符串作为键名；
* `for...in` 循环不仅遍历数字键名，还会遍历手动添加的其他键，甚至包括原型链上的键；
* 某些情况下，`for...in` 循环以任意顺序遍历键名，主要是因为对象在内存中的数据类型决定的。

```javascript
for (let index in fib) {
  console.log(fib[index]);
}
```

### 3.4 `for...of` 的优点

* 有着同 `for...in` 一样的简洁语法，但是没有 `for...in` 那些缺点；
* 不同于 `forEach` 方法，它可以与 `break`、`continue` 和 `return` 配合使用；
* 提供了遍历所有数据结构的统一操作接口。

```javascript
for (let n of fib) {
  if (n > 520)
    break;
  console.log(n);
}
```

当迭代项大于 520 时 `break` 语句会跳出 `for...of` 循环。

## 4. 小结

* 对于数组的处理尽量使用 `for...of` 去迭代数据；
* 如果要遍历的是对象，并且没有顺序的限制可以使用 `for...in` 方式遍历对象更好的处理数据。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
