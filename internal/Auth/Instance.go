package auth

import (
	database "v2/internal/Database"
	"v2/internal/storage"
)

type Authentificator struct {
	DB           *database.DBInstance
	localStorage *storage.Cache
}

func SetupAuthentificator(DB *database.DBInstance, storage *storage.Cache) *Authentificator {
	return &Authentificator{
		DB:           DB,
		localStorage: storage,
	}
}
