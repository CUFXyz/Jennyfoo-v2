package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	auth "v2/internal/Auth"
	database "v2/internal/Database"
	"v2/internal/storage"
	"v2/models"

	"github.com/gin-gonic/gin"
)

type EPHandler struct {
	DB      *database.DBInstance
	Auth    *auth.Authentificator
	Storage *storage.Cache
}

func EPStart(DB *database.DBInstance, Auth *auth.Authentificator, storage *storage.Cache) *EPHandler {
	return &EPHandler{
		DB:      DB,
		Auth:    Auth,
		Storage: storage,
	}
}

func (ep *EPHandler) IndexHandler(c *gin.Context) {

	users, err := ep.DB.GetUsers()
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

func (ep *EPHandler) RegisterHandler(c *gin.Context) {
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
			error.Error(fmt.Errorf("unmarshal | | RegisterHandler: %v", err)),
		)
		return
	}
	cryptedPass, err := ep.Auth.CryptPassword([]byte(user.Password))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("cryptpassword | | RegisterHandler: %v", err)),
		)
		return
	}

	user.Password = cryptedPass
	err = ep.DB.SendUser(user)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("senduser | | RegisterHandler: %v", err)),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		fmt.Sprintf("Successfuly registered user %v", user.Login),
	)
}

func (ep *EPHandler) LoginHandler(c *gin.Context) {
	var user, userd *models.User
	bytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("readall| | LoginHandler: %v", err)),
		)
		return
	}

	err = json.Unmarshal(bytes, &user)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("unmarshal| | LoginHandler: %v", err)),
		)
		return
	}

	userd, err = ep.DB.GetUser(user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("getuser | | LoginHandler: %v", err)),
		)
		return
	}

	if err = ep.Auth.AuthUser(userd.Password, user.Password); err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("authuser | | LoginHandler: %v", err)),
		)
		return
	}

	token := ep.Auth.GenerateToken(userd)
	if token == "" {
		c.JSON(
			http.StatusBadRequest,
			"generatetoken | | LoginHandler: empty token",
		)
		return
	}

	if err := ep.Storage.WriteCache(token, user.Login); err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("writecache | | LoginHandler: %v", err)),
		)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"token":   token,
			"Success": "Generated token",
		},
	)
}

func (ep *EPHandler) PromoteUserHandler(c *gin.Context) {
	var user, userdb *models.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("readall | | PromoteUserHandler: %v", err)),
		)
		return
	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("unmarshal | | PromoteUserHandler: %v", err)),
		)
		return
	}

	userdb, err = ep.DB.GetUser(user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("getuser | | PromoteUserHandler: %v", err)),
		)
		return
	}
	err = ep.DB.PromoteUser(*userdb, "admin")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			error.Error(fmt.Errorf("promoteuser | | PromoteUserHandler: %v", err)),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		fmt.Sprintf("Successfuly promoted user: %v", userdb.Login),
	)
}
