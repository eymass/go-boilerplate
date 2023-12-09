package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/common/logger"
	"github.com/heroku/go-getting-started/service/config"
	"io/ioutil"
	"net/http"
	"reflect"
)

const PROXY_SECRET = "$6ukvh5y8^*99##AySv9"

type RequestBody interface {
	Validate() (bool, string)
}

func SecretMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value := ctx.Request.Header.Get("Authorization")
		if value != PROXY_SECRET {
			logger.ErrorLogger.Printf("bad request")
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("bad request"))
			return
		}
	}
}

func ValidateRequest(dtoType RequestBody) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		interfaceInstance := reflect.New(reflect.TypeOf(dtoType)).Interface()
		requestBytesInterface, exists := ctx.Get(config.RequestBytesKey)
		requestBytes := requestBytesInterface.([]byte)

		if !exists {
			message := fmt.Errorf("ValidateRequest middleware error")
			logger.ErrorLogger.Printf(message.Error())
			ctx.String(http.StatusBadRequest, message.Error())
			return
		}

		err := json.Unmarshal(requestBytes, &interfaceInstance)
		if err != nil {
			message := fmt.Errorf("ValidateRequest middleware error, %s", err.Error())
			logger.ErrorLogger.Printf(message.Error())
			ctx.String(http.StatusBadRequest, message.Error())
			return
		}

		dtoInstance, isValidCast := interfaceInstance.(RequestBody)
		if !isValidCast {
			message := fmt.Errorf("ValidateRequest middleware error not valid cast")
			logger.ErrorLogger.Printf(message.Error())
			ctx.String(http.StatusBadRequest, message.Error())
			return
		}

		isValid, msg := dtoInstance.Validate()
		if isValid == false {
			message := fmt.Errorf("ValidateRequest middleware error not valid, %s", msg)
			logger.ErrorLogger.Printf(message.Error())
			ctx.String(http.StatusBadRequest, message.Error())
			return
		}

		ctx.Set(config.RequestBodyKey, dtoInstance)
	}
}

func SetRequestBytes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBytes, err := ioutil.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()

		if err != nil {
			message := fmt.Errorf("SetRequestBytes middleware error, %s", err.Error())
			logger.ErrorLogger.Printf(message.Error())
			ctx.String(http.StatusBadRequest, message.Error())
			return
		}

		ctx.Set(config.RequestBytesKey, requestBytes)
	}
}
