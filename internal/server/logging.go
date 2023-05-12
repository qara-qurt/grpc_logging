package server

import (
	"context"
	"github.com/qara-qurt/grpc_logging/internal/service"
	log "github.com/qara-qurt/grpc_logging/proto/logging"
)

type LoggingService interface {
	Insert(ctx context.Context, req *log.LogRequest) error
}

type LoggingServer struct {
	log.UnimplementedLoggingServiceServer
	service LoggingService
}

func NewLoggingServer(service *service.Service) *LoggingServer {
	return &LoggingServer{}
}

func (l *LoggingServer) Log(ctx context.Context, request *log.LogRequest) (*log.LogResponse, error) {
	panic("implement me")
}
