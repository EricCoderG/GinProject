package middleware

import (
	"GinProject/common"
	"GinProject/model"
	"GinProject/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 验证通过后获取claim中的userId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
