package middleware

import (
	"ToDoList/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			//token = strings.TrimPrefix(token, "Bearer ")
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //token错误
				fmt.Println(token)
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //过期
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token验证失败",
			})
			c.Abort()
			return
		}
		fmt.Println("token验证成功")
		c.Next()
	}
}
