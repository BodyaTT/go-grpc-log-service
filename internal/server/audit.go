package server

import (
	"context"

	"github.com/BodyaTT/go-grpc-log-service/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *domain.LogRequest) error
}

type AuditServer struct {
	service AuditService
	domain.UnimplementedAuditServiceServer
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *domain.LogRequest) (*domain.Empty, error) {
	err := h.service.Insert(ctx, req)

	return &domain.Empty{}, err
}
