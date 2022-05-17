package routers

/*
   @Author : Mustang Kong
*/

import (
	"fmt"
	"github.com/flamego/flamego"
	"github.com/spf13/viper"
)

// email 路由
func EmailRouter(g *flamego.Flame) {
	routerEmail := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/email")
	g.Group(routerEmail, func() {
		// Upload
		//g.Get("/list", email2.GetEmailListHandler)
		//g.Post("/add", email2.AddEmailHandler)
		//g.Put("/update/:id", email2.UpdateEmailHandler)
		//g.Delete("/delete/:contentId", email2.DeleteEmailHandler)
		//g.Post("/push", email2.AddPushHandler)
	})
}
