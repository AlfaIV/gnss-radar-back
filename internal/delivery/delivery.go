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
	authService authorization.IAuthorizationService
}

func NewApp(auth authorization.IAuthorizationService) *App {
	return &App{
		authService: auth,
	}
}

func (a *App) Run(port string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("The application is running on %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
