package bizrestaurant

import (
	"golang.org/x/net/context"
	"golang/common"
	restaurantmodel "golang/module/restaurant/model"
)

type GetRestaurantStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return data, nil
}
