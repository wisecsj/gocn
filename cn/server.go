// 很意外，golang标准库居然实现了这么一个文件服务器的Handler
// 其实自己实现一个简易的也并不难，也就是实现一个Handler，然后
// 把对应的文件字节流写进socket.不过标准库的FileServer还实现
// 了Range（从而支持客户端进行并行下载和断点续传）。并且访问跟目录
// 的时候，还会很方便地把文件列表用a标签展示出来
package main


import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/home/wisej/桌面"))))
}