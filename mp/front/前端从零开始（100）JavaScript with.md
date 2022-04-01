# JavaScript with

> with 语句就可以扩展一个语句的作用域链。——MDN

with 可以指定代码块的作用域，特别是在访问对象成员时，它使得代码非常简洁。

## 1. 基本语法

```javascript
with (表达式) {
  代码块;
}
```

表达式通常会直接给定一个对象。

```javascript
var person = {
  name: '小明',
  age: 666,
  major: 'english',
};

with (person) {
  console.log(name); // 输出："小明"
  console.log(age); // 输出：666
  console.log(major); // 输出："english"
}
```

可以看到，with 代码块内输出的变量，实际上是 `person` 对象的属性。

在 with 语句中，访问变量会先去看这个变量是不是在给定的对象中作为属性存在，如果存在，则取对象中属性的值，否则继续往上层找。

## 2. 不推荐使用 with

尽管 with 很方便，但 with 会造成诸多的问题。

### 2.1 造成语义不明

先看这段代码：

```javascript
function fn(block, height) {
  with (block) {
    console.log(height);
  }
}
```

这样的代码，有点让阅读代码的人难以理解具体含义。

阅读代码的时候，无法确定 height 要取 block 下的属性，还是要取形参 height。

### 2.2 造成污染

```javascript
function fn(obj) {
  with (obj) {
    b = 1;
  }
}
```

这个场景其实和上面的类似，假如 obj 中没有属性 b ，则会造成上层作用域的污染。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e7ad9b60a72d5dc13760432.jpg)

---

事实上很少有前端开发者会使用 with，在很长篇幅的代码中，with 会让代码逻辑变得不清晰，需要反复确认作用域。

尽可能的避免使用 with ，使用短变量名和合理的空行来使代码变得整洁易懂。

```javascript
var family = [
  {
    seniority: '子',
    name: '小明',
    detail: {
      birth: '1192/01/22',
      sex: 'man',
      hand: 8,
      leg: 44,
      deposit: 9999,
    },
  },
];

// 不好的表述形式
family[0].detail.birth = '1122/22/11';
family[0].detail.sex = '?';
var deposit = family[0].detail.family[0].detail;


// 取出要操作的数据 并控制空行
var detail = family[0].detail;

detail.birth = '2312/22/33';
detail.sex = '未知性别';

var deposit = detail.deposit;
```

## 3. 小结

with 的使用需要谨慎，尽量在可控范围内使用。

使用 with 常见的遇到问题的情况，都是恰巧污染了上层作用域，又恰巧污染的是同名的变量，这也产生的 bug 定位也相对困难。

> 前端框架 Vue 将 template 转化成 render 函数的时候就用到了 with。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
