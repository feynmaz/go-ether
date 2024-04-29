package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      int    `envconfig:"PORT" default:"8080"`
	InfuraURL string `envconfig:"INFURA_URL" required:"true"`
}

var (
	cfg  Config
	once sync.Once
)

func GetDefault() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatal(err)
		}
	})
	return &cfg
}

func (c *Config) String() string {
	str, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("couldn't parse config to string: %s\n", err.Error())
		return ""
	}
	return string(str)
}
