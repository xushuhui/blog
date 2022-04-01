# JavaScript while 语句

> while 语句可以在某个条件表达式为真的前提下，循环执行指定的一段代码，直到那个表达式不为真时结束循环。—— MDN

while 语句也是一种循环语句，也称 while 循环。

while 循环接收一个表达式，当这个表达式结果非 `false` 的时候，就会执行 while 循环的代码块。

## 1. 基本语法

```javascript
while (表达式) {
  表达式结果为真时候执行的代码;
}
```

while 的语法相对简单，其使用的频率没有`for循环`高，可以使用 for 循环完成的都可以使用 while 循环完成，反之亦然。

例如输出 0 到 100 的偶数，使用 while 就可以这样做：

```javascript
var num = 0
while (num <= 100) {
  if (num % 2 === 0) {
    console.log(num);
  }

  num++;
}
```

```javascript
var i;
for (i = 0; i <= 100; i++) {
if (i % 2 === 0) {
    console.log(i);
  }
}
```

对比 for 循环，许多场景两者在同一问题的处理方式上区别不大。

for 循环将初始操作、循环条件、条件判断后要做的事情放在了规定的位置，而 while 循环只是将这些操作换个地方写而已。

事实上 for 循环也可以写得像 while 循环一样：

```javascript
var num = 0;
for (;num <= 100;) {
  if (num % 2 === 0) {
    console.log(num);
  }

  num++;
}
```

换成这样的写法可以说是高度一致了。

## 2. 使用 while 的场景

根据个人的经验，while 比 for 循环用到的少，通常有三种情况：

* 需要“无限循环”

相比 for 循环，while 来做无限循环更直接：

```javascript
for (;;) {
  console.log('我停不下来了!!');
}

console.log('永远也不会执行到这里 :)');

while(true) {
  console.log('我也停不下来了!!');
}
```

那么问题就来了，为什么需要无限循环。

主要是想将复杂的条件拿出来，自己判断条件是否达成，然后使用 `break` 中断循环。

* 有许多不可控的量构成循环条件

当循环条件为一个复杂表达式，而又不需要用到循环计数变量的时候（循环变量不一定非要是一个数字进行累加的）。

* 通过循环对已有变量做操作，并将这个变量作为条件或者构成条件的成员

## 3. 小结

可以使用 while 语句实现的需求，都可以使用 for 语句实现，但 while 更适合复杂循环条件的场景。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
