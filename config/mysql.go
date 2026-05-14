/*
 * @Author: jw ljw1271005234@gmail.com
 * @Date: 2026-05-13 18:22:47
 * @LastEditors: jw ljw1271005234@gmail.com
 * @LastEditTime: 2026-05-13 18:23:28
 * @FilePath: /jw_go/config/mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import "gorm.io/gorm"

var DB *gorm.DB

const (
	MySQLDSN  = "root:xqcz1234@tcp(127.0.0.1:3306)/gin_db?charset=utf8mb4&parseTime=True&loc=Local"
	JwtSecret = "my-secret-key-123"
)
