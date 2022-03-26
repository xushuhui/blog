# 逗号操作符

> 逗号操作符 对它的每个操作数求值（从左到右），并返回最后一个操作数的值。(MDN)

逗号操作符会依此从左到右执行逗号分隔的表达式，并把最后一个表达式的运算结果作为最终结果。

```javascript
var 表达式3的结果 = (表达式1, 表达式2, 表达式3);
```

逗号操作符使用场景很有限，通常会用在 `for` 循环中，同时压缩代码也会用到大量的逗号表达式。

## 1. 运用在 for 循环中

```javascript
var arr = [1, 2, 3, 4, 5, 6];
var i, len;

for (i = 0, len = arr.length; i < len; i++) {
  console.log(arr[i]);
}
```

for 循环的 `初始语句` 部分要做多件事情的时候，就会用逗号操作符。

上述例子在 `初始语句` 部分做了两个赋值操作，对应两个赋值表达式，使用逗号操作符，就会从左到右依次执行，对变量 `i` 和变量 `len` 进行赋值操作。

for 语句的三个部分根据不同的业务都有可能用到逗号操作符，结合逗号操作符可以让表达式更符合语义。

如上述例子中的 `i = 0, len = arr.length;`，可以明确的表示把 i 变成 0 和 把 len 设置成 arr 的长度是初始操作。

但如果表达式太多，或者一个表达式很长，不太建议使用逗号操作符。

```javascript
var person = {
  father: {
    father: {
      father: {
        mother: {
          hobby: ['吃饭', '睡觉', '打游戏'],
        }
      }
    }
  }
};
var i, len;

for (i = 0, len = person.father.father.father.mother.hobby.length; i < len; i++) {
  console.log(person.father.father.father.mother.hobby[i]);
}
```

这里要获取到某个人的`爸爸的爸爸的爸爸的妈妈`的爱好，就要写一长串。

这种情况应该避免，防止一行代码太长，很多代码规范也规定了一行代码不应超过 100 个字符。

可以考虑把`爸爸的爸爸的爸爸的妈妈`单独取出来操作。

```javascript
var person = {
  father: {
    father: {
      father: {
        mother: {
          hobby: ['吃饭', '睡觉', '打游戏'],
        }
      }
    }
  }
};

var hobby = person.father.father.father.mother.hobby;
var i, len;

for (i = 0, len = hobby.length; i < len; i++) {
  console.log(hobby[i]);
}
```

再比如要输出一个二维数组从右上到左下的对角线数据，也可以应用逗号操作符。

```javascript
var arr = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];

var i, j, len;
for (i = 0, j = arr.length - 1; arr[i]; i++, j--) {
  console.log(arr[i][j]);
}
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb1784e09f3ed4417940412.jpg)

输出对角线，就是输出 `arr[0][2]` 、 `arr[1][1]` 、 `arr[2][0]`，利用逗号运算符就能很轻易的做到。

## 2. 用于代码压缩

通常不会手动对代码去进行压缩，都会借助自动化的工具。

代码进行压缩后，体积会有显著的变化，如 3.5.0 版本 `jquery` 未压缩的体积大概是 `280KB` ，压缩后大约为 `80KB`。

代码压缩不仅仅是去除了空格，还会改变语法结构，但通常不影响执行结果。

如以下函数：

```javascript
function encrypt(number) {
  number += 10;

  return number / 2 - 1;
}
```

压缩后的代码可能是这样的：

```javascript
function encrypt(n){return n+=10,n/2-1}
```

这是一个简单的数字加密函数，利用逗号表达式的特性，就可以用于在函数中对连续的几个表达式进行压缩，最后一个表达式的结果就会是函数的返回值。

## 3. 声明多个变量的逗号

在 `JavaScript` 中，同时想声明多个变量，也是使用逗号进行分隔。

```javascript
var i = 1, j, person = {}, arr = [];
```

**这里的逗号并不属于逗号操作符**，这是同时声明多个变量的语法，其具有自己的含义。

## 4. 小结

逗号操作符是很多前端开发者在第一次进行前端逆向工程（反向推测压缩混淆的代码）时碰到的，可以见得一般用的不多。

逗号操作符有时候可以增加代码语义性，但不恰当的使用也会适得其反。如果逗号操作符让一行代码变得很行，应考虑其他方式实现。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
