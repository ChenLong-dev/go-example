package account

import (
	"context"
	"fmt"
	"testserver/model/user"
	"time"

	"testserver/common/errorx"
	"testserver/internal/svc"
	"testserver/internal/types"
	"testserver/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	ctx := logx.ContextWithFields(l.ctx, logx.Field("user", req.Username))
	// todo: add your logic here and delete this line
	logx.Info("hello02: ", req)
	logx.WithContext(ctx).Info("WithContext hello02: ", req)

	resp = &types.LoginResp{}
	userModel := user.NewUserModel(l.svcCtx.Conn)
	u, err := userModel.FindByUsernameAndPwd(l.ctx, req.Username, req.Password)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = errorx.UserNotExist.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.UserNotExist.Msg, err)
		return resp, nil
	}
	if u == nil {
		return nil, errorx.NameOrPwdError
	}
	//登录成功 生成token
	secret := l.svcCtx.Config.Auth.Secret
	expire := l.svcCtx.Config.Auth.Expire
	token, err := tools.GetJwtToken(secret, time.Now().Unix(), expire, u.Id)
	if err != nil {
		resp.Code = errorx.TokenError.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.TokenError.Msg, err)
		return resp, nil
	}
	resp.Code = errorx.Success.Code
	resp.Msg = errorx.Success.Msg
	resp.Token = token
	return
}
