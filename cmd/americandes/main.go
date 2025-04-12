package americandes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	r := gin.Default()

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server_status": "Running",
			"timestamp":     time.Now().Unix(),
		})
	})

	r.GET("/flights/search", func(c *gin.Context) {
		codeFrom := c.Query("code_from")
		codeTo := c.Query("code_to")
		departureDay := c.Query("departure_day")
		departureMonth := c.Query("departure_month")
		departureYear := c.Query("departure_year")

		departureDate, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", departureYear, departureMonth, departureDay))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid departure date"})
			return
		}

		today := time.Now().Truncate(24 * time.Hour)
		if !departureDate.After(today) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date",
			})
			return
		}

		if (codeFrom != MIA && codeFrom != SCL) || (codeTo != MIA && codeTo != SCL) {
			c.JSON(http.StatusOK, FlightsResponse{
				Metadata: FlightsResponseMetadata{
					TotalFlights: 0,
					ServerTime:   time.Now().Format(time.RFC3339),
					Search: FlightsResponseMetadataSearch{
						DepartureDay:   departureDay,
						DepartureMonth: departureMonth,
						DepartureYear:  departureYear,
						CodeFrom:       codeFrom,
						CodeTo:         codeTo,
					},
				},
				Flights: []Flight{},
			})
			return
		}

		flights := flights(codeFrom, codeTo, departureDate)

		c.JSON(http.StatusOK, FlightsResponse{
			Metadata: FlightsResponseMetadata{
				TotalFlights: len(flights),
				ServerTime:   time.Now().Format(time.RFC3339),
				Search: FlightsResponseMetadataSearch{
					DepartureDay:   departureDay,
					DepartureMonth: departureMonth,
					DepartureYear:  departureYear,
					CodeFrom:       codeFrom,
					CodeTo:         codeTo,
				},
			},
			Flights: flights,
		})
	})

	srv := &http.Server{
		Addr:    ":19002",
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
