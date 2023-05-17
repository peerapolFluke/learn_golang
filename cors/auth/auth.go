package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"seamoor/cors/utils"
	"seamoor/ent"
	"seamoor/services"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Next()
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := utils.JwtVerify(auth)
		if err != nil {
			return
		}

		user, err := services.GetUserById(c.Request.Context(), validate.ID)
		ctx := context.WithValue(c.Request.Context(), userCtxKey, user)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}
