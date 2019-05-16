package main

import (
	"lession/crawler/engine"
	"lession/crawler/scheduler"
	"lession/crawler/zhenai/parser"
	"lession/crawler_distributed/config"
	"lession/crawler_distributed/persist/client"
)

// 前端直接用elaticsearch 搜索就好了
func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	URL:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// default
	// saver, err := persist.ItemSaver("xq_1")
	// if err != nil {
	// 	panic(err)
	// }

	// dis, first start rpc-serve
	saver, err := client.ItemSaver(config.ItemSaverRPCAddr)
	if err != nil {
		panic(err)
	}

	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    saver,
	}

	//e.Run(engine.Request{
	//	URL:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseProfileList,
	})

	// http.HandleFunc("/index", frontend.RenderPage)
	// log.Println(http.ListenAndServe(":1332", nil))
}
