package parser

import (
	"github.com/jlguochn/91crawler/fetcher"
	"log"
	"testing"
)

func TestDownloadMp4(t *testing.T) {
	body, err := fetcher.Fetch("https://cdn.91p07.com//mp43/418390.mp4?st=d2GzoMFT5ES2rPf_m77KFw&e=1608515894")
	if err != nil {
		log.Printf("Fetcher:error %v", err)
	}else {
		DownloadMp4(body,"test")
	}
}