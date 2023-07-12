package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Tournament struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

var tournaments = []Tournament{
	{Id: uuid.New(), Name: "Creepcamp League", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 3)},
	{Id: uuid.New(), Name: "League1", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 4)},
	{Id: uuid.New(), Name: "League2", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 5)},
}

func RegisterTournamentEndpoints(router *gin.Engine) {
	router.GET("/tournaments", allTournamentsHandler)
}

func allTournamentsHandler(context *gin.Context) {
	// Get all tournaments and return to client
	context.IndentedJSON(http.StatusOK, tournaments)
}
