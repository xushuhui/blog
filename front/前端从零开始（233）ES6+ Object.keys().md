## ES6+ Object.keys()

## 1. 前言

我们知道迭代对象可以使用 `for...in` 循环来做，但 `for...in` 循环会枚举其原型链上的属性，这使得我们在遍历时需要判断是不是原型链属性。`Object.keys()` 可以接受一个对象返回一个可枚举的数组，数组中的元素的排列顺序和使用 `for...in` 循环遍历返回的顺序是一致的。

`Object.keys()` 在 ES5 中就有此方法，但是在设计上存在一定的缺陷，ES6 对其底层做了重大的更新。比如：在 ES5 中，如果此方法的参数不是对象（而是一个原始值），那么它会抛出 TypeError。在 ES2015 中，非对象的参数将被强制转换为一个对象。

```javascript
// ES5 代码
Object.keys("imooc");  // TypeError: "imooc" is not an object
// ES6 代码
Object.keys("imooc");  // ["0", "1", "2", "3", "4"]
```

现在的浏览器已经基本都支持 ES6 的结果了，下面我们来系统性地认识一下 `Object.keys()`。

## 2. 方法详情

### 2.1 基本语法

`Object.keys()` 方法会返回一个由一个给定对象的自身可枚举属性组成的数组，数组中属性名的排列顺序和正常循环遍历该对象时返回的顺序一致 。

**语法使用：**

```javascript
Object.keys(obj)
```

**参数解释：**

|参数|描述|
|----|----|
|obj|要返回其枚举自身属性的对象。|

### 2.2 基本实例

用于对象时返回对象键值作为数组：

```javascript
var obj =  {
	name:  'imooc',
	type:  'ES6 Wiki'
}
console.log(Object.keys(obj));
// ["name", "type"]
```

用于数组类型：

```javascript
var arr = ['a',  'b',  'c'];
console.log(Object.keys(arr));
// console: ['0', '1', '2']
```

也可以用于类数组中：

```javascript
var obj =  {  0:  'a',  1:  'b',  2:  'c'  };
console.log(Object.keys(obj));  // ['0', '1', '2']
```

键值是数字和字符串混合时，会先进行数值的排序，然后再按添加的顺序排列字符串：

```javascript
var obj =  {  name:  'imooc',  10:  'a',  3:  'b',  age:  7  };
console.log(Object.keys(obj));
// ["3", "10", "name", "age"]
```

`Object.keys()` 不能获取不可枚举属性：

```javascript
// 创建一个obj对象带有一个不可枚举属性
var obj =  Object.create({},  {
	getFoo:  {
		value:  function  ()  {  return  this.foo;  }
	}
});
obj.foo =  1;
console.log(Object.keys(obj));  // ['foo']
```

## 3. 自动排序问题

在说自动排序问题前，我们先来看下三个例子：

```javascript
var obj1 =  {99:  '九十九',  5:  '五',  7:  '七'}
Object.keys(obj1) // ["5", "7", "99"]

var obj2 =  {c:  'z',  a:  'x',  b:  'y'}
Object.keys(obj2) // ["c", "a", "b"]

var obj3 =  {  name:  'imooc',  10:  'a',  3:  'b',  age:  7  };
Object.keys(obj3);  // ["3", "10", "name", "age"]
```

上面的例子可以看出当键值是数字时返回的值会自动排序，即使在混合情况下也会先进行排序后把数字项放在数组中前面，而键值对是字符串时则不会被排序。那当 `Object.keys()` 被调用时内部都发生了什么呢？

通过查阅 ECMA262 规范知道，`Object.keys` 在内部会根据对象的属性名 `key` 的类型进行不同的排序逻辑。分三种情况：

1. 如果属性名的类型是 `Number`，那么 `Object.keys` 返回值是按照 `key` 从小到大排序；
2. 如果属性名的类型是 `String`，那么 `Object.keys` 返回值是按照属性被创建的时间升序排序；
3. 如果属性名的类型是 `Symbol`，那么逻辑同 `String` 相同。

那内部到底发生了什么呢？

### 3.1 将参数转换为对象

在 `Object.keys()` 调用时会根据传入的参数进行类型转换，转换为 Object 类型的值：

|参数类型|结果|
|--------|----|
|Undefined|抛出 TypeError           |
|Null     |抛出 TypeError           |
|Boolean  |返回一个新的 Boolean 对象|
|Number   |返回一个新的 Number 对象 |
|String   |返回一个新的 String 对象 |
|Symbol   |返回一个新的 Symbol 对象 |
|Object   |直接将 Object 返回       |

实例：

```javascript
// 参数是undefined或null
Object.keys(undefined) // Uncaught TypeError: Cannot convert undefined or null to object
// 参数是数值
Object.keys(123) // []
// 参数是字符串
Object.keys('imooc') // ["0", "1", "2", "3", "4"]
```

上面的代码中，参数是数值时为什么会返回空数组呢？是因为数值转换为对象时没有可提取的属性，而字符串在 ES5 时会报错，ES6 进行了修复，因为 String 对象有可提取的属性。看下面两张图：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f153d3d0824406102880074.jpg)

![图片描述](https://xushuhui.gitee.io/image/imooc/5f153d65087cc81302870170.jpg)

## 3.2 获取属性列表

上面我们说到了 `Object.keys()` 会对参数做类型转换，在获取属性的时候会调用内部方法 `EnumerableOwnProperties ( O, kind )` 来计算对象上所有可枚举的属性 `ownKeys`，这里的 `ownKeys` 类型时 `list` 类型，只用于内部实现。

然后声明变量用于存放遍历对象后得到的属性集合 `properties`，`properties` 也是 List 类型，循环对象的 `ownKeys` 将每个元素添加到 `properties` 列表中。最后返回 `properties`。

为什么会对数值进行排序，是因为在调用 `EnumerableOwnProperties(O, kind)` 方法执行时，又会调用 `OrdinaryOwnPropertyKeys(O)` ，对于不同类型的属性，会按不同的顺序放入 `properties` 属性列表中：

1. 先处理类型为数值的属性，**从小到大**放到属性列表中；
2. 再处理类型为字符串的属性，按该属性的**创建顺序**，放到属性列表中；
3. 最后处理类型为 `Symbol` 的属性，按**创建顺序**，放到属性列表中。

这样就知道为什么会对数值进行排序了，是 ECMA262 中 `OrdinaryOwnPropertyKeys(o)` 规定的。其原因是 `OrdinaryOwnPropertyKeys(o)` 内部方法不只是给 `Object.keys()` 使用的，是通用的规则。

最后将 `properties` 列表转化为数组就得到了 `Object.keys()` 的结果。

## 4. 小结

本节主要学习了 `Object.keys()` 方法用于获取对象上可枚举属性，并返回属性的数组，数组中的元素的排列顺序和使用 `for...in` 循环遍历返回的顺序是一致的。这里需要注意的是，如果对象上的属性是数值时，会被排序。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
