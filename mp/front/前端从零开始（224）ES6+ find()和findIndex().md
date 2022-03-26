# ES6+ find () 和 findIndex ()

## 1. 前言

[上一节](https://www.imooc.com/wiki/ES6lesson/arrayincludes.html) 我们学习了 `includes()` 方法用于查找数组，但在数组中我们希望查找一个符合某个条件的元素，在 ES5 中有 `filter` 方法可以用于过滤符合条件的元素，但是 `filter` 返回的是一个数组，其实我们只想得到符合条件的值或者索引。本节将学习 ES6 的 `find` 和 `findIndex` 方法，丰富了数组查询的方法。

## 2. 语法详解

### 2.1 基本语法

这两个方法的使用基本相同，只是它们的返回结果不同。`find` 方法返回的是数组中符合条件的第一个值，`findIndex` 方法则返回符合条件的第一个索引的位置。它们都只是关注第一个查找到的结果，在查找到结果以后就不会继续查找了。

**使用语法：**

```javascript
arr.find(callback[, thisArg])
arr.findIndex(callback[, thisArg])
```

**参数解释：**

|参数|描述|
|----|----|
|callback|一个回调函数，接受数组的每一项并执行该函数，当主动返回 `true` 时，则终止调用|
|thisArg |（可选）执行 `callback` 时作为 `this` 对象的值                              |

`callback` 函数有三个参数：当前元素的值、当前元素的索引，以及数组本身。数组中的每一项元素都会执行一次 `callback` 函数，直到 `callback` 返回 `true` 时，则终止调用，并且把查找的结果返回。否则返回 `undefined`。

如果提供了 `thisArg` 参数，那么它将作为每次 `callback` 函数执行时的 `this`，如果未提供，则使用 `undefined`。

### 2.2 方法示例

以下是 `find` 和 `findIndex` 的使用示例，以便更好地理解这两个方法。

```javascript
var arr = [1, 6, 3, 4, 5]
var target = arr.find(function(item) {
  return item % 2 === 0
})
console.log(target)   // 6

var target = arr.findIndex(function(item) {
  return item % 2 === 0
})
console.log(target)   // 1
```

上面的代码找出数组中是 2 的倍数的项，`find` 方法返回的是数组中符合条件的第一个值 6；`findIndex` 方法返回的是数组中符合条件的第一个索引的位置 1。

## 3. 语法对比

在 ES5 中有 `filter` 方法可以用于查找符合条件的元素，`filter` 会遍历整个数组把符合条件的数组都返回出来，与 `find` 不同的是，它的返回结果是一个符合查询条件的数组。`find` 和 `findIndex` 则只关注数组中有没有符合查询条件的元素，而且只关注查询到的第一个元素。

```javascript
let arr = [1,2,3,4,5]
let find = arr.filter(function(item) {
  return item % 2 === 0
})
console.log(find)   // [2, 4]
```

上面的代码是找到所有满足 2 的倍数的数组，返回的结果同样是一个数组。

## 4. 小结

本节讲解了数组的 `find` 和 `findIndex` 方法的使用，主要注意的是它们的返回的结果不同，`find` 方法返回的是数组中符合条件的第一个值，`findIndex` 方法则返回符合条件的值的第一个索引。ES5 中的 `filter` 方法也可以查询，它返回的是满足条件的整个数组，但这两个方法都只关注查找的值是否存在。在只关注是否存在的情况下，这两个方法的效率要高，丰富了对数组的查找场景。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
