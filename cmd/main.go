package main

import (
	"encoding/json"
	"fmt"

	"github.com/samygp/dummy-rest-api/app/server"
	"github.com/samygp/dummy-rest-api/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Debugf("this")
	config.Init()
	var err error
	buffer, err := json.Marshal(config.Config)
	if err != nil {
		fmt.Printf("Unable to marshal config: %v\n", err)
	} else {
		log.Infof("Current config: %s\n", string(buffer))
	}
	if config.Config.Logger.Level == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	server.Start()
}
