---
title: Go 1.13版本 xerrors 包装错误
published: 2021-01-03 13:59:58
tags: ["Go"]
categories: ["Go"]
---

> 翻译自 https://crawshaw.io/blog/xerrors

# Go 1.13: xerrors

Go 2 系列语言更改的一部分是新的[错误检查提案](https://go.googlesource.com/proposal/+/master/design/29934-error-values.md)。


错误检查提案为其他地方（在 github.com/pkg/errors 等包中）尝试过的错误添加了几个功能，并带有一些新的实现技巧。 该提案已在提示中实施，为 Go 1.13 做准备。 您今天可以通过使用 Go from tip 或使用包 golang.org/x/xerrors 和 Go 1.12 来尝试一下。


额外的功能完全基于库，不涉及对编译器或运行时的更改。 一个重要的新功能是错误包装。


## 一个工作示例：包装“key not found” 错误

我们正在为 [Tailscale](https://tailscale.io/)  构建的产品包括一个名为 taildb 的简单键值存储。 与许多简单的 KV 存储一样，您可以读取键值。 



```go
// Get fetches and unmarshals the JSON blob for the key k into v.
// If the key is not found, Get reports a "key not found" error.
func (tx *Tx) Get(k string, v interface{}) (err error)
```

让我们来谈谈 "key not found."

### 版本 1


第一个 API 版本将"key not found"错误定义为：

```go
var ErrNotFound = errors.New("taildb: key not found")
```

使用taildb的代码可以轻松使用:

```go
var val Value
if err := tx.Get("my-key", &val); err == taildb.ErrNotFound {
	// no such key
} else if err != nil {
	// something went very wrong
} else {
	// use val
}
```

这很好，直到我进行一些调试并遇到一个归结为:

```go
my_http_handler: taildb: key not found
```
这不是一个非常有用的错误消息.

### 版本 2

鉴于 `Get` 方法具有键名，最好将其包含在错误消息中。

所以我遵循了 Go 中的一个常见策略，即在 taildb 包中引入错误类型：

```go
type KeyNotFoundError struct {
	Name string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Errorf("taildb: key %q not found")
}
```

这很好用！检查此特定错误的代码有点混乱，但它可以工作:

```go
var val Value
err := tx.Get("my-key", &val)
if err != nil {
	if _, isNotFound := err.(taildb.KeyNotFoundError); isNotFound {
		// no such key
	} else {
		// something went very wrong
	}
} else {
	// use val
}
```
但这种直接搭配的风格有一个缺陷。如果任何中间代码将信息添加到错误中，我们将无法再检查错误的类型。考虑如下函数:

```go
func accessCheck(tx *taildb.Tx, key string) error {
	var val Value
	if err := tx.Get(key, &val); err != nil {
		return fmt.Errorf("access check: %v", err)
	}
	if !val.AccessGranted {
		return errAccessDenied
	}
	return nil
}
```

在这里，我们在数据库之上实现逻辑，检查用户是否具有某种访问权限。报告 nil 错误将授予访问权限，否则访问将被拒绝。拒绝访问的原因可能是 `!AccessGranted` 或一些底层数据库错误。所有关于错误的文本信息都被保留了，但是使用 `fmt.Errorf` 意味着我们不能再检查访问错误是否是 `KeyNotFoundError`。

### 版本 3

新的 xerrors 库通过提供一个版本的 Errorf 来解决此问题，该版本保留了新错误中的底层错误对象:

```go
	if err := tx.Get(key, &val); err != nil {
		return xerrors.Errorf("access check: %w", err)
	}
```

%w for wrap.

从表面上看，Errorf 的这种实现与 fmt 中的实现完全一样。在底层，保留类型意味着我们现在可以检查 KeyNotFoundError 的原因链:

```go
var val Value
if err := accessCheck(tx, "my-key"); err != nil {
	var notFoundErr taildb.KeyNotFoundError
	if xerrors.As(err, &notFoundErr) {
		// no such key
	} else {
		// something went very wrong
	}
} else {
	// use val
}
```

Great!

### 版本 4

我们可以做得更好。我们替换导出的 KeyNotFoundError 的唯一原因是我们可以在错误消息中添加一些额外的文本，同时使类型可测试。新的 xerrors 为我们提供了一种更简单的方法来做到这一点。

所以让我们回到第一个定义:

```go
var ErrNotFound = errors.New("key not found")
```

在taildb里面我们可以写:

```go
func (tx *Tx) Get(k string, v interface{}) (err error) {
	// ...
	if noSuchKey {
		return xerrors.Errorf("taildb: %q: %w", k, ErrNotFound)
	}
}
```

我们想要的所有信息都在这里。当我们将错误打印到日志时，我们会看到 `taildb: "my-key": key not found`。要检查从 `accessCheck` 返回的错误，我们可以编写:

```go
var val Value
if err := accessCheck(tx, "my-key"); xerrors.Is(err, taildb.ErrNotFound) {
	// no such key
} else if err != nil {
	// something went very wrong
} else {
	// use val
}
```

简单！

## Go 1.13

新的 xerrors 将在 Go 1.13 中升级到标准库的错误包中。

链接不是 xerrors.Errorf，而是直接构建到我们今天使用的 [fmt.Errorf](https://tip.golang.org/pkg/fmt/#Errorf) 函数中：

如果最后一个参数是错误的并且格式字符串以“: %w”结尾，
返回的错误实现 errors.Wrapper 并带有返回它的 Unwrap 方法。

当然，这看起来不错。然而，距离 Go 1.13 仅三个月之遥！在那之后，所有这些新的变化（这篇文章只介绍一个）将在[Go 1 兼容性承诺]（https://golang.org/doc/go1compat）下的标准库中被永久冻结。对于如此高的标准，这个包[可悲地测试不足]（https://godoc.org/golang.org/x/xerrors?importers）。

我鼓励你从今天开始使用 golang.org/x/xerrors，或者更好的是，通过 [从源代码安装](https://golang.org/doc/install/source) 直接针对 Go 提示开始开发。更多的人需要尝试一下。 