# VueRouter 路由别名、重定向

## 1. 前言

本小节我们介绍如何使用 VueRouter 路由别名、重定向。路由别名和重定向在项目中经常使用，本节的学习内容相对简单，相信同学们看完本小节，并对小节中的案例自己实现一遍就可以熟练掌握了。

## 2. 路由重定向

重定向也是通过 routes 配置来完成，可以配置路由重定向到具体路由地址、具名路由或者动态返回重定向目标。

### 2.1 重定向到路由地址

通过属性 redirect 指定重定向的路由地址：

```javascript
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: '/b' }
  ]
})
```

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
  { path: '/', redirect: '/index' },
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

HTML 代码第 12-13 行，我们定义了两个跳转链接。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-17 行，我们定义了路由数组：

1. 根路由，地址为 ‘/’，重定向到路由地址 ‘/index’。
2. 首页路由，地址为 ‘/index’，匹配组件 Index。
3. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 19-21 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 25 行，通过 router 配置参数注入路由。

### 2.2 重定向到具名路由

通过属性 redirect 重定向到具名路由：

```javascript
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: {name: 'name'} }
  ]
})
```

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
  { path: '/', redirect: {name: 'index'} },
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

HTML 代码第 12-13 行，我们定义了两个跳转链接。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-17 行，我们定义了路由数组：

1. 根路由，地址为 ‘/’，重定向到具名路由 ‘index’。
2. 首页路由，地址为 ‘/index’，匹配组件 Index。
3. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 19-21 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 25 行，通过 router 配置参数注入路由。

### 2.3 动态返回重定向目标

属性 redirect 可以接收一个方法，动态返回重定向目标：

```javascript
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: to => {
      // 方法接收 目标路由 作为参数
      // return 重定向的 字符串路径/路径对象
    }}
  ]
})
```

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
{ path: '/', redirect: to => {
  if(Math.random() > 0.5) {
    return '/index'
  }else {
    return {
      name: 'article'
    }
  }
}},
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

HTML 代码第 12-13 行，我们定义了两个跳转链接。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-25 行，我们定义了路由数组：

1. 根路由，地址为 ‘/’，根据随机数的大小重定向到路由 ‘/index’ 或 ‘/article’。
2. 首页路由，地址为 ‘/index’，匹配组件 Index。
3. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 27-29 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 32 行，通过 router 配置参数注入路由。

## 3. 路由别名

“重定向”的意思是，当用户访问 /a 时，URL 将会被替换成 /b，然后匹配路由为 /b，那么“别名”又是什么呢？

/a 的别名是 /b，意味着，当用户访问 /b 时，URL 会保持为 /b，但是路由匹配则为 /a，就像用户访问 /a 一样。

```javascript
const router = new VueRouter({
  routes: [
    { path: '/a', component: A, alias: '/b' }
  ]
})
```

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
  { path: '/index', component: Index, alias: '/' },
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

HTML 代码第 12-13 行，我们定义了两个跳转链接。

HTML 代码第 15 行，我们使用 `<router-view></router-view>` 组件来渲染匹配组件。

JS 代码第 5-7 行，我们定义了组件 Index。

JS 代码第 9-11 行，我们定义了组件 Article。

JS 代码第 13-16 行，我们定义了路由数组：

1. 首页路由，地址为 ‘/index’，匹配组件 Index，路由别名 ‘/’。
2. 文章路由，地址为 ‘/article’，匹配组件 Article。

JS 代码第 18-20 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 24 行，通过 router 配置参数注入路由。

## 4. 小结

本节，我们带大家学习了 VueRouter 路由重定向和别名。主要知识点有以下几点：

* 通过 redirect 属性指定路由重定向地址。
* 通过 alias 属性配置路由别名。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
