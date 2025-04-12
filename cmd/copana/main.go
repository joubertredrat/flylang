package copana

import (
	"context"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	tokenValidityDuration = 10 * time.Minute
)

var (
	tokenData       = ""
	tokenExpiration time.Time
	tokenMutex      sync.Mutex
)

func Run(ctx context.Context) error {
	r := gin.Default()

	r.POST("/webservices/copana/version27_8_5/system/ping.wsd", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"response": "pong",
			"time":     time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.POST("/webservices/copana/version27_8_5/authentication/token.wsd", func(c *gin.Context) {
		tokenMutex.Lock()
		defer tokenMutex.Unlock()

		tokenData = newToken()
		tokenExpiration = time.Now().Add(tokenValidityDuration)

		c.XML(http.StatusOK, gin.H{
			"token":      tokenData,
			"expires_in": "1 minute",
		})
	})

	r.POST("/webservices/copana/version27_8_5/flight/search.wsd", func(c *gin.Context) {
		iataFrom := c.PostForm("iata:departure_code")
		iataTo := c.PostForm("iata:arrival_code")
		dateDeparture := c.PostForm("date:departure")
		token := c.PostForm("token")

		tokenMutex.Lock()
		defer tokenMutex.Unlock()
		if token != tokenData || time.Now().After(tokenExpiration) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		if iataFrom == "" || iataTo == "" || dateDeparture == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing required fields: iata:from, iata:to, or date:departure",
			})
			return
		}

		parsedDate, err := time.Parse("02/01/06", dateDeparture)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format. Use DD/MM/YY (e.g., 12/04/25).",
			})
			return
		}

		c.XML(http.StatusOK, gin.H{
			"message":        "ok",
			"iata_from":      iataFrom,
			"iata_to":        iataTo,
			"date_departure": parsedDate.Format("2006-01-02"),
		})
	})

	srv := &http.Server{
		Addr:    ":19003",
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

func newToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 16)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}
