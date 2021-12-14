# ES6+ WeakMap

## 1. 前言

前面我们已经学了 Set 和对应的 WeakSet，Map 对应也有 WeakMap，在学习 WeakSet 时我们已经接触到弱引的相关知识，本节我们将结合 WeakMap 深入的理解弱引用的相关问题。

由前面学到的 Set 和 WeakSet 具有很多相似的地方，比如它们存放的都是独一无二的元素。所以，对 Map 和 WeakMap 也可以进行类比，WeakMap 中也存放的是键值对。不同的是 WeakMap 的 key 只能是对象，值可以是任意类型的，和 WeakSet 一样 WeakMap 对 key 的引用是弱引用。

## 2. WeakMap 基本用法

WeakMap 像 Map 一样可以接受一个二维数组进行初始化。

```javascript
var wm = new WeakMap([
  [{name: 'imooc'}, 'imooc'],
  [{name: 'lesson'}, 'ES6 Wiki']
])
console.log(wm)
```

上面的代码打印结果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f4c8bab09d6716a12720376.jpg)

从打印的结果可以大概了解 WeakMap 的存储方式，WeakMap 的实例本来就是一个对象。

WeakMap 只提供了四个方法用于操作数据。

|方法名|描述|
|------|----|
| set   | 接收键值对，向 WeakMap 实例中添加元素       |
| get   | 传入指定的 key 获取 WeakMap 实例上的值      |
| has   | 传入指定的 key 查找在 WeakMap 实例中是否存在|
| delete| 传入指定的 key 删除 WeakMap 实例中对应的值  |

看如下实例：

```javascript
var wm1 = new WeakMap();
var wm2 = new WeakMap();
var wm3 = new WeakMap();

var o1 = {name: 'imooc'};
var o2 = function(){};
var o3 = window;

// 使用 set 方法添加元素，value 可以是任意值，包括对象、函数甚至另外一个WeakMap对象
wm1.set(o1, 'ES6 Wiki');
wm1.set(o2, 10);
wm2.set(o1, o2);
wm2.set(o3, null);
wm2.set(wm1, wm2);

wm1.get(o2); // 10
wm2.get(o2); // undefined，wm2 中没有 o2 这个键
wm2.get(o3); // null

wm1.has(o2); // true
wm2.has(o2); // false
wm2.has(o3); // true (即使值是null)

wm3.set(o1, 'lesson is ES6 Wiki!');
wm3.get(o1); // lesson is ES6 Wiki!

wm1.has(o1);   // true
wm1.delete(o1);
wm1.has(o1);   // false
```

上面的实例基本涵盖了 WeakMap 四种方法的基本使用情况，上面也提到了 WeakMap 的 key 只能是对象类型的，如果 WeakMap 的 key 是基本类型数据时就会报错。

```javascript
var wm = new WeakMap();
wm.set('lesson', 'ES6 Wiki');
// Uncaught TypeError: Invalid错误value used as weak map key
```

上面代码中在设置 wm 值时，报错了。从报错类型知道是一个类型错误，弱引用映射的键是无效的。

## 3. WeakMap 使用场景

上节我们学习了 Map 的使用，在 JavaScript 中对对象的引用都是强保留的，这意味着只要持有该对象的引用，垃圾回收机制就不会回收该对象。

```javascript
var obj = {a: 10, b: 88};
```

上面是一个字面量对象，只要我们访问 obj 对象，或者任何地方有引用该对象，这个对象就不会被垃圾回收。而在 ES6 之前 JavaScript 中没有弱引用概念，弱引用的本质上就是不会影响垃圾回收机制。其实，WeakMap 并不是真正意义上的弱引用，只要键仍然存在，它就强引用其上的内容。WeakMap 仅在键被垃圾回收之后，才弱引用它的内容，所以也不用太纠结其中的弱。

在官方上对为什么使用 WeakMap 做了描述，Map 在存储值是有顺序的，这种顺序是通过二维数组的形式来完成的。我们知道 Map 在初始化时接受一个数组，数组中的每一项也是一个数组，这个数组中包含两个值，一个存放的是键，一个存放的是值。新添加的值会添加到数组的末尾，从而使得键值具有索引的含义。在取值时就需要进行遍历，通过索引取出对应的值。

但是这样存在两个很大的缺陷：

1. 赋值和搜索的时间复杂度都是 O (n) （n 是键值对的个数），因为这两个操作都是要遍历整个数组才能完成的；
2. 可能会导致内存泄漏，因为数组会一直引用每个键和值。这种引用使得垃圾回收算法不能回收处理它们，即使没有任何引用存在。

相比之下，原生的 WeakMap 持有的是 “弱引用”，这意味着它不会影响垃圾回收。WeakMap 中的 key 只有在键值存在的情况才会引用，而且只是一个读取操作，并不会对引用的值产生影响。也正因为这样的弱引用关系，导致 WeakMap 中的 key 是不可枚举的，假设 key 是可枚举的，就会对该值产生引用关系，影响垃圾回收。

如果只是单纯地向对象上添加值用于检查某些逻辑判断，又不想影响垃圾回收机制，这个时候就可以使用 WeakMap。这里说一点，在一些框架中已经使用了像 WeakMap 和 WeakSet 这样的数据结构，其中 Vue3 就引入了这样的新数据进行一些必要的逻辑判断，有兴趣的可以去扒扒 Vue3 的源码研究研究。

## 4. 总结

本节主要介绍了 WeakMap 的使用和应用场景，这里要说明的一点是：WeakMap 不算真正意义上的弱引用方式，只要键仍然存在，它就强引用其上的内容。最新的 ES 方案提出了 WeakRef 的 API 作为真正的弱引用方式，现在还处于不稳定期间，也还存在一些问题，如果有兴趣的可以研究一下。最后，在 WeakMap 的使用上，大多数都是用来进行一些必要的逻辑判断的。在 WeakMap 实例上添加一个对已知对象的引用，从而在需要使用时，对该对象进行必要的逻辑判断。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
