package main

import (
	"github.com/xmge/seelog"
	"os"
	"time"
	"fmt"
	"log"
)

const (
	TestLog  = "test.log"	// 项目日志位置
	TestPort = 9999		// 浏览器查看日志的端口
)

// e1：监听单个日志文件
func main()  {
	seelog.See(TestLog, TestPort)	// 在程序开始时 开启 seelog 即可
	yourProject()			// 模拟your项目
}

func yourProject()  {
	for {
		f, err := os.OpenFile(TestLog, os.O_RDWR|os.O_CREATE, 0766);	if err != nil {
			log.Panic(err)
		}
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			testLog := fmt.Sprintf("「模拟日志」[%s] 第[%d]行日志\n", time.Now().String(),i)
			_, err := f.WriteString(testLog)
			if err != nil {
				log.Panic(err)
			}
		}
		if err := f.Close();err != nil {
			log.Panic(err)
		}
		if err := os.Remove(TestLog);err != nil {
			log.Panic(err)
		}
	}

}
