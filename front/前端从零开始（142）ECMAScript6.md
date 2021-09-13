# ECMAScript 6

> 2015 年 6 月 17 日，ECMA 国际组织发布了 ECMAScript 的第六版，该版本正式名称为 ECMAScript 2015，但通常被称为 ECMAScript 6 或者 ES6。(MDN)

ES6 的发布让 `JavaScript` 的便携体验有了里程碑式的飞跃。

各种概念性的特性都被纳入标准，如 `Promise`、`模块化` 。

`Google Chrome` 浏览器支持绝大部分的 ES6 特性，可以直接在浏览器上体验。


## 1. 区别体验

ES6 经常被吹的天花乱坠，实则还是语言相关的知识，他在以前的基础上，对语法、全局对象、特性等做了扩充。

如：变量可以采用 `let` 关键字声明。

### 1.1 let

let 是 ES6 中提供了可以声明变量的关键字。

```javascript
let number = 1;

console.log(number);
```

![图片描述](https://img.mukewang.com/wiki/5f0de9dc092c96b608300170.jpg)

使用 `let` 声明的变量，对上层作用域的污染更少。

```javascript
let number = 10;

if (number < 20) {
  var base1 = 1;
  let base2 = 10;

  number = number + base1 + base2;
} else {
  number = 0;
}

console.log(base1); // 输出：1
console.log(base2); // 输出：ReferenceError: base2 is not defined
```

![图片描述](https://img.mukewang.com/wiki/5f0de9fe09a2edea11440474.jpg)

可以发现在 `if` 之外，`base1` 还是能被访问到的，但是 `base2` 却不行。

这是因为 `ES6` 引入了块级作用域的概念，`let` 声明的变量，只在块级作用域内有效。

### 1.2 对象简洁表示法

在 ES6 之前，对象字面量的属性一定要书写上属性值和属性名。

在 ES6 中，如果对象的属性名和存放属性值的变量相同，则只需要写一次。

```javascript
// 在 ES5 中
function getInfo() {
  // 通过某种方式拿到的数据
  var username = '张三';
  var enemy = '罗老师';
  var gender = '男';
  var age = 12;

  return {
    username: username,
    enemy: enemy,
    gender: gender,
    age: age,
  };
}

// 在 ES6 中
function getInfo() {
  // 通过某种方式拿到的数据
  var username = '张三';
  var enemy = '罗老师';
  var gender = '男';
  var age = 12;

  return {
    username,
    enemy,
    gender,
    age,
  };
}
```

可以看到 ES6 使得代码更清晰，也可以直接看出对象的属性名和存放属性值的变量是同名的。

同样的，方法也提供了简写的方式：

```javascript
var obj = {
  say() { // ES6
    console.log('说话');
  },
  walk: function() { // ES5
    console.log('走路');
  },
},
```

## 2. 兼容性

虽然目前还在维护的最新版的浏览器几乎都支持了大部分 `ES6` 特性，但国内生态还不允许直接将 ES6 代码运行于线上，所以就需要一定的解决方案，使开发者开发过程中全面使用 ES6，但是线上又是运行 `ES5`、`ES3` 的代码。

特性上会采用 `shim` 的方式，大部分情况下概念会将它与 `polyfill` 混用，可以理解成给浏览器打补丁，让旧版的浏览器支持新版的特性，如 ES6 提供的 `Object.assign` 方法，旧版的浏览器是没有的。通过 `polyfill`，使用 `ES5` 将刚该方法实现后，在放到 `Object` 对象下，变相的让浏览器支持新特性。

但还有一些特性是 `polyfill` 很难解决的，特比是语法特性，如 `let` 关键字。

这些特性就会采用 `编译` 的方式来解决，如将 `let` 替换成 `var`，这一块最常用的工具目前是 `babel`。

所以如果项目是需要上线运行，并且目标用户群体范围非常广，尽量不要上线 `ES6` 代码，上线前也做好各个平台的浏览器测试。

## 3. 小结

ES6 的出现，让 `JavaScript` 的代码质量有了质的飞跃，也解决、优化了以前存在的许多问题。

现在找前端岗位的工作，`ES6` 是必备技能，务必要掌握。




### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
