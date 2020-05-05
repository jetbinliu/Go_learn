package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接失败,err:", err)
		return
	}
	defer socket.Close()
	//发送数据
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("连接失败,err:", err)
			return
		}
		fmt.Println("收到护肤信息", string(reply[:n]))
	}
}
