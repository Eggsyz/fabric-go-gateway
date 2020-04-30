package app

import (
	"github.com/fabric-go-gateway/controllers"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  routes
 * @Version: 1.0.0
 * @Date: 12/7/19 11:27 上午
 */

func InitIris(port string) {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	mvc.Configure(app.Party("/api"), func(application *mvc.Application) {
		application.Party("/channel").Handle(new(controllers.ChannelController))
		application.Party("/blockchain").Handle(new(controllers.LedgerController))
		application.Party("/chaincode").Handle(new(controllers.ChaincodeController))
		application.Party("/peer").Handle(new(controllers.PeerController))

	})
	app.Run(iris.Addr("0.0.0.0:" + port))
}
