# TypeScript

> TypeScript 是 JavaScript 的超集。
>
>
> TypeScript 是 JavaScript 类型的超集，它可以编译成纯 JavaScript。
>
>
> TypeScript 可以在任何浏览器、任何计算机和任何操作系统上运行，并且是开源的。

TypeScript 包含了 `JavaScript` 所有的特性，同时做出了扩展，实现了许多还处于提案的或非 `ECMAScript` 的内容。（以下对 TypeScript 简称为 TS）。

除了实现 `ES` 的标准外，TS 最主要的就是加入了类型，通过提供静态类型，这个特性在使得代码变得更健壮外，也可以让 IDE 或者编辑器更容易推导出类型。

## 1. 安装 TypeScript

`TS` 可以使用 `npm` 安装。

首先在空目录初始化一个 `package.json`：

```javascript
npm init -y

npm i typescirpt -D
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f2f94770a73672011280692.jpg)

由于 `TS` 也是在开发过程中才会使用，上线前都会编译成 `JavaScript`，所以作为开发环境下的依赖。

然后修改 `package.json` 的 `scripts` 配置项：

```javascript
// package.json
"scripts": {
  "ts": "tsc ./index.ts"
}
```

`tsc ./index.ts` 命令回去当前目录下找到 `index.ts` 文件，并编译成 `.js` 文件

## 2. 使用 `TS`

首先创建一个 `index.ts`：

```javascript
const PI: number = 3.1415926535;
```

这是一个数字类型常量，通过 `npm run ts` 进行编译：

```javascript
npm run ts
```

这样就把 `.ts` 文件变成了可以运行在浏览器的 `.js` 文件。

类型是 `TS` 最重要的特性之一，特别是在开发过程中，这一特性会帮助开发工具进行类型推断：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f2f94f20a9835f911280692.jpg)

## 3. 配置文件

`TS` 提供了一套配置来描述编译行为，如编译到哪个标准的代码，排除哪些代码不编译，是否开启某些规则，如不允许隐式的出现 `any` 等。

[官方](https://www.typescriptlang.org/docs/handbook/tsconfig-json.html)对配置文件进行了详细说明，通常也会采用配置文件的形式来使用。

## 4. 小结

`TypeScript` 可以说是现阶段前端的必备技能，许多开源框架也采用 `TypeScript` 来编写，如 `Angular`、`Vue` 等。

`TypeScript` 不单单提供了静态类型，还扩充了 `类` 的能力，提供了非常多的特性。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
