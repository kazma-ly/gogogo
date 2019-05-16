package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

// DownloadFile 下载文件
func DownloadFile(url, accessToken string) ([]byte, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// UpdateFile 更新文件
func UpdateFile(itemID, accessToken string, file *os.File) ([]byte, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPut, "https://graph.microsoft.com/v1.0//me/drive/items/"+itemID+"/content", file)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
