package main

import (
	"com.buff/learngo/crawler/engine"
	"com.buff/learngo/crawler/zhenai/parse"
)

/**
程序入口
 */
func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})
}
