package utils

import (
	"os"
)

func WriteLog(str string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open/create file\n err: " + err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(str + "\n")
	if err != nil {
		panic("failed to write into file\n err: " + err.Error())
	}
}
