# JavaScript 对象

> 对象 指包含数据和用于处理数据的指令的数据结构。对象有时也指现实世界中的一些事，例如在赛车游戏当中一辆车或者一幅地图都可以是一个对象。—— MDN

JavaScript 中的对象由`属性`和`方法`组成。

属性可以是任意 JavaScript 中的数据类型，方法则是一个函数。

## 1. 创建对象

对象的格式如下：

```javascript
{
  prop1: 'value1',
  prop2: 666,
  prop3: {},
  method1: function() {
  },
}
```

其中 `prop1`、`prop2`、`prop3` 都是属性，`method1` 是方法，属性是任意的数据类型，方法则是一个函数。

可以看到一个属性或者方法，在其名字与值中间采用冒号 `:` 分隔，属性和方法之间采用逗号 `,` 分隔。

通常属性和方法的名字会被称为`属性名`与`方法名`，属性的值称为`属性值`，方法的值则直接叫做 `方法`。

对象符合 `key/value` 的结构，一个 `key` 对应一个 `value`，这种结构也被称为键值对，属性名和方法名都是 `key`，他们的值都是 `value`。

> 注意：在 JavaScript 的对象中，方法和属性的表示其实是一样的。
>
> 一个属性的值如果是一个函数，则就把这个属性称之为方法，只是称呼上的不同。

对象最简单的创建方式就是使用`对象字面量`。

```javascript
var person = {};
```

以上创建了一个最简单的空对象。

对象在 JavaScript 中的应用范围非常广。

可以试想一下，如果需要用变量描述一个人的信息，包括名字、性别、年龄、双亲信息，同时就要表示这个人的一些行为，如说话。

显然数值、字符串、布尔类型这些数据类型是不太适合描述这样一个人的。

这个时候就可以考虑使用对象，也就是说当需要描述的事物变得复合（无法使用单一的数据类型描述的时候），就可以考虑使用对象存储数据。

```javascript
var person = {
  'name': '小明',
  'age': 17,
  isAdult: false,
  sex: 'man',
  hobby: ['eat', 'sleep', 'play doudou'],
  parents: {
    mather: {
      name: '大红',
    },
    father: {
      name: '大明',
    },
  },
  say: function() {
    console.log('我叫' + this.name + '，我今年' + this.age + '岁了。');
  },
};

console.log(person); // 在控制台可以观察 person 对象
```

上面这个 `person` 变量就是一个对象，用于描述一个人，这个人具有 `name`、`age` 等属性与方法。

在控制台输出对象后，可以对其展开，观察他的内容。

其中部分属性在声明的时候，属性名上加上了引号，这在 JavaScript 中是被允许的，但一般情况下不会加上引号，原因之一是没有必要加，不应该与字符串混淆，另外一点就是可以消除初学者对 JSON 和 JavaScript 对象在理解上的歧义。

## 2. 操作对象

创建对象主要是为了设置、访问对象的属性值，调用对象的方法等。

### 2.1 访问对象的属性值

访问属性有 2 种方式：

1. `对象.属性名`
2. `对象['属性名']`

```javascript
var obj = {
  key: 'value',
  say: function() {
    console.log('never 996');
  },
};

console.log(obj.key); // 输出："value"
console.log(obj['key']); // 输出："value"

obj.say(); // 输出："never 996"
obj['say'](); // 输出："never 996"
```

这两种方式都很常用。第二种通常会应用在需要用变量确定属性名的时候去使用。

```javascript
var person = {
  age: 27,
  name: '鸽手',
};

Object.keys(person).forEach(function(key) {
  console.log(person[key]);
});
```

以上就是使用第二种方式的场景之一，使用 `Object.keys()` 获取到 person 的所有属性名组成的数组，对数组做遍历拿到每一个属性名并放在一个变量中，再通过变量访问到对应的属性值。

当试图访问一个不存在的属性的时候，则会返回 `undefined`。

```javascript
var obj = {};

console.log(obj.value); // 输出：undefined
```

### 2.2 设置对象的属性值

设置属性值也有 2 种方式：

1. `对象.属性名 = 属性值`
2. `对象['属性名'] = 属性值`

设置属性值的方式与访问值很相似，只是多了一个赋值操作。

设置属性值按照如下规则进行：

* 如果对象中不存在这个属性，则创建这个属性并赋值
* 如果对象中存在这个属性，则直接赋值

```javascript
var person = {
  age: 22,
};

person.name = '阿花';
person['hobby'] = ['eat', 'play doudou'];

console.log(person);

person.age = 33;

console.log(person);
```

## 3. 使用特殊的属性名

对象的属性名是可以用任意字符串表示的。

上面有提到，声明属性的时候可以带上引号。

如果不带引号，那属性名必须要符合变量命名的规则，使用引号包裹属性名，则可以使用任意字符串作为属性名。

```javascript
var obj = {
  --^$@@@age: 16,
};
```

上面这样写是会报错的，如果非要使用这样的属性名，就可以加上一对引号，可以双引号，也可以是单引号，使用引号的规则和字符串一致。

```javascript
var obj = {
  '--^$@@@age': 16,
};
```

这种特殊的属性名无法通过 `对象.属性名` 的形式访问。

```javascript
var obj = {
  '--^$@@@age': 16,
};

var val = obj.--^$@@@age;
```

JavaScript 无法解析这种特殊的语法，所以要使用 `对象['属性名']` 的形式访问。

