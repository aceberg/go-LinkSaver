package main

import (
	"github.com/spf13/viper"
)

var configPath = "/data/linksaver/config"

func get_config() (config Conf) {
	viper.SetDefault("DBPATH", "/data/linksaver/db.sqlite")
	viper.SetDefault("GUIIP", "0.0.0.0")
	viper.SetDefault("GUIPORT", "8841")
	viper.SetDefault("THEME", "minty")

    viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
    viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)
	config.Theme = viper.Get("THEME").(string)

	return config
}

func write_config() {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.Set("THEME", AppConfig.Theme)
	viper.WriteConfig()
}