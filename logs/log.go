package logs

import (
	"errors"
	"log"
	"os"
)

var Logger *log.Logger

func MyLogger(err error) {
	var log_path string = "error.log"
	error_file, err2 := os.OpenFile(log_path, os.O_WRONLY|os.O_APPEND, 0664)
	if err2 != nil {
		log.Print("Error", err2)
	}
	Logger = log.New(error_file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	err := errors.New("Server Started")
	MyLogger(err)
}