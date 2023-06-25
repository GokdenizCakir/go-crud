package main

import (
	"github.com/GokdenizCakir/go-crud/controllers"
	"github.com/GokdenizCakir/go-crud/initializers"
	"github.com/GokdenizCakir/go-crud/routes"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables();
	initializers.ConnectToDB();
	initializers.SyncDatabase();
}

func main() {
	router := gin.Default();
	
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true;
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config));

	router.Use(limits.RequestSizeLimiter(1024 * 10));

	router.POST("/login", controllers.Login);
	router.POST("/signup", controllers.Signup);
	
	router.GET("/signout", controllers.Signout);
		
	routes.ApplicantRoutes(router);
	routes.QuestionRoutes(router);
	routes.UserRoutes(router);

	router.Run();
}