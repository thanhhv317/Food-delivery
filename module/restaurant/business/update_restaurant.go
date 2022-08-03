package bizrestaurant

import (
	"errors"
	"golang.org/x/net/context"
	restaurantmodel "golang/module/restaurant/model"
)

type UpdateStore interface {
	updateData(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type updateRestaurantBiz struct {
	store UpdateStore
}

func NewUpdateRestaurantBiz(store UpdateStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData == nil || oldData.Status == 0 {
		return errors.New("Restaurant has been deleted.")
	}

	if err := biz.store.updateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
