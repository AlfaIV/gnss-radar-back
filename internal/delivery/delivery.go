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
	"github.com/Gokert/gnss-radar/pb/gnss-radar/grpc"
	grpc_go "google.golang.org/grpc"
	"log"
	"net"
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

type GnssGrpc struct {
	grpcServ *grpc_go.Server
}

type server struct {
	grpc.UnimplementedGnssServiceServer
}

func NewServer() (*GnssGrpc, error) {
	s := grpc_go.NewServer()
	grpc.RegisterGnssServiceServer(s, &server{})

	return &GnssGrpc{grpcServ: s}, nil
}

func (a *App) Run(port string) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(a.config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ResponseWriterKey, w)
		ctx = context.WithValue(ctx, utils.RequestKey, r)
		srv.ServeHTTP(w, r.WithContext(ctx))
	}))

	log.Printf("The graphql application is running on %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}

func (s *GnssGrpc) ListenAndServeGrpc(network, port string) error {
	listen, err := net.Listen(network, ":"+port)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	log.Printf("The grpc listener is running on %s", port)
	err = s.grpcServ.Serve(listen)
	if err != nil {
		return fmt.Errorf("grpcServ.Serve: %w", err)
	}

	return nil
}

func (s *server) GetStatus(ctx context.Context, req *grpc.GetStatusRequest) (*grpc.GetStatusResponse, error) {
	return &grpc.GetStatusResponse{Result: req.GetIsActual()}, nil
}
