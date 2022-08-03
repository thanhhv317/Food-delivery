package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	restaurantstorage "golang/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := bizrestaurant.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"restaurant": newRestaurant})
	}
}
