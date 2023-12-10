package user

import (
	"context"
	"github.com/heroku/go-getting-started/common/bcrypt"
	"github.com/heroku/go-getting-started/common/logger"
	"time"
)

func Init(ctx context.Context, repo *Repository) {
	err := InitUsers(ctx, repo)
	if err != nil {
		logger.ErrorLogger.Printf("cannot init users %v", err)
	}
}

func InitUsers(ctx context.Context, repo *Repository) error {
	us, err := repo.GetUserByRole(ctx, Master)
	if err != nil {
		logger.ErrorLogger.Printf("error get user by role ", err)
		return err
	}
	if us == nil {
		pass, err := bcrypt.GenerateSecretKey(32)
		passHashed, err := bcrypt.HashPassword(pass)
		if err != nil {
			logger.ErrorLogger.Printf("error generating secret key: %v", err)
			return err
		}
		logger.InfoLogger.Printf("sunny is ", pass)
		result, err := repo.CreateUser(ctx, User{
			Role:      Master,
			Username:  "sunny",
			Password:  passHashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			logger.ErrorLogger.Printf("could not create user", err)
			return err
		}
		if result.InsertedID == nil {
			logger.ErrorLogger.Printf("could not create user")
			return nil
		}
	}
	return nil
}
