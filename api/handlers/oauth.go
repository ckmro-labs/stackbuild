package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
)

//OAuth ...
func OAuth(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// OAuthCallback handles the oauth callback which finishes the auth procedure. It checks for the admin flag for the user, and if found it will set the user as admin for the rest of the session
func OAuthCallback(c *gin.Context) {
	q := c.Request.URL.Query()
	provider := c.Param("provider")
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	logrus.Infof("github user: %v", user)
	if err != nil {
		logrus.Errorf("CompleteUserAuth Err: %v", err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
