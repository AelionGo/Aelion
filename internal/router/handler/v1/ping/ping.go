// Filename: ping.go
// Description: 心跳

package ping

import (
	"context"
	logic "github.com/AelionGo/Aelion/internal/logic/ping"
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PingHandler(svcCtx *svc.ServiceContext) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		req := &types.PingRequest{}

		l := logic.NewPingLogic(ctx, svcCtx)
		resp, _ := l.Ping(req)
		ctx.JSON(consts.StatusOK, resp)
	}
}
