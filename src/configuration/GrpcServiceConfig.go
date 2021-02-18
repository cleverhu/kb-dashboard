package configuration

import (
	"google.golang.org/grpc"
	"knowledgeBase/src/grpcServices"
	"log"
)

type GrpcServiceConfig struct {
}

func NewGrpcServiceConfig() *GrpcServiceConfig {
	return &GrpcServiceConfig{}
}

func (this *GrpcServiceConfig) KbInfoClient() grpcServices.KbInfoServiceClient {
	conn, err := grpc.Dial("dashboard-grpc-server-svc:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return grpcServices.NewKbInfoServiceClient(conn)
}
