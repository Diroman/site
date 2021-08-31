package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
	 Log.Out = file
	} else {
	 Log.Info("Failed to log to file, using default stderr")
	}

	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
}