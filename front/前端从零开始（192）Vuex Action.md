# Action

## 1. 前言

本小节我们将介绍如何使用 `Action`。包括如何定义 Action、分发 Action、mapActions 辅助函数的使用方式。Action 在 Vuex 中会大量使用，学好如何使用 Action 非常重要。`Action` 并不是一个难点，它的使用非常简单，接下来我们就一步步学习它的使用。

## 2. Action 简介

`Action` 类似于 `Mutation`，不同的是：

*     1. Action 提交的是 mutation，而不是直接变更状态。

*     1. Action 可以包含任意异步操作。在 `vuex` 的使用过程中，我们可以将多个 `Mutation` 合并到一个 `Action` 中，也可以通过 `Action` 进行异步操作。

## 3. 基础用法

### 3.1 定义 action

Action 函数接受一个与 store 实例具有相同方法和属性的 context 对象，因此你可以调用 context.commit 提交一个 mutation，或者通过 context.state 和 context.getters 来获取 state 和 getters。

```javascript
const store = new Vuex.Store({
  state: {
    count: 1
  },
  mutations: {
    increment (state) {
      state.count++
    }
  },
  actions: {
    // 同步 action
    increment (context) {
      context.commit('increment')
    },
    // 异步 action
    incrementAsync (context) {
      setTimeout(() => {
        context.commit('increment')
      }, 1000)
    }
  }
})
```

实践中，我们会经常用到 ES2015 的参数解构来简化代码（特别是我们需要调用 commit 很多次的时候）：

```javascript
actions: {
  increment ({ commit }) {
    commit('increment')
  }
}
```

### 3.2 分发 Action

Action 通过 store.dispatch 方法触发：

```javascript
store.dispatch('increment')
```

### 3.3 提交载荷（Payload）

你可以向 store.dispatch 传入额外的参数，即 Actions 的 载荷（payload）：

```javascript
action: {
  increment ({commit}, payload) {
    // 具体 action 内容
  }
}
```

```javascript
store.dispatch('increment', {count: 10})
```

### 3.4 对象风格的提交方式

提交 action 的另一种方式是直接使用包含 type 属性的对象：

```javascript
store.dispatch({
  type: 'increment',
  count: 10
})
```

当使用对象风格的提交方式，整个对象都作为载荷传给 action 函数，因此 handler 保持不变：

