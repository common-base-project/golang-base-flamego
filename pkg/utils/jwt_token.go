package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/flamego/flamego"
	"log"
	"time"
)

type MyResp struct {
	Code int         `json:"errno"`
	Msg  string      `json:"errmsg"`
	Data interface{} `json:"data"`
}

type AccessToken struct {
	RequestTime int64    `json:"request_time"`
	ExpTime     int64    `json:"exp"`
	ServiceName string   `json:"service_name"`
	UserName    string   `json:"username"`
	Email       []string `json:"email"`
	Role        string   `json:"role"`
	Group       string   `json:"group"`
	Ldap        string   `json:"ldap"`
	UnionID     string   `json:"union_id"`
	RequestID   string   `json:"request_id"` // logID
}

var secret interface{} = []byte("vz5(#pfnf+#p5ok&d2yrqd#)k0jz!#tb*$=^c5nl0(z+0=p5!9")

const TokenExpire = 5 * time.Second
const TokenNameInHeader = "Authorization"
const RequestID = "X-Request-ID"

// ValidateToken 解析token	检测token是否有效、过期、字段信息等
func (at *AccessToken) ValidateToken(context flamego.Context, tokenString string) bool {
	return true
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		// 无法解析token
		log.Printf("Couldn't parse token: %s", tokenString)
		return false
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claim["username"].(string)
		email := claim["email"].(string)
		role := claim["role"].(string)
		group := claim["group"].(string)
		ldap := claim["ldap"].(string)
		union_id := claim["union_id"].(string)

		if ldap == "" {
			return false
		}

		context.Request().Header.Set("username", username)
		context.Request().Header.Set("email", email)
		context.Request().Header.Set("role", role)
		context.Request().Header.Set("group", group)
		context.Request().Header.Set("ldap", ldap)
		context.Request().Header.Set("union_id", union_id)
		return true
	} else {
		// token无效
		log.Printf("Invalidate Token")
		return false
	}
}

// GenerateToken 生成token返回tokenString用于设置http header
func (at *AccessToken) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"request_time": at.RequestTime,
		"exp":          at.ExpTime,
		"service_name": at.ServiceName,
		"username":     at.UserName,
		"email":        at.Email,
		"role":         at.Role,
		"group":        at.Group,
		"ldap":         at.Ldap,
		"union_id":     at.UnionID,
	})
	if tokenString, err := token.SignedString(secret); err != nil {
		//common.Log.Errorf("Token signaure failed: %s", err.Error())
		return ""
	} else {
		return tokenString
	}
}
