package log

import (
	"io"
	"log"
	"os"
)

func InitLog() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()

	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiWriter)
}
