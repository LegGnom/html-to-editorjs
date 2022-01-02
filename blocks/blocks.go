package blocks

import (
	"github.com/LegGnom/html-to-editorjs/helpers"
	"github.com/LegGnom/html-to-editorjs/scheme"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func Paragraph(selection *goquery.Selection) *scheme.Block {
	html, _ := selection.Unwrap().Html()

	return &scheme.Block{
		Type: "paragraph",
		Data: scheme.BlockData{
			"text": html,
		},
	}
}

func Header(selection *goquery.Selection) *scheme.Block {
	tagName := helpers.GetTagName(selection)
	html, _ := selection.Unwrap().Html()
	levelString := strings.Replace(tagName, "h", "", 1)

	if level, err := strconv.Atoi(levelString); err == nil {
		return &scheme.Block{
			Type: "header",
			Data: scheme.BlockData{
				"text":  html,
				"level": level,
			},
		}
	}

	return nil
}

func List(selection *goquery.Selection) *scheme.Block {
	tagName := helpers.GetTagName(selection)
	items := make([]string, 0)
	style := "unordered"

	if tagName == "ol" {
		style = "ordered"
	}

	selection.Find("li").Each(func(i int, li *goquery.Selection) {
		html, _ := li.Unwrap().Html()
		if html != "" {
			items = append(items, html)
		}
	})

	if len(items) > 0 {
		return &scheme.Block{
			Type: "list",
			Data: scheme.BlockData{
				"style": style,
				"items": items,
			},
		}
	}
	return nil
}

func Image(selection *goquery.Selection) *scheme.Block {
	tagName := helpers.GetTagName(selection)
	file := scheme.BlockData{}
	caption := ""

	if tagName == "img" {
		file = getImgFile(selection)
	}

	if tagName == "figure" {
		file = getImgFile(selection.Find("img:first-child"))
		caption, _ = selection.Find("figcaption").Unwrap().Html()
	}

	if file["url"] != nil {
		return &scheme.Block{
			Type: "image",
			Data: scheme.BlockData{
				"file":           file,
				"caption":        caption,
				"withBorder":     false,
				"stretched":      false,
				"withBackground": false,
			},
		}
	}

	return nil
}

func Code(selection *goquery.Selection) *scheme.Block {
	if len(selection.Children().Nodes) > 0 {
		if selection.Children().Nodes[0].Data == "code" {
			selection = selection.Children()
		}
	}
	html, _ := selection.Unwrap().Html()
	if html != "" {
		return &scheme.Block{
			Type: "code",
			Data: scheme.BlockData{
				"code": html,
			},
		}
	}

	return nil
}
func Quote(selection *goquery.Selection) *scheme.Block {
	tagName := helpers.GetTagName(selection)
	caption := ""
	html := ""
	if tagName == "figure" {
		caption, _ = selection.Find("figcaption").Clone().Unwrap().Html()

		if elem := selection.Find("blockquote"); elem != nil {
			html, _ = elem.Unwrap().Html()
		}
	}

	if tagName == "blockquote" {
		html, _ = selection.Unwrap().Html()
	}

	if html != "" {
		return &scheme.Block{
			Type: "quote",
			Data: scheme.BlockData{
				"text":      strings.TrimSpace(html),
				"caption":   strings.TrimSpace(caption),
				"alignment": "left",
			},
		}
	}

	return nil
}

func Delimiter(selection *goquery.Selection) *scheme.Block {
	return &scheme.Block{
		Type: "delimiter",
	}
}

func Table(selection *goquery.Selection) *scheme.Block {
	content := make([][]string, 0)

	selection.Find("tr").Each(func(i int, tr *goquery.Selection) {
		line := make([]string, 0)
		tr.Children().Each(func(i int, td *goquery.Selection) {
			elementName := helpers.GetTagName(td)
			if elementName == "td" || elementName == "th" {
				html, _ := td.Unwrap().Html()
				line = append(line, html)
			}
		})
		content = append(content, line)
	})

	if len(content) > 0 {
		return &scheme.Block{
			Type: "table",
			Data: scheme.BlockData{
				"withHeadings": false,
				"content":      content,
			},
		}
	}
	return nil
}
