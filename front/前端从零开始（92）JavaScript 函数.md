# JavaScript 函数

> 在 JavaScript 中，函数是头等 (first-class) 对象，因为它们可以像任何其他对象一样具有属性和方法。它们与其他对象的区别在于函数可以被调用。简而言之，它们是 Function 对象。(MDN)

函数就是一段代码片段，调用函数就是执行函数中的代码。

## 1. 函数的使用

### 1.1 语法

函数使用前通常与变量一样需要先进行声明，用 `function` 关键字定义函数。

```javascript
// 常见的函数的定义方式
function 函数名(参数1, 参数2, ...) {
  代码片段;

  return 返回值;
}

// 调用函数 (执行函数中的代码)
var 函数的返回值 = 函数名(参数1, 参数2, ...);
```

* 调用函数就是执行函数中的代码
* 参数是调用函数的时候传递过去的，在函数执行过程中可以访问到
* 函数执行完毕后可以有一个返回值，调用函数的地方可以接收到这个返回值

### 1.2 调用函数

> 使用 `函数名()` 的方式即可调用一个函数

以下是一个最简单的函数：

```javascript
function say() {
  console.log('hello');
}

say(); // 输出："hello"
```

调用这个函数就会在控制台输出 `hello` 字符串。

**这个函数没有返回值，默认会返回一个 `undefined`**。

### 1.3 带有参数与返回值的函数

> 在声明函数的时候，可以对参数也做上说明

假设有一个需求，**需要一个计算三角形周长的函数**。

计算三角形周长则需要知道三角形三条边各自的长度，然后将他们求和。

定义函数的时候就可以将三条边作为参数进行声明。

```javascript
function calcPerimeter(a, b, c) {
  // a, b, c 分别代表三条边
  var sum = a + b + c;

  return sum;
}

// 调用函数 并将返回值赋值给perimeter
var perimeter = calcPerimeter(3, 4, 5);
```

在调用函数的时可以传递值过去，这些值可以在函数中被访问。

在以上 `calcPerimeter` 函数被调用的时，传递了 `3, 4, 5` 三个值。

三个值对应到函数声明时定义的三个参数 `a, b, c`。

所以函数执行过程中 `sum` 的值为 `3 + 4 + 5`，即 `12`，随后 `sum` 被作为返回值进行返回。

最终变量 `perimeter` 也会被赋值为 12。

## 2. 怎么运用函数

### 2.1 合理包装内容

函数可以对代码进行封装，让逻辑更加清晰。

比如如下代码块：

```javascript
// 改写前
var num = 10;

var flag = false;

var i;
var len;
for (i = 2, len = num - 1; i <= len; i++) {
    if (num % i === 0) {
        flag = true;
        break;
    }
}

console.log(flag);
```

以上代码第一眼可能无法看出具体在做什么，仅需要做一点修改，就能有所改善。

```javascript
// 改写后
function isPrimeNumber(num) {
  var flag = false;

  var i;
  var len;
  for (i = 2, len = num - 1; i <= len; i++) {
    if (num % i === 0) {
      flag = true;
      break;
    }
  }

  return flag;
}


var num = 10;

var result = isPrimeNumber(num);

console.log(result);
```

改写后的代码似乎多了几行，但是将其中核心部分包装成了函数。

通过 `isPrimeNumber` 函数名可以很容易的了解到这一段代码作用是用来判断一个数是否为`质数`。

当然有个前提就是起一个 **可以让大部分人看得懂** 的函数名。

### 2.2 优秀的函数名

优秀的函数名可以帮助他人更容易理解代码，同时当自己一段时间后再回头看代码时，能更容易进入当时写代码时候的思维模式等。

这里提供几个函数命名的建议，具体的命名可以根据团队规范、个人成长等做调整。

#### 2.2.1 拼写准确

准确的拼写十分重要，绝大多数情况下函数名都会是英文单词组成的。

当然许多时候手一快可能就少了一个字母，或者错将 `wrap` 进行乾坤大挪移拼写成了 `warp`。

许多情况是无法避免的，经常需要自检。

当然可以借助一些单词的检查插件，如 `Visual Studio Code` 可以借助 `Code Spell Checker` 插件来检查单词的正确性。

再者碰到想起的函数名但是单词拼写不出来，尽可能翻词典，日积月累能有大量的词汇沉淀。

#### 2.2.2 尽量不使用拼音或者混用拼写

尽量不要使用拼音或者是首字母缩写。

以下函数名或许会造成困扰：

```javascript
function jslsh() {}

function jsNumber() {}
```

以上是`计算两数和`函数的命名，可能只有天和地知道这个是什么意思。

