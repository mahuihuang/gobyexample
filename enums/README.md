# 提问

## 如何定义枚举类型？

```go
const (
    Queen = iota 
    King
)
// 
type Poker int
const (
    Queen Poker = iota 
)
```
