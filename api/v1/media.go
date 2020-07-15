package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/service/upload"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// @Summary 上传图片
// @Description 上传图片
// @Tags 上传图片
// @Accept mpfd
// @Produce json
// @Param image formData file true "图片"
// @Success 200 {string} string "success: {"code": 0, data:{image_url:""}, "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/upload_image [post]
func UploadImage(c *gin.Context) {
	var (
		file  multipart.File
		image *multipart.FileHeader
		err   error
		user  model.User
		media model.Media
	)
	// 获取用户名
	userName := c.GetHeader("username")
	// 获取用户
	user, err = model.GetUserByUserName(userName)
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	// 从form表单中获取图片
	file, image, err = c.Request.FormFile("image")
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if image == nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	// 获取新的图片名称
	imageName := upload.GetImageName(image.Filename)
	// 获取图片路径并创建存放的文件夹
	fullPath := upload.GetImageFullPathAndMKDir()
	// 图片的路径
	src := fullPath + imageName
	// 检查图片的后缀和大小是否符合规范
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		response.FailWithMsg(e.IMAGE_FORMAT_OR_SIZE_ERROR, c)
		return
	}
	err = upload.CheckImage(fullPath)
	if err != nil {
		response.FailWithMsg(e.UPLOAD_CHECK_FILE_ERROR, c)
		return
	}
	// 存放照片
	err = c.SaveUploadedFile(image, src)
	if err != nil {
		response.FailWithMsg(e.UPLOAD_SAVE_FILE_ERROR, c)
		return
	}
	data := make(map[string]interface{})
	//data["image_url"] = upload.GetImageFullUrl(imageName)
	imageUrl := upload.GetImagePath() + imageName
	data["image_url"] = imageUrl
	// 加入照片的存放记录
	media, err = model.AddMedia(imageUrl, user.UserID.String(), model.PICTURE)
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	data["image_id"] = media.MediaID
	response.OkWithData(data, c)
}

// @Summary 上传视屏
// @Description 上传视屏
// @Tags 上传视屏
// @Accept mpfd
// @Produce json
// @Param video formData file true "视屏"
// @Success 200 {string} string "success: {"code": 0, data:{video_url:""}, "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/upload_video [post]
func UploadVideo(c *gin.Context) {
	var (
		file  multipart.File
		video *multipart.FileHeader
		err   error
		user  model.User
		media model.Media
	)
	// 获取用户名
	userName := c.GetHeader("username")
	// 获取用户
	user, err = model.GetUserByUserName(userName)
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	// 获取视屏文件
	file, video, err = c.Request.FormFile("video")
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	if video == nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	// 获取视屏新名称
	videoName := upload.GetVideoName(video.Filename)
	// 获取完整路径并创建路径文件夹
	fullPath := upload.GetVideoFullPathAndMKDir()
	src := fullPath + videoName
	// 检查视屏后缀和大小
	if !upload.CheckVideoExt(videoName) || !upload.CheckVideoSize(file) {
		response.FailWithMsg(e.IMAGE_FORMAT_OR_SIZE_ERROR, c)
		return
	}
	// 检查视屏
	err = upload.CheckVideo(fullPath)
	if err != nil {
		response.FailWithMsg(e.UPLOAD_CHECK_FILE_ERROR, c)
		return
	}
	err = c.SaveUploadedFile(video, src)
	if err != nil {
		response.FailWithMsg(e.UPLOAD_SAVE_FILE_ERROR, c)
		return
	}
	data := make(map[string]interface{})

	videoURL := upload.GetVideoPath() + videoName
	data["video_url"] = videoURL
	// 增加视屏存放记录
	media, err = model.AddMedia(videoURL, user.UserID.String(), model.VIDEO)
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	data["video_id"] = media.MediaID
	response.OkWithData(data, c)
}
