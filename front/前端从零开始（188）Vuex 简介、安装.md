# Vuex 简介、安装

## 1. 简介

本小节我们会介绍数据管理工具 `vuex`。包括什么是 Vuex、Vuex 的安装、以及如何创建和使用 Vuex 数据仓库。

## 2. 什么是 Vuex

`Vuex` 是一个专为 `Vue.js` 应用程序开发的状态管理模式。它采用集中式存储管理应用的所有组件的状态，并以相应的规则保证状态以一种可预测的方式发生变化。

### 2.1 什么是状态管理模式

让我们从一个简单的 Vue 计数应用开始：

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
    <div>
      <button @click="increase">添加一次</button>
      <button @click="decrease">减少一次</button>
    </div>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>

<script type="text/javascript">
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        count: 0
      }
    },
    methods: {
      increase() {
        this.count++
      },
      decrease() {
        this.count--
      }
    }
  })
</script>
</html>

```

这个状态自管理应用包含以下几个部分：

* **state**，驱动应用的数据源；
* **view**，以声明方式将 **state** 映射到视图；
* **actions**，响应在 **view** 上的用户输入导致的状态变化。

以下是一个表示 “单向数据流” 理念的简单示意：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed849de09f23d6a12800866.jpg)

但是，当我们的应用遇到多个组件共享状态时，单向数据流的简洁性很容易被破坏：

* 多个视图依赖于同一状态。
* 来自不同视图的行为需要变更同一状态。

对于问题一，传参的方法对于多层嵌套的组件将会非常繁琐，并且对于兄弟组件间的状态传递无能为力。对于问题二，我们经常会采用父子组件直接引用或者通过事件来变更和同步状态的多份拷贝。以上的这些模式非常脆弱，通常会导致无法维护的代码。

因此，我们为什么不把组件的共享状态抽取出来，以一个全局单例模式管理呢？在这种模式下，我们的组件树构成了一个巨大的 “视图”，不管在树的哪个位置，任何组件都能获取状态或者触发行为！

通过定义和隔离状态管理中的各种概念并通过强制规则维持视图和状态间的独立性，我们的代码将会变得更结构化且易维护。

### 2.2 什么时候使用 Vuex

Vuex 可以帮助我们管理共享状态，并附带了更多的概念和框架。这需要对短期和长期效益进行衡量。

如果您不打算开发大型单页应用，使用 Vuex 可能是繁琐冗余的。确实是如此 —— 如果您的应用够简单，您最好不要使用 Vuex。一个简单的 store 模式就足够您所需了。但是，如果您需要构建一个中大型单页应用，您很可能会考虑如何更好地在组件外部管理状态，Vuex 将会成为自然而然的选择。

## 3. 安装 Vuex

### 3.1 直接下载

我们可以在官网 ([vuex](https://unpkg.com/vuex@3.1.2/dist/vuex.js)) 上直接下载 `vuex`。

在 `Vue` 之后引入 `vuex` 会进行自动安装：

```javascript
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="/path/to/vuex.js"></script>
```

### 3.2 CDN 引用

```javascript
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vuex@3.1.2/dist/vuex.js"></script>
```

### 3.3 NPM 或 Yarn

```javascript
npm install vuex --save
yarn add vuex
```

在一个模块化的打包系统中，您必须显式地通过 `Vue.use()` 来安装 Vuex：

```javascript
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
```

当使用全局 script 标签引用 Vuex 时，不需要以上安装过程。

## 4. 基本示例

每一个 Vuex 应用的核心就是 store（仓库）。“store” 基本上就是一个容器，它包含着你的应用中大部分的状态 (state)。Vuex 和单纯的全局对象有以下两点不同：

1. Vuex 的状态存储是响应式的。当 Vue 组件从 store 中读取状态的时候，若 store 中的状态发生变化，那么相应的组件也会相应地得到高效更新。
2. 你不能直接改变 store 中的状态。改变 store 中的状态的唯一途径就是显式地提交 (commit) mutation。这样使得我们可以方便地跟踪每一个状态的变化，从而让我们能够实现一些工具帮助我们更好地了解我们的应用。

最简单的 `store`:

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
  <div id="app">{{count}}</div>
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
    created() {
      console.log(store.state.count)
    },
    computed: {
	    count() {
		    return store.state.count
	    }
    }
  })
</script>
</html>
```

**代码解释**

在 JS 代码第 4-8 行，通过 new Vuex.Store ({…}) 创建数据仓库。

在 JS 代码第 12 行，我们可以通过 store.state.count 访问仓库中定义的数据。

## 5. 小结

本节，我们带大家学习了 vuex 的基本概念。主要知识点有以下几点：

* Vuex 是一个数据管理工具，我们可以通过它简化组件间的数据共享问题。
* Vuex 的安装和使用方法。
* 使用 new Vuex.Store ({…}) 创建数据仓库。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
