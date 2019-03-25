//验证io.EOF
//golang.conn.read 2>log.log
//发现log.log的长度为
//与buf的长度不一致
//通过指定下标查看起始字符，发现原来是log的时间写入了文件导致的
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		//fmt.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)
		fmt.Println(n, len(buf))

	}
	fmt.Printf("totol bytes:%d total rune:%d [0]:%c [%d]:%d", len(buf), len([]rune(string(buf))), buf[0], len(buf)-1, buf[len(buf)-1])
	log.Println(string(buf))
}
