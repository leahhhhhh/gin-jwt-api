/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:24:25
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-13 18:25:22
 * @FilePath: /jw_go/utils/result.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import "github.com/gin-gonic/gin"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Result{
		Code: 200,
		Msg:  "成功",
		Data: data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(200, Result{
		Code: 401,
		Msg:  msg,
		Data: nil,
	})
}
