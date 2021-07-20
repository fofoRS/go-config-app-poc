package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	viper      *viper.Viper
	envProfile string
}

type PropertyParser func(interface{}) (interface{}, error)

const (
	ENV_PROFILE = "ENV_PROFILE"
	LOCAL_ENV   = "LOCAL"
)

func SetupConfig() AppConfig {
	v := viper.New()
	profile, ok := os.LookupEnv(ENV_PROFILE)
	if ok == true && profile != LOCAL_ENV {
		return AppConfig{envProfile: profile}
	} else {
		err := setupLocalEnv(*v)
		if err != nil {
			panic(err)
		}
	}
	return AppConfig{viper: v, envProfile: profile}
}

func setupLocalEnv(v viper.Viper) error {
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	v.AutomaticEnv()
	return nil
}

func (config *AppConfig) GetValue(key string, parser PropertyParser) (interface{}, error) {
	config.viper.BindEnv(key)
	val := config.viper.Get(key)
	fmt.Printf("%s -- %t\n", val, val)
	val, err := parser(val)
	if err != nil {
		return nil, errors.New("Error parsing value")
	}
	return val, nil
}
