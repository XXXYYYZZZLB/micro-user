package main

import (
	"github.com/XXXYYYZZZLB/micro-user/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	//服务设置
	srv := service.New(
		service.Name("go.micro.service.user"),
		service.Version("latest"),
	)

	srv.Init()
	//创建数据库

	// Register handler
	srv.Handle(new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
