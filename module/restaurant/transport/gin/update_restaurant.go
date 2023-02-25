package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	restaurantstorage "golang/module/restaurant/storage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var updateRestaurant restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&updateRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())

		biz := bizrestaurant.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &updateRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))

	}
}
