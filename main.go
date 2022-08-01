package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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

	log.Println(db)

	// Create
	//n := &Note{
	//	Name: "Note 4",
	//}
	//
	//if err := db.Create(n).Error; err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(n)

	// Get first

	var myNote Note

	if err := db.
		Where("id = ?", 2).
		First(&myNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(myNote)

	// List
	var listNote []Note

	if err := db.Find(&listNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(listNote)

	// update

	myNote.Name = "note updated"

	if err := db.Where("id = ?", 3).Updates(myNote).Error; err != nil {
		log.Fatal(err)
	}

	log.Println(myNote)

	// update name to ""

	emptyString := ""

	if err := db.Where("id = ?", 2).Updates(NoteUpdate{Name: &emptyString}).Error; err != nil {
		log.Fatal(err)
	}

	// Delete

	//if err := db.Table(Note{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Fatal(err)
	//}

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		restaurant := v1.Group("/restaurants")
		{
			// CRUD
			restaurant.POST("", func(c *gin.Context) {
				var newRestaurant RestaurantCreate
				if err := c.ShouldBind(&newRestaurant); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Create(&newRestaurant).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusCreated, gin.H{"restaurant": newRestaurant})
			})

			restaurant.GET("", func(c *gin.Context) {
				var result []Restaurant

				var paging struct {
					Limit int   `json:"limit" form:"limit"`
					Page  int   `json:"page" form:"page"`
					Total int64 `json:"total" form:"total"`
				}

				if err := c.ShouldBind(&paging); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Limit <= 0 {
					paging.Limit = 10
				}

				if paging.Page <= 0 {
					paging.Page = 1
				}

				if err := db.Table(Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Order("id desc").Find(&result).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})

			})

			restaurant.GET("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				var data Restaurant
				if err := db.Where("id =?", id).First(&data).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": data})

			})

			restaurant.PUT("/:id", func(c *gin.Context) {
				var updateRestaurant RestaurantUpdate
				if err := c.ShouldBind(&updateRestaurant); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				id, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id =?", id).Updates(&updateRestaurant).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": updateRestaurant})

			})

			restaurant.DELETE("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				return
			})
		}
	}

	r.Run(":8080")
}
