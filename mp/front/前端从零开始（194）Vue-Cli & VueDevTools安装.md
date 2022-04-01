# Vue-Cli & VueDevTools 安装

## 1. 前言

本小节我们会介绍 Vue 脚手架工具 `Vue-Cli` 以及调试工具 `VueDevTools` 的安装和使用。

## 2. 什么是 Vue-Cli

`vue-cli` 是由 `Vue` 提供的一个官方 cli，专门为单页面应用快速搭建繁杂的脚手架。它是用于自动生成 vue.js+webpack 的项目模板。

### 2.1 安装 Cli

```javascript
// npm 安装
npm install -g @vue/cli

// yarn 安装
yarn global add @vue/cli
```

查看是否安装成功：

```javascript
vue -V
// 正确显示版本号
```

### 2.2 初始化项目

```javascript
vue create vue-learn
```

回车之后会出现以下画面

```javascript
Vue CLI v3.9.3
┌────────────────────────────┐
│  Update available: 3.10.0  │
└────────────────────────────┘
? Please pick a preset: (Use arrow keys)
❯ default (babel, eslint)
  Manually select features
```

* default (babel, eslint) 默认套餐，提供 babel 和 eslint 支持。
* Manually select features 自己去选择需要的功能，提供更多的特性选择。比如如果想要支持 TypeScript ，就应该选择这一项。
* 使用上下方向键来选择需要的选项。
* 使用 manually 来创建项目，选中之后会出现以下画面。

```javascript
Vue CLI v3.9.3
┌────────────────────────────┐
│  Update available: 3.10.0  │
└────────────────────────────┘
? Please pick a preset: Manually select features
? Check the features needed for your project: (Press <space> to select, <a> to t
oggle all, <i> to invert selection)
❯◉ Babel
 ◯ TypeScript
 ◯ Progressive Web App (PWA) Support
 ◯ Router
 ◯ Vuex
 ◯ CSS Pre-processors
 ◉ Linter / Formatter
 ◯ Unit Testing
 ◯ E2E Testing
```

依然是上下键选择，空格键选中。

对于每一项的功能，此处做个简单描述：

* TypeScript 支持使用 TypeScript 书写源码。
* Progressive Web App (PWA) Support PWA 支持。
* Router 支持 vue-router 。
* Vuex 支持 vuex 。
* CSS Pre-processors 支持 CSS 预处理器。
* Linter / Formatter 支持代码风格检查和格式化。
* Unit Testing 支持单元测试。
* E2E Testing 支持 E2E 测试。

第一个 typescript 暂时还不会，先不选，这次选择常用的。

```javascript
◉ Babel
◯ TypeScript
◯ Progressive Web App (PWA) Support
◉ Router
◉ Vuex
◉ CSS Pre-processors
◉ Linter / Formatter
❯◉ Unit Testing
◯ E2E Testing
```

回车之后让选择 CSS 处理器，这里选择 Less。

```javascript
? Use history mode for router? (Requires proper server setup for index fallback
in production) Yes
? Pick a CSS pre-processor (PostCSS, Autoprefixer and CSS Modules are supported
by default): (Use arrow keys)
❯ Sass/SCSS (with dart-sass)
 Sass/SCSS (with node-sass)
 Less
 Stylus
```

接下来选择 eslink，我选择了 eslink+prettier：

```javascript
? Pick a linter / formatter config:
 ESLint with error prevention only
 ESLint + Airbnb config
 ESLint + Standard config
❯ ESLint + Prettier
```

选择代码检查方式，第一个是保存的时候就检查，第二个是提交上传代码的时候才检查。

```javascript
? Pick additional lint features: (Press <space> to select, <a> to toggle all, <i
> to invert selection)
❯◉ Lint on save
 ◯ Lint and fix on commit
```

选择单元测试，这个我不懂，随便先选个 jest：

```javascript
? Pick a unit testing solution:
  Mocha + Chai
❯ Jest
```

配置文件存放的地方，选择 package.json：

```javascript
? Where do you prefer placing config for Babel, PostCSS, ESLint, etc.?
  In dedicated config files
❯ In package.json
```

是否保存这次配置，方便下次直接使用，一般都是选择 Y。

```javascript
? Save this as a preset for future projects? (y/N)
```

配置完成之后就开始创建一个初始项目了：

启动

```javascript
cd vue-learn
npm run serve
```

## 3. 什么是 VueDevTools

vue-devtools 是一款基于 chrome 游览器的插件，用于调试 vue 应用，这可以极大地提高我们的调试效率。

### 3.1 VueDevTools 安装

1. chrome 商店直接安装

vue-devtools 可以从 chrome 商店直接下载安装。我们可以打开 chrome 商店，搜索 vue-dev-tools，点击 “添加至 chrome” 即可。

2. 手动安装

    1. 将 vue-devtools 克隆到本地。

    ```javascript
    git clone https://github.com/vuejs/vue-devtools.git
    ```

    2. 安装项目所需要的安装包

    ```javascript
    npm install
    ```

    3. 编译项目文件

    ```javascript
    npm run build
    ```

    4. 添加至 chrome 浏览器

    ```javascript
    1、游览器输入地址“chrome://extensions/” 进入扩展程序页面，
    2、点击“加载已解压的扩展程序...”按钮
    3、选择vue-devtools>shells下的chrome文件夹。
    ```

### 3.2 VueDevTools 的使用

当我们添加完 vue-devtools 扩展程序之后，我们在调试 vue 应用的时候，chrome 开发者工具中会看一个 vue 的一栏，点击之后就可以看见当前页面 vue 对象的一些信息。vue-devtools 使用起来还是比较简单的，上手非常地容易。

## 4. 小结

在本小节我们介绍了什么是 `vue-cli`，如何安装 `vue-cli`。介绍了调试工具 `VueDevTools` 的安装和使用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
