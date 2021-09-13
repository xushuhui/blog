# ES6+ Generator 基础

## 1. 前言

前面我们花了三节深入地学习了 ES6 的异步解决方案 Promise，本节学习的生成器也是为了解决异步而生的，但是它的出发思路和 Promise 截然不同。

上节我们学习了 ES6 中 [迭代](http://baike.imooc.com/tms/section/add?tid=3976&id=2732&type=0) 的相关内容，并实现了一个迭代器。我们知道实现一个迭代器，我们需要手动添加对象的 `Symbol.iterator` 属性，并需要实现 next 方法。那么有没有什么可以帮助我们自动实现迭代器呢？ES6 给出了生成器的方法来满足我们的需求。我们不需要在对象上添加 `Symbol.iterator` 属性，使用生成器函数就可以实现迭代器的功能。本节我们将学习生成器的相关概念和基础用法。

生成器是一个灵活的结构，能使得一个函数块内部暂停和恢复代码执行的能力。在实际应用中，使用生成器可以自定义迭代器和协程。

## 2. 生成器对象和生成器函数

有些概念是我们必须要理解的，前面在学习迭代器的时候，我们学习了迭代协议和迭代器协议，实现一个迭代器需要满足这两个协议才算是一个真正的迭代器。而本节的生成器和生成器函数也是如此，我们也需要知道生成器对象和生成器函数概念和它们直接的关系。

Generator 就是我们说的生成器，它包含两个概念 `生成器对象和生成器函数`。首先，要理解的是生成器对象和迭代器的关系，生成器对象是遵守迭代协议和迭代器协议实现的 Iterable 接口，可以理解生成器对象其实也是一个迭代器；然后，我们需要理解什么是生成器函数，生成器函数是由 `function *` 来定义的，并且返回结果是一个 Generator 对象。

生成器是一个特殊的函数，在调用后会返回一个生成器对象，这个生成器对象是遵守可迭代协议和迭代器协议实现的 Iterable 接口。生成器可以使用 yield 关键字来暂停执行的生成器函数：

```javascript
function* generator() {
  yield 'a';
  yield 'b';
}

var gen = generator();	// Object [Generator] {}
```

### 2.1 Generator.prototype.next()

生成器的 next () 方法和迭代器返回的结果是一样的，返回了一个包含属性 `done` 和 `value` 的对象，该方法也可以通过接受一个参数用以向生成器传值。

使用 yield 返回的值会被迭代器的 next () 方法捕获：

```javascript
var gen = generator();

gen.next()	// {value: 'a', done: false}
gen.next()	// {value: 'b', done: false}
gen.next()	// {value: undefined, done: true}
```

从上面代码的执行结果可以看出，生成器函数在执行后会返回一个生成器对象，这个生成器对象满足迭代协议和迭代器协议，所以我们可以去手动调用它的 next () 方法去获取每一步的返回值。从这里可以看出，生成器其实就是迭代器的一个应用，并且这个应用会在异步中大放异彩。

### 2.2 Generator.prototype.return()

`return()` 方法返回给定的值并结束生成器。

```javascript
var gen = generator();

gen.next();        // { value: 'a', done: false }
gen.return("imooc"); // { value: "imooc", done: true }
gen.next();        // { value: undefined, done: true }
```

另外，如果对已经完成状态的生成器调用 `return(value)` 则生成器会一直保持在完成状态，如果出入参数，`value` 会设置成传入的参数，`done` 的值不变：

```javascript
var gen = generator();

gen.next(); // { value: 1, done: false }
gen.next(); // { value: 2, done: false }
gen.next(); // { value: undefined, done: true }
gen.return(); // { value: undefined, done: true }
gen.return(1); // { value: 1, done: true }
```

### 2.2 Generator.prototype.throw()

`throw()` 方法用来向生成器抛出异常，并恢复生成器的执行，返回带有 `done` 及 `value` 两个属性的对象。

```javascript
function* generator() {
  while(true) {
    try {
       yield 'imooc'
    } catch(e) {
      console.log("Error caught!");
    }
  }
}
var gen = generator();
gen.next(); // { value: "imooc", done: false }
gen.throw(new Error("error")); // "Error caught!"
```

## 3. Generator 案例

### 3.1 类数组转化

将一个类数组转化为一个真正的数组方式有很多，ES6 提供了 `Array.from()` 可以将类数组转化为数组 。另外在一些函数中可以使用 `[...argument]` 的方式转化类数组。

```javascript
function fn() {
  const arg = [...arguments];
  console.log(arg);
}
fn(1, 2, 3);	// [1, 2, 3]
```

当然我们知道类数组的定义，所以我们自己定义一个类数组，看能不能使用展开运算符将类数组转化为数组：

```javascript
const likeArr = {
  0: 1,
  1: 2,
  length: 2,
}
console.log([...likeArr]);	// Uncaught TypeError: likeArr is not iterable
```

上面代码中我们定义了一个类数组，但是使用展开运算符报错了，提示我们 likeArr 不是一个迭代器。因为在函数中类数组是内部帮我们实现了迭代器的功能，而我们自己定义的类数组是不具有迭代器功能的，那我们来自己实现一个：

```javascript
likeArr[Symbol.iterator] = function() {
  let index = 0;
  return {
    next: () => {
      return { value: this[index], done: index++ === this.length}
    }
  }
}
console.log([...likeArr]);	// [1, 2]
```

上面的代码我们在 likeArr 对象上定义了 `Symbol.iterator` 它具有迭代功能。上面代码中我们需要手动地去实现 next () 方法，这比较麻烦，那能不能简化一下呢？我们的生成器函数就出场了：

```javascript
likeArr[Symbol.iterator] = function* () {
  let index = 0;
  while (index !== this.length) {
    yield this[index++];
  }
}
console.log([...likeArr]);	// [1, 2]
```

上面的代码使用了生成器函数，并且没有去手动实现 next () 方法，从这里我们也能很清楚地知道迭代器和生成器的关系。而且使用生成器函数更加简单方便。

### 3.2 单步获取质数

还有一个案例是面试中经常会考到的：

> 题目：实现一个函数，每次调用返回下一个质数，要求不使用全局变量，且函数本身不接受任何参数

从题目的要求可以知道，这个函数每次调用都会返回一个质数，也就是说每次调用后都会返回一个函数。

首先我们定义一个判断一个数是否为质数的方法：

```javascript
function isPrime(num) {
  for (let i = 2; i <= Math.sqrt(num); i++) {
    if (num % i === 0) {
      return false
    }
  }
  return true
}
```

传统的方式是使用闭包方法来解决：

```javascript
function primeHandler() {
  let prime = 1
  return () => {
    while (true) {
      prime++
      if (isPrime(prime)) {
        return prime
      }
    }
  }
}

const getPrime = primeHandler()
console.log(getPrime());	// 2
console.log(getPrime());	// 3
console.log(getPrime());	// 5
```

既然是单步执行的，那么我们就可以使用迭代器方式实现：

```javascript
var prime = {}
prime[Symbol.iterator] = function() {
  let prime = 1;
  return {
    next() {
      while(true) {
        prime++
        if (isPrime(prime)) {
          return prime;
        }
      }
    }
  }
}
var getPrime = prime[Symbol.iterator]().next;
console.log(getPrime());	// 2
console.log(getPrime());	// 3
```

上一个实例我们知道实现迭代器的方式是很麻烦的，可以使用生成器函数去替代迭代器的功能，所以上面的代码可以使用生成器函数改造如下：

```javascript
function* primeGenerator () {
  let prime = 1
  while (true) {
    prime++
    if (isPrime(prime)) {
      yield prime
    }
  }
}

var getPrime = primeGenerator().next().value
console.log(getPrime());	// 2
console.log(getPrime());	// 3
```

## 4. 小结

本节我们主要学习了生成器的概念和用法，需要生成器对象是由生成器函数返回的结果，生成器对象是遵守迭代协议和迭代器协议实现的 Iterable 接口。生成器其实就是对迭代器的应用。另外，通过两个案例更加深刻地理解了生成器的应用场景，对比了生成器和迭代器的不同。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
