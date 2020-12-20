package parser

import (
	"fmt"
	"github.com/jlguochn/91crawler/engine"
	"io/ioutil"
	"strings"
)

func DownloadMp4(contents []byte, filename string) engine.ParseResult {
	result := engine.ParseResult{}
	filename = strings.Replace(filename, " ", "", -1)
	filename = "./download/" + filename + ".mp4"
	fmt.Println("开始保存")
	ioutil.WriteFile(filename, contents, 0755)
	fmt.Println("保存完成：" + filename)
	return result
}





