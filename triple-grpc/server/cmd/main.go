package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"

	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster_impl"
	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	_ "dubbo.apache.org/dubbo-go/v3/metadata/service/inmemory"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	_ "dubbo.apache.org/dubbo-go/v3/registry/nacos"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	"github.com/cjphaha/handsonlabs-samples/triple-grpc/server/pkg"
	//_ "dubbo.apache.org/dubbo-go/v3/metadata/service/remote"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// survival time
var (
	survivalTimeout = int(3e9)
)

func main() {
	config.SetProviderService(pkg.NewGreeterProvider())
	config.Load()
	initSignal()
}

// elegant ending procedure
func initSignal() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			fmt.Println("provider app exit now...")
			return
		}
	}
}
