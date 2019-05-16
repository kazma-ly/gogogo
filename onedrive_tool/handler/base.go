package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"onedrivetool/defs"
	"onedrivetool/util"
	"os"

	"github.com/julienschmidt/httprouter"
)

// GoMSAuth 前往MS认证
func GoMSAuth(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	http.Redirect(w, r, util.RedirectURL, http.StatusMovedPermanently)
}

// MSCallBack MS的Callback
func MSCallBack(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	q := r.URL.Query()
	code := q.Get("code")

	msResult, err := util.GetAccessToken(code)
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorInternalServer)
		return
	}

	defs.SendSuccessResponse(w, "200", "成功", 200, msResult)
}

// GetOneDriveFile 下载文件
func GetOneDriveFile(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	msResult, err := util.GetAccessToken("")
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorCanNotGetAccessToken)
		return
	}

	bs, err := util.DownloadFile("https://graph.microsoft.com/v1.0/me/drive/items/67A5E9B5809A1843!11017/content", msResult.AccessToken)
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorDownload)
		return
	}
	ioutil.WriteFile(os.TempDir()+"/11223.json", bs, 0666)

	defs.SendSuccessResponse(w, "200", "成功", 200, string(bs))
}

// UploadDriveFile 上传新文件
func UploadDriveFile(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	msResult, err := util.GetAccessToken("")
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorCanNotGetAccessToken)
		return
	}
	f, err := os.OpenFile(os.TempDir()+"/11223.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorInternalServer)
		return
	}

	bs, err := util.UpdateFile("67A5E9B5809A1843!11017", msResult.AccessToken, f)
	if err != nil {
		log.Println(err)
		defs.SendErrorResponse(w, defs.ErrorInternalServer)
		return
	}

	defs.SendSuccessResponse(w, "200", "成功", 200, string(bs))
}
