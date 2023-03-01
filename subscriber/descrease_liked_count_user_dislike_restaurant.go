package subscriber

import (
	"context"
	"golang/component/appctx"
	restaurantstorage "golang/module/restaurant/storage"
	restaurantlikemodel "golang/module/restaurantlike/model"
	"golang/pubsub"
)

func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user dislikes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			likeData := msg.Data().(*restaurantlikemodel.Like)
			return store.DecreaseLikeCount(ctx, likeData.RestaurantId)
		},
	}
}
