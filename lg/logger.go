package lg

import (
	"log"
	"os"
)

type debugLevel uint

const (
	NOLOG         debugLevel = 0
	DEBUGGING     debugLevel = 15
	NORMAL        debugLevel = 7
	ONLY_PROBLEMS debugLevel = 3
	ALL           debugLevel = 31
	VERBOSE       debugLevel = 16
	DEBUG         debugLevel = 8
	INFO          debugLevel = 4
	WARN          debugLevel = 2
	ERROR         debugLevel = 1
)

var (
	Verbose *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warn    *log.Logger
	Error   *log.Logger
)

func Init(Level debugLevel, showFileName bool) (err error) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		return
	}
	defer devNull.Close()

	flag := log.Ldate | log.Ltime
	if showFileName {
		flag |= log.Lshortfile
	}

	if Level&VERBOSE > 0 {
		println("verbose enabled.")
		Verbose = log.New(os.Stdout, "[VRBSE] ", flag)
	} else {
		Verbose = log.New(devNull, "[VRBSE] ", 0)
	}

	if Level&DEBUG > 0 {
		Debug = log.New(os.Stdout, "[DEBUG] ", flag)
	} else {
		Debug = log.New(devNull, "[DEBUG] ", 0)
	}

	if Level&INFO > 0 {
		Info = log.New(os.Stdout, "[INFO ] ", flag)
	} else {
		Info = log.New(devNull, "[INFO ] ", 0)
	}

	if Level&WARN > 0 {
		Warn = log.New(os.Stdout, "[WARN ] ", flag)
	} else {
		Warn = log.New(devNull, "[WARN ] ", 0)
	}

	if Level&ERROR > 0 {
		Error = log.New(os.Stdout, "[ERROR] ", flag)
	} else {
		Error = log.New(devNull, "[ERROR] ", 0)
	}

	return
}
