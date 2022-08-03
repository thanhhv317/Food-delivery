package restaurantstorage

import (
	"golang.org/x/net/context"
	restaurantmodel "golang/module/restaurant/model"
)

func (s *sqlStore) DeleteRestaurant(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
