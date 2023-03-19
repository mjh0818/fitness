package cmd

import (
	"fitness/conf"
	"fitness/dao"
	"fitness/routes"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Start() {
	conf.InitConfig()
	dao.InitDB()
	r := routes.SetRouter()
	if port := viper.GetString("server.port"); port != "" {
		log.Panic(r.Run(":" + port))
	} else {
		log.Panic(r.Run())
	}
}

func Clean() {
	fmt.Println("=====clean=====")
}
