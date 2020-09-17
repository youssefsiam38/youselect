package framework

import (
	"fmt"
	logPkg "log"
	"os"
	// "time"
	"io"
)

type writer io.Writer

var (
	// Log is customized log
	Log func(v ...interface{})
	w   writer
)

func init() {
	tofile := false

	if tofile {
		file, err := os.OpenFile("logs", os.O_WRONLY, os.ModePerm)

		if err != nil {
			fmt.Println("here")
			file, _ = os.Create("logs")
		}
		w = file
	} else {
		w = os.Stdout
	}

	logger := logPkg.New(w, ``, logPkg.Lshortfile|logPkg.LstdFlags)

	Log = logger.Println

}
