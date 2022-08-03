package restaurantstorage

import (
	"golang.org/x/net/context"
	restaurantmodel "golang/module/restaurant/model"
)

func (s *sqlStore) updateData(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}
