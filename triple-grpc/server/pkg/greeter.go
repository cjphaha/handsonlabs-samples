package pkg

import (
	"context"
	"fmt"
	pb "github.com/cjphaha/handsonlabs-samples/triple-grpc/protobuf"
)

type GreeterProvider struct {
	*pb.Dubbo3GreeterProviderBase
}

func NewGreeterProvider() *GreeterProvider {
	return &GreeterProvider{
		Dubbo3GreeterProviderBase: &pb.Dubbo3GreeterProviderBase{},
	}
}

func (g *GreeterProvider) Dubbo3Hello(ctx context.Context, in *pb.Dubbo3HelloRequest) (*pb.Dubbo3HelloReply, error) {
	fmt.Println("######### get server request data :", in.Req)
	fmt.Println("get tri-req-id = ", ctx.Value("tri-req-id"))
	return &pb.Dubbo3HelloReply{Rsp: "Hello " + in.Req}, nil
}

func (g *GreeterProvider) Reference() string {
	return "GrpcGreeterImpl"
}
