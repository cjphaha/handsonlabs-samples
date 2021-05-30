package gate

import (
    "context"
    "github.com/cjphaha/handsonlabs-samples/game/pkg/pojo"
)

type BasketballService struct {
    Send func(ctx context.Context, uid string, data string) (*pojo.Result, error)
}

func (p *BasketballService) Reference() string {
    return "gateConsumer.basketballService"
}
