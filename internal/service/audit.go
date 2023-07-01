package service

import (
	"context"

	"github.com/BodyaTT/go-grpc-log-service/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAudit(repo Repository) *Audit {
	return &Audit{repo: repo}
}

func (a *Audit) Insert(ctx context.Context, req *domain.LogRequest) error {
	item := domain.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return a.repo.Insert(ctx, item)
}
