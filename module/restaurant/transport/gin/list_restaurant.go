package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang/common"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	restaurantstorage "golang/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

type fakeListStore struct {
}

func (fakeListStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	return []restaurantmodel.Restaurant{}, nil
}

func ListRestaurants(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result []restaurantmodel.Restaurant

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := bizrestaurant.NewListRestaurantBiz(store)
		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging, "filter": filter})
	}
}
