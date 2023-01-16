package basefunc

import (
	bc "auto-start/basecmd"
	"fmt"
	"log"
)

//RunStart
func RunStart(input string) (bool, error) {
	//检查目录是否存在
	checkDir, err := CheckDir(input)
	if err != nil {
		log.Println("Check sh Error: ", err)
		//***结束***
		return false, err
	}
	if checkDir {
		cmdRes, err := bc.CmdAndChangeDirToResAllInOne("./", "sh "+input)
		if err != nil {
			log.Println("start "+input+" error:", err)
			return false, err
		} else {
			fmt.Println(cmdRes)
			return true, nil
		}
	} else {
		log.Println("bash not found")
		return false, err
	}

}
