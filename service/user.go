/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:28:19
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-14 15:07:54
 * @FilePath: /jw_go/service/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"errors"
	"gin-demo/config"
	"gin-demo/model"
	"gin-demo/utils"

	"gorm.io/gorm"
)

// 注册用户
func RegisterUser(username, password string) error {
	user := model.User{
		Username: username,
		Password: password,
	}
	return config.DB.Create(&user).Error
}

// 登录并返回token
func LoginUser(username, password string) (string, error) {
	var user model.User
	err := config.DB.Where("username=? AND password=?", username, password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("账号密码错误")
	}
	if err != nil {
		return "", err
	}
	return utils.GenToken(user.ID)
}
