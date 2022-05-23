package main

import (
"fmt"
"github.com/yuin/goldmark"
_ "github.com/yuin/goldmark"
"github.com/yuin/goldmark/ast"
"github.com/yuin/goldmark/extension"
"github.com/yuin/goldmark/parser"
"github.com/yuin/goldmark/text"
"io/ioutil"
)
func Resolver(){
	input := "test/data.md"

	source, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(),
			parser.WithBlockParsers(),
			parser.WithInlineParsers(),
		),
	)
	//var buf bytes.Buffer
	doc := md.Parser().Parse(text.NewReader(source))
	walk(doc)
}
func walk(node ast.Node) {
	for n := node.FirstChild(); n != nil; n = n.NextSibling() {
		fmt.Println(n.Type())
		walk(n)
	}
}
func  main()  {
	fmt.Println("fmbb")
	Resolver()
}
