# ES6+ Object.values()

## 1. 前言

[上一节](https://www.imooc.com/wiki/ES6lesson/keys.html) 我们学习了 `Object.keys()` 可以获取对象上可枚举属性的 key 作为一个数组，对应的是不是有一个让我们获取对象上所有属性值的方法呢？答案是肯定的，ES6 中提供了 `Object.values()` 获取可枚举对象上属性的值作为一个数组。

## 2. 方法详情

`Object.values()` 方法返回一个给定对象自身的所有可枚举属性值的数组，值的顺序与使用 `for...in` 循环的顺序相同 （区别在于 for-in 循环枚举原型链中的属性）。与 `Object.keys()` 相反一个是获取 key，一个是获取 value，其他的使用特性基本相同。

### 2.1 语法使用

```javascript
Object.values(obj);
```

|参数|描述|
|----|----|
|obj|可以返回其枚举自身属性和键的对象|

### 2.2 基本使用

```javascript
var obj = {a: 1, b: 2, c: 3};
Object.values(obj)     // [1, 2, 3]
```

## 3. 实例

### 3.1 参数是数值或布尔值

[上一节](https://www.imooc.com/wiki/ES6lesson/keys.html) 我们知道 `Object.keys()` 的参数会被类型转换为对象，数值对象没有可遍历的对象，所有返回一个空数组。同样，`Object.values()` 和 `Object.keys()` 一样都会进行类型转换，所以传入的是数字或布尔值时，则返回一个空数组

```javascript
console.log(Object.values(123));	// []
Object.values(false)    // []
```

### 3.2 键排序问题

[上一节](https://www.imooc.com/wiki/ES6lesson/keys.html) 我们学习 `Object.keys()` 会对属性是数值的键进行排序，在这个过程中属性对应的值也会跟着改变位置，所有使用 `Object.values()` 返回的数组是按 `Object.keys()` 顺序后的结果展示的，所以得到的值要和排序后的属性一一对应。

```javascript
var obj = {10: 'a', 1: 'b', 7: 'c'};
Object.values(obj)    // ['b', 'c', 'a']
```

### 3.3 参数为字符串

`values()` 传入的对象如果是一个字符串时，则会把字符拆成数组中的单个项，如下：

```javascript
Object.values('abc')    // ['a', 'b', 'c']
```

## 4. 小结

本节主要讲解了 ES6 提供了获取可枚举对象上的属性值的方法，这个方法可以很方便地获取对象上的属性值用于遍历。提高代码的简洁性。这里需要注意的是在获取属性值时会对属性为数值类型的键进行排序，所以对应的属性值也会跟着一起被排序，所以得到的结果会有所不同。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
