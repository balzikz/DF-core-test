package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	conf, err := server.DefaultConfig().Load()
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.New(&conf, log)
	srv.Start()

	for srv.Accept(nil) {
	}
}
