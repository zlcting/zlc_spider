package main

import (
	"zlc_spider/crawler/engine"
	"zlc_spider/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//fmt.Printf("%s\n", all)
}
