package restaurantmodel

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"addres s" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
