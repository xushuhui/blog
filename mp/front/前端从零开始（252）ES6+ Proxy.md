# ES6+ Proxy

## 1. 前言

本节我们将学习 ES6 的新增知识点 ——Proxy，Proxy 是代理的意思。Proxy 是一个对象，用于定义基本操作的自定义行为（如属性查找、赋值、枚举、函数调用等）。这是 MDN 上的定义，但是不容易理解，想要理解 Proxy 我们首先需要知道什么是代理？

在日常开发中比较常见的代理常见有，使用 Charles 代理抓包、nginx 服务器的反向代理，以及 VPN 等，都用到了代理。什么是代理呢？我们先看一张图：

![What is a proxy server? - Seobility Wiki](https://www.seobility.net/en/wiki/images/8/8a/Proxy-Server.png)

上图是客户端访问网络的示意图，客户端不能直接访问网络，它只能先访问代理服务器，只有代理服务器才能有权限访问，然后代理服务器把客户端请求的信息转发给目标服务器，最后代理服务器在接收到目标服务器返回的结果再转发给客户端，这样就完成了整个请求的响应过程。这是现在大多数服务器的架构，我们可以把上图的 Proxy Server 理解为 Nginx。代理有正向代理和反向代理，有兴趣的小伙伴可以去深入了解一下。

本节说的 Proxy 就是作用在 JavaScript 中的一种代理服务，代理的过程其实就是一种对数据的劫持过程，Proxy 可以对我们定义的对象的属性进行劫持，当我们访问或设置属性时，会去调用对应的钩子执行。在 ES5 中我们曾学习过 `Object.defineProperty()` 它的作用和 Proxy 是相同的，但是 `Object.defineProperty()` 存在一些性能问题，Proxy 对其进行了升级和扩展更加方便和易用。本节我们将学习 Proxy 的使用。

## 2. Object.defineProperty()

在学习 Proxy 之前，我们先来回归一下 ES5 中的 `Object.defineProperty()` ，接触过前端框架的同学应该都知道 Vue 和 React，其中 Vue 中的响应式数据底层就是使用 `Object.defineProperty()` 这个 API 来实现的。下面是 `Object.defineProperty()` 的语法。

```javascript
Object.defineProperty(obj, prop, descriptor)
```

`Object.defineProperty()` 会接收三个参数：

* obj 需要观察的对象；
* prop 是 obj 上的属性名；
* descriptor 对 prop 属性的描述。

当我们去观察一个对象时需要在 descriptor 中去定义属性的描述参数。在 descriptor 对象中提供了 get 和 set 方法，当我们访问或设置属性值时会触发对应的函数。

```javascript
var obj = {};
var value = undefined;

Object.defineProperty(obj, "a", {
  get: function() {
    console.log('value:', value)
    return value;
  },
  set: function(newValue) {
    console.log('newValue:', newValue)
    value = newValue;
  },
  enumerable: true,
  configurable: true
});
obj.a;	// value: undefined
obj.a = 20;	//  newValue: 20
```

上面的代码中，我们使用一个变量 value 来保存值，这里需要注意的是，不能直接使用 obj 上的值，否则就会出现死循环。

`Object.defineProperty()` 是 Vue2 的核心， Vue2 在初始化时会对数据进行劫持，如果劫持的属性还是对象的话需要递归劫持。下面我们把 Vue2 中数据劫持的核心代码写出来。

```javascript
var data = {
  name: 'imooc',
  lession: 'ES6 Wiki',
  obj: {
    a: 1
  }
}

observer(data);


function observer(data) {
  if (typeof data !== 'object' || data == null) {
    return;
  }

  const keys = Object.keys(data);

  for (let i = 0; i < keys.length; i++) {
    let key = keys[i];
    let value = obj[key];
    defineReactive(obj, key, value);
  }
}

function defineReactive(obj, key, value) {
  observer(value);

  Object.defineProperty(obj, key, {
    get() {
      return value;
    },
    set(newValue) {
      if (newValue === value) return;
      observer(newValue);
      value = newValue;
    }
  })
}
```

上面代码的核心是 defineReactive 方法，它是递归的核心函数，用于重新定义对象的读写。从上面的代码中我们发现 `Object.defineProperty()` 是有缺陷的，当观察的数据嵌套非常深时，这样是非常耗费性能的，这也是为什么现在 Vue 的作者极力推广 Vue3 的原因之一，Vue3 的底层使用了 Proxy 来代替 `Object.defineProperty()` 那 Proxy 具体有什么好处呢？

## 3. Proxy

首先我们来看下 Proxy 是如何使用的，语法：

```javascript
const p = new Proxy(target, handler)
```

Proxy 对象是一个类，需要通过 new 去实例化一个 Proxy 对象，它接收的参数比较简单，只有两个：

* target：需要使用 Proxy 进行观察的目标对象；
* handler：对目标对象属性进行处理的对象，包含了处理属性的回调函数等。

```javascript
const handler = {
	get: function(obj, prop) {
    return obj[prop];
  },
  set: function(obj, prop, value) {
    return obj[prop] = value;
  }
};

const p = new Proxy({}, handler);
p.a = 1;

console.log(p.a, p.b);      // 1, undefined
```

对比上面的 `Object.defineProperty()` API 直观的看 Proxy 做了一些精简，把对象、属性和值作为 get 和 set 的参数传入进去，不必考虑死循环的问题了。这是直观的感受。

上面我们使用了 `Object.defineProperty()` API 简单地实现了 Vue2 的响应式原理，那么 Vue 使用 Proxy 是怎么实现的呢？它带来了哪些好处呢？下面我们看实现源码：

```javascript
var target = {
  name: 'imooc',
  lession: 'ES6 Wiki',
  obj: {
    a: 1
  }
}
var p = reactive(target);
console.log(p.name);		// 获取值: imooc
p.obj.a = 10;						// 获取值: {a : 1}
console.log(p.obj.a);		// 获取值: {a : 10}


function reactive(target) {
  return createReactiveObject(target)
}

function createReactiveObject(target) {
  // 判断如果不是一个对象的话返回
  if (!isObject(target)) return target

  // target观察前的原对象； proxy观察后的对象：observed
  observed = new Proxy(target, {
    get(target, key, receiver) {
      const res = target[key];
      console.log('获取值:', res)
      // todo: 收集依赖...
      return isObject(res) ? reactive(res) : res
    },
    set(target, key, value, receiver) {
      target[key] = value;
    }
  })

  return observed
}
```

上面的代码是从 Vue3 中摘出来的 reactive 函数的实现，我们可以直观地看到没有对 target 进行递归循环去创建观察对象。而且，当我们对 obj 下的 a 属性设置值时，执行 get 函数，这是为什么呢？这就是 Proxy 的优点，在对 obj 下属性设置值时，首先需要调用 set 方法获取 target 下 obj 的值，然后判断 obj 又是一个对象再去调用 reactive 函数进行观察。这样就不需要递归地去对嵌套数据进行观察了，而是在获取值的时候，判断获取的值是不是一个对象，这样极大地节约了资源。

## 4. 小结

本节主要通过代理和 `Object.defineProperty()` API 的学习来理解 ES6 的新增知识点 ——Proxy，并且通过 Vue2 和 Vue3 实现响应式原理来对比 `Object.defineProperty()` 和 Proxy 的优缺点，从而更深入地理解 Proxy。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
