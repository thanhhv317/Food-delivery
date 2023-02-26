package restaurantmodel

import "golang/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Address string         `json:"address" gorm:"column:addr;"`
	OwnerId int            `json:"owner_id" gorm:"column:owner_id;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isOwnerOrAdmin bool) {
	data.SQLModel.Mask(common.DbTypeRestaurant)
}
