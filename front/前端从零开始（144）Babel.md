# Babel

> Babel 是一个 JavaScript 编译器。
>
>
> Babel 是一个工具链，主要用于将 ECMAScript 2015+ 版本的代码转换为向后兼容的 JavaScript 语法，以便能够运行在当前和旧版本的浏览器或其他环境中。

`Babel` 由 `Node.js` 驱动，可以把 `ES6+` 的代码编译成 `ES5`、`ES3` 的代码。

配合插件和工具，`Babel` 还能处理一些非规范的语法，处理 `JSX`、`TypeScript` 等。

使用 `Babel`，可以让开发者在开发环境享受新特性带来的优质的编码体验，又能让代码在生产环境兼容大部分浏览器或其他宿主环境，如 `Node.js`。

## 1. 安装 Babel

`Babel` 可以直接使用 `npm` 安装。

首先进入到一个空目录，初始化一个项目，然后安装 `Babel`。

```javascript
npm init -y

npm i @babel/core @babel/cli -D
```

![图片描述](https://img.mukewang.com/wiki/5f1d2c980ab43bf211280706.jpg)

安装后可以查看 `package.json` 中的内容，因为安装时候提供了 `-D` 参数，表示安装到开发依赖中，`-D` 是 `--save-dev` 参数的别名。

安装完后就可以利用 `npm scripts` 使用 babel 命令行工具了。

## 2. 使用 Babel

先创建一个带有 `ES6` 语法的 `JS` 文件：

```javascript
// index.js
const fn = (a, b) => a + b;

const added = fn(1, 2);
```

然后修改一下 `package.json` 的 `scripts` 选项。

```javascript
// package.json
"scripts": {
  "compile": "babel index.js"
}
```

保存后就可以进入终端，输入命令 `npm run compile`，`npm` 就会去执行 `compile` 对应的命令，也就是使用 `Babel`，对 `index.js` 文件进行编译。

![图片描述](https://img.mukewang.com/wiki/5f1d2cc20a435c9e11280706.jpg)

但执行后会发现代码并没有变成非 `ES6` 的代码，那是因为没有告诉 `Babel` 想要将现有代码编译成什么样的代码，这时候就需要提供一个配置文件。

### 2.1 babel 配置文件

在项目根目录新建一个 `babel.config.js` 的文件，以前的 `Babel` 版本对应的默认配置文件名是 `.babelrc`，`Babel@7` 之后官方更推荐使用 `babel.config.js`。

然后安装一下 Babel 的预设配插件，Babel 提供了许多预设，官方命名为 `Presets`，现有的 `Preset` 非常多，如 `preset-es2015`、`preset-es2016`、`preset-es2017`，每一个对应了一个 `ES` 的版本，现在 `ES` 每年都会有一个版本，所以插件会越来越多。

`preset-env` 这个预设为解决上述和一些其他问题而被官方推出，如果不给任何配置，这个 `preset` 的效果与 `preset-latest` 相同。

所以要安装一下 `@babel/preset-env` 这个预设。

```javascript
npm i @babel/preset-env -D
```

然后修改配置文件：

```javascript
// babel.config.js

module.exports = {
  presets: [
    '@babel/preset-env',
  ],
};
```

告诉 `Babel`，`presets` 使用 `@babel/preset-env`，`presets` 配置项是个数组，可以有多个 `presets`，目前只用到一个。

做好这些工作，再去命令行跑 `npm run compile`：

![图片描述](https://img.mukewang.com/wiki/5f1d2cfb0a2bd33711280706.jpg)

现在这样就成功的把 `ES6` 代码编译成了 `ES5`。

但现在代码并没有输出成文件，而是直接显示在了终端里，最好是能输出到一个文件里，这个需求改下命令就能做到。

```javascript
// package.json
"scripts": {
  "compile": "babel index.js -o index.compiled.js"
}
```

增加 `-o` 参数，表示要输出到哪个文件，然后再执行 `npm run compile`：

![图片描述](https://img.mukewang.com/wiki/5f1d2d100a440f1811280706.jpg)

> 为什么 `npm run compile` 可以去调用 `babel` 命令行工具？
>
>
> 在 `node_modules` 下有个 `.bin` 目录，在执行一个命令的时候，`npm` 会先去 `node_nodules/.bin` 下找对应的命令行工具，如果有就会调用执行。而 `babel` 会出现在 `.bin` 目录下，又与 `babel` 项目中的 `package.json` 文件的 `bin` 配置项有关，具体的可以再去官方文献中扩展。

## 3. 小结

`Babel` 是一个很重要的工具，是现在编译方案的首选。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
