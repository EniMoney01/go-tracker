package main

import (
	"os"

	"github.com/EniMoney01/go-react-calorie-tracker/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntries)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("entry/delete/:id", routes.DeleteIngredient)

	router.Run(":" + port)
}
