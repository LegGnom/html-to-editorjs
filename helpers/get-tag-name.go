package helpers

import "github.com/PuerkitoBio/goquery"

func GetTagName(selection *goquery.Selection) string {
	if len(selection.Nodes) == 1 {
		return selection.Nodes[0].Data
	}
	return ""
}
