# ES6+ Set

## 1. 前言

上一节我们对 ES6 新的数据结构做了简要的介绍，并没有深入的学习。这节我们将深入地学习 Set 数据结构，了解 Set 的存储方式，可以让我们更深入地理解 Set 数据结构。

## 2. 基本用法

`Set` 提供了用于操作数据的几种方法，使用这些方法可以对 `Set` 上的数据进行增、删、查等操作，具体有以下 5 种方法：

|方法名|描述|
|------|----|
| add   | 向 Set 实例中添加值          |
| delete| 删除 Set 实例中的指定值      |
| clear | 清空 Set 实例                |
| has   | 查找指定的值是否在 Set 实例中|
| size  | 返回集合存储数据的长度       |

`Set` 不像数组可以用索引去查找存储的数据，所以 `Set` 没有修改操作。`Set` 存储的数据都是唯一的，而且可以是任意类型的数据。

在实例化一个 Set 数据结构时，可以接受一个可遍历的对象（一般为数组），作为默认值，返回一个 `Set` 实例：

```javascript
var set = new Set([1, 2, 3])
console.log(set) // Set(3) {1, 2, 3}
```

也可以使用 add 方法去添加元素，并且 add 方法支持链式调用：

```javascript
// 创建一个空的 Set 实例
var set = new Set()        // Set(0) {}
// 添加数据
set.add('es6')             // Set(1) {"es6"}
// 还可以链式添加数据
set.add('imooc').add({age: '7'})
cosnole.log(set)	// 如下图
```

![图片描述](http://img.mukewang.com/wiki/5f428414098e82e705180276.jpg)

上面的例子中，对 `Set` 实例添加元素，可以添加引用类型数据。`Set` 的存储方式是以对象的形式存放的，所以 `Set` 的本质也就是一个对象，它继承 Object，我们可以通过原型链（ `__proto__` ）找到最上层是 Object。

`Set` 还提供了一些方法用于删除、查找和清空操作：

```javascript
// 创建一个带默认值的 Set 实例
var set = new Set([1, 2, 3]);
// delete 删除数据: 删除指定的数据，或者清空数据
console.log(set.delete(3));
// set.size 可以统计多少条数据，类似数组中的length属性
console.log(set.size);        // 2
// clear 会清空Set实例
console.log(set.clear());    // Set(0) {}
// has 方法可以查看某个元素是否包含在 Set 实例中
console.log(set.has('a'));    // false
console.log(set.has(1));    // true
```

## 2. 案例：数组去重

`Set` 一个最常见的操作就是对数组进行去重操作，这也是它诞生的一个主要功能之一。

在 ES5 中数组去重的方法很多，一种是使用 for 循环来对数组中的项进行一一验证，另一种是比较高效的方法就是借助对象的键是唯一的特性进行去重操作。下面我们来实现这两种方法，并对比 Set 方式去重。

### 2.1 使用 for 循环去重

使用 for 循环去重比较简单，建立一个空数组，然后循环目标数组，如果当前循环的元素不在这个新建的数组中，就 push 到这个数组中。下面给出具体的代码：

```javascript
function unique(arr) {
    const newArr = []
    for (let i = 0; i < arr.length; i++) {
        if (newArr.indexOf(arr[i]) === -1) {
            newArr.push(arr[i])
        }
    }
    return newArr
}
```

这种方法比较简单也是最常见的方法，虽然这里只有一次循环，但是 `indexOf` 的查询也是有时间复杂度的，所以不能单纯地认为这里的时间复杂度是 n。另外，由于 `indexOf` 存在一些缺陷这里也可以使用 [`includes()`](http://www.imooc.com/wiki/ES6lesson/arrayincludes.html) 进行替换。

### 2.2 使用 hash 方法去重

使用 hash 方法去重是利用对象的键是唯一的，维护一个以数组中元素为 hash 表的键，由于键是唯一的，所以数组中相同的元素在 hash 表中只会有一个键。

```javascript
  function unique(arr){
    const newArr = [];
    const hash = {};
    for(var i=0; i<arr.length; i++){
      if(!hash[arr[i]]){
        hash[arr[i]] = true;
        newArr.push(arr[i]);
      }
    }
    return newArr;
  }
```

上面的代码时间复杂度是 n，只需要对数组进行一次循环即可。把循环的元素存放在 hash 表中来记录不重复的元素。如果 hash 表中找不到对应的值则在 hash 表中添加一个记录，并把该元素 push 到数组中。这样的方式时间复杂度为 n，但是维持一个 hash 表需要更多的空间。

### 2.3 使用 Set 方法去重

我们知道 Set 是一个集合，其中的元素不能重复，所以可以利用 Set 数据结构的特点进行去重操作。下面给出具体代码：

```javascript
function unique(arr) {
    return [...new Set(arr)]
}
```

上面的代码可以看出，上节我们学习了 Set 和数组直接的转换，可以使用 `...` 语法展开 Set 实例就可以得到数组。当然还可以使用 `Array.from()` 方法把 Set 数据结构转化为数组。这种方式在低版本浏览器是不能运行的。

其实上面三种方式都有一定的缺点：

* 第一种方式，时间复杂度高；
* 第二种方式，空间复杂度高；还有一个致命的缺点是，如果数组中的元素是对象形式，那么就不能使用此方法。因为对象的 key 只能是字符串，其解决方式可以使用 Map 数据结构代替 hash 的存储；
* 第三种方式，需要更高级的浏览器。

## 3. Set 的扩展方法

Set 实例上的数据可以使用 `forEach` 进行遍历：

```javascript
var set = new Set(['a', 'b', 'c'])

set.forEach((item) => {
    console.log(item)
})
// a
// b
// c
```

虽然在 Set 数据结构中没有索引的概念，但是 Set 提供了三个方法用于我们去遍历 Set 数据结构。

|方法名|描述|
|------|----|
| values | 得到 Set 实例中的值作为一个可以遍历的对象            |
| keys   | 得到 Set 实例中的键作为一个可以遍历的对象，键和值相等|
| entries| 得到 Set 实例中的键值作为一个可以遍历的对象          |

通过这三种方法得到的是一个可迭代对象（后面的迭代器小节会具体介绍），看如下实例：

```javascript
var set = new Set(['a', 'b', 'c'])

set.keys();    // SetIterator {"a", "b", "c"}
set.values();  // SetIterator {"a", "b", "c"}
set.entries(); // SetIterator {"a" => "a", "b" => "b", "c" => "c"}

for(var [key, value] of set.entries()) {
    console.log(key, value)
}
// a, a
// b, b
// c, c
```

从上面的代码中，使用这三个方法得到的数据可以被 `for...of` 进行遍历。而且这三个方法得到的数据是一个可以迭代的对象，对象上的键和值是相等的，这也是为什么 Set 中的数据是唯一的原因。

## 4. 小结

本节我们深入地学习了 Set 数据结构的基本用法和它提供的三个扩展方法。另外，通过案例数组去重，深入了解了 Set 数据结构的用途。 Set 数据结构的本质还是一个键值对的对象，只不过键和值是相等的，这也是 Set 集合中元素不能重复的原因之一。并且 Set 是一个构造函数，需要实例化一个 Set 数据结构才能使用。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
