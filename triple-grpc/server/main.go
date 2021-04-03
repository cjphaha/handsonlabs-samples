package main

import (
	"fmt"
	"github.com/apache/dubbo-go/common/logger"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo3"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
	"os"
	"os/signal"
	"syscall"
	"time"
	"triple-grpc/server/pkg"
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
