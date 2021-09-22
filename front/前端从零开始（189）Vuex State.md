# State

## 1. 前言

本小节我们将学习和使用 Vuex 中 state 的概念。包括如何创建 state、组件中获取 state、以及辅助函数 mapState 的使用方法。

## 2. 创建数据仓库

在上一小节中，我们已经给大家写了一个简单的示例，大家一定还记得 `Vuex.Store({...})` 这个方法。在 Vuex 中，我们通过该方法创建一个数据仓库，并把数据 state 传入。例如：

```javascript
const store = new Vuex.Store({
  state: {
    count: 12000,
    name: '慕课网',
    logo: ''
  }
})
```

那么，创建完数据仓库后，我们怎样才能在 Vue 组件中使用它呢？我们知道，要使用 Vue 需要通过 new Vue () 创建一个 Vue 实例，并传入对象的参数。要在 Vue 中使用 store，只需要在创建 Vue 实例的时候将 store 传入即可：

```javascript
var vm = new Vue({
  el: '#app',
  store: store
})
```

## 3. 在 Vue 组件中获得 Vuex 状态

那么我们如何在 Vue 组件中展示状态呢？由于 Vuex 的状态存储是响应式的，从 store 实例中读取状态最简单的方法就是在计算属性中返回某个状态：

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
    <div> {{ count }} </div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      count: 12
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
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

JS 代码第 4-8 行，我们定义了仓库 store。

JS 代码第 11 行，创建 Vue 实例的时候传入 store。

JS 代码第 13-15 行，利用计算属性返回 count。

HTML 中利用插值显示 count 的数据。

## 4. mapState 辅助函数

当一个组件需要获取多个状态时候，将这些状态都声明为计算属性会有些重复和冗余。为了解决这个问题，我们可以使用 mapState 辅助函数帮助我们生成计算属性：

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
    <div> 我是： {{ name }}，我的今年：{{ age }}</div>
    <div>{{countPlusAge}}</div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      name: '句号',
      age: 18
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    data() {
      return {
        count: 1
      }
    },
    computed: Vuex.mapState({
      // 箭头函数可使代码更简练
      name: state => state.name,

       // 传字符串参数 'age' 等同于 `state => state.age`
      age: 'age',

      // 为了能够使用 `this` 获取局部状态，必须使用常规函数
      countPlusAge (state) {
        return state.age + this.count
      }
    })
  })
</script>
</html>

```

**代码解释**

JS 代码第 4-9 行，我们定义了仓库 store。

JS 代码第 12 行，创建 Vue 实例的时候传入 store。

JS 代码第 18-28 行，利用计算属性分别返回 name、age、countPlusAge。

当映射的计算属性的名称与 state 的子节点名称相同时，我们也可以给 mapState 传一个字符串数组。

```javascript
computed: Vuex.mapState([
  // 映射 this.age 为 store.state.age
  'age',
   // 映射 this.name 为 store.state.name
  'name'
])
// ===等同于===
computed:  Vuex.mapState({age:'age', name: 'name'})
//
```

## 5. 对象展开运算符

mapState 函数返回的是一个对象。我们如何将它与局部计算属性混合使用呢？通常，我们需要使用一个工具函数将多个对象合并为一个，以使我们可以将最终对象传给 computed 属性。但是自从有了对象展开运算符，我们可以极大地简化写法：

```javascript
computed: {
  localComputed () { /* ... */ },

  // 使用对象展开运算符将此对象混入到外部对象中
  ...Vuex.mapState({
    // ...
  })
}
```

## 6. 小结

本节，我们带大家学习了 Vuex 中 state 的使用方式。主要知识点有以下几点：

* 在 store 中定义 state 数据。
* 通过 $store.state 访问 state 中的数据。
* 使用 mapState 辅助函数简化获取 state 中数据的写法。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
