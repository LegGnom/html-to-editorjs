package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetContentSelection(payload string) *goquery.Selection {
	r := strings.NewReader(payload)
	query, _ := goquery.NewDocumentFromReader(r)
	return query.Find("body").Children()
}
