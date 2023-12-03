package main

import (
	_ "app/docs" // Importez le dossier docs généré par Swaggo
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @host      localhost:8080

// @securityDefinitions.basic  None
// @BasePath /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	// Swagger setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Hello World endpoint
	r.GET("/hello", HelloHandler)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

// HelloHandler handles the /hello endpoint
// @Summary Returns a greeting with the user's name
// @Produce json
// @Success 200 {object} string "OK"
// @Router /hello [get]
func HelloHandler(c *gin.Context) {
	user := os.Getenv("USER")
	if user == "" {
		user = "Guest"
	}

	message := fmt.Sprintf("Bonjour %s", user)
	c.JSON(http.StatusOK, message)
}
