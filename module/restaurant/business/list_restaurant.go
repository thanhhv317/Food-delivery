package bizrestaurant

import (
	"golang.org/x/net/context"
	"golang/common"
	restaurantmodel "golang/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *ListRestaurantBiz {
	return &ListRestaurantBiz{repo: repo}
}

func (biz *ListRestaurantBiz) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.repo.ListRestaurant(ctx, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	return result, nil

}
