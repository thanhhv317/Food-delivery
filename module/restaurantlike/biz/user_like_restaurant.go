package bizrestaurantlike

import (
	"context"
	"golang/component/asyncjob"
	restaurantlikemodel "golang/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseRestaurantCounterStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store      UserLikeRestaurantStore
	countStore IncreaseRestaurantCounterStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	countStore IncreaseRestaurantCounterStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:      store,
		countStore: countStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// run with job;
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.countStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	//job.SetRetryDurations([]time.Duration{time.Second * 3})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	// run with goroutines
	//go func() {
	//	defer common.Recover()
	//	_ = biz.countStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//}()

	return nil
}
