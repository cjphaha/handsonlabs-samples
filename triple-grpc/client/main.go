package main

import (
	"github.com/apache/dubbo-go/config"
	pb "triple-grpc/protobuf"
	"triple-grpc/client/pkg"
	"context"
	"fmt"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo3"
	_ "github.com/apache/dubbo-go/protocol/grpc"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

var grpcGreeterConsumer = new(pkg.GrpcGreeterConsumer)

func main( )  {
	config.SetConsumerService(grpcGreeterConsumer)
	config.Load()
	test()
}

func test(){
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