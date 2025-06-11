package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	database "v2/internal/Database"
	"v2/models"

	"github.com/gin-gonic/gin"
)

type WebHandler struct {
	DB *database.DBInstance
}

func WebHandlerStart(DB *database.DBInstance) *WebHandler {
	return &WebHandler{
		DB: DB,
	}
}

func (wh *WebHandler) IndexHandler(c *gin.Context) {

	users, err := wh.DB.GetUsers()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(err),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		users,
	)
}

func (wh *WebHandler) RegisterHandler(c *gin.Context) {
	var user models.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			"Cannot read body",
		)
		return
	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	err = wh.DB.SendUser(user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	c.JSON(
		http.StatusOK,
		fmt.Sprintf("Successfuly registered user %v", user.Login),
	)
}
