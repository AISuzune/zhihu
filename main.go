package main

import (
	"zhihu/api"
	"zhihu/boot"
)

func main() {
	boot.ViperSetup("./config/config.yaml")
	boot.LoggerSetup()
	api.InitRouter()
}
