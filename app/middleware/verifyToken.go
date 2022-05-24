package middleware

import (
	"github.com/flamego/flamego"
	"golang-base-flamego/pkg/utils"
	"regexp"
)

var Token = utils.AccessToken{}

func CheckToken() func(context flamego.Context) {
	return func(context flamego.Context) {
		if !isMustApi(context) {
			return
		}

		// 获取头部的 token
		tokenString := context.Request().Header.Get(utils.TokenNameInHeader) // token
		requestId := context.Request().Header.Get(utils.RequestID)           // logID
		Token.RequestID = requestId

		if b := Token.ValidateToken(context, tokenString); b {
			context.Next()
			return
		} else {
			context.ResponseWriter().Write([]byte("{\"errno\": 100001, \"errmsg\": \"access token无效\"}"))
			return
		}
	}
}

// 定义无需登陆检测的接口
func isMustApi(context flamego.Context) bool {
	path := context.Request().RequestURI
	return path != "/api/v1/login" &&
		path != "/api/v1/logout" &&
		!ignoreMatchErr(`/api/v1/reservations/([0-9]+)/checkin`, path) &&
		!ignoreMatchErr(`/api/v1/upload`, path)
}

func ignoreMatchErr(pattern, str string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}
