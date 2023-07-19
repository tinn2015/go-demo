# go-tool-trace

## 不合适

运行缓慢的函数，或者找到大部分CPU时间花费在哪里，术业有专攻，看CPU时间花费，是有专门的工具的 go tool pprof

## 合适

找出程序在一段时间内正在做什么
go tool trace 可以通过 view trace链接提供的其他可视化功能，对于诊断争用问题帮助极大

# 使用方式
```
// 先执行
go run ./main.go

// 查看分析， 这时会打开一个网页
go tool trace myTrace.dat 
```