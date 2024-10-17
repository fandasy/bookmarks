package l

import (
	"time"
	"os"
	"strings"
	"log"
	"io"
)

func Start() {
	logDir := "logs"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
        os.Mkdir(logDir, 0774)
    }

    nowDate := time.Now().Format(time.DateOnly)
	nowTime := strings.ReplaceAll(time.Now().Format(time.TimeOnly), ":", ".")

	file, err := os.Create(logDir + "/" + nowDate + "_" + nowTime + ".txt")
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
}