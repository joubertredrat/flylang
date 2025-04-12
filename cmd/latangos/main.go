package latangos

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	r.GET("/api/flight", func(c *gin.Context) {
		date := c.Query("date")
		origin := c.Query("origin")
		destination := c.Query("destination")

		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format. Use Y-m-d (e.g., 2025-04-12).",
			})
			return
		}

		c.JSON(http.StatusOK, flights(origin, destination, parsedDate))
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
