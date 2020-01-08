package entity

import "time"

// FileInfo 响应的文件信息
type FileInfo struct {
	Name    string    `json:"name"`
	IsDir   bool      `json isDir`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
}
