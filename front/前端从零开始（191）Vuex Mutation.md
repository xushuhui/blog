# Mutation

## 1. 前言

本文我们将介绍如何使用 `mutation`。包括如何定义 mutation、如何触发 mutation、mapMutations 辅助函数的使用方式。`mutation` 是更改 Vuex 中 store 数据状态的唯一方法。在 `vuex` 的使用过程中，我们需要编写大量的 `mutation` 来操作 store 中的数据。所以，学好如何使用 mutation 非常重要。`mutation` 并不是一个难点，它的使用非常简单，接下来我们就一步步学习它的使用。

## 2. 基础用法

### 2.1. 定义 mutation

Vuex 中的 mutation 非常类似于事件：每个 mutation 都有一个字符串的 **事件类型 (type)** 和 一个 **回调函数 (handler)**。这个回调函数就是我们实际进行状态更改的地方，并且它会接受 state 作为第一个参数：

```javascript
const store = new Vuex.Store({
  state: {
    count: 1
  },
  mutations: {
    increment (state) {
      // 变更状态
      state.count++
    }
  }
})
```

### 2.2. 触发 mutation

我们不能直接调用一个 mutation handler。这个选项更像是事件注册：“当触发一个类型为 increment 的 mutation 时，调用此函数。” 要唤醒一个 mutation handler，你需要以相应的 type 调用 store.commit 方法：

```javascript
store.commit('increment')
```

### 2.3. 提交载荷（Payload）

你可以向 store.commit 传入额外的参数，即 mutation 的 载荷（payload）：

```javascript
mutations: {
  incrementByCount (state, n) {
    state.count = state.count + n
  }
}
```

```javascript
store.commit('incrementByCount', 10)
```

在大多数情况下，载荷应该是一个对象，我们通常接收的参数命名为 `payload`，这样可以包含多个字段并且记录的 mutation 会更易读：

```javascript
// 定义 mutation
mutations: {
  incrementByCount (state, payload) {
    state.count = state.count  + payload.count
  }
}
// 触发 mutation
store.commit('incrementByCount', {
  count: 10
})
```

### 2.3. 对象风格的提交方式

提交 mutation 的另一种方式是直接使用包含 type 属性的对象：

```javascript
store.commit({
  type: 'incrementByCount',
  count: 10
})
```

当使用对象风格的提交方式，整个对象都作为载荷传给 mutation 函数，因此 handler 保持不变：

```javascript
mutations: {
  incrementByCount (state, payload) {
    state.count = state.count  + payload.count
  }
}
```

完整示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app">
    <div>购物车数量：{{count}}</div>
    <button @click="add">无参 mutation（+1）</button>
    <button @click="addTen">携参 mutation（+10）</button>
    <button @click="addByObject">携带对象类型的参数 mutation（+5）</button>
    <button @click="commitByObject">对象类型提交 mutation（+3）</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      count: 0
    },
    mutations: {
      // 无参 mutation
      increment(state) {
        state.count++
      },
      // 带参 mutation
      incrementByCount(state, count) {
        state.count = state.count + count
      },
      // 对象类型参数 mutation
      incrementByObject(state, payload) {
        state.count = state.count + payload.count
      },
    },
  })
  var vm = new Vue({
    el: '#app',
    store,
    methods: {
      add() {
        this.$store.commit('increment')
      },
      addTen() {
        this.$store.commit('incrementByCount', 10)
      },
      addByObject() {
        this.$store.commit('incrementByObject', {
          count: 5
        })
      },
      commitByObject() {
        this.$store.commit( {
          type: 'incrementByObject',
          count: 3
        })
      }
    },
    computed: {
      count() {
        return this.$store.state.count
      }
    }
  })
</script>
</html>

```

**代码解释**

JS 代码第 9-11 行，定义了一个无参的 mutation，对 state.count + 1。

JS 代码第 12-14 行，定义一个传入 Number 类型参数 count 的 mutation，对 state.count + count。

JS 代码第 16-18 行，定义一个传入 Object 类型参数 payload 的 mutation，对 state.count + payload.count。

JS 代码第 26 行，提交 mutation increment。

JS 代码第 28 行，提交 mutation incrementByCount 并传入数量 10。

JS 代码第 31-33 行，提交 mutation incrementByObject 并传入参数 {count: 5}。

JS 代码第 36-39 行，以对象的形式提交 mutation incrementByObject。

## 3. mutation 使用注意事项

### 3.1. Mutation 需遵守 Vue 的响应规则

既然 Vuex 的 store 中的状态是响应式的，那么当我们变更状态时，监视状态的 Vue 组件也会自动更新。这也意味着 Vuex 中的 mutation 也需要与使用 Vue 一样遵守以下注意事项：

1. 最好提前在你的 store 中初始化好所有所需属性。

2. 当需要在对象上添加新属性时，你应该：

    * 使用 Vue.set (obj, ‘newProp’, 123), 或者
    * 以新对象替换老对象。例如，利用对象展开运算符我们可以这样写：

```javascript
    state.obj = { ...state.obj, newProp: 123 }
