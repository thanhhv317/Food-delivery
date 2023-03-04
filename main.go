package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/component/appctx"
	"golang/component/uploadprovider"
	"golang/midleware"
	ginrestaurant "golang/module/restaurant/transport/gin"
	"golang/module/restaurantlike/transport/ginrestaurantlike"
	ginupload "golang/module/upload/transport/gin"
	userstorage "golang/module/user/storage"
	"golang/module/user/transport/ginuser"
	"golang/pubsub/localpb"
	"golang/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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

	s3Provider := uploadprovider.NewS3Provider("", "", "", "", "")
	pb := localpb.NewPubSub()

	secretKey := os.Getenv("SYSTEM_SECRET")
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey, pb)

	_ = subscriber.NewEngine(appCtx).Start()

	r := gin.Default()
	r.Use(midleware.Recover(appCtx))

	r.Static("/static", "./static")

	userStorage := userstorage.NewSQLStore(db)
	midAuthorize := midleware.RequiredAuth(appCtx, userStorage)

	v1 := r.Group("/v1")

	{
		v1.POST("/upload", ginupload.Upload(appCtx))

		v1.POST("/register", ginuser.Register(appCtx))
		v1.POST("/authenticate", ginuser.Login(appCtx))

		v1.GET("/profile", midAuthorize, ginuser.Profile(appCtx))

		restaurant := v1.Group("/restaurants", midAuthorize)
		{
			// CRUD
			restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
			restaurant.GET("", ginrestaurant.ListRestaurants(appCtx))
			restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
			restaurant.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
			restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

			restaurant.GET("/:id/like-users", ginrestaurantlike.ListUsers(appCtx))
			restaurant.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
			restaurant.DELETE("/:id/dislike", ginrestaurantlike.UserDislikeRestaurant(appCtx))

		}
	}

	// Config exporter Jaeger
	//je, err := jg.NewExporter(jg.Options{
	//	AgentEndpoint: os.Getenv("JAEGER_AGENT_URL"),
	//	Process:       jg.Process{ServiceName: "G06-Food-Delivery"},
	//})
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//trace.RegisterExporter(je)
	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})

	r.Run(":8080")
	//if err := http.ListenAndServe(
	//	":8080",
	//	&ochttp.Handler{ // moi khi request di vao la phai dang ky tracing
	//		Handler: r,
	//	},
	//	//r,
	//); err != nil {
	//	log.Fatalln(err)
	//}

}
