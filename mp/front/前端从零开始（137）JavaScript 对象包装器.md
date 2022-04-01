# 对象包装器

对象包装器也被称为 `包装器`、`包装对象`.

所有包装器都是内置对象，如 `Number`、`String`、`Boolean` 等。

## 1. 装箱

通常在对一个变量赋值的时候，都会直接给定一个字面量。

```javascript
var string = '996 say no!';

console.log(typeof string); // 输出："string"
```

通过类型的检测，可以知道他是一个字符串。

但又可以访问到一些属性，比如 `length`：

```javascript
var string = '996 say no!';

console.log(typeof string); // 输出："string"
console.log(string.length); // 输出：11
```

照道理讲只有对象才能访问到属性，字符串字面量只是一个值而已。

这里就是因为 `JavaScript` 内部的拆装箱的机制。

当把一个字面量像对象一样操作的时候，`JavaScript` 会进行装箱操作。

可以把上面这份代码理解成下面这份：

```javascript
var string = '996 say no!';

console.log(typeof string); // 输出："string"
console.log((new String(string)).length); // 输出：11
```

其中的 `new String` 就是装箱操作，`String` 就是字符串的对象包装器。

这样将字符串转换成了对象，就能访问到其属性了。

需要注意的是，对一个字面量包装后不会修改原始值，上述例子中的 `string` 变量的值依然是字符串字面量，不会变成对象，所以每一次对一个字面量做访问属性或方法的操作时，都会做一次装箱操作。

许多开发者会考虑频繁装箱的性能影响，其实通常是不必要的，一是现在计算机和浏览器的处理执行速度很快了，几乎可以忽略不计装箱的开销，另外就是业务开发中在没有明确的性能要求下，是不考虑性能的，以完成业务逻辑为主。

## 2. 拆箱

拆箱操作很多时候是`隐式转换`过程中发生的。

如将字符串进行对象相等操作：

```javascript
var obj = {
  toString: function() {
    return '996';
  },
};

console.log(
  '996' == obj,
); // 输出：true
```

上述例子，在字符串字面量 `996` 和对象 `obj` 进行比较的时候，会尝试将右侧的对象转化成字符串，即调用 `toString` 方法（在 ES6 中则会先看有没有部署 Symbol.toPrimitive 方法）。

这个即使拆箱的过程。

## 3. 小结

理解对象包装器的作用，可以更好的理解内置对象的用途。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
