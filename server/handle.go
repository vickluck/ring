package server

import (
	"context"

	"github.com/tnngo/nog"
	"github.com/tnngo/ring/server/proto/provider"
)

type Handle struct{}

func (h *Handle) Regist(ctx context.Context, in *provider.Request) (*provider.Response, error) {
	nog.SL.Debug(1111111)
	return &provider.Response{}, nil
}
