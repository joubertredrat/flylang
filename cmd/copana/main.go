package copana

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	tokenValidityDuration = 2 * time.Minute
)

var (
	tokenData       = ""
	tokenExpiration time.Time
	tokenMutex      sync.Mutex
)

func Run(ctx context.Context) error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/webservices/copana/version27_8_5/system/ping.wsd", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"response": "pong",
			"time":     time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.POST("/webservices/copana/version27_8_5/authentication/token.wsd", func(c *gin.Context) {
		tokenMutex.Lock()
		defer tokenMutex.Unlock()

		simulatedDelay := time.Duration(rand.Intn(700)+1020) * time.Millisecond
		time.Sleep(simulatedDelay)

		tokenData = newToken()
		tokenExpiration = time.Now().Add(tokenValidityDuration)

		c.XML(http.StatusOK, Token{
			Value:  tokenData,
			Expiry: fmt.Sprintf("%d seconds", int(tokenValidityDuration.Seconds())),
		})
	})

	r.POST("/webservices/copana/version27_8_5/flight/search.wsd", func(c *gin.Context) {
		simulatedDelay := time.Duration(rand.Intn(800)+1520) * time.Millisecond
		time.Sleep(simulatedDelay)

		iataFrom := c.PostForm("iata:departure_code")
		iataTo := c.PostForm("iata:arrival_code")
		dateDeparture := c.PostForm("date:departure")
		token := c.PostForm("token")

		tokenMutex.Lock()
		defer tokenMutex.Unlock()
		if token != tokenData || time.Now().After(tokenExpiration) {
			c.XML(http.StatusOK, Error{
				Code:    7,
				Message: "Invalid or expired token",
			})
			return
		}

		if iataFrom == "" || iataTo == "" || dateDeparture == "" {
			c.XML(http.StatusOK, Error{
				Code:    9,
				Message: "Missing required fields: iata:from, iata:to, or date:departure",
			})
			return
		}

		parsedDate, err := time.Parse("02/01/06", dateDeparture)
		if err != nil {
			c.XML(http.StatusOK, Error{
				Code:    15,
				Message: "Invalid date format. Use DD/MM/YY (e.g., 12/04/91).",
			})
			return
		}

		flights := flights(iataFrom, iataTo, parsedDate)
		if len(flights) == 0 {
			c.String(http.StatusOK, "No flights found")
			return
		}

		c.XML(http.StatusOK, flights)
	})

	srv := &http.Server{
		Addr:    ":19003",
		Handler: r,
	}

	routes := r.Routes()
	log.Printf("	Copana Webservice:")
	for _, route := range routes {
		log.Printf("	%-6s http://127.0.0.1%s%s", route.Method, srv.Addr, route.Path)
	}
	log.Printf("")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Stopped Copana.")
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
