package blocks

import (
	"github.com/LegGnom/html-to-editorjs/helpers"
	"github.com/LegGnom/html-to-editorjs/scheme"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestParagraph(t *testing.T) {
	nodes := helpers.GetContentSelection(`
		<p>test node <a href="#x">Link</a></p>
	`)

	nodes.Each(func(i int, selection *goquery.Selection) {
		block := Paragraph(selection)

		if block == nil {
			t.Fatal("Parse fail")
		}

		if block.Data.(scheme.BlockData)["text"] != `test node <a href="#x">Link</a>` {
			t.Fatal("Parse fail")
		}
	})
}

func TestCode(t *testing.T) {
	nodes := helpers.GetContentSelection(`
		<pre><code class="python">x = lambda  x,y :  x * y
print(x(3,4))</code></pre>
	`)

	nodes.Each(func(i int, selection *goquery.Selection) {
		block := Code(selection)

		if block == nil {
			t.Fatal("Parse fail")
		}

		if block.Data.(scheme.BlockData)["code"] != `x = lambda  x,y :  x * y
print(x(3,4))` {
			t.Fatal("Parse fail")
		}
	})
}

func TestCode2(t *testing.T) {
	nodes := helpers.GetContentSelection(`
		<pre>x = lambda  x,y :  x * y
print(x(3,4))</pre>
	`)

	nodes.Each(func(i int, selection *goquery.Selection) {
		block := Code(selection)

		if block == nil {
			t.Fatal("Parse fail")
		}

		if block.Data.(scheme.BlockData)["code"] != `x = lambda  x,y :  x * y
print(x(3,4))` {
			t.Fatal("Parse fail")
		}
	})
}

func TestQuote(t *testing.T) {
	nodes := helpers.GetContentSelection(`
		<figure>
			<blockquote cite="https://www.huxley.net/bnw/four.html">
        		<p>Words can be like X-rays, if you use them properly—they’ll go through anything. You read and you’re pierced.</p>
    		</blockquote>
    		<figcaption>—Aldous Huxley, <cite>Brave New World</cite></figcaption>
		</figure>
	`)

	nodes.Each(func(i int, selection *goquery.Selection) {
		block := Quote(selection)

		if block == nil {
			t.Fatal("Parse fail")
		}

		if block.Data.(scheme.BlockData)["text"] != `<p>Words can be like X-rays, if you use them properly—they’ll go through anything. You read and you’re pierced.</p>` {
			t.Fatal("Parse fail")
		}
	})
}