```

> Tips：以新对象替换老对象替换老对象的方式只能修改 `state` 中的某个属性，而不能替换整个 `state`。想要替换整个 `state`，需要使用 store.replaceState () 的方法：

```javascript
  state.obj = { ...state.obj, newProp: 123 } // OK
  state = {...state, name: '123'} // Error
  store.replaceState({...state, name: '123'}) // OK
```

### 3.2. 使用常量替代 Mutation 事件类型

在日常开发中，我们一般会使用常量替代 mutation 事件类型。这样可以使 linter 之类的工具发挥作用，同时可以让你的代码合作者对整个 app 包含的 mutation 一目了然：

```javascript
 const INCREMENT_COUNT = 'INCREMENT_COUNT'
 const store = new Vuex.Store({
   state: {
     count: 0
   },
   mutations: {
     [INCREMENT_COUNT](state) {
       state.count++
     },
   },
 })
 var vm = new Vue({
   el: '#app',
   store,
   methods: {
     add() {
       this.$store.commit('INCREMENT_COUNT')
     },
   }
 })
```

当然，是否使用用常量取决于个人喜好 —— 在需要多人协作的大型项目中，这会很有帮助。但如果你不喜欢，你完全可以不这样做。

### 3.3 Mutation 必须是同步函数

一条重要的原则就是要记住 mutation 必须是同步函数。为什么？请参考下面的例子：

```javascript
mutations: {
  someMutation (state) {
    api.callAsyncMethod(() => {
      state.count++
    })
  }
}
```

现在想象，我们正在 debug 一个 app 并且观察 devtool 中的 mutation 日志。每一条 mutation 被记录，devtools 都需要捕捉到前一状态和后一状态的快照。然而，在上面的例子中 mutation 中的异步函数中的回调让这不可能完成：因为当 mutation 触发的时候，回调函数还没有被调用，devtools 不知道什么时候回调函数实际上被调用 —— 实质上任何在回调函数中进行的状态的改变都是不可追踪的。

## 4 mapMutations 辅助函数

mapMutations 辅助函数帮助我们简化提交 mutation 的写法。

### 4.1 mapMutations 接收数组格式的参数

`mapMutations` 可以接收一个事件类型 (type) 的数组：

```javascript
...mapMutations([
  // 将 `this.increment()` 映射为 `this.$store.commit('increment')`
  'increment'
]),
```

### 4.2 mapMutations 接收对象格式的参数

在某些情况，我们需要对 `Mutation` 中的函数名重命名以避免和组件内部的变量冲突，这时候我们可以使用对象的方式接收参数：

```javascript
...mapMutations({
  [别名]: [Mutation type]
})
// 例：将 `this.add()` 映射为 `this.$store.commit('increment')`
...mapMutations({
  add: 'increment'
})
```

示例：

```javascript
const INCREMENT_COUNT = 'INCREMENT_COUNT'
const store = new Vuex.Store({
  state: {
    count: 0
  },
  mutations: {
    [INCREMENT_COUNT](state) {
      state.count++
    },
  },
})
var vm = new Vue({
  el: '#app',
  store,
  methods: {
    ...Vuex.mapMutations([INCREMENT_COUNT]),
    ...Vuex.mapMutations({
      add: INCREMENT_COUNT
    })
  }
})
```

## 5. 小结

本小节我们介绍了如何使用 `Mutation` 提交事件来修改 `state` 中的数据。主要知识点有以下几点：

* 在 store 中定义 Mutation 事件。
* 通过 $store.commit 触发 Mutation 事件。
* 通过 mapMutations 方法简化提交 Mutation 的写法。

其中，使用 Mutation 需要注意以下几点：    1. Mutation 必须是一个同步函数。
    2. Mutation 需遵守 Vue 的响应规则：只能通过 Vue.set 添加 state 中的属性，只能通过 store.replaceState 替换这个 state。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
