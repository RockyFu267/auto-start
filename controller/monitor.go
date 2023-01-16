package controller

import (
	bf "auto-start/basefunc"
	"log"
	"time"
)

func StartMonitoring(port string, startPath string, intervalTime int64) {
	for {
		time.Sleep(time.Duration(intervalTime) * time.Second)
		resCon, err := bf.CheckPort(port)
		if err != nil {
			log.Println(err)
			continue
		} else {
			if resCon {
				continue
			} else {
				resStr, err := bf.RunStart(startPath)
				if err != nil {
					log.Println(err)
					continue
				}
				if resStr {
					continue
				} else {
					log.Println(port + " start failed")
					continue
				}
			}
		}
	}
}
