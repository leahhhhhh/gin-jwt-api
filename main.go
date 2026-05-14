// package main

// import (
// 	"errors"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // ======================
// // 1. 统一返回封装
// // ======================
// type Result struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }

// func Success(c *gin.Context, data interface{}) {
// 	c.JSON(http.StatusOK, Result{
// 		Code: 200,
// 		Msg:  "成功",
// 		Data: data,
// 	})
// }

// func Fail(c *gin.Context, msg string) {
// 	c.JSON(http.StatusOK, Result{
// 		Code: 500,
// 		Msg:  msg,
// 		Data: nil,
// 	})
// }

// // ======================
// // 2. 数据库模型（自动建表）
// // ======================
// type User struct {
// 	gorm.Model            // 自带 ID、创建时间、更新时间、删除时间
// 	Username   string     `json:"username" gorm:"size:32;unique"`
// 	Password   string     `json:"password" gorm:"size:64"`
// 	Phone      string     `json:"phone" gorm:"size:16"`
// 	CreatedAt  time.Time  `json:"created_at"`
// 	UpdatedAt  time.Time  `json:"updated_at"`
// 	DeletedAt  *time.Time `json:"deleted_at" gorm:"index"`
// }

// // 全局DB对象
// var db *gorm.DB

// // JWT密钥 自己随便改
// const jwtSecret = "my-secret-key-123"

// // ======================
// // 3. 初始化连接MySQL
// // ======================
// func initMySQL() {
// 	// 这里改成你自己的 MySQL 账号、密码、数据库名
// 	dsn := "root:xqcz1234@tcp(127.0.0.1:3306)/gin_db?charset=utf8mb4&parseTime=True&loc=Local"

// 	var err error
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("数据库连接失败：" + err.Error())
// 	}

// 	// 自动建表
// 	err = db.AutoMigrate(&User{})
// 	if err != nil {
// 		panic("建表失败：" + err.Error())
// 	}
// }

// // 生成JWT
// func GenToken(userId uint) (string, error) {
// 	claims := jwt.MapClaims{
// 		"userId": userId,
// 		"exp":    time.Now().Add(24 * time.Hour).Unix(), // 有效期24小时
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(jwtSecret))
// }

// // JWT中间件 鉴权
// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		auth := c.GetHeader("Authorization")
// 		if auth == "" {
// 			Fail(c, "未登录，请先登录")
// 			c.Abort()
// 			return
// 		}
// 		// 格式 Bearer xxx
// 		if len(auth) < 7 || auth[:7] != "Bearer " {
// 			Fail(c, "token格式错误")
// 			c.Abort()
// 			return
// 		}
// 		tokenStr := auth[7:]
// 		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtSecret), nil
// 		})
// 		if err != nil {
// 			Fail(c, "token无效或已过期")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func main() {
// 	// 初始化数据库
// 	initMySQL()

// 	r := gin.Default()

// 	// 不需要登录的接口
// 	r.POST("/register", func(c *gin.Context) {
// 		var u User
// 		if err := c.ShouldBindJSON(&u); err != nil {
// 			Fail(c, "参数错误")
// 			return
// 		}
// 		db.Create(&u)
// 		Success(c, "注册成功")
// 	})

// 	// 登录接口 颁发token
// 	r.POST("/login", func(c *gin.Context) {
// 		var req User
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			Fail(c, "参数错误")
// 			return
// 		}
// 		var user User
// 		err := db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			Fail(c, "账号密码错误")
// 			return
// 		}
// 		// 生成token
// 		token, _ := GenToken(user.ID)
// 		Success(c, gin.H{"token": token})
// 	})

// 	// 需要登录保护的路由组
// 	authGroup := r.Group("/auth")
// 	authGroup.Use(JWTMiddleware())
// 	{
// 		authGroup.GET("/info", func(c *gin.Context) {
// 			Success(c, "这是需要登录才能访问的用户信息")
// 		})
// 	}
// 	// ======================
// 	// 4. 接口：增删改查
// 	// ======================
// 	// userGroup := r.Group("/user")
// 	// {
// 	// 	// 1. 添加用户
// 	// 	userGroup.POST("/add", func(c *gin.Context) {
// 	// 		var user User
// 	// 		if err := c.ShouldBindJSON(&user); err != nil {
// 	// 			Fail(c, "参数错误")
// 	// 			return
// 	// 		}
// 	// 		// 插入数据库
// 	// 		err := db.Create(&user).Error
// 	// 		if err != nil {
// 	// 			Fail(c, "创建失败："+err.Error())
// 	// 			return
// 	// 		}
// 	// 		Success(c, user)
// 	// 	})

// 	// 	// 2. 根据ID查询用户
// 	// 	userGroup.GET("/info/:id", func(c *gin.Context) {
// 	// 		id := c.Param("id")
// 	// 		var user User
// 	// 		err := db.First(&user, id).Error
// 	// 		if err != nil {
// 	// 			Fail(c, "用户不存在")
// 	// 			return
// 	// 		}
// 	// 		Success(c, user)
// 	// 	})

// 	// 	// 3. 更新用户
// 	// 	userGroup.POST("/update", func(c *gin.Context) {
// 	// 		var user User
// 	// 		if err := c.ShouldBindJSON(&user); err != nil {
// 	// 			Fail(c, "参数错误")
// 	// 			return
// 	// 		}
// 	// 		err := db.Save(&user).Error
// 	// 		if err != nil {
// 	// 			Fail(c, "更新失败")
// 	// 			return
// 	// 		}
// 	// 		Success(c, "更新成功")
// 	// 	})

// 	// 	// 4. 删除用户
// 	// 	userGroup.POST("/delete", func(c *gin.Context) {
// 	// 		var user User
// 	// 		if err := c.ShouldBindJSON(&user); err != nil {
// 	// 			Fail(c, "参数错误")
// 	// 			return
// 	// 		}
// 	// 		err := db.Delete(&user).Error
// 	// 		if err != nil {
// 	// 			Fail(c, "删除失败")
// 	// 			return
// 	// 		}
// 	// 		Success(c, "删除成功")
// 	// 	})
// 	// }

// 	r.Run(":8080")
// }

package main

import (
	"gin-demo/config"
	"gin-demo/model"
	"gin-demo/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() {
	db, err := gorm.Open(mysql.Open(config.MySQLDSN), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	// 自动迁移表结构
	db.AutoMigrate(&model.User{})
	config.DB = db
}

func main() {
	// 初始化数据库
	initMySQL()
	// 初始化路由
	r := router.InitRouter()
	// 启动服务
	r.Run(":8080")
}
