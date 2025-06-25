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

func CaptchaHandler(svcCtx *svc.ServiceContext) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		req := &types.CaptchaRequest{}
		req.Type = ctx.Query("type")
		if req.Type != "1" && req.Type != "2" {
			ctx.JSON(consts.StatusOK, msg.GetResponse(errors.ParamsError, nil))
			return
		}

		l := logic.NewCaptchaLogic(ctx, svcCtx)
		resp, _ := l.Captcha(req)
		ctx.JSON(consts.StatusOK, resp)
	}
}
