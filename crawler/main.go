package main

import (
	"com.buff/Crawler/crawler/crawler/engine"
	"com.buff/Crawler/crawler/crawler/zhenai/parse"
	"com.buff/Crawler/crawler/crawler/scheduler"
)

/**
程序入口
 */
func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})
}
