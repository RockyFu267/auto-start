package basefunc

import (
	"log"
	"net"
	"time"
)

// func CheckPort(input int64) (bool, error) {
// 	return true, nil
// }

//TcpGather net包
func TcpGather(ports string) (bool, error) {
	address := net.JoinHostPort("127.0.0.1", ports)
	// 3 秒超时
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		log.Println(ports+" Connet Refused", err)
		return false, err
	} else {
		if conn != nil {
			log.Println(ports + " Connet Success")
			return true, nil
		} else {
			log.Println(ports+" Connet Refused", err)
			return false, err
		}
	}
}
