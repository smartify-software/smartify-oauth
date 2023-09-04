package smartify_oauth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     "your-client-id",
	ClientSecret: "your-client-secret",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func handleGoogleLogin(c *gin.Context) {
	authURL := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, authURL)
}
func handleGoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
	// Use token to get user info, etc.
}

func main() {
	r := gin.Default()
	r.GET("/login", handleGoogleLogin)
	r.GET("/auth/google/callback", handleGoogleCallback)
	err := r.Run(":8080")
	if err != nil {
		log.Printf("Error in running server: %s", err.Error())
	} // Change the port as per your requirement
}
