package server

import (
	"context"

	"github.com/tnngo/ring/balance"
	"github.com/tnngo/ring/role"
	"github.com/tnngo/ring/server/proto/consumer"
	"github.com/tnngo/ring/server/proto/provider"
)

type Handle struct{}

func (h *Handle) Registry(ctx context.Context, in *provider.Request) (*provider.Response, error) {
	b := balance.New()
	p := &role.Provide{
		MSrvName: in.MsrvName,
		IP:       in.Ip,
		Port:     in.Port,
		ID:       in.Id,
	}
	b.Put(in.MsrvName, p)
	return &provider.Response{}, nil
}

func (h *Handle) Discovery(ctx context.Context, in *consumer.Request) (*consumer.Response, error) {
	return nil, nil
}
