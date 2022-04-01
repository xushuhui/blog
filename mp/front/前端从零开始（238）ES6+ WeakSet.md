# ES6+ WeakSet

## 1. 前言

上一节我们学习了 Set 数据结构，本节将学习与它类似的数据结构 WeakSet，不同的是 WeakSet 存放的数据是一个弱引用类型。在 JavaScript 中，对象的引用是强保留的，这意味着只要持有对象的引用，它就不会被垃圾回收。JavaScript 属于高级语言， 存在 GC 机制，不需要直接地去操作内存，避免了很多问题。同时也让一些内存泄漏的问题变得更加不易察觉，所以 ES6 引入了 WeakSet 和 WeakMap 这样存储弱引用类型的数据结构，是不会阻止它被垃圾回收的。

## 2. 基本用法

WeakSet 对象允许你将一个弱引用对象报错在一个集合中。和 Set 一样，它们都是构造函数，都需要实例化才能使用。WeakSet 可以接收一个可迭代对象作为参数，则该对象的所有迭代值都会被自动添加进生成的 `WeakSet` 对象中。null 被认为是 undefined。

```javascript
const ws = new WeakSet([iterable]);
```

WeakSet 对数据的操作方法相对 Set 是比较少的，只有添加、删除和查找，而且不能被遍历。

|方法名|描述|
|------|----|
|add   |向 WeakSet 实例中添加值          |
|delete|删除 WeakSet 实例中的指定值      |
|has   |查找指定的值是否在 WeakSet 实例中|

WeakSet 存放的一般都是对象的引用，如下实例：

```javascript
var ws = new WeakSet();
var obj1 = {};
var obj2 = {};

ws.add(obj1);
ws.add(obj2);

ws.has(obj1);    // true
ws.has(obj2);   // true

ws.delete(obj1); // 从 set 中删除 obj1 对象
ws.has(obj1);    // false, obj1 对象已经被删除了
ws.has(obj2);    // true, obj2 依然存在
```

## 3. 对比 Set

`WeakSet` 对象是一些对象值的集合，并且其中的每个对象值都只能出现一次。在 `WeakSet` 的集合中是唯一的，这和 Set 对象是一样的。

但 `WeakSet` 和 `Set` 还是有明显的区别的，主要区别有两点：

* 与 `Set` 相比，`WeakSet` 只能是对象的集合，而不能是任何类型的任意值，这个 Set 可以存储任何类型的值不同；
* `WeakSet` 是弱引用：`WeakSet` 对象中存储的对象值都是被弱引用的，如果没有其他的变量或属性引用这个对象值，则这个对象值会被当成垃圾回收掉。正因为这样，`WeakSet` 对象是无法被枚举的，没有办法拿到它包含的所有元素。

当我们去取一个对象进行操作时，外界必然存在对这个对象的引用，否则我们不可能取到这个对象。而弱引用到底是做什么的呢？

首先我们要知道对象的生命周期是，只有对象存在引用就不会被 GC 回收。但有时候我们只是需要这个集合去判断一些逻辑，如果使用 Set 对象的话，就会存在引用，这样实例化的内容就不会被回收。这时，使用 WeakMap 就是有必要的事了。让我们来看下面的这个实例：

```javascript
const requests = new WeakSet();
class ApiRequest {
  constructor() {
    requests.add(this);
  }

  makeRequest() {
    if(!request.has(this)) throw new Error("Invalid access");
    // do work
  }
}
```

上面的代码中，ApiRequest 想验证一下 this 的来源，使用 WeakMap 来存储这个 this 对象，并在 makeRequest 执行时去验证一下是否是 ApiRequest 这个类。这里的 requests 实例并不想参与到 ApiRequest 类中的生命周期中去，它只是作为一个条件判断使用的。如果使用 Set 对象的话，这个实例就会在 ApiRequest 类中存在引用关系，并一直保存在实例中，增加内存的开销，也可能会发生内存泄漏。而使用 WeakMap 则不同它是弱引用，只有在劫持的时候才会被获取到。

## 4. 小结

本节学习了 WeakMap 对象，它用了存储对象的值，存储的值都是弱引用类型。其实，弱引用就是将对象键添加到 WeakSet 上，而 WeakSet 对对象的引用不会影响垃圾回收，也就是说，你持有对象的引用，就可以获取元数据。一旦不再持有对象的引用，即使你仍持有并添加了该对象的引用，也会被垃圾回收。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
