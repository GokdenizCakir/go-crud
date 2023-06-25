package middlewares

import "github.com/gin-gonic/gin"

func RoleBasedAuth(c *gin.Context) {
	c.Get("user")
}