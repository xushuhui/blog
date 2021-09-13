# getter & setter

利用 getter/setter 可以拦截操作对象的属性，在设置属性前或获取属性前做一些事情。

## 1. getter

> get 语法将对象属性绑定到查询该属性时将被调用的函数。(MDN)

```javascript
// 语法
var 对象 = {
  get 属性名() {
    // 做一些事情 balabala
    return 值;
  }
};
```

getter 在获取一个属性时被调用，同时返回 getter 的返回值。

```javascript
const student = {
  score: {
    english: 10,
    chinese: 99,
    math: 6,
  },

  // 总分 通过计算得出
  get totalScore() {
    var score = this.score;

    return score.english + score.chinese + score.math;
  }
};

console.log(student.totalScore); // 115
```

访问 `student.totalScore`，实际上时访问了 `totalScore` 这个 `getter`，`getter` 本质上是个函数，所以可以像写一个函数一样写 `getter`，最后返回 `getter` 的返回值作为访问属性的属性值。

需要注意的是，没有特殊情况，不要在 getter 中访问自身。

```javascript
var obj = {
  key: 1,

  get key() {
    return this.key;
  },
}
```

这样会导致无限访问 `key` 这个 `getter`。

## 2. setter

> 当尝试设置属性时，set 语法将对象属性绑定到要调用的函数。(MDN)

```javascript
// 语法
var 对象 = {
  set 属性名(值) {
    // 做一些事情 balabala
    // this.某个属性 = 值;
  }
};
```

setter 在一个属性被赋值时调用，同时这个值会被作为参数传递给 setter。

```javascript
const student = {
  score: {
    english: 10,
    chinese: 99,
    math: 6,
  },

  // 总分 通过计算得出
  get totalScore() {
    var score = this.score;

    return score.english + score.chinese + score.math;
  },

  set english(value) {
    this.score.english = value;
  },

  set chinese(value) {
    this.score.chinese = value;
  },

  set math(value) {
    this.score.math = value;
  },
};

console.log(student);

student.math = 66;
student.chinese = 150;
student.english = 77;

console.log(student);
console.log(student.totalScore);
```

这里通过三个 `setter` 来设定对应的分数，这样就不用使用 `student.score.学科` 的方式赋值了，可以省略 `score`。

和 `getter` 同理，使用 setter 时，`setter` 名和最终要设置值的属性不应同名，否则会无限设置这个值。

## 3. 使用 Object.defineProperty 设置 setter/getter

利用 `defineProperty` 方法也可以设置 `setter/getter`。

```javascript
var person = {
  _cash: 1,
  _deposit: 99999,
};

Object.defineProperty(person, 'money', {
  get: function() {
    return this._cash + this._deposit;
  },
});
Object.defineProperty(person, 'cash', {
  set: function(val) {
    console.log('现金发生了改变');
    this._cash = val;
  },
});

person.cash = 2;
console.log(person.money);
```

![图片描述](https://img.mukewang.com/wiki/5ec2ae8d0910109213660634.jpg)

## 4. 小结

getter/setter 可以充当属性拦截器的角色，在设置和访问属性的时候做一些额外的事情。

灵活使用 getter/setter 可以使开发变得更有效率，许多框架的核心机制就是灵活、巧妙的使用了 `getter/setter`。

getter/setter 是 `ES5` 中的特性，所以要注意 IE8 并不支持。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
