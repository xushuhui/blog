# Symbol

## 1. 前言

在 ES5 中基础数据类型有 5 种：`Boolean`、`Null`、`Undefined`、`Number`、`String`，ES6 新增了一个基础数据类型 `Symbol` 符号、代号的意思，它是独一无二的，也就是说使用它声明的变量是独一无二的。引入这个数据类型有什么作用呢？

我们知道在 ES5 中， 对象的属性名都是字符串，容易造成属性名冲突。比如，你使用了一个他人提供的对象，但又想为这个对象添加新的方法（mixin 模式），新方法的名字就有可能与现有方法产生冲突，ES6 引入 Symbol 就可以解决这个问题。不仅如此 Symbol 的使用还有很多，在元编程中也发挥很大的作用。下面我们就来看看 Symbol 的 使用。

## 2. 语法详解

### 2.1 基本语法

使用 `Symbol()` 会返回一个独一无二的变量，可以作为对象的 key 存在，返回的值是 `symbol` 类型，该类型具有静态属性和静态方法。

```javascript
Symbol([description])
```

|参数|描述|
|----|----|
|description|（可选）是字符串类型，对 symbol 的描述|

### 2.2 基本语法

`Symbol()` 是一个方法，返回的值是 `symbol` 类型，使用如下：

```javascript
var s1 = Symbol();
var s2 = Symbol();
console.log(s1 === s2)  // false
console.log(typeof s1)	// symbol
```

上面的代码中，使用 `Symbol` 声明的变量 s1 和 s2，虽然它们看上去是同一个 `Symbol` 定义的，但其实是不相等的。

### 2.3 带描述的 Symbol

在 `Symbol` 中可以传入一些参数，来描述定义的 `Symbol` 类型的值。

```javascript
var s1 = Symbol('imooc');
var s2 = Symbol('imooc');
console.log(s1)         // Symbol(imooc)
console.log(s1 === s2)  // false

let s = Symbol({name: 'imooc'});
console.log(s);         // Symbol([object Object])
```

上面的代码中，`Symbol` 接收的参数是一个对 `Symbol` 的描述，即使两个 `Symbol` 接收相同值，两个值也是不一样的。另外，如果传入的描述符是对象类型，内部会将描述的内容进行 `toString` 操作，所以返回的结果是 `[object Object]`。

## 3. 作为对象的 key

`Symbol` 经常会作为对象的属性存在，如果这个属性是用 `symbol` 来声明的，则不可枚举，也不能用 `for...in`、`for...of` 迭代。

对象上的 key，可以用取值表达式 （中括号） 的方式取出来，作为对象上的属性，如下：

```javascript
var s = Symbol('imooc');
var obj = {
  [s]: 1
}
obj     // {Symbol(imooc): 1}
obj[s]  // 1
```

上面的代码，使用 `Symbol` 声明了一个变量，然后作为对象的 key 给它赋值。取值的时候只能使用中括号的方式，因为这里的 `s` 是变量不能使用点的方式。下面是对 obj 对象用 `for...in` 进行的遍历。

```javascript
for(let key in obj) {
  console.log(obj[key])
}
// undefined
```

上面的代码对 obj 对象进行迭代，但是没有打印出对应的值，说明用 `Symbol` 来声明的属性是不可枚举的。如果想要获取到这个属性可以使用 `Object.getOwnPropertySymbols(obj)` 获取。使用 `Object.keys()`、`Object.getOwnPropertyNames()`、`JSON.stringify()` 也是不能返回想要的结果。

```javascript
Object.getOwnPropertySymbols(obj);   // [Symbol(imooc)]
Object.keys(obj);                    // []
Object.getOwnPropertyNames(obj);     // []
JSON.stringify(obj);                 // "{}"
```

从上面的代码中可以看到使用 `Object.getOwnPropertySymbols()` 可以获取对象上所有 `Symbol` 类型的属性，并返回一个数组。

另外，可以通过 `description` 方法获取 `Symbol` 类型描述。

```javascript
var s = Symbol('imooc');
s.description;    // "imooc"
```

## 4. Symbol.for () 和 Symbol.keyFor ()

`Symbol.for(key)` 方法也是声明变量使用，不同的是 `Symbol.for(key)` 是在全局作用域下声明的。它会根据给定的键 `key` 从运行时的 `symbol` 注册表中找到对应的 `symbol`。如果找到了，则返回它。否则，新建一个与该键关联的 `symbol`，并放入全局的 `symbol` 注册表中，如果有已经声明了的 `symbol` 则不回重复声明。

