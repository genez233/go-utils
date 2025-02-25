package dc

import (
	"fmt"
	"net"
)

// GetIP 获取设备IP地址
func GetIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("get ip", err)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println("get ip", ipAddress)
	return ipAddress.IP.String()
}
