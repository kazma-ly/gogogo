package tool

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"

	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// MakeRandomNum 生成随机数
func MakeRandomNum() string {
	seedNano := time.Now().UnixNano()
	seed := rand.NewSource(seedNano) // 用指定值创建一个随机数种子
	ran := seed.Int63()
	return Md5(Md5(strconv.FormatInt(ran, 10)) + Md5(strconv.FormatInt(ran, 10)))
}

// Md5 生成MD5
func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

// SaveMultipartFile 保存SaveMultipartFile
func SaveMultipartFile(multipartFile multipart.File, savePath string) error {
	localFile, err := CheckAndCreate(savePath)
	if err != nil {
		return err
	}
	defer localFile.Close()
	io.Copy(localFile, multipartFile)
	return nil
}

// ClipFileToRect 图片裁剪成正方形
func ClipFileToRect(src *multipart.File, buf *bytes.Buffer) error {
	// file, err := CheckAndCreate(saveFile)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	srcPic, _, err := image.Decode(*src)
	if err != nil {
		return err
	}
	srcSize := srcPic.Bounds().Size()
	srcBouns := srcPic.Bounds()
	var xoffset int            // 裁剪位移
	var yoffset int            // 裁剪位移
	if srcSize.X > srcSize.Y { // 偏移
		xoffset = (srcSize.X - srcSize.Y) / 2
	} else {
		yoffset = (srcSize.Y - srcSize.X) / 2
	}
	dst := image.NewRGBA(srcBouns)
	draw.Draw(dst, srcBouns, srcPic, srcBouns.Min, draw.Src) // 转换成RGBA, image.Point{X: xoffset, Y: yoffset}
	// srcRGBA := dst.SubImage(srcBouns)
	// rgbImg := srcRGBA.(*image.RGBA)
	subImg := dst.SubImage(image.Rect(xoffset, yoffset, srcSize.X-xoffset, srcSize.Y-yoffset)) // 图片裁剪x0 y0 x1 y1
	err = png.Encode(buf, subImg)
	return err
	//white := color.RGBA{255, 255, 255, 255}
}

// PathExist 文件是否存在
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckPath 检查路路径 没有就创建
func CheckPath(p string) error {
	savePath := filepath.ToSlash(p)
	return os.MkdirAll(path.Dir(savePath), os.ModePerm)
}

// CheckAndCreate 检查并且创建文件
func CheckAndCreate(p string) (*os.File, error) {
	err := CheckPath(p)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(p, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
}
