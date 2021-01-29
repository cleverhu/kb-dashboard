package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	_ "knowledgeBase/src/common"
	"knowledgeBase/src/configuration"
	"knowledgeBase/src/controllers"
	"knowledgeBase/src/middlewares"
)

func main() {
	goft.Ignite().
		Config(configuration.NewDBConfig(), configuration.NewKbUserServiceConfig(),configuration.NewRedisConfig()).
		Attach(middlewares.NewCors(),middlewares.NewKbUserIDCheck()).
		Mount("", controllers.NewKbUserController()).
		Launch()
}
