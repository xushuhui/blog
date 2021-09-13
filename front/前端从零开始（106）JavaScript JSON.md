# JavaScript JSON

> JSON 对象包含两个方法：用于解析 JavaScript Object Notation (JSON) 的 parse() 方法，以及将对象 / 值转换为 JSON 字符串的 stringify() 方法。除了这两个方法，JSON 这个对象本身并没有其他作用，也不能被调用或者作为构造函数调用。

JavaScript 内置的 `JSON对象` 用于处理 `JSON` 。

JSON（JavaScript Object Notation）是一种带有格式的文本，JavaScript 中的 `JSON对象` 用于处理这种文本。

JSON 对象只提供了两个方法，一个用于序列化 JSON ，一个用于反序列化 JSON 。

这里的序列化可以理解成`将JavaScirpt对象转换成JSON`，反序列化则是`将JSON转换成JavaScript对象`。

## 1. JSON.parse

`JSON.parse` 用于解析 JSON 格式的字符串，将 JSON 转化成 JavaScript 对象。

```javascript
JSON.parse(JSON字符串, 处理函数);
```

第一个参数是要转换成对象的 JSON 字符串，第二个参数可以不传递 /

```javascript
var str = '{ "name": "baba", "age": 12, "info": { "locate": "浙江" } }';

var user = JSON.parse(str);

console.log(user); // 输出一个 JavaScript 对象
```

传递给 JSON.parse 方法的字符串要符合 JSON 标准，否则会报错。

第二个参数非常有趣，传入的是一个函数，这个函数会在每个 JSON 属性被解析的时候调用，同时会传递属性名和属性值给函数作为参数，传入参数的返回值会作为当前遍历到的属性的新值。

```javascript
var str = '{ "name": "baba", "age": 12, "info": { "locate": "浙江" } }';

var user = JSON.parse(str, function(key, value) {
  console.log(key, value);

  return value;
});
```

可以发现上述例子打印的最后以想，属性名是可以空字符串，属性值是解析完的结果。

修改一下例子，将返回值改成一个固定的值。

```javascript
var str = '{ "name": "baba", "age": 12, "info": { "locate": "浙江" } }';

var user = JSON.parse(str, function(key, value) {
  console.log(key, value);

  return '强行修改值';
});
```

观察输出后，可以发现所有属性都被遍历了，并且赋值成功，但是最终 `user` 也变成了返回的字符串。

这是因为当解析完成后，传入的函数会被最后调用一次，传递进来的值就是最终 JSON.parse 的返回值，所以对其修改后，会影响到的最终结果。

```javascript
var str = '{ "name": "baba", "age": 12, "info": { "locate": "浙江" } }';

var user = JSON.parse(str, function(key, value) {
  console.log(key, value);

  if (key === '') {
    return value;
  }

  return '强行修改值';
});
```

对传递过来的属性名为空字符串 `''` 进行单独处理即可避免这种特殊情况。

> 业务逻辑中很少会用第二个参数来处理解析内容。

## 2. JSON.stringify

JSON.stringify 用于将 JavaScript 对象转换成 JSON 格式的字符串。

```javascript
JSON.stringify(JavaScript对象, 处理函数, 缩进空格字符数量);
```

第一个参数是需要转换成 JSON 字符串的对象。

第二个参数可以是个函数，也可以是个数组。

如果是函数，则每一个属性在处理的时候就被调用这个函数，同时属性名和属性值作为参数传递给这个函数，并且函数的返回值作为这个处理属性的值。

如果是数组，则只有属性名在数组中的属性才会被处理，不传递则默认处理整个对象。

如果第二个参数传递 null ，也就是不做特殊处理，在使用到第三个参数的时候，第二个参数会传递 null 。

第三个参数可以传递数字，也可以传递字符串，传递了这个参数会对结果做格式化，具有一定的格式，参数的值决定格式化的样式。

如果是数字，则使用对应长度的空格来缩进，长度 1 到 10 ，比 1 小则表示不缩进。

如果是字符串，则会使用传入的字符串进行缩进，传入的字符串长度超过 10 ，则会截取前 10 个作为缩进字符。

```javascript
var user = {
  name: '小明',
  age: 14,
  skill: ['HTML', 'Java'],
};

var json = JSON.stringify(user);

console.log(json);
// 输出：{"name":"小明","age":14,"skill":["HTML","Java"]}
```

第二个参数用起来和 parse 方法的第二个参数类似。

```javascript
var user = {
  name: '小明',
  age: 14,
  skill: ['HTML', 'Java'],
};

var json = JSON.stringify(user, function(key, value) {
  console.log(key, vlue);

  return value;
});

console.log(json);
```

根据上述例子可以看到，先输出的属性为空字符串，属性值为被处理对象，所以如果不想操作原对象，需要做特殊处理。

