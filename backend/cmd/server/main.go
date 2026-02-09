package main

import (
	"flag"
	"log"

	"qizhan/backend/internal/biz"
	"qizhan/backend/internal/conf"
	"qizhan/backend/internal/data"
	"qizhan/backend/internal/server"
	"qizhan/backend/internal/service"

	"github.com/go-kratos/kratos/v2"
)

func main() {
	configPath := flag.String("conf", "configs/config.yaml", "config path")
	flag.Parse()

	cfg, err := conf.Load(*configPath)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	db, err := data.NewDB(cfg)
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}

	repo := data.NewAboutRepo(db)
	uc := biz.NewAboutUsecase(repo)
	svc := service.NewAboutService(uc)
	httpSrv := server.NewHTTPServer(cfg.Server.HTTP.Addr, svc)

	app := kratos.New(
		kratos.Name("qizhan.backend"),
		kratos.Server(httpSrv),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
