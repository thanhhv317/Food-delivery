package appctx

import (
	socketio "github.com/googollee/go-socket.io"
	"golang/common"
	"golang/component/uploadprovider"
	"golang/pubsub"
	"golang/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
	GetUser(conn socketio.Conn) common.Requester
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pb             pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string, pb pubsub.Pubsub, rtEngine skio.RealtimeEngine) *appCtx {
	return &appCtx{db: db, uploadProvider: provider, secretKey: secretKey, pb: pb, rtEngine: rtEngine}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB { return ctx.db }

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.Pubsub { return ctx.pb }

func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine { return ctx.rtEngine }

func (ctx *appCtx) GetUser(conn socketio.Conn) common.Requester {
	return ctx.GetRealtimeEngine().GetUser(conn)
}
