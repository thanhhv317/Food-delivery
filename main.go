package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/component/appctx"
	"golang/midleware"
	ginrestaurant "golang/module/restaurant/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type SQLModel struct {
	ID        int       `json:"id" gorm:"column:id;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status    int       `json:"status" gorm:"column:status;"`
}

type Note struct {
	SQLModel
	Name    string `gorm:"column:title;"`
	Address int    `gorm:"column:category_id;"`
}

func (Note) TableName() string { return "notes" }

type Restaurant struct {
	SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

type NoteUpdate struct {
	Name       *string `gorm:"column:title;"`
	CategoryId *int    `gorm:"column:category_id;"`
}

func (NoteUpdate) TableName() string { return Note{}.TableName() }

func main() {
	fmt.Println("Hello world")

	dsn := "root:123456789@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug() // show query

	if err != nil {
		log.Fatal(err)
	}

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(midleware.Recover(appCtx))
	v1 := r.Group("/v1")
	{
		restaurant := v1.Group("/restaurants")
		{
			// CRUD
			restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))

			restaurant.GET("", ginrestaurant.ListRestaurants(appCtx))

			restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

			restaurant.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))

			restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		}
	}

	r.Run(":8080")
}
