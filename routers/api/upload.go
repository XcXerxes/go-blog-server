/*
 * @Description: 上传
 * @Author: leo
 * @Date: 2020-02-23 20:29:16
 * @LastEditors: leo
 * @LastEditTime: 2020-02-23 20:58:52
 */

package api

import (
	"github.com/XcXerxes/go-blog-server/models"
	"github.com/XcXerxes/go-blog-server/pkg/e"
	"github.com/XcXerxes/go-blog-server/pkg/upload"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	data := make(map[string]string)
	code := e.SUCCESS
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		code = e.ERROR
		models.SendResponse(c, code, nil, data)
	}
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
		models.SendResponse(c, code, nil, data)
	}
}
