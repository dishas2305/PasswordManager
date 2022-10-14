package storage

import (
	"io"
	"os"

	"passmanager/config"

	log "github.com/sirupsen/logrus"
)

func ConnectLogrus() {
	if envName := os.Getenv("ENV"); envName != config.Dev && envName != config.Qa && envName != config.Prod {
		log.Println("Log file will be available on terminal console")
	}

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
	var filename string = "logs/logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file. error: %v", err)
	}
	//Set the log format to JSON format
	Formatter := new(log.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	// Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	SetLoggingLevel("info")
	log.SetOutput(io.MultiWriter(f, os.Stdout))
}

func SetLoggingLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Println("No log level found")
	}
}
