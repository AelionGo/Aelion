package user

import (
	"context"
	logic "github.com/AelionGo/Aelion/internal/logic/v1/user"
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/AelionGo/Aelion/pkg/errors"
	"github.com/AelionGo/Aelion/pkg/msg"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func LoginHandler(svcCtx *svc.ServiceContext) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		req := &types.LoginRequest{}
		if err := ctx.BindJSON(req); err != nil {
			ctx.JSON(consts.StatusOK, msg.GetResponse(errors.ParamsError, nil))
			return
		}

		l := logic.NewLoginLogic(ctx, svcCtx)
		resp, _ := l.Login(req)
		ctx.JSON(consts.StatusOK, resp)
	}
}
