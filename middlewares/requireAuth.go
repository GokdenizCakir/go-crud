package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GokdenizCakir/go-crud/initializers"
	"github.com/GokdenizCakir/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization");
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"});
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"});
				return
			}

			var user models.User
			initializers.DB.First(&user, claims["sub"])

			if user.ID == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}

			if !isRoleAllowed(user.Role, allowedRoles) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
				return
			}

			c.Set("user", user);

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func isRoleAllowed(role string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
        if role == allowedRole {
            return true
        }
    }
    return false
}