package main

import (
	"github.com/kristiyankiryakov/go-sms-verify/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router :=gin.Default()

	// Initialize config

	app := api.Config{Router: router}

	// routes
	app.Routes()

	router.Run(":8000")
}
