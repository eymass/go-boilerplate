package Ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/service/config"
	"github.com/heroku/go-getting-started/user"
)

func GetUserRepository(ctx *gin.Context) *user.Repository {
	repo, ok := ctx.Get(config.UserRepositoryKey)
	if !ok {
		return nil
	}
	return repo.(*user.Repository)
}
