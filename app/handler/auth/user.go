/*
@Author : Mustang
*/
package auth

import (
	"golang-base-flamego/app/models/auth"
	"golang-base-flamego/pkg/connection"
	"golang-base-flamego/pkg/curd"
	"golang-base-flamego/pkg/pagination"
	"golang-base-flamego/pkg/response/code"
	resultResp "golang-base-flamego/pkg/response/response"

	"github.com/flamego/flamego"
)

// CreateUserHandler 创建用户 注入auth.User
func CreateUserHandler(User auth.User) string {
	//var User auth.User
	//f.Map(User)

	models := auth.User{}

	if err := curd.Create(&curd.Param{
		Name:     "用户",
		Models:   &models,
		Param:    &User,
		WhereMap: map[string]interface{}{"username": User.Username},
	}); err != nil {
		return resultResp.Response(code.CreateUserError, nil, err.Error())
	}

	return resultResp.Response(code.Success, User, "成功创建用户!")
}

func UpdateUserHandler(User auth.User) {
	models := auth.User{}
	if err := curd.Update(&curd.Param{
		Name:       "用户",
		Models:     &models,
		Param:      &User,
		WhereValue: User.Username,
	}); err != nil {
		resultResp.Response(code.UpdateUserError, nil, err.Error())
	}

	resultResp.Response(code.Success, User, "成功更新用户信息!")
}

func DeleteUserHandler(c flamego.Context) string {
	var User auth.User
	GID := c.Param("id")

	if err := connection.DB.Self.Delete(&auth.User{}, "id = ?", GID).Error; err != nil {
		return resultResp.Response(code.DeleteUserError, nil, err.Error())
	}

	return resultResp.Response(code.Success, User, "成功删除用户!")
}

func UserListHandler(c flamego.Context) string {
	var data auth.User
	var userList []*auth.User
	result, err := pagination.Paging(&pagination.Param{
		DB: connection.DB.Self,
	}, data, &userList)

	if err != nil {
		return resultResp.Response(code.SelectUserError, nil, err.Error())
	}

	return resultResp.Response(nil, result, "成功获取用户列表")
}

func UserDetailHandler(c flamego.Context) string {
	var User auth.User
	userID := c.Param("id")
	err := connection.DB.Self.Where("id = ?", userID).Find(&User).Error
	if err != nil {
		return resultResp.Response(code.SelectUserError, nil, err.Error())
	}

	return resultResp.Response(nil, User, "获取用户详细信息")

}

//// 获取用户部门对应的所有用户
//func DeptUserListHandler(c *gin.Context) {
//	var (
//		err          error
//		userInfo     auth.User
//		deptUserList []auth.User
//	)
//
//	username := c.DefaultQuery("username", "")
//	if username == "" {
//		username = c.GetString("user")
//	}
//
//	err = connection.DB.Self.Model(&auth.User{}).
//		Where("username = ?", username).
//		Find(&userInfo).Error
//	if err != nil {
//		resultResp.Response(code.SelectUserError, nil, err.Error())
//		return
//	}
//
//	err = connection.DB.Self.Model(&auth.User{}).
//		Where("depart_id = ? and username != ?", userInfo.Depart, username).
//		Find(&deptUserList).Error
//	if err != nil {
//		resultResp.Response(code.SelectUserError, nil, err.Error())
//		return
//	}
//
//	resultResp.Response(nil, deptUserList, "")
//}
