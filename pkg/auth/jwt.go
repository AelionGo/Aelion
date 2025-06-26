package auth

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
)

func GetJwtToken(secret string, iat, exp int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat       // 签发时间
	claims["exp"] = iat + exp // 过期时间
	claims["uid"] = uid       // 用户ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJwtToken(secret, tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}

func GetUId(c *app.RequestContext) (string, error) {
	uid, ok := c.Get("uid")
	if !ok {
		return "", jwt.NewValidationError("uid not found in context", jwt.ValidationErrorClaimsInvalid)
	}
	uidStr, ok := uid.(string)
	if !ok {
		return "", jwt.NewValidationError("uid is not a string", jwt.ValidationErrorClaimsInvalid)
	}

	return uidStr, nil
}
