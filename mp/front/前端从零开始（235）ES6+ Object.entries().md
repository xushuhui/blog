# ES6+ Object.entries()

## 1. 前言

前两节我们学习了获取可枚举对象上属性和属性值作为一个数组的方法 `Object.keys()` 和 `Object.values()`，但是我们想把键值对同时获取到怎么办呢？这时 ES6 提供了 `Object.entries()` 方法用于获取可枚举对象上的属性和值的数组。

## 2. 方法详情

`Object.entries()` 会返回一个二维数组，数组中的每一项是可枚举对象上的属性和值的数组。

### 2.1 语法使用

```javascript
Object.entries(obj);
```

|参数|描述|
|----|----|
|obj|可以返回其枚举自身属性和键的对象|

### 2.2 基本使用

返回结果是一个二维数组，数组中的每个元素是一个包含两个元素的数组，第一个元素是属性，第二个元素是属性对应的值。

```javascript
var obj = {a: 1, b: 2, c: 3};
console.log(Object.entries(obj))
// [['a', 1], ['b', 3], ['c', 3]]
```

如果传入的参数是数字或布尔值时，则返回一个空数组

```javascript
Object.entries(50)       // []
Object.entries(false)    // []
```

对 `values()` 返回数组的顺序，会按照属性的数值大小，从小到大排序。

```javascript
var obj = {10: 'a', 1: 'b', 7: 'c'};
console.log(Object.entries(obj))
// [["1", "b"], ["7", "c"], ["10", "a"]]
```

上面的代码中，属性名为数值的属性，是按照数值大小，从小到大遍历的，因此返回值是一个排序后的二维数组。

`values()` 传入的对象如果是一个字符串时，则会把字符拆成数组中的单个项，如下：

```javascript
console.log(Object.entries('abc'))
// [["0", "a"], ["1", "b"], ["2", "c"]]
```

## 3. 案例

### 3.1 用于 for…of 循环获取对象上的键值对

`Object.entries()` 在开发中有很好的用途，可以同时获取对象的键值对进行使用。

其实像使用 `Object.keys()`、 `Object.values()` 和 `Object.entries()` 都是为了遍历对象，然后对对象中的元素进行操作，下面我们来看一下 `Object.entries()` 在 for 循环中的使用。

```javascript
var obj = { name: 'imooc', age: 7, lesson: 'ES6 Wiki' };
for (let i = 0; i < Object.entries(obj).length; i++) {
  var [key, value] = Object.entries(obj)[i]
  console.log(key, value);
}
```

上面的代码中我们可以看出，在循环体内可以通过结构的方式获取对象的属性和值。使用 for 循环时我们要对循环的每一步进行处理。

在下面的章节中我们会学到 `for...of` 循环，它的功能强大，能遍历可迭代的对象，可以替代 `for`、 `forEach` 等循环，并具有扩展性。

```javascript
let obj = { name: 'imooc', age: 7, lesson: 'ES6 Wiki' };

for (let [key, value] of Object.entries(obj)) {
  console.log(`${key}: ${value}`);
}
// name: imooc
// age: 7,
// lesson: ES6 Wiki
```

上面的代码中，我们不用在 for 循环中处理每一项，可以在循环时直接解构出 `Object.entries()` 的值。这样就可以遍历出对象的键值对。

### 3.2 把对象转为 Map 结构

`Object.entries()` 还有一个最重要的功能，就是可以把指定的对象直接转化成 ES6 的 Map 数据结构。

ES6 提供了新的 `Map` 数据结构，它类似于对象，也是键值对的集合，但是 `Map` 的键可以是任意类型（原始类型和对象类型），并且提供了 `set` 、 `get` 方法去设置和获取对象的值。如果想把一个对象转为 `Map` 结构，可以借助 `Object.entries()` 来实现。

```javascript
var obj = { name: 'imooc', age: 7 };
var map = new Map(Object.entries(obj));
console.log(map)            // Map(2) {"name" => "imooc", "age" => 7}
console.log(map.get(name))  // imooc
```

上面的代码中，很好地把已有的对象，转化为 `Map` 对象，而且可以使用 `Map` 的方法获取对象上的数据。

## 4. Object.fromEntries()

`Object.fromEntries()` 是 `Object.entries()` 的反转函数，这样理解起来就比较轻松。它接收的是一个可迭代的对象，类似 `Array`、`Map` 或者其它实现了可迭代协议的对象，结果返回一个新的对象。

### 4.1 数组转化为对象

将一个带有键值对的数组转化成对象。

```javascript
var arr = [ ['a', '0'], ['b', '1'], ['c', '2'] ];
var obj = Object.fromEntries(arr);
console.log(obj); // {a: "0", b: "1", c: "2"}
```

上面的代码中，arr 是一个二维数组，子数组中的每一个项包含键和值，只有这样的数组类型才可以转化为对象。

### 4.2 Map 转化为对象

`Object.fromEntries()` 可以直接将 `Map` 结构的对象转化为一个正常的对象，在开发中是一个非常常用的方法。

```javascript
var map = new Map();
map.set('name', 'imooc');
map.set('age', 7);
console.log(map);     // Map(2) {"name" => "imooc", "age" => 7}
var obj = Object.fromEntries(map);
console.log(obj);     // {name: "imooc", age: 7}
```

上面的代码中，先定义一个 `Map` 数据结构，并给 `map` 添加 `name` 和 `age` 两个属性，然后使用 `Object.fromEntries()` 方法对 `map` 进行操作，最后可以得到一个对象。

## 4. 小结

本节主要讲解了 ES6 提供了获取可枚举对象上的键值对的方法，这个方法在开发中很常用，主要用于 `for...of` 循环和 Map 数据结构的转化。另外，`Object.fromEntries()` 是 `Object.entries()` 的反转函数，他们的操作是相反的，主要用在对 Map 数据结构转换为普通的对象结构。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
