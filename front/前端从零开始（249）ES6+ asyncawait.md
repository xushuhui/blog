# ES6+ async/await

## 1. 前言

前面几节我们已经学习了解决异步的方案 Promise 和 Generator，上一节我们也通过一个案例使用 `Promise + Generator` 实现了一个比较好的异步解决方案。同时我们实现了一个简版的 co 库，让我们在使用 Generator 函数处理异步任务时更加方便，但这不是最完美的解决方案。

本节我们将学习 ES7 推出的 `async/await` 其特性是对 JS 的异步编程进行了重要的改进，在不阻塞主线程的情况下，它给我们提供了使用同步代码的风格来编写异步任务的能力。另外，我们要明确的是 `async/await` 其实是 `Promise + Generator` 的语法糖，为了帮助我们像写同步代码一样书写异步代码，代码风格更优雅，错误捕获也更容易。

本节我们将通过对上一节案例的改造。在不需要 co 库的情况下直接使用 `async/await` 让我们更加深刻地理解异步方案的演变过程。

## 2. 改造上节案例

[上一节](http://baike.imooc.com/tms/section/todo) 我们通过一个案例来讲解 `Promise + Generator` 在实际应用中的使用，通过 Generator 函数 和 yield 让异步代码看起来像同步代码一样执行。但是这样里面存在的一个问题就是生成器函数直接执行，需要手动处理。为了解决深层回调的问题我们借助了 co 库来帮助我们去执行生成器函数，从而解决了回调地狱的问题。下面是上一节的代码。

```javascript
const ajax = function(api) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (api === 'api_1') {
        resolve('api_2');
      }
      if (api === 'api_2') {
        resolve(100);
      }
    }, 0)
  })
}

function * getValue() {
  const api = yield ajax('api_1');
  const value = yield ajax(api);
  return value;
}

co(getValue()).then(res => {
  console.log(res);
})
```

上面的代码中 `getValue` 是生成器函数，不能直接调用，这里用 co 库来进行执行，然后通过 Promise 的链式调用获取执行后的结果。但是这里借助了 co 的库，我们其实最希望的是能像执行普通函数一样直接调用 `getValue` 就能执行并得到结果。 `async/await` 的出现就是为了抹平在调用时所做的额外步骤。那让我们看看 `async/await` 是怎么用的：

```javascript
async function getValue() {
  const api = await ajax('api_1');
  const value = await ajax(api);
  console.log(value)
  return value;
}
getValue()	// 控制台打印 value的值是：100
```

上面的代码中我们可以看出使用 `async/await` 定义的 `getValue` 函数和生成器函数 `*/yield` 定义的基本相同，但是在执行时 `async/await` 定义的函数直接调用即可。从这里我们就能看到 `async/await` 的优点，无需过多的操作非常优雅和简洁。

## 3. 用法

上面我们基本了解了 async 函数，下面我们就来看看它的基本使用和需要注意的地方。

定义一个异步函数时需要使用 `async` 和 `function` 关键字一起来完成，类似生成器函数中的 `yield` 来暂停异步任务，在 async 函数中使用 `await` 关键去等待异步任务返回的结果。

async 函数其本质是 `Promise + Generator` 函数组成的语法糖，它为了减少了 Promise 的链式调用，解放了 Generator 函数的单步执行。主要语法如下：

```javascript
async function name([param[, param[, ... param]]]) {
    statements
}
```

上面代码中的 statements 是函数主体的表达式，async 函数可以通过 `return` 来返回一个值，这个返回值会被包装成一个 Promise 实例，可以被链式调用。下面我们来看两段等价代码。

```javascript
// 下面两段代码时相同的
async function foo() {
   return 100
}
function foo() {
   return Promise.resolve(100)
}

// 下面两段代码时相同的
async function foo() {
  await 1;
}
function foo() {
   return Promise.resolve(1).then(() => undefined)
}
```

上面的两段代码效果时相同的，这里我们就不去探究 async 函数是怎么实现的，其大概原理类似上节写的 co 库，有兴趣的小伙伴可以去 [babel](https://babeljs.io/repl) 上去看看 async 函数编译是什么样子的。

当在 async 函数中返回的是一个普通值或 await 后跟一个普通值时，此时的 async 函数是同步的。在 Promise 中失败是不能被 `try...catch` 捕获的，需要通过 `catch` 的方式来捕获错误。而使用 async 函数则是可以通过 `try...catch` 来捕获。

```javascript
async function foo() {
	return new Error('Throw an error');
}

foo().then(res => {
  console.log(res)
}).catch(err => {
  console.error(err)	// Error: Throw an error
})

async function foo2() {
  try{
    var v = await foo()
    console.log(v)
  } catch(e) {
    console.log(e);	// Error: Throw an error
  }
}
foo2()
```

上面的代码中在执行 `foo()` 直接抛出了一个错误，而 Promise 和 async/await 对错误的捕获是不同的，我们知道 Promise 是通过 `then` 中的失败回调和 `catch`来捕获错误的，而 async 函数使用的是 `try...catch` 更像同步的方式。

### 3.1 错误捕获

但是有个问题，当程序需要同时处理多个异步任务时，那我们使用 `async/await` 怎样捕获那个异步任务出现错误呢？try 块中的代码只要程序出现错误就会抛出错误，但是不知道是哪个异步任务出错了不利于定位问题。如果使用多个 `try...catch` :

```javascript
const task = function (num) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (num === 300) {
        reject('throw error')
      } else {
       	resolve('imooc');
      }
    }, 1000)
  })
}

async function foo() {
  try {
    let res1 = await task(100);
    try {
      let res2 = await task(200);
      try {
        let res3 = await task(300);
      } catch(e) {
        console.log('res3', e)
      }
    } catch(e) {
      console.log('res2', e)
    }
  } catch(e) {
    console.log('res1', e)
  }
}
foo()	// res3 throw error
```

看到上面的代码你是不是觉得很难受啊，又回到了嵌套地狱的原始问题了。async 函数在异常捕获时，没有非常完美的解决方案，这主要源自依赖 `try...catch` 对错误的捕获。但有一些还算比较优雅的解决方案，我们已经知道了 async 函数返回的是一个 Promise 那么我们是不是可以使用 Promise 的 `catch` 来捕获呢？答案是当然的呢。

```javascript
async function foo() {
  let res1 = await task(100).catch(err => console.log('res1', err));
  let res2 = await task(200).catch(err => console.log('res2', err));
  let res3 = await task(300).catch(err => console.log('res3', err));
}
foo()	// res3 throw error
```

上面的代码看起来就比嵌套的 `try...catch` 感觉好很多，这也是一个比较好的解决方式。在使用 `catch` 时需要弄清楚 Promise 和 async 函数之间的关系，不然就很难理解这种写法。

### 3.2 滥用 async/await

既然 `async/await` 这么优雅简洁，那在编程的过程中都使用这个就好啦！其实这里是一个坑，很多时候 `async/await` 都会被滥用导致程序卡顿，执行时间过长。

```javascript
async function foo() {
  let res1 = await task(100);
  let res2 = await task(200);
  let res3 = await task(300);

  return { res1, res2, res3 }
}
foo()
```

在很多时候我们会写成这样的代码，如果后一个任务依赖前一个任务这样写完全没问题，但是如果是三个独立的异步任务，那这样写就会导致程序执行时间加长。这样的代码过于同步化，我们需要牢记的是 await 看起来是同步的，但它仍然属于异步的内容，最终还是走的回调，只是语言底层给我们做了很多工作。

针对没有关联的异步任务我们需要把它们解开，

```javascript
const task = function (num) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
       	resolve('imooc ' + num);
    }, 1000)
  })
}

async function foo() {
  let res1Promes = task(100);
  let res2Promes = task(200);
  let res3Promes = task(300);

  let res1 = await res1Promes;
  let res2 = await res2Promes;
  let res3 = await res3Promes;

	console.log({ res1, res2, res3 })

  return { res1, res2, res3 }
}
foo();	// { res1: 'imooc 100', res2: 'imooc 200', res3: 'imooc 300' }
```

这里需要明白的一点是：为什么要把 task 拿到 await 外面去执行呢？await 的本质就是暂停异步任务，等待返回结果，等到结果返回后就会继续往下执行。还要知道的是每个 task 都是一个异步任务，像之前的那种写法，await 会等待上一个异步任务完成才会走下一个。而我们把 task 拿出来了，也就是每个 task 会按照异步的方式去执行。这个时候三个 task 都已经开始执行了，当遇到 await 就只需要等到任务完成就行。所需要的时间是异步任务中耗时最长的，而不是之前的总和。

## 4. 小结

本节我们主要通过延续上一节的案例，用 async 函数给出了最优的解决方案，从而完善了整个异步演变的过程，让我们更加清晰地理解为什么会有 Promise？为什么会有生成器？为什么会有 async/await？由浅入深层层递进地讲解了 ES6 以后对异步任务处理的演变。然后我们主要学习了 async 函数的基本使用和错误处理的捕获。最后，我们讲解了如果不滥用 async 函数的案例，让我们在以后写程序的过程中更加得心应手。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
