# merrors

将fmt.Errorf输出为调用栈的字符串

## 安装

```shell
  go get -u github.com/kevin2027/merrors
```

## 引入代码

```go
import (
    "github.com/kevin2027/merrors"
)
```

## 使用

```go
err:=MyFunction(arg1, arg2)
fmt.Printf("%s\n", merrors.String(err))
```

## 打印结果

```text
new error: error call
/Users/dz0401008/projects/merrors/merrors_test.go:func1:14: new error: error call
/Users/dz0401008/projects/merrors/merrors_test.go:TestString:21: error call
```
