package server

import (
	"log"
	database "v2/internal/Database"

	"github.com/gin-gonic/gin"
)

type JFServer struct {
	Addr       string
	Engine     *gin.Engine
	Db         *database.DBInstance
	WebHandler *WebHandler
}

func JFServerSetup(addr string) *JFServer {
	Db := database.Connect()
	Engine := gin.Default()
	WebHandler := WebHandlerStart(Db)
	return &JFServer{
		Addr:       addr,
		Engine:     Engine,
		Db:         Db,
		WebHandler: WebHandler,
	}
}

func (jfs *JFServer) Run() {
	jfs.Engine.GET("/", jfs.WebHandler.IndexHandler)
	jfs.Engine.POST("/register", jfs.WebHandler.RegisterHandler)

	log.Fatal(
		jfs.Engine.Run(jfs.Addr),
	)
}
