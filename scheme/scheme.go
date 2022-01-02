package scheme

import (
	"github.com/PuerkitoBio/goquery"
)

type BlockHandler func(selection *goquery.Selection) *Block

type BlockData map[string]interface{}

type Response struct {
	Time    int64   `json:"time"`
	Blocks  []Block `json:"blocks"`
	Version string  `json:"version"`
}

type Block struct {
	Id   *string     `json:"id,omitempty"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
