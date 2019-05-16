package confs

import (
	"encoding/json"
	"io/ioutil"
	"life-service/logx"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var (
	fileInfo map[string]interface{}
	waitG    sync.WaitGroup
	filepath = "confs/conf.json"
)

// ReadConfValue 读取配置文件的信息
func ReadConfValue(key string) interface{} {
	if fileInfo == nil {
		waitG.Add(1)
		go readFileInfo()
		waitG.Wait()
	}
	return fileInfo[key]
}

func readFileInfo() map[string]interface{} {
	defer waitG.Done()
	// 读取文件
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fileInfo = make(map[string]interface{}, 0)
	err = json.Unmarshal(bs, &fileInfo)
	if err != nil {
		panic(err)
	}
	for k, v := range fileInfo {
		log.Println(k, "=", v)
	}
	return fileInfo
}

// NoticeFileChange 监听配置文件的变化 重新加载
func NoticeFileChange() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	// defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					waitG.Add(1)
					go readFileInfo()
				}
			case err := <-watcher.Errors:
				logx.LogInfo("error:", err)
			}
		}
	}()

	err = watcher.Add(filepath)
	if err != nil {
		panic(err)
	}
}