当然，如果是自己写 demo 或者测试代码的时候，其实不需要考虑这么多。

#### 2.2.3 有“状态”的函数名

如碰到函数功能是判断`是否`、`有没有`、`可以`的时候，可以带上一些前缀，比如：

```javascript
// 是否登入
function isLogin() {}
```

同时可以合理的使用动词，比如`打开文件`就可以使用 `openFile` 函数名，具体的状态可以根据语境、函数作用、个人习惯等做调整使用。

#### 2.2.4 合理使用缩写

使用词语的缩写尽量使用通用的缩写

如：

* pwd - password
* mgr - manager
* del - delete
* …

这些缩写大部分开发者是可以看的懂的缩写。

## 3. 函数示例

### 3.1 计算圆的面积

分析：根据圆面积公式 S=π·r·r，其中 S 就是要求的值，即函数的返回值，π 是常量（固定的一个值），半径 r 是未知数，所以 r 就可以设计成参数

```javascript
function circleArea(r) {
    var pi = 3.1415926;

    return pi * r * r;
}

// 计算半径为10的圆的面积
var area = circleArea(10);
```

### 3.2 判断某个 DOM 元素是否含有某个类名

分析：

`某个DOM`和`某个类名`可以说明有两个未知量，可以设计成两个参数。

根据描述也可以确定一个 `某个DOM` 的类型是个 `DOM` 对象，`某个类名`是个字符串

只要拿到这个 DOM 的 `class` 属性，判断里面是不是含有这个类型即可得到结果

```javascript
function hasClass(el, className) {
  // el 是 element的缩写，表示一个dom元素

  // 如果没有元素 则返回
  if (!el) {
      return false;
  }

  // 根据空格分割成数组
  // 可以不使用 split 方法，使用字符串也可以用indexOf匹配
  var classList = el.className.split(' ');

  // 判断是否存在
  if (classList.indexOf(className) >= 0) {
      return true;
  }

  return false;
}
```

## 4. 函数的其他知识

> 以下扩展内容可能需要一定的知识积累，遇到不懂的地方可以停下脚步，先学习下一章节

### 4.1 函数表达式

以上篇幅的函数其实都通过`函数声明`的方式来定义，还有一种方式就是使用函数表达式定义函数。

```javascript
// 函数声明
function add(a, b) {
    return a + b;
}

// 函数表达式
var add = function(a, b) {
    return a + b;
};
```

通过上述例子可以看出写法上的区别就是`函数表达式`是将函数赋值给了变量。

这两种方式创建的函数最大的区别在于，**不能提前调用使用函数表达式创建的函数**

光看句子有点抽象，举个例子？：

```javascript
var num1 = add1(1, 2);

var num2 = add2(3, 4);

// 函数声明
function add1(a, b) {
    return a + b;
}

// 函数表达式
var add2 = function(a, b) {
    return a + b;
};
```

上面一段代码在执行的时候会报 `add2 is not a function` 的错误，表示 `add2` 不是函数，也就是说 `add2` 不能被提前使用，而 `add1` 可以。

具体原因可以查看`执行上下文`章节。

### 4.2 函数作用域

函数有他自己的作用域，函数内声明的变量等`通常情况下`不能被外部访问，但是函数可以访问到外部的变量或者其他函数等

```javascript
var a = 1;

function fn() {
    var b = 2;

    console.log(a); // 输出：1
    console.log(b); // 输出：2
}

fn();

console.log(b); // ReferenceError: b is not defined
```

执行以上代码会报 `b is not defined` 错误。

### 4.3 匿名函数

没有名字的函数就是一个匿名函数

```javascript
var fn = function() {
    console.log('我是一个匿名函数');
};
```

除了在`函数表达式`中会出现匿名函数，还有许多场景。

相对常见的一个就是`自执行匿名函数`，MDN 官方翻译为`立即调用函数表达式`。

`自执行`就是这个函数声明后就会立即执行，自执行的匿名函数通常会被用来`形成独立的作用域`。

如：

```javascript
(function() {
    var num = 1;

    alert(num);
})();
```

这是一个自执行的匿名函数，这个匿名函数是被包裹了一段括号后才被调用的。

以下这段代码会报错：

```javascript
// 报错
function() {
    var num = 1;

    alert(num);
}();
```

浏览器会告诉你必须给函数一个名字。

通过括号包裹一段函数，让`js引擎`识别成他是一个函数表达式，再对他进行执行，就不会报错，这是加括号的原因。

同理，可以使用 `+`，`!` 等运算符代替括号，让一个匿名函数成为一个函数表达式即可。

