package ginuser

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	"net/http"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
