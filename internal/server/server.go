package server

import (
	"fmt"
	log "github.com/qara-qurt/grpc_logging/proto/logging"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer    *grpc.Server
	loggingServer log.LoggingServiceServer
}

func New(loggingServer log.LoggingServiceServer) *Server {
	return &Server{
		grpcServer:    grpc.NewServer(),
		loggingServer: loggingServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.RegisterLoggingServiceServer(s.grpcServer, s.loggingServer)

	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
