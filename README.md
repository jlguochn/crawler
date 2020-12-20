# 91Porn爬虫

使用golang实现的91porn爬虫

运行方式：

```go
go run main.go
```

需要添加能科学上网的HTTP代理，例如 /fetcher/fercher.go 中：

```go
resp,err := HttpGetFromProxy(url,"//127.0.0.1:1087")//替换成你自己的HTTP代理
```

爬取到的视频将自动下载到download目录中。