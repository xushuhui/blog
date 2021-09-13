# JavaScript 字符串

字符串是一种数据类型，由字符组成，用于表示文本数据。

## 1. 创建字符串

使用字符串字面量就可以创建字符串。

```javascript
var str1 = '';

var str2 = "";
```

以上例子创建了两个空字符串。字符串可以用单引号包裹，也可以用双引号包裹，效果是一样的。

大部分前端程序员都会选择单引号，这么做的原因是为了契合 `HTML` 的书写习惯，拼接 HTML 的时候，HTML 元素的属性可以直接使用双引号包裹。

```javascript
var html = '<p class="tip">更多请<a href="/detail" target="_blank">了解详情</a></p>';
```

如果换成双引号则需要转义：

```javascript
var html = "<p class=\"tip\">更多请<a href=\"/detail\" target=\"_blank\">了解详情</a></p>";
```

## 2. 字符串转义

转义可以理解成字面意思，即转换字符的含义。

比如想在字符串里描述换行符，就会使用 `\n`。

```javascript
var str = '第一行\n第二行\n第三行';

console.log(str); // 将会输出换行的字符串
```

通常转义字符都是通过 `\` 开头的。

同理，如果需要在双引号包裹的字符串中使用双引号，或者单引号包裹的字符串中使用单引号，就需要对引号进行转义。

```javascript
var str = '我'要'显'示'一'堆'单'引'号';
```

这样就会报错，JavaScript 无法知道这些引号的含义。

```javascript
var str = '我\'要\'显\'示\'一\'堆\'单\'引\'号';
```

通过 `\n` 表示一个单引号，就可以让 JavaScript 理解，需要在字符串里描述一个单引号。

### 2.1 转义表

JavaScript 支持以下字符的转义：

|\'|单引号|
|\"|双引号|
|\&|和号  |
|\\|反斜杠|
|\n|换行符|
|\r|回车符|
|\t|制表符|
|\b|退格符|
|\f|换页符|
|\'|单引号|
|\"|双引号|
|\&|和号  |
|\\|反斜杠|
|\n|换行符|
|\r|回车符|
|\t|制表符|
|\b|退格符|
|\f|换页符|

## 3. 使用场景

### 3.1 用于描述文案

文案描述是最常见的情景之一，文案本身就是字符串，使用字符串是最契合场景的。

```javascript
function gameover(age) {
  // 判断是否游戏结束
  return age > 300;
}

var isGameover = gameover(666);

if (isGameover) {
  console.log('游戏结束');
}
```

`游戏结束`就是一个确确实实的文案，用于展现。

通过输出日志简单调试的也是，也会用到字符串作为文案进行观察。

```javascript
function something() {
  console.log('循环开始之前');

  var i
  for (i = 0; i < 10; i--) {
    console.log('循环中，第 ' + (i + 1) + ' 次循环');
  }

  console.log('循环结束');
}

something();
```

以上模拟一个程序卡死的状态，通过调试输出很容易发现是由死循环导致的，原因是将循环条件判断后执行的表达式写错了。

### 3.2 拼接 HTML

拼接 `HTML` 是前几年大部分前端程序员做的最多的事情之一。

现在有了许多前端框架，解放了很多 `HTML` 拼接的工作。

早期想渲染一个列表，需要前端程序员在 JavaScript 中拼接 HTML 后再使用。

```javascript
function getList(content) {
  return '<li class="list-item">' + content + '</li>';
}

var arr = [];

var i;
for (i = 1; i <= 10; i++) {
  arr.push(getList('我是第' + i + '条'));
}

document.body.innerHTML = [
  '<ul class="list">',
  arr.join(''),
  '</ul>',
].join('');
```

![图片描述](https://img.mukewang.com/wiki/5e7ae5740a0cb1d920121314.jpg)

> 数组的 `join` 方法会将参数作为分隔符，将数组成员连接成一个字符串，默认的分隔符是逗号。

不论是制作插件，还是业务需求的页面元素、文案调整，都需要拼接 HTML。

### 3.3 其他任意场景

通常字符串可以用在`任意场景`，因为可以人为的赋予他任何含义。

如使用字符串的 `'true'`、`'false'` 来表示布尔值。

```javascript
var isMan = 'false';

if (isMan === 'false') {
  console.log('不是男的');
}
```

或者可以使用字符串描述一个对象：

```javascript
var person = '小明,男,24岁,吃饭、睡觉、打游戏';

name = person.split(',')[0];

