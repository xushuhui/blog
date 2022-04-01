# ES6 实战 1- 实现 Vue3 reactive 源码

## 1. 前言

本节开始我们将进入 ES6 实战课程，首先会花费两节的时间来学习 Vue3 响应式原理，并实现一个基础版的 Vue3 响应式系统；然后通过 Promise 来封装一个真实业务场景中的 ajax 请求；最后我们会聊聊前端开发过程中的编程风格。

本实战主要通过对前面 ES6 的学习应用到实际开发中来，Vue3 的响应式系统涵盖了大部分 ES6 新增的核心 API，

如：Proxy、Reflect、Set/Map、WeakMap、Symbol 等 ES6 新特性的应用。更加深入地学习 ES6 新增 API 的应用场景。

由于篇幅内容有限，本实战不会完全实现 Vue3 响应式系统的所有 API，主要实现 `reactive` 、 `effect` 这四个核心 API，其他 API 可以参考 [vue-next](https://github.com/vuejs/vue-next)

源码。本节的目录结构和命名和 Vue3 源码基本一致，在阅读源码的时候我们能看到作者的思考，和功能细颗粒度的拆分，使得代码更易于扩展和复用。

## 2. 环境配置

### 2.1 rollup 配置

ES6 很多 API 不能在低版本浏览器自己运行，另外我们在开发源码的时候需要大量地使用模块化，以拆分源码的结构。在学习模块化一节时，我们使用了 Webpack 作用打包工具，由于 Vue3 使用的是 rollup，更加适合框架和库的大包，这里我们也和 Vue3 看齐，rollup 最大的特点是按需打包，也就是我们在源码中使用的才会引入，另外 rollup 打包的结果不会产生而外冗余的代码，可以自己阅读。下面我们来看下 rollup 简单的配置：

```javascript
// rollup.config.js
import babel from "rollup-plugin-babel";
import serve from "rollup-plugin-serve";

export default {
  input: "./src/index.js",
  output: {
    format: "umd", // 模块化类型
    file: "dist/umd/reactivity.js",
    name: "VueReactivity", // 打包后的全局变量的名字
    sourcemap: true,
  },
  plugins: [
    babel({
      exclude: "node_modules/**",
    }),
    process.env.ENV === "development"
      ? serve({
          open: true,
          openPage: "/public/index.html",
          port: 3000,
          contentBase: "",
        })
      : null,
  ],
};
```

上面的配置内容和 webpack 很相似，是最基础的编译内容，有兴趣的小伙伴可以去了解一下。[本节源码](https://github.com/fex-teaching/ES6-wiki/tree/master/packages/vue-next) 在 ES6-Wiki 仓库的 vue-next 目录下，在这个项目中可以直接启动，在启动前需要在项目根目录中安装依赖。本项目使用的是 yarn workspace 的工作环境，可以在根目录中共享 npm 包。

### 2.2 调试源码

在开发的过程中需要对我们编写的代码进行调试，这里我们在 public 目录中创建了一个 html 文件用于在浏览器中打开。并且引入了 reactivity 的源码可以参考对比我们实现的 API 的功能，同学在使用时可以打开注释进行验证。

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>
<body>
  <div id="app"></div>

	<!-- 我们自己实现的 reactivity 模块 -->
	<script src="/dist/umd/reactivity.js"></script>

  <!-- vue的 reactivity 模块，测试时可以使用 -->
	<!-- <script src="./vue.reactivity.js"></script> -->

  <script>
    const { reactive, effect } = VueReactivity;
    const proxy = reactive({
      name: 'ES6 Wiki',
    })
    document.getElementById('app').innerHTML = proxy.name;
  </script>
</body>
</html>
```

## 3. reactive 实现

在实现 Vue3 的响应式原理前，我们先来回顾一下 Vue2 的响应式存在什么缺陷，主要有以下三个缺陷：

* 默认会劫持的数据进行递归；
* 不支持数组，数组长度改变是无效的；
* 不能劫持不存在的属性。

Vue3 使用了 Proxy 去实现数据的代理，在实现 Vue3 的响应式原理的同时，我们需要思考 Proxy 会不会存在上面的缺陷，它的缺点又是什么呢？

### 3.1 数据劫持

首先我在 reactive 文件中定义并导出一个 reactive 函数，在 reactive 中返回一个创建响应式对象的方法。`createReactiveObject` 函数主要是为了创建响应式对象使用，在 reactive 的相关 API 中，很多都需要创建响应式对象，这样可以复用，而且更加直观。

```javascript
// vue-next/reactivity-1/reactive.js
import { isObject } from "../shared/index";
import { mutableHandlers } from './baseHandlers';

export function reactive(target: object) {
    // 1.创建响应式对象
    return createReactiveObject(target, mutableHandlers)
}

function createReactiveObject(target, baseHandlers) {
    // 3.对数据进行代理
    const proxy = new Proxy(target, baseHandlers);
    return proxy;
}
```

下面的代码是 Proxy 处理对象的回调，包括 get、set、deleteProperty 等回调方法，具体用法可以参考 Proxy 小节内容。这样我们就实现了拦截数据的功能。

```javascript
// vue-next/reactivity-1/baseHandlers.js
function createGetter() {
  return function get(target, key, receiver) {
    console.log('获取值');
    return target[key];
  }
}

function createSetter() {
  return function get(target, key, value, receiver) {
    console.log('设置值');
    target[key] = value;
  }
}

function deleteProperty(target, key) {
  delete target[key];
}

const get = createGetter()
const set = createSetter()

export const mutableHandlers = {
  get,
  set,
  deleteProperty,
  // has,
  // ownKeys
}
```

在 Vue3 源码中使用 Reflect 来操作对象的，Reflect 和 Proxy 方法一一对应，并且 Reflect 操作后的对象有返回值，这样我们可以对返回值做异常处理等，修改上面的代码如下：

```javascript
// vue-next/reactivity-1/baseHandlers.js
function createGetter() {
  return function get(target, key, receiver) {
    console.log('获取值');
    const res = Reflect.get(target, key, receiver);
    return res;
  }
}

function createSetter() {
  return function get(target, key, value, receiver) {
    console.log('设置值');
    const result = Reflect.set(target, key, value, receiver);
    return result;
  }
}
```

下面是测试用例，可以放在 public/index.html 下执行。

```javascript
const { reactive } = VueReactivity;
const proxy = reactive({
  name: 'ES6 Wiki',
})
proxy.name = 'imooc ES6 wiki';	// 设置值
console.log('proxy.name');		// 获取值
// proxy.name
```

### 3.2 实现响应式逻辑

首先我们需要对传入的参数进行判断，如果不是对象则直接返回。

```javascript
// shared/index.js
const isObject = val => val !== null && typeof val === 'object'

function createReactiveObject(target, baseHandlers) {
    if (!isObject(target)) {
        return target
    }
    ...
}
```

在使用时，用户可能多次代理对象或多次代理过的对象，如：

```javascript
var obj = {a:1, b:2};
var proxy = reactive(obj);
var proxy = reactive(obj);
// 或者
var proxy = reactive(proxy);
var proxy = reactive(proxy);
```

像上面这样的情况我们需要处理，不能多次代理。所以我们这里要将代理的对象和代理后的结果做一个映射表，这样我们在代理时判断此对象是否被代理即可。这里的映射我们用到了 WeakMap 弱引用。

```javascript
export const reactiveMap = new WeakMap();

function createReactiveObject(target, baseHandlers) {
  if (!isObject(target)) {
    return target
  }
  const proxyMap = reactiveMap;
  const existingProxy = proxyMap.get(target);
	// 这里判断对象是否被代理，如果映射表上有，则说明对象已经被代理，则直接返回。
  if (existingProxy) {
    return existingProxy;
  }
  const proxy = new Proxy(target, baseHandlers);
  // 这里在代理过后把对象存入映射表中，用于判断。
  proxyMap.set(target, proxy);
  return proxy;
}
```

上面我们已经基本实现了响应式，但是有个问题，我们只实现了一层响应式，如果是嵌套多层的对象这样就不行了。Vue2 是使用的是深层递归的方式来做的，而我们使用了 Proxy 就不需要做递归操作了。Proxy 在获取值的时候会调 get 方法，这时我们只需要在获取值时判断这个值是不是对象，如果是对象则继续代理。

```javascript
import { isSymbol, isObject } from '../shared';
import { reactive } from './reactive';

function createGetter() {
  return function get(target, key, receiver) {
    const res = Reflect.get(target, key, receiver);
    if (isSymbol(key)) {
      return res;
    }
    console.log("获取值，调用get方法"); // 拦截get方法
    if (isObject(res)) {
      return reactive(res);
    }
    return res;
  };
}
```

在获取值的时候有很多边界值需要特殊处理，这里列出了如果 key 是 symbol 类型的话直接返回结果，当然还有其他场景，同学可以去看 Vue3 的源码。

在我们设置值的时候，如果是新增属性时 Vue2 是不支持的，使用 Proxy 是可以的，但是我们需要知道当前操作是新增还是修改？所以需要判断有无这个属性，如果是修改则肯定有值。一般判断有两种情况：

* 数组新增和修改逻辑，需要先进行数组判断，当我们修改的 key 小于数组的长度时说明是修改，反之则是新增；
* 对象新增和修改逻辑，对象判断是否有属性就比较简单了，直接取值验证即可。

```javascript
// 判断数组
export const isArray = Array.isArray;
export const isIntegerKey = key => '' + parseInt(key, 10) === key;	// 判断key是不是整型
// 使用 Number(key) < target.length 判断数组是不是新增，key 小于数组长度说明有key

// 判断对象是否有某属性
const hasOwnProperty = Object.prototype.hasOwnProperty
export const hasOwn = (val, key) => hasOwnProperty.call(val, key)

// 判断有无key
const hadKey = isArray(target) && isIntegerKey(key) ? Number(key) < target.length : hasOwn(target, key);
```

最终我们可以得到下面的 createSetter 函数。

```javascript
function createSetter() {
  return function get(target, key, value, receiver) {
    const oldValue = target[key];	// 获取旧值
    const hadKey = isArray(target) && isIntegerKey(key) ? Number(key) < target.length : hasOwn(target, key);

    const result = Reflect.set(target, key, value, receiver);

    if (!hadKey) {
      console.log('新增属性');
    } else if (hasChanged(value, oldValue)) {
      console.log('修改属性');
    }

    return result;
  };
}
```

## 4. 小结

以上内容就是 Vue3 中实现 reactive API 的核心源码，文章的完整代码放在了 reactivity-1 目录下。源码中的实现方式可能会有所改变，在对照学习时可以参考 Vue 3.0.0 版本。本节实现响应式的核心是 Proxy 对数据的劫持，通过对 set 和 get 方法的实现来处理各种边界数据问题。在学习过程中需要注意多次代理、设置属性时判断是新增还是修改，这对后面实现 effect 等 API 有很重要的作用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
