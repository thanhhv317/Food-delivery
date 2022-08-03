package restaurantstorage

import (
	"golang.org/x/net/context"
	restaurantmodel "golang/module/restaurant/model"
)

func (s *sqlStore) GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var data restaurantmodel.Restaurant
	if err := db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
