package main

import (
	"com.buff/learngo/crawler/engine"
	"com.buff/learngo/crawler/zhenai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})
}
