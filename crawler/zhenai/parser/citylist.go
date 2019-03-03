package parser

import (
	"regexp"
	"zlc_spider/crawler/engine"
)

const CitylistRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CitylistRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 1 //限制爬虫城市数量 方便测试
	for _, m := range matches {
		result.Items = append(result.Items, "city"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit <= 0 { //限制爬虫城市数量 方便测试
			break
		}
	}
	return result
}
