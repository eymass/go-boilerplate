package api

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/auth"
	"github.com/heroku/go-getting-started/common/logger"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/service/config"
	"github.com/heroku/go-getting-started/user"
	"log"
	"os"
)

func Start() {
	g := gin.New()
	routeGroup := g.Group("/", initiateContextMiddleware())
	authGroup := routeGroup.Group("/auth")
	auth.Routes(authGroup)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	logger.InfoLogger.Printf("starting service %v", port)
	g.Run(":" + port)
}

func initiateContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		database, closer := db.GetDatabase(ctx)
		defer closer()
		repo := user.NewUserRepository(database)
		ctx.Set(config.UserRepositoryKey, &repo)
		ctx.Next()
	}
}
