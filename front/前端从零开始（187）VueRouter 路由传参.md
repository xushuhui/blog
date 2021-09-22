# VueRouter 路由传参

## 1. 前言

本小节我们介绍 VueRouter 路由组件传参。包括 params 传参、query 传参的两种方式。路由传参的知识点非常重要，在日常开发中，我们经常会通过路由传递各种参数，同学们在学完本节后可以将小节中的案例自己动手实现一遍，这样才可以加深印象并熟练掌握。

## 2. params 传参

使用 params 传参数我们可以分为两个步骤：

1. 定义路由以及路由接收的参数。
2. 路由跳转时传入对应参数。

首先，我们先了解如何定义路由接收的参数：

```javascript
const routes = [
  { path: '/detail/:name', name: 'detail', component: Detail },
]
```

使用 `<router-link></router-link>` 的方式跳转路由：

```javascript
 <!-- router-link 跳转 -->
<router-link :to="{name: 'detail', params: {name: 'React 基础学习'}}">2. React 基础学习</router-link>
```

具体示例：

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
      <router-link to="/">首页</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Detail = Vue.component('detail', {
  template: '<div>这是 {{$route.params.name}} 的详情页面</div>'
})

const Article = Vue.component('myArticle', {
  template: `<ul>
              <li>
                <router-link :to="{name: 'detail', params: {name: 'Vue 计算属性的学习'}}">
                  1. Vue 计算属性的学习
                </router-link>
              </li>
              <li>
                <router-link :to="{name: 'detail', params: {name: 'React 基础学习'}}">
                  2. React 基础学习
                </router-link>
              </li>
            </ul>`
})

const routes = [
  { path: '/detail/:name', name: 'detail', component: Detail },
  { path: '/', component: Article }
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

在 JS 代码第 24 行，我们定义了路由 detail，他通过 params 接收一个参数 name。

在组件 Article 中，我们使用 `<router-link>` 链接要跳转的路由并将参数传入。

在组件 Detail 中，我们将传入的课程名称显示出来。

使用 $router 的方式跳转路由：

```javascript
 // $router 跳转
 this.$router.push({
  name: 'detail',
  params: {
    name: 'Vue 教程'
  }
})
```

具体示例：

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
      <router-link to="/">首页</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Detail = Vue.component('detail', {
  template: '<div>这是 {{$route.params.name}} 的详情页面</div>'
})

const Article = Vue.component('myArticle', {
  template: `<ul>
              <li @click="getDetail('Vue 计算属性的学习')">
                1. Vue 计算属性的学习
              </li>
              <li @click="getDetail('React 基础学习')">
                2. React 基础学习
              </li>
            </ul>`,
  methods: {
    getDetail(name) {
      this.$router.push({
        name: 'detail',
        params: {
          name: name
        }
      })
    }
  }
})

const routes = [
  { path: '/detail/:name', name: 'detail', component: Detail },
  { path: '/', component: Article }
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

在 JS 代码第 31 行，我们定义了路由 detail，他通过 params 接收一个参数 name。

在 JS 代码第 19 行，我们定义了方法 getDetail，该方法通过 $router.push 跳转到详情页面，并传入 name 参数。

在组件 Article 中，当我们点击课程名称的时候调用 getDetail 方法。

在组件 Detail 中，我们将传入的课程名称显示出来。

## 3. query 传参

使用 query 传参的方法相对简单，只需要在对应路由跳转的时候传入参数即可：

使用 `<router-link></router-link>` 的方式跳转路由：

```javascript
 <!-- router-link 跳转 -->
<router-link :to="{path: '/detail', query: { id: 1 }}">2. React 基础学习</router-link>
```

具体示例：

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
      <router-link to="/">首页</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Detail = Vue.component('detail', {
  template: '<div>这是 id 为 {{$route.query.id}} 的详情页面</div>'
})

const Article = Vue.component('myArticle', {
  template: `<ul>
              <li>
                <router-link :to="{path: '/detail', query: {id: 1}}">
                  1. Vue 计算属性的学习
                </router-link>
              </li>
              <li>
                <router-link :to="{path: '/detail', query: {id: 2}}">
                  2. React 基础学习
                </router-link>
              </li>
            </ul>`
})

const routes = [
  { path: '/detail', component: Detail },
  { path: '/', component: Article }
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

在组件 Article 中，我们使用 `<router-link>` 链接到要跳转的路由并将参数传入。

在组件 Detail 中，我们通过 $[route.query.id](http://route.query.id) 将传入的课程 ID 显示出来。

使用 $router 的方式跳转路由：

```javascript
 // $router 跳转
 this.$router.push({
  path: '/detail',
  query: {
    id: 2
  }
})
```

具体示例：

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
      <router-link to="/">首页</router-link>
    </div>
    <router-view></router-view>
  </div>
</body>

<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
<script type="text/javascript">

const Detail = Vue.component('detail', {
  template: '<div>这是 id 为 {{$route.query.id}} 的详情页面</div>'
})

const Article = Vue.component('myArticle', {
  template: `<ul>
              <li @click="getDetail(1)">1. Vue 计算属性的学习</li>
              <li @click="getDetail(2)">2. React 基础学习</li>
            </ul>`,
  methods: {
    getDetail(id) {
      this.$router.push({
        path: '/detail',
        query: {
          id: id
        }
      })
    }
  }
})

const routes = [
  { path: '/detail', component: Detail },
  { path: '/', component: Article }
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

在 JS 代码第 19 行，我们定义了方法 getDetail，该方法通过 $router.push 跳转到详情页面，并通过 query 传入参数 id。

在组件 Article 中，当我们点击课程名称的时候调用 getDetail 方法。

在组件 Detail 中，我们通过 $[route.query.id](http://route.query.id) 把传入的课程 ID 显示出来。

## 4. 小结

本节，我们带大家学习了路由传参的两种方式。主要知识点有以下几点：

* 通过 params 传递参数。
* 通过 query 传递参数。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
