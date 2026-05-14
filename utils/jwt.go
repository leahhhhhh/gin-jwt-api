/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:25:38
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-14 15:06:26
 * @FilePath: /jw_go/utils/jwt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"time"

	"gin-demo/config"

	"github.com/golang-jwt/jwt/v4"
)

func GenToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}
