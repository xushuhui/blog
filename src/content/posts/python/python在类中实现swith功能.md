---
title: Python 在类中实现 swith 功能
published: 2019-03-29 14:51:12
tags: [Python]
categories: ["Python"]
top: 0
---

## 问题

Python 中没有 switch 的语法，但是很多时候需要多重条件判断，又不想写多个 if，那只能手动实现了。
实现代码

```python
class RunMethod:
    def post(self,url=None,data=None,header=None):
        print(url)
    def get(self,url=None,data=None,header=None):
        print("get")

    def main(self,method):
        method = getattr(self, method)
        return method

if __name__ == '__main__':
    client = RunMethod()
    client.main("post")("http://www.baidu.com")
```

其中主要用到 getattr 这个函数，用于返回一个对象属性值。

```python
getattr(object, name[, default])
```

- object -- 对象。
- name -- 字符串，对象属性。
- default -- 默认返回值，如果不提供该参数，在没有对应属性时，将触发 AttributeError。
