// Package ping 健康检测逻辑
package ping

import (
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/AelionGo/Aelion/pkg/errors"
	"github.com/AelionGo/Aelion/pkg/msg"
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

func (l *PingLogic) Ping(_ *types.PingRequest) (resp *msg.Response, err error) {
	res := &types.PingResponse{
		Version: "0.1.0",
	}
	resp = msg.GetResponse(errors.OK, res)
	return resp, nil
}
