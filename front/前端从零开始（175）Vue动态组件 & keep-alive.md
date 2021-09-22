## 1. 前言

本小节我们将介绍 Vue 的动态组件，以及缓存 keep-alive 的使用。包括动态组件的使用方法，以及如何使用 keep-alive 实现组件的缓存效果。

## 2. 慕课解释

动态组件是让多个组件使用同一个挂载点，并动态切换。动态组件是 Vue 的一个高级用法，但其实它的使用非常简单。keep-alive 是 vue 的内置组件，能在组件切换过程中将状态保存在内存中，防止重复渲染 DOM。

## 3. 动态组件如何使用

通过使用保留的 `<component>` 元素，动态地把组件名称绑定到它的 `is` 特性，可以实现动态组件：

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
    <component :is="currentView"></component>
    <button @click="changeView('A')">切换到A</button>
    <button @click="changeView('B')">切换到B</button>
    <button @click="changeView('C')">切换到C</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.component('ComponentA', {
    template: '<div> 组件 A </div>',
  })
  Vue.component('ComponentB', {
    template: '<div> 组件 B </div>',
  })
  Vue.component('ComponentC', {
    template: '<div> 组件 C </div>',
  })
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        currentView: 'ComponentB'
      }
    },
    methods: {
      changeView(name) {
        this.currentView = `Component${name}`
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 2 行，我们使用动态组件 component，将当前需要展示的组件名通过变量 currentView 绑定到 component 的 is 属性上。

HTML 代码第 3-5 行，我们定义了三个按钮，通过点击按钮切换 currentView 的值。

JS 代码第 3-11 行，我们定义了组件 ComponentA、ComponentB、ComponentC。

最终的实现效果是：当点击按钮的时候会动态切换展示的组件。

## 4. keep-alive

`keep-alive` 是 `Vue` 提供的一个抽象组件，用来对组件进行缓存，从而节省性能，由于是一个抽象组件，所以在页面渲染完毕后不会被渲染成一个 DOM 元素。被 `keep-alive` 缓存的组件只有在初次渲染时才会被创建，并且当组件切换时不会被销毁。

### 4.1. 基础用法

`keep-alive` 的用法相对简单，直接使用 `keep-alive` 包裹需要缓存的组件即可：

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
    <keep-alive>
      <component :is="currentView"></component>
    </keep-alive>
    <button @click="changeView('A')">切换到A</button>
    <button @click="changeView('B')">切换到B</button>
    <button @click="changeView('C')">切换到C</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.component('ComponentA', {
    template: '<div> 组件 A </div>',
    created() {
      console.log('组件A created')
    }
  })
  Vue.component('ComponentB', {
    template: '<div> 组件 B </div>',
    created() {
      console.log('组件B created')
    }
  })
  Vue.component('ComponentC', {
    template: '<div> 组件 C </div>',
    created() {
      console.log('组件C created')
    }
  })
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        currentView: 'ComponentB'
      }
    },
    methods: {
      changeView(name) {
        this.currentView = `Component${name}`
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 2-3 行，我们使用 keep-alive 包裹动态组件 component，将当前需要展示的组件名通过变量 currentView 绑定到 component 的 is 属性上。

HTML 代码第 5-7 行，我们定义了三个按钮，通过点击按钮切换 currentView 的值。

JS 代码第 3-29 行，我们定义了组件 ComponentA、ComponentB、ComponentC，分别定义了他们的 created 和 beforeDestroy 事件。

之前我们介绍过，`keep-alive` 缓存的组件只有在初次渲染时才会被创建。所以，我们通过修改 currentView 切换组件时，组件的 beforeDestroy 事件不会触发。若该组件是第一次渲染，会触发 created 事件，当再次切换显示该组件时，created 事件不会再次触发。

### 4.2. activated 和 deactivated 生命周期

activated 和 deactivated 和我们之前学习的生命周期函数一样，也是组件的生命周期函数。不过， `activated` 和 `deactivated` 只在 `<keep-alive>` 内的所有嵌套组件中触发。`activated`：进入组件时触发。`deactivated`：退出组件时触发。

示例代码：

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
    <keep-alive>
      <component :is="currentView"></component>
    </keep-alive>
    <button @click="changeView('A')">切换到A</button>
    <button @click="changeView('B')">切换到B</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.component('ComponentA', {
    template: '<div> 组件 A </div>',
    activated() {
      console.log('组件A 被添加')
    },
    deactivated() {
      console.log('组件A 被移除')
    }
  })
  Vue.component('ComponentB', {
    template: '<div> 组件 B </div>',
    activated() {
      console.log('组件B 被添加')
    },
    deactivated() {
      console.log('组件B 被移除')
    }
  })
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        currentView: 'ComponentB'
      }
    },
    methods: {
      changeView(name) {
        this.currentView = `Component${name}`
      }
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码中，我们定义了组件 ComponentA、ComponentB，并分别定义了他们的 activated 和 deactivated 事件函数。

HTML 代码第 2-3 行，我们使用 keep-alive 包裹动态组件 component，将当前需要展示的组件名通过变量 currentView 绑定到 component 的 is 属性上。

HTML 代码第 5-6 行，我们定义了两个按钮，通过点击按钮切换 currentView 的值。当我们切换组件显示时，可以看到这样的打印信息：

```javascript
1. ComponentA -> ComponentB 会打印出：组件A 被移除、组件B 被添加
2. ComponentB -> ComponentA 会打印出：组件B 被移除、组件A 被添加
```

> **TIPS：** 注意，activated 和 deactivated 这两个生命周期函数一定是要在使用了 keep-alive 组件后才会有的，否则不存在。

### 4.3. include 和 exclude

`include` 和 `exclude` 是 keep-alive 的两个属性，允许组件有条件地缓存。

include： 可以是字符串或正则表达式，用来表示只有名称匹配的组件会被缓存。

exclude： 可以是字符串或正则表达式，用来表示名称匹配的组件不会被缓存。

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
    <keep-alive include="ComponentA,ComponentB">
      <component :is="currentView"></component>
    </keep-alive>
    <button @click="changeView('A')">切换到A</button>
    <button @click="changeView('B')">切换到B</button>
    <button @click="changeView('C')">切换到C</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.component('ComponentA', {
    template: '<div> 组件 A </div>',
    created() {
      console.log('组件A created')
    },
    activated() {
      console.log('组件A 被添加')
    },
    deactivated() {
      console.log('组件A 被移除')
    }
  })
  Vue.component('ComponentB', {
    template: '<div> 组件 B </div>',
    created() {
      console.log('组件B created')
    },
    activated() {
      console.log('组件B 被添加')
    },
    deactivated() {
      console.log('组件B 被移除')
    }
  })
  Vue.component('ComponentC', {
    template: '<div> 组件 C </div>',
    created() {
      console.log('组件C created')
    },
    activated() {
      console.log('组件C 被添加')
    },
    deactivated() {
      console.log('组件C 被移除')
    }
  })
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        currentView: 'ComponentB'
      }
    },
    methods: {
      changeView(name) {
        this.currentView = `Component${name}`
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 2-4 行，我们使用 keep-alive 包裹动态组件 component。给 keep-alive 指定需要缓存组件 ComponentA，ComponentB。

在之前的小节我们了解到 `keep-alive` 缓存的组件只有在初次渲染时才会被创建。所以，在案例中，组件 ComponentA 和 ComponentB 的 created 函数只有在第一次组件被创建的时候才会触发，而 ComponentC 的 created 函数当每次组件显示的时候都会触发。

exclude 示例：

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
    <keep-alive exclude="ComponentA,ComponentB">
      <component :is="currentView"></component>
    </keep-alive>
    <button @click="changeView('A')">切换到A</button>
    <button @click="changeView('B')">切换到B</button>
    <button @click="changeView('C')">切换到C</button>
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.component('ComponentA', {
    template: '<div> 组件 A </div>',
    created() {
      console.log('组件A created')
    }
  })
  Vue.component('ComponentB', {
    template: '<div> 组件 B </div>',
    created() {
      console.log('组件B created')
    }
  })
  Vue.component('ComponentC', {
    template: '<div> 组件 C </div>',
    created() {
      console.log('组件C created')
    }
  })
  var vm = new Vue({
    el: '#app',
    data() {
    	return {
        currentView: 'ComponentB'
      }
    },
    methods: {
      changeView(name) {
        this.currentView = `Component${name}`
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 2-4 行，我们使用 keep-alive 包裹动态组件 component。给 keep-alive 指定不需要缓存组件 ComponentA，ComponentB。

## 5. 小结

本节，我们带大家学习了动态组件和缓存组件在项目中的运用。主要知识点有以下几点：

* 使用 `<component :is="component-name"/>` 的方式实现动态组件；
* 使用 keep-alive 实现组件的缓存；
* 使用 include 属性指定需要缓存的组件；
* 使用 exclude 属性指定不需要缓存的组件。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
