/*
@Author : Mustang
*/
package routers

import (
	"fmt"
	"github.com/flamego/binding"
	"golang-base-flamego/app/handler/auth"
	model "golang-base-flamego/app/models/auth"

	"github.com/flamego/flamego"
	"github.com/spf13/viper"
)

// UserRouter 用户
func UserRouter(g *flamego.Flame) {
	authRouterUser := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/user")
	g.Group(authRouterUser, func() {
		g.Post("", binding.JSON(model.User{}), auth.CreateUserHandler)
		g.Get("", auth.UserListHandler)
		g.Put("/{id}", binding.JSON(model.User{}), auth.UpdateUserHandler)
		g.Delete("/{id}", auth.DeleteUserHandler)
		g.Get("/{id}", auth.UserDetailHandler)
		//g.Get("/:id/dept-user", auth.DeptUserListHandler)
	})
}
