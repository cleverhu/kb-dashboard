package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	_ "knowledgeBase/src/common"
	"knowledgeBase/src/configuration"
	"knowledgeBase/src/controllers"
	"knowledgeBase/src/middlewares"
)

func main() {
	//conn, err := grpc.Dial("101.132.107.2:8088", grpc.WithInsecure())
	//fmt.Println(err)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(conn.GetState())
	//return
	goft.Ignite().
		Config(configuration.NewDBConfig(),
			configuration.NewKbUserServiceConfig(),
			configuration.NewRedisConfig(),
			configuration.NewGrpcServiceConfig()).
		Attach(middlewares.NewKbUserIDCheck()).
		Mount("", controllers.NewKbUserController()).
		Launch()
}
