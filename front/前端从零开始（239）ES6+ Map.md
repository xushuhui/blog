# ES6+ Map

## 1. 前言

前面两节我们学习了 Set 的相关内容，本节我们开始学习 Map 数据结构的内容。Map 对象和原生的 Object 类似都是存储数据的，而且也推荐使用 Map 的方式来存储数据。

在 ES5 中使用 `Object` 来存储对象，然而这种存储存在一些问题，比如说 `Object` 中的键是无序的，`Object` 的键只能是字符串类型等等。ES6 提供了 `Map` 数据结构，保存的是键值对，而且能够记住键的插入顺序，并且任何值 （对象或者原始值） 都可以作为一个键或一个值，这样极大地扩展了数据存储的场景。

## 2. Map 使用详情

### 2.1 Map 基本说明

在前面的 数据结构扩展 一节我们已经了解到 Map 的基本使用。和 `Set` 一样，`Map` 也是一个构造函数，不能直接使用，需要通过 `new` 的方式来创建一个 `Map` 实例。

```javascript
var map = Map([iterable]);
```

`Map` 对象在插入数据时是按照顺序来插入的，也就是说在插入时会保持插入的位置，`Object` 的在插入数据时没有顺序概念。`Map` 对象可以被 `for...of` 循环，在每次迭代后会返回一个形式为 `[key，value]` 的数组，这个我们在下面的例子中会说到。

`Map` 的本质其实还是一个对象，并且它也是继承 Object 的，看下面的实例：

```javascript
var map = new Map([["x", 1], ["y", 2]]);
console.log(map instanceof Object);   //true
```

从上面的代码中可以看出 Object 在实例 map 的原型链上。

在创建 `Map` 实例时可以接收一个数组或是一个可遍历的对象作为参数，这个参数内的每一项是键值的组合 `[key, value]` 第一个值时键，第二个值时键的值。

在初始化 Map 对象时，如果默认参数的数组中超过两个以上的值不会被 `Map` 对象读取。

```javascript
var map = new Map([["x", 1, 'a', 'b'], ["y", 2, 'c'], ["z", 3, 'd']]);
console.log(map)  // Map(3) {"x" => 1, "y" => 2, "z" => 3}
```

上面的代码中，从打印的结果可以看出，数组中超过的元素都会被自动忽略。

### 2.2 Map 的属性和方法

`Map` 提供的属性和方法从增、删、改、查几个方面入手，主要有以下 5 种：

|方法名|描述|
|------|----|
|set   |接收键值对，向 Map 实例中添加元素       |
|get   |传入指定的 key 获取 Map 实例上的值      |
|delete|传入指定的 key 删除 Map 实例中对应的值  |
|clear |清空 Map 实例                           |
|has   |传入指定的 key 查找在 Map 实例中是否存在|
|size  |属性，返回 Map 实例的长度               |

`Map` 提供 `size` 属性可以获取 `Map` 实例上的长度

```javascript
var map = new Map([["x", 1], ["y", 2], ["z", 3]]);
console.log(map.size)  // 3
```

`set()` 方法为 `Map` 实例添加或更新一个指定了键（key）和值（value）的键值对。

```javascript
myMap.set(key, value);
```

通常情况下不会一开始就初始化值，而是动态地添加，或更新 `Map` 时需要用到 `set` 方法，可以新增和修改 `Map` 实例的值。而且 key 值可以是任意类型的，查看如下示例：

```javascript
var map = new Map();

var str = 'string';
var obj = {};
var arr = [];
var fun = function() {};

map.set(str, '键的类型字符串');
map.set(obj, '键的类型对象');
map.set(arr, '键的类型数组');
map.set(fun, '键的类型函数');
```

上面的代码中，我们定义了不同类型的变量，使用这些变量为 `map` 添加数据。相比 `Object` 对象，扩展性更强了。另外还可以链式添加键值对：

```javascript
var map = new Map();
map.set('a', 1).set('b', 2).set('c', 3);
console.log(map);   // Map(3) {"a" => 1, "b" => 2, "c" => 3}
```

