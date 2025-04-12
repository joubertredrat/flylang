package latangos

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	FlightNumber string
	Origin       string
	Destination  string
	Departure    time.Time
	Arrival      time.Time
	BasePrice    float64
}

func Run(ctx context.Context) error {
	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) {
		currentTime := time.Now().Format(time.RFC3339)
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   currentTime,
		})
	})

	r.GET("/api/flight", func(c *gin.Context) {
		date := c.Query("date")
		origin := c.Query("origin")
		destination := c.Query("destination")

		c.JSON(http.StatusOK, gin.H{
			"message":     "ok",
			"date":        date,
			"origin":      origin,
			"destination": destination,
		})
	})

	srv := &http.Server{
		Addr:    ":19001",
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
