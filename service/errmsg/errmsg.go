package errmsg

import (
	"errors"
)

var (
	InternalError = errors.New("内部错误，请稍后重试")

	UserNotFoundError             = errors.New("用户不存在")
	UserEmailExistError           = errors.New("邮箱已被使用")
	UserEmailOrPasswordWrongError = errors.New("用户邮箱或密码错误")

	PostNotFoundError = errors.New("文章不存在")

	TokenWrongError      = errors.New("token 不正确")
	TokenParseWrongError = errors.New("token 解析错误")
	TokenNotFoundError   = errors.New("token 没有找到")

	TermNotFoundError = errors.New("分类或标签找不到")
)
