package aeronyx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/flights/search", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		dateStr := c.Query("date")

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
			return
		}

		results := SearchFlights(from, to, date)
		c.JSON(http.StatusOK, results)
	})

	fmt.Println("Starting server on :9002")
	r.Run(":9002")
}
