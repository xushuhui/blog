---
title: PHP 7 新特性
published: 2016-04-03 09:03:42
tags: [PHP]
categories: ["PHP"]
top: 0
---

## 1. 太空船运算符

``` php
echo 1<=>1; //0
echo 2<=>1; //1
echo 1<=>2; //-1
```

## 2. 类型声明

``` php
declare(strict_type=1)//strict_type=1表示开启严格模式
function sumOfInts(int ...$ints):int{
    return array_sum($ints);
}
```

## 3.null 合并操作符

``` php
$page = isset($_GET['page']) ? $_GET['page'] : 0;
$page = $_GET['page'] ?? 0;
```

## 4. 常量数组

``` php
define('ANIMALS',['dog','cat']);
```

## 5.namespace 批量导入

``` php
use Space\{ClassA,ClassB as B,ClassC}
```

## 6.intdiv 函数

``` php
intdiv(10,3);
```

## 7.list 方括号

``` php
$arr = [1,2,3];
list($a,$b,$c) = $arr;

$arr = [1,2,3];
[$a,$b,$c] = $arr;
```
