package main

import (
	"context"
	"fmt"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/cjphaha/handsonlabs-samples/triple-grpc/client/pkg"
	pb "github.com/cjphaha/handsonlabs-samples/triple-grpc/protobuf"

	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster_impl"
	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/grpc"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	_ "dubbo.apache.org/dubbo-go/v3/registry/nacos"
)

var grpcGreeterConsumer = new(pkg.GrpcGreeterConsumer)

func main() {
	config.SetConsumerService(grpcGreeterConsumer)
	config.Load()
	test()
}

func test() {
	in := &pb.Dubbo3HelloRequest{
		Req: "cjp",
	}
	out := &pb.Dubbo3HelloReply{}
	err := grpcGreeterConsumer.Dubbo3Hello(context.Background(), in, out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
