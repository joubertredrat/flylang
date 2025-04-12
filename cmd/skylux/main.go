package skylux

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	r := gin.Default()

	r.GET("/api/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server_status": "Running",
			"timestamp":     time.Now().Unix(),
		})
	})

	r.GET("/api/v2/flights", func(c *gin.Context) {
		departingFrom := c.Query("departingFrom")
		arrivingTo := c.Query("arrivingTo")
		departureDateStr := c.Query("departureDate")
		departureDate, err := time.Parse("20060102", departureDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid departureDate format. Expected format: YYYYMMDD",
			})
			return
		}

		flights := flights(departingFrom, arrivingTo, departureDate)
		c.JSON(http.StatusOK, flights)
	})

	srv := &http.Server{
		Addr:    ":19004",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
