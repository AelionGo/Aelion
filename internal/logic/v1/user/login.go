package user

import (
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/internal/types"
	"github.com/AelionGo/Aelion/models"
	"github.com/AelionGo/Aelion/pkg/auth"
	"github.com/AelionGo/Aelion/pkg/captcha"
	"github.com/AelionGo/Aelion/pkg/consts"
	"github.com/AelionGo/Aelion/pkg/errors"
	"github.com/AelionGo/Aelion/pkg/hash"
	"github.com/AelionGo/Aelion/pkg/msg"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type LoginLogic struct {
	ctx    *app.RequestContext
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx *app.RequestContext, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *msg.Response, err error) {
	//校验参数
	if (req.Email == "" && req.Phone == "") || len(req.Email) > consts.MaxEmailLength || len(req.Phone) > consts.MaxPhoneLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}
	if req.Password == "" || len(req.Password) < consts.MinPasswordLength || len(req.Password) > consts.MaxPasswordLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}

	//获取用户密码
	u := models.NewUserModel()
	var id string
	var hashed string
	if req.Email != "" {
		usr, err := u.GetOneByEmail(req.Email)
		if err != nil {
			resp = msg.GetResponse(errors.UserNotFound, nil)
			return resp, err
		}
		id = usr.Id
		hashed = usr.Password
	} else {
		usr, err := u.GetOneByEmail(req.Phone)
		if err != nil {
			resp = msg.GetResponse(errors.UserNotFound, nil)
			return resp, err
		}
		id = usr.Id
		hashed = usr.Password
	}

	//验证图形验证码
	enabled, err := l.svcCtx.Config.RegisterCaptchaEnabled()
	if err != nil {
		resp = msg.GetResponse(errors.GetConfigItemError, nil)
		return resp, err
	}
	if enabled {
		if !captcha.Verify(req.CaptchaId, req.CaptchaAnswer) {
			resp = msg.GetResponse(errors.CaptchaVerifyError, nil)
			return resp, nil
		}
	}

	//验证密码
	ok := hash.CheckPasswordHash(req.Password, hashed)
	if !ok {
		resp = msg.GetResponse(errors.PasswordError, nil)
		return resp, nil
	}

	//验证用户是否激活
	enabled, err = l.svcCtx.Config.EmailValidationEnabled()
	if err != nil {
		resp = msg.GetResponse(errors.GetConfigItemError, nil)
		return resp, err
	}
	if enabled {
		// TODO: 检查用户激活状态，重发激活邮件

	}

	//生成token
	secret, err := l.svcCtx.Config.JwtSecret()
	if err != nil {
		resp = msg.GetResponse(errors.GetConfigItemError, nil)
		return resp, err
	}
	token, err := auth.GetJwtToken(secret, time.Now().Unix(), 86400, id)
	if err != nil {
		resp = msg.GetResponse(errors.TokenGenerateError, nil)
		return resp, err
	}
	resp = msg.GetResponse(errors.OK, &types.LoginResponse{
		Id:    id,
		Token: token,
	})
	return resp, nil
}
