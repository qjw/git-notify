package utils

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"github.com/codegangsta/cli"
)

func InitLog(c *cli.Context){
	log.SetOutput(os.Stdout)
	if c.GlobalBool("debug") {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)
	}
}
