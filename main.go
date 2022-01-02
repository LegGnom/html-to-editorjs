package html_to_editorjs

import (
	"github.com/LegGnom/html-to-editorjs/blocks"
	"github.com/LegGnom/html-to-editorjs/helpers"
	"github.com/LegGnom/html-to-editorjs/scheme"
	"github.com/PuerkitoBio/goquery"
	"time"
)

var blockHandlers = map[string]scheme.BlockHandler{}

func init() {
	RegistryBlock("p", blocks.Paragraph)
	RegistryBlock("h1", blocks.Header)
	RegistryBlock("h2", blocks.Header)
	RegistryBlock("h3", blocks.Header)
	RegistryBlock("h4", blocks.Header)
	RegistryBlock("h5", blocks.Header)
	RegistryBlock("ul", blocks.List)
	RegistryBlock("ol", blocks.List)
	RegistryBlock("img", blocks.Image)
	RegistryBlock("figure", blocks.Image)
	RegistryBlock("figure", blocks.Quote)
	RegistryBlock("blockquote", blocks.Quote)
	RegistryBlock("hr", blocks.Delimiter)
	RegistryBlock("table", blocks.Table)
}

func RegistryBlock(name string, handler scheme.BlockHandler) {
	blockHandlers[name] = handler
}

func Parse(payload string) string {
	nodes := helpers.GetContentSelection(payload)

	response := scheme.Response{
		Time: time.Now().Unix(),
	}

	nodes.Each(func(i int, selection *goquery.Selection) {
		if len(selection.Nodes) == 1 {
			tagName := helpers.GetTagName(selection)
			handler, ok := blockHandlers[tagName]

			if ok {
				if block := handler(selection); block != nil {
					response.Blocks = append(response.Blocks, *block)
				}
			}
		}

	})

	return helpers.ToString(response)
}
