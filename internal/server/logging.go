package server

import (
	"context"
	log "grpc_logging/proto/logging"
)

//type LoggingService interface {
//	Insert(ctx context.Context, req *log.LogRequest) error
//}

type LoggingServer struct {
}

func NewLoggingServer() *LoggingServer {
	return &LoggingServer{}
}

func (h *LoggingServer) Log(ctx context.Context, request *log.LogRequest) (*log.LogResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *LoggingServer) mustEmbedUnimplementedLoggingServiceServer() {
	//TODO implement me
	panic("implement me")
}
