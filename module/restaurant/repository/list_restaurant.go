package reporestaurant

import (
	"context"
	"golang/common"
	restaurantmodel "golang/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type GetLikedCountStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store      ListRestaurantStore
	likedStore GetLikedCountStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likedStore GetLikedCountStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likedStore: likedStore}
}

func (biz *listRestaurantRepo) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, moreKeys...)

	if err != nil {
		return nil, err
	}

	//resIDs := make([]int, len(result))
	//for i := range resIDs {
	//	resIDs[i] = result[i].ID
	//}
	//
	//if mapResLiked, err := biz.likedStore.GetRestaurantLikes(ctx, resIDs); err == nil {
	//	for i := range result {
	//		result[i].LikedCount = mapResLiked[result[i].ID]
	//	}
	//}

	return result, nil
}