使用链式添加键值对的方式比较简洁，如果需要添加多个值，建议使用这样的方式去添加。

`get()` 方法是接收一个指定的键（key）返回一个 `Map` 对象中与这个指定键相关联的值，如果找不到这个键则返回 `undefined`。

```javascript
myMap.get(key);
```

使用上面的示例，可以通过 `get` 方法获取对应的值：

```javascript
console.log(map.get('string'));  // "键的类型字符串"
console.log(map.get(str));       // "键的类型字符串"
console.log(map.get(obj));       // "键的类型对象"
console.log(map.get(arr));       // "键的类型数组"
console.log(map.get(fun));       // "键的类型数组"
```

上面的代码可以看出，我们可以直接使用键的值去获取 Map 实例上对应的值，也可以通过定义变量的方式去获取。

`has()` 方法是用于判断指定的键是否存在，并返回一个 bool 值，如果指定元素存在于 `Map` 中，则返回 `true`，否则返回 `false`。

```javascript
myMap.has(key);
```

**实例：**

```javascript
var map = new Map();
map.set("a", 11);

map.has("a");  // true
map.has("b");  // false
```

`delete()` 方法用于移除 `Map` 实例上的指定元素，如果 `Map` 对象中存在该元素，则移除它并返回 `true`；否则如果该元素不存在则返回 `false`。

```javascript
myMap.delete(key);
```

**实例：**

```javascript
var map = new Map();
map.set("a", 11);

map.delete("a");  // true
map.has("a");     // false
```

`clear()` 方法会移除 Map 对象中的所有元素，返回 `undefined`。

```javascript
myMap.clear(key);
```

**实例：**

```javascript
var map = new Map();
map.set("a", 11);

map.clear();  // 返回 undefined
```

这里需要注意的是 `clear()` 返回的值是 `undefined` 而不是 `true` 所以如果在判断结果的时候需要注意这一点。

### 2.3 Map 的扩展方法

和 `Set` 数据结构一样，Map 也提供了三个获取 Map 对象的键值以及键值对组合的方法：

|方法名|描述|
|------|----|
|values |返回 Map 实例中的值作为一个可以遍历的对象    |
|keys   |返回 Map 实例中的键作为一个可以遍历的对象    |
|entries|返回 Map 实例中的键值对作为一个可以遍历的对象|

`keys()` 方法是获取 `Map` 实例上的键，并返回一个可迭代（Iterator）的对象。

```javascript
myMap.keys();
var map = new Map();
map.set('a', 1);
map.set('b', 2);
map.set('c', 3);

var keys = map.keys()
console.log(keys.next().value);  // "a"
console.log(keys.next().value);  // "b"
console.log(keys.next().value);  // "c"
```

获取后的 `keys` 结构可以被迭代器上的 `next` 函数获取到对应值。

`values()` 方法是获取 `Map` 实例上元素的值，并返回一个可迭代（Iterator）的对象。

```javascript
myMap.values();
```

**实例：**

```javascript
var map = new Map();
map.set('a', 1);
map.set('b', 2);
map.set('c', 3);

var values = map.values()
console.log(values.next().value);  // 1
console.log(values.next().value);  // 2
console.log(values.next().value);  // 3
```

获取后的 `values` 结构可以被迭代器上的 `next` 函数获取到对应值。

`entries()` 方法返回一个包含 `[key, value]` 的可迭代（Iterator）的对象，返回的迭代器的迭代顺序与 `Map` 实例的插入顺序相同。

```javascript
myMap.entries()
```

**实例：**

```javascript
var map = new Map();
map.set('a', 1);
map.set('b', 2);
map.set('c', 3);

var values = map.values()
console.log(values.next().value);  // 1
console.log(values.next().value);  // 2
console.log(values.next().value);  // 3
```

`keys()`、`values()`、`entries()` 都可以被 `for...of` 循环。

```javascript
var map = new Map([["x", 1], ["y", 2], ["z", 3]]);

for (let value of map.values()) {
  console.log(value);
}
// 1
// 2
// 3

for (let [key, value] of map.entries()) {
  console.log(key + " = " + value);
}
// x = 1
// y = 2
// z = 3
```

