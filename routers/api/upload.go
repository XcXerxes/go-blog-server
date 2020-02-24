/*
 * @Description: 上传
 * @Author: leo
 * @Date: 2020-02-23 20:29:16
 * @LastEditors: leo
 * @LastEditTime: 2020-02-23 20:58:52
 */

package api

import (
	"github.com/XcXerxes/go-blog-server/pkg/app"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	appG := app.Gin{c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}
	err = upload.CheckImage(fullPath)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	err = c.SaveUploadedFile(image, src)
	if err != nil  {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url": upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
