package service

import (
	"context"
	"fmt"
	"github.com/heroku/go-getting-started/api"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/service/config"
	"github.com/heroku/go-getting-started/user"
)

type Service struct {
	Config *config.Cfg
}

func InitService() *Service {
	service := new(Service)
	service.Config = new(config.Cfg)
	ctx := context.Background()
	database, closer := db.GetDatabase(ctx)
	defer closer()
	repo := user.NewUserRepository(database)
	user.Init(ctx, repo)
	return service
}

func (service *Service) Start() {
	fmt.Println("Starting Service")
	api.Start()
}