```javascript
var user = {
  name: '小明',
  age: 14,
  skill: ['HTML', 'Java'],
};

var json = JSON.stringify(user, function(key, value) {
  if (key === '') {
    return value;
  }

  return '我是处理过的值';
});

console.log(json);
```

这样处理后，最终处理完的 JSON 字符串的属性值都是函数的返回值了。

第三个参数会在做一些工具类调试的时候常用到。

```javascript
var obj = [
  {
    path: '/',
    component: 'function() {}',
    children: [
      {
        path: 'note',
        component: 'function() {}',
      },
      {
        path: 'friends',
        component: 'function() {}',
      }
    ]
  },
  {
    path: '*',
    component: 'function() {}',
  }
];

var json1 = JSON.stringify(obj, null);
var json2 = JSON.stringify(obj, null, 2);
var json3 = JSON.stringify(obj, null, '*-*');

console.log(json1); // 没有格式
console.log(json2); // 使用两个空格控制的缩进
console.log(json3); // 使用 *-* 控制的缩进
```

传入参数后就会将处理后的 JSON 字符串进行格式化，缩进部分根据传入的参数值决定。

![图片描述](https://img.mukewang.com/wiki/5e7a44c20a01631711460686.jpg)

## 3. 其他注意点

### 3.1 深拷贝

可以配合 JSON 的两个方法，对对象进行深拷贝。

```javascript
var obj = {prop: 'value'};

var newObj = JSON.parse(JSON.stringify(obj));

newObj.prop = 'new value';

console.log(obj);
console.log(newObj);
```

根据结果可以看到新的对象修改，没有影响到原对象，两者之间不存在引用关系。

### 3.2 序列化规则

使用 JSON.stringify 有些内置规则。

* 如果对象中存在包装对象，则在转换过程中会变成原始值。

```javascript
var obj = {
  string: new String('A promise is a promise.'),
  number: new Number(996),
};

var result = JSON.stringify(obj);

console.log(result); // 输出："{"string":"A promise is a promise.","number":996}"
```

* 如果转换的对象或者对象下的属性存在 toJSON 方法，那么这个方法的返回值会作为转换结果。

```javascript
var user = {
  nickname: 'joker',

  toJSON: function() {
    return 'hahahahahahaha';
  },
}

var result = JSON.stringify(user);

console.log(result); // 输出："hahahahahahaha"
```

可以看到结果为 toJSON 方法的返回值。

* 除了数组以外的对象，转换结果顺序为随机。

```javascript
var obj = {
  b: 2,
  c: 3,
  a: 1,
};
```

如以上对象，转换的结果有可能是以下情况中的一种：

```javascript
"{"a":1,"b":2,"c":3}"
"{"a":1,"c":3,"b":2}"
"{"b":2,"a":1,"c":3}"
"{"b":2,"c":3,"a":1}"
"{"c":3,"b":2,"a":1}"
"{"c":3,"a":1,"b":2}"
```

* undefined、ES6 中的 symbol 值、函数在转换过程中都会被忽略，当然函数如果具有 toJSON 方法依然会优先选择 toJSON 方法的结果。

```javascript
var fn = function() {};
fn.toJSON = function() {return '我是函数'};

var result = JSON.stringify({
	a: fn,
	b: Symbol(1),
	c: undefined,
  d: function() {},
});

console.log(result);
```

* 存在循环引用，则会报错

```javascript
var obj1 = {
  prop1: 1,
};
var obj2 = {
  prop1: 1,
};

obj1.prop2 = obj2;
obj2.prop2 = obj1;

JSON.stringify(obj1); // TypeError: Converting circular structure to JSON
```

两个对象相互引用之后，进行系列化就会抛出错误。

* 在 ES6 中，symbol 可以作为对象的属性值，但在处理的时候都会被忽略。

```javascript
var symbol = Symbol();

var obj = {
  prop1: 'value1',
  [symbol]: 'value2',
};

console.log(obj);

var result = JSON.stringify(obj);

console.log(result); // 输出：{"prop1":"value1"}
```

* null、正负 Infinity、NaN 在序列化时都会被当作 null 。

```javascript
var obj = {
  null: null,
  infinity1: +Infinity,
  infinity2: -Infinity,
  NaN: NaN,
};

var result = JSON.stringify(obj);

console.log(result); // 输出：{"null":null,"infinity1":null,"infinity2":null,"NaN":null}
```

## 4. 小结

JSON 几乎是目前前后端交互最常用的数据格式，所以 JSON 对象使用的频率也很高。

在使用 `JSON.parse` 反序列化的时候，如果 JSON 格式不符合规范，是会报错的，日常开发中建议封装一层 JSON 的方法，将错误集中处理，方便定位与上报错误。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
