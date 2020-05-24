package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	clientMap = make(map[net.Addr]net.Conn)
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		for _, conn := range clientMap {
			conn.Write([]byte("服务器进入维护状态，准备关闭！"))
		}
		os.Exit(1)
	}

}
func ioWithConn(conn net.Conn) {
	client_add := conn.RemoteAddr().String()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	HandleError(err, "conn.Read")
	if n > 0 {
		msg := string(buffer[:0])
		fmt.Println("%s:%s", client_add, msg)
		strs := strings.Split(msg, "#")
		targetAddr := strs[0]
		targMsg := strs[1]
		if targetAddr == "all" {
			for _, conn := range clientMap {
				conn.Write([]byte(targMsg))
			}
		}else{
			//点对点消息
			range clientMap
			}
		}
	}
	// 群发消息

}
func main() {
	// 建立服务端
	lister, err := net.Listen("tcp", "127.0.0.1:8888")
	HandleError(err, "net.Listen")
	//在单独的聊天室和每个客户端聊天
	for {
		//循环接入所有客户端
		conn, err := lister.Accept()
		HandleError(err, "lister.Accept")
		go ioWithConn(conn)
	}

}
