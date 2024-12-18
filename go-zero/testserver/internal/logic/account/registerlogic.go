package account

import (
	"context"
	"fmt"
	"testserver/model/user"
	"time"

	"testserver/common/errorx"
	"testserver/internal/svc"
	"testserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.RegisterResp{}
	userModel := user.NewUserModel(l.svcCtx.Conn)
	u, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("Register FindByUsername err: ", err)
		resp.Code = errorx.UserNotExist.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.UserNotExist.Msg, err)
		return resp, nil
	}
	if u != nil {
		//代表已经注册
		resp.Code = errorx.UserAlreadyExist.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.UserAlreadyExist.Msg, err)
		return resp, nil
	}
	_, err = userModel.Insert(l.ctx, &user.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		resp.Code = errorx.DBError.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.DBError.Msg, err)
		return resp, nil
	}
	resp.Code = errorx.Success.Code
	resp.Msg = errorx.Success.Msg
	return
}
