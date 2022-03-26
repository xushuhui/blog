# delete 操作符

> delete 操作符用于删除对象的某个属性；如果没有指向这个属性的引用，那它最终会被释放。(MDN)

delete 操作符可以删除对象的一个属性。

JavaScript 中的关键字与其他语言略有不同，如 C++ 中的 `delete` 关键字会释放内存，JavaScript 中不会，只有当一个值的引用归零时，才会被释放。

## 1. 使用 delete

delete 操作符在与操作数运算结束后，会返回一个布尔值，成功返回 true。在属性是不可配置的情况下会返回 false，在严格模式下，则会抛出 `TypeError` 异常。

```javascript
var person = {
  age: 16,
};

delete person.age;

console.log(person); // 输出：{}
```

当一个属性为不可配置的时候：

```javascript
var person = {};

Object.defineProperty(person, 'age', {
  value: 17,
  writable: true,
  configurable: false,
});

delete person.age; // 返回false

console.log(person); // 输出：{age: 17}
```

## 2. 在严格模式下的 delete

如果在严格模式下，对一个不可配置属性进行 delete 操作，则会抛出异常。

```javascript
'use strict'; // 开启严格模式
var person = {};

Object.defineProperty(person, 'age', {
  value: 17,
  writable: true,
  configurable: false,
});

delete person.age; // TypeError: Cannot delete property 'age'
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ebd8325096b5f2812420372.jpg)

## 3. 对使用 var 声明的变量进行 delete

使用 var 声明的变量默认是不可配置的，所以对 var 声明的变量进行 `delete` 操作是无效的。

```javascript
var number = 996;

delete number; // false

console.log(number);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ebd82ec093902a809460186.jpg)

这里的 `number` 是 `window`下的一个属性，可以使用 `Object.getOwnPropertyDescriptor` 来查看属性的描述符。

```javascript
var number = 996;

Object.getOwnPropertyDescriptor(window, 'number');
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ebd82f809ca1f1510880240.jpg)

## 4. 使用 delete 删除数组成员

delete 可以用于删除数组成员，并且是真正意义的删除，可以让指定的成员变成 `empty` 。

```javascript
var arr = [2, 4, 6, 8, 10];

delete arr[0]; // true

console.log(arr);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ebd83010921f72608460194.jpg)

> 有关数组 empty 相关的内容可以查阅数组章节。

## 5. 小结

`delete` 操作符就是用来删除对象下的属性，但这个属性还有在其他地方被引用，就不会被释放。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
