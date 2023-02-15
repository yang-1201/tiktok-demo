package authenticate

import (
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	ReqUserInfoKey = "auth_user_info"
)

type UserInfo struct {
	Username string
<<<<<<< HEAD
=======
	UserID   int64
>>>>>>> f3bcb08 (publish完成)
}

func GetAuthUserInfo(c *app.RequestContext) (*UserInfo, error) {
	userInfo, found := c.Get(ReqUserInfoKey)
	if !found {
		return nil, errors.New("cannot get userInfo from context")
	}
	return userInfo.(*UserInfo), nil
}
