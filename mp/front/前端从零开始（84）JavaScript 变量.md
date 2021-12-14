# JavaScript 变量

> 变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量就是`存放一些内容`的`容器`。

对于初学者，理解变量是重要的一环。

从分析`变量`这个名词，可以知道他是一个`可以改变的量`，这里的量就是代表某一种`值`。

在 JavaScript 中，变量就是一个用来存放值的容器，并且可以对容器中的值做修改。

每个变量都有唯一的变量名，使用变量名来区分变量。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7ae7af09c9e23014400452.jpg)

## 1. 声明变量

在 JavaScript 中使用`var`关键字来声明变量。

```javascript
var 存放数字用的变量 = 996;

console.log(存放数字用的变量); // 输出：996
```

上述这段代码就是`申明`了一个名为`存放数字用的变量`的变量，并且将它的值设为`996`。

使用 `console.log`，括号内放置`变量名`，即可将变量的值输出在控制台。

其中 `//` 后面的内容为`注释`，代码执行过程中会被忽略。

虽然使用中文作为变量名在 `chrome` 浏览器下没有报错，但是还是不建议使用。

常规场景中**不会有使用中文名作为变量的情况**。

所以上述例子中的变量名不可取。

```javascript
var number = 996;

console.log(number); // 输出：996
```

将`存放数字用的变量`修改成 `number` ，执行结果是一样的。

## 2. 赋值

给变量设置值的操作称为`赋值`操作。

### 2.1 最简单的赋值操作

```javascript
var result = 0;

console.log(result); // 输出：0
```

这是一个最简单的赋值操作，直接将值赋给变量。

通常只有`一个等号`出现的情况下就存在赋值操作。

### 2.2 将计算结果赋值给变量

```javascript
var result = 2 + 3;

console.log(result); // 输出：5
```

这也是一个赋值操作，只不过等号右边的 `2 + 3` 会被计算出结果（计算的方式和小学开始学习的自然数学一样），再赋给变量 `result`。

将上面这个例子做一个简单的改写：

### 2.3 让变量也参与计算

```javascript
var number1 = 2;
var number2 = 3;

var result = number1 + number2; // 2 + 3

console.log(result); // 输出：5
```

原本 `2 + 3` 这部分也可以被变量所代替，参与计算的就是变量中的值。

### 2.4 改变变量的值

```javascript
var string = '今天加班？';

console.log(string); // 输出：今天加班？

string = '福报！';

console.log(string); // 输出：福报！
```

> 注意：
>
> 这里赋给变量的值和之前有点不一样，是中文文字。
>
> 当需要用变量存放一些“字”的时候，就需要用单引号`'`或者双引号`"`将需要存放的字包裹。
>
> 通常单个字会称之为`字符`，多个字的时候称为`字符串`。
>
> 这里做个了解，具体的会在后续`数据类型`章节详细展开讨论。

这段代码运行后可以在控制台观察到有两个输出，分别对应变量的值。

代码很简单，先声明了一个叫 `string` 的变量，并赋值字符串`今天加班？`并输出，随后修改了他的值，重新赋值了字符串`福报！`。

这是变量最重要的一个特性：`可变`。

## 3. 变量的命名规范

在 JavaScript 中变量名存在一定规范，所有变量名必须符合这些规范，否则程序无法执行。

### 3.1 变量名必须使用`字母`、`下划线(_)`、`美元符号($)`开头

尽管之前的例子有用到中文作为变量名，但是是不推荐的。

```javascript
// 不会报错但是不推荐
var 数字 = 1;
// 错误
var 1number = 1;
// 错误
var number@a = 1;
// 错误
var num+aa = 2;

//下面是正确的方式
var number1 = 1;
var _number = 1;
var $number = 1;
```

以上是一些简单的示例，可以根据规则自己在控制台尝试寻找规则。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7a42720a59717f10810372.jpg)

### 3.2 变量对大小写敏感

```javascript
// 这是两个不同的变量
var firstName = 'Hello';
var firstname = 'hello';
```

以上是两个不同的变量，在 `JavaScript` 中变量是对大小写敏感的。

两个变量名即便字母是相同的，但是大小写不同，就不能算做一个变量。

### 3.3 无法使用关键字作为变量名

`关键字`就是指一些已经被 `JavaScript` 预定义或者保留下来的内容，如声明变量用的关键字 `var` 就不能作为变量名。

```javascript
var var = 1; // Uncaught SyntaxError: Unexpected token 'var'
```

上面这段代码尝试着将 var 作为变量，到控制台运行是会报错的。

## 4. 合理规范的变量名

