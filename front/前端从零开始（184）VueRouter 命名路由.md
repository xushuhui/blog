# VueRouter 命名路由

## 1. 前言

本小节我们介绍如何使用 VueRouter 命名路由。包括如何定义命名路由、如何使用路由名实现路由跳转。本节的学习内容相对简单，相信同学们看完本小节，并对小节中的案例自己实现一遍就可以熟练掌握了。

## 2. 定义路由名

在之前的小节中，我们学习了如何定义一个路由：

```javascript
const router = new VueRouter({
  routes: [
    {
      path: '/user',
      component: '[component-name]'
    }
  ]
})
```

route 对象中有两个属性。path 表示路由地址，component 表示路由显示的组件。我们可以在 route 对象中添加一个 name 属性用来给路由指定一个名字：

```javascript
const router = new VueRouter({
  routes: [
    {
      path: '/user',
      name: 'user',
      component: '[component-name]'
    }
  ]
})
```

### 2.1 `<router-link>` 跳转命名路由

在之前的小节中，我们学习了使用 `<router-link to="path">...</router-link>` 的方式来实现路由跳转。实际上 `router-link` 的 to 属性可以接收一个对象：

```javascript
  <router-link :to="{path: 'path'}">...</router-link>
```

让我们来看一个简单的示例：

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
    <div>
      <router-link :to="{path: '/index'}">首页</router-link>
      <router-link to="/article">文章</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Index = Vue.component('index', {
  template: '<div>Hello，欢迎使用慕课网学习 Vue 教程！</div>',
})

const Article = Vue.component('myArticle', {
  template: `<ul><li>1. Vue 计算属性的学习</li><li>2. React 基础学习</li></ul>`,
})

const routes = [
  { path: '/index', component: Index },
  { path: '/article', component: Article }
]

const router = new VueRouter({
  routes: routes
})

var vm = new Vue({
  el: '#app',
  router: router,
  data() {
    return {}
  }
})
</script>
</html>

```

**代码解释：**

HTML 代码第 12 行，我们定义了首页跳转链接，通过对象的形式给属性 to 赋值。

HTML 代码第 13 行，我们定义了文章跳转链接，通过字符串的形式给属性 to 赋值。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-16 行，我们定义了路由数组：

- 1. 首页路由，地址为 ‘/index’，匹配组件 Index。

- 2. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

除了通过 path 可以链接到路由外，还可以通过路由 name 实现链接跳转：

```javascript
`<router-link :to="{name: 'name'}">...</router-link>`
```

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
    <div>
      <router-link :to="{name: 'index'}">首页</router-link>
      <router-link :to="{name: 'article'}">文章</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Index = Vue.component('index', {
  template: '<div>Hello，欢迎使用慕课网学习 Vue 教程！</div>',
})

const Article = Vue.component('myArticle', {
  template: `<ul><li>1. Vue 计算属性的学习</li><li>2. React 基础学习</li></ul>`,
})

const routes = [
  { path: '/index', name: 'index', component: Index },
  { path: '/article', name: 'article', component: Article }
]

const router = new VueRouter({
  routes: routes
})

var vm = new Vue({
  el: '#app',
  router: router,
  data() {
    return {}
  }
})
</script>
</html>

```

**代码解释：**

HTML 代码第 12-13 行，我们定义了两个跳转链接，通过对象的形式给属性 to 赋值，跳转指定 name 的路由。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-16 行，我们定义了路由数组：

*     1. 首页路由，地址为 ‘/index’， 路由名为 index，匹配组件 Index。

*     1. 文章路由，地址为 ‘/article’， 路由名为 article，匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

### 2.2 编程式导航跳转命名路由

在之前的小节中，我们学习了如何使用 $router 实例来实现编程式的导航。我们也可以使用 $router 实例跳转指定名字的路由地址：

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
    <div>
      <button @click="jump('index')">首页</button>
      <button @click="jump('article')">文章</button>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Index = Vue.component('index', {
  template: '<div>Hello，欢迎使用慕课网学习 Vue 教程！</div>',
})

const Article = Vue.component('myArticle', {
  template: `<ul><li>1. Vue 计算属性的学习</li><li>2. Vue 侦听器的学习</li></ul>`,
})


const routes = [
  { path: '/index', name: 'index', component: Index },
  { path: '/article', name: 'article' , component: Article }
]

const router = new VueRouter({
  routes: routes
})

  var vm = new Vue({
    el: '#app',
    router,
    data() {
    	return {}
    },
    methods: {
      jump(name) {
        this.$router.push({
          name: name
        })
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 12-13 行，我们定义了两个按钮，并给他们点击事件 jump。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-16 行，我们定义了路由数组：

1. 首页路由，地址为 ‘/index’，匹配组件 Index。

2. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

JS 代码第 29-31 行，我们定义来 jump 函数，通过 router.push 实现路由跳转。

## 3. 小结

本节，我们带大家学习了 VueRouter 命名路由的使用方法。主要知识点有以下几点：

* 通过 name 属性指定路由名称。
* 通过 `<router-link>` 跳转指定名称的路由地址。
* 通过 $router 跳转指定名称的路由地址。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
