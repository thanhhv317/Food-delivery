package restaurantmodel

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantUpdate) Validate() error {
	if v := data.Name; v != nil && *v == "" {
		return ErrRestaurantNameCannotBeBlank
	}
	return nil
}
