package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	restaurantstorage "golang/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var updateRestaurant restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&updateRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := bizrestaurant.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &updateRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": updateRestaurant})

	}
}
