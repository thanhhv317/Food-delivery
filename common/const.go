package common

import "golang/pubsub"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

const (
	TopicUserLikeRestaurant    pubsub.Topic = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant pubsub.Topic = "TopicUserDislikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
