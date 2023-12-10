package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/common/bcrypt"
	Ctx "github.com/heroku/go-getting-started/common/ctx"
	"github.com/heroku/go-getting-started/common/logger"
	"log"
	"net/http"
	"time"
)

func Login(ctx *gin.Context) {
	var dto LoginDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	start := time.Now()
	err := dto.Validate()
	if err != nil {
		msg := fmt.Sprintf("validation error %s", err.Error())
		logger.ErrorLogger.Printf(msg)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}

	repo := Ctx.GetUserRepository(ctx)
	us, err := repo.GetUserByUserName(ctx, dto.Username)
	if err != nil {
		msg := fmt.Sprintf("error getting user %s", err.Error())
		logger.ErrorLogger.Printf(msg)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("check user took %f", elapsed)
	if us == nil {
		msg := fmt.Sprintf("username %s does not found", dto.Username)
		logger.ErrorLogger.Printf(msg)
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	// Check the password
	if !bcrypt.CheckPasswordHash(dto.Password, us.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
	start = time.Now()
	token, err := GenerateJWT(us.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error generating JWT token"})
		return
	}
	elapsed = time.Since(start).Seconds()
	log.Printf("JWT took %f", elapsed)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
