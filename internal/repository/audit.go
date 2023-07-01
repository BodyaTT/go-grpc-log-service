package repository

import (
	"context"

	"github.com/BodyaTT/go-grpc-log-service/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{
		db: db,
	}
}

func (a *Audit) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := a.db.Collection("logs").InsertOne(ctx, item)

	return err
}
