package htmlfmt

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
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
