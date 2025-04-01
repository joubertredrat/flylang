package shared

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AirportResponse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func Run() {
	r := gin.Default()

	r.GET("/airports", func(c *gin.Context) {
		response := []AirportResponse{}

		for _, airport := range Airports() {
			response = append(response, AirportResponse{
				Code:    airport.Code,
				Name:    airport.Name,
				City:    airport.City,
				Country: airport.Country,
			})
		}

		c.JSON(http.StatusOK, response)
	})

	fmt.Println("Starting server on :9001")
	r.Run(":9001")
}
