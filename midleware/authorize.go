package midleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"golang/common"
	"golang/component/appctx"
	"golang/component/tokenprovider/jwt"
	usermodel "golang/module/user/model"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuth(appCtx appctx.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		c1, span := trace.StartSpan(c.Request.Context(), "RequireAuth.middleware")

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			span.End()
			panic(err)
		}

		//db := appCtx.GetMaiDBConnection()
		//store := userstore.NewSQLStore(db)
		//
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			span.End()
			panic(err)
		}
		//
		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		user, err := authStore.FindUser(c1, map[string]interface{}{"id": payload.UserId})

		if err != nil {
			span.End()
			panic(err)
		}

		if user.Status == 0 {
			span.End()
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}
		span.End()
		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
