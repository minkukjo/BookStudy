package config

import (
	"flag"
)

var Config *string
var Url string

func init() {
	Config = flag.String("config", "prod", "phase config")
	flag.Parse()

	if *Config == "local" {
		Url = "localhost"
	} else {
		Url = "34.67.130.46"
	}
}