大部分第三方框架都会通过一个自执行的匿名函数包裹代码，与浏览器全局环境隔离，避免污染到全局环境。

### 4.4 具有函数名的函数表达式

函数表达式进行声明的时候也可以使用具名函数

```javascript
var count = function fn(num) {
    console.log('我是一个函数');
};
```

以上这段代码是不会报错的，但是不能通过 `fn` 访问到函数，这里的 `fn` 只能在函数内部进行访问，通常在使用递归的形式做计算的时候会用到这种写法。

```javascript
var count = function fn(num) {
    if (num < 0) {
        return num;
    }

    return fn(num - 1) + num;
}

count(5);
```

上面这个例子，就是在函数内部访问 `fn` 调用自己，使用递归的形式求和。

**注：递归相关的知识可以参考相关文献进行学习**

### 4.5 arguments

> arguments 是一个对应于传递给函数的参数的类数组对象。(MDN)

通常情况下函数都具有 `arguments` 对象，可以在函数内部直接访问到。

他是一个类数组，即长得很像数组，成员都是用数字编号，同时具有 length 属性。

arguments 中存放着当前函数被调用时，传递过来的所有参数，即便不声明参数，也可以通过 arguments 取到传递过来的参数。

```javascript
function sum() {
    console.log(arguments);
}

sum(1, 2, 3, 4);
```

执行上述代码，可以看到在控制台输出了一个对象，存放的就是所有传递过去的参数，利用这一特性，就可以不限制参数个数，或者让函数做中转站（拦截函数），利用 arguments 将参数传递给另一个函数。

如，一个不确定用户输入的参数个数的求和函数：

```javascript
function sum() {
  var total = 0;

  var i;
  var len;
  for (i = 0, len = arguments.length; i < len; i++) {
    total += arguments[i];
  }

  return total;
}

var total = sum(1, 2, 3, 4, 15);
console.log(total); // 输出：25
```

通过循环遍历 `arguments` 对象，就可以得到所有参数，然后做累加就可以达到求和的目的。

### 4.6 函数和方法

方法在本质上是个函数。

通常都能听到“调用一下某个方法”，“取到某个方法的返回值”，这里的方法其实就是一个函数。

一般方法是用来描述对象的某个行为的，但是平时我们会混用，口头交流的时候会经常把函数直接称作方法。

只要自己理解，不需要去纠结函数和方法到底是什么，也不用特意纠正别人的说法，大家都能听得懂就行。

### 4.7 JS DOC 注释

使用 `JS DOC` 描述函数是非常良好的习惯，良好的 `JS DOC` 书写还可以使用工具快速生成文档。

`JS DOC` 对函数的描述大体如下：

```javascript
/**
 * 这是这个求幂函数 计算 x 的 y 次方
 * @param {Number} x - 底数
 * @param {String} y - 指数
 */
function pow(x, y) {
    // ...
}
```

除此之外还可以描述返回值等。

### 4.8 纯函数与副作用

所谓纯函数，就是没有副作用的函数

一个函数从执行开始到结束，没有对外部环境做任何操作，即对外部环境没有任何影响（没有副作用），这样的函数就是纯函数。

纯函数只负责输入输出，对于一种输入只有一种函数返回值。

如果函数中存在 `Math.random` 这种影响返回值的函数，也不能算是纯函数。

```javascript
// 纯函数
function add(a, b) {
  return a + b;
}

// 非纯函数
var person = { name: '小明' };
function changeName {
  person.name = '小红'; // 影响了函数外的内容，产生了副作用
}
```

### 4.9 构造函数

当一个函数与 `new` 关键字一起被调用的时候，就会作为一个构造函数。

```javascript
function Person(name, age) {
    this.name = name;
    this.age = age;
}

Person.prototype.say = function() {
    console.log('我是' + this.name);
};

var person = new Person('阿梅', 12);

person.say();

console.log(person);
```

可以看到当函数作为构造函数调用的时候，默认返回的是一个对象。

细心的读者仔细观察就能发现，构造函数的默认返回值是函数体内的 this。

事实上构造函数的执行有一定流程：

1. 创建一个空对象，将函数的 this 指向这个空对象
2. 执行函数
3. 如果函数没有指定返回值，则直接返回 this（一开始创建的空对象），否则返回指定返回值

理解这个流程，就能理解构造函数的返回值。

具体的函数的 `prototype` 属性等可以参阅`原型`章节。

## 5. 小结

函数特性相对较多，也是 JavaScript 的核心之一。

函数可以用于封装代码，提供代码的复用率和可读性，在大部分情况下，当两段代码具有超高相似度时，应当设计成函数，不同的部分使用参数进行区分。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
