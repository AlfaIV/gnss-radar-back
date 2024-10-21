package delivery

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/Gokert/gnss-radar/internal/delivery/graphql"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/service"
	"log"
	"net/http"
)

type App struct {
	config generated.Config
}

func NewApp(service2 *service.Service) *App {
	return &App{
		config: generated.Config{Resolvers: graph.NewResolver(service2)},
	}
}

func (a *App) Run(port string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(a.config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ResponseWriterKey, w)
		ctx = context.WithValue(ctx, utils.RequestKey, r)
		srv.ServeHTTP(w, r.WithContext(ctx))
	}))

	log.Printf("The application is running on %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
