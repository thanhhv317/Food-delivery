package bizrestaurant

import (
	"golang.org/x/net/context"
	restaurantmodel "golang/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	DeleteRestaurant(ctx context.Context, id int) error
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	_, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	error := biz.store.DeleteRestaurant(ctx, id)
	if err != nil {
		return error
	}
	return nil
}
