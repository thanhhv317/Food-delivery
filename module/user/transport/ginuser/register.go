package ginuser

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	"golang/component/hasher"
	userbiz "golang/module/user/biz"
	usermodel "golang/module/user/model"
	userstorage "golang/module/user/storage"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
