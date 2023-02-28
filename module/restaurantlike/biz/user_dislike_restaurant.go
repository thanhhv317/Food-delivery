package bizrestaurantlike

import (
	"context"
	"golang/common"
	"golang/component/asyncjob"
	restaurantlikemodel "golang/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseCounterStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store      UserDislikeRestaurantStore
	countStore DecreaseCounterStore
}

func NewUserDislikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	countStore DecreaseCounterStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:      store,
		countStore: countStore,
	}
}

func (biz *userDislikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Delete(ctx, data.UserId, data.RestaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	//go func() {
	//	defer common.Recover()
	//	_ = biz.countStore.DecreaseLikeCount(ctx, data.RestaurantId)
	//}()

	// Job & Job Manager

	go func() {
		defer common.Recover()

		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.countStore.DecreaseLikeCount(ctx, data.RestaurantId)
		})

		//job.SetRetryDurations([]time.Duration{time.Second * 3})

		_ = asyncjob.NewGroup(true, job).Run(ctx)

	}()

	return nil
}
