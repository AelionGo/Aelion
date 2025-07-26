package user

import (
	"context"
	logic "github.com/AelionGo/Aelion/internal/logic/v1/user"
	"github.com/AelionGo/Aelion/internal/svc"
	types "github.com/AelionGo/Aelion/internal/types/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func InfoHandler(svcCtx *svc.ServiceContext) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		req := &types.InfoRequest{}
		req.Id = ctx.Query("id")

		l := logic.NewInfoLogic(ctx, svcCtx)
		resp, _ := l.Info(req)
		ctx.JSON(consts.StatusOK, resp)
	}
}
