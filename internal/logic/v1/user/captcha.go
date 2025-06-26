package user

import (
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/AelionGo/Aelion/pkg/captcha"
	"github.com/AelionGo/Aelion/pkg/errors"
	"github.com/AelionGo/Aelion/pkg/msg"
	"github.com/cloudwego/hertz/pkg/app"
)

type CaptchaLogic struct {
	ctx    *app.RequestContext
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx *app.RequestContext, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req *types.CaptchaRequest) (resp *msg.Response, err error) {
	if req.Type == "1" {
		enabled, err := l.svcCtx.Config.RegisterCaptchaEnabled()
		if err != nil {
			resp = msg.GetResponse(errors.GetConfigItemError, nil)
			return resp, err
		}
		if !enabled {
			resp = msg.GetResponse(errors.OK, &types.CaptchaResponse{
				Enabled: false,
				Id:      "",
				B64s:    "",
			})
			return resp, nil
		}
	} else {
		enabled, err := l.svcCtx.Config.LoginCaptchaEnabled()
		if err != nil {
			resp = msg.GetResponse(errors.GetConfigItemError, nil)
			return resp, err
		}
		if !enabled {
			resp = msg.GetResponse(errors.OK, &types.CaptchaResponse{
				Enabled: false,
				Id:      "",
				B64s:    "",
			})
			return resp, nil
		}
	}
	id, b64s, err := captcha.Generate()
	if err != nil {
		resp = msg.GetResponse(errors.CaptchaGenerateError, nil)
		return resp, err
	}
	resp = msg.GetResponse(errors.OK, &types.CaptchaResponse{
		Enabled: true,
		Id:      id,
		B64s:    b64s,
	})
	return resp, nil
}
