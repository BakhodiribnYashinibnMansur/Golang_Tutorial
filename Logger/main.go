package main

import (
	"logger/log"
)

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
}
