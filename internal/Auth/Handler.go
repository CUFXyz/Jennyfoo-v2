package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Authentificator) AuthentificatorHandler(c *gin.Context) {
	authInfo := c.GetHeader("Authorization")
	if authInfo == "" {
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"Error": "No token provided",
			},
		)
		c.Abort()
		return
	}

	token, err := a.localStorage.GetValue(authInfo)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
	}

	if token != authInfo {
		c.JSON(
			http.StatusForbidden,
			"Invalid token",
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		"Successful session",
	)

	c.Next()
}
