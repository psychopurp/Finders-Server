package upload

import (
	"finders-server/global"
	"finders-server/service/file"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// GetVideoFullUrl：获取视屏完整访问URL
func GetVideoFullUrl(name string) string {
	return global.CONFIG.AppSetting.PrefixUrl + "/" + GetVideoPath() + name
}

// GetVideoName：获取视屏名称
func GetVideoName(name string) string {
	ext := path.Ext(name)
	timeStamp := time.Now().Unix()
	timeFileName := strconv.FormatInt(timeStamp, 10)

	return timeFileName + ext
}

// GetVideoPath：获取视屏路径
func GetVideoPath() string {
	// imagesavepath: upload/images/
	return global.CONFIG.AppSetting.VideoSavePath + GetTimePath()
}

// GetVideoFullPath：获取视屏完整路径
func GetVideoFullPath() string {
	// runtimerootpath: runtime/
	// imagesavepath: upload/images/
	prefix := global.CONFIG.AppSetting.RuntimeRootPath + global.CONFIG.AppSetting.VideoSavePath
	return prefix
}

func GetVideoFullPathAndMKDir() string {
	// runtimerootpath: runtime/
	// imagesavepath: upload/images/
	prefix := global.CONFIG.AppSetting.RuntimeRootPath + GetVideoPath()
	_ = file.MkDirByYearMonthDay(prefix)
	return prefix
}

// CheckVideoExt：检查视屏后缀
func CheckVideoExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range global.CONFIG.AppSetting.VideoAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckVideoSize：检查视屏大小
func CheckVideoSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		global.LOG.Warning(err)
		return false
	}
	return size <= float64(global.CONFIG.AppSetting.VideoMaxSize)
}

// CheckVideo：检查视屏
func CheckVideo(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
