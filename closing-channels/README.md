# 提问

## 读取被关闭的 channel 会出现问题吗？

不会，读取的数据是 channel 的零值

## 写入数据到被关闭的 channel 会出现什么问题？

出现 panic
