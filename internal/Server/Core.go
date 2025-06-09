package server

import (
	"log"
	database "v2/internal/Database"

	"github.com/gin-gonic/gin"
)

type JFServer struct {
	Addr   string
	Engine *gin.Engine
	Db     *database.DBInstance
}

func JFServerSetup(addr string) *JFServer {
	return &JFServer{
		Addr:   addr,
		Engine: gin.Default(),
		Db:     database.Connect(),
	}
}

func (jfs *JFServer) Run() {
	jfs.Engine.GET("/", jfs.IndexHandler)
	jfs.Engine.POST("/register", jfs.RegisterHandler)
	log.Fatal(
		jfs.Engine.Run(jfs.Addr),
	)
}
