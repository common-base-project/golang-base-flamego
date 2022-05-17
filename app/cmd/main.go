/*
  @Author : Mustang Kong
*/

package main

import (
	"fmt"
	"golang-base-flamego/app/models"
	"golang-base-flamego/app/router"
	"golang-base-flamego/pkg/config"
	_ "golang-base-flamego/pkg/config"
	"golang-base-flamego/pkg/connection"
	"golang-base-flamego/pkg/logger"
	"golang-base-flamego/pkg/service/auth_rsync"
	"log"
	"net/http"
	"os"

	"github.com/flamego/flamego"
	"github.com/ory/graceful"
	"github.com/spf13/viper"
)

// init all
func init() {
	fmt.Println("#########1", os.Getenv("ENV_SERVER_MODE"))
	config.EnvMode = os.Getenv("ENV_SERVER_MODE")
	config.Initial()
	fmt.Println("#########2", config.EnvMode)
	logger.Initial()
	connection.Initial()
}

// @title golang-common-base API docs
// @version 0.0.1
// @contact.name Mustang Kong
// @contact.email mustang2247@gmail.com
// http://localhost:9088/api/v1/swagger/index.html
func main() {
	// 同步用户和部门数据
	go auth_rsync.Main()

	// 同步数据结构
	models.AutoMigrateTable()

	f := flamego.New()

	// 加载路由
	router.Load(f)

	// 启动端口
	port := viper.GetString("server.port")
	log.Print("server start at port:", port)

	// 启动服务
	server := graceful.WithDefaults(
		&http.Server{
			Addr:    port,
			Handler: f,
		},
	)
	log.Print(flamego.Env())
	log.Println("main: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		// 处理错误
	}
	log.Println("main: Server was shutdown gracefully")

	// log.Println("Server is running...")
	// log.Println(http.ListenAndServe(viper.GetString(`server.port`), f))

	// g := gin.New()

	// if config.EnvMode == "prod" {
	// 	gin.SetMode(gin.ReleaseMode)
	// } else if config.EnvMode == "staging" {
	// 	gin.SetMode(gin.TestMode)
	// } else {
	// 	gin.SetMode(gin.DebugMode)
	// }

	// // 加载路由
	// router.Load(g)

	// // 运行程序
	// err := g.Run(viper.GetString(`server.port`))
	// if err != nil {
	// 	logger.Error("启动失败")
	// 	panic(fmt.Sprintf("程序启动失败：%v", err))
	// }

	defer connection.DB.Close()
}
