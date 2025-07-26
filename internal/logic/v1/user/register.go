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
	"github.com/google/uuid"
	"time"
)

type RegisterLogic struct {
	ctx    *app.RequestContext
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx *app.RequestContext, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *msg.Response, err error) {
	//校验参数
	if (req.Email == "" && req.Phone == "") || len(req.Email) > consts.MaxEmailLength || len(req.Phone) > consts.MaxPhoneLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}
	if req.Password == "" || len(req.Password) < consts.MinPasswordLength || len(req.Password) > consts.MaxPasswordLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}
	if req.Nickname == "" || len(req.Nickname) > consts.MaxNicknameLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}
	if len(req.Avatar) > consts.MaxAvatarLength {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}

	//验证邮箱或手机号是否已存在
	u := models.NewUserModel()
	if req.Email != "" {
		_, err := u.GetOneByEmail(req.Email)
		if err == nil {
			resp = msg.GetResponse(errors.EmailExists, nil)
			return resp, nil
		}
	}
	if req.Phone != "" {
		_, err := u.GetOneByPhone(req.Phone)
		if err == nil {
			resp = msg.GetResponse(errors.PhoneExists, nil)
			return resp, nil
		}
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

	hashed, err := hash.HashPassword(req.Password)
	if err != nil {
		resp = msg.GetResponse(errors.PasswordHashError, nil)
		return resp, err
	}

	user := &models.User{
		Id:       uuid.New().String(),
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashed,
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
	}
	if req.Group != "" {
		g := models.NewGroupModel()
		_, err := g.GetOneByID(req.Group)
		if err != nil {
			resp = msg.GetResponse(errors.ParamsError, nil)
			return resp, err
		}
		user.Group = req.Group
	} else {
		group, err := l.svcCtx.Config.DefaultGroup()
		if err != nil {
			resp = msg.GetResponse(errors.GetConfigItemError, nil)
			return resp, err
		}
		user.Group = group
	}

	enabled, err = l.svcCtx.Config.EmailValidationEnabled()
	if err != nil {
		resp = msg.GetResponse(errors.GetConfigItemError, nil)
		return resp, err
	}
	if enabled {
		user.Status = models.UserNeedActivation

		// TODO: 发送激活邮件，并且定期删除未激活用户

	} else {
		user.Status = models.UserActive
	}

	err = u.Create(user)
	if err != nil {
		resp = msg.GetResponse(errors.DatabaseError, nil)
		return resp, err
	}

	//生成token
	secret, err := l.svcCtx.Config.JwtSecret()
	if err != nil {
		resp = msg.GetResponse(errors.GetConfigItemError, nil)
		return resp, err
	}
	token, err := auth.GetJwtToken(secret, time.Now().Unix(), 86400, user.Id)
	if err != nil {
		resp = msg.GetResponse(errors.TokenGenerateError, nil)
		return resp, err
	}
	resp = msg.GetResponse(errors.OK, &types.RegisterResponse{
		Id:    user.Id,
		Token: token,
	})
	return resp, nil
}
