package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"

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
	webAppUrl := viper.GetString("webapp.url")
	webAppPort := viper.GetInt("webapp.port")
	gatewayPort := viper.GetInt("gateway.port")

	envStr := fmt.Sprintf("%s%s", "environment: ", env)
	webAppUrlStr := fmt.Sprintf("%s%s", "webapp.url: ", webAppUrl)
	webAppPortStr := fmt.Sprintf("%s%d", "webapp.port: ", webAppPort)
	gatewayPortStr := fmt.Sprintf("%s%d", "gateway.port: ", gatewayPort)

	fmt.Println(envStr)
	fmt.Println(webAppUrlStr)
	fmt.Println(webAppPortStr)
	fmt.Println(gatewayPortStr)

	r := gin.Default()

	webAppRoute := ""

	if env != "prod" {
		webAppRoute = fmt.Sprintf("/%s/*proxyPath", viper.GetString("environment"))
	} else {
		webAppRoute = "/*proxyPath"
	}

	r.Any(webAppRoute, func(c *gin.Context) {

		target := fmt.Sprintf("http://%s:%d", webAppUrl, webAppPort) // The target server URL
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			// Rewrite the request URL here
			req.URL.Path = c.Param("proxyPath")
			// Optionally, you can modify headers or other parts of the request
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Start the server
	port := fmt.Sprintf("%s%d", ":", gatewayPort)
	r.Run(port)
}
