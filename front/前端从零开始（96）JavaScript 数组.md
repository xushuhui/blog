# JavaScript 数组

> 数组是一种使用整数作为键 (integer-key-ed) 属性和长度 (length) 属性之间关联的常规对象。此外，数组对象还继承了 Array.prototype 的一些操作数组的便捷方法。——MDN

数组可以存放多个值。数组的`数`不是指数字，而是指`数据`，数组可以存放任意数据类型的值。

从理论上讲，在内存允许的情况下，数组的长度是无限的。

![图片描述](https://img.mukewang.com/wiki/5e7ae0cd095da6a214400768.jpg)

## 1. 创建数组

数组由中括号 `[]` 包裹，每一项之间用逗号 `,` 分隔。

```javascript
[第一项, 第二项, 第三项, ... , 第n项]
```

创建数组最常见的方式是使用字面量。

```javascript
var hobby = ['java', 'c', 'c++', 'python', 'javascript', 1, 2, 3, true, false];

console.log(hobby);
```

通过直接描述一个数组的方式就可以创建一个数组。

## 2. 数组长度

数组的长度可以通过 `length` 属性获取。

```javascript
var arr = [1, 2, 3, 4, 5];

console.log(arr.length); // 5
```

## 3. 访问数组成员

数组的每一项可以称之为`成员`。

数组成员可以通过`下标`访问，下标也可以称为`索引`。

下标可以理解成数组成员的编号，编号从 0 开始，到数组长度 -1 。

```javascript
var arr = ['第一项', '第二项', '第三项'];

var item1 = arr[0]; // 从0开始编号 第一项的下标是0
var item2 = arr[1];
var item3 = arr[2];

console.log(item1); // 输出：第一项
console.log(item2); // 输出：第二项
console.log(item3); // 输出：第三项
```

通过`数组[下标]` 的方式即可访问到成员。

## 4. 遍历数组

数组遍历主要有两种方式：

* for 循环
*  forEach 方法

### 4.1 for 循环

使用 for 循环，从数组下标 0 循环到最后一位，过程中通过下标访问成员。

```javascript
var arr = ['第一项', '第二项', '第三项', '第四项', '第五项'];

var i, len;
for (i = 0, len = arr.length - 1; i <= len; i++) {
  console.log(arr[i]);
}
```

### 4.2 forEach 方法

`forEach` 是数组原型上的方法，所有数组都具有此方法。

forEach 接收一个函数作为参数，在遍历每一项的时候，这个函数会被调用，同时将当前遍历到的项、当前项的下标（索引）、遍历的数组作为函数参数传递过来。

```javascript
var arr = ['第一项', '第二项', '第三项', '第四项', '第五项'];

arr.forEach(function(item, index, arr) {
  console.log('第' + (index + 1) + '项的值是：' + item);
});
```

第二个参数的值也是从 0 开始的。

通常第三个参数用到的比较少，没有用到可以没必要接收这个参数。

## 5. 修改数组成员的值

成员的值可以通过`数组[下标] = 值`的方式进行修改。

```javascript
var arr = ['一', '三', '三', '四'];

console.log(arr[1]); // 输出："三"

arr[1] = '二';

console.log(arr[1]); // 输出："二"
```

## 6. 增减数组项

数组提供了很多方式来对成员进行增减操作，也可以依靠其他特性来增加减少数组成员。

### 6.1 增加

#### 6.1.1 直接给指定位置赋值

通过下标，直接访问到一个不存在的成员，然后进行赋值，就可以为数组增加一项。

```javascript
var arr = ['jquery', 'react'];

arr[1] = 'vue';
arr[2] = 'angular';

console.log(arr[2]); // 输出："angular"
console.log(arr.length); // 输出：3
```

#### 6.1.2 push 方法

push 方法接收任意个参数，这些参数会依次添加到数组的末尾，添加完后返回数组新的长度。

```javascript
var arr = [1];

var length = arr.push(2, 3);

console.log(arr); // 输出：[1, 2, 3]
console.log(length); // 输出：3
```

通常不会用到这个返回的长度，可以不需要接收返回值。

#### 6.1.3 unshift 方法

unshift 接收任意个参数，这些参数会被添加到数组头部，添加完后返回数组新的长度。

```javascript
var arr = [3];

var length = arr.unshift(1, 2);

console.log(arr); // 输出：[1, 2, 3]
console.log(length); // 输出：3
```

### 6.2 删除

#### 6.2.1 pop 方法

pop 方法会删除数组最后一项，并将删除项作为返回值。

```javascript
var arr = ['c++', 'java', 'javascript'];

var lastOne = arr.pop();

console.log(lastOne); // 输出："javascript"
```

如果数组是空的，调用 pop 会返回 `undefined` 。

#### 6.2.2 shift 方法

shift 方法会删除数组的第一项，并将删除项作为返回值。

```javascript
var arr = ['996', '007'];

const first = arr.shift();

console.log(first); // 输出："996"
```

与 pop 一样，如果是数组为空的情况下，会返回 `undefined` 。

### 6.3 在数组中间删除或添加值

splice 方法可以在任意位置添加或删除值。

这个方法接受任意个参数，前两个为从哪里开始（从 0 开始计数），删除几个，从第三个开始则是要添加的项。

```javascript
arr.splice(从第几个开始, 要删除几个, 要添加的项目1, 要添加的项目2, ..., 要添加的项目n);
```

如果不需要删除，只需要往数组中间插入值，只需要传递 0 给第二个参数即可。

```javascript
// 在第二项之后插入3, 4, 5
var arr = [1, 2, 6];

arr.splice(2, 0, 3, 4, 5);
```

因为第一个参数是从 0 开始计数，所以在第二项之后，就是要插入在第三项的位置，所以传递 2 ，不需要删除项目，所以第二个参数传递 0 ，之后就是要插入的项。

> 注意：往数组中间插入数据的情况相对较少，数组做这种操作是比较低效的，小量的操作对性能的影响可以忽略不计，但有超大量非首尾的增减操作，应考虑使用`链表`。

删除项只需要指定从第几项开始删除，要删除几项即可。

```javascript
// 去除 '996'、'加班'
var arr = ['早睡早起', '朝九晚五', '996', '加班'];

arr.splice(2, 2);
```

'996’和’加班’是连续的，所以一个 splice 就可以删除掉这 2 项，2 个参数的意思就是从第 2 项开始，删除 2 项。

## 7. 清空数组

将数组所有成员全部删除就达到了清空的效果。

```javascript
var arr = [1, 2, 3, 4];

arr.splice(0, arr.length);
```

当然也可以使用 `pop` 一个个删除，但是通常都不会用这种方式。

清空数组最常用的方式是重新给变量赋值。

```javascript
var arr = ['red', 'green', 'blur'];

arr = [];

console.log(arr); // 输出空数组：[]
```

通过给变量赋值一个新的空数组，达到清空数组的目的，但是这样会改变引用，新赋值的数组和之前的数组并不是同一个。

另一种方式可以让保持对当前数组的引用。

```javascript
var arr = ['yellow', 'black'];

arr.length = 0;
```

通过给数组的 `length` 属性重新赋值，也可以达到清空数组的效果。

这种方式相对灵活，假如需求是保留三项、五项，只需要给 length 赋值对应的值即可。

## 8. 使用 Array 创建数组

内建对象 `Array` 也可以用来创建数组。

```javascript
var arr = new Array();
```

如果什么参数都不传递，则返回一个空数组。

传参则有 2 种情况：

* 如果只传一个参数，并且这个参数的类型为数字，则会创建长度为这个数字的数组；
* 传入其他类型的一个或者多个参数，则会将这些参数组合成数组。

```javascript
var arr = new Array(10);

console.log(arr); // 输出：[empty × 10]
console.log(arr.length); // 输出：10
```

在控制台可以观察到这个数组长度为 10，但均为 `empty` 。

如果尝试着访问其中某一项，得到的值是 `undefined` 。

```javascript
var arr1 = new Array('item1');

var arr2 = new Array(1, 2, 'item3');

console.log(arr1); // 输出：["item1"]
console.log(arr2); // 输出：[1, 2, "item3"]
```

这样创建的数组，成员与传参一致。

## 9. 数组中的 undefined 与 empty

在数组中 undefined 与 empty 是有区别的。

使数组项变成 `empty` 通常有两种方式。

* 使用 Array 对象同时提供了长度创建出来的数组
* 使用 `delete` 关键字删除的数组项

```javascript
var arr1 = new Array(10);

arr1[0] = 1;

var arr2 = [1, 2, 3, 4, 5];

delete arr2[3];

console.log(arr1);
console.log(arr2);
```

![图片描述](https://img.mukewang.com/wiki/5e7adc990a1fe2d413760566.jpg)

虽然 empty 的项在访问的时候返回的是 undefined ，但其本身只做简单占位，\ 是遍历不到的。

```javascript
var arr = [1, undefined, 3, 4, 5];

delete arr[3];

arr.forEach(function(item, index) {
  console.log(index, item);
});

var i = 0;
for (i in arr) {
  console.log(i, arr[i]);
}
```

上面两种遍历的方式都是遍历不到 `empty` 项的，而 `undefined` 是可以遍历到的，这是最主要的区别。

## 10. 使用数组

数组非常常用，大量的 HTML 字符串在做拼接的时候，就会使用到数组。

除了用于简单的存储数据，数组也可以被用来解决问题。

### 10.1 判断括号是否匹配

判断一个数学式子的括号匹配是否合法，如 `(1 + 2) * (3 + 4))` ，这个式子就是不合法的。

校验括号合法不单单要看左括号和右括号的数量是否相等，还要看括号的顺序， `))((` 这样的括号顺序一定是错误的。

利用 JavaScript 数组的特性，可以很容易的实现。

```javascript
// 空数组
var stack = [];

// 一个式子
var equation = '(1 + (2 - 3) * 4) / (1 - 3)';

var i, len;
for (i = 0, len = equation.length; i < len; i++) {
  if (equation[i] === '(') { // 如果碰到左括号
    // 往数组里放个1
    stack.push(1);
  } else if (equation[i] === ')') { // 如果碰到右括号
    if (!stack.length) { // 判断数组长度，如果是0，则肯定是出错的，数组长度0的时候说明没有左括号，没有左括号是不可能出现右括号的
      // 随便放一个1
      stack.push(1);
      break;
    }
    // 如果数组不是空的 就从数组末尾拿一个走。
    stack.pop();
  }
}

// 判断数组长度
if (!stack.length) { // 如果数组已经空了，说明括号都一一对应，式子里的括号没问题。
  console.log('括号合法');
} else {
  console.log('括号不合法');
}
```

使用数组实现的具体思路就是，碰到左括号就往数组里放一个成员，碰到一个右括号就拿掉一个成员。

这样如果最后有剩下，说明括号没有一一成对。

`(1+2*(3+4))*1` 如这样一个式子：

```javascript
初始化操作：
  定义数组 arr为空

从式子第一个字符开始循环
  第一次循环：
    数组的值为 []
    得到字符"("
    判断是左括号，往数组里放一个1，表示碰到了左括号

  第二次循环
    数组的值为 [1]
    得到字符"+"
    既不是左括号，又不是右括号，进行下一轮循环，不做操作

  第三次循环
  第四次循环
  第五次循环
    与第二次循环基本一致

  第六次循环
    数组的值为 [1]
    得到字符"("
    是左括号 往数组里再放一个1，表示碰到了左括号

  第七次循环
    数组值为 [1, 1]
    得到字符"3"
    不是左括号，也不是右括号，进行下一轮循环

  第八次循环
  第九次循环
    与第七次一致

  第十次循环
    数组的值为 [1, 1]
    得到字符")"
    是右括号，从数组末尾拿掉一项

  第十一次循环
    数组的值为 [1]
    得到字符")"
    是右括号，从数组末尾拿掉一项

  第十二次循环
  第十三次循环
    数组值为 []
    都不是括号，不做操作
循环结束

判断数组值，如果是空的，说明括号匹配完了，显然 (1+2*(3+4))*1 是合法的
```

当然这种判断有局限性，假如碰到 `((1+)2)-3` 这种非括号造成不合法的式子，就判断不出来了。

> 其实这里用到了`栈`这种数据结构，这个问题在栈的应用上很经典，是算法入门常见面试题之一。

## 11. 类数组

类数组并不是数组，而是长得像数组的对象。

```javascript
var fakeArray = {
  0: '第一项',
  1: '第二项',
  3: '第三项',
  length: 3,
};

console.log(fakeArray[0]); // 输出："第一项"
console.log(fakeArray.length); // 输出：3
```

上述例子中的 `fakeArray` 就是一个类数组，属性是以类型数组的`下标`的形式存在，同时也具有 `length` 属性。

这种类数组对象，也被称为 `array-like对象` ，部分文献也称为`伪数组`。

类数组对象可以转化为数组，许多方法在设计时也会考虑支持类数组。

## 12. 小结

JavaScript 中的数组非常灵活，可以存放任意类型、任意长度 （内存足够的情况下） 的数据，其下标从 0 开始，最大到数组长度减去 1 ，并提供了一系列方法，来完成增、删、改、查操作。

数组项的 `empty` 和 `undefined` 的区别，是面试中常问的问题。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
