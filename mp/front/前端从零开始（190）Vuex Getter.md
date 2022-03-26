# Getter

## 1. 前言

本小节我们将介绍 Vuex 中 getter 的使用方式。包括如何定义 getter、使用 getter、辅助函数 mapGetters 的使用。Getter 在项目中的使用非常普通，学会使用 Getter 可以避免我们重复的通过 state 获取数据。同学们在学完本小节后可以多尝试写一些 Getter 来巩固本节的知识点。

## 2. 慕课解释

> Vuex 允许我们在 store 中定义 “getter”（可以认为是 store 的计算属性）。就像计算属性一样，getter 的返回值会根据它的依赖被缓存起来，而且只有当它的依赖值发生了改变才会被重新计算。— 官方定义

我们可以把 Getter 理解成是封装好的获取数据的方法，在方法内部我们可以对 state 中的数据做一些相应的处理，最后返回我们想要的数据。

## 3. 用法

### 3.1 通过属性访问

Getter 接受 state 作为其第一个参数，我们可以对 state 中的数据做相应的处理，最终返回我们想要的数据：

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
    <div v-for="item in skillList" :key="item.name">{{item.name}}</div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      name: '句号',
      age: 18,
      skill: [
        {name: 'Vue', type: 1},
        {name: 'React', type: 1},
        {name: 'JAVA', type: 2},
        {name: 'Webpack', type: 3},
        {name: 'Node', type: 1}
      ]
    },
    getters: {
      skillList: state => {
        return state.skill.filter(item => item.type === 1)
      }
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    computed: {
      skillList() {
        return this.$store.getters.skillList
      }
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 16-20 行，我们定义了 Getter 方法 skillList，skillList 内部我们返回状态 skill 中 type 为 1 的数据。

JS 代码第 26-28 行，我们通过 $store.getters 获取 skillList 的返回值。

Getter 也可以接受其他 getter 作为第二个参数：

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
    <div v-for="item in skillList" :key="item.name">{{item.name}}</div>
    <div>我有 {{count}} 个技能包</div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      name: '句号',
      age: 18,
      skill: [
        {name: 'Vue', type: 1},
        {name: 'React', type: 1},
        {name: 'JAVA', type: 2},
        {name: 'Webpack', type: 3},
        {name: 'Node', type: 1}
      ]
    },
    getters: {
      skillList: state => {
        return state.skill.filter(item => item.type === 1)
      },
      skillCount: (state, getters) => {
        return getters.skillList.length
      },
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    computed: {
      skillList() {
        return this.$store.getters.skillList
      },
      count() {
        return this.$store.getters.skillCount
      }
    }
  })
</script>
</html>

```

**代码解释**

JS 代码第 16-23 行，我们定义了 Getter 方法 skillList 和 skillCount，skillList 内部我们返回 skill 数据 中 type 为 1 的数组，skillCount 内部我们通过 getters 获取 skillList 的数组长度。

JS 代码第 28-30 行，我们通过 $store.getters 获取 skillList 的返回值。

JS 代码第 31-33 行，我们通过 $store.getters 获取 skillCount 的返回值。

### 3.2 通过方法访问

在上一个例子中我们只能通过 skillList 获取 type 为 1 的数据列表，那么如果我想获取 type 为 2 的数据呢？同学们可能会说：我们在定义一个 skillList2 不就好了！确实这样可以满足需要，但是，如果又有 type = 3、type = 4 等等其他的呢？难道我们还要继续写 skillList3、skillList4 吗？

其实 getter 除了可以直接返回数据之外，也可以通过让 getter 返回一个函数，来实现给 getter 传参。在对 store 里的数组进行查询时非常有用。

示例：

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
    <div v-for="item in skillList" :key="item.name">{{item.name}}</div>
    <div>我有 {{count}} 个技能包</div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      name: '句号',
      age: 18,
      skill: [
        {name: 'Vue', type: 1},
        {name: 'React', type: 1},
        {name: 'JAVA', type: 2},
        {name: 'Webpack', type: 3},
        {name: 'Node', type: 1}
      ]
    },
    getters: {
      skillList: state => (type) => {
        return state.skill.filter(item => item.type === type)
      },
      skillCount: (state, getters) => (type) => {
        return getters.skillList(type).length
      },
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    computed: {
      skillList() {
        return this.$store.getters.skillList(2)
      },
      count() {
        return this.$store.getters.skillCount(2)
      }
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 17-19 行，我们定义了 Getter 方法 skillList，skillList 返回一个函数，该函数接收一个 type 参数，函数内部返回 state.skill 中对应 type 的数组。

JS 代码第 20-22 行，我们定义了 Getter 方法 skillCount，skillCount 返回一个函数，该函数接收一个 type 参数，函数内部获取 getters.skillList 的值，并返回数组长度。

JS 代码第 28-30 行，我们通过 $store.getters.skillList 传入参数 type 获取 skillList 的返回值。

JS 代码第 31-33 行，我们通过 $store.getters.skillCount 传入参数 type 获取 skillCount 的返回值。

## 4. mapGetters 辅助函数

mapGetters 辅助函数仅仅是将 store 中的 getter 映射到局部计算属性：

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
    <div v-for="item in skillList" :key="item.name">{{item.name}}</div>
    <div>我有 {{skillCount}} 个技能包</div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
<script type="text/javascript">
  const store = new Vuex.Store({
    state: {
      name: '句号',
      age: 18,
      skill: [
        {name: 'Vue', type: 1},
        {name: 'React', type: 1},
        {name: 'JAVA', type: 2},
        {name: 'Webpack', type: 3},
        {name: 'Node', type: 1}
      ]
    },
    getters: {
      skillList: state => {
        return state.skill.filter(item => item.type === 1)
      },
      skillCount: (state, getters)  => {
        return getters.skillList.length
      },
    }
  })
  var vm = new Vue({
    el: '#app',
    store,
    computed: {
      ...Vuex.mapGetters([
        'skillList',
        'skillCount'
      ])
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 29-32 行我们通过 Vuex.mapGetters 获取 skillList 和 skillCount 的值。

如果你想将一个 getter 属性另取一个名字，可以使用对象形式：

```javascript
Vuex.mapGetters({
  skillListAlias: 'skillList',
  skillCountAlias: 'skillCount'
})
```

## 5. 小结

本节，我们带大家学习了 Vuex 中 Getter 的使用方式。主要知识点有以下几点：

* 在 store 中定义 Getter 数据。
* 通过 $store.getter 访问 getter。
* 通过让 getter 返回一个函数给 getter 传参。
* 使用 mapGetters 辅助函数简化获取 getter 的写法。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
