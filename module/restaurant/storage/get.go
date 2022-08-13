package restaurantstorage

import (
	"golang.org/x/net/context"
	"golang/common"
	restaurantmodel "golang/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var data restaurantmodel.Restaurant
	if err := db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
