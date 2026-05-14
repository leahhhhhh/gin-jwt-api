/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:27:14
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-14 15:07:15
 * @FilePath: /jw_go/middleware/jwt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"gin-demo/config"
	"gin-demo/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			utils.Fail(c, "未登录，请先登录")
			c.Abort()
			return
		}
		if len(auth) < 7 || auth[:7] != "Bearer " {
			utils.Fail(c, "token格式错误")
			c.Abort()
			return
		}
		tokenStr := auth[7:]
		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			utils.Fail(c, "token无效或已过期")
			c.Abort()
			return
		}
		c.Next()
	}
}
