package conf

import (
	"github.com/spf13/viper"

	ccLog "cc/application/log"
)

func initConfig(configFileName, configPath string) {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		ccLog.Fatal(err.Error())
	}
}

func YamlConfig(configFile, configPath string) {
	viper.SetConfigType("yaml")
	initConfig(configFile, configPath)
}

func TomlConfig(configFile, configPath string) {
	viper.SetConfigType("toml")
	initConfig(configFile, configPath)
}

func JsonConfig(configFile, configPath string) {
	viper.SetConfigType("json")
	initConfig(configFile, configPath)
}

func IniConfig(configFile, configPath string) {
	viper.SetConfigType("ini")
	initConfig(configFile, configPath)
}

func EnvConfig(configFile, configPath string) {
	viper.SetConfigType("dotenv")
	initConfig(configFile, configPath)
}
