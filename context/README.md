# 提问

## 在此练习中，提前结束的客户端无法收到响应，还有必要保留下面的代码吗？

```go
internalError := http.StatusInternalServerError
http.Error(w, err.Error(), internalError)
```
