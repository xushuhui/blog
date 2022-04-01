# VueRouter 基础使用

## 1. 前言

本小节我们介绍如何在 Vue 项目中使用 VueRouter。包括 VueRouter 的下载、什么是 VueRouter、如何使用 VueRouter 配置一个单页应用。其中，学习使用 VueRouter 配置一个单页应用是本节的重点。同学们在学完本节课程之后需要自己多尝试配置路由。

## 2. 慕课解释

> Vue Router 是 Vue.js 官方的路由管理器。它和 Vue.js 的核心深度集成，让构建单页面应用变得易如反掌。 — 官方定义
>
>
> VueRouter 是 SPA（单页应用）的路径管理器，它允许我们通过不同的 URL 访问不同的内容。

## 3 安装 VueRouter

### 3.1 直接下载

我们可以在官网 ([VueRouter](https://unpkg.com/vue-router/dist/vue-router.js)) 上直接下载 `VueRouter`。

在 `Vue` 之后引入 `VueRouter` 会进行自动安装：

```javascript
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="/(具体文件地址)/vue-router.js"></script>
```

### 3.2 CDN 引用

```javascript
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
```

### 3.3 NPM 或 Yarn

```javascript
npm install vue-router
or
yarn add vue-router
```

在一个模块化的打包系统中，您必须显式地通过 Vue.use () 来安装 Vuex：

```javascript
import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
```

在本章节的 VueRouter 学习中，我们都将使用 CDN 的方式引入路由。

## 4. 使用 VueRouter

用 Vue.js + VueRouter 创建单页应用，是非常简单的。使用 Vue.js ，我们已经可以通过组合组件来组成应用程序，当你要把 VueRouter 添加进来，我们需要做的是，将组件 (components) 映射到路由 (routes)，然后告诉 VueRouter 在哪里渲染它们。

在使用 VueRouter 之前，我们需要先了解 VueRouter 的两个内置组件：

1. `<router-link>`：该组件用于设置一个导航链接，切换不同 HTML 内容。 to 属性为目标地址，即要显示的内容。例：`<router-link to="/index">首页</router-link>`；
2. `<router-view>`：该组件将渲染路由匹配到的组件内容。

接下来我们看一个基本例子：

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
  template: '<div>Hello，欢迎使用学习 Vue 教程！</div>',
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

HTML 代码第 12-13 行，我们定义了两个跳转链接；

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件；

JS 代码第 5-7 行，我们定义了组件 Index；

JS 代码第 9-11 行，我们定义了组件 Article；

JS 代码第 13-16 行，我们定义了路由数组：

1. 首页路由，地址为 ‘/index’，匹配组件 Index；
2. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

## 5. 小结

本节，我们带大家学习了 VueRouter。主要知识点有以下几点：

* 通过 CDN、NPM、Yarn 等方式下载 VueRouter；
* 使用 VueRouter 配置一个简单的单页应用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
