# for 语句

for 语句是循环语句中的一种。

for 语句可以使程序在某一个条件下重复执行一段代码。

## 1. 基本语法

for 语句相对于 if 语句稍微复杂，通常为以下格式：

```javascript
for (初始语句; 条件; 条件为真值时执行的语句) {
  // 循环体
}
```

`初始语句`会在循环开始前进行执行。

`条件`会在每次循环结束后执行，结果影响循环语句是否要继续执行。

`条件为真时执行的语句`通常会用来影响下一次计算条件的结果。

## 2. 为什么需要循环语句

循环的应用非常广泛，如果有学习过数据结构与算法，会发现大部分的算法都需要循环介入，否则会使程序变得冗余复杂。

如我们需要生成一段 HTML 列表代码：

```javascript
<div id="container"></div>

<script>
  var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];

  var li1 = '<li>' + arr[0] + '</li>';
  var li2 = '<li>' + arr[1] + '</li>';
  var li3 = '<li>' + arr[2] + '</li>';
  var li4 = '<li>' + arr[3] + '</li>';
  var li5 = '<li>' + arr[4] + '</li>';
  var li6 = '<li>' + arr[5] + '</li>';
  var li7 = '<li>' + arr[6] + '</li>';
  var li8 = '<li>' + arr[7] + '</li>';
  var li9 = '<li>' + arr[8] + '</li>';

  var ul = '<ul>'+ li1 + li2 + li3 + li4 + li5 + li6 + li7 + li8 + li9 + '</ul>';

  document.getElementById('container').innerHTML = ul;
</script>
```

可以发现这里生成列表的代码很冗余，都是重复的操作，如果使用 for 循环，代码可以改成这样：

```javascript
<div id="container"></div>
<script>
  var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];

  var lis = '';

  // 单独拿到数组的长度
  var len = arr.length;
  // 声明变量i用来计数
  var i;
  for (i = 0; i < len; i++) {
    lis = lis + ('<li>' + arr[i] + '</li>');
  }

  var ul = '<ul>' + lis + '</ul>';

  document.getElementById('container').innerHTML = ul;
</script>
```

这里使用 for 循环生成了一个列表，效果和上述声明九个变量然后做拼接的方式是一样的。

循环开始前先声明的变量`i`用于计数，表示当前循环到第几次。

在循环开始前先将`i`设置为 0，这只会执行一次，随后比较当前循环的次数是否小于数组长度，如果比数组长度小则执行`i++`，`i++`会先使用`i`的值，再做累加 ( i = i + 1) 操作，随后执行循环体，重复上述操作。

具体的流程如下：

len 的值为 9

* 第一次循环    * i 的值为 0，i 是小于 len 变量的，所以 arr[0] 的值就是 1，这个时候 lis 被累加了`<li>1</li>`。

* 第二次循环    * i 的值为 1，i 是小于 len 变量的，所以 arr[1] 的值就是 2，这个时候 lis 被累加了`<li>2</li>`。

* 第三次循环    * i 的值为 2，i 是小于 len 变量的，所以 arr[2] 的值就是 3，这个时候 lis 被累加了`<li>3</li>`。

* 第四次循环    * i 的值为 3，i 是小于 len 变量的，所以 arr[3] 的值就是 4，这个时候 lis 被累加了`<li>4</li>`。

* 第五次循环    * i 的值为 4，i 是小于 len 变量的，所以 arr[4] 的值就是 5，这个时候 lis 被累加了`<li>5</li>`。

* 第六次循环    * i 的值为 5，i 是小于 len 变量的，所以 arr[5] 的值就是 6，这个时候 lis 被累加了`<li>6</li>`。

* 第七次循环    * i 的值为 6，i 是小于 len 变量的，所以 arr[6] 的值就是 7，这个时候 lis 被累加了`<li>7</li>`。

* 第八次循环    * i 的值为 7，i 是小于 len 变量的，所以 arr[7] 的值就是 8，这个时候 lis 被累加了`<li>8</li>`。

* 第九次循环    * i 的值为 8，i 是小于 len 变量的，所以 arr[8] 的值就是 9，这个时候 lis 被累加了`<li>9</li>`。

* 第十次循环    * i 的值为 9，i 等于 len 变量，不再小于 len，条件不成立，循环结束。

虽然流程看起来复杂，但是代码的可维护性得到了提高，冗余代码也减少了，如果这个时候`li`标签需要加一些属性，如`class`或者`style`，只需要修改循环体中的一行代码即可。

这种形式的 for 循环还有一种语法：

```javascript
for (初始语句; 条件; 条件为真值时执行的语句) 需要循环的语句;
```

和 if 语句很像，这种属于行循环语句，这种用到的比较少，因为代码的可阅读性比较低，而且一般用到循环的场景都不止一行代码。

## 3. for … in

for…in 循环可以用来遍历对象的属性名。

```javascript
var obj = {
  name: '小红',
  age: 12,
  hobby: ['打篮球', '唱歌'],
};

for (key in obj) {
  console.log(obj[key]);
}

// 输出：
//   "小红"
//   12
//   ["打篮球", "唱歌"]
```

每一次遍历拿到的 key 就是对象的某一个属性名，当属性名被遍历完后会自动退出循环。

有部分 key 是无法遍历到的，具体规则可以参阅对象章节。

## 4. 无限循环

```javascript
for (;;) {
  console.log('loop...');
}
```

这样的循环语句会陷入无限循环。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7ae94e0a91c46413760596.jpg)

**大部分无限循环会让浏览器卡死，需要强制退出浏览器！**

## 5. 循环应用的例子

### 5.1 判断一个数是不是质数

```javascript
var num = 17;

var flag = false;

var len;
var i;

for (i = 2, len = 17 -1; i <= len; i++) {
  if (num % i === 0) {
    flag = true;
    break; // break可以中断循环
  }
}

if (flag) {
  console.log(num + '不是质数');
} else {
  console.log(num + '是质数');
}

// 输出："17是质数"
```

首先要知道什么是质数，质数就是只能被 1 和本身整除的数。

所以如果要判断`num`是不是质数，只需要去掉头尾，从`2`循环到`num - 1`，用`num`对每一个循环数做取余操作，如果存在余数为 0 的，就说明中间有个数可以被整除，那就不是质数，反之就是质数。

### 5.2 计算阶乘

```javascript
var num = 4;

var result = 1;

var i;
for (i = num; i > 1; i--) {
  result = result * i;
}

console.log(result); // 输出：24
```

阶乘是所有小于及等于某一数的正整数的积，如 4 的阶乘，在数学中表示为`4!`，结果为`4 * 3 * 2 * 1`。

在代码中，就可以用一个变量来保存每次做乘法的结果，如 4 的阶乘，就可以用变量`result`记录结果，初始值为 1，循环可以从 4 循环到 1，每次将循环到的值乘以`result`，循环结束后就可以得到结果。

## 6. 小结

循环语句和条件语句一样，是给语言带来处理业务能力的重要特性之一。

通常如果需要连续执行多次的重复操作，都应该考虑使用循环来解决。

1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
