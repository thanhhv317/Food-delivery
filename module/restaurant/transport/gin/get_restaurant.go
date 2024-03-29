package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	bizrestaurant "golang/module/restaurant/business"
	restaurantstorage "golang/module/restaurant/storage"
	"net/http"
)

func GetRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		// Recover in go routine
		//go func() {
		//	defer common.Recover()
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		biz := bizrestaurant.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), int(id.GetLocalID()))

		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
