package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/order-web/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)
		zap.S().Infof("current user id %d ,role %d",currentUser.ID,currentUser.AuthorityId)  //currentUser.ID  注意字段名大小写  不是Id
		if currentUser.AuthorityId != 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
