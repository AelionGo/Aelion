package auth

import (
	"context"
	"github.com/AelionGo/Aelion/internal/svc"
	authx "github.com/AelionGo/Aelion/pkg/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(svcCtx *svc.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := string(c.Request.Header.Peek("Authorization"))
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "Missing or invalid Authorization header",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret, err := svcCtx.Config.JwtSecret()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve JWT secret",
			})
			return
		}
		claims, err := authx.ParseJwtToken(secret, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token",
			})
			return
		}

		// 将用户信息设置到上下文中
		c.Set("uid", claims["uid"])

		c.Next(ctx)
	}
}
