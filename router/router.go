/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-14 11:30:24
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-14 15:07:46
 * @FilePath: /jw_go/router/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package router

import (
	"gin-demo/controller"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 公共接口
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 需要JWT鉴权
	auth := r.Group("/auth")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/info", controller.UserInfo)
	}

	return r
}
