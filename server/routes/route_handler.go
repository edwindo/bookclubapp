package routes

import (
	"net/http"
	"github.com/YoshiRussell/bookclubapp/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:	[]string{"*"},
		AllowMethods: 	[]string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:	[]string{"Authorization"},
	}))
	r.GET("/", rootHandler)
	r.GET("/mydashboard", middleware.Auth0Middleware(), dashboardHandler)
	r.GET("/epic", epicHandler)
	return r
}

// Serves index.html containing react script
func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a placeholder for the homepage.")
}

// This should only be reached by authenticated users
// Must include Authorization header with access token in request
func dashboardHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"username" : "Yoshi",
		"pageNumber" : 22,
	})
}

func epicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "fortnite",
	})
}
