package logic

import (
	"context"

	"go-zero/greet/internal/svc"
	"go-zero/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (*types.Response, error) {
	message := "Hello, " + req.Name + "!"
	return &types.Response{
		Message: message,
	}, nil
}
