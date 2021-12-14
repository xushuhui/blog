# JavaScript switch 语句

> switch 语句评估一个表达式，将表达式的值与 case 子句匹配，并执行与该情况相关联的语句。—— MDN

switch 是另一种控制流程的方式，根据条件执行不同的代码块。

能用 switch 实现的都可以用 if 实现。

## 1. 基本语法

```javascript
  switch (表达式) {
    case 表达式结果为值1的时候:
      做的事情;
      break;
    case 表达式结果为值2的时候:
      做的事情;
      break;
    case ...:
      做的事情;
      break;
    case 表达式结果为值n的时候:
      做的事情;
      break;
    default:
      上面一个情况都没中的时候做的事情;
  }
```

switch 语句在执行的时候会先接受一个表达式，最后根据表达式的结果进行条件的匹配，即 `case` 后面的值。

```javascript
var num = 3;

switch (num + 1) {
  case 2:
    console.log('case的值是2');
    break;
  case 3:
    console.log('case的值是3');
    break;
  case 4:
    console.log('case的值是4');
  case 5:
    console.log('case的值是5');
  default:
    console.log('没有匹配到值');
}
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7a2f060a247f7a15440746.jpg)

`num` 为 3 ，所以加上 1 之后为 4 ，case 匹配到的就是 `4` ，所以输出了`case的值是4`，但是紧接着后面的`case的值是5`与`没有匹配到值`也被输出了。

这是因为分支内碰到 `break` 才会中断执行，如果不中断，即便后面的条件不匹配了，里面的代码块还是会被继续执行。

**需要注意的是 case 后面的值与表达式的结果在比较的时候是使用严格相等 (===) 的**。

## 2. default 的位置可以不固定

default 不一定要写在末尾，但通常推荐写在末尾。

```javascript
switch (1 > 2) {
  default:
    console.log('我是default');
  case true:
    console.log('1 不可能大于 2，肯定是代码写错了');
    break;
}
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7a2f710a313c5315440432.jpg)

在所有 case 都不匹配的时候，就会回去走 default 。

需要注意的是，default 语句块里也需要加 `break` ，不然会继续往下执行，直至碰到 `break` ，大部分情况下只有末尾的分支不需要加 `break` ，因为已经是最后一个分支了。

## 3. 灵活使用 break

switch 语句与 break 的特性结合可以很灵活。

如以下场景：

* 服务端返回了用户信息，当用户的 VIP 等级为 1、2、3 的时候，显示初级 VIP，VIP 等级为 4、5 的时候，显示中级 VIP，当 VIP 等级为 6 的时候，显示高级 VIP，否则显示普通会员。

```javascript
var user = { vip: 1 };

switch (user.vip) {
  case 1:
  case 2:
  case 3:
    console.log('初级vip');
    break;
  case 4:
  case 5:
    console.log('中级vip');
    break;
  case 6:
    console.log('高级vip');
    break;
  default:
    console.log('普通会员');
}

// 输出："初级vip"
```

利用没有 break 就会往下执行的特点，可以给条件归类。相比 `if` 语句，使用 switch 有更强的表现力。

* 页面中有一排图片，共 4 张，当用户选择了某一张后，隐藏这张图片前面的所有图片，取消选择后显示所有图片。

```javascript
function showImage(index) {
  console.log('显示第' + index + '图片');
}

function hideAllImage() {
  console.log('先隐藏所有图片');
}

// 点击事件
function event(e) {
  var selected = e.index; // 0表示没有选择 1表示选择第一张 以此类推

  hideAllImage();

  switch (selected) {
    default:
    case 1:
      showImage(1);
    case 2:
      showImage(2);
    case 3:
      showImage(3);
    case 4:
      showImage(4);
  }
}

event({ index: 0 });
```

其实这个需求，使用 switch 并不是最适合的场景，假如图片一多，上百上千张，这种方式就显得有些愚蠢了。

这提供这种实现方式不是为了显示他有多好，而是为了能在思考问题的时候，可以想到有这样的方案，来评判是不是更适合现有业务场景。

**没有最好的方案，只有最适合的方案。**

## 4. 小结

switch 语句可以做到的，if 语句都可以做到，实际开发应结合具体业务做选择。

依据 switch 的 break 和 default 特性，常常可以很方便的实现`其他方式需要大量额外代码`的需求。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
