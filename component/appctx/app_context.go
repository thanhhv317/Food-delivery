package appctx

import (
	"golang/component/uploadprovider"
	"golang/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pb             pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string, pb pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, uploadProvider: provider, secretKey: secretKey, pb: pb}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB { return ctx.db }

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.Pubsub { return ctx.pb }
