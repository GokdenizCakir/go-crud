package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/GokdenizCakir/go-crud/initializers"
	"github.com/GokdenizCakir/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct{
		Email string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash the password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user);
	 
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		});
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(ctx *gin.Context) {
	var body struct{
		Email string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	var user models.User;
	result := initializers.DB.First(&user, "email = ?", body.Email)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		});
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password));
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		});
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")));
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign token",
		});
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode);
	ctx.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
	})
}

func Signout(c *gin.Context) {
	_, exists := c.Get("user");
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	c.Set("user", nil);
	c.SetCookie("Authorization", "", -1, "", "", false, true);

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed out successfully",
	})
}

func GetUsers(ctx *gin.Context) {
	var users []models.User;

	initializers.DB.Find(&users);

	ctx.JSON(200, gin.H{
		"usersFound": len(users),
		"users": users,
	})
}