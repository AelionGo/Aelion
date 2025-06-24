// Package ping 健康检测逻辑
package ping

import (
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/cloudwego/hertz/pkg/app"
)

type PingLogic struct {
	ctx    *app.RequestContext
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx *app.RequestContext, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(_ *types.PingRequest) (resp *types.PingResponse, err error) {
	resp = &types.PingResponse{
		Version: "0.1.0",
	}
	return resp, nil
}
