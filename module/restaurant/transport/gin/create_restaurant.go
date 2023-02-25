package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	restaurantstorage "golang/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&newRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		biz := bizrestaurant.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(newRestaurant.ID))
	}
}
