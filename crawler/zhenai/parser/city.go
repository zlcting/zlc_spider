package parser

import (
	"regexp"
	"zlc_spider/crawler/engine"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	//cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "user"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
	}
	return result
}
