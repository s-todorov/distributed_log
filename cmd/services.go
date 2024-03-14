package cmd

import (
	"distributed_log/cmd/config"
	"distributed_log/internal/common/server/http"
	"distributed_log/internal/event"
	"distributed_log/internal/log"
	"distributed_log/internal/log/ports"
	"distributed_log/internal/log/service"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/go-multierror"
	"log/slog"
	"time"
)

// ServicesConfiguration is an alias for a function that will take in a pointer to an Services and modify it
type ServicesConfiguration func(os *Services) error

// Services is a implementation of the Services
type Services struct {
	IndexRepo    event.Store
	IndexService *service.IndexService
	EventHandler *event.Handler

	Logs   []*log.DistributedLog
	Router *chi.Mux
	Routes ports.Routes
	Grpc   *ports.GrpcServer
}

// NewServices takes a variable amount of ServicesConfiguration functions and returns a new Services
// Each ServicesConfiguration will be called in the order they are passed in
func NewServices(cfgs ...ServicesConfiguration) (*Services, error) {

	os := &Services{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithApiService(conf config.Config) ServicesConfiguration {
	return func(os *Services) error {
		os.Router = http.NewRouter()
		os.Routes = ports.NewRoutes(os.Router, os.EventHandler)
		return nil
	}
}

func WithGrpcService(conf config.Config) ServicesConfiguration {
	return func(os *Services) error {
		grpcServer, err := ports.NewGrpcServer(os.EventHandler, 5002)
		if err != nil {
			return err
		}
		os.Grpc = grpcServer

		return nil
	}
}

func WithIndexService(conf config.Config) ServicesConfiguration {
	return func(os *Services) error {
		os.IndexService = service.NewIndexService(os.IndexRepo)
		os.EventHandler = event.NewHandler(os.IndexService)

		//os.eventHandler.EventBus.Subscribe(string(event.EventInsert), os.IndexService)
		return nil
	}
}

//func WithLogStore(conf config.Config) ServicesConfiguration {
//	return func(os *Services) error {
//		store, err := log.NewIndexStore("C:\\Git\\github.com\\elastic_middleware\\log")
//		if err != nil {
//			slog.Error(err.Error(), err)
//		}
//		os.EventHandler.WithStoreLog(store)
//		return nil
//	}
//}

func WithDsLog(conf log.Config, dataDir string, id string) ServicesConfiguration {
	return func(os *Services) error {

		l, err := log.NewDistributedLog(dataDir, id, conf)
		if err != nil {
			slog.Error(err.Error(), err)
		}
		if len(os.Logs) == 0 {
			err := l.WaitForLeader(3 * time.Second)
			slog.Error(err.Error(), err)
		} else {
			err = os.Logs[0].Join(
				fmt.Sprintf("%d", id), conf.Raft.BindAddr,
			)
		}

		os.Logs = append(os.Logs, l)
		return nil
	}
}

func (os *Services) Stop() {

	err := os.EventHandler.Stop()

	var merr *multierror.Error
	if errors.As(err, &merr) {
		for _, er := range merr.Errors {
			slog.Error(er.Error(), er)
		}
	}

	os.Grpc.GracefulStop()

	os.EventHandler.EventBus.Stop()
}
