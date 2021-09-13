# ES6+ includes()

## 1. 前言

在字符串中我们学习了 `includes ()` 方法，在数组中同样存在 `includes ()` 方法，用来查找数组。`includes ()` 的存在是为了取代 `indexOf ()` 方法而设计的， `indexOf()` 在数组查找时存在一定缺陷，对于数组中元素是 `undefined`、`NaN` 时查找的结果是有问题的。为了保持语法的一致性和简洁性 `indexOf ()` 方法也是有必要的，本节我们就来学习数组中的 `includes()` 方法。

## 2. 方法详解

`includes ()` 方法用于查找一个数组中是否包含一个指定的元素，并返回一个布尔值，如果包含返回 true 否则返回 false。

**使用语法：**

```javascript
arr.includes(valueToFind[, fromIndex])
```

**参数解释：**

|参数|描述|
|----|----|
|valueToFind|需要查找的目标值                                                                                                                                                                                 |
|fromIndex  |（可选）从 `fromIndex` 索引处开始查找 `valueToFind`。如果为负值，则按升序从 `array.length + fromIndex` 的索引开始搜 （就是从末尾开始往前跳 `fromIndex` 的绝对值个索引，然后往后搜寻）。默认为 0。|

**语法实例：**

```javascript
var arr = ['imooc', 'ES6', 'wiki'];

console.log(arr.includes('ES6'));     // true
```

和字符串中的 `includes()` 方法一样，当没有参数时，`includes()` 方法会把第一个参数置成 `undefined`，注意，不同的是这里的 `undefined` 不是字符串 “undefined”。如下实例：

```javascript
[undefined].includes();		// true

['undefined'].includes();	// false
```

## 3. indexOf () 的问题

`indexOf()` 在查询数组中元素时存在一些问题，下面我们就来看看为什么 ES6 要引入 `includes()` 方法。

在 ES5 中使用 `indexOf()` 方法在数组中可以找到一个给定元素的第一个索引，如果不存在，则返回 -1。但是查找数组时存在一定缺陷，`indexOf` 不能判断数组中是否有 `NaN`，对于数组中的空项也不能判断。

```javascript
var arr1 = [,,,,,];
var arr2 = [null, undefined, NaN];
console.log(arr1[0], arr1[1])		// undefined undefined
arr1.indexOf(undefined)            // -1
arr2.indexOf(NaN);    // -1
```

上面的代码可以看到，在第 1 行中数组的每一项都是空的， 使用 `indexOf()` 查找返回的结果为 -1，没有查到 `undefined` 值，但从第 3 行打印的结果可以看到其实空数组的每一项都是 `undefined`。另外，还有个问题 `indexOf()` 不能解决，数组中有 NaN 时查询不了，返回的结果也是 -1。ES6 的`includes()` 可以完美解决上面的问题，看如下示例：

```javascript
[,,,,,].includes(undefined)           // true
[null, undefined, NaN].includes(NaN)]   // true
```

从上面的代码可以看出，使用 `includes()` 查询可以得到正确的结果。

indexOf 返回的是数值型的，而 includes 返回的是布尔型的，方便逻辑判断。如下实例：

```javascript
var arr = ['imooc', 'ES6', 'wiki'];
if (arr.includes('ES6')) {
  // todo
}
if (arr.indexOf('ES6') !== -1) {
  // todo
}
```

## 4. 使用场景

### 4.1 参数简介

`includes()` 方法返回一个布尔值，表示某个数组是否包含给定的值，如果给定的值是 `NaN` 也是可以判断的。

```javascript
[1, 2, 3].includes(2);            // true
[1, 2, 3].includes(4);            // false
[1, 2, NaN].includes(NaN);        // true
[undefined].includes(undefined)   // true
```

该方法的第二个参数表示搜索的起始位置，包括当前的位置，如果第二个参数大于或等于数组的长度时，则返回 `false`。

```javascript
[1, 2, 3].includes(3, 3);   // false
[1, 2, 3].includes(3, 2);   // true
```

### 4.2 第二个参数为负值时

如果第二个参数为负值时，计算数组的长度和第二个参数之和小于 0，则整个数组都会被搜索。

```javascript
var arr = ['a', 'b', 'c'];

arr.includes('a', -10);   // true
arr.includes('b', -10);   // true
arr.includes('a', -2);    // false
```

arr 的数组长度是 3，第二个参数是 - 10，计算之和为 -7 小于 0，则整个数组都会被搜索。

### 4.3 作为通用方法

我们看到在字符串和数组中都有 `includes()` 方法，其有意设计为通用方法。它不要求 `this` 值是数组对象，所以它可以被用于其他类型的对象 （比如类数组对象）。下面的例子展示了 在函数的 arguments 对象上调用的 `includes()` 方法。

```javascript
function fn() {
  console.log([].includes.call(arguments, 'a'));
  console.log([].includes.call(arguments, 'd'));
}
fn('a', 'b', 'c');
// true
// false
```

上面的代码中，includes 方法接收 arguments 对象，并且正确地返回相应的结果。

## 5. 小结

本节讲解了数组的 `includes()` 方法的使用，并对比了 ES5 中的 `indexOf()` 并回答了为什么 `includes()` 就是比较好的选择 ，总结有以下几点：

1. `includes()` 返回的是布尔值更容易做条件判断；
2. `indexOf()` 不能对数组空项和数组项为 NaN 的元素进行查找，而 `includes()` 可以很好地解决这个问题；
3. `includes()` 被设计为通用方法，可以在类数组中使用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
