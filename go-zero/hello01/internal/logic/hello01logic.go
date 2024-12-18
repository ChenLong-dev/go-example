package logic

import (
	"context"

	"hello01/internal/svc"
	"hello01/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Hello01Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHello01Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Hello01Logic {
	return &Hello01Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Hello01Logic) Hello01(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{
		Name:    req.Name,
		Message: "hello01",
	}
	return
}
