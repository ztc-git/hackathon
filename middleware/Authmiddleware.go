package middleware

import (
	"github.com/gin-gonic/gin"
	"hackathon/initDB"
	"hackathon/jwt"
	"hackathon/model"
	"net/http"
	"strings"
)



func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized,  gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,  gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证后获取claims 中userID
		userId := claims.UserId
		var user model.User
		initDB.Db.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,  gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
