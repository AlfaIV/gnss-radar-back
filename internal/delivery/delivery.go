package delivery

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/Gokert/gnss-radar/internal/delivery/graphql"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	authorization "github.com/Gokert/gnss-radar/internal/service"
	"log"
	"net/http"
)

type App struct {
	config generated.Config
}

func NewApp(auth authorization.IAuthorizationService) *App {
	return &App{
		config: generated.Config{Resolvers: graph.NewResolver(auth)},
	}
}

func (a *App) Run(port string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(a.config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("The application is running on %s", port)
	if err := http.ListenAndServe(":"+port, srv); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
