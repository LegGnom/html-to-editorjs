Simple converter from html to json for [editor.js][editor]

## Installation

```bash
go get github.com/LegGnom/html-to-editorjs
```

## Usage

```go
package main

import (
	"encoding/json"
    "fmt"
    "github.com/LegGnom/html-to-editorjs"
)

func main() {
	htmlString := `
        <h2>Editor.js</h2>
        <p>Hey. Meet the new Editor. On this page you can see it in action — try to edit this text.</p>
        <h3>Key features</h3>
        <ul>
            <li>It is a block-styled editor</li>
            <li>It returns clean data output in JSON</li>
            <li>Designed to be extendable and pluggable with a simple API</li>
        </ul>
	`

	html_to_editorjs.RegistryAll()
	j, _ := json.MarshalIndent(html_to_editorjs.Parse(html), "", "   ")
	fmt.Println(string(j))
}
```

It will generate the following output:

```json 
{
   "time": 1641162829,
   "blocks": [
      {
         "type": "header",
         "data": {
            "level": 2,
            "text": "Editor.js"
         }
      },
      {
         "type": "paragraph",
         "data": {
            "text": "Hey. Meet the new Editor. On this page you can see it in action — try to edit this text."
         }
      },
      {
         "type": "header",
         "data": {
            "level": 3,
            "text": "Key features"
         }
      },
      {
         "type": "list",
         "data": {
            "items": [
               "It is a block-styled editor",
               "It returns clean data output in JSON",
               "Designed to be extendable and pluggable with a simple API"
            ],
            "style": "unordered"
         }
      }
   ],
   "version": ""
}
```

### Adding new handlers
```go
package main

import (
	"fmt"
	"github.com/LegGnom/html-to-editorjs"
	"github.com/LegGnom/html-to-editorjs/scheme"
	"github.com/PuerkitoBio/goquery"
)


func main(t *testing.T) {
	html := `
		<pre>My code block</pre>
	`
	html_to_editorjs.RegistryBlock("pre", CodeHandler)
	fmt.Println(html_to_editorjs.Parse(html))
}

func CodeHandler(selection *goquery.Selection) *scheme.Block {
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
```

## Supported handlers
* **blocks.Paragraph** - tag `p`
* **blocks.Header** - tags `h1`, `h2`, `h3`, `h4`, `h5`
* **blocks.List** - tags `ul`, `ol`
* **blocks.Image** - tags `img`, `figure` 
* **blocks.Quote** - tags `figure`, `blockquote`
* **blocks.Delimiter** - tag `hr`
* * **blocks.Table** - tag `table`

If you need to add a handler to another tag, for example to div:

```go
html_to_editorjs.RegistryBlock("div", blocks.Image)
```

## License

The [BSD 3-Clause license][bsd], the same as the [Go language][golic].

[editor]: https://editorjs.io/
[golic]: http://golang.org/LICENSE
[bsd]: http://opensource.org/licenses/BSD-3-Clause
