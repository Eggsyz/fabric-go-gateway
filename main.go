package main

import (
	"github.com/fabric-go-gateway/app"
	"github.com/fabric-go-gateway/config"
	"github.com/fabric-go-gateway/logging"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020-03-27 15:43
 */

func main() {
	config.InitConfig("conf/app.yaml")
	logging.Init(logging.Config{
		LogLevel: config.Conf.Level,
		FilePath: config.Conf.LogPath,
	})
	app.InitIris(config.Conf.Port)
}
