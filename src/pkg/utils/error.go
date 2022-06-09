package utils

import (
	"bytes"
	"fmt"
	"log"
)

func Error(err error) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "ERROR: ", log.LstdFlags|log.Llongfile)
		infof  = func(info string) {
			_ = logger.Output(3, info)
		}
	)

	infof(err.Error())
	fmt.Print(&buf)
}
