package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	bizrestaurant "golang/module/restaurant/business"
	restaurantstorage "golang/module/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		biz := bizrestaurant.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
		return
	}
}
