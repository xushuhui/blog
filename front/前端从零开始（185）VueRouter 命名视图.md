# VueRouter 命名视图

## 1. 前言

本小节我们介绍如何使用 VueRouter 命名视图。包括如何定义命名视图、如何使用命名视图。本节的学习内容相对简单，相信同学们看完本小节，并对小节中的案例自己实现一遍就可以熟练掌握了。

## 2. 定义视图名

### 2.1 默认视图

在之前的小节中，我们学习了如何使用 `<router-view/>` 来承载路由分发的内容。我们并没有给 `<router-view/>` 指定一个 name 属性，实际上他有一个默认的属性 default，我们以一个简单的实例来验证这一点：

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
      <router-link to="/index">首页</router-link>
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
  { path: '/index', components: {default: Index} },
  { path: '/article', components: {default: Article} }
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

HTML 代码第 12-13 行，我们定义了两个跳转链接。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-16 行，我们定义了路由数组：

- 1. 首页路由，地址为 ‘/index’，默认视图匹配组件 Index。

- 2. 文章路由，地址为 ‘/article’，默认视图匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

### 2.2 具名视图

除了使用默认视图名外，我们还可以给视图指定一个名字：

```javascript
  <router-view name="name"/>
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
      <router-link to="/index">首页</router-link>
      <router-link to="/article">文章</router-link>
    </div>
    <router-view name="view"></router-view>
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
  { path: '/index', components: {view: Index} },
  { path: '/article', components: {view: Article} }
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

**代码解释**

我们对上述案例做一个简单的修改：

1. 指定 `<router-view />` 的视图名为 view。
2. 定义路由信息的时候，指定视图 view 匹配对应组件。

## 3. 小结

本节，我们带大家学习了 VueRouter 命名视图的使用方法。主要知识点有以下几点：

* 通过 name 属性指定视图名称。
* 通过路由 `components` 指定各具名视图对应匹配的组件。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
