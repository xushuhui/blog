# JavaScript instanceof

> instanceof 运算符用于检测构造函数的 prototype 属性是否出现在某个实例对象的原型链上。(MDN)

instanceof 是另一种检测类型的手段，但通常用于检测对象之间的关系，如某个对象是不是由某个构造函数生成的。

```javascript
function Person(name) {
  this.name = name;
}

var person = new Person('小明');

console.log(
  person instanceof Person,
); // 输出：true
```

## 1. 语法

```javascript
对象 instanceof 构造函数;
```

虽然语法是这样的，其实左侧可以是任意数据类型，但右侧必须是一个函数。

否则会报如下错误：

```javascript
[] instanceof {};

// Uncaught TypeError: Right-hand side of 'instanceof' is not callable
```

错误大致意思是 `instanceof` 的右操作数不能被调用。

在 `JavaScript` ，可被调用的目前只有函数。

## 2. 注意点

使用 `instanceof` 检测的时候，不一定只有一个为 `true` 的结果。

```javascript
function Person(name) {
  this.name = name;
}

var person = new Person('小明');

console.log(
  person instanceof Person,
  person instanceof Object,
); // 输出：true
```

因为 `instanceof` 实际上是去`左操作数的原型链上寻找有没有右操作数的原型`。

`person` 的原型链上既匹配到 `Person.prototype` 又能匹配到 `Object.prototype`，所以都能返回 `true`。

使用的时候要注意这个问题，如判断某个对象的原型链上是否有 `Object.prototype` 的时候，要考虑到一些其他对象。

```javascript
[] instanceof Object; // true
```

数组的原型链上也是有 `Object.prototype` 的，所以做一些检测的时候要考虑一些特殊情况。

## 3. 小结

instanceof 可以用来检测对象和构造函数之间的关系，其检测的原理是左操作数的原型上是否有右操作数的 `prototype` 属性，所以要注意一些检测的特殊情况。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
