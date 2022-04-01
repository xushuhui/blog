# ES6+ Reflect（一）

## 1. 前言

任何一门语言在走向成熟的过程都是趋向精细化和规范化，在 API 设计之初满足当时的需求场景。随着前端的飞速发展，软件复杂度的提升，很多 API 在使用过程中存在很多使用别扭的情况，不符合软件开发的规范。上一节我们学习了 Proxy，Proxy 的设计目的是为了取代 Object.defineProperty，优化性能，使得数据劫持的过程更加规范。

本节我们将学习 ES6 新的全局对象 —— Reflect（反射），首先我们要了解一下，为什么会新添加这么一个全局对象？Reflect 上的一些函数基本上都可以在 Object 上找到，找不到的，也是可以通过对对象命令式的操作去实现的；那么为什么还要新添加一个呢？本节我们将学习 Reflect 的相关知识。

## 2. 基础知识

Reflect 是一个内置的对象，它提供了拦截 JavaScript 操作的方法。这些方法与 Proxy 中的 handlers 方法相同。与大多数全局对象不同 `Reflect` 并非一个构造函数，所以不能通过 new 运算符对其进行调用，或者将 `Reflect` 对象作为一个函数来调用。`Reflect` 的所有属性和方法都是静态的（类似 `JSON` 或者 `Math` 等对象）。

### 2.1 基本用法

Reflect 可以检查对象上是否存在特定属性，可以使用 `Reflect.has()` 方法检测。

```javascript
let key = Symbol.for('a');
const obj = {
  name: 'imooc',
  lession: 'ES6 Wiki',
  [key]: 100
}

console.log(Reflect.has(obj, 'name'));	// true
console.log(Reflect.has(obj, 'age'));		// false
```

可以使用 `Reflect.get()` 方法获取对象上的属性值。

```javascript
console.log(Reflect.get(obj, 'name'));	// imooc
```

可以使用 `Reflect.set()` 方法为对象添加一个新的属性。

```javascript
const res = Reflect.set(obj, 'age', 7);
console.log(res);		// true
console.log(obj);		// {name: "imooc", lession: "ES6 Wiki", age: 7}
```

使用 `Reflect.ownKeys()` 方法获取对象上的自有属性。

```javascript
console.log(Object.keys(obj));	// ["name", "lession"]

console.log(Reflect.ownKeys(obj));	// ["name", "lession", Symbol(a)]
```

上面的代码可以看出，使用 `Object.keys()` 获取不到属性是 Symbol 的值。

### 2.2 返回值

Reflect 对象上的方法并不是专门为对象设计的，而是在语言层面的，它可以拿到语言内部的方法，和 Proxy 的结合可以实现元编程。并且每个操作都是有返回值的，上节我们使用 Proxy 简单地实现了 Vue3 的响应式。但是在 Vue3 源码中获取和设置对象上的属性使用的是 Reflect，Reflect 会返回一个状态表示获取和设置的成功与否。

```javascript
// const res = target[key]; // 上节代码
const res = Reflect.get(target, key);	// 获取target上属性key的值

// target[key] = value;	// 上节代码
const result = Reflect.set(target, key, value);	// 设置目标对象key属性的值
```

上面的两段代码是 Vue3 中的源码，因为在源码中需要知道获取或赋值的结果，因为可能获取失败。在 ES5 中如果想要监听劫持属性操作的结果需要使用 `try...catch` 的方式。

```javascript
try {
  Object.defineProperty(obj, prop, descriptor);
  // success
} catch (e) {
  // failure
}
```

Reflect 在操作对象时是有返回结果的，而 Object.defineProperty 是没有返回结果的，如果失败则会抛出异常，所以需要使用 `try...catch` 来捕获异常。

### 2.3 使用函数代替命令式

Object 中操作数据时，有一些是命令式的操作，如：`delete obj.a` 、`name in obj` ，Reflect 则将一些命令式的操作如 `delete`，`in` 等使用函数来替代，这样做的目的是为了让代码更加好维护，更容易向下兼容；也避免出现更多的保留字。

```javascript
// ES5
'assign' in Object // true
// ES6
Reflect.has(Object, 'assign') // true

delete obj.name;	// ES5
Reflect.deleteProperty(obj, 'name');	// ES6
```

## 3. 静态方法

Reflect 的出现是为了取代 Object 中一些属于语言层面的 API，这些 API 在 Object 上也是可以找到的，并且它们的功能基本是相同的。上面我们也提到了 Reflect 和 Proxy 中 handlers 的方法是一一对应的，在很多场景中它门都是配套使用的。这里我们就来学习一下 Reflect 提供的静态方法：

### 3.1 Reflect.get()

`Reflect.get()` 方法是从对象中读取属性的值，类似 ES5 中属性访问器语法： `obj[key]` ，但是它是通过调用函数来获得返回结果的。

**语法：**

```javascript
Reflect.get(target, propertyKey[, receiver])
```

* target：需要取值的目标对象；
* propertyKey：需要获取的值的键值；
* receiver：如果 target 对象中指定了 getter，receiver 则为 getter 调用时的 this 值。

如果目标值 target 类型不是 `Object`，则抛出一个 `TypeError`。

```javascript
// Object
var obj = { a: 1, b: 2 };
Reflect.get(obj, "a"); // 1

// Array
Reflect.get(["a", "b", "c"], 1); // "one"
```

第三个参数 receiver 是 this 所在的上下文，不传时指的是当前对象，如果传如一个人对象则 this 指向该对象。下面我们来看个实例：

