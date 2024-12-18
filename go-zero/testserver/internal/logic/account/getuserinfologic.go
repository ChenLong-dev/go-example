package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testserver/common/errorx"
	"testserver/model/user"

	"testserver/internal/svc"
	"testserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	//如果认证通过 可以从ctx中获取jwt payload
	//userId, err := l.ctx.Value("userId").(json.Number).Int64()
	//if err != nil {
	//	return nil, errorx.TokenInvalid
	//}
	resp = &types.UserInfoResp{}
	userId := req.Id
	u, err := user.NewUserModel(l.svcCtx.Conn).FindOne(l.ctx, userId)
	if err != nil && (errors.Is(err, user.ErrNotFound) ||
		errors.Is(err, sql.ErrNoRows)) {
		resp.Code = errorx.UserNotExist.Code
		resp.Msg = fmt.Sprintf("%s: %v", errorx.UserNotExist.Msg, err)
		return resp, nil
	}
	resp.Code = errorx.Success.Code
	resp.Msg = errorx.Success.Msg
	resp.Id = userId
	resp.Username = u.Username
	return
}
