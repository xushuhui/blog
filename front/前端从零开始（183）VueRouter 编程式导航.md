# VueRouter 编程式导航

## 1. 前言

本小节我们介绍如何使用 VueRouter 编程式导航。包括 push、replace、go 三种方法的使用和区别。其中了解和掌握三种方法使用方式以及他们之间的区别是本节的重点。本节的内容相对容易，同学们只需要在学完本节的内容后稍加记忆，并通过一两个案例进行调试，相信一定可以对这三种方法的使用游刃有余。

## 2. router.push

在之前的小节中，我们的路由跳转是通过标签 `<router-link>` 来完成的。但有时候，我们可能需要通过一个普通的 onClick 事件来完成跳转。router.push 就可以帮我们实现这一点。

### 2.1 基本用法

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
  { path: '/index', component: Index },
  { path: '/article', component: Article }
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
        this.$router.push(name)
      }
    }
  })
</script>
</html>

```

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

### 2.2 对象格式的参数

在上一个例子中，我们通过 router.push 的方法实现了路由跳转，该方法接收跳转路径作为参数。实际上，router.push 也可以接收一个描述地址的对象作为参数。例如：

```javascript

// 字符串形式的参数
router.push('home')

// 通过路径描述地址
router.push({ path: 'home' })

// 通过命名的路由描述地址
router.push({ name: 'user' }})

```

当以对象形式传递参数的时候，还可以有一些其他属性，例如查询参数 params、query。路由传参我们将有一个专门的小节来学习，在这里同学们只需要有一个印象即可。

## 3. router.replace

跟 router.push 很像，唯一的不同就是，它不会向 history 添加新记录，而是跟它的方法名一样 —— 替换掉当前的 history 记录。我们将上一个例子中的 jump 函数稍加修改：

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
  { path: '/index', component: Index },
  { path: '/article', component: Article }
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
        this.$router.replace(name)
      }
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 29-31 行，我们定义来 jump 函数，通过 router.replace 实现路由跳转。

## 4. router.go

这个方法的参数是一个整数，意思是在 history 记录中向前或者后退多少步。例如：

```javascript

// 在浏览器记录中前进一步
router.go(1)

// 后退一步记录
router.go(-1)

// 前进 3 步记录
router.go(3)

// 如果 history 记录不够用，路由将不会进行跳转
router.go(-100)
router.go(100)

```

接下来我们仍然对第一个案例稍加修改：

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
      <router-link to="index">首页</router-link>
      <router-link to="article">文章</router-link>
    </div>
    <button @click="go(1)">前进一步</button>
    <button @click="go(-1)">后路一步</button>
    <button @click="go(3)">前进三步</button>
    <button @click="go(-3)">后路三步</button>
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
  { path: '/index', component: Index },
  { path: '/article', component: Article }
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
      go(n) {
        this.$router.go(n)
      }
    }
  })
</script>
</html>

```

**代码解释：**

HTML 代码第 15-18 行，我们定义了四个按钮，并给他们点击事件 go。

JS 代码第 29-31 行，我们定义来 go 函数，通过 router.go 实现路由跳转。

## 5. 小结

本节，我们带大家学习了 VueRouter 如何通过方法来实现跳转。主要知识点有以下几点：

* 通过 router.push 跳转到指定路由。
* 通过 router.replace 替换当前路由记录跳转指定路由。
* 通过 router.go 实现路由的前进、后退功能。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
