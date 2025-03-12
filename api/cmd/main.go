package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scallyt/neko/internal/db"
	router "github.com/scallyt/neko/internal/routes"
)

func main() {
	r := gin.Default()

	router.RegisterRoutes(r)

	db.ConnectDB()

	r.Run("localhost:8080")
}
