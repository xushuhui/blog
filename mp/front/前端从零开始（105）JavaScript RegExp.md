# JavaScript RegExp

> RegExp 构造函数创建了一个正则表达式对象，用于将文本与一个模式匹配。

RegExp 的实例用于生成一个正则表达式，利用正则表达式从字符串中匹配想要的内容。

正则表达式不是 JavaScript 的一个子内容，也并非 JavaScript 独有，需要学习正则表达式可以查阅相对应的文献。

## 1. 创建实例

创建一个 RegExp 实例，只需要将其当作构造函数使用：

```javascript
var regexp = new RegExp(/^a*$/);

regexp.test('aaaa'); // true
regexp.test('a'); // true
regexp.test('a1'); // false
```

上面这个例子创建了一个规则为`从开头到结尾必须是任意个a`正则表达式。

注意上面的例子中实际上是把一个`正则表达式字面量`作为参数传递给了 RegExp 构造函数。

`test` 方法可以用来校验某个字符串能否使用这个正则表达式匹配到目标。

如果不想传递字面量，也可以传递一个正则表达式字符串，这个时候不需要使用 `/` 包裹，而字面量是需要 `/` 包裹的。

```javascript
var regexp = new RegExp('^a*$'); // 不需要使用 / 将表达式包裹起来

regexp.test('aaaa'); // true
regexp.test('a'); // true
regexp.test('a1'); // false
```

这样创建出来的和传递正则表达式字面量的效果一样。

传递字符串的时候还能传递第二个参数，作为正则表达式的符号，部分文献也称其为描述符。

```javascript
var regexp1 = new RegExp('^a*$', 'i');
var regexp2 = new RegExp('^a*$');

var str = 'AAAAA';

console.log(regexp1.test(str)); // 输出：true
console.log(regexp2.test(str)); // 输出：false
```

符号 `i` 表示忽略大小写，所以 `regexp2` 无法在 `str` 中匹配到值。

如果需要多个符号，则将多个符号放在一起作为字符串即可。

```javascript
var regexp = new RegExp('^a*$', 'igm');
```

## 2. 字面量

通常构造函数会在不确定表达式的内容情况下使用，预定义好的正则表达式通常都会用字面量来表示。

正则表达式的字面量使用一对 `/` 进行包裹。

```javascript
var regexp = /^a&/;
```

这里不需要引号进行包裹，使用引号就变成了字符串。

如果需要加入符号，则跟在末尾即可。

```javascript
var regexp1 = /^a*$/g;
var regexp2 = /^a*$/ig;
```

## 3. 符号

在 ES6 之前，有三种符号。

* `g` 全局匹配，找到所有匹配，而不是在第一个匹配后停止

```javascript
var regexp1 = /a/g;
var regexp2 = /a/;

var str = 'abcdeabcde';

console.log(str.match(regexp1)); // 匹配到两个 a
console.log(str.match(regexp2)); // 只匹配到一个，并返回相应信息
```

可以看到，`regexp1` 能匹配到两个 a。

* `i` 忽略大小写

```javascript
var regexp1 = /apple/i;
var regexp2 = /apple/;

var str = 'AN APPLE A DAY KEEPS THE DOCTOR AWAY.';

console.log(str.match(regexp1)); // 可以找到一个
console.log(str.match(regexp2)); // 找不到 输出：null
```

regexp2 没有忽略大小写，所以是无法匹配到 `apple` 的。

* `m` 多行匹配

多行匹配模式下，开头和末尾就不是整个字符串的开头和末尾了，而是一行的开头和末尾。

> 目前 ES6 提供了三种新的描述符，分别为 u（Unicode 模式），y（粘连模式），s（dotAll 模式）。

## 4. 常用的正则表达式汇总

> 正则表达式不一定通用，可能需要结合业务的实际场景来做调整。

### 4.1 URL

```javascript
/(http[s]?:\/\/)?[^\s(["<,>]*\.[^\s[",><]*/
```

### 4.2 纯数字

```javascript
/^[0-9]*$/
```

### 4.3 邮箱

```javascript
/\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/
```

### 4.4 中文与全角符号

```javascript
/[\u3000-\u301e\ufe10-\ufe19\ufe30-\ufe44\ufe50-\ufe6b\uff01-\uffee]/
```

### 4.5 身份证（不验证是否合法）

```javascript
/\d{15}(\d\d[0-9xX])?/
```

### 4.6 仅包含英文字母的字符串

```javascript
/^[A-Za-z]+$/
```

### 4.7 正整数

```javascript
/^\d+$/
```

### 4.8 负整数

```javascript
/^((-\d+)|(0+))$/
```

### 4.9 数字（正负数、小数）

```javascript
/^(\-|\+)?\d+(\.\d+)?$/
```

### 4.10 IPv4

```javascript
/^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$/
```

## 5. 小结

正则表达式字面量需要使用 `/` 包裹，通常字面量会用于写死固定的正则表达式，如果需要动态生成，都会使用构造函数的方式。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
