package ginupload

import (
	"github.com/gin-gonic/gin"
	"golang/common"
	"golang/component/appctx"
	uploadbusiness "golang/module/upload/business"
)

func Upload(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// upload to static folder

		//fileHeader, err := c.FormFile("file")
		//
		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//
		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(common.ErrInternal(err))
		//}
		//
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		//	Id:        0,
		//	Url:       "http://localhost:8080/static/" + fileHeader.Filename,
		//	Width:     0,
		//	Height:    0,
		//	CloudName: "",
		//	Extension: "",
		//}))
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appContext.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
