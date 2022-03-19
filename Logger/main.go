package main

import (
	"logger/log"
)

type logData struct {
	info string      `json:"info"`
	data interface{} `json:"data"`
}
type PostgresInfo struct {
	user     string `json:"user"`
	password string `json:"password"`
}

func main() {

	logger := log.GetLogger()
	logger.Trace("Something very low level.")
	logger.Debug("Useful debugging information.")
	logger.Info("Something noteworthy happened!")
	logger.Warn("You should probably take a look at this.")
	logger.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after loggerging
	logger.Fatal("Bye.")
	// Calls panic() after loggerging
	logger.Panic("I'm bailing.")

	// logdata := logData{info: "psql data", data: PostgresInfo{user: "mrb", password: "0224"}}

	// logger.GetLoggerWithField("INFO PSQL DATA", logdata).Info("INFO PSQL")
}
