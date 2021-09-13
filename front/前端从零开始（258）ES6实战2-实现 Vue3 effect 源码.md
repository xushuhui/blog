# ES6 实战 2- 实现 Vue3 effect 源码

## 1. 前言

上一节我们实现了 Vue3 的数据劫持功能，并对一些边界值做了处理。但是，当数据改变了我们希望更新试图，这个时候虽然我们能劫持到数据的变化但是没有做任何处理，我们需要对数据的获取和修改增加更新的逻辑，并提供一个 API 给业务用来响应式的处理数据的变化。Vue3 中提供了 effect，当 effect 回调函数中引用的响应式数据变化时，会触发 effect 回调函数的执行，相当于 vue2 中的 watcher。我们来看下面的应用示例：

```javascript
// /vue-next/public/index.html
<script src="./vue.reactivity.js"></script>
<script>
  const { reactive, effect } = VueReactivity;
  const proxy = reactive({
    name: 'ES6 Wiki',
  })

  effect(() => {
    document.getElementById('app').innerHTML = proxy.name;
  })

  setTimeout(() => {
    proxy.name = 'imooc ES6 Wiki 实战'
  }, 1000)
</script>
```

上面的代码中我们引入了 Vue3 的 reactivity 库，初始化网页内容后，在 1 秒以后更新网页中的内容。本节我们就来实现 effect 这个 API 的功能，本节源码参考: [ES6 Wiki](https://github.com/fex-teaching/ES6-wiki/tree/master/packages/vue-next/src/reactivity-2) 。

## 2. effect 实现

### 2.1 创建响应式 effect

effect 在 Vue3 的响应式系统中是一个非常关键的函数，后面的 ref、computed 等函数都会用到 effect 中的功能。在 Vue3 中的 effect 会接受不了两个参数：

```javascript
effect(fn, options)
```

基于 Vue3 响应式 API 的 effect 特点，需要将 effect 变成一个响应式函数，effect 的响应式就是当数据变化时 fn 会自动执行。实现 effect 这个函数的一个目标就是，将 effect 回调函数中所有引用了响应式数据的属性收集起来，并和 effect 的回调函数关联上，在数据变化时在执行 effect 的回调函数。也就是上面的测试案例中，proxy 对象的 name 属性在 effect 的回调函数中。要想让 effect 成为响应式的，就需要将 name 和 effect 关联起来，当 name 的值变化了，就执行 effect 的回调函数。

在本节 options 没用到，但是在 computed 中会使用到，本节使用了 `options.lazy` 属性，用于判断是否在第一次的时候执行回调函数中的内容。effect 中是默认执行回调函数的。

如果要把 effect 变成响应式，需要定义一个创建响应式的方法（createReactiveEffect）用于创建一个 effect 函数。createReactiveEffect 执行后会返回一个 effect 函数，在 createReactiveEffect 函数中会默认执行 fn。

```javascript
export function effect(fn, options){
  const effect = createReactiveEffect(fn, options)
  if (!options.lazy) {
    effect()
  }
  return effect
}

function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    return fn();	// 用户创建的回调函数，fn函数内部会对响应式数据进行取值操作
  }
  return effect
}
```

我们定义一个全局变量 activeEffect，这样做是为了把 effect 存起来，方便后面调用，在取值的时候就可以拿到这个 activeEffect。

```javascript
let activeEffect;
function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    activeEffect = effect;
    return fn();
  }
  return effect
}
```

### 2.2 属性和 effect 关联

怎么才能让属性和这个函数进行关联呢？首先我们要创建一个收集函数（track）用于收集属性 key 和 effect 回调函数的关联，并且只有在 effect 中使用到的 key，更新时才会执行 effect 中的回调，所以我们在收集依赖时需要先判断。

```javascript
function track(target, key) {
  if (activeEffect === viod 0) {
    return;
  }
}
```

什么时候进行收集呢？effect 回调函数会默认执行，在获取值的时候对响应式对象上的 key 进行依赖收集，也就是在 createGetter 函数中进行收集。

```javascript
function createGetter() {
  return function get(target, key, receiver) {
    const res = Reflect.get(target, key, receiver);
    if (isSymbol(key)) {
      return res;
    }

    // 依赖收集
    track(target, key);

    if (isObject(res)) {
      return reactive(res);
    }
    return res;
  };
}
```

如何关联呢？就是需要在 target 上的 key 中存放若干个 effect，那这要怎么存放呢？这时我们想到了 WeakMap，创建一个 WeakMap 来保持 target 上的需要关联 effect 的属性。同时，

下面的伪代码数据结构是我们希望存放在 WeakMap 中的映射，其中 target 是目标对象。

```javascript
{
  target1: {
    key: [effect, effect]
  },
  target2: {
    key: [effect, effect]
  }
}
```

在存放 effect 时可能还需要给 effect 加上一些标识，如：id、deps、options 等，后面会用到。

```javascript
Let uid = 0;
function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    activeEffect = effect;
    return fn();
  }
  effect.id = uid++;
  effect.deps = [];
  effect.options = opntions;
  return effect
}

const targetMap = new WeakMap();
function track(target, key) {
  if (activeEffect === undefined) {
    return;
  }
  // 目标是创建一个映射：{target1: {name: [effect, effect]},target2: {name: [effect, effect]}}
  let depsMap = targetMap.get(target);	// depsMap存放target的值，是一个Map对象
  if(!depsMap) {	// 如果targetMap中没用target对象，则创建一个。
    targetMap.set(target, (depsMap = new Map()));
  }
  let dep = depsMap.get(key);	// 获取depsMap对象中属性是target上的key值
  if(!dep) {
    depsMap.set(key, (dep = new Set())); // 存放effect的集合
  }
  if(!dep.has(effect)) {
    dep.add(activeEffect);
    activeEffect.deps.push(dep);
  }
}
```

上面的代码中，收集目标对象上所有的依赖，在 effect 的回调函数中没有使用到的属性，就不需要进行依赖收集。在执行完创建响应式 effec 函数 createReactiveEffect 后需要把 activeEffect 置为 null。

```javascript
function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    try {
      activeEffect = effect;
    	return fn();
    } finally {
      activeEffect = null;
    }
  }
  return effect
}
```

上面的代码中 finally 是一定会执行的。在 effect 回调函数中嵌套使用 effect，并且在嵌套的 effect 后还有响应式数据，如果是下面这种写法，`state.c = 300` 将不会收集。

```javascript
effect(() => {
  state.a = 100;
  effect(() => {
    state.b = 200;
  })
  state.c = 300;
})
```

这个时候我们就需要创建一个存放栈的数组（effectStack）来存放 activeEffect，执行完毕后也不用赋值 null 了，通过出栈的形式把最后一个移除，让当前的 activeEffect 值等于 effectStack 最后一个值 `effectStack[effectStack.length-1]` 。这样我们在执行完创建响应式 effect 函数时，控制权又会交到上一层的 activeEffect 上，这样上面代码中的 `state.c=300` 就会被收集到第一层的 effect 中去。具体执行代码如下：

```javascript
const effectStack = [];
function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    try {
      activeEffect = effect;
      effectStack.push(activeEffect);
    	return fn();
    } finally {
      effectStack.pop();
      activeEffect = effectStack[effectStack.length - 1];
    }
  }
  return effect
}
```

使用栈的还有一个好处可以防止递归执行，在 effect 如果有数据持续变化是如： `state.a++` 这样的逻辑就会形成递归。这时需要处理为只执行一次，增加一个条件判断，如下代码：

```javascript
function createReactiveEffect(fn, options) {
  const effect = function reactiveEffect() {
    if (!effectStack.includes(effect)) {	// 防止死循环
      try {
        activeEffect = effect;
        effectStack.push(activeEffect);
        return fn();
      } finally {
        effectStack.pop();
        activeEffect = effectStack[effectStack.length - 1];
      }
    }
  }
  return effect
}
```

### 2.3 执行收集的函数

上面的内容是依赖收集的过程，主要在响应式数据获取时执行，也就是在调用 createGetter 的时候执行，那么依赖收集完后，当数据发生变化的时候，需要让收集的回调函数依次执行。而执行这样收集函数的过程是在 createSetter 中完成，因为在这里是更新数据的过程。上节中我们在 createSetter 中预留了新增和更新属性的判断：

```javascript
function createSetter() {
  return function get(target, key, value, receiver) {
		...
    if (!hadKey) {
      console.log('新增属性');
      trigger(target, 'ADD', key, value)
    } else if (hasChanged(value, oldValue)) {
      console.log('更新属性');
      trigger(target, 'SET', key, value, oldValue)
    }

    return result;
  };
}
```

Vue3 中执行依赖的函数是 trigger，这个函数一共接受五个参数，在执行 trigger 时会传入修改数据的类型：新增（ADD）和更新（SET），这是 Vue 为了处理不同场景而设置的属性。这里我们先创建 tigger 函数，首先需要判断在 targetMap 中是否有被依赖的对象，没有则直接返回。

```javascript
export function trigger(target, type, key, newValue, oldValue) {
  const depsMap = targetMap.get(target)
  if (!depsMap) {
    return
  }
}
```

如何让依赖的 effect 执行呢？

* 首先要判断 key 是不是 undefined；
* 获取 key 中的 effect 函数，并执行。

```javascript
export function trigger(target, type, key, newValue, oldValue) {
  const depsMap = targetMap.get(target)
  if (!depsMap) {
    return
  }
  const run = (effects) => {
    if (effects) {
      effects.forEarch(effect => effect())
    }
  }
  if (key == void 0) {
    run(depsMap.get(key));
  }
}
```

上面是对对象的处理，但是在处理数组的时候还会有问题，如下代码：

```javascript
const state = reactive([1,2,3]);
effect(() => {
  document.getElementById('app').innerHTML = state[2];
})

setTimeout(() => {
  state.length = 1;
}, 1000)
```

上面的代码中，数据变化是直接更新数组的长度，而在 effect 中没有使用 length 属性，所以在更新 length 属性时不会触发 `run(depsMap.get(key));` 的依次执行，这样 length 改变 effect 回调函数不会执行，视图也不会被更新。这时就需要对属性是 length 的数组进行验证，如果直接更新的是数组的长度就需要单独处理：

```javascript
export function trigger(target, type, key, newValue, oldValue) {
  const depsMap = targetMap.get(target)
  if (!depsMap) {
    return
  }
  const run = (effects) => {
    if (effects) {
      effects.forEarch(effect => effect())
    }
  }
  if (key === 'length' && isArray(target)) {
    depsMap.forEarch((deps, key) => {
      if(key === 'length' || key >= newValue) {	// newValue是更新后的值，
        run(deps)
      }
    })
  } else {
    if (key == void 0) {
      run(depsMap.get(key));
    }
  }
}
```

上面的代码是在修改数组 length 属性时，让收集依赖的函数执行。还有一种情况，是在 effect 回调中没有直接取索引的值，而且在修改数组时，直接在超过数组长度的位置上新增一个元素。

```javascript
const state = reactive([1,2,3]);
effect(() => {
  document.getElementById('app').innerHTML = state;
})

setTimeout(() => {
  state[5] = 5;
}, 1000)
```

在这种情况下也没有索引 key 进行收集，但是确实使用数组的索引增加了值。这时我们就需要借助 trigger 中的 type 类型来进行处理，当对数组索引进行添加操作时，需要触发数组的更新。

```javascript
export function trigger(target, type, key, newValue, oldValue) {
  const depsMap = targetMap.get(target)
  if (!depsMap) {
    return
  }
  const run = (effects) => {
    if (effects) {
      effects.forEarch(effect => effect())
    }
  }
  if (key === 'length' && isArray(target)) {
    depsMap.forEarch((deps, key) => {
      if(key === 'length' || key >= newValue) {	// newValue是更新后的值，
        run(deps)
      }
    })
  } else {
    if (key == void 0) {
      run(depsMap.get(key));
    }
    switch (type) {
      case 'ADD':
        if(isArray(target)) {
          if(isIntergerKey) {	// 判断key是否是索引类型
            run(depsMap.get('length'));	// 新增属性时直接触发length收集的依赖即可
          }
        }
        break;
    }
  }
}
```

这样我们就基本上实现了 effect 的响应式的源码。

## 小结

本节我们主要实现了 Vue3 中 effect 函数，它是一个响应式的函数，在源码实现过程中需要注意几点：

* 使用 WeakMap 数据结构来存放 target 上的 key 和 effect 的关系；
* 对 effect 的嵌套处理时，引入了栈的方式来控制当前的 activeEffect 值；
* 在使用数组时，在对 length 直接修改等操作时进行特殊的处理。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
