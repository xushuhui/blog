# ES6+ includes()

## 1. 前言

字符串查找一直都是程序中的常用操作，在 ES5 中查找一个字符串是否包含另一个字符串，一般有两种思维。一是使用正则的方式来匹配，二是使用 ES5 的方式，如 indexOf、lastIdexOf、search。使用正则比较简单，但是需要对正则语法有一定的了解。一般使用 `indexOf()` 来进行字符串的查找 ，它会返回查询后第一次出现的指定值的索引，如果未找到该值，则返回 -1。

ES6 新增方法 `includes()` 方法，来替代 `indexOf()` 弥补它的不足。而 ES6 提供了新的方法 `includes()` 可以更好地进行判断，而不需要根据返回的索引进行判断。

## 2. 方法详情

`includes()` 方法主要用于查询字符串，判断一个字符串是否包含另外一个字符串，其返回值是如果包含该字符串则返回 true 否则返回 false。

**使用语法：**

```javascript
str.includes(searchString[, position])
```

**参数说明：**

|参数|描述|
|----|----|
|searchString|需要查找的目标字符串                                         |
|position    |（可选） 从当前字符串的哪个索引位置开始搜寻子字符串，默认值为 0|

**实例：**

```javascript
var str = 'hello world';

console.log(str.includes('hello'));     // true
console.log(str.includes('hello', 3));  // false
```

在没有传入确切的查询字符串时，`searchString` 会被强制设置成 “undefined”，然后在当前字符串中查找这个值。

```javascript
'undefined'.includes('');			// 返回 true
'undefined'.includes(); 	  	// 返回 true
'undefine'.includes();		  	// 返回 false
'imooc ES6 wiki'.includes();	// 返回 false
```

上面的代码最容易在面试中出现，考察你对 `includes()` 方法的理解程度。第 1 行返回 `true` 很容易理解，查询一个空字符串嘛这个没问题，但第 2 行返回的结果也是 `true`，这说明查询的字符串在 “undefined” 字符串中，还不能说被设置成了字符串 “undefined” 。

第 3 、 4 行代码中也没有传值，但返回的结果为 `false`，从而证明了在没有传值时，第一个参数的值被设置成字符串 “undefined”。

在没传值时和 `indexOf()` 的查询结果是一致的。下面我们看 `indexOf()` 在没有参数时是一个什么样的结果。

```javascript
'undefined'.indexOf();		// 返回 0
'undefine'.indexOf();		  // 返回 -1
```

上面的代码中，第 1 行返回的结果是 0，是查询结果的位置，第 2 行返回的是 -1，说明没有查询到。虽然返回的结果不一样，但是意义是一样的。`includes()` 可以替换 `indexOf()` 使用，`includes()` 好处在于它可以直接判断，而 `indexOf()` 还需要对结果进行对比，如下实例：

```javascript
const str = 'imooc ES6 wiki';
if (str.includes('ES6')) {
  // todo

if (str.indexOf('ES6') !== -1) {
  // todo
}
```

上面代码中的两个 if 判断是一个意思，但是使用 `includes()` 很简洁，这也是 ES6 设计的初衷。

## 3. 使用场景

`includes()` 方法的引入是为了代替 `indexOf()` 作为字符串的查询的方法使用。

### 3.1 区分大小写

`includes()` 方法是区分大小写的。

```javascript
'imooc es6'.includes('imooc');  // true
'imooc es6'.includes('Imooc');  // false
```

### 3.2 一个参数

`includes()` 在只有一个参数时，会从字符串的第一个字符开始查找。

```javascript
var str = "I love imooc.";

console.log(str.includes("I love"));    // true
console.log(str.includes("imooc"));     // true
console.log(str.includes("eimooc"));    // false
```

### 3.3 两个参数

当 `includes()` 有第二个参数的时候，会从第二个参数作为索引的位置开始。

```javascript
var str = "I love imooc.";

console.log(str.includes("love", 3));  // false
console.log(str.includes("ove", 3));   // true
```

第二个参数的意思是，查找字符串开始的位置，例子中的 3 就是查找的位置，所以查找的目标字符串是 `ove imooc.`。

当第二个参数是负数时，只要查找的字符串在目标字符串里，无论是多少，都会返回 true。

```javascript
var str = "I love imooc.";

console.log(str.includes("love", -1));    // true
console.log(str.includes("ove", -100));   // true
console.log(str.includes("Love", -1));    // false
```

## 4. 注意事项

在使用 `includes()` 时需要注意类型转换的一些问题：

### 4.1 includes 会做类型转换

```javascript
let numStr = '2020';

numStr.includes('2');  // true
numStr.includes(2);    // true
```

在这个例子中 `numStr` 是一个字符串，判断字符串 `2` 和数字 `2` 都是能返回正确的结果，这里 `includes()` 会把数字转换成字符串 ‘2’ 然后再执行查询操作。

### 4.2 不能对 Number 类型直接使用

```javascript
let numStr = 2020;

numStr.includes(2);    // Uncaught TypeError: numStr.includes is not a function
```

从这个例子可以看出，`includes` 是字符串上的方法，而这里直接使用在数值类型上所以会报语法错误。如果要使用 `includes` 来查询，就必须把数字转化成字符串，然后进行查询。

```javascript
let numStr = 2020;
("" + numStr).includes(0) // true
```

这里对 `numStr` 前加一个空字符串可以进行类型转换。

## 5. 小结

本节讲解了字符串的 `includes()` 方法的使用，总结以下几点：

1. 在没有传参时，查询字符串会被设置成 “undefined”；
2. `includes()` 区分大小写；
3. 当有第二个参数时，则会从第二个参数作为索引的位置开始查找，并包含当前位置的字符；
4. 当第二个参数是负数时，只要查找的字符串在目标字符串里，无论是多少，都会返回 true；
5. 判断数字时，需要把数字转换成字符串类型才能查询。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
