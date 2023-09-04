package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:8080/home",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func handleLogout(c *gin.Context) {
	// Invalidate the session or token
	c.Set("user", nil)

	// Redirect to homepage
	c.Redirect(http.StatusFound, "/")
}

func handleLogin(c *gin.Context) {
	authURL := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, authURL)
}

func handleHome(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func main() {
	r := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/", redirectToHome)
	r.GET("/login", handleLogin)
	r.GET("/home", handleHome)
	err := r.Run(":8080")
	if err != nil {
		log.Printf("Error in running server: %s", err.Error())
	} // Change the port as per your requirement
}

func redirectToHome(context *gin.Context) {
	context.Redirect(http.StatusFound, "/home")
}

func welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Welcome to the server!"})
}
