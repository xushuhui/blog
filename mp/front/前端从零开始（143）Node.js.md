# Node.js

> Node.js 是一个基于 Chrome V8 引擎的 JavaScript 运行时。

`Node.js` 并不是运行在浏览器里的一个库或框架。

`Node.js` 可以提供了一系列服务端能力，如 `HTTP 服务`、`读写本地文件`等，开发者可以利用 `JavaScript` 来使用这些能力，因为前端开发者的主要语言就是 `JavaScript`，所以利用 `Node.js` 可以降低学习成本，让前端开发者更容易接触到服务端开发。

## 1. 安装 Node.js

`Node.js` 需要单独安装，进入 `Node.js`[官网](https://nodejs.org/en/download/) 获取对应平台的安装包下载即可。

如果是为了学习使用，建议使用最新版，支持更多的特性。

`Node.js` 的安装过程和普通软件相似，安装完毕后可以通过命令行测试是否安装成功。

### 1.1 windows 下打开命令提示符

windows 下可以直接在开始中进行搜索，搜索命令提示符，打开搜索结果。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d1eff09383ce308530636.jpg)

打开后在命令提示符中输入 `node -v` 并回车，如果有正确输出安装的 `Node.js` 的版本号，则表示安装成功。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d1fe609e3fb3e06940370.jpg)

### 1.2 Mac OS 下打开终端

在 `Mac OS` 操作系统下，打开 `聚焦搜索` ，输入 `终端` 后回车即可打开终端。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d1f7c09e393bb13600860.jpg)

打开终端后输入 `node -v`，如果正确输出了版本号，则表示安装成功。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d1f9309e345e911700730.jpg)

> 因为两个平台下的命令几乎一致，后续内容不再区分平台，统一使用 Mac OS 下的终端。

## 2. npm

`Node.js` 安装后，会同时安装 `npm`，和查看 `Node.js` 的版本一样，在终端里输入 `npm -v`，即可查看到 `npm` 的版本号。

```javascript
npm -v
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d206109039ca706980174.jpg)

`npm` 的全称是 `Node Package Manager`，翻译过来就是 `node.js` 的包管理工具。

一个项目开发过程中需要许多第三方的代码，比如一些框架，这些框架大部分都会被做一个 `npm 包` 发布，通过 `npm` 命令行工具就可以安装到本地项目中，进行使用。

同样的可以自己开发一些 `包`，发布到 `npm`，然后造福社会。

### 2.1 创建一个包

一个 `npm 包` 由一个 `package.json` 文件描述。

`package.json` 所在的位置通常会被作为项目的根目录。

可以通过 `npm` 提供的命令创建一个 `package.json`，可以先创建一个项目目录，然后在终端中进入到这个目录，使用 `npm init -y` 命令，就可以创建一个最简单的 `package.json`。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d20a009081cd210980626.jpg)

当然现在的工程化的前端项目，也会用 `package.json` 来描述项目信息，来管理依赖、工作里等，一个包不一定要发送到 npm 上。

### 2.2 package.json 简析

`package.json` 中有许多项目，描述了不同的信息。

这里介绍几个常用的字段。

#### 2.1.1 devDependencies

```javascript
{
  "devDependencies": {
    "@babel/core": "^7.2.2",
    "@babel/plugin-proposal-decorators": "^7.3.0"
  }
}
```

`devDependencies` 记录了一些开发依赖，这些依赖在生产环境不会使用。

如一些代码检查工具，因为只有在开发、编译阶段需要检查代码，最终跑上线的代码已经在运行了，不需要再进行语法、规范检查。

#### 2.1.2 dependencies

```javascript
"dependencies": {
  "md5": "^2.2.1"
}
```

`dependencies` 记录了生产、开发环境都会用到的依赖。

如 `jquery`，这样实实在在跑在项目里，支撑起项目功能的依赖。

#### 2.1.3 scripts

```javascript
{
  "scripts": {
    "dev": "echo \"development\""
  }
}
```

`scripts` 可以说是直接接触到的最常用的一个配置项。

配置 `script` 可以完成一些简单的工作流，或者把复杂的命令配置为一个别名。

如配置的 `dev` 项，就可以通过在终端输入 `npm run dev` 来调用。

同时 `script` 还提供了前置和后置钩子，具体可以参阅文档。

## 3. 体验 Node.js

知道了 `npm` 和 `node.js` 的关系，了解了 `package.json` 的作用，就可以来尝试使用 `Node.js` 做应用了。

### 3.1 读写文件

在终端使用 `node js文件.js` 就可以使用 `Node.js` 执行 `.js` 文件。

在 `Node.js` 中，处理文件需要使用 `fs` 模块。这个模块是 `Node.js` 自带的，可以直接引入。

首先创建一个 `.js` 文件，然后在同级目录下，创建任意的文本文件。

```javascript
// app.js
var fs = require('fs');

var text = fs.readFileSync('./text.txt', 'utf-8');

console.log(text);
```

写完这三行代码，保存后就可以去终端执行代码了。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d20c60ad937e111080630.jpg)

`fs` 模块的 `readFileSync` 方法，表示可以同步的读取一个文件，然后将读取到的文件放在 `text` 变量中。

随后将读取到的文件进行输出。

`fs` 文件非常重要，一些特殊的配置文件，如 `YAML` 配置文件，就需要 `fs` 模块读取，在或者是静态资源，也需要用 `fs` 模块读取或者写入。

`fs` 模块通过 `require` 引入，`Node.js` 支持 `commonjs规范`，通过 `commonjs规范` 来处理模块。

> 新版的 `Node.js` 已经支持 `ES Module`，需要修改 `package.json` 中的 `type` 选项为 `module`。

### 3.2 使用 npm 包

`md5` 是很常用的加密算法，但通常又不可能自己去实现一遍，快速迭代的项目可以 `拿来主义`，有现成的方案直接拿来用。

在 `npm` 上有一个 [md5](https://www.npmjs.com/package/md5) 包，就可以拿来计算 `md5`。

首先在一个空目录创建一个 `package.json`，用来描述项目信息，然后安装 `md5`。

```javascript
npm init -y

npm i md5
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d20e70a67e24210920522.jpg)

然后新建一个 `.js` 文件，尝试着使用 `md5` 这个包。

```javascript
// app.js
var md5 = require('md5');

var password = '123456';

var encode = md5(password);

console.log(encode);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1d20fa09a5299008140174.jpg)

安装好的包直接通过 `require` 引入，然后跟着包的文档使用即可。

## 4. 小结

`Node.js` 目前生态主要集中在前端工具上，大众的前段工程化的解决方案，都是由 `Node.js` 来驱动完成。

作为前端开发者，基础的使用 `Node.js` 已经是必备技能，许多针对项目的自动化的流程工具，都需要前端开发者自己动手实现。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