```javascript
actions: {
  increment ({commit}, payload) {
    // 具体 action 内容
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
    <button @click="add">同步 +1</button>
    <button @click="addAsync">1s后 +1</button>
    <button @click="addAsyncParams">2s后 +1</button>
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
      increment(state) {
        state.count++
      }
    },
    actions: {
      increment ({commit}) {
        commit('increment')
      },
      incrementAsync ({commit}) {
        setTimeout(() => {
          commit('increment')
        }, 1000)
      },
      incrementAsyncParams ({commit}, payload) {
        setTimeout(() => {
          commit('increment')
        }, payload.time)
      }
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    methods: {
      add() {
        this.$store.dispatch('increment')
      },
      addAsync() {
        this.$store.dispatch({
          type: 'incrementAsync',
        })
      },
      addAsyncParams() {
        this.$store.dispatch('incrementAsyncParams', {
          time: 2000
        })
      },
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

JS 代码第 9-11 行，我们定义了 mutation 事件 increment，事件对 state.count + 1。

JS 代码第 15-17 行，我们定义了同步 Action increment，Action 中直接提交事件 increment。

JS 代码第 18-22 行，我们定义了异步 Action incrementAsync，1 秒后提交事件 increment。

JS 代码第 23-27 行，我们定义了接收参数的异步 Action incrementAsyncParams。

JS 代码第 35 行，分发 Action 事件 increment。

JS 代码第 38-40 行，以对象的形式分发 Action 事件 incrementAsync。

JS 代码第 43-45 行，分发 Action 事件 incrementAsyncParams，并传入对应参数。

## 4 mapActions 辅助函数

mapActions 辅助函数帮助我们简化提交 action 的写法。

### 4.1 mapActions 接收数组格式的参数

`mapActions` 可以接收一个 action 事件名的数组：

```javascript
...mapActions([
  // 将 `this.increment()` 映射为 `this.$store.dispatch('increment')`
  'increment'
]),
```

### 4.2 mapActions 接收对象格式的参数

在某些情况，我们需要对 `Action` 中的函数名重命名以避免和组件内部的变量冲突，这时候我们可以使用对象的方式接收参数：

```javascript
...mapActions({
  [别名]: [Action name]
})
// 例：将 `this.add()` 映射为 `this.$store.dispatch('increment')`
...mapActions({
  add: 'increment'
})
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
    <button @click="increment">添加</button>
    <button @click="add">别名添加</button>
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
      increment(state) {
        state.count++
      },
    },
    actions: {
      increment({commit}) {
        commit('increment')
      }
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    methods: {
      ...Vuex.mapActions([
        'increment'
      ]),
      ...Vuex.mapActions({
        add: 'increment'
      })
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

JS 代码第 23-25 行，我们通过 mapActions 将 this.increment () 映射为 this.$store.dispatch (‘increment’)。

JS 代码第 26-28 行，我们通过 mapActions 将 this.add () 映射为 this.$store.dispatch (‘increment’)。

## 5. 组合 Action

Action 通常是异步的，有时候我们需要知道 action 什么时候结束，并在结束后进行相应的其他操作。更重要的是，我们可以组合多个 action，以处理更加复杂的异步流程。

首先，你需要明白 store.dispatch 可以处理被触发的 action 的处理函数返回的 Promise，并且 store.dispatch 仍旧返回 Promise：

```javascript
actions: {
  actionA ({ commit }) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        commit('someMutation')
        resolve()
      }, 1000)
    })
  }
}
```

现在我们可以：

```javascript
store.dispatch('actionA').then(() => {
  // ...
})
```

在另外一个 action 中也可以：

```javascript
actions: {
  // ...
  actionB ({ dispatch, commit }) {
    return dispatch('actionA').then(() => {
      commit('someOtherMutation')
    })
  }
}
```

最后，如果我们利用 async /await，我们可以如下组合 action：

```javascript
// 假设 getData() 和 getOtherData() 返回的是 Promise
actions: {
  async actionA ({ commit }) {
    commit('increment', await getData())
  },
  async actionB ({ dispatch, commit }) {
    await dispatch('actionA') // 等待 actionA 完成
    commit('increment', await getOtherData())
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
    <button @click="addAsync">添加</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  function getData() {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        console.log('getData success')
        resolve()
      }, 1000)
    })
  }
  function getOtherData() {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        console.log('getOtherData success')
        resolve()
      }, 2000)
    })
  }
  const store = new Vuex.Store({
    state: {
      count: 0
    },
    mutations: {
      increment(state) {
        state.count++
      }
    },
    actions: {
      async actionA ({ commit }) {
        commit('increment', await getData())
      },
      async actionB ({ dispatch, commit }) {
        await dispatch('actionA') // 等待 actionA 完成
        commit('increment', await getOtherData())
      }
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    methods: {
      addAsync() {
        this.$store.dispatch('actionB')
      },
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

JS 代码第 4-19 行，我们定义函数 getData 和 getOtherData。

JS 代码第 29-31 行，定义 actionA，当 getData 函数执行完成之后 commit increment 事件。

JS 代码第 32-35 行，定义 actionB，当 dispatch (actionA) 执行完成之后 commit increment 事件。

## 6. 小结

本小节我们介绍了如何使用 `Action` 来操作 `mutation` 或者进行异步操作。主要知识点有以下几点：

* 在 store 中定义 Action 事件。
* 通过 $store.dispatch 分发 Action 事件。
* 通过 Action 处理异步操作、合并处理 Mutation。
* 使用 mapActions 方法简化分发 Action 的写法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
