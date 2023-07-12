package main

import (
	"fmt"
	"net/http"
	"wc3_game_tracker/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)
	api.RegisterTournamentEndpoints(router.Group("/api/v1/"))
	err := router.Run("0.0.0.0:8080")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return
	}
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", nil)
}
