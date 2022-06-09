package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Apps     AppsConfig     `json:"apps"`
	Database DatabaseConfig `json:"database"`
}

type AppsConfig struct {
	Name     string `json:"name"`
	HttpPort string `json:"httpPort"`
}

type DatabaseConfig struct {
	Host               string        `json:"host"`
	Username           string        `json:"username"`
	Password           string        `json:"password"`
	Port               int           `json:"port"`
	Name               string        `json:"name"`
	MinIdleConnections int           `json:"minIdleConnections"`
	MaxOpenConnections int           `json:"maxOpenConnections"`
	ConnMaxLifeTime    time.Duration `json:"connMaxLifeTime"`
}

func NewConfig(path string) (config *Config) {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	return
}
