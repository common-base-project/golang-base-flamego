package router

/*
  @Author : Mustang Kong
*/

import (
	"golang-base-flamego/app/models/auth"
	"golang-base-flamego/app/router/routers"
	_ "golang-base-flamego/docs"
	"golang-base-flamego/pkg/connection"

	"github.com/flamego/cors"
	"github.com/flamego/flamego"
	"github.com/spf13/viper"
)

// Load 加载路由
func Load(g *flamego.Flame) {
	// 404
	g.NotFound(func() string {
		return "API地址不存在"
	})

	g.Get("/",
		cors.CORS(),
		func(c flamego.Context) string {
			return "This endpoint can be shared cross-origin"
		},
	)

	//cors， 跨域
	// config := cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowOrigins:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }
	// 注册zap相关中间件
	// g.Use(cors.New(config))
	//g.Use(logger.GinLogger(), logger.GinRecovery(true))
	// g.Use(utils.CostTime())

	// ========================文件配置===============================
	filePath := viper.GetString("filePath")
	// _, err := tools.CreateDictByPath(filePath)
	// if err != nil {
	// 	log.Error("创建目录失败，请手动创建![%v]\n", err)
	// 	return
	// }
	// log.Infof("创建目录成功: %s", filePath)

	// staticPath := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/upload")
	// 静态文件地址 http://localhost:port/api/v1/upload/fileid.jpg
	g.Use(flamego.Renderer())

	// 全局注入db
	db := connection.DB.Self
	g.Map(db)

	g.Map(auth.User{})

	g.Use(flamego.Recovery())
	g.Use(flamego.Logger())

	g.Use(
		flamego.Static(
			flamego.StaticOptions{
				Directory: filePath,
			},
		))

	// swagger api docs
	//g.Get(fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/swagger/*any"), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// jwt 检查
	// g.Use(middleware.CheckToken())

	// user
	routers.UserRouter(g)

	// email
	routers.EmailRouter(g)
}

func Logger() {
	panic("unimplemented")
}
