package main

import (
	"logger/log"
)

func main() {

	logger := log.GetLogger()
	logger.Info("Info")
	// logger.Fatal(errors.New("Error"))
	logger.Debug("Debug")
	// logger.Error("Error")
}
