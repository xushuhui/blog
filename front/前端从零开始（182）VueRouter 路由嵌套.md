# VueRouter 路由嵌套

## 1. 前言

本小节我们介绍如何嵌套使用 VueRouter。嵌套路由在日常的开发中非常常见，如何定义和使用嵌套路由是本节的重点。同学们在学完本节课程之后需要自己多尝试配置路由。

## 2. 配置嵌套路由

实际项目中的应用界面，通常由多层嵌套的组件组合而成。同样地，URL 中各段动态路径也按某种结构对应嵌套的各层组件，例如：

```javascript
/article/vue                          /article/react
+------------------+                  +-----------------+
| Article          |                  | Article         |
| +--------------+ |                  | +-------------+ |
| | Vue          | |  +------------>  | | React       | |
| |              | |                  | |             | |
| +--------------+ |                  | +-------------+ |
+------------------+                  +-----------------+
```

借助 vue-router，使用嵌套路由配置，就可以很简单地表达这种关系。

在上一小节中我们学习了如何配置一个路由信息：

```javascript
  {
    path: '路由地址',
    component: '渲染组件'
  }
```

要配置嵌套路由，我们需要在配置的参数中使用 children 属性：

```javascript
  {
    path: '路由地址',
    component: '渲染组件',
    children: [
      {
        path: '路由地址',
        component: '渲染组件'
      }
    ]
  }
```

### 2.1 基本使用

接下来我们对上一小节的例子来做一个改造：在文章页面，我们对文章进行分类，提供两个链接按钮 vue、react，点击可以跳转到对应的文章列表，具体代码示例如下：

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
  template: `<div>
      <div>
        <router-link to="/article/vue">vue</router-link>
        <router-link to="/article/react">react</router-link>
      </div>
      <router-view></router-view>
    </div>`,
})

const VueArticle = Vue.component('vueArticle', {
  template: `<ul><li>1. Vue 基础学习</li><li>2. Vue 项目实战</li></ul>`,
})

const ReactArticle = Vue.component('reactArticle', {
  template: `<ul><li>1. React 基础学习</li><li>2. React 项目实战</li></ul>`,
})

const routes = [
  { path: '/index', component: Index },
  {
    path: '/article',
    component: Article ,
    children: [
      {
        path: 'vue',
        component: VueArticle ,
      },
      {
        path: 'react',
        component: ReactArticle ,
      }
    ]
  }
]

const router = new VueRouter({
  routes: routes
})

  var vm = new Vue({
    el: '#app',
    router,
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

JS 代码第 9-17 行，我们定义了组件 Article，组件内部使用 `<router-link></router-link>` 定义出来两个跳转链接，使用 `<router-view></router-view>` 来渲染匹配组件。

JS 代码第 19-21 行，我们定义了组件 VueArticle.

JS 代码第 23-25 行，我们定义了组件 ReactArticle。

JS 代码第 27-43 行，我们定义了路由数组，在 ‘/article’ 中配置来嵌套路由 children

JS 代码第 44-46 行，创建 router 实例，然后传 `routes` 配置。

JS 代码第 49 行，通过 router 配置参数注入路由。

### 2.2 定义路由地址

在上述的例子中，我们通过 ‘/article/vue’ 来访问嵌套路由，但是有时候你可能不希望使用嵌套路径，这时候我们可以对上面例子中的配置信息做一点修改：

```javascript
const routes = [
  { path: '/index', component: Index },
  {
    path: '/article',
    component: Article ,
    children: [
      {
        path: '/vueArticle',
        component: VueArticle ,
      },
      {
        path: '/reactArticle',
        component: ReactArticle ,
      }
    ]
  }
]
```

以 **‘/’** 开头的嵌套路径会被当作根路径，因此，我们访问 ‘/vueArticle’ 就相当于访问之前的 ‘/article/vue’。

## 3. 小结

本节，我们带大家学习了 VueRouter 嵌套路由的使用方法，主要知识点有以下几点：

* 通过路由配置的 children 属性定义和使用嵌套路由。
* 通过修改路由配置的 path 属性定义嵌套路由的跳转地址。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
