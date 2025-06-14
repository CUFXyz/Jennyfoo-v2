package auth

import (
	"net/http"
	"v2/models"

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
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	login, err := a.localStorage.GetValue(authInfo)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := a.DB.GetUser(&models.User{Login: login})
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if login != user.Login {
		c.JSON(
			http.StatusForbidden,
			"Invalid token",
		)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(
		http.StatusOK,
		"Successful session",
	)

	c.Next()
}

func (a *Authentificator) AdministratorHandler(c *gin.Context) {
	authInfo := c.GetHeader("Authorization")
	login, err := a.localStorage.GetValue(authInfo)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
		return
	}

	user, err := a.DB.GetUser(&models.User{Login: login})
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
		return
	}

	if user.Role != "admin" {
		c.JSON(
			http.StatusForbidden,
			"You don't have permission to watch this page",
		)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(
		http.StatusOK,
		"Welcome admin",
	)

	c.Next()

}
