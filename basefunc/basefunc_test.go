package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_CheckDir(t *testing.T) {
	res, err := CheckDir("/Users/fuao/Downloads/付傲.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_TcpGather(t *testing.T) {
	res, err := TcpGather("1087")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_RunStart(t *testing.T) {
	res, err := RunStart("/Users/fuao/Desktop/code/github/auto-start/test.sh")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}
