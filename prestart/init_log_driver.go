package prestart

import (
	"github.com/hitokoto-osc/Vive/config"
	log "github.com/sirupsen/logrus"
	"os"
)

// The Log driver is served by logrus
func initLogDriver() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999Z",
	})
	if config.Debug {
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.ErrorLevel)
	}
}