注意在循环 `entries()` 结果的时候，因为每一项是包含键值的数组，可以通过 `[key, value]` 这种数组结构的方式把键值结构出来直接使用。

## 3. Map 和 Object

`Map` 和 `Object` 有非常多的相似的地方，`Map` 的出现也是为了弥补 `Object` 的不足。 `Object` 的键只能是字符串，`Map` 的键可以是任意类型的值（包括对象），所以 `Map` 是一种更完善的 Hash 结构实现。

### 3.1 键无序问题

Object 的键是无序的，当键可以隐式转换为数值时，在循环的时候就会被优先排序。这也是为什么要求最好不要使用 Number 类型作为对象的属性。

```javascript
var obj = {
  c: 'C',
  3: 3,
  a: 'A',
  1: 1,
}

for (let key in obj) {
  console.log(key, obj[key])
}
// 1 1
// 3 3
// c C
// a A
```

`Map` 会记录插入的顺序，存放的是键值对的组合，并且不会做类型转换。`Map` 可以用 forEach 循环。

```javascript
var map = new Map();
map.set('c', 'C').set(3, 3).set('a', 'A').set(1, 1);
map.forEach((item, key) => {
    console.log(key, item, typeof key)
})
// c C string
// 3 3 "number"
// a A string
// 1 1 "number"
```

从上面的代码中，使用 typeof 去检查 key 的数据类型，可以看出 `Map` 并不会对键做类型转换。

### 3.2 应用场景

Object 不仅是存储数据用的，它还可以有自己的内部逻辑。属性的值是一个函数时，是可以被执行的，并且可以通过 this 拿到对象上的属性。

```javascript
var obj = {
    id: 1,
    desc: "imooc ES6 wiki",
    print: function(){
        console.log(this.desc)
    }
}
console.log(obj.print()); //"imooc ES6 wiki"
```

所以，尽管 Map 相对于 Object 有很多优点，但 Object 在某些场景更易于使用，比如上面的实例。毕竟 Object 是 JavaScript 中最基础的概念，给出使用场景的几个参考。

**使用 Object 的场景：**

1. 如果你知道所有的 key，它们都为字符串或整数（或是 Symbol 类型），而你只需要一个简单的结构去存储这些数据，Object 无疑是一个非常好的选择。另外，获取对象的值时 Object 的性能要优于 Map；
2. 如果需要在对象中保持自己独有的逻辑和属性，比如上面实例，只能使用 Object；
3. JSON 可以直接转为 Object，但 Map 不行。因此，在某些我们必须使用 JSON 的情况下，应将 Object 视为首选。

**使用 Map 的场景：**

* Map 是一个纯哈希结构，使用 `delete` 对 Object 的属性进行删除操作时存在很多性能上的问题。所以，在有大量数据，或是多数据进行增、删操作的场景中，使用 Map 更合适；
* Map 会保留所有元素的插入顺序。在数据迭代时不会有乱序的情况。所以，如果考虑到元素迭代或顺序，使用 Map 更好；
* Map 的键值可以是任意类型，而且不会做类型转换。所以，在键值不确定的情况下，保证键值不被隐式转换的情况下可以优选 Map。

## 4. 判断键值相等的问题

因为在 Map 对象中键可以是任意值，所以对键的说明有以下几点：

1. `NaN` 是与 `NaN` 相等的（虽然 `NaN !== NaN`），剩下所有其它的值是根据 `===` 运算符的结果判断是否相等；
2. 在目前的 ECMAScript 规范中，`-0` 和 `+0` 被认为是相等的，尽管这在早期的草案中并不是这样。

## 5. 小结

本节我们深入地学习了 Map 的使用情况，并对比 Object 给出了几个使用场景的参考。学习完本章你需要知道以下几点：

* 键的类型可以是任意的，可以使用 `function`、对象等等作为 key；
* 键是有顺序，根据添加的顺序决定的；
* Map 是一个完善的 Hash 结构，在存放大数据，或在频繁增、删键值时表现优异。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