> 刚开始学习的读者，现在去深究如何命名一个变量还有些尚早，因为结合了具体的需求场景才能体会到一个好的变量名的重要性。可以先在此做个了解。

对于变量名，除了上面提到的变量命名的规范，最需要注意的就是给变量起一个`有意义`的名字。

如求和：

```javascript
var num1 = 1;
var num2 = 2;
var num3 = 3;
var num4 = 4;

var count = num1 + num2 + num3 + num4;
```

其中`num`是`number`的缩写，是很常用的一种缩写。

`count`则是总数，表示求和的结果。

如果将上述例子做如下修改：

```javascript
var a = 1;
var b = 2;
var c = 3;
var d = 4;

var e = a + b + c + d;
```

缺少了有意义的变量名就比较难看出代码具体在做什么。当然这段代码本身意义就不大，场景太过简单。

刚才提到的缩写，其实也是要注意的一点，缩写上一定要使用通用的缩写，如含有`fn`表示一个功能或者函数，`avg` 表示平均值，`pwd` 表示密码，`i18n` 为国际化。

这些缩写比较通用，大部分开发者都可以看得懂。随着编码经验的增加，会在他人代码里见到大量的缩写，从而累积到自己的大脑的缩写库中。

最后需要注意的一点是业务中尽量不要含有中文拼音或中文拼音的缩写，排开鄙视链的原因，最大的问题是会让变量名变得冗长难懂。

以上内容在写 demo 或者测试功能的时候可以不需要考虑，写 demo 等大部分情况是为了验证自己的猜想。

```javascript
// 不合理的变量名
var ln = 'World'; // last name
var zs = 0; // 总数
var jinNianDeNianShouRu = 1999999999; // 今年的年收入
```

以上是针对`变量名的意义`展开的讨论。

还有需要注意的是变量命名的格式，大部分前端程序员会使用`驼峰命名法`，也就是第一个字母小写，后续如果有新的单词来进行构成，单词的第一个字符都大写。

如：

```javascript
var firstName = 'Hello';

var lastName = 'world';

var createAt = 1577895179196;

var userInfo = '用户信息'; // Info => Information

var isPaidUser = '是否付费用户';
```

可以见到上面的变量，从构成变量名的第二个单词开始，首字母都是大写，这就是驼峰命名的格式，本 Wiki 所有变量名使用的就是这种格式。

当然还有大驼峰，就是第一个字母也大写。

除此之外最常用的还有使用下划线分隔变量，如 `user_info`，还有按功能来划分的变量名，如使用`匈牙利命名`法，这里不再做展开。

## 5. 有关变量的其他知识

### 5.1 变量的默认值

变量在声明的时候，如果没有赋值，则变量就会有一个默认值 `undefined`。

```javascript
var total;

console.log(total); // 输出：undefined
```

`undefined` 是一种是数据类型，具体内容可以参考`数据类型`章节。

### 5.2 同时声明多个变量

使用一个 `var` 关键字就可以直接声明多个变量。

```javascript
var num1 = 0, num2 = 1;

// 通常会换行，方便阅读代码
var num3 = 2,
    num4 = 3,
    num5 = 4,
    num6,
    num7 = 6;
```

在一个变量声明后，使用逗号分隔，紧接着声明下一个变量即可。

通常使用一个 `var` 声明多个变量的时候也会换行，方便后续阅读，并保持代码格式上的整洁清晰，防止一行过长。

### 5.3 变量在 window 上

在最外层声明的变量（不包括 ES6 模块的情况），实际上是变成了 window 对象的一个属性。

```javascript
var number = 996;

console.log(number); // 输出：996
console.log(window.number); // 输出：996
```

上述代码执行后输出的两个内容是一样的，都为 996。有关 window 对象的内容可以参考 `BOM` 章节。

细心的读者应该会注意到`最外层`这个条件，因为变量还有可能声明在函数之中，函数有自己独立的作用域，通常在函数中使用 `var` 关键字声明的变量，只在函数中有效。

至于为什么可以省略 `window` 直接访问到变量，可以参考`作用域链`章节。

### 5.4 不使用 var 关键字声明的变量

假如不使用 `var` 关键字，直接创建变量并赋值：

```javascript
total = 10;

console.log(total); // 输出：10
```

在控制台运行后会发现其实并没有报错，输出的结果也正常。

在非`ES6模块`中，这样创建的变量和使用 `var` 创建的变量除了不能提前使用之外，没有其他大的区别，会被直接作为 window 对象的属性，成为全局变量。

即便是在函数或者其他存在块级作用域的环境中，这样声明的变量也会作为全局变量。

