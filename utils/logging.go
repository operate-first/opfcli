package utils

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ConfigureLogging() {
	loglevel, err := strconv.Atoi(os.Getenv("OPF_LOGLEVEL"))
	if err != nil {
		loglevel = 1
	}

	switch {
	case loglevel >= 2:
		log.SetLevel(log.DebugLevel)
	case loglevel >= 1:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}