```javascript
let s1 = Symbol.for('imooc');
let s2 = Symbol.for('imooc');
function fn() {
  return Symbol.for('imooc');
}
console.log(s1, s2)       // Symbol(imooc) Symbol(imooc)
console.log(s1 === s2)    // true
console.log(fn() === s1)    // true
```

上面的代码中可以看出来，使用 `Symbol.for(key)` 无论在哪里进行声明，都不会影响它们的值。

`Symbol.keyFor()` 通过 `key` 值获取 `symbol` 的描述：

```javascript
let s1 = Symbol.for('imooc');
console.log(Symbol.keyFor(s1))  // imooc
```

## 5 实战案例

### 5.1 解决属性重名

在现实中姓名重复是很常见的，但是在 JavaScript 对象中，属性名是唯一的存在。如果定义一个对象中有重复的属性则会被覆盖，这个现象叫做 “引用消除”。我们看下面的例子：

```javascript
var person = {
  Tom: {sex: '男', age: 18},
  David: {sex: '男', age: 17},
  David: {sex: '女', age: 16},
}
console.log(person); // {Tom: {sex: "男", age: 18}, David: {sex: "女", age: 16}}
```

可以看到我们定义了一个 person 对象，第一个 David 对象被后一个 David 引用消除了，所以只有第二个 David 的数据。如果要解决这个问题可以使用 `Symbol` 来实现

```javascript
var person = {
  Tom: {sex: '男', age: 18},
  [Symbol('David')]: {sex: '男', age: 17},
  [Symbol('David')]: {sex: '女', age: 16},
}
console.log(person)
// {Tom: {sex: "男", age: 18}, Symbol(David): {sex: "男", age: 17}, Symbol(David): {sex: "女", age: 16}}
```

这样就可以解决属性名冲突的问题，需要注意的是使用这样的方式定义对象数据存在一个问题，就是使用 `for...in` 或者使用 `Object.keys()` 遍历时 Symbol 属性的数据不会被遍历到，上文有具体说明。所以，如果想要遍历到对象的值可以通过 `Reflect.ownKeys()` 去获取对象的 key，然后进行循环操作。

```javascript
for (let key of Reflect.ownKeys(person)) {
	console.log(person[key])
}
// {sex: "男", age: 18}
// {sex: "男", age: 17}
// {sex: "女", age: 16}
```

### 5.2 消除魔术字符串

**魔术字符串** 指的是在代码中多次出现与代码形成强耦合的某一个具体的字符串或者数值，看下面的例子：

```javascript
function getArea(shape, options) {
  let area = 0;

  switch (shape) {
   	case 'Circle':
      area = 3.14 * Math.pow(options.radius, 2)
      break;
    case 'Square':
      area = options.width * options.height;
      break;
  }
  return area;
}

getArea('Circle', { radius: 10 });		// 314
getArea('Square', { width: 10, height: 10 });	 // 100
```

上面的代码中 ‘Circle’ 和 ‘Triangle’ 就属于魔术字符串，常见的消除魔术字符串的方法就是使用变量替代，如下：

```javascript
const shapeType = {
  circle: 'Circle',
  triangle: 'Square'
}
function getArea(shape, options) {
  let area = 0;

  switch (shape) {
   	case shapeType.circle:
      area = 3.14 * Math.pow(options.radius, 2)
      break;
    case shapeType.square:
      area = options.width * options.height;
      break;
  }
  return area;
}

getArea(shapeType.circle, { radius: 10 });		// 314
getArea(shapeType.square, { width: 10, height: 10 });	 // 100
```

上面的代码中就消除了代码的强耦合，其实我们不关注 shapeType 属性的值，只要他们不同即可，有了 Symbol 这时我们就可以使用 Symbol 进行描述，如下更改：

```javascript
const shapeType = {
  circle: Symbol('Circle'),
  triangle: Symbol('Square')
}
```

## 6. 小结

本节学习了 ES6 新增数据类型 `Symbol`，使用它可以声明一个独一无二的变量，通常会作为对象的属性存在，解决属性名冲突的问题。注意这个属性是不能被迭代的，如果想要迭代它可以使用 `Reflect.ownKeys()` 的方式去获取 key 值。最后介绍了 `Symbol` 在实战中的应用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
