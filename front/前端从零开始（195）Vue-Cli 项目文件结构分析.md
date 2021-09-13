# Vue-Cli 项目文件结构分析

## 1. 前言

在上一小节中，我们介绍了 Vue-Cli 初始化项目，本小节我们一起来分析以下 Vue-Cli 创建项目的文件结构。

## 2. 目录结构

首先我们先看以下用 Vue-Cli 创建项目的整体结构：

![图片描述](https://img.mukewang.com/wiki/5ed845a109b4cac203360708.jpg)

项目相关的代码，我们都放在 src 的文件夹中，该文件夹中的结构如下：

* assets 是资源文件夹，通常我们会把图片资源放在里面。
* components 文件夹通常会放一些组件。
* router 文件夹里面放的是 VueRouter 的相关配置。
* store 文件夹里面放的是 Vuex 的相关配置。
* views 文件夹里面通常放置页面的 .vue 文件。
* App.vue 定义了一个根组件。
* main.js 是项目的入口文件。

Vue-Cli 给我们提供了一个默认的项目文件结构，当然你并不是一定要按照这个文件结构来编写项目，你完全可以根据项目的需要或者个人喜好对项目结构加以改写。

## 3. 运行项目

我们打开之前通过脚手架创建的项目，在项目的根目录下运行：

```javascript
npm run serve
```

出现界面后：

![图片描述](https://img.mukewang.com/wiki/5ed84578083edff013560334.jpg)

我们可以打开浏览器预览项目：

![图片描述](https://img.mukewang.com/wiki/5ed8458708a747b713841494.jpg)

## 4. 详细分析

可能，同学们对这样一个项目结构还不是很熟悉，接下来，我们将对主要的几个项目文件详细分析。

### 4.1 main.js 入口文件分析

打开 main.js 入口文件，我们可以看到这段代码：

```javascript
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
```

我们通过 import 引入了 App 组件、VueRouter 配置 router、Vuex 配置 store。

```javascript
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
```

通过 new Vue () 创建 Vue 实例，并将 router、store 配置传入。通过 render 函数渲染组件 App。并将 Vue 实例挂载到 id 为 app 的 div 上。

### 4.2 router 文件分析

打开 router/index.js 文件，我们可以看到路由配置信息：

```javascript
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    component: () =>
      import( "../views/About.vue")
  }
];
```

定义了两个路由：

* 路由 ‘/’ 匹配组件 Home。
* 路由 ‘/about’ 匹配组件 About。

### 4.3 store 文件分析

打开 store/index.js 文件，我们可以看到 Vuex 的配置信息：

```javascript
import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {}
});
```

## 5. 小结

在本小节我们介绍了脚手架根据初始化项目的文件结构。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
