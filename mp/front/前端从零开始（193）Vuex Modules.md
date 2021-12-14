# Modules

## 1. 前言

本节我们将介绍如何将 store 中的数据按模块划分。在复杂的大型项目中，如果将所有的数据都存在一个 state 对象中，那将使得 store 对象变得非常大，难于管理。这时候，使用 module 将变得异常重要。Modules 并非难点，接下来我们就一步步介绍 modules 的使用。

## 2. 如何使用

### 2.1 基本用法

Module 其实是一个对象，它和我们 new Vuex.Store ({…}) 传入的对象格式相同。例如：

```javascript
const moduleA = {
  state: { ... },
  mutations: { ... },
  actions: { ... },
  getters: { ... }
}

const moduleB = {
  state: { ... },
  mutations: { ... },
  actions: { ... }
}

const store = new Vuex.Store({
  modules: {
    a: moduleA,
    b: moduleB
  }
})

store.state.a // -> moduleA 的状态
store.state.b // -> moduleB 的状态

```

### 2.2 模块的局部状态

对于模块内部的 mutation 和 getter，接收的第一个参数是模块的局部状态对象。

```javascript
const moduleA = {
  state: { count: 0 },
  mutations: {
    increment (state) {
      // 这里的 `state` 对象是当前模块的局部状态
      state.count++
    }
  },

  getters: {
    doubleCount (state) {
      // 这里的 `state` 对象是当前模块的局部状态
      return state.count * 2
    }
  }
}
```

同样，对于模块内部的 action，局部状态通过 context.state 暴露出来，根节点状态则为 context.rootState：

```javascript
const moduleA = {
  // ...
  actions: {
    incrementIfOddOnRootSum ({ state, commit, rootState }) {
      if ((state.count + rootState.count) % 2 === 1) {
        commit('increment')
      }
    }
  }
}
```

对于模块内部的 getter，根节点状态会作为第三个参数暴露出来：

```javascript
const moduleA = {
  // ...
  getters: {
    sumWithRootCount (state, getters, rootState) {
      return state.count + rootState.count
    }
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
    <div>模块 A 数量：{{moduleACount}}</div>
    <div>根节点 数量：{{rootCount}}</div>
    <div>数量总和：{{countSum}}</div>
    <button @click="addModuleCount">模块 A + 1</button>
    <button @click="addRootToModule">添加 root 至模块</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const moduleA = {
    state: {
      count: 18
    },
    getters: {
      countSum(state, getters, rootState) {
        return state.count + rootState.count
      }
    },
    mutations: {
      addModuleCount(state) {
        state.count++
      },
      addModuleByCount(state, payload) {
        state.count = state.count + payload.count
      }
    },
    actions: {
      addRootToModule({state, commit, rootState}) {
        commit('addModuleByCount', {count: rootState.count})
      }
    }
  }
  const store = new Vuex.Store({
    modules: {
      a: moduleA,
    },
    state: {
      count: 20
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    computed: {
      countSum() {
        return this.$store.getters.countSum
      },
      moduleACount() {
        return this.$store.state.a.count
      },
      rootCount() {
        return this.$store.state.count
      }
    },
    methods: {
      addModuleCount() {
        this.$store.commit('addModuleCount')
      },
      addRootToModule() {
        this.$store.dispatch('addRootToModule')
      }
    }
  })
</script>
</html>

```

**代码解释**

JS 代码第 4-26 行，我们定义了模块 moduleA。

JS 代码第 9-11 行，在 moduleA 定义 getter countSum。

JS 代码第 13-20 行，在 moduleA 定义 mutations。

JS 代码第 21-25 行，在 moduleA 定义 actions。

JS 代码第 27-34 行，我们定义了 store，并将 moduleA 传入 modules 的属性中。

## 3. 小结

本小节我们介绍了如何使用 `Modules` 进行模块化。主要有以下知识点：

* 如何定义一个模块 module。
* 在 store 中利用 modules 属性传入定义的模块 module。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
