/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:23:50
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-13 18:24:02
 * @FilePath: /jw_go/model/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:32;unique"`
	Password string `json:"password" gorm:"size:64"`
}
