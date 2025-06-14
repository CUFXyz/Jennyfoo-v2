package server

import (
	"log"
	auth "v2/internal/Auth"
	database "v2/internal/Database"
	"v2/internal/storage"

	"github.com/gin-gonic/gin"
)

type JFServer struct {
	Addr            string
	Engine          *gin.Engine
	Db              *database.DBInstance
	EPHandler       *EPHandler
	Authentificator *auth.Authentificator
}

func JFServerSetup(addr string) *JFServer {
	localStorage := storage.NewCache()
	Db := database.Connect()
	Engine := gin.Default()
	Authentificator := auth.SetupAuthentificator(Db, localStorage)
	EPHandler := EPStart(Db, Authentificator, localStorage)

	return &JFServer{
		Addr:            addr,
		Engine:          Engine,
		Db:              Db,
		EPHandler:       EPHandler,
		Authentificator: Authentificator,
	}
}

func (jfs *JFServer) Run() {
	signedin := jfs.Engine.Group("/", jfs.EPHandler.Auth.AuthentificatorHandler)
	admin := signedin.Group("/admin", jfs.EPHandler.Auth.AdministratorHandler)

	admin.GET("/userlist", jfs.EPHandler.IndexHandler)
	admin.POST("/promote", jfs.EPHandler.PromoteUserHandler)

	jfs.Engine.POST("/register", jfs.EPHandler.RegisterHandler)
	jfs.Engine.POST("/login", jfs.EPHandler.LoginHandler)

	log.Fatal(
		jfs.Engine.Run(jfs.Addr),
	)
}
