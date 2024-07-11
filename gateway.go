package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func setConfig() {
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config") // path to look for the config file in
	viper.AddConfigPath(".")       // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

}

func main() {
	setConfig()
	env := fmt.Sprintf("%s%s", "environment: ", viper.GetString("environment"))
	webAppUrl := fmt.Sprintf("%s%s", "webapp.url: ", viper.GetString("webapp.url"))
	webAppPortLocal := fmt.Sprintf("%s%d", "webapp.port.local: ", viper.GetInt("webapp.port.local"))
	webAppPortDev := fmt.Sprintf("%s%d", "webapp.port.dev: ", viper.GetInt("webapp.port.dev"))
	webAppPortStaging := fmt.Sprintf("%s%d", "webapp.port.staging: ", viper.GetInt("webapp.port.staging"))
	webAppPortProd := fmt.Sprintf("%s%d", "webapp.port.prod: ", viper.GetInt("webapp.port.prod"))

	gatewayPort := fmt.Sprintf("%s%d", "gateway.port: ", viper.GetInt("gateway.port"))

	jenkinsUrl := fmt.Sprintf("%s%s", "jenkins.url: ", viper.GetString("jenkins.url"))
	jenkinsPort := fmt.Sprintf("%s%d", "jenkins.port: ", viper.GetInt("jenkins.port"))

	fmt.Println(env)
	fmt.Println(webAppUrl)
	fmt.Println(webAppPortLocal)
	fmt.Println(webAppPortDev)
	fmt.Println(webAppPortStaging)
	fmt.Println(webAppPortProd)

	fmt.Println(gatewayPort)

	fmt.Println(jenkinsUrl)
	fmt.Println(jenkinsPort)
}
