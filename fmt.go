package htmlfmt

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Format reformats an input HTML document or error
func Format(src io.Reader, out io.Writer, prefix, indent string) error {
	doc, err := html.Parse(src)
	if err != nil {
		return err
	}
	if err := Render(out, doc, prefix, indent); err != nil {
		return err
	}
	out.Write([]byte{'\n'})
	return nil
}

// FormatFragment reformats an HTML fragment
func FormatFragment(src io.Reader, out io.Writer, prefix, indent string) error {
	root := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}
	nodes, err := html.ParseFragment(src, root)
	if err != nil {
		return err
	}

	for _, n := range nodes {
		if err := Render(out, n, prefix, indent); err != nil {
			return err
		}
	}
	out.Write([]byte{'\n'})
	return nil
}

// FormatBytes reformats an input HTML document as []byte
//
func FormatBytes(src []byte, prefix, indent string) []byte {
	reader := bytes.NewReader(src)
	writer := &bytes.Buffer{}
	if err := Format(reader, writer, prefix, indent); err != nil {
		return nil
	}
	return writer.Bytes()
}