```javascript
var obj = {
  '--^$@@@age': 16,
};

var val = obj['--^$@@@age'];

console.log(val); // 输出：16
```

特殊的属性名场景比较少，如统计字符串的场景。

```javascript
var counter = {};

var strs = [
  '#@T%TGFDSgfdsgsf',
  '#@T%TGFDSgfdsgsf',
  '123fdafeafewa',
  '123fdafeafewa',
  '#@T%TGFDSgfdsgsf',
];

strs.forEach(function(item) {
  if (item in counter) {
    counter[item]++;
  } else {
    counter[item] = 0;
  }
});

console.log(counter);
```

strs 是由字符串组成的数组，即需要统计的一组数据。

利用对象的特性来对字符串分类计数。

## 4. 其他创建对象的方法

除了字面量的方式，还有许多创建对象的方式。

### 4.1 使用 Object 对象

使用 `new Object()` 或者 `Object()` 的方式也可以创建一个对象

```javascript
var obj1 = new Object();
var obj2 = new Object; // 如果没有参数 可以不带括号

var obj3 = Object();
```

上面的方式都可以创建一个空对象。

比较有趣的是可以给 `Object` 传递一个对象字面量作为参数，返回的对象的属性与传入的对象字面量的属性一致。

```javascript
var obj1 = new Object({
  age: 11,
  name: '长睫毛',
});

var obj2 = Object({
  age: 12,
  name: '小酒窝',
});

console.log(obj1, obj2);
```

### 4.2 构造函数

使用构造函数，也可以创建对象。

```javascript
function Car(color, maxSpeed) {
  this.color = color;
  this.maxSpeed = maxSpeed;
}

Car.prototype.bibi = function() {
  console.log('哔哔！');
};

var car = new Car('red', 9999999);

console.log(car);
```

以上例子使用构造函数创建了一个`速度超级快的车`对象。

### 4.3 Object.create

使用 `Object.create` 也可以创建一个新对象，但是必须传递一个对象作为参数。

```javascript
var parent = {
  walk: function() {
    console.log('走路');
  },
};

var son = Object.create(parent);

console.log(parent === son);

son.walk();
```

`Object.create` 会根据传递过去的对象生成一个新的对象，作为参数传递的对象会作为新对象的原型。

## 5. 遍历对象

### 5.1 for … in

```javascript
var me = {height: 180, weight: 70};

var i;
for (i in me) {
  console.log(i);
  console.log(me[i]);
}
```

使用 `for ... in` 可以遍历对象的所有 key 值，也就是属性名，取到所有的属性就可以访问到所有的属性值。

需要注意的是 `for ... in` 循环只遍历可枚举属性，同时对象原型上的也会被访问到。

```javascript
var me = {height: 180, weight: 70};

var you = Object.create(me);

you.age = 11;

var i;
for (i in you) {
  console.log(i);
}
```

上面这个例子就把 `me` 和 `you` 中的所有属性都遍历出来了。

可以使用 `Object.prototype.hasOwnProperty` 来判断一个属性是否只处于其本身而不在原型上。

```javascript
var me = {height: 180, weight: 70};

var you = Object.create(me);

you.age = 11;

var i;
for (i in you) {
  if (you.hasOwnProperty(i)) {
    console.log(i);
  }
}
```

这样就只会输出 `age` 了。

### 5.2 Object.keys

`Object.keys` 会返回对象上的所有可枚举属性组成的数组。

```javascript
var gugugu = {
  name: '?王',
  hobby: '放鸽子',
};

var keys = Object.keys(gugugu);

console.log(keys);

keys.forEach(function(key) {
  console.log(gugugu[key])
});
```

通过遍历属性组成的数组来遍历对象。

### 5.3 Object.getOwnPropertyNames

使用 `Object.getOwnPropertyNames` 也可以获取到由属性组成的数组，但数组会包括不可枚举的属性。

```javascript
var gugugu = {
  name: '?王',
  hobby: '放鸽子',
};

var desc = Object.create(null);
desc.enumerable = false; // 是否可枚举 默认就是false
desc.value = '最强鸽手';

Object.defineProperty(gugugu, 'nickname', desc);

console.log(Object.keys(gugugu));
console.log(Object.getOwnPropertyNames(gugugu));
```

使用 `getOwnPropertyNames` 比使用 `keys` 多出一个不可枚举的 `nickname` 属性。

> 注意：ES6 还提供了`Object.values`、`Object.entries`、`for ... of`、`Reflect.ownKeys`等特性，结合这些特性也能遍历一个对象。

## 6. 创建绝对纯净的对象

> 纯净对象仅为本篇中的称呼方式，这种特殊的对象没有特定的称呼。

纯净对象即原型为 null 的对象。

使用 `Object.create(null)` 来创建纯净对象。

```javascript
var obj1 = Object.create(null);

console.log(obj1);

var obj2 = {};

console.log(obj2);
```

可以尝试在控制台中对比这两个对象，纯净对象是没有原型的，无法调用 `toString`、`hasOwnProperty`、`valueOf` 这些原型上的方法。

大部分使用纯净对象的场景是使用 `Object.defineProperty` 为对象创建属性的时候，属性的描述需要一个绝对干净的对象，防止特殊的属性对描述造成影响。

另外的使用场景就是当作一个字典使用，防止原型上的内容对字典产生干扰。

## 7. 小结

对象最常用的两种创建方式，就是使用字面量和构造函数。

创建对象的时候应合理规划属性名和方法名，根据业务来确定如何使用对象。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
