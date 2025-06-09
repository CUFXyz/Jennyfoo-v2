package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (jfs *JFServer) IndexHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"Status": "OK",
		},
	)
}

func (jfs *JFServer) RegisterHandler(c *gin.Context) {

}
