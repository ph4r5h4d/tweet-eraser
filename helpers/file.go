package helpers

import (
	"log"
	"os"
)

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
