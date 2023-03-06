package skio

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"golang/common"
	"golang/component/tokenprovider/jwt"
	userstorage "golang/module/user/storage"
	"golang/module/user/transport/skiouser"
	"gorm.io/gorm"
	"sync"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	SecretKey() string
	GetUser(conn socketio.Conn) common.Requester
}

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	GetUser(conn socketio.Conn) common.Requester
	//Run(ctx AppContext, engine *gin.Engine) error
	//Emit(userId int) error
}

// A -x-> [0x111]{b: 0x222} (Retain count: 1)
//					||
//					||
//					||
// B -x-> [0x222]{a: 0x111} (Retain count: 1)

type rtEngine struct {
	server      *socketio.Server
	storage     map[int][]AppSocket  // userId : []AppSocket
	storageConn map[string]AppSocket // clientId : AppSocket
	locker      *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage:     make(map[int][]AppSocket),
		storageConn: make(map[string]AppSocket),
		locker:      new(sync.RWMutex),
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	//appSck.Join("order-{ordID}")

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.storageConn[appSck.ID()] = appSck

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}

	delete(engine.storageConn, appSck.ID())
}

func (engine *rtEngine) GetUser(conn socketio.Conn) common.Requester {
	ask, _ := engine.GetAppSocket(conn)
	return ask
}

func (engine *rtEngine) GetAppSocket(conn socketio.Conn) (AppSocket, error) {
	if appSocket, ok := engine.storageConn[conn.ID()]; ok {
		return appSocket, nil
	}

	return nil, errors.New("socket not found")
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	engine.locker.RLock()
	defer engine.locker.RUnlock()
	if scks, ok := engine.storage[userId]; ok {
		return scks
	}

	return sockets
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}

func (engine *rtEngine) Run(appCtx AppContext, r *gin.Engine) error {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr(), s.ID())

		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		if appSck, err := engine.GetAppSocket(s); err == nil {
			engine.removeAppSocket(appSck.GetUserId(), appSck)
		}

		fmt.Println("closed", reason)
	})

	// Setup

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMaiDBConnection()
		store := userstorage.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}

		user.Mask(false)

		// Important: New AppSocket
		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.ID, appSck)

		s.Emit("authenticated", user)

		//appSck.Join(user.GetRole()) // the same
		//if user.GetRole() == "admin" {
		//	appSck.Join("admin")
		//}

		//server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))
	})

	server.OnEvent("/", string(common.TopicUserUpdateLocation), skiouser.RealtimeUserUploadLocation(appCtx))

	go server.Serve()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
