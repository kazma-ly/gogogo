package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"onedrivetool/confs"
	"onedrivetool/defs"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	callback = confs.MyConf.MSconf.Callback
	clientid = confs.MyConf.MSconf.Clientid
	scope    = confs.MyConf.MSconf.Scope
	sec      = confs.MyConf.MSconf.Sec
	// RedirectURL oauth2跳转的URL
	RedirectURL = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?" +
		"client_id=" + clientid +
		"&scope=" + scope +
		"&response_type=code" +
		"&redirect_uri=" + callback
	msCache = cache.New(29*time.Minute, 30*time.Minute)
	// MS_RESPONSE_CAHCE_KEY 缓存的key
	MS_RESPONSE_CAHCE_KEY = "MS_RESPONSE_CAHCE_KEY"
	// MS_REFRESH_KEY
	MS_REFRESH_KEY = ""
)

// GetAccessToken 获取accessToken
func GetAccessToken(code string) (defs.MSResponse, error) {
	var msRes defs.MSResponse

	res, has := msCache.Get(MS_RESPONSE_CAHCE_KEY)

	if has {
		return res.(defs.MSResponse), nil
	}

	client := http.DefaultClient
	urlval := url.Values{}
	urlval.Add("client_id", clientid)
	urlval.Add("redirect_uri", callback)
	urlval.Add("client_secret", sec)
	if code != "" {
		urlval.Add("code", code)
		urlval.Add("grant_type", "authorization_code")
	} else {
		urlval.Add("refresh_token", MS_REFRESH_KEY)
		urlval.Add("grant_type", "refresh_token")
	}

	resp, err := client.PostForm("https://login.microsoftonline.com/common/oauth2/v2.0/token", urlval)
	if err != nil {
		return msRes, err
	}
	defer resp.Body.Close()

	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msRes, err
	}

	err = json.Unmarshal(resByte, &msRes)

	if err != nil {
		return msRes, err
	}

	msCache.Add(MS_RESPONSE_CAHCE_KEY, msRes, time.Duration(msRes.ExpiresIn)*time.Second)
	MS_REFRESH_KEY = msRes.RefreshToken

	return msRes, nil
}
