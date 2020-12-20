package parser

import (
	"github.com/jlguochn/91crawler/engine"
	"regexp"
)

//const videoRe = `http://.+//mp43/.+e=[0-9]+`
//const titleRe = `<title>\s(.+)`
const videoRe = `<title>\s+(.+)|(http://.+//mp43/.+e=[0-9]+)`

func ParseVideo(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(videoRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	result.Items = append(result.Items, string(matches[0][1]))
	result.Requests = append(result.Requests, engine.Request{
		Url:        string(matches[1][0]),
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return DownloadMp4(bytes, string(matches[0][1]))
		},
	})

	return result
}
