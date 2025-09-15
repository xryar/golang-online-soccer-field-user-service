package config

import (
	"os"
	"user-service/common/util"

	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port                  int            `json:"port"`
	AppName               string         `json:"appName"`
	AppEnv                string         `json:"appEnv"`
	SignatureKey          string         `json:"signatureKey"`
	Database              DatabaseConfig `json:"database"`
	RateLimiterMaxRequest float64        `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int            `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string         `json:"jwtSecretKey"`
	JwtExpirationTime     int            `json:"jwtExpirationTime"`
}

type DatabaseConfig struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Errorf("failed to bind config: %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}
