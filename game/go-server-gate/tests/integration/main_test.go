package integration

import (
    "os"
    "testing"
    "time"

    hessian "github.com/apache/dubbo-go-hessian2"
    "github.com/apache/dubbo-go/config"

    _ "github.com/apache/dubbo-go/protocol/dubbo"
    _ "github.com/apache/dubbo-go/registry/protocol"

    _ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
    _ "github.com/apache/dubbo-go/filter/filter_impl"

    _ "github.com/apache/dubbo-go/cluster/cluster_impl"
    _ "github.com/apache/dubbo-go/cluster/loadbalance"
    // _ "github.com/apache/dubbo-go/metadata/service/remote"
    _ "github.com/apache/dubbo-go/metadata/service/inmemory"
    _ "github.com/apache/dubbo-go/registry/zookeeper"

    "github.com/cjphaha/handsonlabs-samples/game/pkg/consumer/gate"
    "github.com/cjphaha/handsonlabs-samples/game/pkg/pojo"
)

var gateProvider = new(gate.BasketballService)

func TestMain(m *testing.M) {
    config.SetConsumerService(gateProvider)

    hessian.RegisterPOJO(&pojo.Result{})
    config.Load()
    time.Sleep(3 * time.Second)

    os.Exit(m.Run())
}