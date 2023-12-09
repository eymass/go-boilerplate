package service

import (
	"fmt"
	"github.com/heroku/go-getting-started/api"
	"github.com/heroku/go-getting-started/service/config"
)

type Service struct {
	Config *config.Cfg
}

func InitService() *Service {
	service := new(Service)
	service.Config = new(config.Cfg)
	return service
}

func (service *Service) Start() {
	fmt.Println("Starting Service")
	api.Start()
}
