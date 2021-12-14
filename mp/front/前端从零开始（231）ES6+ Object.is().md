# ES6+ [Object.is](http://Object.is)()

## 1. 前言

在 ES5 中判断两个值是否相等基本都是使用 `==` 或 `===` 来判断，ES6 提供了 `Object.is()` 方法来判断两个值是否相同，这个方法弥补了使用等号方式判断所存在的问题。

## 2. 方法详情

### 2.1 基本语法

`Object.is()` 会接收两个需要判断的参数，最后返回一个布尔值，如果相同则返回 `true` 否则返回 `false`。

**语法使用：**

```javascript
Object.is(value1, value2);
```

**参数解释：**

|参数|描述|
|----|----|
| value1| 第一个需要比较的值|
| value2| 第二个需要比较的值|

### 2.2 基本使用

下面通过一些案例来说明一下 `Object.is()` 的使用：

```javascript
Object.is('imooc', 'imooc');  // true
Object.is('imooc', 'mooc');   // false

Object.is(window, window);    // true
Object.is([], []);           // false

var foo = { a: 1 };
var bar = { a: 1 };
var obj = foo;
Object.is(foo, foo);         // true
Object.is(foo, bar);         // false
Object.is(foo, obj);         // true

Object.is(null, null);       // true

// 特例
Object.is(0, -0);            // false
Object.is(0, +0);            // true
Object.is(-0, -0);           // true
Object.is(NaN, 0/0);         // true
```

上面的代码中，需要注意的是，在对象判断时，如果判断的数据是引用类型，即使两个对象值是一样的，但是它们也会返回 `false`，这是由于对象的存储方式决定的。在判断 0 和 - 0 的时候，二者不是同一个值，所以返回 `false`。

## 3. 对比等号 (=)

判断值相等，一般有两种 `==`（双等） 和 `===` （三等），他们之间有所不同。在值对比的过程中，类型转换在中间起到重要的作用，在对比过程中必须考虑到，值是否被类型转换了。

### 3.1 == 运算符

`==` 是一个非严格相等的对比运算符，它只会对比两边的操作数是否相等，相等则会返回 `true`。如果对比的操作数类型不同，则会自动将值进行隐式转换为一种常见的类型，然后才进行相等性比较。我们来看如下的实例：

```javascript
0 == -0 			// true
0 == '0' 			// true
0 == false          // true
0 == ''        		// true
"" == false         // true
null == undefined   // true
1 == '1'  		    // true
true == 'true' 		// false
NaN == 'NaN' 		// false
NaN == NaN 			// false

{"name": "imooc"} == {"name": "imooc"} // false

let a = {"name": "imooc"}
let b = a
console.log(a == b) // true
```

上面的代码中，列出了 `==` 判断大部分场景，但是这样的判断方式存在严谨性，在类型不同的时候会做类型转换，这样不利于判断两个值是否真实相等。我们总结以下几点：

* `NaN` 不等于包含它自身在内的任何值；
* 0 和 - 0 相等；和 false 也是相等的，和空字符串也是相等的；
* `null` 等于 `null` 和 `undefined`；
* 操作的值可以被自动转换为 `String`、`Boolean`、`Number` 三种类型；
* `String` 类型的比较会区分操作值大小写；
* 两个操作值如果引用同一个对象，返回 `true`，否则 `false`；
* 6 个虚值 (`null`, `undefined`, ‘’ , `0` , `NaN` , `false`)。

### 3.2 === 运算符

`===` 是严格相等的，被称作全等操作符，和 `==` 很相似，区别在于 `===` 不执行隐式类型转换。只有当两个操作值的值与类型都相等的前提下，才会返回 `true`。但是一些操作值的判断还是会有问题：

```javascript
+0 === -0             // true
true === true 				// true
null === null 				// true
NaN === NaN 					// false, NaN永远不等于NaN

1 === '1' 						// false, 值类型不同：数值和字符串
true === 'true' 			// false
null === undefined 		// false

'Imooc' === 'imooc' 	// false, 严格区分大小写

null
```

上面的代码中，存在两个问题，+0 和 - 0 是全等的，虽然它们没有进行隐式转化，但是它们是带符号的，我们其实也不希望它们是相等的。`NaN` 不等于任何值，包括它自己，这样我们就无法判断两个 NaN 值相等了。根据上面的例子，我们可以总结以下几点：

* `NaN` 不等于包含它在内的任何值；
* 0 等于 - 0，+0 和 -0 也相等；
* `null` 等于 `null`，但不等于 `undefined`；
* `String` 严格区分大小写；
* 两个操作值如果引用同一个对象，返回 `true`，否则 `false`。

### 3.3 [Object.is](http://object.is/)()

`Object.is()` 被称为同值相等的比较，在两个值进行比较时用到了很多规则，比如上面 `===` 在判断带符号的 0 时返回的都是 `true`，`NaN` 和任何值都不相等，但 `Object.is()` 给出了截然相反的结果：

```javascript
Object.is(0, -0);     // false
Object.is(NaN, NaN);  // true
```

下面我们来看下，`Object.is()` 都有哪些规则：

1. 当操作值都没有被定义时，这时它们的值是 `undefined` ，通过 `Object.is()` 判断的结果为 `true`。

```javascript
let a
let b
Object.is(a,b) // true
```

1. `Object.is()` 也是严格区分大小写的。

```javascript
Object.is('Imooc', 'Imooc') // true
Object.is('Imooc', 'imooc') // false
```

1. 操作值的类型必须相同，无论是什么值，只要类型不一样就会返回 `false`。

```javascript
Object.is(null, 'null') // false
Object.is(undefined, 'undefined') // false
```

1. 判断引用类型的值时，引用类型的地址相同时，则相等，与 `==` 和 `===` 判断的结果是一样的。

```javascript
let a = {"name": "imooc"}
let b = a

Object.is(a, b) // true

Object.is({"name": "imooc"}, {"name": "imooc"}) // false

Object.is(window, window) 		// true, 只有一个window全局变量
```

1. 操作数是 0、+0、-0 的比较，0 和 +0 是相等的，因为正号可以省略，但是 0 和 -0 是不相等的，这样就和 `===` 判断的结果不一样了。

```javascript
Object.is(0, +0) 	// true
Object.is(0, -0) 	// false
```

1. 当两个操作值是 `NaN` 时，使用 `Object.is()` 返回的结果是 `true` 这个和 `===` 返回的结果不一样，如果计算的结果是 `NaN` 的话，返回的结果也是 `true`。

```javascript
Object.is(NaN, NaN) 	// true
Object.is(NaN, 0/0) 	// true
```

## 4. 小结

本节讲解了判断两个值是否相同的方法 `Object.is()`，这个方法弥补了使用等号（=）判断的不足，更加准确地判断两个值是否相同。同时处理 0 和 - 0 以及 `NaN` 值的判断，所以在开发中，尽量使用 `Object.is()` 方法去判断两个相同的值，会更加准确明了。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
