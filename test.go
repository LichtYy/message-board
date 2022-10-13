// package main

// import (
// 	"fmt"
// 	"net"
// )

//	func main() {
//		//监听端口
//		listen, _ := net.Listen("tcp", `:86`)
//		//等待连接连接并建立连接
//		conn, _ := listen.Accept()
//		//http半双工需要先读取里面的数据
//		buf := make([]byte, 1024)
//		l, _ := conn.Read(buf)
//		fmt.Println(string(buf[:l]))
//		//通过连接发送数据
//		//http 协议格式
//		// conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain;charset=UTF-8\r\n\r\n数据开始：net模拟http"))
//		conn.Write([]byte("HTTP/1.1 200 OK Content-Type: text/plain;charset=UTF-8\r\n\r\n<div style=\"color: red;\">hh</div>"))
//		conn.Write([]byte("HTTP/1.1 200 OK Content-Type: text/plain;charset=UTF-8\r\n\r\n<div style=\"color: green;\">huhu</div>"))
//		conn.Close()
//	}
package main

import (
	"net"
)

func main() {
	//监听端口
	listen, _ := net.Listen("tcp", `:88`)
	for {
		//等待连接连接并建立连接
		conn, _ := listen.Accept()
		//通过连接发送数据
		conn.Write([]byte(`hello 我是服务端`))
	}
}
