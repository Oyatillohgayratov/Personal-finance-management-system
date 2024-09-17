package main

import (
	"context"
	"log"
	"net"
	"notifaction-service/internal/config"
	"notifaction-service/internal/kafka"
	notifactionpb "notifaction-service/protos/notifaction"
	"notifaction-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	consumer := kafka.ConnKafka(cfg)

	notificationServer := service.NewNotificationServer(consumer)
	go notificationServer.ListenForKafkaMessages()

	lis, err := net.Listen("tcp", cfg.NotifactionServer.Http.Host+":"+cfg.NotifactionServer.Http.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(grpcServer)

	notifactionpb.RegisterNotificationServiceServer(grpcServer, notificationServer)

	log.Println("main: server running  port", cfg.NotifactionServer.Http.Port)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("failed while listening:", (err))
	}

}

func grpcInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
		return nil, err
	}
	return m, nil
}
