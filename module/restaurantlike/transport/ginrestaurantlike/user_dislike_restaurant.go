package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	restaurantstorage "golang/module/restaurant/storage"
	bizrestaurantlike "golang/module/restaurantlike/biz"
	restaurantlikemodel "golang/module/restaurantlike/model"
	restaurantlikestorage "golang/module/restaurantlike/storage"
	"net/http"
)

// DELETE /v1/restaurants/:id/dislike (RPC)
// OR DELETE /v1/restaurants/:id/liked-users

func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
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
		countStore := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := bizrestaurantlike.NewUserDislikeRestaurantBiz(store, countStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
