package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	bizrestaurant "golang/module/restaurant/business"
	restaurantstorage "golang/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := bizrestaurant.NewGetRestaurantBiz(store)
		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
