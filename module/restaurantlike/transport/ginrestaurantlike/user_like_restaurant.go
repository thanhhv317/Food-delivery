package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	bizrestaurantlike "golang/module/restaurantlike/biz"
	restaurantlikemodel "golang/module/restaurantlike/model"
	restaurantlikestorage "golang/module/restaurantlike/storage"
	"net/http"
)

// POST /v1/restaurants/:id/like (RPC)
// OR POST /v1/restaurants/:id/liked-users

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMaiDBConnection())
		//countStore := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := bizrestaurantlike.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
