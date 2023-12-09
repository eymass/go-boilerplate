package auth

import "github.com/gin-gonic/gin"

// gin router for auth

func Routes(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/login", Login)
}
