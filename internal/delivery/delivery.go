package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"

	"github.com/Gokert/gnss-radar/internal/pkg/middleware"
	"github.com/Gokert/gnss-radar/internal/pkg/pythoncodegen"

	gnss_radar "github.com/Gokert/gnss-radar/gen/go/api/proto/gnss-radar"
	graph "github.com/Gokert/gnss-radar/internal/delivery/graphql"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/service"
	grpc_go "google.golang.org/grpc"
)

type App struct {
	config          generated.Config
	mux             http.ServeMux
	hardwareService service.IHardware
	httpServer      http.Server
	middleware      middleware.IMiddlewareService
}

func NewApp(service2 *service.Service, hardwareService service.IHardware, middleware middleware.IMiddlewareService) *App {
	return &App{
		config:          generated.Config{Resolvers: graph.NewResolver(service2)},
		mux:             *http.NewServeMux(),
		hardwareService: hardwareService,
		middleware:      middleware,
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

	//http.Handle("/", a.middleware.CheckAuthorize(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", a.middleware.CallMiddlewares()(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", a.middleware.SetResponseRequest(srv))

	log.Printf("The graphql application is running on %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}

func (a *App) HardwareHandlers(port string) error {
	a.mux.Handle("/hardware/spectrum", http.HandlerFunc(a.AddSpectrum))
	a.mux.Handle("/hardware/power", http.HandlerFunc(a.AddPower))
	a.mux.Handle("/hardware/upload", http.HandlerFunc(a.UploadSP3))
	a.mux.Handle("/hardware/pair_measurement", http.HandlerFunc(a.AddPairMeasurement))
	a.mux.Handle("/codegen/download", http.HandlerFunc(a.CodeGenDownload))

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w port is NaN", err)
	}
	portStr := strconv.Itoa(portNum + 1)
	log.Printf("Hardware Port %s", portStr)
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

	err := a.hardwareService.CompareDeviceToken(r.Context(), req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.hardwareService.AddSpectrum(r.Context(), req)
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

	err := a.hardwareService.CompareDeviceToken(r.Context(), req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.hardwareService.AddPower(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *App) AddPairMeasurement(w http.ResponseWriter, r *http.Request) {
	var req model.PairMeasurementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.hardwareService.CompareDeviceToken(r.Context(), req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.hardwareService.AddPairMeasurement(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) UploadSP3(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	randomID := uuid.New().String()
	saveDir := filepath.Join("sp3", randomID)
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		savePath := filepath.Join(saveDir, fileHeader.Filename)
		out, err := os.Create(savePath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Failed to copy file", http.StatusInternalServerError)
			return
		}
	}

	err = a.hardwareService.UploadSP3(r.Context(), saveDir)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process files: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Files uploaded and processed successfully")
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

func (a *App) CodeGenDownload(w http.ResponseWriter, r *http.Request) {
	var downloadCodeReq model.CodeRecieverInput
	if err := json.NewDecoder(r.Body).Decode(&downloadCodeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch downloadCodeReq.TypeLang {
	case "python":
		codeFile, err := pythoncodegen.SaveCodeToFile(model.PythonGenConfig{
			BaseURL:   model.HardwareHandlersBaseAddress,
			Token:     downloadCodeReq.Token,
			SampleNum: 256,
		}, downloadCodeReq.Token+"_"+downloadCodeReq.TypeLang)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.py", downloadCodeReq.Token+"_"+downloadCodeReq.TypeLang))
		http.ServeFile(w, r, codeFile.Name())
		os.Remove(codeFile.Name())
	default:
		http.Error(w, "Unsupported code generation language", http.StatusBadRequest)
	}
}

func (s *server) GetStatus(ctx context.Context, req *gnss_radar.GetStatusRequest) (*gnss_radar.GetStatusResponse, error) {
	return &gnss_radar.GetStatusResponse{Result: req.GetIsActual()}, nil
}