console.log(name);
```

通过一定的规则来确定字符串的含义，如上面就是用逗号分隔，含义依次为`姓名，性别，年龄，爱好`。

取值的时候通过 `split` 方法，将字符串按指定的字符分隔成数组。

这种情况通常会用在后台需要存储简单的数据结构，前端拿到的数据也许就是这样的，需要自己切分。

**正常情况下前端开发者在开发过程中不建议这样做，因为有更好的数据类型来描述这些内容，JavaScript 提供了布尔值、对象来更好的、灵活的应对这些场景。**

## 4. 访问字符串

**字符串是无法修改的**，只能进行访问。

```javascript
var str = '我是字符串';

console.log(str); // 输出："我是字符串"
```

这是直接访问整个字符串，还可以访问中间某一个字符。

```javascript
var str = '一二三四五六七，7654321';

// 获取第一个字符
console.log(str[0]); // 输出："一"

// 获取最后一个字符
console.log(str[str.length - 1]); // 输出："1"
console.log(str.split('').pop()); // 输出："1"
```

通过`字符串[下标]`的形式可以访问到某一个字符，这种访问通常会用在字符串遍历上。

## 5. 获取字符串长度

字符串可以直接通过 `length` 属性获取长度。

```javascript
var string = '1234567';

var len = string.length;

console.log(len); // 输出：7
```

## 6. 常用的字符串拼接

### 6.1 使用 + 连接字符串

最常见的字符串拼接就是使用 `+` 符号：

```javascript
var str1 = '我是';
var str2 = '小明';

console.log(str1 + str2);
```

### 6.2 数组的 join 方法

还有一种是使用数组来拼接字符串：

```javascript
var str1 = '我是';
var str2 = '一只鱼';

console.log([str1, str2].join('')); // 输出：我是一只鱼
```

这种方式也很常见。

> 注意：许多文献中会提到这样拼接字符串的性能更好，效率更高。大部分浏览器确实如此，其引擎针对性的做了特殊优化，当然只有在非常大量的字符串拼接时才能感知到性能上的区别。

### 6.3 字符串的 concat 方法

```javascript
var str1 = '教练';
var str2 = '我想';
var str3 = '写代码';

var str4 = str1.concat(str2, str3);

console.log(str4);
```

使用 `concat` 可以接受任意个字符串作为参数，最后会拼接成一个字符串。

> ES6 提供了模版字符串，在模版字符串中拼接更加灵活。

## 7. 使用 String 对象创建字符串

`String` 对象也可以创建字符串，但不常用。

```javascript
var str = new String('do sth.');

console.log(typeof str); // 输出：object
console.log(str.concat('gugu')); // 输出：do sth.gugu
```

可以观察到，使用 `String` 对象生成的字符串，实际上是一个`对象`，

但使用 `String` 对象生成的字符串对象使用的时候基本和字符串无异。这部分会涉及到部分装箱拆箱的知识，具体可以查阅相关章节。

## 8. 与字符串相关的常用方法

以下是一些常用方法，使用频率较高：

|方法|描述|
|----|----|
|[replace](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/replace)                    | replace 方法返回一个由替换值替换一些或所有匹配的模式后的新字符串。                                            |
|[match](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/match)                        | match 方法检索返回一个字符串匹配正则表达式的的结果。                                                          |
|[split](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/split)                        | split 方法使用指定的分隔符字符串将一个String对象分割成子字符串数组，以一个指定的分割字串来决定每个拆分的位置  |
|[substring](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/substring)                | substring 方法返回一个字符串在开始索引到结束索引之间的一个子集, 或从开始索引直到字符串的末尾的一个子集。      |
|[toLocaleLowerCase](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/toLocaleLowerCase)| toLocaleLowerCase 方法根据任何指定区域语言环境设置的大小写映射，返回调用字符串被转换为小写的格式。            |
|[toLocaleUpperCase](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/toLocaleUpperCase)| toLocaleUpperCase 使用本地化（locale-specific）的大小写映射规则将输入的字符串转化成大写形式并返回结果字符串。 |

## 8. 小结

字符串常常用于拼接 HTML ，且前端开发者习惯使用单引号来包裹字符串。

大量的字符串拼接，利用数组的 `join` 比使用 `+` 连接字符串性能更好，前提是非常大量。

使用 `join` 更大的好处是多行的时候换行方便。

字符串基本可以胜任任何数据场景，但一般不会这么做，因为 JavaScript 提供了多种数据类型来应对各种业务场景。

4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
