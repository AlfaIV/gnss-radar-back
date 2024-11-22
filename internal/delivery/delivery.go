package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gnss_radar "github.com/Gokert/gnss-radar/gen/go/api/proto/gnss-radar"
	graph "github.com/Gokert/gnss-radar/internal/delivery/graphql"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/service"
	grpc_go "google.golang.org/grpc"
)

type App struct {
	config          generated.Config
	mux             http.ServeMux
	hardwareService service.IHardware
	httpServer      http.Server
}

func NewApp(service2 *service.Service, hardwareService service.IHardware) *App {
	return &App{
		config:          generated.Config{Resolvers: graph.NewResolver(service2)},
		mux:             *http.NewServeMux(),
		hardwareService: hardwareService,
	}
}

type GnssGrpc struct {
	grpcServ *grpc_go.Server
}

type server struct {
	gnss_radar.UnimplementedGnssServiceServer
}

func NewServer() (*GnssGrpc, error) {
	s := grpc_go.NewServer()
	gnss_radar.RegisterGnssServiceServer(s, &server{})

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

func (a *App) HardwareHandlers(port string) error {
	a.mux.Handle("/hardware/spectrum", http.HandlerFunc(a.AddSpectrum))
	a.mux.Handle("/hardware/power", http.HandlerFunc(a.AddPower))

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w port is NaN", err)
	}
	portStr := strconv.Itoa(portNum + 1)
	log.Printf("Port %s", portStr)
	a.httpServer = http.Server{
		Addr:    ":" + portStr,
		Handler: &a.mux,
	}
	if err := a.httpServer.ListenAndServe(); err != nil {
		log.Printf("http.ListenAndServe: %s", err.Error())
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}
	log.Printf("The hardware application is running on %s", portStr)

	return nil
}

func (a *App) AddSpectrum(w http.ResponseWriter, r *http.Request) {
	var req model.SpectrumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.hardwareService.AddSpectrum(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) AddPower(w http.ResponseWriter, r *http.Request) {
	var req model.PowerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.hardwareService.AddPower(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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

func (s *server) GetStatus(ctx context.Context, req *gnss_radar.GetStatusRequest) (*gnss_radar.GetStatusResponse, error) {
	return &gnss_radar.GetStatusResponse{Result: req.GetIsActual()}, nil
}
