package task

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

//限速
const maxLimit = 1<<63 - 1

var limit uint64 = maxLimit

//kb
const kb = 1024

func httpLimit(url, i string) {
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	var buff = make([]byte, kb*10)
	var temp uint64
	var f, _ = os.OpenFile(i+"qq.exe", os.O_CREATE|os.O_RDONLY|os.O_TRUNC, 0655)
	t := time.Now().Unix()
	var now = t
	for true {
		n, err := data.Body.Read(buff)
		if err != nil {
			f.Write(buff[:n])
			data.Body.Close()
			break
		}
		f.Write(buff[:n])
		temp += uint64(n)
		if temp >= limit {
			if (time.Now().Unix() - t) <= 1 {
				time.Sleep(time.Second)
				t = time.Now().Unix()
			}
			temp = 0
		}
	}
	f.Close()
	fmt.Println(time.Now().Unix()-now, limit)
}
