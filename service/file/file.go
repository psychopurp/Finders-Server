package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize：获取文件大小 MB
func GetSize(f multipart.File) (float64, error) {
	content, err := ioutil.ReadAll(f)

	return float64(len(content)) / float64(1024*1024), err
}

// GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist：检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return !os.IsNotExist(err)
}

// CheckPermission：检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir：如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir：新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open：打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MkDirByYearMonthDay(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	return nil
}
