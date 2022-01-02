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

## License

The [BSD 3-Clause license][bsd], the same as the [Go language][golic]. Cascadia's license is [here][caslic].

[editor]: https://editorjs.io/
[golic]: http://golang.org/LICENSE
[bsd]: http://opensource.org/licenses/BSD-3-Clause
[caslic]: https://github.com/andybalholm/cascadia/blob/master/LICENSE
