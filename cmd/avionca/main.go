package avionca

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Run(ctx context.Context) error {
	graphql := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	http.Handle("/graphql", graphql)
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	srv := &http.Server{
		Addr: ":19002",
	}

	log.Printf("	Avionca GraphQL API:")
	log.Printf("	http://127.0.0.1%s/grapqhl", srv.Addr)
	log.Printf("	http://127.0.0.1%s/playground", srv.Addr)
	log.Printf("")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Stopped Avionca.")
	return srv.Shutdown(shutdownCtx)
}
