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

// GetImageFullUrl：获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return global.CONFIG.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName：获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	timeStamp := time.Now().UnixNano()
	timeFileName := strconv.FormatInt(timeStamp, 10)

	return timeFileName + ext
}

// GetImagePath：获取图片路径
func GetImagePath() string {
	// imagesavepath: upload/images/
	return global.CONFIG.AppSetting.ImageSavePath + GetTimePath()
}

func GetTimePath() string {
	return time.Now().Format("2006/01/02/")
}

// GetImageFullPath：获取图片完整路径
func GetImageFullPath() string {
	// runtimerootpath: runtime/
	// imagesavepath: upload/images/
	prefix := global.CONFIG.AppSetting.RuntimeRootPath + global.CONFIG.AppSetting.ImageSavePath
	return prefix
}

func GetImageFullPathAndMKDir() string {
	// runtimerootpath: runtime/
	// imagesavepath: upload/images/
	prefix := global.CONFIG.AppSetting.RuntimeRootPath + GetImagePath()
	_ = file.MkDirByYearMonthDay(prefix)
	return prefix
}

// CheckImageExt：检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range global.CONFIG.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize：检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		global.LOG.Warning(err)
		return false
	}
	return size <= float64(global.CONFIG.AppSetting.ImageMaxSize)
}

// CheckImageSize：检查图片大小
func CheckImageSizeForMulti(size int64) bool {
	return size <= global.CONFIG.AppSetting.ImageMaxSize*1024*1024
}

// CheckImage：检查图片
func CheckImage(src string) error {
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
