package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"restful_go_project/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host" env-default:"localhost"`
		Port       string `yaml:"port" env-default:"27817"`
		Database   string `yaml:"database" env-default:"user-service"`
		AuthDB     string `yaml:"auth_db" env-default:""`
		Username   string `yaml:"username" env-default:""`
		Password   string `yaml:"password" env-default:""`
		Collection string `yaml:"collection" env-default:"localhost"`
	} `yaml:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logging.Init()
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
