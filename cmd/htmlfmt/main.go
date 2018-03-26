package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/client9/htmlfmt"
)

var (
	flagFragment = kingpin.Flag("fragment", "parse fragment").Short('f').Bool()
	flagPrefix   = kingpin.Flag("prefix", "prefix for each line").Short('p').String()
	flagIndent   = kingpin.Flag("indent", "indent market").Short('i').String()
)

func main() {
	var nodes []*html.Node
	var err error

	kingpin.Parse()
	if !*flagFragment {
		err := htmlfmt.Format(os.Stdin, os.Stdout, *flagPrefix, *flagIndent)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	root := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}

	log.Printf("parsing fragment")
	nodes, err = html.ParseFragment(os.Stdin, root)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodes {
		if err := htmlfmt.Render(os.Stdout, n, *flagPrefix, *flagIndent); err != nil {
			log.Fatal(err)
		}
	}
	os.Stdout.Write([]byte{'\n'})
}
