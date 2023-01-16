package basefunc

import (
	bc "auto-start/basecmd"
	"log"
	"net"
	"time"
)

func CheckPort(input string) (bool, error) {
	cmdRes, err := bc.CmdAndChangeDirToResAllInOne("./", "ss -lntp | grep "+input)
	if err != nil {
		log.Println("check port "+input+" error:", err)
		return false, err
	}
	if len(cmdRes) == 0 {
		return false, nil
	}
	return true, nil
}

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
			//log.Println(ports + " Connet Success")
			return true, nil
		} else {
			log.Println(ports+" Connet Refused", err)
			return false, err
		}
	}
}
