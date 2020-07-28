package file

import (
	"io/ioutil"
	"mime/multipart"
	"path"
)

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

//获取文件后缀
func GetExt(filename string) string {
	return path.Ext(filename)
}
