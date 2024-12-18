package logic

import (
	"context"

	"hello02/internal/svc"
	"hello02/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Hello02Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHello02Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Hello02Logic {
	return &Hello02Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Hello02Logic) Hello02(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{
		Message: "xxxxxxx",
	}
	logx.Info("hello02: ", req)
	logx.WithContext(l.ctx).Info("WithContext hello02: ", req)
	return
}
