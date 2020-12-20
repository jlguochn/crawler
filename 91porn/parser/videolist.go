package parser

import (
	"github.com/jlguochn/91crawler/engine"
	"regexp"
)

const videoListRe = `http://91porn.com/view_video.php.+category=rf`

func ParseVideoList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(videoListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	var visit = map[string]bool{}
	for _, m := range matches {
		//result.Items = append(result.Items, "name:"+string(m[2]))
		url := string(m[0])
		if visit[url] == false{
			result.Requests = append(result.Requests, engine.Request{
				Url:        url,
				ParserFunc: ParseVideo,
			})
			visit[url] = true
		}

	}
	return result
}
