package parser

import (
	"regexp"
	"zlc_spider/crawler/engine"
	"zlc_spider/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Age = extractString(contents, ageRe)
	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(c []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
