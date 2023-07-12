package main

import (
	"fmt"
	"net/http"
	"wc3_game_tracker/api"

	"github.com/gin-gonic/gin"
)

// type Players struct {
// 	ID   uuid.UUID `json:"id"`
// 	Name string    `json:"name"`

// }

// type Game struct {
// 	Id      uuid.UUID `json:"id"`
// 	Name    string    `json:"name"`
// 	Version string    `json:"version"`
// 	Map     string    `json:"map"`
// }

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)
	api.RegisterTournamentEndpoints(router)
	err := router.Run("0.0.0.0:8080")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return
	}
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", nil)
}
