package grpc

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	articlev1 "github.com/mahdimehrabi/m1-article-proto/gen/go/article/article"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"m1-article-service/application/grpc/server"
	"m1-article-service/domain/repository/article/pgx"
	"m1-article-service/domain/service/article"
	"m1-article-service/infrastructure/godotenv"
	"m1-article-service/infrastructure/log/zerolog"
	"net"
)

func Boot() {
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()

	conn, err := pgxpool.New(context.Background(), env.DATABASE_HOST)
	if err != nil {
		log.Fatal(err)
	}

	articleRepo := pgx.NewArticleRepository(env, conn)
	if err != nil {
		log.Fatal(err)
	}
	loggerService := article.NewService(logger, articleRepo)

	lis, err := net.Listen("tcp", env.ServerAddr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	articleServer := server.NewArticleServer(logger, loggerService)
	articlev1.RegisterArticleServiceServer(grpcServer, articleServer)

	reflection.Register(grpcServer)
	logger.Info(fmt.Sprintf("running grpc server on: %s ⛴", env.ServerAddr))
	err = grpcServer.Serve(lis)
	log.Fatal(err)
}
