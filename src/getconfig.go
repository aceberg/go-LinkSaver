package main

import (
	"github.com/spf13/viper"
)

func get_config() (config Conf) {
	viper.SetDefault("DBPATH", "data/db.sqlite")
	viper.SetDefault("GUIIP", "localhost")
	viper.SetDefault("GUIPORT", "8841")
	viper.SetDefault("THEME", "minty")

    viper.SetConfigFile("data/config")
	viper.SetConfigType("env")
    viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)
	config.Theme = viper.Get("THEME").(string)

	return config
}