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

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Process()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := bizrestaurantlike.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}