```javascript
let obj = {
  name: 'imooc',
  lesson: 'ES5 Wiki',
  get info() {
    console.log(`这是慕课 ${this.lesson}`);
    return 0
  }
};
Reflect.get(obj, 'info');	// 这是慕课 ES5 Wiki
Reflect.get(obj, 'info', {lesson: 'ES6 Wiki'});	// 这是慕课 ES5 Wiki
```

### 3.2 Reflect.set()

`Reflect.set()` 是在一个对象上设置一个属性，类似 ES5 中属性设置语法：`obj[key] = value` ，它也是通过调用函数的方式来对对象设置属性的。

**语法：**

```javascript
Reflect.set(target, propertyKey, value[, receiver])
```

* target：表示要操作的目标对象；
* propertyKey：表示要设置的属性名；
* value：表示设置的属性值；
* receiver：表示的是一个 this 值，如果我们在设置值的时候遇到 setter 函数，那么这个 receiver 值表示的就是 setter 函数中的 this 值。

这个函数会返回一个 Boolean 值，表示在目标对象上设置属性是否成功。

```javascript
// Object
var obj = {};
Reflect.set(obj, "name", "imooc"); // true
console.log(obj.name); // "imooc"

// Array
var arr = ["a", "b", "c"];
Reflect.set(arr, 2, "C"); // true
console.log(arr); // ["a", "b", "C"]
```

使用可以截断数组：

```javascript
var arr = ["a", "b", "c"];
Reflect.set(arr, "length", 2); // true
console.log(arr);	// ["a", "b"]
```

当有 receiver 参数时，如果 receiver 对象中有 propertyKey 属性，则会使用 receiver 对象中的值。

```javascript
Reflect.set(obj, 'lession', 'ES5 Wiki', {lession: 'ES6 Wiki', age: 18});
console.log(obj);	// {name: "imooc", lesson: "ES5 Wiki"}
```

### 3.3 Reflect.deleteProperty()

`Reflect.deleteProperty()` 方法允许删除对象的属性。它类似 ES5 中的 `delete` 操作符，但它也是一个函数，通过调用函数来实现。

**语法：**

```javascript
Reflect.deleteProperty(target, propertyKey)
```

* target：表示要操作的目标对象；
* propertyKey：表示要删除的属性。

这个函数的返回值是一个 Boolean 值，如果成功的话，返回 true；失败的话返回 false。我们来看下面的实例：

```javascript
var obj = {
    name: 'imooc',
    lession: 'ES6 Wiki'
};

var r1 = Reflect.deleteProperty(obj, 'name');
console.log(r1); // true
console.log(obj); // {lession: "ES6 Wiki"}

var r2 = Reflect.deleteProperty(Object.freeze(obj), 'lession');
console.log(r2); // false
```

上面的例子中使用 `Object.freeze()` 方法来冻结 obj 对象使之不能被修改。

### 3.4 Reflect.has()

`Reflect.has()` 方法可以检查一个对象上是否含有特定的属性，这个方法相当于 ES5 的 `in` 操作符。

**语法：**

```javascript
Reflect.has(target, propertyKey)
```

* target：表示要操作的目标对象；
* propertyKey： 属性名，表示需要检查目标对象是否存在此属性。

这个函数的返回结果是一个 Boolean 值，如果存在就返回 true，不存在就返回 false。当然如果目标对象 (target) 不是一个对象，那么就会抛出一个异常。

```javascript
Reflect.has({x: 0}, "x"); // true
Reflect.has({x: 0}, "y"); // false

// 如果该属性存在于原型链中，也返回true
Reflect.has({x: 0}, "toString");	// true
```

这方法也可检查构造函数的属性。

```javascript
function A(name) {
    this.name = name || 'imooc';
}
// 在原型上添加方法
A.prototype.getName = function() {
    return this.name;
};

var a = new A();

console.log('name' in a); // true
console.log('getName' in a); // true

let r1 = Reflect.has(a, 'name');
let r2 = Reflect.has(a, 'getName');
console.log(r1, r2); // true true
```

### 3.5 Reflect.ownKeys()

`Reflect.ownKeys()` 返回一个由目标对象自身的属性键组成的数组。

**语法：**

```javascript
Reflect.ownKeys(target)
```

* target：表示目标对象

如果这个目标对象不是一个对象那么这个函数就会抛出一个异常。这个数组的值等于 `Object.getOwnPropertyNames(target).concat(Object.getOwnPropertySymbols(target))` 我们来看下面的实例：

```javascript
let a = Symbol.for('a');
let b = Symbol.for('b');

let obj = {
    [a]: 10,
    [b]: 20,
    key1: 30,
    key2: 40
};

let arr1 = Object.getOwnPropertyNames(obj);
console.log(arr1); // [ 'key1', 'key2' ]
let arr2 = Object.getOwnPropertySymbols(obj);
console.log(arr2); // [ Symbol(a), Symbol(b) ]
let arr3 = Reflect.ownKeys(obj);
console.log(arr3); // [ 'key1', 'key2', Symbol(a), Symbol(b) ]
```

## 4. 小结

本节主要学习了 ES6 新增的全局对象 `Reflect` ，它的目的是为了分离 Object 中属于语言部分的内容，每个使用 `Reflect` 下的方法操作的对象都要返回值。 `Reflect` 对象和 `Proxy` 下的方法是一一对应的，二者配合可以实现很多功能。Vue3 中的数据响应就是使用的它们。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
