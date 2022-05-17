package result

/*
  @Author : Mustang Kong
*/

import (
	"encoding/json"
	"fmt"
	"golang-base-flamego/pkg/logger"
	httpCode "golang-base-flamego/pkg/response/code"
)

type ResponseData struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func Response(err error, data interface{}, resultText string) string {
	code, message := httpCode.DecodeErr(err)

	if err != nil {
		message = fmt.Sprintf("%s，错误：%v", message, resultText)
	}

	if err == nil && resultText != "" {
		message = resultText
	}

	// write log
	if code != httpCode.Success.Errno {
		logger.Error(message)
	}

	res := ResponseData{
		Errno:  code,
		Errmsg: message,
		Data:   data,
	}

	json, _ := json.Marshal(res)

	// always return http.StatusOK
	return string(json)
	// c.JSON(http.StatusOK, ResponseData{
	// 	Errno:  code,
	// 	Errmsg: message,
	// 	Data:   data,
	// })
}
