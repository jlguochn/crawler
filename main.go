package main

import (
	"github.com/jlguochn/91crawler/91porn/parser"
	"github.com/jlguochn/91crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://91porn.com/v.php?category=rf&viewtype=basic&page=1",
		ParserFunc: parser.ParseVideoList,
	})
}
