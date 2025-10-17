# 提问

## 如何定义嵌套结构体？

```go
type base struct {
    num int
}

type container struct {
    base
    s string
}
```
