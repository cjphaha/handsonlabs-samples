package pkg

import (
	"context"
	dubbo3 "github.com/dubbogo/triple/pkg/triple"
	pb "triple-grpc/protobuf"
)

type GrpcGreeterConsumer struct {
	Dubbo3Hello func(ctx context.Context, in *pb.Dubbo3HelloRequest, out *pb.Dubbo3HelloReply) error
}

func (u *GrpcGreeterConsumer) Reference() string {
	return "GrpcGreeterImpl"
}

func (u *GrpcGreeterConsumer) GetDubboStub(tc *dubbo3.TripleConn) pb.Dubbo3GreeterClient {
	return pb.NewDubbo3GreeterDubbo3Client(tc)
}