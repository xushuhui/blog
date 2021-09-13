# new 运算符与构造函数

当一个函数被 `new` 运算符调用的时候，这个函数就会被称为构造函数。

任何函数都能被 `new` 运算符调用，但是一般会从设计上将一个函数考虑为构造函数，提供给 `new` 运算符调用。

```javascript
function Human(name, gender) {
  this.name = name;
  this.gender = gender;
}

var human = new Human();
```

## 1. 构造函数的作用

构造函数的主要作用是用于生成对象。

有其他面向对象语言开发经验的同学可能会觉得使用 `new` 运算符的语法和创建类的示例很像，其实本质是不一样的。

结合原型的特性，在 `JavaScript` 中也能实现类似于类的一套机制。

> 关于构造函数和原型的处理关系，原型章节已经有详细介绍，具体内容可以参考原型章节。

## 2. new 运算符的运算机制

使用 new 运算符调用函数的时，背后有一套运行机制，这套机制说明了构造函数是怎么产生对象的。

当 new 运算符调用函数时，大致会进行以下几个操作：

1. 创建一个空对象
2. 将函数的 this 指向这个空对象
3. 执行函数
4. 如果函数没有指定返回值，则直接返回 this（一开始创建的空对象），否则返回指定返回值

```javascript
function Person(name, gender, age) {
  this.name = name;
  this.gender = gender;
  this.age = age;
}

var person = new Person('小明', '男', 17);

console.log(person.name);
```

这样就能理解为什么使用 `new` 操作符可以生成对象了。

这个机制也是面试的高频题。

## 3. 小结

构造函数用于生成对象，理解构造函数和原型机制非常重要，不但是面试中的高频题，也可以提升编写高质量、可复用的代码的能力。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
