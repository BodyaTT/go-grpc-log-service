package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BodyaTT/go-grpc-log-service/internal/config"
	"github.com/BodyaTT/go-grpc-log-service/internal/repository"
	"github.com/BodyaTT/go-grpc-log-service/internal/server"
	"github.com/BodyaTT/go-grpc-log-service/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditServer := server.NewAuditServer(auditService)

	srv := server.New(auditServer)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
