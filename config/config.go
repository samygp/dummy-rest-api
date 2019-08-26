package config

import (
	"github.com/jinzhu/configor"
)

// Config represents the configuration of the entire application.
var Config = struct {
	App struct {
		Name string `json:"name" default:"Image monitor"`
	} `json:"app"`
	Logger struct {
		Level string `json:"level" default:"debug"`
	} `json:"logger"`
	Server struct {
		MaxRequests  int    `json:"maxrequests" default:"0"`
		Port         string `json:"port" default:"32123"`
		ReadTimeout  int64  `json:"readtimeout" default:"30"`
		WriteTimeout int64  `json:"writetimeout" default:"30"`
	} `json:"server"`
}{}

// Init config
func Init() {
	if err := configor.New(&configor.Config{ENVPrefix: "-"}).Load(&Config); err != nil {
		panic(err)
	}
}
