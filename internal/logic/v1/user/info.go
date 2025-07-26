package user

import (
	"github.com/AelionGo/Aelion/internal/svc"
	types "github.com/AelionGo/Aelion/internal/types/user"
	"github.com/AelionGo/Aelion/models"
	"github.com/AelionGo/Aelion/pkg/errors"
	"github.com/AelionGo/Aelion/pkg/msg"
	"github.com/cloudwego/hertz/pkg/app"
)

type InfoLogic struct {
	ctx    *app.RequestContext
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx *app.RequestContext, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoRequest) (resp *msg.Response, err error) {
	//校验参数
	if req.Id == "" || len(req.Id) > 128 {
		resp = msg.GetResponse(errors.ParamsError, nil)
		return resp, nil
	}

	u := models.NewUserModel()
	//验证权限
	id, ok := l.ctx.Get("uid")
	if !ok {
		resp = msg.GetResponse(errors.PermissionDenied, nil)
		return resp, nil
	}
	if id != req.Id {
		//验证是否为管理员
		user, err := u.GetOneByID(id.(string))
		if err != nil {
			resp = msg.GetResponse(errors.DatabaseError, nil)
			return resp, err
		}
		g := models.NewGroupModel()
		group, err := g.GetOneByID(user.Group)
		if err != nil {
			resp = msg.GetResponse(errors.DatabaseError, nil)
			return resp, err
		}
		if group.Type != models.GroupTypeAdmin {
			resp = msg.GetResponse(errors.PermissionDenied, nil)
			return resp, nil
		}
	}

	user, err := u.GetOneByID(req.Id)
	if err != nil {
		resp = msg.GetResponse(errors.UserNotFound, nil)
		return resp, err
	}

	res := msg.GetResponse(errors.OK, &types.InfoResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		Group:    user.Group,
	})
	return res, nil
}
