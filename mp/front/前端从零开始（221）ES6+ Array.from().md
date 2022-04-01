# ES6+ Array.from()

## 1. 前言

在前端开发中经常会遇到类数组，但是我们不能直接使用数组的方法，需要先把类数组转化为数组。本节介绍 ES6 数组的新增方法 `Array.from()`，该方法用于将类数组对象（array-like）和可遍历的对象（iterable）转换为真正的数组进行使用。

## 2. 方法详情

### 2.1 基本语法

`Array.from()` 方法会接收一个类数组对象然后返回一个真正的数组实例，返回的数组可以调用数组的所有方法。

**语法使用：**

```javascript
Array.from(arrayLike[, mapFn[, thisArg]])
```

**参数解释：**

|参数|描述|
|----|----|
| arrayLike| 想要转换成数组的类数组对象或可迭代对象              |
| mapFn    | 如果指定了该参数，新数组中的每个元素会执行该回调函数|
| thisArg  | 可选参数，执行回调函数 mapFn 时 this 对象           |

### 2.2 类数组转化

所谓类数组对象，就是指可以通过索引属性访问元素，并且对象拥有 length 属性，类数组对象一般是以下这样的结构：

```javascript
var arrLike = {
  '0': 'apple',
  '1': 'banana',
  '2': 'orange',
  length: 3
};
```

在 ES5 中没有对应的方法将类数组转化为数组，但是可以借助 call 和 apply 来实现：

```javascript
var arr = [].slice.call(arrLike);
// 或
var arr = [].slice.apply(arrLike);
```

有了 ES6 的 `Array.from()` 就更简单了，对类数组对象直接操作，即可得到数组。

```javascript
var arr = Array.from(arrLike);
console.log(arr)  // ['apple', 'banana', 'orange']
```

### 2.3 第二个参数 —— 回调函数

在 `Array.from` 中第二个参数是一个类似 map 函数的回调函数，该回调函数会依次接收数组中的每一项作为传入的参数，然后对传入值进行处理，最得到一个新的数组。

`Array.from(obj, mapFn, thisArg)` 也可以用 map 改写成这样 `Array.from(obj).map(mapFn, thisArg)`。

```javascript
var arr = Array.from([1, 2, 3], function (x) {
  return 2 * x;
});
var arr = Array.from([1, 2, 3]).map(function (x) {
  return 2 * x;
});
//arr: [2, 4, 6]
```

上面的例子展示了，`Array.from` 的参数可以使用 `map` 方法来进行替换，它们是等价的操作。

### 2.4 第三个参数 ——this

`Array.from` 中第三个参数可以对回调函数中 this 的指向进行绑定，该参数是非常有用的，我们可以将被处理的数据和处理对象分离，将各种不同的处理数据的方法封装到不同的的对象中去，处理方法采用相同的名字。

在调用 Array.from 对数据对象进行转换时，可以将不同的处理对象按实际情况进行注入，以得到不同的结果，适合解耦。

```javascript
let obj = {
  handle: function(n){
    return n + 2
  }
}

Array.from([1, 2, 3, 4, 5], function (x){
  return this.handle(x)
}, obj)
// [3, 4, 5, 6, 7]
```

定义一个 `obj` 对象可以认作是，`Array.from` 回调函数中处理数据的方法集合，`handle` 是其中的一个方法，把 `obj` 作为第三个参数传给 `Array.from` 这样在回调函数中可以通过 `this` 来拿到 `obj` 对象。

### 2.5 从字符串里生成数组

`Array.from()` 在传入字符串时，会把字符串的每一项都拆成单个的字符串作为数组中的一项。

```javascript
Array.from('imooc');
// [ "i", "m", "o", "o", "c" ]
```

### 2.6 从 Set 中生成数组

用 `Set` 定义的数组对象，可以使用 `Array.from()` 得到一个正常的数组。

```javascript
const set = new Set(['a', 'b', 'c', 'd']);
Array.from(set);
// [ "a", "b", "c", "d" ]
```

上面的代码中创建了一个 Set 数据结构，把实例传入 `Array.from()` 可以得到一个真正的数组。

### 2.7 从 Map 中生成数组

`Map` 对象保存的是一个个键值对，`Map` 中的参数是一个数组或是一个可迭代的对象。 `Array.from()` 可以把 Map 实例转换为一个二维数组。

```javascript
const map = new Map([[1, 2], [2, 4], [4, 8]]);

Array.from(map);  // [[1, 2], [2, 4], [4, 8]]
```

## 3. 使用案例

### 3.1 创建一个包含从 0 到 99 (n) 的连续整数的数组

1. 一般情况下我们可以使用 for 循环来实现。

```javascript
var arr = [];
for(var i = 0; i <= 99; i++) {
  arr.push(i);
}
```

这种方法的主要优点是最直观了，性能也最好的，但是很多时候我们不想使用 for 循环来进行操作。

1. 使用 Array 配合 map 来实现。

```javascript
var arr = Array(100).join(' ').split('').map(function(item,index){return index});
```

Array (100) 创建了一个包含 100 个空位的数组，但是这样创建出来的数组是没法进行迭代的。所以要通过字符串转换，覆盖 undefined，最后调用 map 修改元素值。

1. 使用 es6 的 `Array.from` 实现。

```javascript
var arr = Array.from({length:100}).map(function(item,index){return index});
```

`Array.from({length:100})` 可以定义一个可迭代的数组，数组的每一项都是 undefined，这样就非常方便的定义出所需要的数组了，但是这样定义的数组性能最差，具体可以参考 [constArray](https://jsperf.com/constarray/4) 的测试结果。

### 3.2 数组去重合并

```javascript
function combine(){
  let arr = [].concat.apply([], arguments);  //没有去重复的新数组
  return Array.from(new Set(arr));
}

var m = [1, 2, 2], n = [2,3,3];
console.log(combine(m,n));                     // [1, 2, 3]
```

首先定义一个去重数组函数，通过 concat 把传入的数组进行合并到一个新的数组中去，通过 new Set () 可以对 arr 进行去重操作，再使用 `Array.from()` 返回一个拷贝后的数组。

## 4. 小结

本节讲解了字符串的 `Array.from()` 方法的使用，用于将类数组对象和可迭代的对象转化真正的数组，在编程中主要用于更加方便的初始化一个有默认值的数组，还可以用于将获取的 html 的 DOM 对象转化为数组，可以使用数组方法进行操作。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
