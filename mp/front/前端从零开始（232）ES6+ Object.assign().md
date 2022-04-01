# ES6+ Object.assign()

## 1. 前言

Object 对象的更新 把一个对象复制到另一个对象，在 ES5 中需要循环对象进行拷贝

## 2. 方法详情

### 2.1 基本语法

**语法使用：**

```javascript
Object.assign(target, ...sources)
```

**参数解释：**

|参数|描述|
|----|----|
|target |需要拷贝到的目标对象|
|sources|源对象              |

### 2.2 拷贝对象

```javascript
let target = {};
let source = {a: 1, b: 2, c: 3};
Object.assign(target, source);
target.d = 4;
console.log(target)   // {a: 1, b: 2, c: 3, d: 4}
console.log(source)   // {a: 1, b: 2, c: 3}
```

上面的代码可以看出，Object.assign () 的主要用法就是把源对象拷贝到指定的对象上去，目标对象的更改不会影响源对象。

### 2.3 合并对象

```javascript
let target = {a: 1};
let source1 = {b: 2};
let source2 = {c: 3};
Object.assign(target, source1, source2);
console.log(target);  // {a: 1, b: 2, c: 3}
```

上面的代码可以看出，Object.assign () 不会把目标对象清空，会合并后面所有的对象上的值。

### 2.4 覆盖前面的值

```javascript
let target = {a: 1, b: 1};
let source1 = {b: 2, c: 2};
let source2 = {c: 3};
Object.assign(target, source1, source2);
console.log(target);  // {a: 1, b: 2, c: 3}
```

如果后面的源对象上有相同的值，后面的源对象会覆盖前面对象上的值，这一点需要注意。

## 3. 浅拷贝

`Object.assign()` 的拷贝属于浅拷贝，也就是说它只拷贝对下的第一层的属性值。如果这个值是一个对象类型，那么 `Object.assign()` 不会对该对象进行深拷贝，也就是说，拷贝后的对象下的这个对象类型是源对象和拷贝后的对象共有的，无论谁（源对象或拷贝后对象）对这个对象下的值进行修改，另一个对象（源对象或拷贝后对象）也会共享这个改变。看下面的例子更清晰的表达：

```javascript
var target = {};
var source = {a: 1, b: {c: 2, d: 3}};
Object.assign(target, source);
target.a = 5;
target.b.c = 9;
console.log(target)   // {a: 5, b: {c: 9, d: 3}}
console.log(source)   // {a: 1, b: {c: 9, d: 3}}
```

上面的代码中，源对象 source 是个两层的字面量对象，b 也是一个对象。使用 `Object.assign()` 拷贝给目标对象 target，拷贝后对 target 对象下的值进行修改，然后打印目标对象和源对象。从打印的结果可以看出，对 target 第一层的 a 进行修改时，源对象是不会改变。但是对 target 下的 b 对象下的值进行修改时，因为 b 也是一个对象，所以源对象中的值也被修改了。到这里可以看出，`Object.assign()` 没有对 b 进行拷贝。

如果需要深拷贝则需要，需要递归地使用去 `Object.assign()` 来拷贝对象。

## 4. 基本类型的合并

当合并的源对象是基本类型时，这些基本类型会作为最后对象上的值，而键则以数字递增，其中如果值是 null 和 undefined 时会被忽略。看如下实例：

```javascript
var s1 = "abc";
var s2 = true;
var s3 = 10;
var s4 = Symbol("foo")
var obj = Object.assign({}, s1, null, s2, undefined, s3, s4);
console.log(obj); // { "0": "a", "1": "b", "2": "c" }
```

## 5. 拷贝异常时会被打断

在拷贝时如果发生异常，则拷贝会被终止，并报错，前面已经被拷贝的不会被影响可以继续使用，但后面没有被拷贝的则不能被使用。

```javascript
var target = Object.defineProperty({}, "a", {
  value: 1,
  writable: false
});
Object.assign(target, {b: 2}, {a: 3}, {c: 4});
// Uncaught TypeError: Cannot assign to read only property 'a' of object '
console.log(target.b);  // 2
console.log(target.c);  // undefined
```

上面的代码中，定义了目标对象 target 上的属性 a 是只读的，也就是不能不被修改，在合并代码时，源对象上有 a，则报了 a 是对象上的只读属性不能被 assign 操作。从后面的打印结果可以看出，b 已经被拷贝到目标对象上了可以正常使用，但由于拷贝中发生异常，最后一个对象没有被拷贝，所以 c 的值是 undefined。

## 6. 小结

本章讲解了用于合并对象的方法 `Object.assign()` 主要有以下几点需要注意的。

* Object.assign () 属于浅拷贝；
* 合并对象时后面的对象会覆盖前面的对象；
* 拷贝时发生异常，前面已拷贝的不会受到影响，异常后面的对象则不会被拷贝。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
