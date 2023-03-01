package subscriber

import "golang/common"

//func Setup(appCtx appctx.AppContext) {
//	RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
//	RunPushNotificationAfterUserLikeRestaurant(appCtx, context.Background())
//}

func (engine *subscriberEngine) setup() {
	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		PushNotificationAfterUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DecreaseLikeCountAfterUserDislikeRestaurant(engine.appCtx),
	)
}
