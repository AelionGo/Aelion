// Package captcha 提供图形验证码的生成和验证
package captcha

import (
	"github.com/mojocn/base64Captcha"
	"golang.org/x/image/colornames"
)

var store = base64Captcha.DefaultMemStore

func Generate() (string, string, error) {
	driver := base64Captcha.NewDriverMath(80, 240, 5, 80, &colornames.White, nil, []string{})
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	if err != nil {
		return "", "", err
	}
	// 将答案存储在内存中
	err = store.Set(id, answer)
	if err != nil {
		return "", "", err
	}

	return id, b64s, nil
}

func Verify(id, answer string) bool {
	return store.Verify(id, answer, true)
}
