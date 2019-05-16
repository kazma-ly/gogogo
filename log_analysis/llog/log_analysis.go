package llog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//type JsonTime time.Time
//
////实现它的json序列化方法
//func (jt JsonTime) MarshalJSON() ([]byte, error) {
//	var stamp = fmt.Sprintf("\"%s\"", time.Time(jt).Format("2006-01-02 15:04:05"))
//	return []byte(stamp), nil
//}

//func (jt JsonTime) UnmarshalJSON(data []byte) error {
//	t, err := time.Parse("\"2006-01-02 15:04:05\"", string(data))
//	return err
//}

type (
	// LogAnalysis 日志分析
	LogAnalysis struct {
		analysisQueue chan AnalysisItem
		regMap        map[string]reflect.Type
	}

	// AnalysisItem 分析对象
	AnalysisItem struct {
		data string
		t    reflect.Type
	}

	// BaseLog 日志model
	BaseLog struct {
		Level string      `json:"level"`
		Msg   interface{} `json:"msg"`
		Time  time.Time   `json:"time"`
	}
)

func New() *LogAnalysis {
	return &LogAnalysis{
		regMap:        make(map[string]reflect.Type),
		analysisQueue: make(chan AnalysisItem, 10),
	}
}

// AddAnalysisType 添加分析类型
func (la *LogAnalysis) AddAnalysisType(regtype string, model reflect.Type) {
	la.regMap[regtype] = model
}

// Run 启动 开始处理PrepareAnalysis
func (la *LogAnalysis) Run() {
	go processLogChan(la)
}

func processLogChan(la *LogAnalysis) {
	for {
		select {
		case analysisItem := <-la.analysisQueue:
			val := analysisItem.data
			if val != "" {
				v := reflect.New(analysisItem.t)
				res := v.Interface()
				err := json.Unmarshal([]byte(val), res)
				if err != nil {
					fmt.Println(err)
				} else {
					ProcessLog(res)
				}
			}
		}
	}
}

// PrepareAnalysis 把日志放进队列
func (la *LogAnalysis) PrepareAnalysis(regtype, data string) {
	item := AnalysisItem{
		data: data,
		t:    la.regMap[regtype],
	}
	la.analysisQueue <- item
}
