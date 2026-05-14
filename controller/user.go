/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-14 11:29:49
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-14 15:07:27
 * @FilePath: /jw_go/controller/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"gin-demo/model"
	"gin-demo/service"
	"gin-demo/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}
	err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		utils.Fail(c, "注册失败")
		return
	}
	utils.Success(c, "注册成功")
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}
	token, err := service.LoginUser(req.Username, req.Password)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, gin.H{"token": token})
}

func UserInfo(c *gin.Context) {
	utils.Success(c, "这是登录后才能访问的用户信息")
}
