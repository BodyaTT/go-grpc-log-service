package server

import (
	"fmt"
	"net"

	"github.com/BodyaTT/go-grpc-log-service/pkg/domain"

	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv     *grpc.Server
	auditServer domain.AuditServiceServer
}

func New(auditServer domain.AuditServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	domain.RegisterAuditServiceServer(s.grpcSrv, s.auditServer)

	return s.grpcSrv.Serve(lis)
}
