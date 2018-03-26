package htmlfmt

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// write a newline followed by any prefix or indentation
func newline(w writer, prefix, indent string, depth int) error {
	if err := w.WriteByte('\n'); err != nil {
		return err
	}
	if _, err := w.WriteString(prefix); err != nil {
		return err
	}
	for i := 0; i < depth; i++ {
		if _, err := w.WriteString(indent); err != nil {
			return err
		}
	}
	return nil
}

// should emit a newline?
func openNewline(n *html.Node) bool {
	if spanElement[n.Data] {
		return false
	}
	switch n.DataAtom {
	case atom.P, atom.Pre:
		return false
	}

	/// if any is a inline element, then no newline
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if spanElement[c.Data] {
			return false
		}
	}
	c := n.FirstChild
	if c == nil {
		// render as <foo></foo>
		return false
	}
	if c.Type == html.TextNode {
		return len(strings.TrimSpace(c.Data)) == 0
	}

	return true
}

// inline elements
var spanElement = map[string]bool{
	"a":      true,
	"abbr":   true,
	"b":      true,
	"bdi":    true,
	"bdo":    true,
	"br":     true,
	"cite":   true,
	"code":   true,
	"data":   true,
	"del":    true,
	"dfn":    true,
	"em":     true,
	"i":      true,
	"kbd":    true,
	"mark":   true,
	"nobr":   true,
	"q":      true,
	"rp":     true,
	"rt":     true,
	"rtc":    true,
	"ruby":   true,
	"s":      true,
	"samp":   true,
	"small":  true,
	"span":   true,
	"strong": true,
	"sub":    true,
	"sup":    true,
	"time":   true,
	"tt":     true,
	"u":      true,
	"var":    true,
	"wbr":    true,
}
