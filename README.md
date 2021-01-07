# 91爬虫
91为国外知名视频网站
使用golang实现的91爬虫

运行方式：

```go
go run main.go
```

需要添加能科学上网的HTTP代理，/fetcher/fercher.go 中：

```go
resp,err := HttpGetFromProxy(url,"//127.0.0.1:1087")//替换成你自己的HTTP代理
```

爬取到的视频将自动下载到/download/目录中（需手动在项目根目录下创建/download/目录）。



##免责声明
本程序仅供学习和技术交流。