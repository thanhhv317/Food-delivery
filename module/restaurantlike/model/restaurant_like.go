package restaurantlikemodel

import (
	"fmt"
	"golang/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"-" gorm:"column:restaurant_id;"`
	UserId       int                `json:"-" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *Like) GetUserId() int {
	return l.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDislikeRestaurant"),
	)
}

//type User struct {
//	ID        int           `json:"-" gorm:"column:id;"`
//	FakeId    common.UID    `json:"id" gorm:"-"`
//	LastName  string        `json:"last_name" gorm:"column:last_name;"`
//	FirstName string        `json:"first_name" gorm:"column:first_name;"`
//	Role      string        `json:"role" gorm:"column:role;"`
//	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
//}
//
//func (User) TableName() string {
//	return "users"
//}
//
//func (u *User) Mask(isAdmin bool) {
//	u.FakeId = common.NewUID(uint32(u.ID), common.DbTypeUser, 1)
//}
