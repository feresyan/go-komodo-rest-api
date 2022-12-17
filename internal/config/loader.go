package config

import (
	"github.com/spf13/viper"
	"gokomodo-assesment/internal/config/env"
	"log"
	"os"
	"strings"
)

type Config struct {
	EnvSystem env.Environment
}

var configuration Config

func init() {
	var err error
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(dir + "/internal/config/env")
	viper.SetConfigType("yaml")
	viper.SetConfigName("env.yml")
	err = viper.MergeInConfig()
	if err != nil {
		log.Println("Cannot read env config: %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	viper.Unmarshal(&configuration)

}

func GetConfig() *Config {
	return &configuration
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
