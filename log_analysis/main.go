package main

import (
	"log_analysis/lfile"
	"log_analysis/llog"
	"reflect"
	"time"
)

func main() {

	r := lfile.New("D:/temp/test.txt")

	analysis := llog.New()

	analysis.AddAnalysisType("default", reflect.TypeOf(llog.BaseLog{}))

	analysis.Run()

	for {
		val := r.GetLine()
		analysis.PrepareAnalysis("default", val)
		time.Sleep(1 * time.Second)
	}

}