### 5.5 连续赋值

```javascript
var a = b = 1;
```

假如把上面这行代码拆开来可以理解成是这样的：

```javascript
b = 1;
var a = b;
```

看似没什么问题，许多开发者也会用这种方式同时声明多个变量，但如果在函数或者独立的作用域中，`b` 就会成为全局变量，造成全局命名空间的污染。

### 5.6 重复声明变量

按照之前说的，变量在声明的时候如果没有赋值，则会是 `undefined`，这个规则在重复声明的情况下不适用。

```javascript
var num = 1;
var num;

console.log(num); // 输出：1
```

观察上面这个例子输出的结果，可以发现变量 `num` 的值并没有改变。

但是如果重新声明的同时做赋值操作，值就会改变。

```javascript
var num = 1;
var num = 3;

console.log(num); // 输出：3
```

这个例子输出的结果，就是再次声明并赋值后的值。

### 5.7 提前使用变量

```javascript
console.log(number); // 输出：undefined

var number = 1;
```

这个例子先输出了 `number` 的值，再声明并对其进行赋值。

代码并没有报错，但如果没有第二行声明，只输出 `number`：

```javascript
console.log(number); // Uncaught ReferenceError: number is not defined
```

这样子会爆出`变量未定义`的错误，说明变量是可以被提前使用，只是没有值，或者说是 `undefined` 默认值。

具体原因可以参考`执行上下文`章节。

这里简单的解释可以理解成，在浏览器执行的时候，会把代码调整成如下样子：

```javascript
var number;

console.log(number); // 这个时候 number 还没有被赋值，所以输出 undefined

number = 1;
```

### 5.8 常量

常量就是定义并赋值后再也不能修改的量，通常一些不会改变的量，如配置、物理值等会声明为常量，在 ES6 之前是没有提供常量这一特性的。

但是根据常量自身的特性，`定义赋值后不能被修改`，就可以通过一些方式来模拟常量。

第一种就是采用约定的形式，通常常量都是大写，不同单词之间用下划线分隔。

```javascript
var PI = 3.1415926535;

var DB_ACCOUNT = 'root';
var DB_PASSWORD = 'root';
```

这种方式定义的`常量`本质上还是变量，值还是可以修改的，但因为命名格式采用`国际惯例`，一眼就能看出是常量，不会对其修改。

这种方式是最简单的方式，但不安全。

第二种方式就是利用对象下属性的描述来控制可写性，将对象的属性设置为只读。

```javascript
var CONFIG = {};

Object.defineProperty(CONFIG, 'DB_ACCOUNT', {
  value: 'root',
  writable: false,
});

console.log(CONFIG.DB_ACCOUNT); // 输出：root

CONFIG.DB_ACCOUNT = 'guest';

console.log(CONFIG.DB_ACCOUNT); // 因为不可被改写，所以输出：root
```

这种方式将常量都放在一个对象下，通过`Object.defineProperty`定义属性，设定其`writable`为`false`，就可以防止被改写。

但有一个问题，`CONFIG`自身这个对象可能被修改。

换一个思路，既然在最外层声明的变量是放在`window`上的，那可以用这个方式往 window 上挂不可改写的属性。

```javascript
Object.defineProperty(window, 'DB_ACCOUNT', {
  value: 'root',
  writable: false,
});

console.log(DB_ACCOUNT); // 输出：root

DB_ACCOUNT = 'guest';

console.log(DB_ACCOUNT); // 因为不可被改写，所以输出：root
```

通常情况下 window 对象是不可被修改的，这样常量的安全系数就变得非常高，但缺点是可能性较差，通过一点修改可以提升可读性。

```javascript
var define = function(name, value) {
  Object.defineProperty(window, name, {
    value: value,
    writable: false,
  });
};

define('DB_ACCOUNT', 'root');
define('DB_PASSWORD', 'root');
```

只要约定好使用 `define` 函数定义的都为常量即可。

还有两种方式，就是结合`Object.seal`与`Object.freeze`的特性来声明常量。

前者可以使对象不能再被扩充，但是所有属性还需要再手动设置不可写，后者可以让对象不能扩充，属性也不能修改。

这里对这两个原生方法不再做过多描述，有兴趣可以查阅相关 API 资料。

## 6. 小结

变量就是存放值的容器。

变量名存在一些命名规则：

* 变量名必须使用`字母`、`下划线(_)`、`美元符号($)`开头；
* 变量对大小写敏感；
* 无法使用关键字作为变量名。

同时起变量名的时候需要有意义，靠近上下文场景。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
