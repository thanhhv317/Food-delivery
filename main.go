package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/component/appctx"
	"golang/component/uploadprovider"
	"golang/memcache"
	"golang/midleware"
	ginrestaurant "golang/module/restaurant/transport/gin"
	"golang/module/restaurantlike/transport/ginrestaurantlike"
	ginupload "golang/module/upload/transport/gin"
	userstorage "golang/module/user/storage"
	"golang/module/user/transport/ginuser"
	"golang/pubsub/localpb"
	"golang/skio"
	"golang/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	r := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	s3Provider := uploadprovider.NewS3Provider("", "", "", "", "")
	pb := localpb.NewPubSub()

	rtEngine := skio.NewEngine()

	secretKey := os.Getenv("SYSTEM_SECRET")
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey, pb, rtEngine)
	rtEngine.Run(appCtx, r)

	_ = subscriber.NewEngine(appCtx).Start()

	r.Use(midleware.Recover(appCtx))

	r.Static("/static", "./static")
	r.StaticFile("/demo/", "./demo.html")

	userStorage := userstorage.NewSQLStore(db)
	userCaching := memcache.NewUserCaching(memcache.NewCaching(), userStorage)
	midAuthorize := midleware.RequiredAuth(appCtx, userCaching)

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

	//r.Run(":8080")

	//startSocketIOServer(r, appCtx)

	if err := http.ListenAndServe(
		":8080",
		r,
		//r,
	); err != nil {
		log.Fatalln(err)
	}
}

//func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
//	server := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("Socket connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		s.Join("Shipper")
//		//server.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
//
//		return nil
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		log.Println("test:", msg)
//	})
//
//	//go func() {
//	//	for range time.NewTicker(time.Second).C {
//	//		server.BroadcastToRoom("/", "Shipper", "test", "Ahihi")
//	//	}
//	//}()
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//		// Remove socket from socket engine (from app context)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//
//		// Validate token
//		// If false: s.Close(), and return
//
//		// If true
//		// => UserId
//		// Fetch db find user by Id
//		// Here: s belongs to who? (user_id)
//		// We need a map[user_id][]socketio.Conn
//
//		db := appCtx.GetMaiDBConnection()
//		store := userstorage.NewSQLStore(db)
//		//
//		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//		//
//		payload, err := tokenProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//		//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		log.Println("test:", msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//	})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
