package restaurantmodel

import (
	"golang/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name       string             `json:"name" gorm:"column:name;"`
	Address    string             `json:"address" gorm:"column:addr;"`
	OwnerId    int                `json:"owner_id" gorm:"column:owner_id;"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	User       *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerId;preload:false"`
	LikedCount int                `json:"liked_count" gorm:"column:liked_count"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isOwnerOrAdmin bool) {
	data.GenUID(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isOwnerOrAdmin)
	}
}
