# JavaScript 的 break 与 continue

break 与 continue 用来控制循环流程。

## 1. break

break 用来中断循环，在 for 循环和 while 循环中都适用。

如：从列表中寻找 id 为 n 的项（每一项的 id 是唯一的）

```javascript
var list = [{ id: 1 }, { id: 4 }, { id: 9 }, { id: 2 }];

var n = 9;

var i = 0, len = list.length;
for (; i < len; i++) {
  console.log(i);
  if (list[i].id === n) {
    console.log(list[i]);
    break;
  }
}
```

其实即便没有 break，上面的语句也可以正常执行，因为 id 是唯一的，即便将整个数组遍历完，也只有一个目标项。

但是如果整个列表有**上万**条数据，除了获取 id ，还有一些消耗性能的操作，那 break 就变得很关键。

假如最好的情况，寻找的项就在第一项，这时候碰到 break ，剩下无用的寻找操作就都不会执行了，也没有必要执行。

将 break 应用在 while 中也是同理，满足某些条件的时候用来中断 while 循环。

如：游戏中的怪物产生了 4 次连击，伤害由四个随机 100 至 2000 的值累加构成，假如在累加的伤害过程中值达到 5000，则直接取当前累加结果作为最终伤害。

```javascript
var total = 0; // 累积伤害
var time = 1; // 累加次数

while (time <= 4) {
  var one = Math.floor(Math.random() * 2001 + 100);

  console.log(one);

  total += one;

  if (total > 5000) {
    break;
  }

  time++;
}
```

通过 break，可以很容易的达到这个需求。

当然不使用 break 也是可以的，可以在达到条件后，将 `time` 变量累加到一个循环条件不成立的情况，也会跳出循环，但是建议使用 break ，否则可能还要去控制 break 之后的逻辑，因为 break 之后不一定就是循环结束了，也许还有其他操作。

> 在 switch 语句中也有 break 参与，详细的作用可以参阅 switch 语句章节。

## 2. continue

当循环中碰到 continue ，则会跳过这次循环，进入下一次循环。

如：输出 0 至 100 之间的奇数

```javascript
var i;
for (i = 0; i <= 100; i++) {
  if (i % 2 === 0) {
    continue;
  }

  console.log(i);
}
```

当碰到偶数的时候，则跳过这次循环，反之则为奇数的情况，会输出值。

这个例子比较简单，当一个逻辑复杂的循环体出现的时候，continue 可以让代码块变得稍微简洁。

如：对用户列表进行批量操作，只操作 id (id 一定是数字） 末尾为 1 的用户。

```javascript
// 这是一份伪代码
var users = [ ... 很多用户 ]; //

var i, len;
for (i = 0, len = users.length; i < len; i++) {
  var user = users[i];

  if (user.id % 10 === 1) { // 取到末尾
    continue;
  }

  if (user.edition) { // 如果用户的版本不是免费的 0是免费 其他则是收费
    // 做一些操作
  }

  if (user.money) { // 如果用户有余额
    // 做一些操作
  }

  // ...
}
```

像这样，通过 continue 就可以让流程更清晰。

在 while 中的应用是一样的，continue 也可以在 while 中跳过当前循环。

使用 while 输出 0 至 100 的奇数：

```javascript
var num = -1;
while (num < 100) {
  num++;

  if (num % 2 === 0) {
    continue;
  }

  console.log(num);
}
```

整个流程与 for 循环非常相似。

> 许多代码规范不提倡使用 continue ，因为开发者水平的参差不齐，常有开发者在使用 continue 的时候，让整体逻辑产生了跳跃性，这不利于未来的代码维护者来对代码做阅读理解。

## 3. 小结

break 和 continue 适用于 for 循环和 while 循环。

continue 的使用没有 break 频繁，使用 continue 的时候需要注意，尽量不要使程序跳跃幅度过大，不利于后期维护。

所有需要使用 continue 的地方，都可以不使用 continue 实现。

switch 语句中也可以使用 break 来结束 switch 语句的执行。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
