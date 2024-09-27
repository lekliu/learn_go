package errors

import (
	"demo/src/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeCMSSuccess = 200
	errorCodeCMSFailed  = 401
)

var (
	ErrorCMSFailed = errors.New(
		config.ServiceNameCMS, "操作异常", errorCodeCMSFailed,
	)
	ErrorCMSLogin = errors.New(
		config.ServiceNameCMS, "登录异常", errorCodeCMSFailed,
	)
	ErrorCMSFailedParam = errors.New(
		config.ServiceNameCMS, "参数异常", errorCodeCMSFailed,
	)
	//TODO
)
