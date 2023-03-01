package subscriber

import (
	"golang.org/x/net/context"
	"golang/component/appctx"
	restaurantstorage "golang/module/restaurant/storage"
	restaurantlikemodel "golang/module/restaurantlike/model"
	"golang/pubsub"
	"log"
)

//
//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//
//	store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
//
//	go func() {
//		defer common.Recover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(*restaurantlikemodel.Like)
//			_ = store.IncreaseLikeCount(ctx, likeData.RestaurantId)
//		}
//	}()
//}

//func RunPushNotificationAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//
//	//store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
//
//	go func() {
//		defer common.Recover()
//		for {
//			msg := <-c
//			log.Println("PushNotificationAfterUserLikeRestaurant:", msg)
//		}
//	}()
//}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) func(*pubsub.Message, context.Context) error {
//	return func(msg *pubsub.Message, ctx context.Context) error {
//		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
//
//		likeData := msg.Data().(*restaurantlikemodel.Like)
//		return store.IncreaseLikeCount(ctx, likeData.RestaurantId)
//	}
//}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			likeData := msg.Data().(*restaurantlikemodel.Like)
			return store.IncreaseLikeCount(ctx, likeData.RestaurantId)
		},
	}
}

func PushNotificationAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification when user likes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			//store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			//likeData := msg.Data().(*restaurantlikemodel.Like)
			log.Println("PushNotificationAfterUserLikeRestaurant:", msg)
			return nil
		},
	}
}
